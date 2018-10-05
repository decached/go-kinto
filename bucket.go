package kinto

import (
	"bytes"
	"encoding/json"
)

type Bucket struct {
	ID           string `json:"id"`
	LastModified int64  `json:"last_modified,omitempty"`
}

type bucketReq struct {
	Data Bucket `json:"data,omitempty"`
}

type bucketRes struct {
	Data Bucket     `json:"data"`
	Perm Permission `json:"permissions"`
}

type bucketsRes struct {
	Data []Bucket `json:"data"`
}

func (kc KintoClient) CreateBucket(bucket string) (Bucket, Permission, error) {
	path := kc.buildURI(BUCKETS_URI)

	req := bucketReq{Data: Bucket{ID: bucket}}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		return Bucket{}, Permission{}, err
	}

	var res bucketRes
	err = kc.session.Request("POST", path, nil, bytes.NewReader(reqJSON), &res)
	if err != nil {
		return Bucket{}, Permission{}, err
	}
	return res.Data, res.Perm, nil
}

func (kc KintoClient) GetBuckets(opts Options) ([]Bucket, error) {
	path := kc.buildURI(BUCKETS_URI)

	var res bucketsRes
	err := kc.session.Request("GET", path, opts, nil, &res)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (kc KintoClient) GetBucket(bucket string, opts Options) (Bucket, Permission, error) {
	path := kc.buildURI(BUCKET_URI, bucket)

	var res bucketRes
	err := kc.session.Request("GET", path, opts, nil, &res)
	if err != nil {
		return Bucket{}, Permission{}, err
	}
	return res.Data, res.Perm, nil
}
