package blogrenderer

import (
	"bytes"
	"embed"
	"github.com/yuin/goldmark"
	"io"
	"strings"
	"text/template"
)

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	temp, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: temp}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	var err error
	p.Body, err = r.ConvertMarkDownToHTML(p.Body)
	if err != nil {
		return err
	}
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}
	return nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

func (r *PostRenderer) ConvertMarkDownToHTML(body string) (string, error) {
	md := goldmark.New()
	var buf bytes.Buffer
	if err := md.Convert([]byte(body), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}
