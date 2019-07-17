package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"

type CommandBufferAllocateInfo struct{}

type CommandBufferBeginInfo struct{}

type CommandBufferInheritanceInfo struct{}

type SubmitInfo struct{}

type CommandPoolCreateInfo struct{}

type CommandPool uintptr

type CommandBuffer uintptr

type CommandPoolTrimFlags uint32

type CommandPoolResetFlags uint32

type CommandBufferResetFlags uint32

func CreateCommandPool(device Device, createInfo CommandPoolCreateInfo, allocator *AllocationCallbacks) (CommandPool, error) {
	return 0, nil
}

func TrimCommandPool(device Device, commandPool CommandPool, flags CommandPoolTrimFlags) {
}

func ResetCommandPool(device Device, commandPool CommandPool, flags CommandPoolResetFlags) error {
	return nil
}

func DestroyCommandPool(device Device, commandPool CommandPool, allocator *AllocationCallbacks) {

}

func AllocateCommandBuffers(device Device, allocateInfo CommandBufferAllocateInfo) ([]CommandBuffer, error) {
	return nil, nil
}

func ResetCommandBuffer(commandBuffer CommandBuffer, flags CommandBufferResetFlags) error {
	return nil
}

func FreeCommandBuffers(device Device, commandPool CommandPool, commandBuffers []CommandBuffer) {

}

func BeginCommandBuffer(commandBuffer CommandBuffer, beginInfo CommandBufferBeginInfo) error {
	return nil
}

func EndCommandBuffer(commandBuffer CommandBuffer) error {
	return nil
}

func QueueSubmit(queue Queue, submits []SubmitInfo, fence Fence) error {
	return nil
}

func CmdExecuteCommands(commandBuffer CommandBuffer, commandBuffers []CommandBuffer) {

}

func CmdSetDeviceMask(commandBuffer CommandBuffer, deviceMask uint32) {

}
