package blogrenderer_test

import (
	"GoWithTests/blogrenderer"
	"bytes"
	approvals "github.com/approvals/go-approval-tests"
	"io"
	"testing"
)

func TestRenderer(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title: "hello world",
			Body: `# Добро пожаловать!

Это пример текста, написанного с использованием **Markdown**. 

## Что это такое?

Markdown — это простой язык разметки, который позволяет быстро создавать структурированный текст.

### Преимущества:

- Легкость в использовании
- Читаемость как в сыром виде, так и после обработки
- Поддержка списков, ссылок, изображений и многого другого`,
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := blogrenderer.NewPostRenderer()

	t.Run("It converts to single post HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err != nil {
			t.Fatal(err)
		}
		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())

	})

	t.Run("It renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}
		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRenderer(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
