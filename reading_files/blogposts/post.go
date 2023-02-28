package blogposts

import "io"

type Post struct {
	Title       string
	Description string
}

func newPost(postFile io.Reader) (Post, error) {
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(postData)[7:]} // remove "Title: "
	return post, nil
}
