package splitit

import (
	"context"
	"sync"
	"time"
)

type sessionIDHandler struct {
	apiClient            *APIClient
	mu                   sync.Mutex
	isManual             bool
	isStale              bool
	sessionID            string
	err                  error
	lastSessionIDRefresh time.Time
}

func newSessionIDHandler(apiClient *APIClient) *sessionIDHandler {
	return &sessionIDHandler{
		isStale:   true,
		apiClient: apiClient,
	}
}

func (sih *sessionIDHandler) SetSessionID(sessionID string) {
	// Manual session ID management for (for example) flexfields
	sih.mu.Lock()
	defer sih.mu.Unlock()
	sih.isManual = true
	sih.sessionID = sessionID
}

func (sih *sessionIDHandler) GetSessionID(ctx context.Context) (sessionID string, err error) {
	if ctx.Value(noSessionIdKey{}) != nil {
		return "", nil
	}
	sih.mu.Lock()
	defer sih.mu.Unlock()
	if !sih.isManual && sih.isStale {
		lr, _, err := sih.apiClient.LoginApi.LoginPost(ctx, LoginRequest{
			UserName: sih.apiClient.cfg.username,
			Password: sih.apiClient.cfg.password,
		})
		sih.sessionID = lr.SessionId
		sih.err = err

		sih.isStale = false
		sih.lastSessionIDRefresh = time.Now()
	}
	return sih.sessionID, sih.err
}

func (sih *sessionIDHandler) InvalidateSessionID() {
	sih.mu.Lock()
	defer sih.mu.Unlock()
	sih.isStale = true
}
