package gcse

import (
	"github.com/daviddengcn/go-villa"
	"testing"
)

func TestMemDB_Bug_Sync(t *testing.T) {
	path := villa.Path(".").Join("testmemdb.gob")
	if path.Exists() {
		path.Remove()
	}

	db := NewMemDB(".", "testmemdb")
	db.Put("s", 1)
	err := db.Sync()
	if err != nil {
		t.Error(err)
	}

	villa.AssertEquals(t, "Exists", path.Exists(), true)
	if err := path.Remove(); err != nil {
		t.Error(err)
	}
	villa.AssertEquals(t, "Exists", path.Exists(), false)

	//if err := db.Load(); err != nil {
	//	t.Error(err)
	//}
}