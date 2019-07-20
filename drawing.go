package vulkan

// #include <vulkan/vulkan.h>
import "C"
import (
	"unsafe"
)

func CmdDraw(commandBuffer CommandBuffer, vertexCount, instanceCount, firstVertex, firstInstance uint32) {
	C.vkCmdDraw(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.uint32_t)(vertexCount),
		(C.uint32_t)(instanceCount),
		(C.uint32_t)(firstVertex),
		(C.uint32_t)(firstInstance),
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
				Color:        ClearColorValueUint{0, 0, 0, 255},
				DepthStencil: ClearDepthStencilValue{},
			},
		}
	}
	C.vkCmdClearAttachments(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.uint32_t)(len(clearAttachments)),
		(*C.VkClearAttachment)(unsafe.Pointer(&_clearAttachments[0])),
		(C.uint32_t)(len(rects)),
		(*C.VkClearRect)(unsafe.Pointer(&rects[0])),
	)
}
