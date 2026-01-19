//go:build !windows

package service

import "window-service-watcher/internal/domain"

type MockManager struct{}

// CheckStatus implements [domain.ServiceManager].
func (m *MockManager) CheckStatus() (domain.ServiceStatus, error) {
	return domain.ServiceStatus{
		Name:      "MockService",
		Status:    "Running",
		IsHealthy: true,
	}, nil
}

// Connect implements [domain.ServiceManager].
func (m *MockManager) Connect() error {
	panic("unimplemented")
}

// Disconnect implements [domain.ServiceManager].
func (m *MockManager) Disconnect() error {
	panic("unimplemented")
}

// StartLogWatcher implements [domain.ServiceManager].
func (m *MockManager) StartLogWatcher(filePath string, onLog func(string), onError func(error)) {
	panic("unimplemented")
}

// StopLogWatcher implements [domain.ServiceManager].
func (m *MockManager) StopLogWatcher() {
	panic("unimplemented")
}

func NewManager() domain.ServiceManager {
	return &MockManager{}
}
