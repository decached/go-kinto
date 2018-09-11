package kinto

import (
	"github.com/h2non/gock"
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

	got, _ := kc.GetBuckets(nil)
	assertJSON(got, res.Data, t)
}

func TestKintoClient_GetBucketsWithOpts(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	opts := Options{
		"sort":  "-id",
	}

	var res bucketsRes
	fixture("/buckets/buckets.json", &res)
	gock.New(TEST_BASE_URI).Get(kc.buildURI(BUCKETS_URI)).MatchParams(opts).Reply(200).JSON(res)

	got, _ := kc.GetBuckets(opts)
	assertJSON(got, res.Data, t)
}

func TestKintoClient_GetBucket(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	var res bucketRes
	fixture("/buckets/bucket.json", &res)
	gock.New(TEST_BASE_URI).Get(kc.buildURI(BUCKET_URI, TEST_BUCKET)).Reply(200).JSON(res)

	got, _ := kc.GetBucket(TEST_BUCKET, nil)
	assertJSON(got, res.Data, t)
}
