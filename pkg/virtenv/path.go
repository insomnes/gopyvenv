package virtenv

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/insomnes/gopyvenv/pkg/infra"
)

const rootPath = "/"

func absPathContains(src string, target string) bool {
	msg := "Comparing SRC" + src + " to TARGET " + target
	infra.DebugLog(msg)

	if len(target) < len(src) {
		infra.DebugLog("Target is too short")
		return false
	}

	for len(target) >= len(src) {
		if target == src {
			infra.DebugLog("Target is source")
			return true
		}
		target = filepath.Dir(target)
		infra.DebugLog("New TARGET is " + target)
		if target == rootPath {
			infra.DebugLog("We are in FS root, so we should stop here")
			return false
		}
	}
	infra.DebugLog("Target is too short")

	return false
}

func searchScriptRecursively(cwd string, venvDirsBase []string, script string) string {
	// We don't want errors on non existing script for some reason
	for len(cwd) > 1 {
		infra.DebugLog("CWD: " + cwd)
		venvDirs := cwdCombinations(cwd, venvDirsBase)
		infra.DebugLog(
			fmt.Sprintf("Searching for script %s in venv dirs: %v", script, venvDirs),
		)
		for _, vd := range venvDirs {
			scriptPath := filepath.Join(cwd, vd, script)
			if _, err := os.Stat(scriptPath); err == nil {
				infra.DebugLog("Found activate script file at " + scriptPath)
				return scriptPath
			}
		}
		cwd = filepath.Dir(cwd)
	}

	return ""
}

var combDels = [...]string{"_", "-"}

func cwdCombinations(cwd string, venvDirs []string) []string {
	if cwd == "/" {
		return venvDirs
	}
	projName := filepath.Base(cwd)

	infra.DebugLog("Cut cwd to: " + projName)
	combined := make([]string, 0, len(venvDirs)*(len(combDels)+1))
	for _, vd := range venvDirs {
		combined = append(combined, vd)
		if projName == "." || projName == "" {
			continue
		}
		for _, d := range combDels {
			combined = append(combined, projName+d+vd)
		}
	}
	infra.DebugLog(fmt.Sprintf("Combined venv dirs: %v", combined))
	return combined
}
