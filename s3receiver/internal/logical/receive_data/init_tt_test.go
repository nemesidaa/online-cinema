package receivedata_test

import (
	receivedata "s3receiver/internal/logical/receive_data"
	"testing"
)

func TestReciveData1(t *testing.T) {
	t.Run("Normal word", func(t *testing.T) {
		entryData := []receivedata.Object{{0, 'l'}, {1, 'e'}, {2, 'g'}}
		current, err := receivedata.SearchSortByWord(entryData)
		if err != nil {
			t.Logf("Error, %s", err.Error())
		}
		excepted := "leg"
		if current != excepted {
			t.Logf("Wanted %s, got %s", excepted, current)
			t.FailNow()
		} else {
			t.Logf("current is %s. PASS", current)
		}
	})

	t.Run("Zero letters", func(t *testing.T) {
		entryData := []receivedata.Object{}
		current, err := receivedata.SearchSortByWord(entryData)
		if err != nil {
			t.Logf("Error, %s", err.Error())
		}
		excepted := ""
		if current != excepted {
			t.Logf("Wanted %s, got %s", excepted, current)
			t.FailNow()
		} else {
			t.Logf("current is %s. PASS", current)
		}
	})

	t.Run("One letter", func(t *testing.T) {
		entryData := []receivedata.Object{{0, 'X'}}
		current, err := receivedata.SearchSortByWord(entryData)
		if err != nil {
			t.Logf("Error, %s", err.Error())
		}
		excepted := "X"
		if current != excepted {
			t.Logf("Wanted %s, got %s", excepted, current)
			t.FailNow()
		} else {
			t.Logf("current is %s. PASS", current)
		}
	})
}
