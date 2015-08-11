package photo_fetcher

import (
	"image"
	"net/http"
	"sync"
)

type PhotoFetcher struct {
	image_urls []string
}

func New(urls []string) (f *PhotoFetcher) {
	f = new(PhotoFetcher)
	f.image_urls = urls

	return
}

func (f *PhotoFetcher) Run() []image.Image {
	images := make([]image.Image, len(f.image_urls))

	var wg sync.WaitGroup

	for index, url := range f.image_urls {
		wg.Add(1)

		go downloadImage(url, index, images, &wg)
	}

	wg.Wait()

	return images
}

func downloadImage(url string, index int, images []image.Image, wg *sync.WaitGroup) {
	defer wg.Done()

	res, _ := http.Get(url)
	defer res.Body.Close()

	img, _, _ := image.Decode(res.Body)
	images[index] = img
}
