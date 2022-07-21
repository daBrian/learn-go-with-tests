package blogposts

import (
	"io/fs"
)

func NewPostsFromFs(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(system fs.FS, fn string) (post Post, err error) {
	var postFile fs.File
	postFile, err = system.Open(fn)
	if err != nil {
		return Post{}, err
	}
	defer func(postFile fs.File) {
		err = postFile.Close()
		if err != nil {
			post = Post{}
		}
	}(postFile)
	post, err = newPost(postFile)
	return
}
