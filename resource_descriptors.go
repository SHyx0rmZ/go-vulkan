package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

// Descriptor SetLayout

type PipelineLayoutCreateFlags uint32

type PipelineLayoutCreateInfo struct {
	Type               StructureType
	Next               uintptr
	Flags              PipelineLayoutCreateFlags
	SetLayouts         []DescriptorSetLayout
	PushConstantRanges []PushConstantRange
}

func (info *PipelineLayoutCreateInfo) C(_info *pipelineLayoutCreateInfo) freeFunc {
	*_info = pipelineLayoutCreateInfo{
		Type:                   info.Type,
		Next:                   info.Next,
		Flags:                  info.Flags,
		SetLayoutCount:         uint32(len(info.SetLayouts)),
		PushConstantRangeCount: uint32(len(info.PushConstantRanges)),
	}
	var ps []unsafe.Pointer
	if _info.SetLayoutCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.SetLayoutCount) * unsafe.Sizeof(DescriptorSetLayout(0))))
		ps = append(ps, p)
		for i, layout := range info.SetLayouts {
			*(*DescriptorSetLayout)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(DescriptorSetLayout(0)))) = layout
		}
		_info.SetLayouts = (*DescriptorSetLayout)(p)
	}
	if _info.PushConstantRangeCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.PushConstantRangeCount) * unsafe.Sizeof(PushConstantRange{})))
		ps = append(ps, p)
		for i, pushConstantRange := range info.PushConstantRanges {
			*(*PushConstantRange)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(PushConstantRange{}))) = pushConstantRange
		}
		_info.PushConstantRanges = (*PushConstantRange)(p)
	}
	return freeFunc(func() {
		for i := len(ps); i > 0; i-- {
			C.free(ps[i-1])
		}
	})
}

type pipelineLayoutCreateInfo struct {
	Type                   StructureType
	Next                   uintptr
	Flags                  PipelineLayoutCreateFlags
	SetLayoutCount         uint32
	SetLayouts             *DescriptorSetLayout
	PushConstantRangeCount uint32
	PushConstantRanges     *PushConstantRange
}

type DescriptorSetLayout uintptr

type PushConstantRange struct {
	StageFlags ShaderStageFlags
	Offset     uint32
	Size       uint32
}

func CreatePipelineLayout(device Device, createInfo PipelineLayoutCreateInfo, allocator *AllocationCallbacks) (PipelineLayout, error) {
	var pipelineLayout PipelineLayout
	result := Result(C.vkCreatePipelineLayout(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkPipelineLayoutCreateInfo)(unsafe.Pointer(&createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkPipelineLayout)(unsafe.Pointer(&pipelineLayout)),
	))
	if result != Success {
		return 0, result
	}
	return pipelineLayout, nil
}

func DestroyPipelineLayout(device Device, pipelineLayout PipelineLayout, allocator *AllocationCallbacks) {
	C.vkDestroyPipelineLayout(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineLayout)(unsafe.Pointer(pipelineLayout)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}
