package blogposts_test

import (
	"GoWithTests/reading-files"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPost(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)
	t.Run("Test normal Data", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)
		if err != nil {
			t.Fatal(err)
		}
		assertPost(t, posts[0], blogposts.Post{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: `Hello
World`})

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})
	t.Run("Test Error ", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(&blogposts.StubFailingFS{})

		if err == nil {
			t.Errorf("Must be return error: %v", err)
		}
	})

}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
