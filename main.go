package main

import (
	"embed"
	"log"
	"log/slog"
	"os"
	"time"
	"window-service-watcher/internal/app"
	"window-service-watcher/internal/config"
	"window-service-watcher/internal/service"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

var trayIcon []byte

func init() {
	// Register a custom event whose associated data type is string.
	// This is not required, but the binding generator will pick up registered events
	// and provide a strongly typed JS/TS API for them.
	application.RegisterEvent[string]("time")
}

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {
	logFile, _ := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(logFile)

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	srvMgr := service.NewManager()
	myApp := app.NewApp(*cfg, srvMgr)

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "Zen Service Manager",
		Description: "A simple service manager built with Wails",
		Services: []application.Service{
			application.NewService(myApp),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ActivationPolicy: application.ActivationPolicyAccessory,
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
		Windows:  application.WindowsOptions{},
		Logger:   slog.Default(),
		LogLevel: slog.LevelDebug,
	})

	systemTray := app.SystemTray.New()

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	mainWindow := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:          "Zen Service Watcher",
		DisableResize:  false,
		Hidden:         false,
		BackgroundType: application.BackgroundTypeTranslucent,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		Windows: application.WindowsWindow{
			Theme:                   application.Dark,
			BackdropType:            application.Mica,
			WindowMaskDraggable:     true,
			ResizeDebounceMS:        200,
			WindowDidMoveDebounceMS: 200,
		},
		Linux:            application.LinuxWindow{},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})

	if trayIcon != nil {
		systemTray.SetIcon(trayIcon)
	}
	trayMenu := app.NewMenu()
	trayMenu.Add("Open Dashboard").OnClick(func(_ *application.Context) {
		mainWindow.Show()
		mainWindow.Focus()
	})
	trayMenu.AddSeparator()
	trayMenu.Add("Exit").OnClick(func(_ *application.Context) {
		app.Quit()
	})

	systemTray.SetMenu(trayMenu)
	systemTray.OnClick(func() {
		if mainWindow.IsVisible() {
			mainWindow.Hide()
		} else {
			mainWindow.Show()
			mainWindow.Focus()
		}
	})

	mainWindow.RegisterHook(events.Windows.WindowClosing, func(e *application.WindowEvent) {
		e.Cancel()
		mainWindow.Hide()

		app.Dialog.Info().SetTitle("Still Running").SetMessage("Zen Watcher is running in background").Show()
	})

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.Event.Emit("time", now)
			time.Sleep(time.Second)
		}
	}()

	// Run the application. This blocks until the application has been exited.
	errApp := app.Run()
	// If an error occurred while running the application, log it and exit.
	if errApp != nil {
		log.Fatal(errApp)
	}
}
