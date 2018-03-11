package kinto

import (
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestKintoClient_CreateBucket(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	var res bucketRes
	fixture("/buckets/bucket.json", &res)
	gock.New(TEST_BASE_URI).Post(kc.buildURI(BUCKETS_URI)).Reply(200).JSON(res)

	got, _ := kc.CreateBucket(TEST_BUCKET)
	assertJSON(got, res.Data, t)
}

func TestKintoClient_GetBuckets(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	var res bucketsRes
	fixture("/buckets/buckets.json", &res)
	gock.New(TEST_BASE_URI).Get(kc.buildURI(BUCKETS_URI)).Reply(200).JSON(res)

	got, _ := kc.GetBuckets()
	assertJSON(got, res.Data, t)
}

func TestKintoClient_GetBucket(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	var res bucketRes
	fixture("/buckets/bucket.json", &res)
	gock.New(TEST_BASE_URI).Get(kc.buildURI(BUCKET_URI, TEST_BUCKET)).Reply(200).JSON(res)

	got, _ := kc.GetBucket(TEST_BUCKET)
	assertJSON(got, res.Data, t)
}
