// ELF相关数据结构

package linker

import (
	"unsafe"
	"bytes"
)

const EhdrSize = int(unsafe.Sizeof(Ehdr{}))
const ShdrSize = int(unsafe.Sizeof(Shdr{}))
const SymSize = int(unsafe.Sizeof(Sym{}))

// ELF头部
type Ehdr struct {
	Ident     [16]uint8
	Type      uint16		// 文件类型：1=可重定位，2=可执行，3=共享目标文件，4=转储镜像（core image）
	Machine   uint16		// 架构类型：2=SPARC，3=x86，4=68K，等等
	Version   uint32		// 文件版本，总是1
	Entry     uint64		// 入口地址（若为可执行文件）
	PhOff     uint64		// 程序头在文件中的位置（不存在则为0）
	ShOff     uint64		// 区段头在文件中的位置（不存在则为0）
	Flags     uint32		// 体系结构相关的标志，总为0
	EhSize    uint16		// 该ELF头部的大小
	PhEntSize uint16		// 程序头表项的大小
	PhNum     uint16		// 程序头表项的个数（不存在则为0）
	ShEntSize uint16		// 区段头表项的大小
	ShNum     uint16		// 区段头表项的个数（不存在则为0）
	ShStrndx  uint16		// 保存区段名称字符串的区段的序号
}

// 区段头部 section header
type Shdr struct {
	Name      uint32		// 名称。区段名称是一个字符串，这里存储的是名称字符串在字符串表中的索引 
	Type      uint32		// 区段类型	
	Flags     uint64		// 标志位
	Addr      uint64		// 若该区段可加载，则为被加载的内存基址，否则为0
	Offset    uint64		// 区段起始点在文件中的位置
	Size      uint64		// 区段大小（字节为单位）
	Link      uint32		// 保存关联信息的区段号，若没有则为0
	Info      uint32		// 一些区段相关的信息
	AddrAlign uint64		// 移动区域的对齐粒度
	EntSize   uint64		// 若该区段是一张表，表示每一个表项的大小
}

// 符号表
type Sym struct {
	Name  uint32
	Info  uint8
	Other uint8
	Shndx uint16
	Val   uint64
	Size  uint64
}

func ElfGetName(strTab []byte, offset uint32) string {
	length := uint32(bytes.Index(strTab[offset:], []byte{0}))
	return string(strTab[offset : offset+length])	
}