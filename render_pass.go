package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type FramebufferCreateInfo struct{}

type Framebuffer uintptr

type RenderPassBeginInfo struct{}

type SubpassContents struct{}

func CreateRenderPass(device Device, createInfo RenderPassCreateInfo, allocator *AllocationCallbacks) (RenderPass, error) {
	var renderPass RenderPass
	_createInfo := renderPassCreateInfo{
		Type:            C.VK_STRUCTURE_TYPE_RENDER_PASS_CREATE_INFO,
		Next:            createInfo.Next,
		Flags:           createInfo.Flags,
		AttachmentCount: uint32(len(createInfo.Attachments)),
		SubpassCount:    uint32(len(createInfo.Subpasses)),
		DependencyCount: uint32(len(createInfo.Dependencies)),
	}
	if _createInfo.AttachmentCount > 0 {
		p := C.malloc(C.size_t(uintptr(_createInfo.AttachmentCount) * unsafe.Sizeof(AttachmentDescription{})))
		defer C.free(p)
		for i, attachment := range createInfo.Attachments {
			*(*AttachmentDescription)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(AttachmentDescription{}))) = attachment
		}
		_createInfo.Attachments = (*AttachmentDescription)(p)
	}
	if _createInfo.SubpassCount > 0 {
		p := C.malloc(C.size_t(uintptr(_createInfo.SubpassCount) * unsafe.Sizeof(subpassDescription{})))
		defer C.free(p)
		var ps []unsafe.Pointer
		defer func() {
			for _, p := range ps {
				C.free(p)
			}
		}()
		for i, subpass := range createInfo.Subpasses {
			_subpass := subpassDescription{
				Flags:                   subpass.Flags,
				PipelineBindPoint:       subpass.PipelineBindPoint,
				InputAttachmentCount:    uint32(len(subpass.InputAttachments)),
				ColorAttachmentCount:    uint32(len(subpass.ColorAttachments)),
				PreserveAttachmentCount: uint32(len(subpass.PreserveAttachments)),
			}
			sp := C.malloc(C.size_t(unsafe.Sizeof(AttachmentReference{})))
			ps = append(ps, sp)
			*(*AttachmentReference)(sp) = subpass.DepthStencilAttachment
			_subpass.DepthStencilAttachment = (*AttachmentReference)(sp)
			if _subpass.InputAttachmentCount > 0 {
				sp := C.malloc(C.size_t(uintptr(_subpass.InputAttachmentCount) * unsafe.Sizeof(AttachmentReference{})))
				ps = append(ps, sp)
				for i, reference := range subpass.InputAttachments {
					*(*AttachmentReference)(unsafe.Pointer(uintptr(sp) + uintptr(i)*unsafe.Sizeof(AttachmentReference{}))) = reference
				}
				_subpass.InputAttachments = (*AttachmentReference)(sp)
			}
			if _subpass.ColorAttachmentCount > 0 {
				sp := C.malloc(C.size_t(uintptr(_subpass.ColorAttachmentCount) * unsafe.Sizeof(AttachmentReference{})))
				ps = append(ps, sp)
				for i, reference := range subpass.ColorAttachments {
					*(*AttachmentReference)(unsafe.Pointer(uintptr(sp) + uintptr(i)*unsafe.Sizeof(AttachmentReference{}))) = reference
				}
				_subpass.ColorAttachments = (*AttachmentReference)(sp)
				if len(subpass.ResolveAttachments) != 0 && len(subpass.ResolveAttachments) != len(subpass.ColorAttachments) {
					// don't do anything, just let it crash
					// return 0, fmt.Errorf("count mismatch") // todo
				}
				if len(subpass.ResolveAttachments) > 0 {
					sp := C.malloc(C.size_t(uintptr(_subpass.ColorAttachmentCount) * unsafe.Sizeof(AttachmentReference{})))
					ps = append(ps, sp)
					for i, reference := range subpass.ResolveAttachments {
						*(*AttachmentReference)(unsafe.Pointer(uintptr(sp) + uintptr(i)*unsafe.Sizeof(AttachmentReference{}))) = reference
					}
					_subpass.ResolveAttachments = (*AttachmentReference)(sp)
				}
			}
			if _subpass.PreserveAttachmentCount > 0 {
				sp := C.malloc(C.size_t(uintptr(_subpass.PreserveAttachmentCount) * unsafe.Sizeof(uint32(0))))
				ps = append(ps, sp)
				for i, attachment := range subpass.PreserveAttachments {
					*(*uint32)(unsafe.Pointer(uintptr(sp) + uintptr(i)*unsafe.Sizeof(uint32(0)))) = attachment
				}
				_subpass.PreserveAttachments = (*uint32)(sp)
			}
			*(*subpassDescription)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(subpassDescription{}))) = _subpass
		}
		_createInfo.Subpasses = (*subpassDescription)(p)
	}
	if _createInfo.DependencyCount > 0 {
		p := C.malloc(C.size_t(uintptr(_createInfo.DependencyCount) * unsafe.Sizeof(SubpassDependency{})))
		defer C.free(p)
		for i, dependency := range createInfo.Dependencies {
			*(*SubpassDependency)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(SubpassDependency{}))) = dependency
		}
		_createInfo.Dependencies = (*SubpassDependency)(p)
	}
	result := C.vkCreateRenderPass(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkRenderPassCreateInfo)(unsafe.Pointer(&_createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkRenderPass)(unsafe.Pointer(&renderPass)),
	)
	if result != C.VK_SUCCESS {
		return 0, fmt.Errorf("CreateRenderPass")
	}
	return renderPass, nil
}

func DestroyRenderPass(device Device, renderPass RenderPass, allocator *AllocationCallbacks) {
	C.vkDestroyRenderPass(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkRenderPass)(unsafe.Pointer(renderPass)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

func CreateFramebuffer(device Device, createInfo FramebufferCreateInfo, allocator *AllocationCallbacks) (Framebuffer, error) {
	return 0, nil
}

func DestroyFramebuffer(device Device, framebuffer Framebuffer, allocator *AllocationCallbacks) {

}

func CmdBeginRenderPass(commandBuffer CommandBuffer, renderPassBegin RenderPassBeginInfo, contents SubpassContents) {

}

func GetRenderAreaGranularity(device Device, renderPass RenderPass, granularity Extent2D) {

}

func CmdNextSubpass(commandBuffer CommandBuffer, contents SubpassContents) {

}

func CmdEndRenderPass(commandBuffer CommandBuffer) {

}
