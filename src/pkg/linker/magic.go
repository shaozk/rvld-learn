// 魔术值检查

package linker

import (
	"bytes"
)

// 检查是否为ELF文件（判断首四个字节是否为"\177ELF"）
func CheckMagic(contents []byte) bool {
	return bytes.HasPrefix(contents, []byte("\177ELF"))
}