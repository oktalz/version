package version

import (
	"errors"
	"runtime/debug"
	"strings"
)

// Repo holds the repository path of the build.
//
// Version holds the version string of the build.
//
// CommitDate holds the commit date of the build.
var (
	Repo       = ""
	Version    = "dev"
	CommitDate = ""
)

// ErrBuildDataNotReadable is returned when build data cannot be read.
var ErrBuildDataNotReadable = errors.New("not able to read build data")

// Set populates the Repo, Version, and CommitDate variables with build information.
//
// It returns an error if the build data cannot be read.
func Set() error {
	buildinfo, ok := debug.ReadBuildInfo()
	if !ok {
		return ErrBuildDataNotReadable
	}
	Repo = buildinfo.Main.Path
	CommitDate = get(buildinfo, "vcs.time")

	Version = strings.Replace(buildinfo.Main.Version, "(devel)", "dev", 1)

	return nil
}

func get(buildInfo *debug.BuildInfo, key string) string {
	for _, setting := range buildInfo.Settings {
		if setting.Key == key {
			return setting.Value
		}
	}
	return ""
}
