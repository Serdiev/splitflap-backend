package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func GetUrl[resType any](url string, res resType) (*resType, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
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
