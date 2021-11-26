package mongoifc

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SingleResult is an interface for `mongo.SingleResult` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#SingleResult
type SingleResult interface {
	Decode(v interface{}) error
	DecodeBytes() (bson.Raw, error)
	Err() error
}

type singleResult struct {
	sr *mongo.SingleResult
}

func (s *singleResult) Decode(v interface{}) error {
	return s.sr.Decode(v)
}

func (s *singleResult) DecodeBytes() (bson.Raw, error) {
	return s.sr.DecodeBytes()
}

func (s *singleResult) Err() error {
	return s.sr.Err()
}

func wrapSingleResult(sr *mongo.SingleResult) SingleResult {
	return &singleResult{sr: sr}
}
