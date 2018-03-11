package kinto

type Perm struct {
	Read  []string `json:"read"`
	Write []string `json:"write"`
}
