package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

type CommandBufferLevel uint32

const (
	CommandBufferLevelPrimary CommandBufferLevel = iota
	CommandBufferLevelSecondary
)

type CommandBufferAllocateInfo struct {
	Type               StructureType
	Next               uintptr
	CommandPool        CommandPool
	Level              CommandBufferLevel
	CommandBufferCount uint32
}

type CommandBufferUsageFlagBits uint32
type CommandBufferUsageFlags = CommandBufferUsageFlagBits

const (
	CommandBufferUsageOneTimeSubmitBit CommandBufferUsageFlagBits = 1 << iota
	CommandBufferUsageRenderPassContinueBit
	CommandBufferUsageSimultaneousUseBit
)

type CommandBufferBeginInfo struct {
	Type            StructureType
	Next            uintptr
	Flags           CommandBufferUsageFlags
	InheritanceInfo *CommandBufferInheritanceInfo
}

type CommandBufferInheritanceInfo struct{}

type SubmitInfo struct {
	Type             StructureType
	Next             uintptr
	WaitSemaphores   []Semaphore
	WaitDstStageMask []PipelineStageFlags
	CommandBuffers   []CommandBuffer
	SignalSemaphores []Semaphore
}

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
		for _, p := range ps {
			C.free(p)
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

type CommandPoolCreateFlagBits uint32
type CommandPoolCreateFlags = CommandPoolCreateFlagBits

const (
	CommandPoolCreateTransient CommandPoolCreateFlagBits = 1 << iota
	CommandPoolCreateResetCommandBuffer
	CommandPoolCreateProtected
)

type CommandPoolCreateInfo struct {
	Type             StructureType
	Next             uintptr
	Flags            CommandPoolCreateFlags
	QueueFamilyIndex uint32
}

type CommandPool uintptr

type CommandBuffer uintptr

type CommandPoolTrimFlags uint32

type CommandPoolResetFlags uint32

type CommandBufferResetFlags uint32

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

func QueueSubmit(queue Queue, submits []SubmitInfo, fence Fence) error {
	_submits := make([]submitInfo, len(submits))
	var fs []freeFunc
	for i, submit := range submits {
		fs = append(fs, submit.C(&_submits[i]))
	}
	defer func() {
		for _, f := range fs {
			f()
		}
	}()
	result := Result(C.vkQueueSubmit(
		(C.VkQueue)(unsafe.Pointer(queue)),
		(C.uint32_t)(len(submits)),
		(*C.VkSubmitInfo)(unsafe.Pointer(&_submits[0])),
		(C.VkFence)(unsafe.Pointer(fence)),
	))
	if result != Success {
		return result
	}
	return nil
}

func CmdExecuteCommands(commandBuffer CommandBuffer, commandBuffers []CommandBuffer) {

}

func CmdSetDeviceMask(commandBuffer CommandBuffer, deviceMask uint32) {

}
