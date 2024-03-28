package validation

import (
	"errors"
	"net/url"
)

func ValidateUrl(website string) error {
	if website == "" {
		return errors.New("Url cannot be empty")
	}

	_, err := url.ParseRequestURI(website)
	if err != nil {
		return err
	}

	return nil
}
