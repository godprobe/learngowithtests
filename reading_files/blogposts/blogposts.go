package blogposts

import "io/fs"

type Post struct {
}

func NewPostsFromFS(fileSystem fs.FS) []Post {
	var posts []Post
	dir, _ := fs.ReadDir(fileSystem, ".")

	for range dir {
		posts = append(posts, Post{})
	}

	return posts
}
