package virtenv

import (
	"os"
	"path/filepath"
)

const rootPath = "/"

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
		if target == rootPath {
			debugLog("We are in FS root, so we should stop here")
			return false
		}
	}
	debugLog("Target is too short")

	return false
}

func searchScriptRecursively(cwd string, venvDirs []string, script string) string {
	// We don't want errors on non existing script for some reason
	for len(cwd) > 1 {
		for _, vd := range venvDirs {
			scriptPath := filepath.Join(cwd, vd, script)
			if _, err := os.Stat(scriptPath); err == nil {
				debugLog("Found activate script file at " + scriptPath)
				return scriptPath
			}
		}
		cwd = filepath.Dir(cwd)
	}

	return ""
}
