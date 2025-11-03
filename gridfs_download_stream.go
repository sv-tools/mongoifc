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

// Close is a wrapper for `mongo.GridFSDownloadStream.Close` method
func (stream *gridFSDownloadStream) Close() error {
	return stream.sm.Close()
}

// GetFile is a wrapper for `mongo.GridFSDownloadStream.GetFile` method
func (stream *gridFSDownloadStream) GetFile() *mongo.GridFSFile {
	return stream.sm.GetFile()
}

// Read is a wrapper for `mongo.GridFSDownloadStream.Read` method
func (stream *gridFSDownloadStream) Read(p []byte) (int, error) {
	return stream.sm.Read(p)
}

// Skip is a wrapper for `mongo.GridFSDownloadStream.Skip` method
func (stream *gridFSDownloadStream) Skip(skip int64) (int64, error) {
	return stream.sm.Skip(skip)
}

func wrapGridFSDownloadStream(sm *mongo.GridFSDownloadStream) GridFSDownloadStream {
	return &gridFSDownloadStream{sm: sm}
}
