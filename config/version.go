package config

import (
	"fmt"
	"runtime"
	"strings"
)

var (
	GitCommit   string
	VersionDesc string
)

const (
	Version = "2.1.0-pre"
)

func init() {
	gitCommit := ""
	if len(GitCommit) > 0 {
		gitCommit = "; git " + GitCommit
	}
	VersionDesc = fmt.Sprintf("git-lfs/%s (GitHub; %s %s; go %s%s)",
		Version,
		runtime.GOOS,
		runtime.GOARCH,
		strings.Replace(runtime.Version(), "go", "", 1),
		gitCommit,
	)
}
