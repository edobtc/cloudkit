package config

import "os"

const (
	DefaultNodeHost = "127.0.0.1:8332"

	DefaultSettingsPathName = ".btc-cloudkit"
)

const (
	// DefaultThing
	DefaultThing = "defaultThing"

	// LambdaStuff
	LambdaServerPortKey     = "_LAMBDA_SERVER_PORT"
	DefaultLambdaServerPort = "8181"

	// Filesystem Stuff
	// DefaultProvider is the default provider to use
	// if no config is provided to override the default
	// when deploying the service, or sent with a header
	// that lets a single request
	DefaultProvider = "aws/s3"

	// DefaultNamespace is the default key/filepath prefix to
	// namespace files to, it can be logs/files/data
	DefaultNamespace = "files"

	// MaxFileCount is the max number of files allowed
	// in any single upload
	MaxFileCount = 5

	DefaultSSHKeyName = "cloudkit-ssh-key"
)

func SetLambdaRuntimeDefaults() {
	serverPort := os.Getenv(LambdaServerPortKey)
	if serverPort == "" {
		os.Setenv(LambdaServerPortKey, DefaultLambdaServerPort)
	}
}
