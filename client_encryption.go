package mongoifc

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// ClientEncryption is an interface for `mongo.ClientEncryption` structure
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#ClientEncryption
type ClientEncryption interface {
	AddKeyAltName(
		ctx context.Context,
		id bson.Binary,
		keyAltName string,
	) SingleResult
	Close(ctx context.Context) error
	CreateDataKey(
		ctx context.Context,
		kmsProvider string,
		opts ...options.Lister[options.DataKeyOptions],
	) (bson.Binary, error)
	CreateEncryptedCollection(
		ctx context.Context,
		db Database,
		coll string,
		createOpts options.Lister[options.CreateCollectionOptions],
		kmsProvider string,
		masterKey any,
	) (Collection, bson.M, error)
	Decrypt(ctx context.Context, val bson.Binary) (bson.RawValue, error)
	DeleteKey(ctx context.Context, id bson.Binary) (*mongo.DeleteResult, error)
	Encrypt(
		ctx context.Context,
		val bson.RawValue,
		opts ...options.Lister[options.EncryptOptions],
	) (bson.Binary, error)
	EncryptExpression(
		ctx context.Context,
		expr any,
		result any,
		opts ...options.Lister[options.EncryptOptions],
	) error
	GetKey(ctx context.Context, id bson.Binary) SingleResult
	GetKeyByAltName(ctx context.Context, keyAltName string) SingleResult
	GetKeys(ctx context.Context) (Cursor, error)
	RemoveKeyAltName(
		ctx context.Context,
		id bson.Binary,
		keyAltName string,
	) SingleResult
	RewrapManyDataKey(
		ctx context.Context,
		filter any,
		opts ...options.Lister[options.RewrapManyDataKeyOptions],
	) (*mongo.RewrapManyDataKeyResult, error)
}

type clientEncryption struct {
	ce *mongo.ClientEncryption
}

// AddKeyAltName is a wrapper for `mongo.ClientEncryption.AddKeyAltName` method
func (c *clientEncryption) AddKeyAltName(ctx context.Context, id bson.Binary, keyAltName string) SingleResult {
	return wrapSingleResult(c.ce.AddKeyAltName(ctx, id, keyAltName))
}

// Close is a wrapper for `mongo.ClientEncryption.Close` method
func (c *clientEncryption) Close(ctx context.Context) error {
	return c.ce.Close(ctx)
}

// CreateDataKey is a wrapper for `mongo.ClientEncryption.CreateDataKey` method
func (c *clientEncryption) CreateDataKey(
	ctx context.Context,
	kmsProvider string,
	opts ...options.Lister[options.DataKeyOptions],
) (bson.Binary, error) {
	return c.ce.CreateDataKey(ctx, kmsProvider, opts...)
}

// CreateEncryptedCollection is a wrapper for `mongo.ClientEncryption.CreateEncryptedCollection` method
func (c *clientEncryption) CreateEncryptedCollection(
	ctx context.Context,
	db Database,
	coll string,
	createOpts options.Lister[options.CreateCollectionOptions],
	kmsProvider string,
	masterKey any,
) (Collection, bson.M, error) {
	col, doc, err := c.ce.CreateEncryptedCollection(ctx, UnWrapDatabase(db), coll, createOpts, kmsProvider, masterKey)
	if err != nil {
		return nil, nil, err
	}
	return wrapCollection(col, db.(*database)), doc, err
}

// Decrypt is a wrapper for `mongo.ClientEncryption.Decrypt` method
func (c *clientEncryption) Decrypt(ctx context.Context, val bson.Binary) (bson.RawValue, error) {
	return c.ce.Decrypt(ctx, val)
}

// DeleteKey is a wrapper for `mongo.ClientEncryption.DeleteKey` method
func (c *clientEncryption) DeleteKey(ctx context.Context, id bson.Binary) (*mongo.DeleteResult, error) {
	return c.ce.DeleteKey(ctx, id)
}

// Encrypt is a wrapper for `mongo.ClientEncryption.Encrypt` method
func (c *clientEncryption) Encrypt(
	ctx context.Context,
	val bson.RawValue,
	opts ...options.Lister[options.EncryptOptions],
) (bson.Binary, error) {
	return c.ce.Encrypt(ctx, val, opts...)
}

// EncryptExpression is a wrapper for `mongo.ClientEncryption.EncryptExpression` method
func (c *clientEncryption) EncryptExpression(
	ctx context.Context,
	expr any,
	result any,
	opts ...options.Lister[options.EncryptOptions],
) error {
	return c.ce.EncryptExpression(ctx, expr, result, opts...)
}

// GetKey is a wrapper for `mongo.ClientEncryption.GetKey` method
func (c *clientEncryption) GetKey(ctx context.Context, id bson.Binary) SingleResult {
	return wrapSingleResult(c.ce.GetKey(ctx, id))
}

// GetKeyByAltName is a wrapper for `mongo.ClientEncryption.GetKeyByAltName` method
func (c *clientEncryption) GetKeyByAltName(ctx context.Context, keyAltName string) SingleResult {
	return wrapSingleResult(c.ce.GetKeyByAltName(ctx, keyAltName))
}

// GetKeys is a wrapper for `mongo.ClientEncryption.GetKeys` method
func (c *clientEncryption) GetKeys(ctx context.Context) (Cursor, error) {
	cr, err := c.ce.GetKeys(ctx)
	if err != nil {
		return nil, err
	}
	return wrapCursor(cr), nil
}

// RemoveKeyAltName is a wrapper for `mongo.ClientEncryption.RemoveKeyAltName` method
func (c *clientEncryption) RemoveKeyAltName(
	ctx context.Context,
	id bson.Binary,
	keyAltName string,
) SingleResult {
	return wrapSingleResult(c.ce.RemoveKeyAltName(ctx, id, keyAltName))
}

// RewrapManyDataKey is a wrapper for `mongo.ClientEncryption.RewrapManyDataKey` method
func (c *clientEncryption) RewrapManyDataKey(
	ctx context.Context,
	filter any,
	opts ...options.Lister[options.RewrapManyDataKeyOptions],
) (*mongo.RewrapManyDataKeyResult, error) {
	return c.ce.RewrapManyDataKey(ctx, filter, opts...)
}

// NewClientEncryption creates a new ClientEncryption
// Documentation: https://pkg.go.dev/go.mongodb.org/mongo-driver/v2/mongo#NewClientEncryption
func NewClientEncryption(
	keyVaultClient Client,
	opts ...options.Lister[options.ClientEncryptionOptions],
) (ClientEncryption, error) {
	ce, err := mongo.NewClientEncryption(UnWrapClient(keyVaultClient), opts...)
	if err != nil {
		return nil, err
	}
	return &clientEncryption{ce: ce}, nil
}
