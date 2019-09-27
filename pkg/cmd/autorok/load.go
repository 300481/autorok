package autorok

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-yaml/yaml"
)

const (
	YAML             = "yaml"
	JSON             = "json"
	wrongFormatError = "Got wrong format."
	failedHttpGet    = "HTTP Get failed."
)

func loadBytes(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	log.Printf("Get URL: %s", url)
	r, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode > 299 {
		return nil, errors.New(failedHttpGet)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func loadObject(url, format string, object interface{}) error {
	b, err := loadBytes(url)
	if err != nil {
		return err
	}

	log.Printf("Use format '%s' to parse object.", format)
	switch format {
	case YAML:
		err = yaml.Unmarshal(b, object)
	case JSON:
		err = json.Unmarshal(b, object)
	default:
		return errors.New(wrongFormatError)
	}

	if err == nil {
		log.Println("Load Object done")
	}

	return err
}
