package utils

import (
	"fmt"
	"strings"
)

func PathMethod(method string) string {
	return strings.ToUpper(fmt.Sprintf("%s:%s", method, "SEQUENCE"))
}
