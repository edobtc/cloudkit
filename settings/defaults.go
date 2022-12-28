package settings

const (
	// defaultURL is the url default to set for
	// settings, we should eventually update this to be
	// whatever the prod url to build for is
	//
	// it can be updated via
	// btck settings set-url --url=btck.staging.blah.com
	defaultURL = "http://127.0.0.1:8081"

	// SettingsPath
	defaultSettingsPath = ".btck/settings.yaml"
)
