package blogposts_test

import (
	"errors"
	blogposts "github.com/daBrian/learn-go-with-tests/reading-files"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Goodbye, World!")
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: Post 2")},
	}

	posts, err := blogposts.NewPostsFromFs(fs)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d want %d", len(posts), len(fs))
	}
	got := posts[0]
	want := blogposts.Post{Title: "Post 1"}

	assert.Equal(t, want, got)
}
