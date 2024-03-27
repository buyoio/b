package k9s

import (
	"bytes"
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/buyoio/b/pkg/binaries"
	"github.com/buyoio/b/pkg/binary"
)

func NewK9s(options *binaries.BinaryOptions) *binary.Binary {
	if options == nil {
		options = &binaries.BinaryOptions{
			Context: context.Background(),
		}
	}
	return &binary.Binary{
		Context:    options.Context,
		Envs:       options.Envs,
		Tracker:    options.Tracker,
		Name:       "k9s",
		GitHubRepo: "derailed/k9s",
		GitHubFile: fmt.Sprintf(
			"k9s_%s_%s.tar.gz",
			string(append(bytes.ToUpper([]byte{runtime.GOOS[0]}), runtime.GOOS[1:]...)),
			runtime.GOARCH,
		),
		VersionF: binary.GithubLatest,
		IsTarGz:  true,
		VersionLocalF: func(b *binary.Binary) (string, error) {
			s, err := b.Exec("version", "-s")
			if err != nil {
				return "", err
			}
			v := strings.Split(strings.SplitN(s, "\n", 2)[0], " ")
			return v[len(v)-1], nil
		},
	}
}
