package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
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

func (info *BufferCreateInfo) C(_info *bufferCreateInfo) freeFunc {
	*_info = bufferCreateInfo{
		Type:                  info.Type,
		Next:                  info.Next,
		Flags:                 info.Flags,
		Size:                  info.Size,
		Usage:                 info.Usage,
		SharingMode:           info.SharingMode,
		QueueFamilyIndexCount: uint32(len(info.QueueFamilyIndices)),
		QueueFamilyIndices:    nil,
	}
	if _info.QueueFamilyIndexCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.QueueFamilyIndexCount) * unsafe.Sizeof(uint32(0))))
		for i, index := range info.QueueFamilyIndices {
			*(*uint32)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(uint32(0)))) = index
		}
		return freeFunc(func() {
			C.free(p)
		})
	}
	return freeFunc(nil)
}

type bufferCreateInfo struct {
	Type                  StructureType
	Next                  uintptr
	Flags                 BufferCreateFlags
	Size                  DeviceSize
	Usage                 BufferUsageFlags
	SharingMode           SharingMode
	QueueFamilyIndexCount uint32
	QueueFamilyIndices    *uint32
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

func CreateBuffer(device Device, createInfo BufferCreateInfo, allocator *AllocationCallbacks) (Buffer, error) {
	var buffer Buffer
	var _createInfo bufferCreateInfo
	defer createInfo.C(&_createInfo).Free()
	result := Result(C.vkCreateBuffer(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkBufferCreateInfo)(unsafe.Pointer(&_createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkBuffer)(unsafe.Pointer(&buffer)),
	))
	if result != Success {
		return 0, result
	}
	return buffer, nil
}

func DestroyBuffer(device Device, buffer Buffer, allocator *AllocationCallbacks) {
	C.vkDestroyBuffer(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkBuffer)(unsafe.Pointer(buffer)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

func CmdBindVertexBuffers(commandBuffer CommandBuffer, firstBinding uint32, buffers []Buffer, offsets []DeviceSize) {
	C.vkCmdBindVertexBuffers(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.uint32_t)(firstBinding),
		(C.uint32_t)(len(buffers)),
		(*C.VkBuffer)(unsafe.Pointer(&buffers[0])),
		(*C.VkDeviceSize)(unsafe.Pointer(&offsets[0])),
	)
}

type SharingMode uint32

const (
	SharingModeExclusive SharingMode = iota
	SharingModeConcurrent
)
