package renderer

import (
	"bytes"
	"context"
	// "html/template"
	// "regexp"
	"github.com/yuin/goldmark"
	// "github.com/yuin/goldmark/parser"
)
// var urlRE = regexp.MustCompile(`https?://[^\s]+`)
// var linkTmpl = template.Must(template.New("link").Parse(`<a href="{{.}}">{{.}}</a>`))


// Render は受け取った文書を HTML に変換する
func Render(ctx context.Context, src string) (string, error) {
	// TODO: これはサンプル実装 (URL の自動リンク) です
	var w bytes.Buffer
	source := []byte(src)

	if err := goldmark.Convert(source, &w); err != nil {
		panic(err)
	}

	return w.String(), nil
	

	// html := urlRE.ReplaceAllStringFunc(src, func(url string) string {
	// 	err := linkTmpl.ExecuteTemplate(&w, "link", url)
	// 	if err != nil {
	// 		return url
	// 	}
	// 	return w.String()
	// })
	// return html, nil
}
