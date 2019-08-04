package vulkan

type ShaderModule uintptr

type ShaderModuleCreateFlags uint32

type ShaderModuleCreateInfo struct {
	Type  StructureType
	Next  uintptr
	Flags ShaderModuleCreateFlags
	Code  []byte
}
