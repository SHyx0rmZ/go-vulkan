package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

type FramebufferCreateFlags uint32

type FramebufferCreateInfo struct {
	Type        StructureType
	Next        uintptr
	Flags       FramebufferCreateFlags
	RenderPass  RenderPass
	Attachments []ImageView
	Width       uint32
	Height      uint32
	Layers      uint32
}

type framebufferCreateInfo struct {
	Type            StructureType
	Next            uintptr
	Flags           FramebufferCreateFlags
	RenderPass      RenderPass
	AttachmentCount uint32
	Attachments     *ImageView
	Width           uint32
	Height          uint32
	Layers          uint32
}

type Rect2D struct {
	Offset Offset2D
	Extent Extent2D
}

type Framebuffer uint64

type RenderPassBeginInfo struct {
	Type        StructureType
	Next        uintptr
	RenderPass  RenderPass
	Framebuffer Framebuffer
	RenderArea  Rect2D
	ClearValues []ClearValue
}

func (info *RenderPassBeginInfo) C(_info *renderPassBeginInfo) freeFunc {
	*_info = renderPassBeginInfo{
		Type:            info.Type,
		Next:            info.Next,
		RenderPass:      info.RenderPass,
		Framebuffer:     info.Framebuffer,
		RenderArea:      info.RenderArea,
		ClearValueCount: uint32(len(info.ClearValues)),
		ClearValues:     nil,
	}
	if _info.ClearValueCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.ClearValueCount) * unsafe.Sizeof(clearValue{})))
		for i, cv := range info.ClearValues {
			_clearValue := clearValue{
				//DepthStencil: cv.DepthStencil,
			}
			switch color := cv.Color.(type) {
			case ClearColorValueFloat:
				*(*float32)(unsafe.Pointer(&_clearValue.Color[0])) = color[0]
				*(*float32)(unsafe.Pointer(&_clearValue.Color[1])) = color[1]
				*(*float32)(unsafe.Pointer(&_clearValue.Color[2])) = color[2]
				*(*float32)(unsafe.Pointer(&_clearValue.Color[3])) = color[3]
			case ClearColorValueInt:
				*(*int32)(unsafe.Pointer(&_clearValue.Color[0])) = color[0]
				*(*int32)(unsafe.Pointer(&_clearValue.Color[1])) = color[1]
				*(*int32)(unsafe.Pointer(&_clearValue.Color[2])) = color[2]
				*(*int32)(unsafe.Pointer(&_clearValue.Color[3])) = color[3]
			case ClearColorValueUint:
				*(*uint32)(unsafe.Pointer(&_clearValue.Color[0])) = color[0]
				*(*uint32)(unsafe.Pointer(&_clearValue.Color[1])) = color[1]
				*(*uint32)(unsafe.Pointer(&_clearValue.Color[2])) = color[2]
				*(*uint32)(unsafe.Pointer(&_clearValue.Color[3])) = color[3]
			}
			*(*clearValue)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(clearValue{}))) = _clearValue
		}
		_info.ClearValues = (*clearValue)(p)
		return freeFunc(func() {
			C.free(p)
		})
	}
	return freeFunc(nil)
}

type renderPassBeginInfo struct {
	Type            StructureType
	Next            uintptr
	RenderPass      RenderPass
	Framebuffer     Framebuffer
	RenderArea      Rect2D
	ClearValueCount uint32
	ClearValues     *clearValue
}

type SubpassContents uint32

const (
	SubpassContentsInline SubpassContents = iota
	SubpassContentsSecondaryCommandBuffers
)

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
			for i := len(ps); i > 0; i-- {
				C.free(ps[i-1])
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
	result := Result(C.vkCreateRenderPass(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkRenderPassCreateInfo)(unsafe.Pointer(&_createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkRenderPass)(unsafe.Pointer(&renderPass)),
	))
	if result != Success {
		return 0, result
	}
	return renderPass, nil
}

func DestroyRenderPass(device Device, renderPass RenderPass, allocator *AllocationCallbacks) {
	C.vkDestroyRenderPass(
		(C.VkDevice)(unsafe.Pointer(device)),
		*(*C.VkRenderPass)(unsafe.Pointer(&renderPass)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

func CreateFramebuffer(device Device, createInfo FramebufferCreateInfo, allocator *AllocationCallbacks) (Framebuffer, error) {
	var framebuffer Framebuffer
	_createInfo := framebufferCreateInfo{
		Type:            createInfo.Type,
		Next:            createInfo.Next,
		Flags:           createInfo.Flags,
		RenderPass:      createInfo.RenderPass,
		AttachmentCount: uint32(len(createInfo.Attachments)),
		Width:           createInfo.Width,
		Height:          createInfo.Height,
		Layers:          createInfo.Layers,
	}
	if _createInfo.AttachmentCount > 0 {
		p := C.malloc(C.size_t(uintptr(_createInfo.AttachmentCount) * unsafe.Sizeof(ImageView(0))))
		defer C.free(p)
		for i, attachment := range createInfo.Attachments {
			*(*ImageView)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(ImageView(0)))) = attachment
		}
		_createInfo.Attachments = (*ImageView)(p)
	}
	result := Result(C.vkCreateFramebuffer(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkFramebufferCreateInfo)(unsafe.Pointer(&_createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkFramebuffer)(unsafe.Pointer(&framebuffer)),
	))
	if result != Success {
		return 0, result
	}
	return framebuffer, nil
}

func DestroyFramebuffer(device Device, framebuffer Framebuffer, allocator *AllocationCallbacks) {
	C.vkDestroyFramebuffer(
		(C.VkDevice)(unsafe.Pointer(device)),
		*(*C.VkFramebuffer)(unsafe.Pointer(&framebuffer)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

func CmdBeginRenderPass(commandBuffer CommandBuffer, beginInfo RenderPassBeginInfo, contents SubpassContents) {
	var _beginInfo renderPassBeginInfo
	beginInfo.C(&_beginInfo)
	//defer beginInfo.C(&_beginInfo).Free()
	C.vkCmdBeginRenderPass(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		(*C.VkRenderPassBeginInfo)(unsafe.Pointer(&_beginInfo)),
		(C.VkSubpassContents)(contents),
	)
}

func GetRenderAreaGranularity(device Device, renderPass RenderPass) Extent2D {
	var granularity Extent2D
	C.vkGetRenderAreaGranularity(
		(C.VkDevice)(unsafe.Pointer(device)),
		*(*C.VkRenderPass)(unsafe.Pointer(&renderPass)),
		(*C.VkExtent2D)(unsafe.Pointer(&granularity)),
	)
	return granularity
}

func CmdNextSubpass(commandBuffer CommandBuffer, contents SubpassContents) {
	C.vkCmdNextSubpass(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		(C.VkSubpassContents)(contents),
	)
}

func CmdEndRenderPass(commandBuffer CommandBuffer) {
	C.vkCmdEndRenderPass(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
	)
}
