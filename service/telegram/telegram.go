package telegram

import (
	"strings"
	"sync"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nikoksr/onelog"
	nopadapter "github.com/nikoksr/onelog/adapter/nop"

	"github.com/nikoksr/notify/v2"
)

var _ notify.Service = (*Service)(nil)

const (
	// ModeHTML is one of the modes for sending messages.
	ModeHTML = telegram.ModeHTML
	// ModeMarkdown is the default mode for sending messages.
	ModeMarkdown = telegram.ModeMarkdown
)

func defaultMessageRenderer(conf *SendConfig) string {
	var builder strings.Builder

	builder.WriteString(conf.Subject)
	builder.WriteString("\n\n")
	builder.WriteString(conf.Message)

	return builder.String()
}

// Service is the telegram service. It is used to send messages to Telegram chats.
type Service struct {
	client *telegram.BotAPI

	name          string
	mu            sync.RWMutex
	logger        onelog.Logger
	renderMessage func(conf *SendConfig) string
	dryRun        bool
	continueOnErr bool

	// Telegram specific
	chatIDs   []int64
	parseMode string
}

func (s *Service) applyOptions(opts ...Option) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, opt := range opts {
		opt(s)
	}
}

// New creates a new telegram service. It returns an error if the telegram client could not be created.
func New(token string, opts ...Option) (*Service, error) {
	client, err := telegram.NewBotAPI(token)
	if err != nil {
		return nil, asNotifyError(err)
	}

	s := &Service{
		client:        client,
		name:          "telegram",
		logger:        nopadapter.NewAdapter(),
		renderMessage: defaultMessageRenderer,
		parseMode:     ModeMarkdown,
	}

	s.applyOptions(opts...)

	return s, nil
}

// Name returns the name of the service.
func (s *Service) Name() string {
	return s.name
}

// AddRecipients adds chat IDs that should receive messages.
func (s *Service) AddRecipients(chatIDs ...int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.chatIDs = append(s.chatIDs, chatIDs...)
	s.logger.Debug().Int("count", len(chatIDs)).Int("total", len(s.chatIDs)).Msg("Recipients added")
}
