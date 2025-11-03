package mongoifc

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// SingleResult is an interface for `mongo.SingleResult` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#SingleResult
type SingleResult interface {
	Decode(v any) error
	Err() error
	Raw() (bson.Raw, error)
}

type singleResult struct {
	sr *mongo.SingleResult
}

// Decode is a wrapper for `mongo.SingleResult.Decode` method
func (s *singleResult) Decode(v any) error {
	return s.sr.Decode(v)
}

// Raw is a wrapper for `mongo.SingleResult.Raw` method
func (s *singleResult) Raw() (bson.Raw, error) {
	return s.sr.Raw()
}

// Err is a wrapper for `mongo.SingleResult.Err` method
func (s *singleResult) Err() error {
	return s.sr.Err()
}

func wrapSingleResult(sr *mongo.SingleResult) SingleResult {
	return &singleResult{sr: sr}
}

// NewSingleResultFromDocument is a wrapper for NewSingleResultFromDocument function of the mongodb
// to return SingleResult
// https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#NewSingleResultFromDocument
func NewSingleResultFromDocument(document any, err error, registry *bson.Registry) SingleResult {
	sr := mongo.NewSingleResultFromDocument(document, err, registry)
	return wrapSingleResult(sr)
}
