package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
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
	var _createInfo pipelineLayoutCreateInfo
	defer createInfo.C(&_createInfo).Free()
	result := Result(C.vkCreatePipelineLayout(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkPipelineLayoutCreateInfo)(unsafe.Pointer(&_createInfo)),
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

type DescriptorSetLayoutCreateFlagBits uint32
type DescriptorSetLayoutCreateFlags = DescriptorSetLayoutCreateFlagBits

type DescriptorSetLayoutCreateInfo struct {
	Type     StructureType
	Next     uintptr
	Flags    DescriptorSetLayoutCreateFlags
	Bindings []DescriptorSetLayoutBinding
}

func (info *DescriptorSetLayoutCreateInfo) C(_info *descriptorSetLayoutCreateInfo) freeFunc {
	*_info = descriptorSetLayoutCreateInfo{
		Type:         info.Type,
		Next:         info.Next,
		Flags:        info.Flags,
		BindingCount: uint32(len(info.Bindings)),
		Bindings:     nil,
	}
	if _info.BindingCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.BindingCount) * unsafe.Sizeof(DescriptorSetLayoutBinding{})))
		for i, binding := range info.Bindings {
			*(*DescriptorSetLayoutBinding)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(DescriptorSetLayoutBinding{}))) = binding
		}
		_info.Bindings = (*DescriptorSetLayoutBinding)(p)
		return freeFunc(func() {
			C.free(p)
		})
	}
	return freeFunc(nil)
}

type descriptorSetLayoutCreateInfo struct {
	Type         StructureType
	Next         uintptr
	Flags        DescriptorSetLayoutCreateFlags
	BindingCount uint32
	Bindings     *DescriptorSetLayoutBinding
}

type DescriptorSetLayoutBinding struct {
	Binding           uint32
	DescriptorType    DescriptorType
	DescriptorCount   uint32
	StageFlags        ShaderStageFlags
	ImmutableSamplers *Sampler
}

type DescriptorType uint32

const (
	DescriptorTypeSampler DescriptorType = iota
	DescriptorTypeCombinedImageSampler
	DescriptorTypeSampledImage
	DescriptorTypeStorageImage
	DescriptorTypeUniformTexelBuffer
	DescriptorTypeStorageTexelBuffer
	DescriptorTypeUniformBuffer
	DescriptorTypeUniformBufferDynamic
	DescriptorTypeStorageBuffer
	DescriptorTypeStorageBufferDynamic
	DescriptorTypeImageAttachment
)

type Sampler uintptr

func CreateDescriptorSetLayout(device Device, createInfo DescriptorSetLayoutCreateInfo, allocator *AllocationCallbacks) (DescriptorSetLayout, error) {
	var layout DescriptorSetLayout
	var _createInfo descriptorSetLayoutCreateInfo
	defer createInfo.C(&_createInfo).Free()
	result := Result(C.vkCreateDescriptorSetLayout(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkDescriptorSetLayoutCreateInfo)(unsafe.Pointer(&_createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkDescriptorSetLayout)(unsafe.Pointer(&layout)),
	))
	if result != Success {
		return 0, result
	}
	return layout, nil
}

func DestroyDescriptorSetLayout(device Device, layout DescriptorSetLayout, allocator *AllocationCallbacks) {
	C.vkDestroyDescriptorSetLayout(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDescriptorSetLayout)(unsafe.Pointer(layout)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

type DescriptorPool uintptr

type DescriptorPoolCreateFlagBits uint32
type DescriptorPoolCreateFlags = DescriptorPoolCreateFlagBits

const (
	DescriptorPoolCreateFreeDescriptorSetBit DescriptorPoolCreateFlagBits = 1 << iota
)

type DescriptorPoolCreateInfo struct {
	Type      StructureType
	Next      uintptr
	Flags     DescriptorPoolCreateFlags
	MaxSets   uint32
	PoolSizes []DescriptorPoolSize
}

func (info *DescriptorPoolCreateInfo) C(_info *descriptorPoolCreateInfo) freeFunc {
	*_info = descriptorPoolCreateInfo{
		Type:          info.Type,
		Next:          info.Next,
		Flags:         info.Flags,
		MaxSets:       info.MaxSets,
		PoolSizeCount: uint32(len(info.PoolSizes)),
		PoolSizes:     nil,
	}
	if _info.PoolSizeCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.PoolSizeCount) * unsafe.Sizeof(DescriptorPoolSize{})))
		for i, size := range info.PoolSizes {
			*(*DescriptorPoolSize)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(DescriptorPoolSize{}))) = size
		}
		_info.PoolSizes = (*DescriptorPoolSize)(p)
		return freeFunc(func() {
			C.free(p)
		})
	}
	return freeFunc(nil)
}

type descriptorPoolCreateInfo struct {
	Type          StructureType
	Next          uintptr
	Flags         DescriptorPoolCreateFlags
	MaxSets       uint32
	PoolSizeCount uint32
	PoolSizes     *DescriptorPoolSize
}

type DescriptorPoolSize struct {
	Type            DescriptorType
	DescriptorCount uint32
}

func CreateDescriptorPool(device Device, createInfo DescriptorPoolCreateInfo, allocator *AllocationCallbacks) (DescriptorPool, error) {
	var pool DescriptorPool
	var _createInfo descriptorPoolCreateInfo
	defer createInfo.C(&_createInfo).Free()
	result := Result(C.vkCreateDescriptorPool(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkDescriptorPoolCreateInfo)(unsafe.Pointer(&_createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkDescriptorPool)(unsafe.Pointer(&pool)),
	))
	if result != Success {
		return 0, result
	}
	return pool, nil
}

func DestroyDescriptorPool(device Device, pool DescriptorPool, allocator *AllocationCallbacks) {
	C.vkDestroyDescriptorPool(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDescriptorPool)(unsafe.Pointer(pool)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

type DescriptorSetAllocateInfo struct {
	Type           StructureType
	Next           uintptr
	DescriptorPool DescriptorPool
	SetLayouts     []DescriptorSetLayout
}

func (info *DescriptorSetAllocateInfo) C(_info *descriptorAllocateInfo) freeFunc {
	*_info = descriptorAllocateInfo{
		Type:               info.Type,
		Next:               info.Next,
		DescriptorPool:     info.DescriptorPool,
		DescriptorSetCount: uint32(len(info.SetLayouts)),
		SetLayouts:         nil,
	}
	if _info.DescriptorSetCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.DescriptorSetCount) * unsafe.Sizeof(DescriptorSetLayout(0))))
		for i, layout := range info.SetLayouts {
			*(*DescriptorSetLayout)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(DescriptorSetLayout(0)))) = layout
		}
		_info.SetLayouts = (*DescriptorSetLayout)(p)
		return freeFunc(func() {
			C.free(p)
		})
	}
	return freeFunc(nil)
}

type descriptorAllocateInfo struct {
	Type               StructureType
	Next               uintptr
	DescriptorPool     DescriptorPool
	DescriptorSetCount uint32
	SetLayouts         *DescriptorSetLayout
}

type DescriptorSet uintptr

func AllocateDescriptorSets(device Device, allocateInfo DescriptorSetAllocateInfo) ([]DescriptorSet, error) {
	sets := make([]DescriptorSet, len(allocateInfo.SetLayouts))
	var _allocateInfo descriptorAllocateInfo
	defer allocateInfo.C(&_allocateInfo).Free()
	result := Result(C.vkAllocateDescriptorSets(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkDescriptorSetAllocateInfo)(unsafe.Pointer(&_allocateInfo)),
		(*C.VkDescriptorSet)(unsafe.Pointer(&sets[0])),
	))
	if result != Success {
		return nil, result
	}
	return sets, nil
}

func FreeDescriptorSets(device Device, pool DescriptorPool, sets []DescriptorSet) {
	C.vkFreeDescriptorSets(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDescriptorPool)(unsafe.Pointer(pool)),
		(C.uint32_t)(len(sets)),
		(*C.VkDescriptorSet)(unsafe.Pointer(&sets[0])),
	)
}

type WriteDescriptorSet struct {
	Type            StructureType
	Next            uintptr
	DstSet          DescriptorSet
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
	DescriptorType  DescriptorType
	ImageInfo       []DescriptorImageInfo
	BufferInfo      []DescriptorBufferInfo
	TexelBufferView []BufferView
}

func (set *WriteDescriptorSet) C(_set *writeDescriptorSet) freeFunc {
	*_set = writeDescriptorSet{
		Type:            set.Type,
		Next:            set.Next,
		DstSet:          set.DstSet,
		DstBinding:      set.DstBinding,
		DstArrayElement: set.DstArrayElement,
		DescriptorCount: set.DescriptorCount,
		DescriptorType:  set.DescriptorType,
		ImageInfo:       nil,
		BufferInfo:      nil,
		TexelBufferView: nil,
	}
	var ps []unsafe.Pointer
	if len(set.ImageInfo) > 0 {
		p := C.malloc(C.size_t(uintptr(len(set.ImageInfo)) * unsafe.Sizeof(DescriptorImageInfo{})))
		ps = append(ps, p)
		for i, info := range set.ImageInfo {
			*(*DescriptorImageInfo)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(DescriptorImageInfo{}))) = info
		}
		_set.ImageInfo = (*DescriptorImageInfo)(p)
	}
	if len(set.BufferInfo) > 0 {
		p := C.malloc(C.size_t(uintptr(len(set.BufferInfo)) * unsafe.Sizeof(DescriptorBufferInfo{})))
		ps = append(ps, p)
		for i, info := range set.BufferInfo {
			*(*DescriptorBufferInfo)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(DescriptorBufferInfo{}))) = info
		}
		_set.BufferInfo = (*DescriptorBufferInfo)(p)
	}
	if len(set.TexelBufferView) > 0 {
		p := C.malloc(C.size_t(uintptr(len(set.BufferInfo)) * unsafe.Sizeof(BufferView(0))))
		ps = append(ps, p)
		for i, view := range set.TexelBufferView {
			*(*BufferView)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(BufferView(0)))) = view
		}
		_set.TexelBufferView = (*BufferView)(p)
	}
	return freeFunc(func() {
		for _, p := range ps {
			C.free(p)
		}
	})
}

type writeDescriptorSet struct {
	Type            StructureType
	Next            uintptr
	DstSet          DescriptorSet
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
	DescriptorType  DescriptorType
	ImageInfo       *DescriptorImageInfo
	BufferInfo      *DescriptorBufferInfo
	TexelBufferView *BufferView
}

type DescriptorImageInfo struct {
	Sampler     Sampler
	ImageView   ImageView
	ImageLayout ImageLayout
}

type DescriptorBufferInfo struct {
	Buffer Buffer
	Offset DeviceSize
	Range  DeviceSize
}

type BufferView uintptr

type CopyDescriptorSet struct {
	Type            StructureType
	Next            uintptr
	SrcSet          DescriptorSet
	SrcBinding      uint32
	SrcArrayElement uint32
	DstSet          DescriptorSet
	DstBinding      uint32
	DstArrayElement uint32
	DescriptorCount uint32
}

func UpdateDescriptorSets(device Device, descriptorWrites []WriteDescriptorSet, descriptorCopies []CopyDescriptorSet) freeFunc {
	writeCount := uint32(len(descriptorWrites))
	copyCount := uint32(len(descriptorCopies))
	var writes unsafe.Pointer
	var copies unsafe.Pointer
	var fs []freeFunc
	if writeCount > 0 {
		p := C.malloc(C.size_t(uintptr(writeCount) * unsafe.Sizeof(writeDescriptorSet{})))
		fs = append(fs, freeFunc(func() {
			C.free(p)
		}))
		for i, write := range descriptorWrites {
			fs = append(fs, write.C((*writeDescriptorSet)(unsafe.Pointer(uintptr(p)+uintptr(i)*unsafe.Sizeof(writeDescriptorSet{})))))
		}
		writes = p
	}
	if copyCount > 0 {
		fmt.Println("copies")
		copies = unsafe.Pointer(&descriptorCopies[0])
	}
	fmt.Println(copyCount, "count", copies)
	fmt.Println(writeCount, "writes", writes)
	C.vkUpdateDescriptorSets(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.uint32_t)(writeCount),
		(*C.VkWriteDescriptorSet)(writes),
		(C.uint32_t)(copyCount),
		(*C.VkCopyDescriptorSet)(copies),
	)
	fmt.Println("asd")
	return freeFunc(func() {
		for _, f := range fs {
			f()
		}
	})
}

func CmdBindDescriptorSets(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, layout PipelineLayout, firstSet uint32, descriptorSets []DescriptorSet, dynamicOffsets []uint32) {
	var sets unsafe.Pointer
	var offsets unsafe.Pointer
	if len(descriptorSets) > 0 {
		sets = unsafe.Pointer(&descriptorSets[0])
	}
	if len(dynamicOffsets) > 0 {
		offsets = unsafe.Pointer(&dynamicOffsets[0])
	}
	C.vkCmdBindDescriptorSets(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkPipelineBindPoint)(pipelineBindPoint),
		(C.VkPipelineLayout)(unsafe.Pointer(layout)),
		(C.uint32_t)(firstSet),
		(C.uint32_t)(len(descriptorSets)),
		(*C.VkDescriptorSet)(sets),
		(C.uint32_t)(len(dynamicOffsets)),
		(*C.uint32_t)(offsets),
	)
}
