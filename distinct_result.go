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

// Decode is a wrapper for `mongo.DistinctResult.Decode` method
func (s *distinctResult) Decode(v any) error {
	return s.dr.Decode(v)
}

// Raw is a wrapper for `mongo.DistinctResult.Raw` method
func (s *distinctResult) Raw() (bson.RawArray, error) {
	return s.dr.Raw()
}

// Err is a wrapper for `mongo.DistinctResult.Err` method
func (s *distinctResult) Err() error {
	return s.dr.Err()
}

func wrapDistinctResult(dr *mongo.DistinctResult) DistinctResult {
	return &distinctResult{dr: dr}
}
