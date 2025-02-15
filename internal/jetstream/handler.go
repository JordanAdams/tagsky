package jetstream

import (
	"encoding/json"
	"fmt"

	"github.com/bluesky-social/jetstream/pkg/models"
)

type MessageHandler func(msg []byte) error

type EventHandler func(msg models.Event) error

type CommitHandler func(msg models.Event) error

type Handler struct {
	messageHandlers []MessageHandler
	eventHandlers   []EventHandler
	commitHandlers  []EventHandler
}

func NewHandler() (handler *Handler) {
	return &Handler{
		messageHandlers: []MessageHandler{},
		eventHandlers:   []EventHandler{},
		commitHandlers:  []EventHandler{},
	}
}

func (h *Handler) HandleMessage(f MessageHandler) {
	h.messageHandlers = append(h.messageHandlers, f)
}

func (h *Handler) HandleEvent(f EventHandler) {
	h.eventHandlers = append(h.eventHandlers, f)
}

func (h *Handler) HandleCommit(f EventHandler) {
	h.commitHandlers = append(h.commitHandlers, f)
}

func (h *Handler) Handle(msg []byte) error {
	for _, mh := range h.messageHandlers {
		err := mh(msg)
		if err != nil {
			return fmt.Errorf("failed to handle message: %w", err)
		}
	}

	var event models.Event
	err := json.Unmarshal(msg, &event)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal message: %w", err)
	}

	for _, eh := range h.eventHandlers {
		err := eh(event)
		if err != nil {
			return fmt.Errorf("failed to handle event: %w", err)
		}
	}

	if event.Commit != nil {
		for _, ch := range h.commitHandlers {
			err := ch(event)
			if err != nil {
				return fmt.Errorf("failed to handle commit: %w", err)
			}
		}
	}

	return nil
}
