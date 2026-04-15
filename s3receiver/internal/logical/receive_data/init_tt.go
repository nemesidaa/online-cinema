package receivedata

import "errors"

type Object struct {
	Index  int
	Letter byte
}

func SearchSortByWord(f []Object) (string, error) {
	res := make([]byte, len(f))
	for _, v := range f {
		if v.Index > len(f)-1 {
			return "", errors.New("Outranged")
		}
		res[v.Index] = v.Letter
	}

	return string(res), nil
}
