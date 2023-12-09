package vulkan

// #include <stdlib.h>
// #include <vulkan/vulkan.h>
import "C"
import (
	"unsafe"
)

type Buffer uintptr

type BufferCreateFlagBits uint32
type BufferCreateFlags = BufferCreateFlagBits

const (
	BufferCreateSparseBindingsBit BufferCreateFlagBits = 1 << iota
	BufferCreateSparseResidencyBit
	BufferCreateSparseAliasedBit
	BufferCreateProtectedBit
)

type BufferCreateInfo struct {
	Type               StructureType
	Next               uintptr
	Flags              BufferCreateFlags
	Size               DeviceSize
	Usage              BufferUsageFlags
	SharingMode        SharingMode
	QueueFamilyIndices []uint32
}

type BufferUsageFlagBits uint32
type BufferUsageFlags = BufferUsageFlagBits

const (
	BufferUsageTransferSrcBit BufferUsageFlagBits = 1 << iota
	BufferUsageTransferDstBit
	BufferUsageUniformTexelBufferBit
	BufferUsageStorageTexelBufferBit
	BufferUsageUniformBufferBit
	BufferUsageStorageBufferBit
	BufferUsageIndexBufferBit
	BufferUsageVertexBufferBit
	BufferUsageIndirectBufferBit
)

type BufferCopy struct {
	SrcOffset DeviceSize
	DstOffset DeviceSize
	Size      DeviceSize
}

func CmdCopyBuffer(commandBuffer CommandBuffer, srcBuffer, dstBuffer Buffer, regions []BufferCopy) {
	var regionPtr unsafe.Pointer
	if len(regions) > 0 {
		regionPtr = unsafe.Pointer(&regions[0])
	}
	C.vkCmdCopyBuffer(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		*(*C.VkBuffer)(unsafe.Pointer(&srcBuffer)),
		*(*C.VkBuffer)(unsafe.Pointer(&dstBuffer)),
		(C.uint32_t)(len(regions)),
		(*C.VkBufferCopy)(regionPtr),
	)
}

type CopyBufferInfo2 struct {
	Type      StructureType
	Next      unsafe.Pointer
	SrcBuffer Buffer
	DstBuffer Buffer
	Regions   []BufferCopy2
}

type copyBufferInfo2 struct {
	Type        StructureType
	Next        unsafe.Pointer
	SrcBuffer   Buffer
	DstBuffer   Buffer
	RegionCount uint32
	RegionPtr   *BufferCopy2
}

func (info *CopyBufferInfo2) C(_info *copyBufferInfo2) freeFunc {
	var regionPtr unsafe.Pointer
	if len(info.Regions) > 0 {
		regionPtr = copySliceToC(nil, info.Regions)
	}
	*_info = copyBufferInfo2{
		Type:        info.Type,
		Next:        info.Next,
		SrcBuffer:   info.SrcBuffer,
		DstBuffer:   info.DstBuffer,
		RegionCount: uint32(len(info.Regions)),
		RegionPtr:   (*BufferCopy2)(regionPtr),
	}
	return func() {
		C.free(regionPtr)
	}
}

type BufferCopy2 struct {
	Type      StructureType
	Next      unsafe.Pointer
	SrcOffset DeviceSize
	DstOffset DeviceSize
	Size      DeviceSize
}

func CmdCopyBuffer2(commandBuffer CommandBuffer, info CopyBufferInfo2) {
	var _info copyBufferInfo2
	info.C(&_info)
	C.vkCmdCopyBuffer2(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		(*C.VkCopyBufferInfo2)(unsafe.Pointer(&_info)),
	)
}
