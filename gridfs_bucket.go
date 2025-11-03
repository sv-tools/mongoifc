package mongoifc

import (
	"context"
	"io"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// GridFSBucket is an interface for `mongo.GridFSBucket` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#GridFSBucket
type GridFSBucket interface {
	Delete(ctx context.Context, fileID any) error
	DownloadToStream(ctx context.Context, fileID any, stream io.Writer) (int64, error)
	DownloadToStreamByName(
		ctx context.Context,
		filename string,
		stream io.Writer,
		opts ...options.Lister[options.GridFSNameOptions],
	) (int64, error)
	Drop(ctx context.Context) error
	Find(
		ctx context.Context,
		filter any,
		opts ...options.Lister[options.GridFSFindOptions],
	) (Cursor, error)
	GetChunksCollection() Collection
	GetFilesCollection() Collection
	OpenDownloadStream(ctx context.Context, fileID any) (GridFSDownloadStream, error)
	OpenDownloadStreamByName(
		ctx context.Context,
		filename string,
		opts ...options.Lister[options.GridFSNameOptions],
	) (GridFSDownloadStream, error)
	OpenUploadStream(
		ctx context.Context,
		filename string,
		opts ...options.Lister[options.GridFSUploadOptions],
	) (GridFSUploadStream, error)
	OpenUploadStreamWithID(
		ctx context.Context,
		fileID any,
		filename string,
		opts ...options.Lister[options.GridFSUploadOptions],
	) (GridFSUploadStream, error)
	Rename(ctx context.Context, fileID any, newFilename string) error
	UploadFromStream(
		ctx context.Context,
		filename string,
		source io.Reader,
		opts ...options.Lister[options.GridFSUploadOptions],
	) (bson.ObjectID, error)
	UploadFromStreamWithID(
		ctx context.Context,
		fileID any,
		filename string,
		source io.Reader,
		opts ...options.Lister[options.GridFSUploadOptions],
	) error
}

type gridFSBucket struct {
	bt *mongo.GridFSBucket
}

// Delete is a wrapper for `mongo.GridFSBucket.Delete` method
func (g *gridFSBucket) Delete(ctx context.Context, fileID any) error {
	return g.bt.Delete(ctx, fileID)
}

// DownloadToStream is a wrapper for `mongo.GridFSBucket.DownloadToStream` method
func (g *gridFSBucket) DownloadToStream(ctx context.Context, fileID any, stream io.Writer) (int64, error) {
	return g.bt.DownloadToStream(ctx, fileID, stream)
}

// DownloadToStreamByName is a wrapper for `mongo.GridFSBucket.DownloadToStreamByName` method
func (g *gridFSBucket) DownloadToStreamByName(
	ctx context.Context,
	filename string,
	stream io.Writer,
	opts ...options.Lister[options.GridFSNameOptions],
) (int64, error) {
	return g.bt.DownloadToStreamByName(ctx, filename, stream, opts...)
}

// Drop is a wrapper for `mongo.GridFSBucket.Drop` method
func (g *gridFSBucket) Drop(ctx context.Context) error {
	return g.bt.Drop(ctx)
}

// Find is a wrapper for `mongo.GridFSBucket.Find` method
func (g *gridFSBucket) Find(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.GridFSFindOptions],
) (Cursor, error) {
	cr, err := g.bt.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	return wrapCursor(cr), nil
}

// GetChunksCollection is a wrapper for `mongo.GridFSBucket.GetChunksCollection` method
func (g *gridFSBucket) GetChunksCollection() Collection {
	return WrapCollection(g.bt.GetChunksCollection())
}

// GetFilesCollection is a wrapper for `mongo.GridFSBucket.GetFilesCollection` method
func (g *gridFSBucket) GetFilesCollection() Collection {
	return WrapCollection(g.bt.GetFilesCollection())
}

// OpenDownloadStream is a wrapper for `mongo.GridFSBucket.OpenDownloadStream` method
func (g *gridFSBucket) OpenDownloadStream(ctx context.Context, fileID any) (GridFSDownloadStream, error) {
	ds, err := g.bt.OpenDownloadStream(ctx, fileID)
	if err != nil {
		return nil, err
	}
	return wrapGridFSDownloadStream(ds), nil
}

// OpenDownloadStreamByName is a wrapper for `mongo.GridFSBucket.OpenDownloadStreamByName` method
func (g *gridFSBucket) OpenDownloadStreamByName(
	ctx context.Context,
	filename string,
	opts ...options.Lister[options.GridFSNameOptions],
) (GridFSDownloadStream, error) {
	ds, err := g.bt.OpenDownloadStreamByName(ctx, filename, opts...)
	if err != nil {
		return nil, err
	}
	return wrapGridFSDownloadStream(ds), nil
}

// OpenUploadStream is a wrapper for `mongo.GridFSBucket.OpenUploadStream` method
func (g *gridFSBucket) OpenUploadStream(
	ctx context.Context,
	filename string,
	opts ...options.Lister[options.GridFSUploadOptions],
) (GridFSUploadStream, error) {
	us, err := g.bt.OpenUploadStream(ctx, filename, opts...)
	if err != nil {
		return nil, err
	}
	return wrapGridFSUploadStream(us), nil
}

// OpenUploadStreamWithID is a wrapper for `mongo.GridFSBucket.OpenUploadStreamWithID` method
func (g *gridFSBucket) OpenUploadStreamWithID(
	ctx context.Context,
	fileID any,
	filename string,
	opts ...options.Lister[options.GridFSUploadOptions],
) (GridFSUploadStream, error) {
	us, err := g.bt.OpenUploadStreamWithID(ctx, fileID, filename, opts...)
	if err != nil {
		return nil, err
	}
	return wrapGridFSUploadStream(us), nil
}

// Rename is a wrapper for `mongo.GridFSBucket.Rename` method
func (g *gridFSBucket) Rename(ctx context.Context, fileID any, newFilename string) error {
	return g.bt.Rename(ctx, fileID, newFilename)
}

// UploadFromStream is a wrapper for `mongo.GridFSBucket.UploadFromStream` method
func (g *gridFSBucket) UploadFromStream(
	ctx context.Context,
	filename string,
	source io.Reader,
	opts ...options.Lister[options.GridFSUploadOptions],
) (bson.ObjectID, error) {
	return g.bt.UploadFromStream(ctx, filename, source, opts...)
}

// UploadFromStreamWithID is a wrapper for `mongo.GridFSBucket.UploadFromStreamWithID` method
func (g *gridFSBucket) UploadFromStreamWithID(
	ctx context.Context,
	fileID any,
	filename string,
	source io.Reader,
	opts ...options.Lister[options.GridFSUploadOptions],
) error {
	return g.bt.UploadFromStreamWithID(ctx, fileID, filename, source, opts...)
}

func wrapGridFSBucket(bt *mongo.GridFSBucket) GridFSBucket {
	return &gridFSBucket{bt: bt}
}
