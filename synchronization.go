package vulkan

// #include <vulkan/vulkan.h>
import "C"
import (
	"time"
	"unsafe"
)

type Fence uintptr

type FenceCreateFlagBits uint32
type FenceCreateFlags FenceCreateFlagBits

const (
	FenceCreateSignaledBit FenceCreateFlagBits = 1 << iota
)

type FenceCreateInfo struct {
	Type  StructureType
	Next  uintptr
	Flags FenceCreateFlags
}

type EventCreateFlags uint32

type EventCreateInfo struct {
	Type  StructureType
	Next  uintptr
	Flags EventCreateFlags
}

type Semaphore uintptr

type SemaphoreCreateFlags uint32

type SemaphoreCreateInfo struct {
	Type  StructureType
	Next  uintptr
	Flags SemaphoreCreateFlags
}

type Event uintptr

type MemoryBarrier uintptr

type BufferMemoryBarrier uintptr

type ImageMemoryBarrier uintptr

func CreateFence(device Device, createInfo FenceCreateInfo, allocator *AllocationCallbacks) (Fence, error) {
	var fence Fence
	result := Result(C.vkCreateFence(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkFenceCreateInfo)(unsafe.Pointer(&createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkFence)(unsafe.Pointer(&fence)),
	))
	if result != Success {
		return 0, result
	}
	return fence, nil
}

func DestroyFence(device Device, fence Fence, allocator *AllocationCallbacks) {
	C.vkDestroyFence(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkFence)(unsafe.Pointer(fence)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

func GetFenceStatus(device Device, fence Fence) error {
	result := Result(C.vkGetFenceStatus(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkFence)(unsafe.Pointer(fence)),
	))
	if result != Success {
		return result
	}
	return nil
}

func ResetFences(device Device, fences []Fence) error {
	result := Result(C.vkResetFences(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.uint32_t)(len(fences)),
		(*C.VkFence)(unsafe.Pointer(&fences[0])),
	))
	if result != Success {
		return result
	}
	return nil
}

func WaitForFences(device Device, fences []Fence, waitAll bool, timeout time.Duration) error {
	var _waitAll C.VkBool32 = C.VK_FALSE
	if waitAll {
		_waitAll = C.VK_TRUE
	}
	result := Result(C.vkWaitForFences(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.uint32_t)(len(fences)),
		(*C.VkFence)(unsafe.Pointer(&fences[0])),
		_waitAll,
		(C.uint64_t)(timeout),
	))
	if result != Success {
		return result
	}
	return nil
}

func CreateSemaphore(device Device, createInfo SemaphoreCreateInfo, allocator *AllocationCallbacks) (Semaphore, error) {
	var semaphore Semaphore
	result := Result(C.vkCreateSemaphore(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkSemaphoreCreateInfo)(unsafe.Pointer(&createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkSemaphore)(unsafe.Pointer(&semaphore)),
	))
	if result != Success {
		return 0, result
	}
	return semaphore, nil
}

func DestroySemaphore(device Device, semaphore Semaphore, allocator *AllocationCallbacks) {
	C.vkDestroySemaphore(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkSemaphore)(unsafe.Pointer(semaphore)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

func CreateEvent(device Device, createInfo EventCreateInfo, allocator *AllocationCallbacks) (Event, error) {
	var event Event
	result := Result(C.vkCreateEvent(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkEventCreateInfo)(unsafe.Pointer(&createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkEvent)(unsafe.Pointer(&event)),
	))
	if result != Success {
		return 0, result
	}
	return event, nil
}

func DestroyEvent(device Device, event Event, allocator *AllocationCallbacks) {
	C.vkDestroyEvent(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkEvent)(unsafe.Pointer(event)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

func GetEventStatus(device Device, event Event) error {
	result := Result(C.vkGetEventStatus(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkEvent)(unsafe.Pointer(event)),
	))
	if result != Success {
		return result
	}
	return nil
}

func SetEvent(device Device, event Event) error {
	result := Result(C.vkSetEventStatus(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkEvent)(unsafe.Pointer(event)),
	))
	if result != Success {
		return result
	}
	return nil
}

func ResetEvent(device Device, event Event) error {
	result := Result(C.vkResetEventStatus(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkEvent)(unsafe.Pointer(event)),
	))
	if result != Success {
		return result
	}
	return nil
}

func CmdSetEvent(commandBuffer CommandBuffer, event Event, stageMask PipelineStageFlags) error {
	result := Result(C.vkCmdSetEvent(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkEvent)(unsafe.Pointer(event)),
		(C.VkPipelineStageFlags)(stageMask),
	))
	if result != Success {
		return result
	}
	return nil
}

func CmdResetEvent(commandBuffer CommandBuffer, event Event, stageMask PipelineStageFlags) error {
	result := Result(C.vkCmdResetEvent(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkEvent)(unsafe.Pointer(event)),
		(C.VkPipelineStageFlags)(stageMask),
	))
	if result != Success {
		return result
	}
	return nil
}

func CmdWaitEvents(commandBuffer CommandBuffer, events []Event, srcStageMask, dstStageMask PipelineStageFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) {
	C.vkCmdWaitEvents(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.uint32_t)(len(events)),
		(*C.VkEvent)(unsafe.Pointer(&events[0])),
		(C.VkPipelineStageFlags)(srcStageMask),
		(C.VkPipelineStageFlags)(dstStageMask),
		(C.uint32_t)(len(memoryBarriers)),
		(*C.VkMemoryBarrier)(unsafe.Pointer(&memoryBarriers[0])),
		(C.uint32_t)(len(bufferMemoryBarriers)),
		(*C.VkBufferMemoryBarrier)(unsafe.Pointer(&bufferMemoryBarriers[0])),
		(C.uint32_t)(len(imageMemoryBarriers)),
		(*C.VkImageMemoryBarrier)(unsafe.Pointer(&imageMemoryBarriers[0])),
	)
}

func CmdPipelineBarrier(commandBuffer CommandBuffer, srcStageMask, dstStageMask PipelineStageFlags, dependencyFlags DependencyFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) {
	C.vkCmdPipelineBarrier(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkPipelineStageFlags)(srcStageMask),
		(C.VkPipelineStageFlags)(dstStageMask),
		(C.VkDependencyFlags)(dependencyFlags),
		(C.uint32_t)(len(memoryBarriers)),
		(*C.VkMemoryBarrier)(unsafe.Pointer(&memoryBarriers[0])),
		(C.uint32_t)(len(bufferMemoryBarriers)),
		(*C.VkBufferMemoryBarrier)(unsafe.Pointer(&bufferMemoryBarriers[0])),
		(C.uint32_t)(len(imageMemoryBarriers)),
		(*C.VkImageMemoryBarrier)(unsafe.Pointer(&imageMemoryBarriers[0])),
	)
}

func QueueWaitIdle(queue Queue) error {
	result := Result(C.vkQueueWaitIdle(
		(C.VkQueue)(unsafe.Pointer(queue)),
	))
	if result != Success {
		return result
	}
	return nil
}

func DeviceWaitIdle(device Device) error {
	result := Result(C.vkQueueWaitIdle(
		(C.VkDevice)(unsafe.Pointer(device)),
	))
	if result != Success {
		return result
	}
	return nil
}
