//go:build windows

package windows

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
func (w *WindowsManager) CheckStatus() (domain.ServiceStatus, error) {
	panic("unimplemented")
}

// Connect implements [domain.ServiceManager].
func (w *WindowsManager) Connect() error {
	panic("unimplemented")
}

// Disconnect implements [domain.ServiceManager].
func (w *WindowsManager) Disconnect() error {
	panic("unimplemented")
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

// RestartService implements [domain.ServiceManager].
func (w *WindowsManager) RestartService(serviceName string) error {
	err := w.StopService(serviceName)
	if err != nil {
		return fmt.Errorf("could not stop service: %w", err)
	}
	err = w.StartService(serviceName)
	if err != nil {
		return fmt.Errorf("could not start service: %w", err)
	}
	return nil
}

// StartLogWatcher implements [domain.ServiceManager].
func (w *WindowsManager) StartLogWatcher(filePath string, onLog func(string), onError func(error)) {
	panic("unimplemented")
}

// StartService implements [domain.ServiceManager].
func (w *WindowsManager) StartService(serviceName string) error {
	s, err := w.mgr.OpenService(serviceName)
	if err != nil {
		return fmt.Errorf("could not access service: %w", err)
	}
	defer s.Close()
	return s.Start()
}

// StopLogWatcher implements [domain.ServiceManager].
func (w *WindowsManager) StopLogWatcher() {
	panic("unimplemented")
}

// StopService implements [domain.ServiceManager].
func (w *WindowsManager) StopService(serviceName string) error {
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
