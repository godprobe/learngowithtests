package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	"blogrenderer"

	approvals "github.com/approvals/go-approval-tests"
)

var (
	cases = []blogrenderer.Post{
		{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		},
		{
			Title:       "a new world",
			Body:        "This is another post",
			Description: "This is another description",
			Tags:        []string{"tdd", "go", "experiments"},
		},
	}
)

func TestRender(t *testing.T) {
	buf := bytes.Buffer{}

	// POST RENDERER
	for _, post := range cases {
		t.Run("convert a single post into HTML", func(t *testing.T) {
			buf.Reset()

			postRenderer, err := blogrenderer.NewPostRenderer()
			if err != nil {
				t.Fatal(err)
			}

			if err := postRenderer.Render(&buf, post); err != nil {
				t.Fatal(err)
			}

			approvals.VerifyString(t, buf.String())
		})
	}

	// INDEX RENDERER
	t.Run("render the index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		postRenderer, err := blogrenderer.NewPostRenderer()
		if err != nil {
			t.Fatal(err)
		}

		posts := []blogrenderer.Post{{Title: "hello world"}, {Title: "a new world"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
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
