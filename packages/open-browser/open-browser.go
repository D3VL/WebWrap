package browser

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
    "D3VL/WebWrap/packages/logging"
    "os"
    "syscall"   
)


// copilot wrote alot of these paths, expect them to be wrong -- chrome and msedge are the only ones that have been tested.
var browsers = map[string]map[string]string{
    "chrome-x86": {
        "linux": "/usr/bin/google-chrome",
        "windows": "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe",
        "darwin": "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
        "flags": "--app={{url}}",
    },
    "chrome-x64": {
        "linux": "/usr/bin/google-chrome",
        "windows": "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
        "darwin": "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
        "flags": "--app={{url}}",
    },
    "chromium-x86": {
        "linux": "/usr/bin/chromium",
        "windows": "C:\\Program Files (x86)\\Chromium\\Application\\chromium.exe",
        "darwin": "/Applications/Chromium.app/Contents/MacOS/Chromium",
        "flags": "--app={{url}}",
    },
    "chromium-x64": {
        "linux": "/usr/bin/chromium",
        "windows": "C:\\Program Files\\Chromium\\Application\\chromium.exe",
        "darwin": "/Applications/Chromium.app/Contents/MacOS/Chromium",
        "flags": "--app={{url}}",
    },
    "msedge-x86": {
        "linux": "/usr/bin/msedge",
        "windows": "C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe",
        "darwin": "/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
        "flags": "--app={{url}}",
    },
    "msedge-x64": {
        "linux": "/usr/bin/msedge",
        "windows": "C:\\Program Files\\Microsoft\\Edge\\Application\\msedge.exe",
        "darwin": "/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
        "flags": "--app={{url}}",
    },
    "msedge": {
        "linux": "/usr/bin/microsoft-edge",
        "windows": "C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe",
        "darwin": "/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
        "flags": "--app={{url}}",
    },
    "brave-x86": {
        "linux": "/usr/bin/brave",
        "windows": "C:\\Program Files (x86)\\BraveSoftware\\Brave-Browser\\Application\\brave.exe",
        "darwin": "/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
        "flags": "--app={{url}}",
    },
    "brave-x64": {
        "linux": "/usr/bin/brave",
        "windows": "C:\\Program Files\\BraveSoftware\\Brave-Browser\\Application\\brave.exe",
        "darwin": "/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
        "flags": "--app={{url}}",
    },
    "vivaldi-x86": {
        "linux": "/usr/bin/vivaldi",
        "windows": "C:\\Program Files (x86)\\Vivaldi\\Application\\vivaldi.exe",
        "darwin": "/Applications/Vivaldi.app/Contents/MacOS/Vivaldi",
        "flags": "--app={{url}}",
    },
    "vivaldi-x64": {
        "linux": "/usr/bin/vivaldi",
        "windows": "C:\\Program Files\\Vivaldi\\Application\\vivaldi.exe",
        "darwin": "/Applications/Vivaldi.app/Contents/MacOS/Vivaldi",
        "flags": "--app={{url}}",
    },
}

var commandRun = map[string][]string{
    "linux": {"/bin/sh", "-c"},
    "windows": {},// {"cmd", "/c", "start", "/W", ""},
    "darwin": {"/bin/sh", "-c"},
}

func runCommand(cmd *exec.Cmd)  (int, error) {
    err := cmd.Run()
    if err != nil {
        if e2, ok := err.(*exec.ExitError); ok {
            if s, ok := e2.Sys().(syscall.WaitStatus); ok {
                return int(s.ExitCode), nil
            }
        }
        return 0, err
    } else {
        return 0, nil
    }
}

func Open(url string) {    
    var err error
    var didOpen bool

    // chromeBrowsers := []string{"google-chrome", "chrome.exe", "chromium", "msedge", "brave", "vivaldi"}
    // chromeFlags := []string{"--app", url}

    // browsers in order of preference
    var browserOrder = []string{
        "chrome-x86",
        "chrome-x64",
        "chromium-x86",
        "chromium-x64",
        "msedge-x86",
        "msedge-x64",
        "msedge",
        "brave-x86",
        "brave-x64",
        "vivaldi-x86",
        "vivaldi-x64",
    }


    // check if <BROWSER> is installed, if so open with flags
    // else open with os default browser

    // loop through browsers in order
    for _, browser := range browserOrder {
        // check if browser is installed
        var browserPath = browsers[browser][runtime.GOOS]
        log.Debug("Checking if " + browser + " is installed in " + browserPath)

        if _, err := os.Stat(browserPath); err == nil {
            // browser is installed, open with flags
            log.Debug("Opening with " + browser)
            
            var cmd *exec.Cmd
            
            var browserFlags = browsers[browser]["flags"]
            browserFlags = strings.Replace(browserFlags, "{{url}}", url, -1)
            log.Debug("Flags: " + browserFlags)

            var commandArr = commandRun[runtime.GOOS]
            commandArr = append(commandArr, browserPath)
            commandArr = append(commandArr, strings.Split(browserFlags, " ")...)


            log.Debug("Command: " + strings.Join(commandArr, " "))

            cmd = exec.Command(commandArr[0], commandArr[1:]...)
            err = cmd.Start()

            if err == nil {
                didOpen = true
                break
            } else {
                log.Info("Error opening browser automatically, please open " + url + " in your browser")
                fmt.Println(err)
            }
        }
    }

    // if no browser is installed, open with os default browser
    if (err == nil && didOpen == false) {
        log.Debug("Opening with OS default browser")

        var cmd *exec.Cmd

        switch runtime.GOOS {
            case "linux":
                cmd = exec.Command("xdg-open", url)
            case "windows":
                cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
            case "darwin":
                cmd = exec.Command("open", url)
            default:
                err = fmt.Errorf("unsupported platform")
        }

        if (err == nil) {
            err = cmd.Start()
        }
    }        

    if (err != nil) {
        log.Info("Error opening browser automatically, please open " + url + " in your browser")
    }
}