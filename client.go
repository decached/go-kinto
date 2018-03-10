package kinto

import (
	"fmt"
)

type KintoClient struct {
	session *Session
}

func NewClient(baseURL string, user string, password string) (*KintoClient, error) {
	session, err := NewSession(baseURL, user, password)

	if err != nil {
		return nil, err
	}

	kc := &KintoClient{session: session}
	return kc, nil
}

func (c KintoClient) buildURI(format string, params ...interface{}) string {
	return fmt.Sprintf(format, params...)
}

func (c KintoClient) HeartBeat() (interface{}, error) {
	path := c.buildURI(HEARTBEAT_URI)
	var status map[string]bool
	err := c.session.Request("GET", path, nil, &status)
	return status, err
}
