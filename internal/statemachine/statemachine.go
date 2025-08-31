package statemachine

import (
	h "splitflap-backend/internal/handlers"
	"time"
)

type StateHandler struct {
	App *h.Application
}

var s = StateHandler{}

func Initiate(app *h.Application) {
	s = StateHandler{App: app}

	stateHandlers := map[h.DisplayState]func() bool{
		h.Idle:   s.idleState,
		h.Stocks: s.initAvanzaHandler,
	}

	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		if !app.IsIdleState() {
			continue
		}

		for _, stateHandler := range stateHandlers {
			swappedState := stateHandler()
			if swappedState {
				break
			}
		}
	}
}

func (s *StateHandler) idleState() bool {
	return false
}
