package vulkan

// #include <vulkan/vulkan.h>
import "C"
import (
	"unsafe"
)

type ShaderModule uintptr

type ShaderModuleCreateFlags uint32

type ShaderModuleCreateInfo struct {
	Type  StructureType
	Next  uintptr
	Flags ShaderModuleCreateFlags
	Code  []byte
}

type shaderModuleCreateInfo struct {
	Type     StructureType
	Next     uintptr
	Flags    ShaderModuleCreateFlags
	CodeSize C.size_t
	Code     *byte
}

func CreateShaderModule(device Device, createInfo ShaderModuleCreateInfo, allocator *AllocationCallbacks) (ShaderModule, error) {
	var shaderModule ShaderModule
	_createInfo := shaderModuleCreateInfo{
		Type:     createInfo.Type,
		Next:     createInfo.Next,
		Flags:    createInfo.Flags,
		CodeSize: C.size_t(len(createInfo.Code)),
		Code:     (*byte)(unsafe.Pointer(&createInfo.Code[0])),
	}
	result := Result(C.vkCreateShaderModule(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkShaderModuleCreateInfo)(unsafe.Pointer(&_createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkShaderModule)(unsafe.Pointer(&shaderModule)),
	))
	if result != Success {
		return 0, result
	}
	return shaderModule, nil
}

func DestroyShaderModule(device Device, shaderModule ShaderModule, allocator *AllocationCallbacks) {
	C.vkDestroyShaderModule(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkShaderModule)(unsafe.Pointer(shaderModule)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}
