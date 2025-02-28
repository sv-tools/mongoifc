package mongoifc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/sv-tools/mongoifc"
)

func TestWrapCollection_UnWrapCollection(t *testing.T) {
	t.Parallel()
	cl := connect(t)
	mcl := mongoifc.UnWrapClient(cl)
	name := fmt.Sprintf("test_%d", time.Now().Unix())
	db := mcl.Database(name)
	orig := db.Collection("test")
	_, err := orig.InsertOne(t.Context(), bson.M{"orig": "foo"})
	require.NoError(t, err)
	wrapped := mongoifc.WrapCollection(orig)
	_, err = wrapped.InsertOne(t.Context(), bson.M{"wrapped": "foo"})
	require.NoError(t, err)
	require.Equal(t, orig, mongoifc.UnWrapCollection(wrapped))
}
