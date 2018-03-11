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
	Data        Bucket `json:"data"`
	Permissions Perm   `json:"permissions"`
}

type bucketsRes struct {
	Data []Bucket `json:"data"`
}

func (c KintoClient) CreateBucket(bucket string) (Bucket, error) {
	path := c.buildURI(BUCKETS_URI)

	req := bucketReq{Data: Bucket{ID: bucket}}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		return Bucket{}, err
	}

	var res bucketRes
	err = c.session.Request("POST", path, bytes.NewReader(reqJSON), &res)
	if err != nil {
		return Bucket{}, err
	}
	return res.Data, nil
}

func (c KintoClient) GetBuckets() ([]Bucket, error) {
	path := c.buildURI(BUCKETS_URI)

	var res bucketsRes
	err := c.session.Request("GET", path, nil, &res)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (c KintoClient) GetBucket(bucket string) (Bucket, error) {
	path := c.buildURI(BUCKET_URI, bucket)

	var res bucketRes
	err := c.session.Request("GET", path, nil, &res)
	if err != nil {
		return Bucket{}, err
	}
	return res.Data, nil
}
