package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	src := `# Test\n`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, `<h1>Test</h1>`, html)
}