package virtenv

import (
	"os"
	"path"
	"path/filepath"
	"reflect"
	"testing"
)

var (
	script   = "fake-script"
	venvDirs = []string{"venv"}
)

func createDir(dir string) error {
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func createVenvDir(t *testing.T) string {
	venvDir := filepath.Join(t.TempDir(), venvDirs[0])
	err := createDir(venvDir)
	if err != nil {
		t.Fatal(err)
	}
	return venvDir
}

func createVenvDirWithScript(t *testing.T) (string, string) {
	venvDir := createVenvDir(t)
	scriptPath := filepath.Join(venvDir, script)
	f, err := os.Create(scriptPath)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	return venvDir, scriptPath
}

func TestExistingScript(t *testing.T) {
	venvDir, scriptPath := createVenvDirWithScript(t)
	cwd := path.Dir(venvDir)

	calculatedPath := searchScriptRecursively(cwd, venvDirs, script)
	if calculatedPath != scriptPath {
		t.Error("Cant find script, but should")
	}
}

func TestExistingScriptInParent(t *testing.T) {
	venvDir, scriptPath := createVenvDirWithScript(t)
	cwd := path.Dir(venvDir)

	cwdChild := path.Join(cwd, "tmp")
	calculatedWithChild := searchScriptRecursively(cwdChild, venvDirs, script)
	if calculatedWithChild != scriptPath {
		t.Error("Cant find script from child, but should")
	}
}

func TestNonExistingScript(t *testing.T) {
	venvDir := createVenvDir(t)
	cwd := path.Dir(venvDir)

	calculatedPath := searchScriptRecursively(cwd, venvDirs, script)
	if calculatedPath != "" {
		t.Error("Found something as scrip but should not")
	}
}

func TestNonSense(t *testing.T) {
	pseudoCwd := "/foo/bar/baz/who/is/../here/./for/../once"
	calculatedPath := searchScriptRecursively(pseudoCwd, venvDirs, script)
	if calculatedPath != "" {
		t.Error("Found something as scrip but should not")
	}
}

type absPathContainsTestCase struct {
	name     string
	src      string
	dst      string
	expected bool
}

var absPathContainsCases = []absPathContainsTestCase{
	{"two roots", "/", "/", true},
	{"contains", "/foo", "/foo/bar", true},
	{"does not containl", "/foo/bar", "/baz/meow", false},
	{"reversed", "/foo/bar", "/foo", false},
	{"not abs", ".", "/foo", false},
	{"same name", "/foo/bar/baz", "/bar/foo/baz", false},
}

func TestAbsPathContains(t *testing.T) {
	for _, testCase := range absPathContainsCases {
		contains := absPathContains(testCase.src, testCase.dst)
		if contains != testCase.expected {
			t.Errorf("Failed test: %s => %t != %t", testCase.name, contains, testCase.expected)
		}
	}
}

type cwdCombinationsTestCase struct {
	name     string
	cwd      string
	venvDirs []string
	expected []string
}

var cwdCombintationsCases = []cwdCombinationsTestCase{
	{
		"simple",
		"/foo/bar",
		[]string{"venv"},
		[]string{"venv", "bar_venv", "bar-venv"},
	},
	{
		"multiple venvDirs",
		"/foo/bar",
		[]string{"venv1", "venv2"},
		[]string{"venv1", "bar_venv1", "bar-venv1", "venv2", "bar_venv2", "bar-venv2"},
	},
	{
		"nested cwd",
		"/foo/bar/baz",
		[]string{"venv"},
		[]string{"venv", "baz_venv", "baz-venv"},
	},
	{
		"no venvDirs",
		"/foo/bar",
		[]string{},
		[]string{},
	},
	{
		"root cwd",
		"/",
		[]string{"venv"},
		[]string{"venv"},
	},
	{
		"empty cwd",
		"",
		[]string{"venv"},
		[]string{"venv"},
	},
}

func TestCwdCombinations(t *testing.T) {
	for _, tc := range cwdCombintationsCases {
		combined := cwdCombinations(tc.cwd, tc.venvDirs)
		if !reflect.DeepEqual(combined, tc.expected) {
			t.Errorf("Failed test: %s => %v != %v", tc.name, combined, tc.expected)
		}
	}
}
