package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SearchIndexView is an interface for `mongo.SearchIndexView` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#SearchIndexView
type SearchIndexView interface {
	CreateMany(
		ctx context.Context,
		models []mongo.SearchIndexModel,
		opts ...*options.CreateSearchIndexesOptions,
	) ([]string, error)
	CreateOne(
		ctx context.Context,
		model mongo.SearchIndexModel,
		opts ...*options.CreateSearchIndexesOptions,
	) (string, error)
	DropOne(
		ctx context.Context,
		name string,
		opts ...*options.DropSearchIndexOptions,
	) error
	List(
		ctx context.Context,
		searchIdxOpts *options.SearchIndexesOptions,
		opts ...*options.ListSearchIndexesOptions,
	) (Cursor, error)
	UpdateOne(
		ctx context.Context,
		name string,
		definition interface{},
		opts ...*options.UpdateSearchIndexOptions,
	) error
}

type searchIndexView struct {
	siv *mongo.SearchIndexView
}

func (i *searchIndexView) CreateMany(
	ctx context.Context,
	models []mongo.SearchIndexModel,
	opts ...*options.CreateSearchIndexesOptions,
) ([]string, error) {
	return i.siv.CreateMany(ctx, models, opts...)
}

func (i *searchIndexView) CreateOne(
	ctx context.Context,
	model mongo.SearchIndexModel,
	opts ...*options.CreateSearchIndexesOptions,
) (string, error) {
	return i.siv.CreateOne(ctx, model, opts...)
}

func (i *searchIndexView) DropOne(
	ctx context.Context, name string,
	opts ...*options.DropSearchIndexOptions,
) error {
	return i.siv.DropOne(ctx, name, opts...)
}

func (i *searchIndexView) List(
	ctx context.Context,
	searchIdxOpts *options.SearchIndexesOptions,
	opts ...*options.ListSearchIndexesOptions,
) (Cursor, error) {
	cr, err := i.siv.List(ctx, searchIdxOpts, opts...)
	if err != nil {
		return nil, err
	}

	return wrapCursor(cr), nil
}

func (i *searchIndexView) UpdateOne(
	ctx context.Context,
	name string,
	definition interface{},
	opts ...*options.UpdateSearchIndexOptions,
) error {
	return i.siv.UpdateOne(ctx, name, definition, opts...)
}

func wrapSearchIndexView(siv *mongo.SearchIndexView) SearchIndexView {
	return &searchIndexView{siv: siv}
}
