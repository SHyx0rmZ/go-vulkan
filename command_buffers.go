package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
// #include <string.h>
import "C"
import (
	"unsafe"
)

func (info *SubmitInfo) C(_info *submitInfo) freeFunc {
	*_info = submitInfo{
		Type:                 info.Type,
		Next:                 info.Next,
		WaitSemaphoreCount:   uint32(len(info.WaitSemaphores)),
		WaitSemaphores:       nil,
		WaitDstStageMask:     nil,
		CommandBufferCount:   uint32(len(info.CommandBuffers)),
		CommandBuffers:       nil,
		SignalSemaphoreCount: uint32(len(info.SignalSemaphores)),
		SignalSemaphores:     nil,
	}
	var ps []unsafe.Pointer
	if _info.WaitSemaphoreCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.WaitSemaphoreCount) * unsafe.Sizeof(Semaphore(0))))
		ps = append(ps, p)
		for i, semaphore := range info.WaitSemaphores {
			*(*Semaphore)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(Semaphore(0)))) = semaphore
		}
		_info.WaitSemaphores = (*Semaphore)(p)
	}
	if _info.WaitSemaphoreCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.WaitSemaphoreCount) * unsafe.Sizeof(PipelineStageFlags(0))))
		ps = append(ps, p)
		for i, mask := range info.WaitDstStageMask {
			*(*PipelineStageFlags)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(PipelineStageFlags(0)))) = mask
		}
		_info.WaitDstStageMask = (*PipelineStageFlags)(p)
	}
	if _info.CommandBufferCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.CommandBufferCount) * unsafe.Sizeof(CommandBuffer(0))))
		ps = append(ps, p)
		for i, semaphore := range info.CommandBuffers {
			*(*CommandBuffer)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(CommandBuffer(0)))) = semaphore
		}
		_info.CommandBuffers = (*CommandBuffer)(p)
	}
	if _info.SignalSemaphoreCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.SignalSemaphoreCount) * unsafe.Sizeof(Semaphore(0))))
		ps = append(ps, p)
		for i, semaphore := range info.SignalSemaphores {
			*(*Semaphore)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(Semaphore(0)))) = semaphore
		}
		_info.SignalSemaphores = (*Semaphore)(p)
	}
	return freeFunc(func() {
		for i := len(ps); i > 0; i-- {
			C.free(ps[i-1])
		}
	})
}

type submitInfo struct {
	Type                 StructureType
	Next                 uintptr
	WaitSemaphoreCount   uint32
	WaitSemaphores       *Semaphore
	WaitDstStageMask     *PipelineStageFlags
	CommandBufferCount   uint32
	CommandBuffers       *CommandBuffer
	SignalSemaphoreCount uint32
	SignalSemaphores     *Semaphore
}

func CreateCommandPool(device Device, createInfo CommandPoolCreateInfo, allocator *AllocationCallbacks) (CommandPool, error) {
	var commandPool CommandPool
	result := Result(C.vkCreateCommandPool(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkCommandPoolCreateInfo)(unsafe.Pointer(&createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkCommandPool)(unsafe.Pointer(&commandPool)),
	))
	if result != Success {
		return 0, result
	}
	return commandPool, nil
}

func TrimCommandPool(device Device, commandPool CommandPool, flags CommandPoolTrimFlags) {
}

func ResetCommandPool(device Device, commandPool CommandPool, flags CommandPoolResetFlags) error {
	return _not_implemented
}

func DestroyCommandPool(device Device, commandPool CommandPool, allocator *AllocationCallbacks) {
	C.vkDestroyCommandPool(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkCommandPool)(unsafe.Pointer(commandPool)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

func AllocateCommandBuffers(device Device, allocateInfo CommandBufferAllocateInfo) ([]CommandBuffer, error) {
	commandBuffers := make([]CommandBuffer, allocateInfo.CommandBufferCount)
	result := Result(C.vkAllocateCommandBuffers(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkCommandBufferAllocateInfo)(unsafe.Pointer(&allocateInfo)),
		(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffers[0])),
	))
	if result != Success {
		return nil, result
	}
	return commandBuffers, nil
}

func ResetCommandBuffer(commandBuffer CommandBuffer, flags CommandBufferResetFlags) error {
	return _not_implemented
}

func FreeCommandBuffers(device Device, commandPool CommandPool, commandBuffers []CommandBuffer) {
	C.vkFreeCommandBuffers(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkCommandPool)(unsafe.Pointer(commandPool)),
		(C.uint32_t)(len(commandBuffers)),
		(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffers[0])),
	)
}

func BeginCommandBuffer(commandBuffer CommandBuffer, beginInfo CommandBufferBeginInfo) error {
	result := Result(C.vkBeginCommandBuffer(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(*C.VkCommandBufferBeginInfo)(unsafe.Pointer(&beginInfo)),
	))
	if result != Success {
		return result
	}
	return nil
}

func EndCommandBuffer(commandBuffer CommandBuffer) error {
	result := Result(C.vkEndCommandBuffer(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
	))
	if result != Success {
		return result
	}
	return nil
}

func QueueSubmit(queue Queue, submits []SubmitInfo, fence Fence) (freeFunc, error) {
	_submits := make([]submitInfo, len(submits))
	var fs []freeFunc
	for i, submit := range submits {
		fs = append(fs, submit.C(&_submits[i]))
	}
	ff := freeFunc(func() {
		for i := len(fs); i > 0; i-- {
			fs[i-1]()
		}
	})
	result := Result(C.vkQueueSubmit(
		(C.VkQueue)(unsafe.Pointer(queue)),
		(C.uint32_t)(len(submits)),
		(*C.VkSubmitInfo)(unsafe.Pointer(&_submits[0])),
		(C.VkFence)(unsafe.Pointer(fence)),
	))
	if result != Success {
		return ff, result
	}
	return ff, nil
}

func CmdExecuteCommands(commandBuffer CommandBuffer, commandBuffers []CommandBuffer) {

}

func CmdSetDeviceMask(commandBuffer CommandBuffer, deviceMask uint32) {

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

func CmdBindIndexBuffer(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, indexType uint32) {
	C.vkCmdBindIndexBuffer(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkBuffer)(unsafe.Pointer(buffer)),
		(C.VkDeviceSize)(offset),
		(C.VkIndexType)(indexType),
	)
}

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

type ImageCreateFlags uint32

type ImageCreateInfo struct {
	Type               StructureType
	Next               uintptr
	Flags              ImageCreateFlags
	ImageType          ImageType
	Format             Format
	Extent             Extent3D
	MipLevels          uint32
	ArrayLayers        uint32
	Samples            SampleCountFlagBits
	Tiling             ImageTiling
	Usage              ImageUsageFlags
	SharingMode        SharingMode
	QueueFamilyIndices []uint32
	InitialLayout      ImageLayout
}

func (info *ImageCreateInfo) C(_info *imageCreateInfo) freeFunc {
	*_info = imageCreateInfo{
		Type:                  info.Type,
		Next:                  info.Next,
		Flags:                 info.Flags,
		ImageType:             info.ImageType,
		Format:                info.Format,
		Extend:                info.Extent,
		MipLevels:             info.MipLevels,
		ArrayLayers:           info.ArrayLayers,
		Samples:               info.Samples,
		Tiling:                info.Tiling,
		Usage:                 info.Usage,
		SharingMode:           info.SharingMode,
		QueueFamilyIndexCount: uint32(len(info.QueueFamilyIndices)),
		QueueFamilyIndices:    nil,
		InitialLayout:         info.InitialLayout,
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

type imageCreateInfo struct {
	Type                  StructureType
	Next                  uintptr
	Flags                 ImageCreateFlags
	ImageType             ImageType
	Format                Format
	Extend                Extent3D
	MipLevels             uint32
	ArrayLayers           uint32
	Samples               SampleCountFlagBits
	Tiling                ImageTiling
	Usage                 ImageUsageFlags
	SharingMode           SharingMode
	QueueFamilyIndexCount uint32
	QueueFamilyIndices    *uint32
	InitialLayout         ImageLayout
}

func CreateImage(device Device, createInfo ImageCreateInfo, allocator *AllocationCallbacks) (Image, error) {
	var image Image
	var _createInfo imageCreateInfo
	defer createInfo.C(&_createInfo).Free()
	result := Result(C.vkCreateImage(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkImageCreateInfo)(unsafe.Pointer(&_createInfo)),
		(*C.VkAllocationCallbacks)(allocator),
		(*C.VkImage)(unsafe.Pointer(&image)),
	))
	if result != Success {
		return 0, result
	}
	return image, nil
}

func DestroyImage(device Device, image Image, allocator *AllocationCallbacks) {
	C.vkDestroyImage(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkImage)(unsafe.Pointer(image)),
		(*C.VkAllocationCallbacks)(allocator),
	)
}

func GetImageMemoryRequirements(device Device, image Image) MemoryRequirements {
	var memoryRequirements MemoryRequirements
	C.vkGetImageMemoryRequirements(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkImage)(unsafe.Pointer(image)),
		(*C.VkMemoryRequirements)(unsafe.Pointer(&memoryRequirements)),
	)
	return memoryRequirements
}

func BindImageMemory(device Device, image Image, memory DeviceMemory, memoryOffset DeviceSize) error {
	result := Result(C.vkBindImageMemory(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkImage)(unsafe.Pointer(image)),
		(C.VkDeviceMemory)(unsafe.Pointer(memory)),
		(C.VkDeviceSize)(memoryOffset),
	))
	if result != Success {
		return result
	}
	return nil
}

// todo
func CmdCopyBufferToImage(commandBuffer CommandBuffer, srcBuffer Buffer, dstImage Image, dstImageLayout ImageLayout, regions []BufferImageCopy) {
	C.vkCmdCopyBufferToImage(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkBuffer)(unsafe.Pointer(srcBuffer)),
		(C.VkImage)(unsafe.Pointer(dstImage)),
		(C.VkImageLayout)(dstImageLayout),
		(C.uint32_t)(len(regions)),
		(*C.VkBufferImageCopy)(unsafe.Pointer(&regions[0])),
	)
}

// todo
func CmdCopyImageToBuffer(commandBuffer CommandBuffer, srcImage Image, srcImageLayout ImageLayout, dstBuffer Buffer, regions []BufferImageCopy) {
	C.vkCmdCopyImageToBuffer(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkImage)(unsafe.Pointer(srcImage)),
		(C.VkImageLayout)(srcImageLayout),
		(C.VkBuffer)(unsafe.Pointer(dstBuffer)),
		(C.uint32_t)(len(regions)),
		(*C.VkBufferImageCopy)(unsafe.Pointer(&regions[0])),
	)
}

func CreateSampler(device Device, createInfo SamplerCreateInfo, allocator *AllocationCallbacks) (Sampler, error) {
	var sampler Sampler
	result := Result(C.vkCreateSampler(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkSamplerCreateInfo)(unsafe.Pointer(&createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkSampler)(unsafe.Pointer(&sampler)),
	))
	if result != Success {
		return 0, result
	}
	return sampler, nil
}

func DestroySampler(device Device, sampler Sampler, allocator *AllocationCallbacks) {
	C.vkDestroySampler(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkSampler)(unsafe.Pointer(sampler)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}
