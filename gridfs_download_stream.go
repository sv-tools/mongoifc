package mongoifc

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// GridFSDownloadStream is an interface for `mongo.GridFSDownloadStream` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#GridFSDownloadStream
type GridFSDownloadStream interface {
	Close() error
	GetFile() *mongo.GridFSFile
	Read(p []byte) (int, error)
	Skip(skip int64) (int64, error)
}

type gridFSDownloadStream struct {
	sm *mongo.GridFSDownloadStream
}

func (stream *gridFSDownloadStream) Close() error {
	return stream.sm.Close()
}

func (stream *gridFSDownloadStream) GetFile() *mongo.GridFSFile {
	return stream.sm.GetFile()
}

func (stream *gridFSDownloadStream) Read(p []byte) (int, error) {
	return stream.sm.Read(p)
}

func (stream *gridFSDownloadStream) Skip(skip int64) (int64, error) {
	return stream.sm.Skip(skip)
}

func wrapGridFSDownloadStream(sm *mongo.GridFSDownloadStream) GridFSDownloadStream {
	return &gridFSDownloadStream{sm: sm}
}
