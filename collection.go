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
	Data Collection `json:"data"`
	Perm Permission `json:"permissions"`
}

type collectionsRes struct {
	Data []Collection `json:"data"`
}

func (kc KintoClient) CreateCollection(bucket string, collection string) (Collection, Permission, error) {
	path := kc.buildURI(COLLECTIONS_URI, bucket)

	req := collectionReq{Data: Collection{ID: collection}}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		return Collection{}, Permission{}, err
	}

	var res collectionRes
	err = kc.session.Request("POST", path, nil, bytes.NewReader(reqJSON), &res)
	if err != nil {
		return Collection{}, Permission{}, err
	}
	return res.Data, res.Perm, nil
}

func (kc KintoClient) GetCollections(bucket string, opts Options) ([]Collection, error) {
	path := kc.buildURI(COLLECTIONS_URI, bucket)

	var res collectionsRes
	err := kc.session.Request("GET", path, opts, nil, &res)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (kc KintoClient) GetCollection(bucket string, collection string) (Collection, Permission, error) {
	path := kc.buildURI(COLLECTION_URI, bucket, collection)

	var res collectionRes
	err := kc.session.Request("GET", path, nil, nil, &res)
	if err != nil {
		return Collection{}, Permission{}, err
	}
	return res.Data, res.Perm, nil
}
