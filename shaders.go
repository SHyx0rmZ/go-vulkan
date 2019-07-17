package vulkan

type ShaderModule uintptr

type ShaderModuleCreateInfo struct{}

func CreateShaderModule(device Device, createInfo ShaderModuleCreateInfo, allocator *AllocationCallbacks) (ShaderModule, error) {
	return 0, nil
}

func DestroyShaderModule(device Device, shaderModule ShaderModule, allocator *AllocationCallbacks) {

}
