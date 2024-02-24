package kvstore_test

import (
	"testing"

	"github.com/guraslan/kvstore"
)

func generateDB(s kvstore.Store) (kvstore.Store, error) {
	s, err := s.Store("golang", "1.20")
	if err != nil {
		return s, err
	}

	s, err = s.Store("istio", "1.21")
	if err != nil {
		return s, err
	}

	s, err = s.Store("k8s", "1.29")
	if err != nil {
		return s, err
	}

	return s, nil
}

func Test_RetrieveReturnsExistingKey(t *testing.T) {
	s, err := kvstore.Open("testdata/global.db")
	if err != nil {
		t.Error(err)
	}
	got := s.Retrieve("golang")
	want := "1.21"

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func Test_RetrieveReturnsEmptyIfKeyNotExists(t *testing.T) {
	s, err := kvstore.Open("testdata/global.db")
	if err != nil {
		t.Error(err)
	}
	got := s.Retrieve("nonexistent")
	want := ""

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

func Test_StoreSavesNewKeyValue(t *testing.T) {
	s, err := kvstore.Open(t.TempDir() + "/global.db")
	if err != nil {
		t.Error(err)
	}
	want := "ABC12345"
	s, err = s.Store("couponcode", want)
	if err != nil {
		t.Error(err)
	}
	got := s.Retrieve("couponcode")
	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}

}

func Test_StoreIncrementsRevisionNumber(t *testing.T) {
	s, err := kvstore.Open(t.TempDir() + "/global.db")
	if err != nil {
		t.Error(err)
	}

	key := "dostoyevsky"
	s, err = s.Store(key, "KaramazovBrothers")
	if err != nil {
		t.Error(err)
	}
	s, err = s.Store(key, "NotesFromUnderground")
	if err != nil {
		t.Error(err)
	}
	i, err := s.FindKV(key)
	if err != nil {
		t.Error(err)
	}
	got := s.Data[i].Revision
	want := 1

	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func Test_StoreValueWithSpaceNotCorruptDB(t *testing.T) {
	s, err := kvstore.Open(t.TempDir() + "/global.db")
	if err != nil {
		t.Error(err)
	}

	s, err = generateDB(s)
	if err != nil {
		t.Error(err)
	}

	s, err = s.Store("golang", "b c")
	if err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err)
	}
	got := len(s.Data)
	want := 3

	t.Log(s.Data)
	if want != got {

		t.Errorf("want %v got %v", want, got)
	}
}

func Test_StoreValueAfterValueWithSpaceNotCorruptDB(t *testing.T) {
	s, err := kvstore.Open(t.TempDir() + "/global.db")
	if err != nil {
		t.Error(err)
	}

	s, err = generateDB(s)
	if err != nil {
		t.Error(err)
	}

	s, err = s.Store("k8s", "1 29")
	if err != nil {
		t.Error(err)
	}
	
	s, err = s.Store("test", "1 23")
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error(err)
	}
	got := len(s.Data)
	want := 4

	t.Log(s.Data)
	if want != got {

		t.Errorf("want %v got %v", want, got)
	}
}
