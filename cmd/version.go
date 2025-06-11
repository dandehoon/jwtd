package cmd

import (
	"fmt"
	"runtime"
)

// Build-time variables that can be set using -ldflags
var (
	VERSION    = "dev"     // Version can be set at build time
	COMMIT     = "unknown" // Git commit hash
	BUILD_TIME = "unknown" // Build timestamp
)

// GetVersionInfo returns detailed version information
func GetVersionInfo() string {
	return fmt.Sprintf("jwtd version %s\nCommit: %s\nBuilt: %s\nGo version: %s\nOS/Arch: %s/%s",
		VERSION, COMMIT, BUILD_TIME, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
