package app

import (
	"context"
	"window-service-watcher/internal/domain"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx     context.Context
	manager domain.ServiceManager
}

func NewApp(mgr domain.ServiceManager) *App {
	return &App{
		manager: mgr,
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	err := a.manager.Connect()
	if err != nil {
		wailsRuntime.LogError(a.ctx, "Failed to connect to service manager: "+err.Error())
	}
}

func (a *App) Shutdown(ctx context.Context) {
	a.manager.Disconnect()
}

func (a *App) GetServiceStatus() (domain.ServiceStatus, error) {
	status, err := a.manager.CheckStatus()
	if err != nil {
		wailsRuntime.LogError(a.ctx, "Check Status Error: "+err.Error())
		return domain.ServiceStatus{
			Status:    "Error",
			IsHealthy: false,
		}, err
	}
	return status, nil
}

func (a *App) WatchLogs(filePath string) {
	a.manager.StartLogWatcher(filePath, func(line string) {
		wailsRuntime.EventsEmit(a.ctx, "new-log", line)
	}, func(err error) {
		wailsRuntime.EventsEmit(a.ctx, "log-error", err.Error())
	})
}
