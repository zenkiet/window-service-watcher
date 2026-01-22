package app

import (
	"context"
	"time"
	"window-service-watcher/internal/domain"
)

type ServiceWatcher struct {
	cfg        domain.Config
	mgr        domain.ServiceManager
	lastStatus map[string]domain.ServiceStatus
	updates    chan domain.ServiceStatus
}

func NewServiceWatcher(cfg domain.Config, mgr domain.ServiceManager) *ServiceWatcher {
	return &ServiceWatcher{
		cfg:        cfg,
		mgr:        mgr,
		lastStatus: make(map[string]domain.ServiceStatus),
		updates:    make(chan domain.ServiceStatus, 10),
	}
}

func (sw *ServiceWatcher) Updates() <-chan domain.ServiceStatus {
	return sw.updates
}

func (sw *ServiceWatcher) Start(ctx context.Context) error {
	sw.checkAll()

	ticket := time.NewTicker(2 * time.Second)
	defer ticket.Stop()

	for {
		select {
		case <-ctx.Done():
			close(sw.updates)
			return nil
		case <-ticket.C:
			sw.checkAll()
		}
	}
}

func (sw *ServiceWatcher) checkAll() {
	for _, cfg := range sw.cfg.Services {
		sw.CheckServices(cfg)
	}
}

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

	oldStatus, exists := sw.lastStatus[cfg.ID]
	hasChanged := !exists || oldStatus.Status != status.Status || oldStatus.IsHealthy != status.IsHealthy

	if hasChanged {
		select {
		case sw.updates <- status:
		default:
		}
		sw.lastStatus[cfg.ID] = status
	}
}
