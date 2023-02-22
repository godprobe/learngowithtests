package blogposts_test

import (
	"testing"
	"testing/fstest"
)

func TestNewBlogPost(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %v posts, wanted %v posts", len(posts), len(fs))
	}
}
