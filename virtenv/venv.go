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

type VenvMeta struct {
	venv           Venv
	cwd            string
	venvDirs       []string
	activateScript string
}

func getCommandOnVenv(vm VenvMeta) string {
	venv, cwd := vm.venv, vm.cwd
	venvDirs, scriptToSearch := vm.venvDirs, vm.activateScript

	if !venv.Active {
		script := searchScriptRecursively(cwd, venvDirs, scriptToSearch)
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
	meta := VenvMeta{
		venv:           venv,
		cwd:            cwd,
		venvDirs:       venvDirs,
		activateScript: activateScript,
	}
	return getCommandOnVenv(meta)
}
