# Photo Fetcher
## Parallel image downloads with Go routines
With Photo Fetcher it's easy to download multiple images from different
locations and store them in memory as an image.Image object.

## Installation

```
$ go get -u github.com/jelmersnoeck/photo-fetcher
```

## Usage

```go
    fetcher := New([]string{
        "https://www.google.co.uk/images/srpr/logo11w.png",
    })
    fetcher.Run()
```
