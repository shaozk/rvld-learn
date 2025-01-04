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
	Parent *File
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

func OpenLibrary(filepath string) *File {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil
	} 

	return &File{
		Name:		filepath,
		Contents:	contents,
	}
}

func FindLibrary(ctx *Context, name string) *File {
	for _, dir := range ctx.Args.LibraryPaths {
		stem := dir + "/lib" + name + ".a"
		if f := OpenLibrary(stem); f != nil {
			return f
		}
	}

	utils.Fatal("library no find")
	return nil
}