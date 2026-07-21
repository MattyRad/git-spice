package gitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.abhg.dev/gs/internal/gateway/gitlab"
)

func TestAccessValueName(t *testing.T) {
	t.Run("known", func(t *testing.T) {
		assert.Equal(t, "admin", accessValueName(gitlab.AdminPermissions))
	})

	t.Run("unknown", func(t *testing.T) {
		assert.Equal(t, "999", accessValueName(gitlab.AccessLevelValue(999)))
	})
}

func TestRepository_NavigationReference(t *testing.T) {
	repo := &Repository{}
	assert.Equal(t, "!42+", repo.NavigationReference(&MR{Number: 42}))
}

func TestRepository_ComparisonURL(t *testing.T) {
	r := &Repository{
		owner: "example",
		repo:  "repo",
		forge: &Forge{Options: Options{URL: "https://gitlab.example.com"}},
	}
	assert.Equal(t,
		"https://gitlab.example.com/example/repo/-/compare/main...feat",
		r.ComparisonURL("main", "feat"),
	)
}
