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

type QueryPool uintptr

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
	MaxDriverNameSize         = 256
	MaxDriverInfoSize         = 256
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
	ErrorOutOfPoolMemory                        Result = -1000069000
	ErrorInvalidExternalHandle                  Result = -1000072003
	ErrorSurfaceLostKHR                         Result = -1000000000
	ErrorNativeWindowInUseKHR                   Result = -1000000001
	SuboptimalKHR                               Result = -1000001003
	ErrorOutOfDateKHR                           Result = -1000001004
	ErrorIncompatibleDisplayKHR                 Result = -1000003001
	ErrorValidationFailedEXT                    Result = -1000011001
	ErrorInvalidShaderNV                        Result = -1000012000
	ErrorInvalidDRMFormatModifierPlaneLayoutEXT Result = -1000158000
	ErrorFragmentationEXT                       Result = -1000161000
	ErrorNotPermittedEXT                        Result = -1000174001
	ErrorInvalidDeviceAddressEXT                Result = -1000244000
	ErrorFullScreenExclusiveModeLostEXT         Result = -1000255000
)

const (
	ErrorOutOfPoolMemoryKHR       = ErrorOutOfPoolMemory
	ErrorInvalidExternalHandleKHR = ErrorInvalidExternalHandle
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
	case ErrorSurfaceLostKHR:
		return "surface lost"
	case ErrorNativeWindowInUseKHR:
		return "native window in use"
	case SuboptimalKHR:
		return "suboptimal"
	case ErrorOutOfDateKHR:
		return "out of date"
	case ErrorIncompatibleDisplayKHR:
		return "incompatible display"
	case ErrorValidationFailedEXT:
		return "validation failed"
	case ErrorInvalidShaderNV:
		return "invalid shader"
	case ErrorInvalidDRMFormatModifierPlaneLayoutEXT:
		return "invalid DRM format modifier plane layout"
	case ErrorFragmentationEXT:
		return "fragmentation"
	case ErrorNotPermittedEXT:
		return "not permitted"
	case ErrorInvalidDeviceAddressEXT:
		return "invalid device address"
	case ErrorFullScreenExclusiveModeLostEXT:
		return "full screen exclusive mode lost"
	default:
		panic(fmt.Sprintf("unknown result: %#v", r))
	}
}

type AllocationCallbacks C.VkAllocationCallbacks

// SystemAllocationScope - Allocation scope
//
// Most Vulkan commands operate on a single object, or there is a sole object that is being created or manipulated.
// When an allocation uses an allocation scope of SystemAllocationScopeObject or SystemAllocationScopeCache, the
// allocation is scoped to the object being created or manipulated.
//
// When an implementation requires host memory, it will make callbacks to the application using the most specific
// allocator and allocation scope available.
//
//   - If an allocation is scoped to the duration of a command, the allocator will use the SystemAllocationScopeCommand
//     allocation scope. The most specific allocator available is used: if the object being created or manipulated has
//     an allocator, that object's allocator will be used, else if the parent Device has an allocator it will be used,
//     else if the parent Instance has an allocator it will be used. Else,
//
//   - If an allocation is associated with a ValidationCacheEXT or PipelineCache object, the allocator will use the
//     SystemAllocationScopeCache allocation scope. The most specific allocator available is used (cache, else device,
//     else instance). Else,
//
//   - If an allocation is scoped to the lifetime of an object, that object is being created or manipulated by the
//     command, and that object's type is not Device or Instance, the allocator will use an allocation scope of
//     SystemAllocationScopeObject. The most specific allocator available is used (object, else device, else instance).
//     Else,
//
//   - If an allocation is scoped to the lifetime of a device, the allocator will use an allocation scope of
//     SystemAllocationScopeDevice. The most specific allocator available is used (device, else instance). Else,
//
//   - If the allocation is scoped to the lifetime of an instance and the instance has an allocator, its allocator will
//     be used with an allocation scope of SystemAllocationScopeInstance.
//
// - Otherwise an implementation will allocate memory through an alternative mechanism that is unspecified.
type SystemAllocationScope uint32

const (
	// SystemAllocationScopeCommand specifies that the allocation is scoped to the duration of the Vulkan
	// command.
	SystemAllocationScopeCommand SystemAllocationScope = iota

	// SystemAllocationScopeObject specifies that the allocation is scoped to the lifetime of the Vulkan object
	// that is being created or used.
	SystemAllocationScopeObject

	// SystemAllocationScopeCache specifies that the allocation is scoped to the lifetime of a PipelineCache
	// or ValidationCacheEXT object.
	SystemAllocationScopeCache

	// SystemAllocationScopeDevice specifies that the allocation is scoped to the lifetime of the Vulkan device.
	SystemAllocationScopeDevice

	// SystemAllocationScopeInstance specifies that the allocation is scoped to the lifetime of the Vulkan
	// instance.
	SystemAllocationScopeInstance
)

type InternalAllocationType uint32

const (
	InternalAllocationTypeExecutable InternalAllocationType = iota
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
	Type         StructureType
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

type structureHeader struct {
	Type StructureType
	Next unsafe.Pointer
}

func NextChainInterface[chainable interface {
	init(*chainable)
	alloc() (chainable, unsafe.Pointer)
	copy(chainable)
}](v chainable) *chainable {
	return &v
}

func chain[chainable interface {
	init(*chainable)
	alloc() (chainable, unsafe.Pointer)
	copy(chainable)
}](f func(), elems ...chainable) {
	chain2[chainable, func(chainable, *chainable), func(chainable) (chainable, unsafe.Pointer), func(chainable, chainable)](chainable.init, chainable.alloc, chainable.copy, f, elems...)
}

func chain2[
	chainable interface{},
	initFunc func(chainable, *chainable),
	allocFunc func(chainable) (chainable, unsafe.Pointer),
	copyFunc func(chainable, chainable),
](
	init initFunc,
	alloc allocFunc,
	copy copyFunc,
	f func(),
	elems ...chainable,
) {
	if len(elems) == 0 {
		f()
		return
	}

	for idx := range elems[1:] {
		defer init(elems[idx], &elems[idx+1])
	}
	var ip = elems[0]
	for _, e := range elems[1:] {
		iface, ptr := alloc(e)
		copy(iface, e) // todo: honor next chains
		defer C.free(ptr)
		defer copy(e, iface)
		init(ip, (*chainable)(ptr))
		ip = iface
	}
	init(ip, nil)

	f()
}

// copySliceToC creates a new copy of slice backed by C memory and sets *ptr
// to its address. It also returns the address, so it can be freed later.
func copySliceToC[T any](ptr **T, slice []T) unsafe.Pointer {
	var t T
	p := C.malloc(C.size_t(uintptr(len(slice)) * unsafe.Sizeof(t)))
	copy(unsafe.Slice((*T)(p), len(slice)), slice)
	*ptr = (*T)(p)
	return p
}
