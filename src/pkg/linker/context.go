package linker

type ContextArgs struct {
	Output 			string
	Emulation 		MachineType
	LibratyPaths 	[]string
}

type Context struct {
	Args 		ContextArgs
}

func NewContext() *Context {
	return &Context{
		Args: ContextArgs{
			Output: "a.out",
			Emulation: MachineTypeNone,
		},
	}
}