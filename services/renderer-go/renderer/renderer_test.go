package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Header(t *testing.T) {
	src := `# Test`
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, `<h1>Test</h1>`, html)
}
func Test_(t *testing.T) {
	src := ``
	
}