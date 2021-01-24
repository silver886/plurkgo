package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

// IntSlice is slice of int and can be unmarshal from null.
type IntSlice []int

// UnmarshalJSON implement Unmarshaler for slice of int.
func (i *IntSlice) UnmarshalJSON(b []byte) error {
	fmt.Println(strings.Trim(string(b), "\""))
	fmt.Println(string(b))
	if strings.Trim(string(b), "\"") == "null" {
		i = nil
		return nil
	}
	return json.Unmarshal(b, i)
}
