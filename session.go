package kinto

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Session struct {
	baseURL  string
	client   *http.Client
	user     string
	password string
}

func NewSession(baseURL string, user string, password string) (*Session, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	session := &Session{
		baseURL: baseURL,
		client:  client,
		user:user,
		password:password,
	}
	return session, nil
}

func (s Session) Request(method string, path string, body io.Reader, target interface{}) error {
	req, err := http.NewRequest(method, s.baseURL+path, body)
	if err != nil {
		return err
	}

	req.SetBasicAuth(s.user, s.password)
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
