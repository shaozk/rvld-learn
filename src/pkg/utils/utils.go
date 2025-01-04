// 工具类

package utils

import (
	"fmt"
	"runtime/debug"
	"os"
	"bytes"
	"encoding/binary"
	"strings"
)

// 致命错误
func Fatal(v any) {
	fmt.Printf("[rvld] \033[0;1;31mfatal:\033[0m %v\n", v)
	debug.PrintStack()
	os.Exit(1)
}

// 必须非空
func MustNo(err error) {
	if err != nil {
		Fatal(err)
	}
}

// 读取字节数据（字节->任意数据结构）
func Read[T any](data []byte) (val T) {
	reader := bytes.NewReader(data)
	err := binary.Read(reader, binary.LittleEndian, &val)
	MustNo(err)
	return
}

// 断言
func Assert(condition bool) {
	if !condition {
		Fatal("Assert failed")
	}
}

func RemovePrefix(s, prefix string) (string, bool) {
	if strings.HasPrefix(s, prefix) {
		s = strings.TrimPrefix(s, prefix)
		return s, true
	}
	return s, false
}
