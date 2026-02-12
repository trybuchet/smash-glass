package services

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

var (
	WailsApp       *application.App           // Handle to the Wails application
	WailsWindow    *application.WebviewWindow // Handle to the Wails webview window
	StudioWindow   *application.WebviewWindow // Handle to the Stream Studio window
	HeadlessWindow *application.WebviewWindow // Handle to headless overlay window
)
