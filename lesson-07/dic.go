package dictionary

import "errors"

type Dictionary map[string]string

func (dictionary Dictionary) Get(key string) (error, string) {
	value, ok := dictionary[key]

	if ok {
		return nil, value
	}

	return errors.New("Key not found"), ""
}

func (dictionary Dictionary) Add(key string, value string) error {
	keyNotFoundError, _ := dictionary.Get(key)
	if keyNotFoundError != nil {
		dictionary[key] = value
		return nil
	}
	return errors.New("The key already exists")
}
