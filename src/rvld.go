// 主函数

package main

import (
	"github.com/shaozk/rvld-learn/pkg/linker"
	"github.com/shaozk/rvld-learn/pkg/utils"
	"os"
)

func main() {
	if len(os.Args) < 2{
		utils.Fatal("Wrong args")
	}
	
	filename := os.Args[1]
	file := linker.MustNewFile(filename)
	inputfile := linker.NewInputFile(file)
	
	for _, str := range inputfile.ElfSections {
		println(str.Name)
	}	
	utils.Assert(len(inputfile.ElfSections) == 11)
	

}