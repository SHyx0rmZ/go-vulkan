package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

var _not_implemented = errors.New("not implemented")

type DeviceSize uint64

const (
	NullHandle                = 0
	LodClampNone              = 1000.0
	RemainingMipLevels        = ^uint32(0)
	RemainingArrayLayers      = ^uint32(0)
	WholeSize                 = ^DeviceSize(0)
	AttachmentUnused          = ^uint32(0)
	True                      = uint32(1)
	False                     = uint32(0)
	QueueFamilyIgnored        = ^uint32(0)
	SubpassExternal           = ^uint32(0)
	MaxPhysicalDeviceNameSize = 256
	UUIDSize                  = 16
	MaxMemoryTypes            = 32
	MaxMemoryHeaps            = 16
	MaxExtensionNameSize      = 256
	MaxDescriptionSize        = 256
)

// 1.1
const (
	MaxDeviceGroupSize  = 32
	LUIDSize            = 8
	QueueFamilyExternal = ^uint32(0) - 1
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
)

const (
	ErrorOutOfPoolMemory                     Result = -1000069000
	ErrorInvalidExternalHandle               Result = -1000072003
	ErrorSurfaceLost                         Result = -1000000000
	ErrorNativeWindowInUse                   Result = -1000000001
	Suboptimal                               Result = -1000001003
	ErrorOutOfDate                           Result = -1000001004
	ErrorIncompatibleDisplay                 Result = -1000003001
	ErrorValidationFailed                    Result = -1000011001
	ErrorInvalidShader                       Result = -1000012000
	ErrorInvalidDRMFormatModifierPlaneLayout Result = -1000158000
	ErrorFragmentation                       Result = -1000161000
	ErrorNotPermitted                        Result = -1000174001
	ErrorInvalidDeviceAddress                Result = -1000244000
	ErrorFullScreenExclusiveModeLost         Result = -1000255000
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
	case ErrorOutOfPoolMemory:
		return "out of pool memory"
	case ErrorInvalidExternalHandle:
		return "invalid external handle"
	case ErrorSurfaceLost:
		return "surface lost"
	case ErrorNativeWindowInUse:
		return "native window in use"
	case Suboptimal:
		return "suboptimal"
	case ErrorOutOfDate:
		return "out of date"
	case ErrorIncompatibleDisplay:
		return "incompatible display"
	case ErrorValidationFailed:
		return "validation failed"
	case ErrorInvalidShader:
		return "invalid shader"
	case ErrorInvalidDRMFormatModifierPlaneLayout:
		return "invalid DRM format modifier plane layout"
	case ErrorFragmentation:
		return "fragmentation"
	case ErrorNotPermitted:
		return "not permitted"
	case ErrorInvalidDeviceAddress:
		return "invalid device address"
	case ErrorFullScreenExclusiveModeLost:
		return "full screen exclusive mode lost"
	default:
		panic(fmt.Sprintf("unknown result: %#v", r))
	}
}

type AllocationCallbacks C.VkAllocationCallbacks

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

type SubpassDescriptionFlags uint32

type PipelineBindPoint uint32

const (
	PipelineBindPointGraphics PipelineBindPoint = iota
	PipelineBindPointCompute
)

func (info *ShaderModuleCreateInfo) C(_info *shaderModuleCreateInfo) freeFunc {
	*_info = shaderModuleCreateInfo{
		Type:     info.Type,
		Next:     info.Next,
		Flags:    info.Flags,
		CodeSize: (C.size_t)(len(info.Code)),
	}
	p := C.CBytes(info.Code)
	_info.Code = (*byte)(p)
	return freeFunc(func() {
		C.free(p)
	})
}

type shaderModuleCreateInfo struct {
	Type     StructureType
	Next     uintptr
	Flags    ShaderModuleCreateFlags
	CodeSize C.size_t
	Code     *byte
}

func CreateShaderModule(device Device, createInfo ShaderModuleCreateInfo, allocator *AllocationCallbacks) (ShaderModule, error) {
	var shaderModule ShaderModule
	var _createInfo shaderModuleCreateInfo
	defer createInfo.C(&_createInfo).Free()
	result := Result(C.vkCreateShaderModule(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkShaderModuleCreateInfo)(unsafe.Pointer(&_createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkShaderModule)(unsafe.Pointer(&shaderModule)),
	))
	if result != Success {
		return 0, result
	}
	return shaderModule, nil
}

func DestroyShaderModule(device Device, shaderModule ShaderModule, allocator *AllocationCallbacks) {
	C.vkDestroyShaderModule(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkShaderModule)(unsafe.Pointer(shaderModule)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}
