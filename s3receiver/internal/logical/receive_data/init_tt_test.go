package receivedata_test

import (
	receivedata "s3receiver/internal/logical/receive_data"
	"testing"
)

func TestReciveData1(t *testing.T) {
	t.Run("Custom run 1", func(t *testing.T) {
		entryData := []receivedata.Object{{1, 'X'}, {2, 'U'}, {3, 'I'}}
		if receivedata.SearchSortByWord(entryData) != "XUI" {
			excepted := "XUI"
			current := receivedata.SearchSortByWord(entryData)
			t.Logf("Wanded %s, got %s", excepted, current)
			t.FailNow()
		}
	})
}
func TestReciveData2(t *testing.T) {
	t.Run("Custom run 1", func(t *testing.T) {
		entryData := []receivedata.Object{{1, 'X'}}
		current := receivedata.SearchSortByWord(entryData)
		excepted := "X"
		if current != excepted {
			t.Logf("Wanted %s, got %s", excepted, current)
			t.FailNow()
		} else {
			t.Logf("current is %s. PASS", current)
		}
	})
}
func TestReciveData3(t *testing.T) {
	t.Run("Custom run 1", func(t *testing.T) {
		entryData := []receivedata.Object{}
		current := receivedata.SearchSortByWord(entryData)
		excepted := ""
		if current != excepted {
			t.Logf("Wanted %s, got %s", excepted, current)
			t.FailNow()
		} else {
			t.Logf("current is %s. PASS", current)
		}
	})
}
