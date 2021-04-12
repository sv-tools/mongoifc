package mongoifc

import "go.mongodb.org/mongo-driver/mongo"

// IndexView is an interface for `mongo.IndexView` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#IndexView
type IndexView interface {
	WrappedIndexView() *mongo.IndexView
}

type indexView struct {
	iv *mongo.IndexView
}

func (i *indexView) WrappedIndexView() *mongo.IndexView {
	return i.iv
}

func wrapIndexView(iv *mongo.IndexView) IndexView {
	return &indexView{iv: iv}
}
