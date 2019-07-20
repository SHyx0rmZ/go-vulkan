package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
// #include <string.h>
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

type MemoryAllocateInfo struct {
	Type            StructureType
	Next            uintptr
	AllocationSize  DeviceSize
	MemoryTypeIndex uint32
}

type DeviceMemory uintptr

func AllocateMemory(device Device, allocateInfo MemoryAllocateInfo, allocator *AllocationCallbacks) (DeviceMemory, error) {
	var deviceMemory DeviceMemory
	result := Result(C.vkAllocateMemory(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkMemoryAllocateInfo)(unsafe.Pointer(&allocateInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkDeviceMemory)(unsafe.Pointer(&deviceMemory)),
	))
	if result != Success {
		return 0, result
	}
	return deviceMemory, nil
}

func FreeMemory(device Device, memory DeviceMemory, allocator *AllocationCallbacks) {
	C.vkFreeMemory(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDeviceMemory)(unsafe.Pointer(memory)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

type MemoryMapFlags uint32

func MapMemory(device Device, memory DeviceMemory, offset, size DeviceSize, flags MemoryMapFlags) (uintptr, error) {
	var data uintptr
	result := Result(C.vkMapMemory(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDeviceMemory)(unsafe.Pointer(memory)),
		(C.VkDeviceSize)(offset),
		(C.VkDeviceSize)(size),
		(C.VkMemoryMapFlags)(flags),
		(*unsafe.Pointer)(unsafe.Pointer(&data)),
	))
	if result != Success {
		return 0, result
	}
	return data, nil
}

func UnmapMemory(device Device, memory DeviceMemory) {
	C.vkUnmapMemory(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkDeviceMemory)(unsafe.Pointer(memory)),
	)
}

type MemoryRequirements struct {
	Size           DeviceSize
	Alignment      DeviceSize
	MemoryTypeBits uint32
}

func GetBufferMemoryRequirements(device Device, buffer Buffer) MemoryRequirements {
	var memoryRequirements MemoryRequirements
	C.vkGetBufferMemoryRequirements(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkBuffer)(unsafe.Pointer(buffer)),
		(*C.VkMemoryRequirements)(unsafe.Pointer(&memoryRequirements)),
	)
	return memoryRequirements
}

func GetPhysicalDeviceMemoryProperties(device PhysicalDevice) PhysicalDeviceMemoryProperties {
	var properties PhysicalDeviceMemoryProperties
	C.vkGetPhysicalDeviceMemoryProperties(
		(C.VkPhysicalDevice)(unsafe.Pointer(device)),
		(*C.VkPhysicalDeviceMemoryProperties)(unsafe.Pointer(&properties)),
	)
	return properties
}

type PhysicalDeviceMemoryProperties struct {
	MemoryTypeCount uint32
	MemoryTypes     [32]MemoryType
	MemoryHeapCount uint32
	MemoryHeaps     [16]MemoryHeap
}

type MemoryHeapFlagBits uint32
type MemoryHeapFlags = MemoryHeapFlagBits

const (
	MemoryHeapDeviceLocalBit MemoryHeapFlagBits = 1 << iota
	MemoryHeapMultiInstanceBit
)

type MemoryHeap struct {
	Size  DeviceSize
	Flags MemoryHeapFlags
}

type MemoryPropertyFlagBits uint32
type MemoryPropertyFlags = MemoryPropertyFlagBits

const (
	MemoryPropertyDeviceLocalBit MemoryPropertyFlagBits = 1 << iota
	MemoryPropertyHostVisibleBit
	MemoryPropertyHostCoherentBit
	MemoryPropertyHostCachedBit
	MemoryPropertyLazilyAllocatedBit
	MemoryPropertyProtectedBit
)

type MemoryType struct {
	PropertyFlags MemoryPropertyFlags
	HeapIndex     uint32
}

func BindBufferMemory(device Device, buffer Buffer, memory DeviceMemory, offset DeviceSize) error {
	result := Result(C.vkBindBufferMemory(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkBuffer)(unsafe.Pointer(buffer)),
		(C.VkDeviceMemory)(unsafe.Pointer(memory)),
		(C.VkDeviceSize)(offset),
	))
	if result != Success {
		return result
	}
	return nil
}

func Memcpy(dst unsafe.Pointer, src unsafe.Pointer, size uintptr) {
	C.memcpy(
		dst,
		src,
		C.size_t(size),
	)
}

type SharingMode uint32

const (
	SharingModeExclusive SharingMode = iota
	SharingModeConcurrent
)
