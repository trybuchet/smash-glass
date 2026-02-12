package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"SmashGlass/services"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	application.RegisterEvent[string]("time")
}

func main() {

	mode := "overlay"
	serverMode := "false"
	windowMode := "false"

	_ = godotenv.Load()

	serverMode = os.Getenv("SERVER_MODE")
	windowMode = os.Getenv("WINDOW_MODE")
	serverPort := os.Getenv("SERVER_PORT")
	discordClientId := os.Getenv("DISCORD_CLIENT_ID")

	serverService := services.NewServerService()
	configService := services.NewConfigService()
	hotkeyService := services.NewHotkeyService()
	discordService := services.NewDiscordService(discordClientId)

	if serverMode == "true" {
		port, err := strconv.Atoi(serverPort)
		if err != nil {
			log.Println("Invalid SERVER_PORT:", err)
		} else {
			go serverService.StartServer(port)
		}
	}

	windowService := services.NewWindowService((windowMode == "true"))

	var shutdownOnce sync.Once

	shutdown := func(reason string) {
		shutdownOnce.Do(func() {
			fmt.Println("[SHUTDOWN]", reason)

			hotkeyService.UnregisterAll()

			if serverMode == "true" {
				serverService.StopServer()
			}
		})
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		<-ctx.Done()
		shutdown("os signal")
	}()

	app := application.New(application.Options{
		Name:        "Smash Glass",
		Description: "Smash Soda Overlay",
		Services: []application.Service{
			application.NewService(windowService),
			application.NewService(hotkeyService),
			application.NewService(&services.HookService{}),
			application.NewService(services.NewPluginService()),
			application.NewService(services.NewStyleService()),
			application.NewService(configService),
			application.NewService(discordService),
			application.NewService(services.NewFileService()),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
		ShouldQuit: func() bool {
			shutdown("wails ShouldQuit")
			return true
		},
		OnShutdown: func() {
			shutdown("wails OnShutdown")
		},
	})

	width, height := windowService.GetResolution()

	window := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title: "Smash Glass",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundType:             application.BackgroundTypeTransparent,
		BackgroundColour:           application.NewRGBA(0, 0, 0, 0),
		URL:                        "/?mode=" + mode,
		Zoom:                       1,
		Width:                      width,
		Height:                     height,
		ZoomControlEnabled:         true,
		DefaultContextMenuDisabled: false,
		AlwaysOnTop:                !(windowMode == "true"),
		IgnoreMouseEvents:          !(windowMode == "true"),
		DevToolsEnabled:            true,
	})

	services.WailsApp = app
	services.WailsWindow = window

	err := app.Run()
	if err != nil {
		shutdown("app.Run error")
		log.Println(err)
	}
}
