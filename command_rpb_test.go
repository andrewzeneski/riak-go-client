package riak

import (
	rpbRiak "github.com/basho/riak-go-client/rpb/riak"
	rpbRiakDT "github.com/basho/riak-go-client/rpb/riak_dt"
	rpbRiakKV "github.com/basho/riak-go-client/rpb/riak_kv"
	rpbRiakSCH "github.com/basho/riak-go-client/rpb/riak_search"
	rpbRiakYZ "github.com/basho/riak-go-client/rpb/riak_yokozuna"
	proto "github.com/golang/protobuf/proto"
	"reflect"
	"testing"
)

func TestEnsureCorrectRequestAndResponseCodes(t *testing.T) {
	var cmd Command
	var msg proto.Message
	// Misc commands
	// Ping
	cmd = &PingCommand{}
	if expected, actual := rpbCode_RpbPingReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbPingResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if cmd.getResponseProtobufMessage() != nil {
		t.Error("expected nil response protobuf message")
	}
	// StartTls
	cmd = &startTlsCommand{}
	if expected, actual := rpbCode_RpbStartTls, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbStartTls, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if cmd.getResponseProtobufMessage() != nil {
		t.Error("expected nil response protobuf message")
	}
	// Auth
	cmd = &authCommand{}
	if expected, actual := rpbCode_RpbAuthReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbAuthResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if cmd.getResponseProtobufMessage() != nil {
		t.Error("expected nil response protobuf message")
	}
	// FetchBucketProps
	cmd = &FetchBucketPropsCommand{}
	if expected, actual := rpbCode_RpbGetBucketReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbGetBucketResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiak.RpbGetBucketResp); !ok {
		t.Errorf("error casting %v to RpbGetBucketResp", reflect.TypeOf(msg))
	}
	// StoreBucketProps
	cmd = &StoreBucketPropsCommand{}
	if expected, actual := rpbCode_RpbSetBucketReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbSetBucketResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if msg != nil {
		t.Error("expected nil response protobuf message")
	}

	// KV commands
	// FetchValue
	cmd = &FetchValueCommand{}
	if expected, actual := rpbCode_RpbGetReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbGetResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakKV.RpbGetResp); !ok {
		t.Errorf("error casting %v to RpbGetResp", reflect.TypeOf(msg))
	}
	// StoreValue
	cmd = &StoreValueCommand{}
	if expected, actual := rpbCode_RpbPutReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbPutResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakKV.RpbPutResp); !ok {
		t.Errorf("error casting %v to RpbPutResp", reflect.TypeOf(msg))
	}
	// DeleteValue
	cmd = &DeleteValueCommand{}
	if expected, actual := rpbCode_RpbDelReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbDelResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if msg != nil {
		t.Error("expected nil response protobuf message")
	}
	// ListBuckets
	cmd = &ListBucketsCommand{}
	if expected, actual := rpbCode_RpbListBucketsReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbListBucketsResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakKV.RpbListBucketsResp); !ok {
		t.Errorf("error casting %v to RpbListBucketsResp", reflect.TypeOf(msg))
	}
	// ListKeys
	cmd = &ListKeysCommand{}
	if expected, actual := rpbCode_RpbListKeysReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbListKeysResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakKV.RpbListKeysResp); !ok {
		t.Errorf("error casting %v to RpbListKeysResp", reflect.TypeOf(msg))
	}
	// FetchPreflist
	cmd = &FetchPreflistCommand{}
	if expected, actual := rpbCode_RpbGetBucketKeyPreflistReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbGetBucketKeyPreflistResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakKV.RpbGetBucketKeyPreflistResp); !ok {
		t.Errorf("error casting %v to RpbGetBucketKeyPreflistResp", reflect.TypeOf(msg))
	}
	// SecondaryIndexQuery
	cmd = &SecondaryIndexQueryCommand{}
	if expected, actual := rpbCode_RpbIndexReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbIndexResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakKV.RpbIndexResp); !ok {
		t.Errorf("error casting %v to RpbIndexResp", reflect.TypeOf(msg))
	}
	// MapReduce
	cmd = &MapReduceCommand{}
	if expected, actual := rpbCode_RpbMapRedReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbMapRedResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakKV.RpbMapRedResp); !ok {
		t.Errorf("error casting %v to RpbMapRedResp", reflect.TypeOf(msg))
	}

	// YZ commands
	// StoreIndex
	cmd = &StoreIndexCommand{}
	if expected, actual := rpbCode_RpbYokozunaIndexPutReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbPutResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if cmd.getResponseProtobufMessage() != nil {
		t.Error("expected nil response protobuf message")
	}
	// FetchIndex
	cmd = &FetchIndexCommand{}
	if expected, actual := rpbCode_RpbYokozunaIndexGetReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbYokozunaIndexGetResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakYZ.RpbYokozunaIndexGetResp); !ok {
		t.Errorf("error casting %v to RpbYokozunaIndexGetResp", reflect.TypeOf(msg))
	}
	// DeleteIndex
	cmd = &DeleteIndexCommand{}
	if expected, actual := rpbCode_RpbYokozunaIndexDeleteReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbDelResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if cmd.getResponseProtobufMessage() != nil {
		t.Error("expected nil response protobuf message")
	}
	// StoreSchema
	cmd = &StoreSchemaCommand{}
	if expected, actual := rpbCode_RpbYokozunaSchemaPutReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbPutResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if cmd.getResponseProtobufMessage() != nil {
		t.Error("expected nil response protobuf message")
	}
	// FetchSchema
	cmd = &FetchSchemaCommand{}
	if expected, actual := rpbCode_RpbYokozunaSchemaGetReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbYokozunaSchemaGetResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakYZ.RpbYokozunaSchemaGetResp); !ok {
		t.Errorf("error casting %v to RpbYokozunaSchemaGetResp", reflect.TypeOf(msg))
	}
	// Search
	cmd = &SearchCommand{}
	if expected, actual := rpbCode_RpbSearchQueryReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_RpbSearchQueryResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakSCH.RpbSearchQueryResp); !ok {
		t.Errorf("error casting %v to RpbSearchQueryResp", reflect.TypeOf(msg))
	}

	// CRDT commands
	// UpdateCounter
	cmd = &UpdateCounterCommand{}
	if expected, actual := rpbCode_DtUpdateReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_DtUpdateResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakDT.DtUpdateResp); !ok {
		t.Errorf("error casting %v to DtUpdateResp", reflect.TypeOf(msg))
	}
	// FetchCounter
	cmd = &FetchCounterCommand{}
	if expected, actual := rpbCode_DtFetchReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_DtFetchResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakDT.DtFetchResp); !ok {
		t.Errorf("error casting %v to DtFetchResp", reflect.TypeOf(msg))
	}
	// UpdateSet
	cmd = &UpdateSetCommand{}
	if expected, actual := rpbCode_DtUpdateReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_DtUpdateResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakDT.DtUpdateResp); !ok {
		t.Errorf("error casting %v to DtUpdateResp", reflect.TypeOf(msg))
	}
	// FetchSet
	cmd = &FetchSetCommand{}
	if expected, actual := rpbCode_DtFetchReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_DtFetchResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakDT.DtFetchResp); !ok {
		t.Errorf("error casting %v to DtFetchResp", reflect.TypeOf(msg))
	}
	// UpdateMap
	cmd = &UpdateMapCommand{}
	if expected, actual := rpbCode_DtUpdateReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_DtUpdateResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakDT.DtUpdateResp); !ok {
		t.Errorf("error casting %v to DtUpdateResp", reflect.TypeOf(msg))
	}
	// FetchMap
	cmd = &FetchMapCommand{}
	if expected, actual := rpbCode_DtFetchReq, cmd.getRequestCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	if expected, actual := rpbCode_DtFetchResp, cmd.getResponseCode(); expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
	msg = cmd.getResponseProtobufMessage()
	if _, ok := msg.(*rpbRiakDT.DtFetchResp); !ok {
		t.Errorf("error casting %v to DtFetchResp", reflect.TypeOf(msg))
	}
}
