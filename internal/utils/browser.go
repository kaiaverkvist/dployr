package utils

import (
	"errors"
	"os/exec"
	"runtime"
)

// OpenUserBrowser uses platform specific logic to open a browser with a specific URL.
func OpenUserBrowser(url string) (err error) {
	// Platform switch statement.
	// This performs different techniques to open a URL in the user-default browser.
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return errors.New("unrecognized platform")
	}
}
