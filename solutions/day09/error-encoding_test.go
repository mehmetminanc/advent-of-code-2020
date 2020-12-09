package main

import (
	"strconv"
	"testing"
)

func TestInt64(t *testing.T) {
	i := int64(36861719751332)
	parsed, _ := strconv.ParseInt("36861719751332", 10, 64)
	if i != parsed {
		t.Errorf("Overflowing: 36861719751332")
	}
}
