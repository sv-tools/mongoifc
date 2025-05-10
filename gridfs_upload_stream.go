package mongoifc

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// GridFSUploadStream is an interface for `mongo.GridFSUploadStream` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#GridFSUploadStream
type GridFSUploadStream interface {
	Abort() error
	Close() error
	Write(p []byte) (int, error)

	GetFileID() any
	SetFileID(fileID any)
}

type gridFSUploadStream struct {
	sm *mongo.GridFSUploadStream
}

func (s *gridFSUploadStream) Abort() error {
	return s.sm.Abort()
}

func (s *gridFSUploadStream) Close() error {
	return s.sm.Close()
}

func (s *gridFSUploadStream) Write(p []byte) (int, error) {
	return s.sm.Write(p)
}

func (s *gridFSUploadStream) GetFileID() any {
	return s.sm.FileID
}

func (s *gridFSUploadStream) SetFileID(fileID any) {
	s.sm.FileID = fileID
}

func wrapGridFSUploadStream(sm *mongo.GridFSUploadStream) GridFSUploadStream {
	return &gridFSUploadStream{sm: sm}
}
