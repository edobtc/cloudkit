package settings

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var settings = NewSettingsWithDefaults()

// NewSettingsWithDefaults returns a settings
// object with data defaults set, primarily
// to encapsulate defaults that can be
// over-ridden by those in the know via
// btck settings set-url --key=value
func NewSettingsWithDefaults() Settings {
	return Settings{
		Data: Data{
			URL: defaultURL,
		},
	}
}

// Settings is the container for the
// cli settings
type Settings struct {
	Data Data
}

// ToYAML returns the settings serialized back as it's
// saved YAML format, primarily for printing to stdout
func (s Settings) ToYAML() ([]byte, error) {
	return yaml.Marshal(&s)
}

// Data is the settings entry which holds
// settings data, currently this is only the current
// JWT token for making requests
type Data struct {
	Token string `yaml:"token"`
	URL   string `yaml:"url"`
}

// SavePath is a helper for setting
// up the user's home path
func SavePath() (string, error) {
	if path := os.Getenv("SETTINGS_PATH"); path != "" {
		return fmt.Sprintf("%s/%s", path, defaultSettingsPath), nil
	}

	return pathToHome()
}

func pathToHome() (string, error) {
	p, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", p, defaultSettingsPath), nil
}

// Read will read the settings from
// the settings path
func Read() (*Settings, error) {
	present, err := Exists()
	if err != nil {
		return nil, err
	}

	if !present {
		err = Reset()
		if err != nil {
			return nil, err
		}

		return &settings, nil
	}

	path, err := SavePath()
	if err != nil {
		return &settings, err
	}

	f, err := os.ReadFile(path)
	if err != nil {
		log.Error(err)
		return &settings, err
	}
	err = yaml.Unmarshal(f, &settings)
	if err != nil {
		log.Error(err)
		return &settings, err
	}

	return &settings, nil
}

// Exists test if any settings file exists
func Exists() (bool, error) {
	path, err := SavePath()
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(path); err == nil {
		return true, nil
	}

	return false, nil
}

// Load will load an existing settings file
// from the default save location
func Load() (*Settings, error) {
	path, err := SavePath()
	if err != nil {
		return nil, err
	}

	s := &Settings{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&s); err != nil {
		return nil, err
	}

	return s, nil
}

// Write will write a settings structure
// to the default file save location
func Write() error {
	d, err := yaml.Marshal(settings)
	if err != nil {
		return err
	}

	path, err := SavePath()
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0770)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(d)
	if err != nil {
		return err
	}

	return f.Sync()
}

// Reset will reinitialize the settings var
// with the defaults, useful for debugging or
// completely resetting customized settings
// and here to be called from the cli to reset a
// user's workstation
func Reset() error {
	settings = NewSettingsWithDefaults()
	return Write()
}

// Remove will completely remove a settings file
func Remove() error {
	path, err := SavePath()
	if err != nil {
		return err
	}
	return os.Remove(path)
}
