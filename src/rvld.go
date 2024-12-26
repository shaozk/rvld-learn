// 主函数

package main

import (
	"github.com/shaozk/rvld-learn/pkg/linker"
	"github.com/shaozk/rvld-learn/pkg/utils"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		utils.Fatal("Wrong args")
	}
	
	filename := os.Args[1]
	file := linker.MustNewFile(filename)
	objFile := linker.NewObjectFile(file)
	objFile.Parse()
	utils.Assert(len(objFile.ElfSections) == 11)
	utils.Assert(objFile.FirstGlobal ==  10)
	utils.Assert(len(objFile.ElfSyms) == 12)
	
	for _, sym := range objFile.ElfSyms {
		println(linker.ElfGetName(objFile.SymbolStrtab, sym.Name))
	}

}