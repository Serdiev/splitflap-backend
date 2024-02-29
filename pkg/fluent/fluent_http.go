package fluent

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type FluentHttp struct {
	url    string
	client *http.Client
	method string

	payload []byte

	headers map[string]string

	handlers     map[int]func(bytes []byte) error
	errorHandler func(bytes []byte) error
}

func create() *FluentHttp {
	c := &FluentHttp{
		headers:      map[string]string{},
		handlers:     map[int]func(bytes []byte) error{},
		client:       http.DefaultClient,
		errorHandler: defaultErrorHandler(),
	}
	return c
}

func defaultErrorHandler() func(bytes []byte) error {
	return func(bytes []byte) error {
		return errors.New(string(bytes))
	}
}

func Get(url string) *FluentHttp {
	c := create()
	c.url = url
	c.method = http.MethodGet
	return c
}

func Post(url string, payload []byte) *FluentHttp {
	c := create()
	c.url = url
	c.method = http.MethodPost
	c.payload = payload
	return c
}

func (c *FluentHttp) WithClient(customClient *http.Client) *FluentHttp {
	c.client = customClient
	return c
}

func (c *FluentHttp) WithHeader(key string, value string) *FluentHttp {
	c.headers[key] = value
	return c
}

func (c *FluentHttp) WithAuthorizationHeader(value string) *FluentHttp {
	c.headers["Authorization"] = value
	return c
}

func (c *FluentHttp) WithContentType(value string) *FluentHttp {
	c.headers["Content-Type"] = value
	return c
}

func (c *FluentHttp) WithClientCredentials(clientId string, clientSecret string) *FluentHttp {
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(clientId+":"+clientSecret))
	c.headers["Authorization"] = authHeader
	return c
}

func (c *FluentHttp) OnSuccess(handler func(bytes []byte) error) *FluentHttp {
	c.handlers[http.StatusOK] = handler
	return c
}

func (c *FluentHttp) OnError(handler func(bytes []byte) error) *FluentHttp {
	c.errorHandler = handler
	return c
}

func (c *FluentHttp) OnStatusCode(statusCode int, handler func(bytes []byte) error) *FluentHttp {
	c.handlers[statusCode] = handler
	return c
}

func (c *FluentHttp) Execute() error {
	req, err := http.NewRequest(c.method, c.url, bytes.NewBuffer(c.payload))
	if err != nil {
		return err
	}

	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	handler, exists := c.handlers[resp.StatusCode]
	if exists {
		return handler(respBody)
	}

	if c.errorHandler == nil {
		return errors.New("missing error handler")
	}

	return c.errorHandler(respBody)
}

func BytesToStruct[T interface{}](bytes []byte) (T, error) {
	var resp T
	err := json.Unmarshal(bytes, &resp)
	return resp, err
}
