package mongoifc_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/sv-tools/mongoifc"
)

func TestWrapDatabase_UnWrapDatabase(t *testing.T) {
	t.Parallel()
	cl := connect(t)
	mcl := mongoifc.UnWrapClient(cl)
	name := fmt.Sprintf("test_%d", time.Now().Unix())
	orig := mcl.Database(name)
	require.NoError(t, orig.CreateCollection(context.Background(), "orig"))
	wrapped := mongoifc.WrapDatabase(orig)
	require.NoError(t, wrapped.CreateCollection(context.Background(), "wrapped"))
	require.Equal(t, orig, mongoifc.UnWrapDatabase(wrapped))
}
