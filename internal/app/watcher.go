package app

import (
	"context"
	"sync"
	"time"
	"window-service-watcher/internal/domain"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type ServiceWatcher struct {
	cfg        domain.Config
	mgr        domain.ServiceManager
	lastStatus map[string]domain.ServiceStatus
	mu         sync.RWMutex
}

func NewServiceWatcher(cfg domain.Config, mgr domain.ServiceManager) *ServiceWatcher {
	return &ServiceWatcher{
		cfg:        cfg,
		mgr:        mgr,
		lastStatus: make(map[string]domain.ServiceStatus),
	}
}

func (sw *ServiceWatcher) Start(ctx context.Context) error {
	ticket := time.NewTicker(2 * time.Second)
	defer ticket.Stop()

	for _, cfg := range sw.cfg.Services {
		sw.CheckServices(cfg)
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticket.C:
			for _, cfg := range sw.cfg.Services {
				sw.CheckServices(cfg)
			}
		}
	}
}

// func (sw *ServiceWatcher) Stop() error {}

// func (sw *ServiceWatcher) watchLogs(svc domain.ServiceConfig) error {}

func (sw *ServiceWatcher) CheckServices(cfg domain.ServiceConfig) {
	statusStr, isHealthy, err := sw.mgr.GetServiceState(cfg.ServiceName)

	status := domain.ServiceStatus{
		ID:        cfg.ID,
		Name:      cfg.Name,
		Status:    statusStr,
		IsHealthy: isHealthy,
	}

	if err != nil {
		status.Status = "Error"
		status.IsHealthy = false
	}

	sw.mu.Lock()
	oldStatus, exists := sw.lastStatus[cfg.ID]

	hasChanged := !exists || oldStatus.Status != status.Status || oldStatus.IsHealthy != status.IsHealthy

	if hasChanged {
		sw.lastStatus[cfg.ID] = status
		application.Get().Event.Emit("service-update-"+cfg.ID, status)
		application.Get().Logger.Info("Service status changed", "service", cfg.Name, "status", status.Status, "healthy", status.IsHealthy)
	}
	sw.mu.Unlock()
}
