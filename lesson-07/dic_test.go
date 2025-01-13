package dictionary

import "testing"

func TestGetDictionaryKey(t *testing.T) {
	want := "value"
	dictionary := Dictionary{"key": want}

	_, got := dictionary.Get("key")

	if got != want {
		t.Errorf("Got %s but expected %s", got, want)
	}
}

func TestAddToDictionary(t *testing.T) {
	t.Run("Add to dictionary", func(t *testing.T) {
		want := "value"
		key := "key"
		dictionary := Dictionary{}

		dictionary.Add(key, want)

		_, got := dictionary.Get(key)

		if got != want {
			t.Errorf("Got %s but expected %s", got, want)
		}
	})

	t.Run("Throw error when duplicated", func(t *testing.T) {
		want := "value"
		key := "key"
		dictionary := Dictionary{key: want}

		err := dictionary.Add(key, want)

		if err == nil {
			t.Errorf("Expected to receive error")
		}

	})
}
