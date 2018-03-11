package kinto

import (
	"fmt"
	"net/url"
)

type KintoClient struct {
	session *Session
}

type Options map[string]string

func NewClient(baseURL string, user string, password string) (*KintoClient, error) {
	session, err := NewSession(baseURL, user, password)

	if err != nil {
		return nil, err
	}

	kc := &KintoClient{session: session}
	return kc, nil
}

func (kc KintoClient) buildURI(format string, params ...interface{}) string {
	return fmt.Sprintf(format, params...)
}

func (kc KintoClient) HeartBeat() (interface{}, error) {
	path := kc.buildURI(HEARTBEAT_URI)
	var status map[string]bool
	err := kc.session.Request("GET", path, nil, nil, &status)
	return status, err
}

func buildParams(opts Options) string {
	params := url.Values{}
	for key, value := range opts {
		params.Add(key, value)
	}
	return params.Encode()
}
