package services

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32              = windows.NewLazySystemDLL("user32.dll")
	enumWindowsProc     = user32.NewProc("EnumWindows")
	setWindowsHookEx    = user32.NewProc("SetWindowsHookExW")
	unhookWindowsHookEx = user32.NewProc("UnhookWindowsHookEx")
	callNextHookEx      = user32.NewProc("CallNextHookEx")
	getWindowText       = user32.NewProc("GetWindowTextW")
	getWindowTextLength = user32.NewProc("GetWindowTextLengthW")
	isWindow            = user32.NewProc("IsWindow")
	hookHandle          windows.Handle
	targetWindowName    string
	stopMonitoringChan  = make(chan bool)
)

const (
	WH_CBT        = 5
	HCBT_SETFOCUS = 9
)

type HookService struct{}

func (hs *HookService) GetWindowNames() ([]string, error) {
	var windowNames []string

	// Callback function for EnumWindows
	callback := syscall.NewCallback(func(hwnd syscall.Handle, lParam uintptr) uintptr {
		// Get the length of the window text
		textLen, _, _ := getWindowTextLength.Call(uintptr(hwnd))
		if textLen > 0 {
			// Allocate buffer and retrieve the text
			buf := make([]uint16, textLen+1)
			getWindowText.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&buf[0])), textLen+1)
			windowTitle := syscall.UTF16ToString(buf)
			windowNames = append(windowNames, windowTitle)
		}
		return 1 // Continue enumeration
	})

	// Call EnumWindows to enumerate top-level windows
	ret, _, err := enumWindowsProc.Call(callback, 0)
	if ret == 0 {
		return nil, fmt.Errorf("EnumWindows failed: %v", err)
	}

	return windowNames, nil
}

func (hs *HookService) StartWatchingFocus(windowName string) error {
	targetWindowName = strings.ToLower(windowName)

	// Define the hook procedure
	hookProc := syscall.NewCallback(func(code int, wParam, lParam uintptr) uintptr {
		if code == HCBT_SETFOCUS {
			// Check the window name
			bufLen, _, _ := getWindowTextLength.Call(wParam)
			if bufLen > 0 {
				buf := make([]uint16, bufLen+1)
				getWindowText.Call(wParam, uintptr(unsafe.Pointer(&buf[0])), bufLen+1)
				windowTitle := strings.ToLower(syscall.UTF16ToString(buf))
				if strings.Contains(windowTitle, targetWindowName) {
					fmt.Printf("Focus changed to target window: %s\n", windowTitle)
				}
			}
		}
		// Call the next hook in the chain
		ret, _, _ := callNextHookEx.Call(0, uintptr(code), wParam, lParam)
		return ret
	})

	// Set the hook
	h, _, err := setWindowsHookEx.Call(
		uintptr(WH_CBT),
		hookProc,
		0,
		uintptr(windows.GetCurrentThreadId()),
	)
	if h == 0 {
		return fmt.Errorf("failed to set hook: %v", err)
	}

	hookHandle = windows.Handle(h)
	fmt.Printf("Started watching focus for window: %s\n", windowName)
	go func() {
		<-stopMonitoringChan
		unhookWindowsHookEx.Call(uintptr(hookHandle))
		fmt.Printf("Stopped watching focus for window: %s\n", windowName)
	}()
	return nil
}

func (hs *HookService) StopWatchingFocus() {
	stopMonitoringChan <- true
}
