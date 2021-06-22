package main

import (
	"errors"
	"github.com/petrostrak/Testing-with-Go/15-Reading-Files"
	"io/fs"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I failed.")
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"helloWorld.md":  {Data: []byte("hi")},
		"helloWorld2.md": {Data: []byte("hola")},
	}

	posts, err := NewPostsFromFS(StubFailingFS{})

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d, but want %d posts", len(posts), len(fs))
	}
}
