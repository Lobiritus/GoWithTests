package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebSiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestWebSiteChecker(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://yandex.ru",
		"waat://furhurterwe.geds"}

	want := map[string]bool{
		"http://google.com":       true,
		"http://yandex.ru":        true,
		"waat://furhurterwe.geds": false,
	}

	got := CheckWebsites(mockWebSiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func SLowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(SLowStubWebsiteChecker, urls)
	}

}
