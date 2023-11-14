package mongoifc

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo"
)

// SingleResult is an interface for `mongo.SingleResult` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#SingleResult
type SingleResult interface {
	Decode(v interface{}) error
	DecodeBytes() (bson.Raw, error)
	Err() error
	Raw() (bson.Raw, error)
}

type singleResult struct {
	sr *mongo.SingleResult
}

func (s *singleResult) Decode(v interface{}) error {
	return s.sr.Decode(v)
}

// Raw returns the document represented by this SingleResult as a bson.Raw. If
// there was an error from the operation that created this SingleResult, both
// the result and that error will be returned. If the operation returned no
// documents, this will return (nil, ErrNoDocuments).
func (s *singleResult) Raw() (bson.Raw, error) {
	return s.sr.Raw()
}

// DecodeBytes will return the document represented by this SingleResult as a bson.Raw. If there was an error from the
// operation that created this SingleResult, both the result and that error will be returned. If the operation returned
// no documents, this will return (nil, ErrNoDocuments).
//
// Deprecated: Use [SingleResult.Raw] instead.
func (s *singleResult) DecodeBytes() (bson.Raw, error) {
	return s.sr.DecodeBytes()
}

func (s *singleResult) Err() error {
	return s.sr.Err()
}

func wrapSingleResult(sr *mongo.SingleResult) SingleResult {
	return &singleResult{sr: sr}
}

// NewSingleResultFromDocument is a wrapper for NewSingleResultFromDocument function of the mongodb
// to return SingleResult
// https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#NewSingleResultFromDocument
func NewSingleResultFromDocument(document interface{}, err error, registry *bsoncodec.Registry) SingleResult {
	sr := mongo.NewSingleResultFromDocument(document, err, registry)
	return wrapSingleResult(sr)
}
