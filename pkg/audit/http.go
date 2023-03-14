package audit

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var header = http.Header{
	"User-Agent": {"DevCrawler/0.1"},
}

func getDomain(s string) string {
	r, _ := url.Parse(s)
	return fmt.Sprintf("%s://%s", r.Scheme, r.Host)
}

func checkActiveUrl(url string) (uint16, error) {
	c := http.Client{}
	res, err := c.Head(url)
	if err != nil {
		return 0, err
	}

	return uint16(res.StatusCode), nil
}

func fetchPageDocument(url string) ([]byte, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("Cannot initialize new request")
	}

	req.Header = header

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, ioErr := io.ReadAll(res.Body)

	if ioErr != nil {
		return nil, errors.New("Cannot read response body from request")
	}

	return body, nil
}
