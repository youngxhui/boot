package util

import (
	"fmt"
	"testing"
)

func TestGetConfPath(t *testing.T) {
	path := GetConfPath()
	fmt.Println(path)
}
