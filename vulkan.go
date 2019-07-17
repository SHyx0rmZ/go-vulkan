package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type Result int32

const (
	Success Result = iota
	NotReady
	Timeout
	EventSet
	EventReset
	Incomplete
)

const (
	ErrorOutOfHostMemory Result = -iota - 1
	ErrorOutOfDeviceMemory
	ErrorInitializationFailed
	ErrorDeviceLost
	ErrorMemoryMapFailed
	ErrorLayerNotPresent
	ErrorExtensionNotPresent
	ErrorFeatureNotPresent
	ErrorIncompatibleDriver
	ErrorTooManyObjects
	ErrorFormatNotSupported
	ErrorFragmentedPool
	ErrorOutOfPoolMemory
	ErrorInvalidExternalHandle
)

func (r Result) Error() string {
	switch r {
	case Success:
		return "success"
	case NotReady:
		return "not ready"
	case Timeout:
		return "timeout"
	case EventSet:
		return "event set"
	case EventReset:
		return "event reset"
	case Incomplete:
		return "incomplete"
	case ErrorOutOfHostMemory:
		return "out of host memory"
	case ErrorOutOfDeviceMemory:
		return "out of device memory"
	case ErrorInitializationFailed:
		return "initialization failed"
	case ErrorDeviceLost:
		return "device lost"
	case ErrorMemoryMapFailed:
		return "memory map failed"
	case ErrorLayerNotPresent:
		return "layer not present"
	case ErrorExtensionNotPresent:
		return "extension not present"
	case ErrorFeatureNotPresent:
		return "feature not present"
	case ErrorIncompatibleDriver:
		return "incompatible driver"
	case ErrorTooManyObjects:
		return "too many objects"
	case ErrorFormatNotSupported:
		return "format not supported"
	case ErrorFragmentedPool:
		return "fragmented pool"
	//case ErrorOutOfPoolMemory:
	//	return "out of pool memory"
	//case ErrorInvalidExternalHandle:
	//	return "invalid external handle"
	default:
		panic(fmt.Sprintf("unknown result: %#v", r))
	}
}

type StructureType uint32

// cat s.txt | head -n 49 | sed 's/[A-Z]/\L&/g;s/_\([a-z]\)/\U\1/g;s/^\s\+vk//;s/\s=\s[0-9]\+,$//'

const (
	StructureTypeApplicationInfo StructureType = iota
	StructureTypeInstanceCreateInfo
	StructureTypeDeviceQueueCreateInfo
	StructureTypeDeviceCreateInfo
	StructureTypeSubmitInfo
	StructureTypeMemoryAllocateInfo
	StructureTypeMappedMemoryRange
	StructureTypeBindSparseInfo
	StructureTypeFenceCreateInfo
	StructureTypeSemaphoreCreateInfo
	StructureTypeEventCreateInfo
	StructureTypeQueryPoolCreateInfo
	StructureTypeBufferCreateInfo
	StructureTypeBufferViewCreateInfo
	StructureTypeImageCreateInfo
	StructureTypeImageViewCreateInfo
	StructureTypeShaderModuleCreateInfo
	StructureTypePipelineCacheCreateInfo
	StructureTypePipelineShaderStageCreateInfo
	StructureTypePipelineVertexInputStateCreateInfo
	StructureTypePipelineInputAssemblyStateCreateInfo
	StructureTypePipelineTessellationStateCreateInfo
	StructureTypePipelineViewportStateCreateInfo
	StructureTypePipelineRasterizationStateCreateInfo
	StructureTypePipelineMultisampleStateCreateInfo
	StructureTypePipelineDepthStencilStateCreateInfo
	StructureTypePipelineColorBlendStateCreateInfo
	StructureTypePipelineDynamicStateCreateInfo
	StructureTypeGraphicsPipelineCreateInfo
	StructureTypeComputePipelineCreateInfo
	StructureTypePipelineLayoutCreateInfo
	StructureTypeSamplerCreateInfo
	StructureTypeDescriptorSetLayoutCreateInfo
	StructureTypeDescriptorPoolCreateInfo
	StructureTypeDescriptorSetAllocateInfo
	StructureTypeWriteDescriptorSet
	StructureTypeCopyDescriptorSet
	StructureTypeFramebufferCreateInfo
	StructureTypeRenderPassCreateInfo
	StructureTypeCommandPoolCreateInfo
	StructureTypeCommandBufferAllocateInfo
	StructureTypeCommandBufferInheritanceInfo
	StructureTypeCommandBufferBeginInfo
	StructureTypeRenderPassBeginInfo
	StructureTypeBufferMemoryBarrier
	StructureTypeImageMemoryBarrier
	StructureTypeMemoryBarrier
	StructureTypeLoaderInstanceCreateInfo
	StructureTypeLoaderDeviceCreateInfo
)

type AllocationCallbacks C.VkAllocationCallbacks

type ImageViewCreateInfo struct {
	Type             StructureType
	Next             uintptr
	Flags            C.VkImageViewCreateFlags
	Image            Image
	ViewType         ImageViewType
	Format           C.VkFormat
	Components       ComponentMapping
	SubresourceRange ImageSubresourceRange
}

type ImageView uintptr

type ImageViewType uint32

const (
	ImageViewType1D ImageViewType = iota
	ImageViewType2D
	ImageViewType3D
	ImageViewTypeCube
	ImageViewType1DArray
	ImageViewType2DArray
	ImageViewTypeCubeArray
)

type ComponentMapping struct {
	R ComponentSwizzle
	G ComponentSwizzle
	B ComponentSwizzle
	A ComponentSwizzle
}

type ComponentSwizzle uint32

const (
	ComponentSwizzleIdentity ComponentSwizzle = iota
	ComponentSwizzleZero
	ComponentSwizzleOne
	ComponentSwizzleR
	ComponentSwizzleG
	ComponentSwizzleB
	ComponentSwizzleA
)

type ImageSubresourceRange struct {
	AspectMask     ImageAspectFlags
	BaseMIPLevel   uint32
	LevelCount     uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

type ImageAspectFlags uint32

const (
	ImageAspectColorBit ImageAspectFlags = 1 << iota
	ImageAspectDepthBit
	ImageAspectStencilBit
	ImageAspectMetadataBit
	ImageAspectPlane0Bit
	ImageAspectPlane1Bit
	ImageAspectPlane2Bit
)

func CreateImageView(device Device, createInfo ImageViewCreateInfo, allocator *AllocationCallbacks) (ImageView, error) {
	var view ImageView
	result := C.vkCreateImageView(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkImageViewCreateInfo)(unsafe.Pointer(&createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkImageView)(unsafe.Pointer(&view)),
	)
	if result != C.VK_SUCCESS {
		return 0, fmt.Errorf("CreateImageView")
	}
	return view, nil
}

func DestroyImageView(device Device, imageView ImageView, allocator *AllocationCallbacks) {
	C.vkDestroyImageView(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkImageView)(unsafe.Pointer(imageView)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

type RenderPass uintptr

type RenderPassCreateInfo struct {
	Type         C.VkStructureType
	Next         uintptr
	Flags        C.VkRenderPassCreateFlags
	Attachments  []AttachmentDescription
	Subpasses    []SubpassDescription
	Dependencies []SubpassDependency
}

type renderPassCreateInfo struct {
	Type            C.VkStructureType
	Next            uintptr
	Flags           C.VkRenderPassCreateFlags
	AttachmentCount uint32
	Attachments     *AttachmentDescription
	SubpassCount    uint32
	Subpasses       *subpassDescription
	DependencyCount uint32
	Dependencies    *SubpassDependency
}

type AttachmentDescription struct {
	Flags          AttachmentDescriptionFlags
	Format         C.VkFormat
	Samples        SampleCountFlagBits
	LoadOp         AttachmentLoadOp
	StoreOp        AttachmentStoreOp
	StencilLoadOp  AttachmentLoadOp
	StencilStoreOp AttachmentStoreOp
	InitialLayout  ImageLayout
	FinalLayout    ImageLayout
}

type AttachmentDescriptionFlags uint32

const (
	AttachmentDescriptionMayAliasBit AttachmentDescriptionFlags = 1 << iota
)

type AttachmentLoadOp uint32

const (
	AttachmentLoadOpLoad AttachmentLoadOp = iota
	AttachmentLoadOpClear
	AttachmentLoadOpDontCare
)

type AttachmentStoreOp uint32

const (
	AttachmentStoreOpStore AttachmentStoreOp = iota
	AttachmentStoreOpDontCare
)

type SubpassDescription struct {
	Flags                  SubpassDescriptionFlags
	PipelineBindPoint      PipelineBindPoint
	InputAttachments       []AttachmentReference
	ColorAttachments       []AttachmentReference
	ResolveAttachments     []AttachmentReference
	DepthStencilAttachment AttachmentReference
	PreserveAttachments    []uint32
}

type subpassDescription struct {
	Flags                   SubpassDescriptionFlags
	PipelineBindPoint       PipelineBindPoint
	InputAttachmentCount    uint32
	InputAttachments        *AttachmentReference
	ColorAttachmentCount    uint32
	ColorAttachments        *AttachmentReference
	ResolveAttachments      *AttachmentReference
	DepthStencilAttachment  *AttachmentReference
	PreserveAttachmentCount uint32
	PreserveAttachments     *uint32
}

type AttachmentReference struct {
	Attachment uint32
	Layout     ImageLayout
}

type SubpassDependency struct {
	SrcSubpass      uint32
	DstSubpass      uint32
	SrcStageMask    PipelineStageFlags
	DstStageMask    PipelineStageFlags
	SrcAccessMask   AccessFlags
	DstAccessMask   AccessFlags
	DependencyFlags DependencyFlags
}

type SubpassDescriptionFlags uint32

type PipelineBindPoint uint32

const (
	PipelineBindPointGraphics PipelineBindPoint = iota
	PipelineBindPointCompute
)

type PipelineStageFlagBits uint32
type PipelineStageFlags PipelineStageFlagBits

const (
	PipelineStageTopOfPipeBit PipelineStageFlagBits = 1 << iota
	PipelineStageDrawIndirectBit
	PipelineStageVertexInputBit
	PipelineStageVertexShaderBit
	PipelineStageTessellationControlShaderBit
	PipelineStageTessellationEvaluationShaderBit
	PipelineStageGeometryShaderBit
	PipelineStageFragmentShaderBit
	PipelineStageEarlyFragmentTestsBit
	PipelineStageLateFragmentTestsBit
	PipelineStageColorAttachmentOutputBit
	PipelienStageComputeShaderBit
	PipelineStageTransferBit
	PipelineStageBottomOfPipeBit
	PipelineStageHostBit
	PipelineStageAllGraphicsBit
	PipelineStageAllCommandsBit
)

type AccessFlagBits uint32
type AccessFlags AccessFlagBits

const (
	AccessIndirectCommandReadBit AccessFlagBits = 1 << iota
	AccessIndexReadBit
	AccessVertexAttributeReadBit
	AccessUniformReadBit
	AccessInputAttachmentReadBit
	AccessShaderReadBit
	AccessShaderWriteBit
	AccessColorAttachmentReadBit
	AccessColorAttachmentWriteBit
	AccessDepthStencilAttachmentReadBit
	AccessDepthStencilAttachmentWriteBit
	AccessTransferReadBit
	AccessTransferWriteBit
	AccessHostReadBit
	AccessHostWriteBit
	AccessMemoryReadBit
	AccessMemoryWriteBit
)

type DependencyFlagBits uint32
type DependencyFlags DependencyFlagBits

const (
	DependencyByRegionBit DependencyFlagBits = 1 << iota
	DependencyDeviceGroupBit
	DependencyViewLocalBit
)

type ImageLayout uint32

const (
	ImageLayoutUndefined ImageLayout = iota
	ImageLayoutGeneral
	ImageLayoutColorAttachmentOptimal
	ImageLayoutDepthStencilAttachmentOptimal
	ImageLayoutDepthStencilReadOnlyOptimal
	ImageLayoutShaderReadOnlyOptimal
	ImageLayoutTranserSrcOptimal
	ImageLayoutTransferDstOptimal
	ImageLayoutPreinitialized
)

const (
	ImageLayoutDepthReadOnlyStencilAttachmentOptimal ImageLayout = 1000117000 + iota
	ImageLayoutDepthAttachmentStencilReadOnlyOptimal
)

const (
	ImageLayoutPresentSrcKHR    ImageLayout = 1000001002
	ImageLayoutSharedPresentKHR ImageLayout = 1000111000
)

type SampleCountFlagBits uint32
type SampleCountFlags SampleCountFlagBits

const (
	SampleCount1Bit SampleCountFlagBits = 1 << iota
	SampleCount2Bit
	SampleCount4Bit
	SampleCount8Bit
	SampleCount16Bit
	SampleCount32Bit
	SampleCount64Bit
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
