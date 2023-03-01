package blogrenderer_test

import (
	"bytes"
	"testing"

	"blogrenderer"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
		aSecondPost = blogrenderer.Post{
			Title:       "a new world",
			Body:        "This is another post",
			Description: "This is another description",
			Tags:        []string{"tdd", "go", "experiments"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

		buf.Reset()
		err = blogrenderer.Render(&buf, aSecondPost)

		if err != nil {
			t.Fatal(err)
		}

		got = buf.String()
		want = `<h1>a new world</h1>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
