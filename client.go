package kinto

import (
	"fmt"
)

type KintoClient struct {
	session    *Session
}

func NewClient(baseURL string, user string, password string) (*KintoClient, error) {
	session, err := NewSession(baseURL, user, password)

	if err != nil {
		return nil, err
	}

	kc := &KintoClient{
		session: session,
	}

	return kc, nil
}
