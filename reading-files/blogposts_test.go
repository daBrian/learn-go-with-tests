package blogposts_test

import (
	"errors"
	blogposts "github.com/daBrian/learn-go-with-tests/reading-files"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/fs"
	"testing"
	"testing/fstest"
)

type MockFS struct {
	mock.Mock
}

func (s *MockFS) Open(name string) (file fs.File, err error) {
	args := s.Called(name)
	if args.Get(0) != nil {
		file = args.Get(0).(fs.File)
	}
	if args.Get(1) != nil {
		err = args.Error(1)
	}
	return
}

type StubDirEntry struct {
	mock.Mock
}

func (f *StubDirEntry) ReadDir(_ int) ([]fs.DirEntry, error) {
	return []fs.DirEntry{f}, nil
}

func (f *StubDirEntry) Close() error {
	args := f.Called()
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (f *StubDirEntry) Name() string {
	return "filename"
}

func TestNewBlogPosts(t *testing.T) {
	fakeFs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: Post 2")},
	}

	posts, err := blogposts.NewPostsFromFs(fakeFs)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fakeFs) {
		t.Errorf("got %d want %d", len(posts), len(fakeFs))
	}
	want := []blogposts.Post{{Title: "Post 1"}, {Title: "Post 2"}}

	assert.Equal(t, want, posts)
}
func TestNoBlogPosts(t *testing.T) {
	fakeFs := fstest.MapFS{}
	posts, err := blogposts.NewPostsFromFs(fakeFs)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fakeFs) {
		t.Errorf("got %d want %d", len(posts), len(fakeFs))
	}
	assert.Empty(t, posts)
}

func TestReadDirError(t *testing.T) {
	fakeFs := new(MockFS)
	want := errors.New("Goodbye, World!")
	fakeFs.On("Open", mock.Anything).Return(nil, want).Once()
	posts, err := blogposts.NewPostsFromFs(fakeFs)
	assert.Nil(t, posts)
	assert.Error(t, err)
	assert.Same(t, want, err)
	fakeFs.AssertExpectations(t)
}

func TestReadFileError(t *testing.T) {
	want := errors.New("Goodbye, World!")
	fakeFs := new(MockFS)
	fakeDir := new(StubDirEntry)
	fakeFs.On("Open", mock.Anything).Return(fakeDir, nil).Once()
	fakeFs.On("Open", mock.Anything).Return(nil, want).Once()
	fakeDir.On("Close").Return(nil).Once()
	posts, err := blogposts.NewPostsFromFs(fakeFs)
	assert.Nil(t, posts)
	assert.Same(t, want, err)
	fakeFs.AssertExpectations(t)
}
func (f *StubDirEntry) IsDir() bool {
	panic("implement me")
}

func (f *StubDirEntry) Type() fs.FileMode {
	panic("implement me")
}

func (f *StubDirEntry) Info() (fs.FileInfo, error) {
	panic("implement me")
}

func (f *StubDirEntry) Stat() (fs.FileInfo, error) {
	panic("implement me")
}

func (f *StubDirEntry) Read(_ []byte) (int, error) {
	panic("implement me")
}
