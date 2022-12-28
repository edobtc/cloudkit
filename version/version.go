package version

import (
	"fmt"
)

// Version The main version number
// that is being run at the moment.
var Version = "0.0.1"

// Prerelease is a pre-release marker for the version.
// If this is "" (empty string) then it means that
// it is a final release. Otherwise, this is a pre-release
// such as "dev" (in development), "beta", "rc1", etc.
var Prerelease = "dev"

// String returns the complete version string, including prerelease
func String() string {
	if Prerelease != "" {
		return fmt.Sprintf("%s-%s", Version, Prerelease)
	}
	return Version
}
