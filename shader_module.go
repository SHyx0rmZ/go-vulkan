package vulkan

type ShaderModule uint64

type ShaderModuleCreateFlags uint32

type ShaderModuleCreateInfo struct {
	Type  StructureType
	Next  uintptr
	Flags ShaderModuleCreateFlags
	Code  []byte
}
