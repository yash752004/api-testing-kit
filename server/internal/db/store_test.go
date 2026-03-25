package db

import "testing"

func TestNewStoreNilPool(t *testing.T) {
	t.Parallel()

	if store := NewStore(nil); store != nil {
		t.Fatalf("expected nil store when pool is nil")
	}
}
