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

func (s *singleResult) Decode(v any) error {
	return s.sr.Decode(v)
}

// Raw returns the document represented by this SingleResult as a bson.Raw. If
// there was an error from the operation that created this SingleResult, both
// the result and that error will be returned. If the operation returned no
// documents, this will return (nil, ErrNoDocuments).
func (s *singleResult) Raw() (bson.Raw, error) {
	return s.sr.Raw()
}

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
