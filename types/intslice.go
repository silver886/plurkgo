package types

import (
	"strconv"
	"strings"
)

// IntSlice is slice of int responsed by Plurk API 2.0.
// Its format is `|int|...`
type IntSlice []int

// UnmarshalJSON implement Unmarshaler for slice of int from Plurk API 2.0.
func (i *IntSlice) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), "\"")

	if str == "null" {
		i = nil
		return nil
	}

	for _, itemStr := range strings.Split(strings.Trim(str, "|"), "||") {
		if item, err := strconv.Atoi(itemStr); err != nil {
			return err
		} else {
			*i = IntSlice(append([]int(*i), item))
		}
	}

	return nil
}
