package kinto

import (
	"bytes"
	"encoding/json"
)

type Collection struct {
	ID           string `json:"id"`
	LastModified int64  `json:"last_modified,omitempty"`
	CacheExpires int64  `json:"cache_expires,omitempty"`
}

type collectionReq struct {
	Data Collection `json:"data,omitempty"`
}

type collectionRes struct {
	Data        Collection `json:"data"`
	Permissions Perm       `json:"permissions"`
}

type collectionsRes struct {
	Data []Collection `json:"data"`
}

func (c KintoClient) CreateCollection(bucket string, collection string) (Collection, error) {
	path := c.buildURI(COLLECTIONS_URI, bucket)

	req := collectionReq{Data: Collection{ID: collection}}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		return Collection{}, err
	}

	var res collectionRes
	err = c.session.Request("POST", path, bytes.NewReader(reqJSON), &res)
	if err != nil {
		return Collection{}, err
	}
	return res.Data, nil
}

func (c *KintoClient) GetCollections(bucket string) ([]Collection, error) {
	path := c.buildURI(COLLECTIONS_URI, bucket)

	var res collectionsRes
	err := c.session.Request("GET", path, nil, &res)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (c KintoClient) GetCollection(bucket string, collection string) (Collection, error) {
	path := c.buildURI(COLLECTION_URI, bucket, collection)

	var res collectionRes
	err := c.session.Request("GET", path, nil, &res)
	if err != nil {
		return Collection{}, err
	}
	return res.Data, nil
}
