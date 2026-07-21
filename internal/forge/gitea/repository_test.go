package gitea

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.abhg.dev/gs/internal/forge"
)

func TestRepository_ChangeURL(t *testing.T) {
	r := &Repository{
		owner: "scotty",
		repo:  "warp-core",
		forge: &Forge{Options: Options{URL: "https://gitea.example.com"}},
	}
	assert.Equal(t,
		"https://gitea.example.com/scotty/warp-core/pulls/42",
		r.ChangeURL(&PR{Number: 42}),
	)
}

func TestRepository_ComparisonURL(t *testing.T) {
	r := &Repository{
		owner: "scotty",
		repo:  "warp-core",
		forge: &Forge{Options: Options{URL: "https://gitea.example.com"}},
	}
	assert.Equal(t,
		"https://gitea.example.com/scotty/warp-core/compare/main...feat",
		r.ComparisonURL("main", "feat"),
	)
}

// Verify interface implementations at compile time.
var (
	_ forge.Repository        = (*Repository)(nil)
	_ forge.WithChangeURL     = (*Repository)(nil)
	_ forge.WithComparisonURL = (*Repository)(nil)
)
