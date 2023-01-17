package audit

import (
	"errors"
	"io"
	"net/http"
)

func FetchPageDocument(url string) ([]byte, error) {
	client := http.Client{}
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, errors.New("Cannot initialize new request")
  }

  req.Header = http.Header{
    "User-Agent": {"DevCrawler/0.1"},
  }

  res, err := client.Do(req)
  if err != nil {
    return nil, errors.New("Something went wrong while fetchin website document") 
  }

  defer res.Body.Close()
  body, ioErr := io.ReadAll(res.Body)

  if ioErr != nil {
    return nil, errors.New("Something went wrong while parsing response body")
  }

  return body, nil 
}

