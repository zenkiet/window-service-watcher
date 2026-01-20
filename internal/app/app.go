package app

import (
	"context"
	"window-service-watcher/internal/domain"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type App struct {
	manager domain.ServiceManager
}

func NewApp(mgr domain.ServiceManager) *App {
	return &App{
		manager: mgr,
	}
}

func (a *App) OnStartup(ctx context.Context, options application.ServiceOptions) error {
	err := a.manager.Connect()
	if err != nil {
		application.Get().Logger.Error("Failed to connect to service manager: " + err.Error())
		return err
	}
	return nil
}

func (a *App) Shutdown(ctx context.Context) {
	a.manager.Disconnect()
}

func (a *App) GetServiceStatus() (domain.ServiceStatus, error) {
	status, err := a.manager.CheckStatus()
	if err != nil {
		application.Get().Logger.Error("Check Status Error: " + err.Error())
		return domain.ServiceStatus{
			Status:    "Error",
			IsHealthy: false,
		}, err
	}
	return status, nil
}

func (a *App) WatchLogs(filePath string) {
	a.manager.StartLogWatcher(filePath, func(line string) {
		application.Get().Event.Emit("new-log", line)
	}, func(err error) {
		application.Get().Event.Emit("log-error", err.Error())
	})
}
