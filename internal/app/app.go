package app

import (
	"context"

	"window-service-watcher/internal/domain"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type App struct {
	config  domain.Config
	manager domain.ServiceManager
	watcher *ServiceWatcher

	cancelWatch context.CancelFunc
}

func NewApp(cfg domain.Config, mgr domain.ServiceManager) *App {
	return &App{
		config:  cfg,
		manager: mgr,
		watcher: NewServiceWatcher(cfg, mgr),
	}
}

func (a *App) OnStartup(ctx context.Context, options application.ServiceOptions) error {
	if err := a.manager.Connect(); err != nil {
		application.Get().Logger.Error("Service Manager Connect Error: " + err.Error())
		return err
	}

	watchCtx, cancel := context.WithCancel(context.Background())
	a.cancelWatch = cancel

	go func() {
		if err := a.watcher.Start(watchCtx); err != nil {
			application.Get().Logger.Error("Service Watcher Start Error: " + err.Error())
		}
	}()

	return nil
}

func (a *App) Shutdown(ctx context.Context) {
	if a.cancelWatch != nil {
		a.cancelWatch()
	}
	a.manager.Disconnect()
}

func (a *App) GetConfig() domain.Config {
	return a.config
}

func (a *App) StartService(id string) error {
	cgf, ok := a.findConfigByID(id)
	if !ok {
		return context.DeadlineExceeded
	}

	err := a.manager.StartService(cgf.ServiceName)
	if err != nil {
		application.Get().Logger.Error("Start Service Error: " + err.Error())
		return err
	}
	return nil
}

func (a *App) StopService(id string) error {
	cgf, ok := a.findConfigByID(id)
	if !ok {
		return context.DeadlineExceeded
	}

	err := a.manager.StopService(cgf.ServiceName)
	if err != nil {
		application.Get().Logger.Error("Stop Service Error: " + err.Error())
		return err
	}
	return nil
}

func (a *App) RestartService(id string) error {
	cgf, ok := a.findConfigByID(id)
	if !ok {
		return context.DeadlineExceeded
	}

	err := a.manager.RestartService(cgf.ServiceName)
	if err != nil {
		application.Get().Logger.Error("Restart Service Error: " + err.Error())
		return err
	}
	return nil
}

func (a *App) findConfigByID(id string) (domain.ServiceConfig, bool) {
	for _, cfg := range a.config.Services {
		if cfg.ID == id {
			return cfg, true
		}
	}
	return domain.ServiceConfig{}, false
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
