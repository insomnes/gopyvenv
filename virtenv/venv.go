package virtenv

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	virtualEnvVarKey string = "VIRTUAL_ENV"
	activateScript          = "bin/activate"
	deactivateCmd           = "deactivate 2> /dev/null || :"
	emptyCmd                = ""
)

type Venv struct {
	Active   bool
	VenvPath string
}

func getVenv() Venv {
	venvPath := os.Getenv(virtualEnvVarKey)
	return Venv{Active: len(venvPath) > 0, VenvPath: venvPath}
}

func getCommandOnVenv(venv Venv, cwd, script string) string {
	if !venv.Active {
		if script == "" {
			return emptyCmd
		}
		return fmt.Sprintf("source %s", script)
	}

	venvParentPath := filepath.Dir(venv.VenvPath)
	debugLog(fmt.Sprintf("Venv parent path: %s", venvParentPath))
	if venvParentPath == "." {
		debugLog("Venv parent is '.' something is broken")
		return emptyCmd
	}

	if absPathContains(venvParentPath, cwd) {
		debugLog(fmt.Sprintf("Venv parent: %s contains cwd: %s", venvParentPath, cwd))
		return emptyCmd
	}

	if script != "" && !absPathContains(venvParentPath, script) {
		debugLog("We are in dir with new script, activating it: " + script)
		return fmt.Sprintf("%s && source %s", deactivateCmd, script)
	}

	debugLog(fmt.Sprintf("Venv parent: %s DOES NOT contain cwd: %s", venvParentPath, cwd))
	// In case of some kind of broken situation where we dont have deactivate()
	return deactivateCmd
}

func GetCommand(venvDirs []string) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic("How the hell cant we get cwd?")
	}

	venv := getVenv()
	debugLog(fmt.Sprintf("Venv status: %v", venv))

	script := searchScriptRecursively(cwd, venvDirs, activateScript)

	return getCommandOnVenv(venv, cwd, script)
}
