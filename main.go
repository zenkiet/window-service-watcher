package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"window-service-watcher/internal/app"
	"window-service-watcher/internal/service"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	srvMgr := service.NewManager()

	a := app.NewApp(srvMgr)

	// Create application with options
	err := wails.Run(&options.App{
		Title:       "Zen Service Watcher",
		Width:       320,
		Height:      80,
		Frameless:   true,
		AlwaysOnTop: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        a.Startup,
		OnShutdown:       a.Shutdown,
		Bind: []interface{}{
			a,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    true,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
