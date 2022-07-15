package blogposts

import "io"

type Post struct {
	Title string
}

func newPost(postFile io.Reader) Post {
	postData, _ := io.ReadAll(postFile)
	post := Post{Title: string(postData)[7:]}
	return post
}
