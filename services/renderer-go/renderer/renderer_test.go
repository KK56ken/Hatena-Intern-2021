package renderer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Header(t *testing.T) {
	src := "# Test"
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t, "<h1>Test</h1>\n", html)
}
func Test_List(t *testing.T) {
	src := "- list1"
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t ,"<ul>\n<li>list1</li>\n</ul>\n", html)
}
func Test_Link(t *testing.T) {
	src := "[はてなブックマーク](https://b.hatena.ne.jp/)"
	html, err := Render(context.Background(), src)
	assert.NoError(t, err)
	assert.Equal(t ,"<p><a href=\"https://b.hatena.ne.jp/\">はてなブックマーク</a></p>\n", html)
}