package app

import (
	"context"
	"fmt"

	"window-service-watcher/internal/domain"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type App struct {
	cfg     domain.Config
	mgr     domain.ServiceManager
	svcMap  map[string]domain.ServiceConfig
	watcher *ServiceWatcher

	cancelWatch context.CancelFunc
}

func NewApp(cfg domain.Config, mgr domain.ServiceManager) *App {
	sMap := make(map[string]domain.ServiceConfig)
	for _, svc := range cfg.Services {
		sMap[svc.ID] = svc
	}

	return &App{
		cfg:     cfg,
		mgr:     mgr,
		svcMap:  sMap,
		watcher: NewServiceWatcher(cfg, mgr),
	}
}

func (a *App) OnStartup(ctx context.Context, options application.ServiceOptions) error {
	if err := a.mgr.Connect(); err != nil {
		application.Get().Logger.Error("Service Manager Connect Error: " + err.Error())
		return err
	}

	watchCtx, cancel := context.WithCancel(context.Background())
	a.cancelWatch = cancel

	go a.watcher.Start(watchCtx)

	go func() {
		for static := range a.watcher.Updates() {
			application.Get().Event.Emit("service-update-"+static.ID, static)
			application.Get().Logger.Info("Service Update: " + static.ID + " Status: " + static.Status)
		}
	}()

	return nil
}

func (a *App) Shutdown(ctx context.Context) {
	if a.cancelWatch != nil {
		a.cancelWatch()
	}
	a.mgr.Disconnect()
}

func (a *App) GetConfig() domain.Config {
	return a.cfg
}

func (a *App) StartService(id string) error {
	cfg, ok := a.svcMap[id]
	application.Get().Logger.Info("Starting Service: " + cfg.ServiceName)
	if !ok {
		return fmt.Errorf("service config not found for ID: %s", id)
	}

	err := a.mgr.StartService(cfg.ServiceName)
	if err != nil {
		application.Get().Logger.Error("Start Service Error: " + err.Error())
		return err
	}
	return nil
}

func (a *App) StopService(id string) error {
	cfg, ok := a.svcMap[id]
	if !ok {
		return fmt.Errorf("service config not found for ID: %s", id)
	}

	err := a.mgr.StopService(cfg.ServiceName)
	if err != nil {
		application.Get().Logger.Error("Stop Service Error: " + err.Error())
		return err
	}
	return nil
}

// func (a *App) GetServiceStatus() (domain.ServiceStatus, error) {
// 	status, err := a.manager.CheckStatus()
// 	if err != nil {
// 		application.Get().Logger.Error("Check Status Error: " + err.Error())
// 		return domain.ServiceStatus{
// 			Status:    "Error",
// 			IsHealthy: false,
// 		}, err
// 	}
// 	return status, nil
// }

// func (a *App) WatchLogs(filePath string) {
// 	a.manager.StartLogWatcher(filePath, func(line string) {
// 		application.Get().Event.Emit("new-log", line)
// 	}, func(err error) {
// 		application.Get().Event.Emit("log-error", err.Error())
// 	})
// }
