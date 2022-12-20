package common

import "encoding/json"

func MustMarshalIndent(v interface{}, prefix, indent string) []byte {
	res, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		panic(err)
	}
	return res
}
