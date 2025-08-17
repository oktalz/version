package version

import (
	"errors"
	"runtime/debug"
	"strings"
)

// Version holds the version string of the build.
//
// Repo holds the repository path of the build.
//
// Modified indicates whether the build was modified after tag in vcs.
//
// Commit holds the commit hash of the build.
//
// CommitDate holds the commit date of the build.
var (
	Version    = "dev"
	Repo       = ""
	Modified   = ""
	Commit     = ""
	CommitDate = ""
)

// ErrBuildDataNotReadable is returned when build data cannot be read.
var ErrBuildDataNotReadable = errors.New("not able to read build data")

// Set populates the Version, Repo, Modified, Commit and CommitDate variables with build info data.
//
// It returns an error if the build data cannot be read.
func Set() error {
	buildinfo, ok := debug.ReadBuildInfo()
	if !ok {
		return ErrBuildDataNotReadable
	}
	Version = strings.Replace(buildinfo.Main.Version, "(devel)", "dev", 1)

	Repo = buildinfo.Main.Path
	Commit = get(buildinfo, "vcs.revision")
	if len(Commit) > 8 {
		Commit = Commit[:8]
	}
	CommitDate = get(buildinfo, "vcs.time")
	Modified = get(buildinfo, "vcs.modified")

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
