package mongoifc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/sv-tools/mongoifc/v2"
)

func TestWrapDatabase_UnWrapDatabase(t *testing.T) {
	t.Parallel()
	cl := connect(t)
	mcl := mongoifc.UnWrapClient(cl)
	name := fmt.Sprintf("test_%d", time.Now().Unix())
	orig := mcl.Database(name)
	require.NoError(t, orig.CreateCollection(t.Context(), "orig"))
	wrapped := mongoifc.WrapDatabase(orig)
	require.NoError(t, wrapped.CreateCollection(t.Context(), "wrapped"))
	require.Equal(t, orig, mongoifc.UnWrapDatabase(wrapped))
}
