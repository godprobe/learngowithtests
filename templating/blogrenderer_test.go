package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	"blogrenderer"

	approvals "github.com/approvals/go-approval-tests"
)

var (
	cases = []struct {
		Title       string
		Body        string
		Description string
		Tags        []string
	}{
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

			// got := buf.String()
			// want := fmt.Sprintf("<h1>%s</h1>\r\n", post.Title)
			// want += fmt.Sprintf("<p>%s</p>\r\n", post.Description)
			// want += "Tags: <ul>"
			// for _, tag := range post.Tags {
			// 	want += fmt.Sprintf("<li>%s</li>", tag)
			// }
			// want += "</ul>"
			// if got != want {
			// 	t.Errorf("got '%s' want '%s'", got, want)
			// }

			approvals.VerifyString(t, buf.String())
		})
	}

	// INDEX RENDERER
	t.Run("render the index of posts", func(t *testing.T) {
		buf.Reset()
		postRenderer, err := blogrenderer.NewPostRenderer()
		if err != nil {
			t.Fatal(err)
		}

		posts := []blogrenderer.Post{{Title: "Hello world"}, {Title: "Hello again"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<ol><li><a href="/posts/hello-world">Hello world</a></li><li><a href="/posts/hello-again">Hello again</a></li></ol>`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
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
