package photo_fetcher

import "testing"

// This needs an internet connection at the moment. Need to dig into testing and
// mocking more.
func TestRun(t *testing.T) {
	fetcher := New([]string{
		"https://www.google.co.uk/images/srpr/logo11w.png",
	})
	fetcher.Run()
}
