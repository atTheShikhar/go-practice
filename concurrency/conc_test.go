package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "https://example.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://example.com",
		"https://google.com",
		"https://microsoft.com",
	}

	want := map[string]bool{
		"https://example.com":   false,
		"https://google.com":    true,
		"https://microsoft.com": true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v, but got %v", want, got)
	}
}

func slowStubWebsiteChecker(url string) bool {
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
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
