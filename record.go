package kinto

import (
	"bytes"
	"encoding/json"
)

type Record interface{}

type recordReq struct {
	Data Record `json:"data,omitempty"`
}

type recordRes struct {
	Data Record     `json:"data"`
	Perm Permission `json:"permissions"`
}

type recordsRes struct {
	Data []Record `json:"data"`
}

func (kc KintoClient) CreateRecord(bucket string, collection string, data Record) (Record, Permission, error) {
	path := kc.buildURI(RECORDS_URI, bucket, collection)

	req := recordReq{Data: data}
	reqJSON, err := json.Marshal(req)
	if err != nil {
		return 0, Permission{}, err
	}

	var res recordRes
	err = kc.session.Request("POST", path, nil, bytes.NewReader(reqJSON), &res)
	if err != nil {
		return 0, Permission{}, err
	}
	return res.Data, res.Perm, nil
}

func (kc KintoClient) GetRecords(bucket string, collection string, opts Options) ([]Record, error) {
	path := kc.buildURI(RECORDS_URI, bucket, collection)

	var res recordsRes
	err := kc.session.Request("GET", path, opts, nil, &res)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (kc KintoClient) GetRecord(bucket string, collection string, recordID string) (Record, Permission, error) {
	path := kc.buildURI(RECORD_URI, bucket, collection, recordID)

	var res recordRes
	err := kc.session.Request("GET", path, nil, nil, &res)
	if err != nil {
		return 0, Permission{}, err
	}
	return res.Data, res.Perm, nil
}
