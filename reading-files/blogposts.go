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
		posts = append(posts, getPost(fileSystem, f))
	}
	return posts, nil
}

func getPost(system fs.FS, f fs.DirEntry) Post {
	postFile, _ := system.Open(f.Name())
	defer postFile.Close()
	postData, _ := io.ReadAll(postFile)
	post := Post{Title: string(postData)[7:]}
	return post
}
