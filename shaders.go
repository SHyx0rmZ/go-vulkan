package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
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

func (info *ShaderModuleCreateInfo) C(_info *shaderModuleCreateInfo) freeFunc {
	*_info = shaderModuleCreateInfo{
		Type:     info.Type,
		Next:     info.Next,
		Flags:    info.Flags,
		CodeSize: (C.size_t)(len(info.Code)),
	}
	p := C.CBytes(info.Code)
	_info.Code = (*byte)(p)
	return freeFunc(func() {
		C.free(p)
	})
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
	var _createInfo shaderModuleCreateInfo
	defer createInfo.C(&_createInfo).Free()
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
