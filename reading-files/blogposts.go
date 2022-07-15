package blogposts

import (
	"io/fs"
)

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
