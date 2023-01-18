package audit

import (
	"errors"
	"net/url"
)



func ValidateUrl(website string) (error) {
  if website == "" {
    return errors.New("Url is missing")
  }
  
  _, err := url.ParseRequestURI(website)
  if err != nil {
    return err 
  }

  return nil
}
