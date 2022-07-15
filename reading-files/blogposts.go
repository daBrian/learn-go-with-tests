package blogposts

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFs(fileSystem fs.FS) ([]Post, error) {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for _, f := range dir {
		posts = append(posts, getPost(fileSystem, f.Name()))
	}
	return posts, nil
}

func getPost(system fs.FS, fn string) Post {
	postFile, _ := system.Open(fn)
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile io.Reader) Post {
	postData, _ := io.ReadAll(postFile)
	post := Post{Title: string(postData)[7:]}
	return post
}
