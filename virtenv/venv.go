package virtenv

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	virtualEnvVarKey string = "VIRTUAL_ENV"
	activateScript   string = "bin/activate"
)

var debugOn = os.Getenv("GOVENV_DEBUG")

func debugLog(msg string) {
	if debugOn != "1" {
		return
	}
	fmt.Println(msg)
}

type Venv struct {
	Active   bool
	VenvPath string
}

func getVenv() Venv {
	venvPath := os.Getenv(virtualEnvVarKey)
	return Venv{Active: len(venvPath) > 0, VenvPath: venvPath}
}

func absPathContains(src string, target string) bool {
	msg := "Comparing SRC" + src + " to TARGET " + target
	debugLog(msg)

	if len(target) < len(src) {
		debugLog("Target is too short")
		return false
	}

	for len(target) >= len(src) {
		if target == src {
			debugLog("Target is source")
			return true
		}
		target = filepath.Dir(target)
		debugLog("New TARGET is " + target)
		if target == "/" {
			debugLog("We are in /, so we should stop here")
			return false
		}
	}
	debugLog("Target is too short")

	return false
}

func searchVenvScript(cwd string, venvDirs []string) string {
	for len(cwd) > 1 {
		for _, vd := range venvDirs {
			fullActScriptPath := filepath.Join(cwd, vd, activateScript)
			if _, err := os.Stat(fullActScriptPath); err == nil {
				return fullActScriptPath
			}
		}
		cwd = filepath.Dir(cwd)
	}

	return ""
}

func GetCommand() string {
	venvDirs := []string{"venv", ".venv"}
	cwd, err := os.Getwd()
	if err != nil {
		panic("How the hell cant we get cwd?")
	}

	activeVenv := getVenv()
	debugLog(fmt.Sprintf("Active venv status: %v", activeVenv))

	if !activeVenv.Active {
		script := searchVenvScript(cwd, venvDirs)
		if script == "" {
			return ""
		}
		return fmt.Sprintf("source %s", script)
	}

	venvParentPath := filepath.Dir(activeVenv.VenvPath)
	debugLog(fmt.Sprintf("Venv parent path: %s", venvParentPath))
	if venvParentPath == "." {
		debugLog("Venv parent is '.' something is broken")
		return ""
	}

	if absPathContains(venvParentPath, cwd) {
		debugLog(fmt.Sprintf("Venv parent: %s contains cwd: %s", venvParentPath, cwd))
		return ""
	}

	debugLog(fmt.Sprintf("Venv parent: %s DOES NOT contain cwd: %s", venvParentPath, cwd))
	// In case of some kind of broken situation where we dont have deactivate()
	return "deactivate 2> /dev/null || :"
}
