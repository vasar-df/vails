package worlds

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/mcdb"
	"github.com/df-mc/goleveldb/leveldb/opt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

// Manager manages multiple worlds, dragonfly does not have multi-world management itself,
// so we must implement it ourselves.
type Manager struct {
	s *server.Server

	folderPath string

	log *logrus.Logger

	worldsMu sync.RWMutex
	worlds   map[string]*world.World
}

// New ...
func New(server *server.Server, folderPath string, log *logrus.Logger) *Manager {
	_ = os.Mkdir(folderPath, 0644)
	defaultWorld := server.World()
	return &Manager{
		s:          server,
		folderPath: folderPath,
		log:        log,
		worlds: map[string]*world.World{
			defaultWorld.Name(): defaultWorld,
		},
	}
}

// DefaultWorld ...
func (m *Manager) DefaultWorld() *world.World {
	return m.s.World()
}

// Worlds ...
func (m *Manager) Worlds() []*world.World {
	m.worldsMu.RLock()
	worlds := make([]*world.World, 0, len(m.worlds))
	for _, w := range m.worlds {
		worlds = append(worlds, w)
	}
	m.worldsMu.RUnlock()
	return worlds
}

// AssertWorld ...
func (m *Manager) AssertWorld(name string) *world.World {
	m.worldsMu.RLock()
	defer m.worldsMu.RUnlock()
	if w, ok := m.worlds[name]; ok {
		return w
	}
	panic(fmt.Errorf("expected world %v, but is not loaded", name))
}

// World ...
func (m *Manager) World(name string) (*world.World, bool) {
	m.worldsMu.RLock()
	w, ok := m.worlds[name]
	m.worldsMu.RUnlock()
	return w, ok
}

// LoadWorld ...
func (m *Manager) LoadWorld(folderName, worldName string) error {
	if _, ok := m.World(worldName); ok {
		return fmt.Errorf("world is already loaded")
	}

	log := m.log.WithField("dimension", "overworld")
	log.Debugf("Loading world...")
	p, err := mcdb.New(m.folderPath+"/"+folderName, opt.DefaultCompression)
	if err != nil {
		return fmt.Errorf("error loading world: %v", err)
	}

	w := world.Config{
		Dim:      world.Overworld,
		Log:      m.log,
		ReadOnly: true,
		Provider: p,
	}.New()

	w.SetTickRange(0)
	w.SetTime(6000)
	w.StopTime()

	w.StopWeatherCycle()
	w.SetDefaultGameMode(world.GameModeSurvival)
	w.Handle(&Handler{})

	m.worldsMu.Lock()
	m.worlds[worldName] = w
	m.worldsMu.Unlock()

	log.Debugf(`Loaded world "%v".`, w.Name())
	return nil
}

// UnloadWorld ...
func (m *Manager) UnloadWorld(w *world.World) error {
	if w == m.DefaultWorld() {
		return fmt.Errorf("the default world cannot be unloaded")
	}

	if _, ok := m.World(w.Name()); !ok {
		return fmt.Errorf("world isn't loaded")
	}

	m.log.Debugf("Unloading world '%v'\n", w.Name())
	for _, p := range m.s.Players() {
		if p.World() == w {
			m.DefaultWorld().AddEntity(p)
			p.Teleport(m.DefaultWorld().Spawn().Vec3Middle())
		}
	}

	m.worldsMu.Lock()
	delete(m.worlds, w.Name())
	m.worldsMu.Unlock()

	if err := w.Close(); err != nil {
		return fmt.Errorf("error closing world: %v", err)
	}
	m.log.Debugf("Unloaded world '%v'\n", w.Name())
	return nil
}

// Close ...
func (m *Manager) Close() error {
	m.worldsMu.Lock()
	for _, w := range m.worlds {
		// Let dragonfly close this.
		if w == m.DefaultWorld() {
			continue
		}

		m.log.Debugf("Closing world '%v'\n", w.Name())
		if err := w.Close(); err != nil {
			return err
		}
	}
	m.worlds = nil
	m.worldsMu.Unlock()
	return nil
}
