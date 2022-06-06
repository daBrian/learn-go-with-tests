package blogposts_test

import (
	blogposts "github.com/daBrian/learn-go-with-tests/reading-files"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts := blogposts.NewPostsFromFs(fs)
	if len(posts) != len(fs) {
		t.Errorf("got %d want %d", len(posts), len(fs))
	}
}
