package ntfy

import (
	"net/http"

	"github.com/nikoksr/onelog"
)

// Option is a function that can be used to configure the ntfy service.
type Option = func(*Service)

// WithClient sets the ntfy client. This is useful if you want to use a custom client.
func WithClient(client *http.Client) Option {
	return func(s *Service) {
		s.client = client
		s.logger.Debug().Msg("Ntfy client set")
	}
}

// WithLogger sets the logger. The default logger is a no-op logger.
func WithLogger(logger onelog.Logger) Option {
	return func(s *Service) {
		logger = logger.With("service", s.Name()) // Add service name to logger
		s.logger = logger
		s.logger.Debug().Msg("Logger set")
	}
}

// WithRecipients sets the topics that should receive messages. You can add more topics by calling AddRecipients.
func WithRecipients(topics ...string) Option {
	return func(s *Service) {
		s.topics = topics
		s.logger.Debug().Int("count", len(topics)).Int("total", len(s.topics)).Msg("Recipients set")
	}
}

// WithName sets the name of the service. The default name is "ntfy".
func WithName(name string) Option {
	return func(s *Service) {
		s.name = name
		s.logger.Debug().Str("name", name).Msg("Service name set")
	}
}

// WithMessageRenderer sets the message renderer. The default function will put the subject and message on separate lines.
//
// Example:
//
//	ntfy.WithMessageRenderer(func(conf *SendConfig) string {
//		var builder strings.Builder
//
//		builder.WriteString(conf.subject)
//		builder.WriteString("\n")
//		builder.WriteString(conf.message)
//
//		return builder.String()
//	})
func WithMessageRenderer(builder func(conf *SendConfig) string) Option {
	return func(s *Service) {
		s.renderMessage = builder
		s.logger.Debug().Msg("Message renderer set")
	}
}

// WithDryRun sets the dry run flag. If set to true, messages will not be sent.
func WithDryRun(dryRun bool) Option {
	return func(s *Service) {
		s.dryRun = dryRun
		s.logger.Debug().Bool("dry-run", dryRun).Msg("Dry run set")
	}
}

// WithContinueOnErr sets the continue on error flag. If set to true, the service will continue sending the message to
// the next recipient even if an error occurred.
func WithContinueOnErr(continueOnErr bool) Option {
	return func(s *Service) {
		s.continueOnErr = continueOnErr
		s.logger.Debug().Bool("continue-on-error", continueOnErr).Msg("Continue on error set")
	}
}

// WithAPIBaseURL sets the API base URL. The default is "https://ntfy.sh/".
func WithAPIBaseURL(url string) Option {
	return func(s *Service) {
		s.apiBaseURL = url
		s.logger.Debug().Str("url", url).Msg("API base URL set")
	}
}

// WithParseMode sets the parse mode for sending messages. The default is ModeText.
func WithParseMode(mode Mode) Option {
	return func(s *Service) {
		s.parseMode = mode
		s.logger.Debug().Str("mode", string(mode)).Msg("Parse mode set")
	}
}

// WithPriority sets the priority for sending messages. The default is PriorityDefault.
func WithPriority(priority Priority) Option {
	return func(s *Service) {
		s.priority = priority
		s.logger.Debug().Int("priority", int(priority)).Msg("Priority set")
	}
}

// WithTags sets the tags for sending messages. The default is no tags.
func WithTags(tags ...string) Option {
	return func(s *Service) {
		s.tags = tags
		s.logger.Debug().Int("count", len(tags)).Int("total", len(s.tags)).Msg("Tags set")
	}
}

// WithIcon sets the icon for sending messages. The default is "" (no icon).
func WithIcon(icon string) Option {
	return func(s *Service) {
		s.icon = icon
		s.logger.Debug().Str("icon", icon).Msg("Icon set")
	}
}

// WithDelay sets the delay for sending messages. The default is "" (no delay).
func WithDelay(delay string) Option {
	return func(s *Service) {
		s.delay = delay
		s.logger.Debug().Str("delay", delay).Msg("Delay set")
	}
}

// WithClickAction sets the click action for sending messages. The default is "" (no click action).
func WithClickAction(action string) Option {
	return func(s *Service) {
		s.clickAction = action
		s.logger.Debug().Str("clickAction", action).Msg("Click action set")
	}
}