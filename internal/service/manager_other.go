//go:build !windows

package service

import (
	"window-service-watcher/internal/domain"
)

type MockManager struct {
	// logCancel context.CancelFunc
	// logMutex  sync.Mutex
}

// GetServiceMetrics implements [domain.ServiceManager].
func (m *MockManager) GetServiceMetrics(serviceName string) (*domain.ServiceMetrics, error) {
	return &domain.ServiceMetrics{
		PID:        1234,
		CreateTime: 16251588000000,
		CPUUsage:   2.5,
		MemUsage:   104857600,
	}, nil
}

// GetServiceState implements [domain.ServiceManager].
func (m *MockManager) GetServiceState(serviceName string) (domain.Status, error) {
	return domain.RUNNING, nil
}

// StartService implements [domain.ServiceManager].
func (m *MockManager) StartService(serviceName string) error {
	return nil
}

// StopService implements [domain.ServiceManager].
func (m *MockManager) StopService(serviceName string) error {
	return nil
}

// StartLogWatcher implements [domain.ServiceManager].
// func (m *MockManager) StartLogWatcher(filePath string, onLog func(string), onError func(error)) {
// 	if filePath == "" {
// 		return
// 	}

// 	m.StopLogWatcher()
// 	ctx, cancel := context.WithCancel(context.Background())

// 	m.logMutex.Lock()
// 	m.logCancel = cancel
// 	m.logMutex.Unlock()

// 	go func(ctx context.Context) {
// 		t, err := tail.TailFile(filePath, tail.Config{
// 			Follow: true,
// 			ReOpen: true,
// 			Poll:   true, // window often use polling
// 		})
// 		if err != nil {
// 			onError(err)
// 			return
// 		}
// 		defer func() {
// 			t.Cleanup()
// 			t.Stop()
// 		}()

// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return
// 			case line, ok := <-t.Lines:
// 				if !ok {
// 					return
// 				}
// 				if line.Err != nil {
// 					onError(line.Err)
// 					continue
// 				}
// 				onLog(line.Text)
// 			}
// 		}
// 	}(ctx)
// }

// Connect implements [domain.ServiceManager].
func (m *MockManager) Connect() error {
	return nil
}

// Disconnect implements [domain.ServiceManager].
func (m *MockManager) Disconnect() error {
	return nil
}

// StopLogWatcher implements [domain.ServiceManager].
// func (m *MockManager) StopLogWatcher() {
// 	m.logMutex.Lock()
// 	defer m.logMutex.Unlock()

// 	if m.logCancel != nil {
// 		m.logCancel()
// 		m.logCancel = nil
// 	}
// }

func NewManager() domain.ServiceManager {
	return &MockManager{}
}
