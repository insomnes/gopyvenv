package virtenv

import (
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

func searchScriptRecursively(cwd string, venvDirs []string, script string) string {
	// We don't want errors on non existing script for some reason
	for len(cwd) > 1 {
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
