package statemachine

import (
	h "splitflap-backend/internal/handlers"
	"time"
)

func Initiate(app *h.Application) {
	s := StateHandler{App: app}

	stateHandlers := map[h.DisplayState]func() bool{
		h.Idle:    s.idleState,
		h.Spotify: s.initSpotifyStateHandler,
		h.Stocks:  s.initAvanzaHandler,
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

type StateHandler struct {
	App *h.Application
}

func (s *StateHandler) idleState() bool {
	return false
}
