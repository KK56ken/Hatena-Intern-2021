package renderer

import (
	"bytes"
	"context"
	"github.com/yuin/goldmark"
)

// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	// TODO: これはサンプル実装 (URL の自動リンク) です
	var w bytes.Buffer
	source := []byte(src)

	if err := goldmark.Convert(source, &w); err != nil {
		panic(err)
	}

	return w.String(), nil
}
