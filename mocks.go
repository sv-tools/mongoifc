//go:generate go tool mockgen -destination=mocks/gomock/mocks.go -package mocks . ChangeStream,Client,ClientEncryption,Collection,Cursor,Database,DistinctResult,GridFSBucket,GridFSDownloadStream,GridFSUploadStream,IndexView,SearchIndexView,Session,SingleResult
//go:generate go tool mockery --all --with-expecter --srcpkg github.com/sv-tools/mongoifc/v2 --output mocks/mockery --disable-version-string --case underscore
package mongoifc
