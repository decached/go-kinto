package kinto

type Permission struct {
	Read  []string `json:"read,omitempty"`
	Write []string `json:"write,omitempty"`
}
