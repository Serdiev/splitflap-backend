package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func GetUrl[resType any](url string, res resType) (*resType, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		respBody, err := io.ReadAll(resp.Body)

		val := new(resType)

		err = json.Unmarshal(respBody, val)
		if err != nil {
			return nil, errors.New("could not unmarshal")
		}
		return val, nil
	}

	return nil, errors.New("unknown error")
}
