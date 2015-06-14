// +build darwin
package osarch

import (
	"fmt"
	"runtime"
)

func OsArch() string {
	s := fmt.Sprint(runtime.GOOS, "/", runtime.GOARCH)
	return s
}
