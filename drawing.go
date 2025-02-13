package vulkan

// #include <vulkan/vulkan.h>
import "C"
import (
	"unsafe"
)

func CmdDraw(commandBuffer CommandBuffer, vertexCount, instanceCount, firstVertex, firstInstance uint32) {
	C.vkCmdDraw(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		(C.uint32_t)(vertexCount),
		(C.uint32_t)(instanceCount),
		(C.uint32_t)(firstVertex),
		(C.uint32_t)(firstInstance),
	)
}

func CmdDrawIndexed(commandBuffer CommandBuffer, indexCount, instanceCount, firstIndex uint32, vertexOffset int32, firstInstance uint32) {
	C.vkCmdDrawIndexed(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		(C.uint32_t)(indexCount),
		(C.uint32_t)(instanceCount),
		(C.uint32_t)(firstIndex),
		(C.int32_t)(vertexOffset),
		(C.uint32_t)(firstInstance),
	)
}

func CmdDrawIndirect(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, drawCount, stride uint32) {
	C.vkCmdDrawIndirect(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		*(*C.VkBuffer)(unsafe.Pointer(&buffer)),
		(C.VkDeviceSize)(offset),
		(C.uint32_t)(drawCount),
		(C.uint32_t)(stride),
	)
}

func CmdDrawIndexedIndirect(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, drawCount, stride uint32) {
	C.vkCmdDrawIndexedIndirect(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		*(*C.VkBuffer)(unsafe.Pointer(&buffer)),
		(C.VkDeviceSize)(offset),
		(C.uint32_t)(drawCount),
		(C.uint32_t)(stride),
	)
}

func CmdDrawIndirectCount(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount, stride uint32) {
	C.vkCmdDrawIndirectCount(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		*(*C.VkBuffer)(unsafe.Pointer(&buffer)),
		(C.VkDeviceSize)(offset),
		*(*C.VkBuffer)(unsafe.Pointer(&countBuffer)),
		(C.VkDeviceSize)(countBufferOffset),
		(C.uint32_t)(maxDrawCount),
		(C.uint32_t)(stride),
	)
}

func CmdDrawIndexedIndirectCount(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize, countBuffer Buffer, countBufferOffset DeviceSize, maxDrawCount, stride uint32) {
	C.vkCmdDrawIndexedIndirectCount(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		*(*C.VkBuffer)(unsafe.Pointer(&buffer)),
		(C.VkDeviceSize)(offset),
		*(*C.VkBuffer)(unsafe.Pointer(&countBuffer)),
		(C.VkDeviceSize)(countBufferOffset),
		(C.uint32_t)(maxDrawCount),
		(C.uint32_t)(stride),
	)
}

func CmdDispatch(commandBuffer CommandBuffer, groupCountX, groupCountY, groupCountZ uint32) {
	C.vkCmdDispatch(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		(C.uint32_t)(groupCountX),
		(C.uint32_t)(groupCountY),
		(C.uint32_t)(groupCountZ),
	)
}

func CmdDispatchIndirect(commandBuffer CommandBuffer, buffer Buffer, offset DeviceSize) {
	C.vkCmdDispatchIndirect(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		*(*C.VkBuffer)(unsafe.Pointer(&buffer)),
		(C.VkDeviceSize)(offset),
	)
}

func CmdDispatchBase(commandBuffer CommandBuffer, baseGroupX, baseGroupY, baseGroupZ, groupCountX, groupCountY, groupCountZ uint32) {
	C.vkCmdDispatchBase(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		(C.uint32_t)(baseGroupX),
		(C.uint32_t)(baseGroupY),
		(C.uint32_t)(baseGroupZ),
		(C.uint32_t)(groupCountX),
		(C.uint32_t)(groupCountY),
		(C.uint32_t)(groupCountZ),
	)
}

type ClearAttachment struct {
	AspectMask      ImageAspectFlags
	ColorAttachment uint32
	ClearValue      ClearValue
}

type clearAttachment struct {
	AspectMask      ImageAspectFlags
	ColorAttachment uint32
	ClearValue      clearValue
}

type ClearRect struct {
	Rect           Rect2D
	BaseArrayLayer uint32
	LayerCount     uint32
}

func CmdClearAttachments(commandBuffer CommandBuffer, clearAttachments []ClearAttachment, rects []ClearRect) {
	_clearAttachments := make([]clearAttachment, len(clearAttachments))
	for i, ca := range clearAttachments {
		_clearAttachments[i] = clearAttachment{
			AspectMask:      ca.AspectMask,
			ColorAttachment: ca.ColorAttachment,
			ClearValue: clearValue{
				//Color: ClearColorValueUint{255, 0, 0, 255},
				Color: ClearColorValueFloat{1, 0, 0, 1},
				//DepthStencil: ClearDepthStencilValue{},
			},
		}
	}
	C.vkCmdClearAttachments(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		(C.uint32_t)(len(clearAttachments)),
		(*C.VkClearAttachment)(unsafe.Pointer(&_clearAttachments[0])),
		(C.uint32_t)(len(rects)),
		(*C.VkClearRect)(unsafe.Pointer(&rects[0])),
	)
}

func CmdPushConstants(commandBuffer CommandBuffer, layout PipelineLayout, stageFlags ShaderStageFlags, offset, size uint32, values unsafe.Pointer) {
	C.vkCmdPushConstants(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		*(*C.VkPipelineLayout)(unsafe.Pointer(&layout)),
		(C.VkShaderStageFlags)(stageFlags),
		(C.uint32_t)(offset),
		(C.uint32_t)(size),
		values,
	)
}
