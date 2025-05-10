package mongoifc

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// DistinctResult is an interface for `mongo.DistinctResult` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#DistinctResult
type DistinctResult interface {
	Decode(v any) error
	Err() error
	Raw() (bson.RawArray, error)
}

type distinctResult struct {
	dr *mongo.DistinctResult
}

func (s *distinctResult) Decode(v any) error {
	return s.dr.Decode(v)
}

// Raw returns the document represented by this SingleResult as a bson.Raw. If
// there was an error from the operation that created this SingleResult, both
// the result and that error will be returned. If the operation returned no
// documents, this will return (nil, ErrNoDocuments).
func (s *distinctResult) Raw() (bson.RawArray, error) {
	return s.dr.Raw()
}

func (s *distinctResult) Err() error {
	return s.dr.Err()
}

func wrapDistinctResult(dr *mongo.DistinctResult) DistinctResult {
	return &distinctResult{dr: dr}
}
