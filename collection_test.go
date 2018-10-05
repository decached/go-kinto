package kinto

import (
	"github.com/h2non/gock"
	"testing"
)

func TestKintoClient_CreateCollection(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	var res collectionRes
	fixture("/collections/collection.json", &res)
	gock.New(TEST_BASE_URI).Post(kc.buildURI(COLLECTIONS_URI, TEST_BUCKET)).Reply(200).JSON(res)

	gotCol, gotPerm, _ := kc.CreateCollection(TEST_BUCKET, TEST_COLLECTION)
	assertJSON(gotCol, res.Data, t)
	assertJSON(gotPerm, res.Perm, t)
}

func TestKintoClient_GetCollections(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	var res collectionsRes
	fixture("/collections/collections.json", &res)
	gock.New(TEST_BASE_URI).Get(kc.buildURI(COLLECTIONS_URI, TEST_BUCKET)).Reply(200).JSON(res)

	got, _ := kc.GetCollections(TEST_BUCKET, nil)
	assertJSON(got, res.Data, t)
}

func TestKintoClient_GetCollectionsWithOpts(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	opts := Options{
		"field": "id",
		"sort":  "-id",
	}

	var res collectionsRes
	fixture("/collections/collections.json", &res)
	gock.New(TEST_BASE_URI).Get(kc.buildURI(COLLECTIONS_URI, TEST_BUCKET)).MatchParams(opts).Reply(200).JSON(res)

	got, _ := kc.GetCollections(TEST_BUCKET, opts)
	assertJSON(got, res.Data, t)
}

func TestKintoClient_GetCollection(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	var res collectionRes
	fixture("/collections/collection.json", &res)
	gock.New(TEST_BASE_URI).Get(kc.buildURI(COLLECTION_URI, TEST_BUCKET, TEST_COLLECTION)).Reply(200).JSON(res)

	gotCol, gotPerm, _ := kc.GetCollection(TEST_BUCKET, TEST_COLLECTION)
	assertJSON(gotCol, res.Data, t)
	assertJSON(gotPerm, res.Perm, t)
}
