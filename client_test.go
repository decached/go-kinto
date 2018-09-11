package kinto

import (
	"encoding/json"
	"fmt"
	"github.com/h2non/gock"
	"io/ioutil"
	"reflect"
	"testing"
)

const (
	TEST_USER       = "user"
	TEST_PASSWORD   = "password"
	TEST_BASE_URI   = "http://localhost:8888/v1"
	TEST_BUCKET     = "test-bucket"
	TEST_COLLECTION = "test-collection"
	TEST_RECORD_ID  = "66b18d0c-a3a8-4be7-8ce0-1aceac89ff75"
)

func assert(got interface{}, want interface{}, t *testing.T) {
	if !reflect.DeepEqual(got, want) {
		t.Error(fmt.Sprintf("Want %s, Got %s", want, got))
	}
}

func assertJSON(got interface{}, want interface{}, t *testing.T) {
	gotS, _ := json.Marshal(got)
	wantS, _ := json.Marshal(want)
	if string(gotS) != string(wantS) {
		t.Error(fmt.Sprintf("Want %s, Got %s", want, got))
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func fixture(file string, target interface{}) {
	b1, _ := ioutil.ReadFile("fixtures" + file)
	json.Unmarshal(b1, &target)
}

func KintoClientSetup() *KintoClient {
	kc, err := NewClient(TEST_BASE_URI, TEST_USER, TEST_PASSWORD)
	checkErr(err)
	return kc
}

func KintoClientTearDown(kc *KintoClient) {
	gock.Off()
}

func TestKintoClient_HeartBeat(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	var want interface{}
	fixture("/heartbeat.json", &want)
	gock.New(TEST_BASE_URI).Get(kc.buildURI(HEARTBEAT_URI)).Reply(200).JSON(want)

	got, _ := kc.HeartBeat()
	assertJSON(got, want, t)
}

func TestKintoClient_Flush(t *testing.T) {
	kc := KintoClientSetup()
	defer KintoClientTearDown(kc)

	var want interface{}
	gock.New(TEST_BASE_URI).Get(kc.buildURI(FLUSH_URI)).Reply(200).JSON(want)

	got, _ := kc.Flush()
	assertJSON(got, want, t)
}