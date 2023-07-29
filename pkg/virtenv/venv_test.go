package virtenv

import "testing"

type getCommandOnVenvTestCase struct {
	name     string
	venv     Venv
	cwd      string
	script   string
	expected string
}

var getCommandOnVenvCases = []getCommandOnVenvTestCase{
	{
		name:     "not active no script",
		venv:     Venv{false, ""},
		cwd:      "",
		script:   "",
		expected: emptyCmd,
	},
	{
		name:     "not active with script",
		venv:     Venv{false, "/foo"},
		cwd:      "",
		script:   "/foo/bar",
		expected: "source /foo/bar",
	},
	{
		name:     "active nonsense path",
		venv:     Venv{true, "123"},
		cwd:      "",
		script:   "",
		expected: emptyCmd,
	},
	{
		name:     "active and contains",
		venv:     Venv{true, "/foo/venv"},
		cwd:      "/foo/bar",
		script:   "/foo/venv/bin/activate",
		expected: emptyCmd,
	},
	{
		name:     "active not contains no script",
		venv:     Venv{true, "/foo/venv"},
		cwd:      "/bar/baz",
		script:   "",
		expected: deactivateCmd,
	},
	{
		name:     "active not contains new script",
		venv:     Venv{true, "/some/venv"},
		cwd:      "/foo",
		script:   "/foo/venv/bin/activate",
		expected: deactivateCmd + " && source /foo/venv/bin/activate",
	},
}

func TestGetCommandOnVenv(t *testing.T) {
	for _, testCase := range getCommandOnVenvCases {
		command := getCommandOnVenv(testCase.venv, testCase.cwd, testCase.script)
		if command != testCase.expected {
			t.Errorf("%s => %s != %s", testCase.name, command, testCase.expected)
		}
	}
}
