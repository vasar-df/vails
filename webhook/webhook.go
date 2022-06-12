package webhook

import (
	"github.com/justtaldevelops/webhook"
)

// request is a request containing the target hook and the webhook message.
type request struct {
	hook    *webhook.Hook
	webhook webhook.Webhook
}

// webhooks is a channel for sending webhooks to the staff discord.
var webhooks = make(chan request)

// init initializes the goroutine that sends webhooks to Discord.
func init() {
	go func() {
		for w := range webhooks {
			_ = w.hook.Send(w.webhook)
		}
	}()
}

// Send pushes the given request into the webhook channel to be processed.
func Send(hook *webhook.Hook, message webhook.Webhook) {
	webhooks <- request{hook, message}
}
