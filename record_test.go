package kinto

import (
	"github.com/h2non/gock"
	"testing"
)

func TestKintoClient_CreateRecord(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	var res recordRes
	fixture("/records/record.json", &res)
	gock.New(TEST_BASE_URI).Post(kc.buildURI(RECORDS_URI, TEST_BUCKET, TEST_RECORD_ID)).Reply(200).JSON(res)

	gotRec, gotPerm, _ := kc.CreateRecord(TEST_BUCKET, TEST_RECORD_ID, res.Data)
	assertJSON(gotRec, res.Data, t)
	assertJSON(gotPerm, res.Perm, t)
}

func TestKintoClient_GetRecords(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	var res recordsRes
	fixture("/records/records.json", &res)
	gock.New(TEST_BASE_URI).Get(kc.buildURI(RECORDS_URI, TEST_BUCKET, TEST_COLLECTION)).Reply(200).JSON(res)

	got, _ := kc.GetRecords(TEST_BUCKET, TEST_COLLECTION, nil)
	assertJSON(got, res.Data, t)
}

func TestKintoClient_GetRecordsWithOpts(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	opts := Options{
		"field": "id",
		"sort":  "-id",
	}

	var res recordsRes
	fixture("/records/records.json", &res)
	gock.New(TEST_BASE_URI).Get(kc.buildURI(RECORDS_URI, TEST_BUCKET, TEST_COLLECTION)).MatchParams(opts).Reply(200).JSON(res)

	got, _ := kc.GetRecords(TEST_BUCKET, TEST_COLLECTION, opts)
	assertJSON(got, res.Data, t)
}

func TestKintoClient_GetRecord(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	var res recordRes
	fixture("/records/record.json", &res)
	gock.New(TEST_BASE_URI).Get(kc.buildURI(RECORD_URI, TEST_BUCKET, TEST_COLLECTION, TEST_RECORD_ID)).Reply(200).JSON(res)

	gotRec, gotPerm, _ := kc.GetRecord(TEST_BUCKET, TEST_COLLECTION, TEST_RECORD_ID)
	assertJSON(gotRec, res.Data, t)
	assertJSON(gotPerm, res.Perm, t)
}
