//go:build windows

package service

import (
	"fmt"
	"window-service-watcher/internal/domain"

	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/mgr"
)

type WindowsManager struct {
	mgr *mgr.Mgr
}

// CheckStatus implements [domain.ServiceManager].
// func (w *WindowsManager) CheckStatus() (domain.ServiceStatus, error) {
// }

// Connect implements [domain.ServiceManager].
func (w *WindowsManager) Connect() error {
	if w.mgr != nil {
		return nil
	}

	m, err := mgr.Connect()
	if err != nil {
		return fmt.Errorf("failed to connect to service manager: %w", err)
	}
	w.mgr = m
	return nil
}

// Disconnect implements [domain.ServiceManager].
func (w *WindowsManager) Disconnect() error {
	if w.mgr == nil {
		return nil
	}

	err := w.mgr.Disconnect()
	w.mgr = nil
	return err
}

// GetServiceState implements [domain.ServiceManager].
func (w *WindowsManager) GetServiceState(serviceName string) (string, bool, error) {
	if w.mgr == nil {
		return "Disconnected", false, fmt.Errorf("service manager not connected")
	}

	s, err := w.mgr.OpenService(serviceName)
	if err != nil {
		return "Not Found", false, fmt.Errorf("service not found: %w", err)
	}
	defer s.Close()

	status, err := s.Query()
	if err != nil {
		return "Unknown", false, fmt.Errorf("error querying service status: %w", err)
	}

	state := "Unknown"
	isHealthy := false

	switch status.State {
	case windows.SERVICE_STOPPED:
		state = "Stopped"
	case windows.SERVICE_START_PENDING:
		state = "Starting"
	case windows.SERVICE_STOP_PENDING:
		state = "Stopping"
	case windows.SERVICE_RUNNING:
		state = "Running"
		isHealthy = true
	case windows.SERVICE_CONTINUE_PENDING:
		state = "Resuming"
	case windows.SERVICE_PAUSE_PENDING:
		state = "Pausing"
	case windows.SERVICE_PAUSED:
		state = "Paused"
	default:
		state = "Unknown"
	}

	return state, isHealthy, nil
}

// StartLogWatcher implements [domain.ServiceManager].
// func (w *WindowsManager) StartLogWatcher(filePath string, onLog func(string), onError func(error)) {
// }

// StartService implements [domain.ServiceManager].
func (w *WindowsManager) StartService(serviceName string) error {
	if w.mgr == nil {
		return fmt.Errorf("service manager not connected")
	}

	s, err := w.mgr.OpenService(serviceName)
	if err != nil {
		return fmt.Errorf("could not access service: %w", err)
	}
	defer s.Close()
	return s.Start()
}

// StopLogWatcher implements [domain.ServiceManager].
// func (w *WindowsManager) StopLogWatcher() {
// }

// StopService implements [domain.ServiceManager].
func (w *WindowsManager) StopService(serviceName string) error {
	if w.mgr == nil {
		return fmt.Errorf("service manager not connected")
	}

	s, err := w.mgr.OpenService(serviceName)
	if err != nil {
		return fmt.Errorf("could not access service: %w", err)
	}
	defer s.Close()
	status, err := s.Control(windows.SERVICE_CONTROL_STOP)
	if err != nil {
		return fmt.Errorf("could not stop service: %w", err)
	}
	if status.State != windows.SERVICE_STOPPED {
		return fmt.Errorf("service did not stop successfully, current state: %d", status.State)
	}
	return nil
}

func NewManager() domain.ServiceManager {
	return &WindowsManager{}
}
