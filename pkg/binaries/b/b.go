package b

import (
	"context"
	"fmt"
	"runtime"

	"github.com/buyoio/b/pkg/binaries"
	"github.com/buyoio/b/pkg/binary"
)

func Binary(options *binaries.BinaryOptions) *binary.Binary {
	if options == nil {
		options = &binaries.BinaryOptions{
			Context: context.Background(),
		}
	}
	// https://github.com/buyoio/b/releases/download/v1.0.0/b-linux-amd64.tar.gz
	return &binary.Binary{
		Context:    options.Context,
		Envs:       options.Envs,
		Tracker:    options.Tracker,
		Version:    options.Version,
		Name:       "b",
		GitHubRepo: "buyoio/b",
		GitHubFile: fmt.Sprintf("b-%s-%s.tar.gz", runtime.GOOS, runtime.GOARCH),
		VersionF:   binary.GithubLatest,
		IsTarGz:    true,
		VersionLocalF: func(b *binary.Binary) (string, error) {
			return b.Exec("--version")
		},
	}
}
