package vulkan

import (
	"time"
)

type Fence uintptr

type FenceCreateInfo struct{}

type EventCreateInfo struct{}

type Event uintptr

type MemoryBarrier uintptr

type BufferMemoryBarrier uintptr

type ImageMemoryBarrier uintptr

func CreateFence(device Device, createInfo FenceCreateInfo, allocator *AllocationCallbacks) (Fence, error) {
	return 0, nil
}

func DestroyFence(device Device, fence Fence, allocator *AllocationCallbacks) {

}

func GetFenceStatus(device Device, fence Fence) error {
	return nil
}

func ResetFences(device Device, fences []Fence) error {
	return nil
}

func WaitForFences(device Device, fences []Fence, waitAll bool, timeout time.Duration) error {
	return nil
}

func CreateSemaphore(device Device, createInfo SemaphoreCreateInfo, allocator *AllocationCallbacks) (Semaphore, error) {
	return 0, nil
}

func DestroySemaphore(device Device, semaphore Semaphore, allocator *AllocationCallbacks) {

}

func CreateEvent(device Device, createInfo EventCreateInfo, allocator *AllocationCallbacks) (Event, error) {
	return 0, nil
}

func DestroyEvent(device Device, event Event, allocator *AllocationCallbacks) {

}

func GetEventStatus(device Device, event Event) error {
	return nil
}

func SetEvent(device Device, event Event) error {
	return nil
}

func ResetEvent(device Device, event Event) error {
	return nil
}

func CmdSetEvent(commandBuffer CommandBuffer, event Event, stageMask PipelineStageFlags) error {
	return nil
}

func CmdResetEvent(commandBuffer CommandBuffer, event Event, stageMask PipelineStageFlags) error {
	return nil
}

func CmdWaitEvents(commandBuffer CommandBuffer, events []Event, srcStageMask, dstStageMask PipelineStageFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) {

}

func CmdPipelineBarrier(commandBuffer CommandBuffer, srcStageMask, dstStageMask PipelineStageFlags, dependencyFlags DependencyFlags, memoryBarriers []MemoryBarrier, bufferMemoryBarriers []BufferMemoryBarrier, imageMemoryBarriers []ImageMemoryBarrier) {

}

func QueueWaitIdle(queue Queue) error {
	return nil
}

func DeviceWaitIdle(device Device) error {
	return nil
}
