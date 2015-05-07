package tplfunc

import (
	"strconv"
)

func Intval(v interface{}) int {
	if r, ok := v.(string); ok {
		res, err := strconv.Atoi(r)
		if err != nil {
			return 0
		} else {
			return res
		}
	}

	if r, ok := v.(int); ok {
		return r
	}
	return 0
}
