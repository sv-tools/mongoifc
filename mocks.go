//go:generate go tool mockgen -destination=mocks/gomock/mocks.go -package mocks . ChangeStream,Client,Collection,Cursor,Database,IndexView,Session,SingleResult,SessionContext,ClientEncryption
//go:generate go tool mockery --all --with-expecter --srcpkg github.com/sv-tools/mongoifc --output mocks/mockery --disable-version-string --case underscore
package mongoifc
