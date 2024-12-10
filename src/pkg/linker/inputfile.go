// 输入文件数据结构

package linker

import (
	"github.com/shaozk/rvld-learn/pkg/utils"

)

// 输入文件
type InputFile struct {
	File	*File			///< 文件指针
	ElfSections	[] Shdr		///< 区段头部表
}


// 创建输入文件
// [in] file 文件指针
// [out] 输入文件
func NewInputFile(file *File) InputFile {
	f := InputFile{File : file}
	if len(file.Contents) < EhdrSize {
		utils.Fatal("File too small")
	}
	
	if !CheckMagic(file.Contents) {
		utils.Fatal("Not a ELF file")
	}

	ehdr := utils.Read[Ehdr](file.Contents)
	contents := file.Contents[ehdr.ShOff:]
	shdr := utils.Read[Shdr](contents)

	numSections := int64(ehdr.ShNum)
	if numSections == 0 {
		numSections = int64(shdr.Size)
	}
	
	f.ElfSections = []Shdr{shdr}
	for numSections > 1 {
		contents = contents[ShdrSize:]
		f.ElfSections = append(f.ElfSections, utils.Read[Shdr](contents))
		numSections--
	}
	return f
}