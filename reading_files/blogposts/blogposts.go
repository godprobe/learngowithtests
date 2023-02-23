package blogposts

import "testing/fstest"

type Post struct {
}

func NewPostsFromFS(fileSystem fstest.MapFS) []Post {
	var posts []Post
	dir, _ := fs.ReadDir(fileSystem, ".")

	for range dir {
		posts = append(posts, Post{})
	}

	return posts
}
