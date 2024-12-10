// 文件数据结构

package linker

import (
	"github.com/shaozk/rvld-learn/pkg/utils"
	"os"
)

// 文件
type File struct {
	Name string
	Contents []byte
}

// 创建新文件
// [in] filename 文件路径
// [out] 文件指针
func MustNewFile(filename string) *File {
	contents, err := os.ReadFile(filename)
	utils.MustNo(err)
	return &File{
		Name:	filename,
		Contents:	contents,
	}
}