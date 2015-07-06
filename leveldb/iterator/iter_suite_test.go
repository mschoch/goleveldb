package iterator_test

import (
	"testing"

	"github.com/mschoch/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
