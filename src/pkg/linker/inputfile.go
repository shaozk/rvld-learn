// 输入文件数据结构

package linker

import (
	"github.com/shaozk/rvld-learn/pkg/utils"
	"debug/elf"
	"fmt"

)

// 输入文件
type InputFile struct {
	File	*File			///< 文件指针
	ElfSections	[]Shdr		///< 区段头部表
	ElfSyms []Sym 			///< 符号表
	FirstGlobal int64
	ShStrtab	[]byte
	SymbolStrtab []byte
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

	shstrndx := int64(ehdr.ShStrndx)
	if (ehdr.ShStrndx == uint16(elf.SHN_XINDEX)) {
		shstrndx = int64(shdr.Link)
	}

	f.ShStrtab = f.GetBytesFromIdx(shstrndx)
	return f
}

func (f *InputFile) GetBytesFromShdr(s *Shdr) []byte {
	end := s.Offset + s.Size
	if uint64(len(f.File.Contents)) < end {
		utils.Fatal(fmt.Sprintf("section header is out of range: %d", s.Offset))
	}
	return f.File.Contents[s.Offset : end]
}

func (f *InputFile) GetBytesFromIdx(idx int64) []byte {
	return f.GetBytesFromShdr(&f.ElfSections[idx])
}

func (f *InputFile) FillUpElfSyms(s *Shdr) {
	bs := f.GetBytesFromShdr(s)
	nums := len(bs) / SymSize
	f.ElfSyms = make([]Sym, 0, nums)
	for nums > 0 {
		f.ElfSyms = append(f.ElfSyms, utils.Read[Sym](bs))
		bs = bs[SymSize:]
		nums--
	}
}

func (f *InputFile) FindSection(ty uint32) *Shdr {
	for i := 0; i < len(f.ElfSections); i++ {
		shdr := &f.ElfSections[i]
		if shdr.Type == ty {
			return shdr
		}
	}
	return nil
}