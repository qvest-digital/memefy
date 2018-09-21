package rest

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type VideoHandler func(name string, source io.Reader) (err error)

type Fetcher struct {
	handler    VideoHandler
	baseUrl    string
	httpClient http.Client
}

type FetcherOption func(Fetcher) Fetcher

func NewFetcher(baseUrl string, handler VideoHandler, opts ...FetcherOption) Fetcher {
	fetcher := Fetcher{handler: handler, baseUrl: baseUrl}
	for i := range opts {
		fetcher = opts[i](fetcher)
	}
	return fetcher
}

func WithHttpClient(client http.Client) FetcherOption {
	return func(fetcher Fetcher) Fetcher {
		fetcher.httpClient = client
		return fetcher
	}
}

func (f Fetcher) Do(memes []string) error {
	errGroup := errgroup.Group{}
	for i := range memes {
		errGroup.Go(func() error {
			url := f.baseUrl + memes[i]
			resp, err := f.httpClient.Get(url)
			if err != nil {
				return fmt.Errorf("Could not get meme %s", memes[i])
			}
			defer resp.Body.Close()
			err = f.handler(memes[i], resp.Body)
			if err != nil {
				return fmt.Errorf("Could not handle meme %s: %s", memes[i], err.Error())
			}
			return nil
		})
		return errGroup.Wait()
	}
}
