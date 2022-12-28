package settings

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	fb  = "fallback"
	ptr = "some pointer string"
)

func TestNewSettingsWithDefaultsURL(t *testing.T) {
	s := NewSettingsWithDefaults()
	assert.Equal(t, s.Data.URL, defaultURL)
}

func TestNewSettingsWithDefaultsToken(t *testing.T) {
	s := NewSettingsWithDefaults()
	assert.Equal(t, s.Data.Token, "")
}

var toYAMLTests = []struct {
	input    Settings // input
	expected string   // expected result
}{
	{
		NewSettingsWithDefaults(),
		"data:\n  token: \"\"\n  url: http://127.0.0.1:8081\n",
	},

	{
		Settings{
			Data: Data{
				Token: "secret-token",
			},
		},
		"data:\n  token: secret-token\n  url: \"\"\n",
	},

	{
		Settings{
			Data: Data{
				Token: "secret-token-2",
				URL:   "https://btck.edobtc.platform.com",
			},
		},
		"data:\n  token: secret-token-2\n  url: https://btck.edobtc.platform.com\n",
	},
}

func TestToYaml(t *testing.T) {
	for _, tt := range toYAMLTests {
		data, err := tt.input.ToYAML()
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, string(data), tt.expected)
	}
}

func TestDefaultDoesNotExist(t *testing.T) {
	settingsPath := "/tmp/blah"
	os.Setenv("SETTINGS_PATH", settingsPath)
	Remove()
	present, err := Exists()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, present, false)
	os.Unsetenv("SETTINGS_PATH")
}

func TestResetDoesExist(t *testing.T) {
	settingsPath := "/tmp/blah"
	os.Setenv("SETTINGS_PATH", settingsPath)
	Reset()
	present, err := Exists()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, present, true)
	os.Unsetenv("SETTINGS_PATH")
}

func TestSavePath(t *testing.T) {
	settingsPath := "/tmp"
	os.Setenv("SETTINGS_PATH", settingsPath)
	s, err := SavePath()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, s, fmt.Sprintf("%s/%s", settingsPath, defaultSettingsPath))
	os.Unsetenv("SETTINGS_PATH")
}

func TestSavePathHome(t *testing.T) {
	settingsPath := "/tmp"
	s, err := SavePath()
	if err != nil {
		t.Error(err)
	}
	assert.NotEqual(t, s, fmt.Sprintf("%s/%s", settingsPath, defaultSettingsPath))
}

func TestFSOperationsNewFileImitializedWithDefaults(t *testing.T) {
	settingsPath := "/tmp"

	os.Setenv("SETTINGS_PATH", settingsPath)

	s, err := Read()
	if err != nil {
		t.Error(err)
	}

	defaultSettings := NewSettingsWithDefaults()

	assert.Equal(t, s, &defaultSettings)
	Remove()
	os.Unsetenv("SETTINGS_PATH")
}

func TestFSOperationsNewFileSetAndWrite(t *testing.T) {
	settingsPath := "/tmp"
	token := "secret-token"

	os.Setenv("SETTINGS_PATH", settingsPath)

	s, err := Read()
	if err != nil {
		t.Error(err)
	}

	s.Data.Token = token

	err = Write()
	if err != nil {
		t.Error(err)
	}

	ns, err := Load()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, ns.Data.Token, token)
	Remove()
	os.Unsetenv("SETTINGS_PATH")
}

func TestFSOperationsNewFileSetAndReset(t *testing.T) {
	settingsPath := "/tmp"
	token := "secret-token"

	os.Setenv("SETTINGS_PATH", settingsPath)

	s, err := Read()
	if err != nil {
		t.Error(err)
	}

	s.Data.Token = token

	err = Write()
	if err != nil {
		t.Error(err)
	}

	ns, err := Load()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, ns.Data.Token, token)

	// Reset should set back to defaults
	Reset()

	rs, err := Load()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, rs.Data.Token, "")

	os.Unsetenv("SETTINGS_PATH")
}
