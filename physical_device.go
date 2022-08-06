package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"bytes"
	"unsafe"
)

type QueueFlagBit uint32
type QueueFlags = QueueFlagBit

const (
	QueueGraphicsBit QueueFlagBit = 1 << iota
	QueueComputeBit
	QueueTransferBit
	QueueSparseBindingBit
	QueueProtectedBit
)

type QueueFamilyProperties struct {
	QueueFlags                  QueueFlags
	QueueCount                  uint32
	TimestampValidBits          uint32
	MinImageTransferGranularity Extent3D
}

type QueueFamilyProperties2 struct {
	Type StructureType
	Next uintptr
	QueueFamilyProperties
}

func GetPhysicalDeviceQueueFamilyProperties(physicalDevice PhysicalDevice) []QueueFamilyProperties {
	var count uint32
	C.vkGetPhysicalDeviceQueueFamilyProperties(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		nil,
	)
	queueFamilyProperties := make([]QueueFamilyProperties, count)
	C.vkGetPhysicalDeviceQueueFamilyProperties(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		(*C.VkQueueFamilyProperties)(unsafe.Pointer(&queueFamilyProperties[0])),
	)
	return queueFamilyProperties[:count:count]
}

func GetPhysicalDeviceQueueFamilyProperties2(physicalDevice PhysicalDevice) []QueueFamilyProperties2 {
	var count uint32
	C.vkGetPhysicalDeviceQueueFamilyProperties2(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		nil,
	)
	queueFamilyProperties := make([]QueueFamilyProperties2, count)
	for i := range queueFamilyProperties {
		queueFamilyProperties[i].Type = StructureTypeQueueFamilyProperties2
	}
	C.vkGetPhysicalDeviceQueueFamilyProperties2(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		(*C.VkQueueFamilyProperties2)(unsafe.Pointer(&queueFamilyProperties[0])),
	)
	return queueFamilyProperties[:count:count]
}

type LayerName [MaxExtensionNameSize]byte
type Description [MaxDescriptionSize]byte
type ExtensionName [MaxExtensionNameSize]byte

func (x LayerName) String() string {
	return string(x[:bytes.IndexByte(x[:], 0)])
}

func (x Description) String() string {
	return string(x[:bytes.IndexByte(x[:], 0)])
}

func (x ExtensionName) String() string {
	return string(x[:bytes.IndexByte(x[:], 0)])
}

type LayerProperties struct {
	LayerName             LayerName
	SpecVersion           Version
	ImplementationVersion Version
	Description           Description
}

func EnumerateInstanceLayerProperties() ([]LayerProperties, error) {
	var count uint32
	result := Result(C.vkEnumerateInstanceLayerProperties(
		(*C.uint32_t)(unsafe.Pointer(&count)),
		nil,
	))
	if result != Success {
		return nil, result
	}
	if count == 0 {
		return nil, nil
	}
	properties := make([]LayerProperties, count)
	result = Result(C.vkEnumerateInstanceLayerProperties(
		(*C.uint32_t)(unsafe.Pointer(&count)),
		(*C.VkLayerProperties)(unsafe.Pointer(&properties[0])),
	))
	if result != Success {
		return nil, result
	}
	return properties[:count:count], nil
}

func EnumerateDeviceLayerProperties(physicalDevice PhysicalDevice) ([]LayerProperties, error) {
	var count uint32
	result := Result(C.vkEnumerateDeviceLayerProperties(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		nil,
	))
	if result != Success {
		return nil, result
	}
	if count == 0 {
		return nil, nil
	}
	properties := make([]LayerProperties, count)
	result = Result(C.vkEnumerateDeviceLayerProperties(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		(*C.VkLayerProperties)(unsafe.Pointer(&properties[0])),
	))
	if result != Success {
		return nil, result
	}
	return properties[:count:count], nil
}

type ExtensionProperties struct {
	ExtensionName ExtensionName
	SpecVersion   Version
}

func EnumerateInstanceExtensionProperties(layerName string) ([]ExtensionProperties, error) {
	var _layerName *C.char
	if layerName != "" {
		_layerName = C.CString(layerName)
		defer C.free(unsafe.Pointer(_layerName))
	}
	var count uint32
	result := Result(C.vkEnumerateInstanceExtensionProperties(
		_layerName,
		(*C.uint32_t)(unsafe.Pointer(&count)),
		nil,
	))
	if result != Success {
		return nil, result
	}
	if count == 0 {
		return nil, nil
	}
	properties := make([]ExtensionProperties, count)
	result = Result(C.vkEnumerateInstanceExtensionProperties(
		_layerName,
		(*C.uint32_t)(unsafe.Pointer(&count)),
		(*C.VkExtensionProperties)(unsafe.Pointer(&properties[0])),
	))
	if result != Success {
		return nil, result
	}
	return properties[:count:count], nil
}

func EnumerateDeviceExtensionProperties(physicalDevice PhysicalDevice, layerName string) ([]ExtensionProperties, error) {
	var _layerName *C.char
	if layerName != "" {
		_layerName = C.CString(layerName)
		defer C.free(unsafe.Pointer(_layerName))
	}
	var count uint32
	result := Result(C.vkEnumerateDeviceExtensionProperties(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		_layerName,
		(*C.uint32_t)(unsafe.Pointer(&count)),
		nil,
	))
	if result != Success {
		return nil, result
	}
	if count == 0 {
		return nil, nil
	}
	properties := make([]ExtensionProperties, count)
	result = Result(C.vkEnumerateDeviceExtensionProperties(
		*(*C.VkPhysicalDevice)(unsafe.Pointer(&physicalDevice)),
		_layerName,
		(*C.uint32_t)(unsafe.Pointer(&count)),
		(*C.VkExtensionProperties)(unsafe.Pointer(&properties[0])),
	))
	if result != Success {
		return nil, result
	}
	return properties[:count:count], nil
}

func CreateDevice(physicalDevice PhysicalDevice, info DeviceCreateInfo, allocator *AllocationCallbacks) (Device, error) {
	var device Device
	_info := deviceCreateInfo{
		Type:                  info.Type,
		Next:                  info.Next,
		Flags:                 info.Flags,
		QueueCreateInfoCount:  uint32(len(info.QueueCreateInfos)),
		EnabledLayerCount:     uint32(len(info.EnabledLayers)),
		EnabledExtensionCount: uint32(len(info.EnabledExtensions)),
	}
	if _info.QueueCreateInfoCount > 0 {
		const sizeOfDeviceQueueCreateInfo = unsafe.Sizeof(deviceQueueCreateInfo{})
		const sizeOfFloat32 = unsafe.Sizeof(float32(0))
		var l uintptr
		for _, info := range info.QueueCreateInfos {
			l += uintptr(len(info.QueuePriorities)) * sizeOfFloat32
		}
		p := C.malloc(C.size_t(len(info.QueueCreateInfos))*C.size_t(sizeOfDeviceQueueCreateInfo) + C.size_t(l))
		var o uintptr
		for _, info := range info.QueueCreateInfos {
			*(*deviceQueueCreateInfo)(unsafe.Add(p, o)) = deviceQueueCreateInfo{
				Type:             info.Type,
				Next:             info.Next,
				Flags:            info.Flags,
				QueueFamilyIndex: info.QueueFamilyIndex,
				QueueCount:       uint32(len(info.QueuePriorities)),
				QueuePriorities:  (*float32)(unsafe.Add(p, o+sizeOfDeviceQueueCreateInfo)),
			}
			o += sizeOfDeviceQueueCreateInfo
			for _, priority := range info.QueuePriorities {
				*(*float32)(unsafe.Add(p, o)) = priority
				o += sizeOfFloat32
			}
		}
		_info.QueueCreateInfos = (*deviceQueueCreateInfo)(p)
		defer func() {
			C.free(unsafe.Pointer(_info.QueueCreateInfos))
		}()
	}
	if info.EnabledFeatures != nil {
		p := C.malloc(C.size_t(unsafe.Sizeof(PhysicalDeviceFeatures{})))
		defer C.free(p)
		_info.EnabledFeatures = (*PhysicalDeviceFeatures)(p)
		*_info.EnabledFeatures = *info.EnabledFeatures
	}
	defer fillNames(info.EnabledLayers, &_info.EnabledLayerCount, &_info.EnabledLayerNames).Free()
	defer fillNames(info.EnabledExtensions, &_info.EnabledExtensionCount, &_info.EnabledExtensionNames).Free()
	result := Result(C.vkCreateDevice(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.VkDeviceCreateInfo)(unsafe.Pointer(&_info)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkDevice)(unsafe.Pointer(&device)),
	))
	if result != Success {
		return NullHandle, result
	}
	return device, nil
}

type PhysicalDeviceProperties2KHR struct {
	Type C.VkStructureType
	Next *PhysicalDeviceProperties2KHR
	PhysicalDeviceProperties
}

type PhysicalDeviceProperties2 PhysicalDeviceProperties2KHR

type PhysicalDeviceName [MaxPhysicalDeviceNameSize]byte

func (x PhysicalDeviceName) String() string {
	return string(x[:bytes.IndexByte(x[:], 0)])
}

type UUID [UUIDSize]byte

const hextable = "0123456789abcdef"

func (x UUID) String() string {
	return string([]byte{
		hextable[x[0]>>4],
		hextable[x[0]&15],
		hextable[x[1]>>4],
		hextable[x[1]&15],
		hextable[x[2]>>4],
		hextable[x[2]&15],
		hextable[x[3]>>4],
		hextable[x[3]&15],
		'-',
		hextable[x[4]>>4],
		hextable[x[4]&15],
		hextable[x[5]>>4],
		hextable[x[5]&15],
		'-',
		hextable[x[6]>>4],
		hextable[x[6]&15],
		hextable[x[7]>>4],
		hextable[x[7]&15],
		'-',
		hextable[x[8]>>4],
		hextable[x[8]&15],
		hextable[x[9]>>4],
		hextable[x[9]&15],
		'-',
		hextable[x[10]>>4],
		hextable[x[10]&15],
		hextable[x[11]>>4],
		hextable[x[11]&15],
		hextable[x[12]>>4],
		hextable[x[12]&15],
		hextable[x[13]>>4],
		hextable[x[13]&15],
		hextable[x[14]>>4],
		hextable[x[14]&15],
		hextable[x[15]>>4],
		hextable[x[15]&15],
	})
}

type PhysicalDeviceProperties struct {
	APIVersion        uint32
	DriverVersion     uint32
	VendorID          uint32
	DeviceID          uint32
	DeviceType        C.VkPhysicalDeviceType
	DeviceName        PhysicalDeviceName
	PipelineCacheUUID UUID
	Limits            C.VkPhysicalDeviceLimits
	SparseProperties  C.VkPhysicalDeviceSparseProperties
}

func GetPhysicalDeviceProperties2(physicalDevice PhysicalDevice) PhysicalDeviceProperties2 {
	properties := PhysicalDeviceProperties2{
		Type: (C.VkStructureType)(StructureTypePhysicalDeviceProperties2),
	}
	C.vkGetPhysicalDeviceProperties2(
		*(*C.VkPhysicalDevice)(unsafe.Pointer(&physicalDevice)),
		(*C.VkPhysicalDeviceProperties2)(unsafe.Pointer(&properties)),
	)
	return properties
}

type SurfaceFormat struct {
	Format     Format
	ColorSpace ColorSpace
}

func GetPhysicalDeviceSurfaceFormats(physicalDevice PhysicalDevice, surface Surface) ([]SurfaceFormat, error) {
	var count uint32
	result := Result(C.vkGetPhysicalDeviceSurfaceFormatsKHR(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(C.VkSurfaceKHR)(unsafe.Pointer(surface)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		nil,
	))
	if result != Success {
		return nil, result
	}
	formats := make([]SurfaceFormat, count)
	result = Result(C.vkGetPhysicalDeviceSurfaceFormatsKHR(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(C.VkSurfaceKHR)(unsafe.Pointer(surface)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		(*C.VkSurfaceFormatKHR)(unsafe.Pointer(&formats[0])),
	))
	if result != Success {
		return nil, result
	}
	return formats[:count:count], nil
}

//go:generate go run stringer.go -type PresentMode -output present_mode_string.go
type PresentMode uint32

const (
	PresentModeImmediate PresentMode = iota
	PresentModeMailbox
	PresentModeFIFO
	PresentModeFIFORelaxed
)

const (
	PresentModeSharedDemandRefresh     PresentMode = 1000111000
	PresentModeSharedContinuousRefresh PresentMode = 1000111001
)

func GetPhysicalDeviceSurfacePresentModes(physicalDevice PhysicalDevice, surface Surface) ([]PresentMode, error) {
	var count uint32
	result := Result(C.vkGetPhysicalDeviceSurfacePresentModesKHR(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(C.VkSurfaceKHR)(unsafe.Pointer(surface)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		nil,
	))
	if result != Success {
		return nil, result
	}
	modes := make([]PresentMode, count)
	result = Result(C.vkGetPhysicalDeviceSurfacePresentModesKHR(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(C.VkSurfaceKHR)(unsafe.Pointer(surface)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		(*C.VkPresentModeKHR)(unsafe.Pointer(&modes[0])),
	))
	if result != Success {
		return nil, result
	}
	return modes[:count:count], nil
}

type SurfaceTransformFlags = SurfaceTransformFlagBits

type SurfaceTransformFlagBits uint32

const (
	SurfaceTransformIdentityBit SurfaceTransformFlagBits = 1 << iota
	SurfaceTransformRotate90Bit
	SurfaceTransformRotate180Bit
	SurfaceTransformRotate270Bit
	SurfaceTransformHorizontalMirrorBit
	SurfaceTransformHorizontalMirrorRotate90Bit
	SurfaceTransformHorizontalMirrorRotate180BIt
	SurfaceTransformHorizontalMirrorRotate270Bit
	SurfaceTransformInheritBit
)

type SurfaceCapabilities struct {
	MinImageCount           uint32
	MaxImageCount           uint32
	CurrentExtent           Extent2D
	MinImageExtent          Extent2D
	MaxImageExtent          Extent2D
	MaxImageArrayLayers     uint32
	SupportedTransforms     SurfaceTransformFlags
	CurrentTransform        SurfaceTransformFlagBits
	SupportedCompositeAlpha CompositeAlphaFlags
	SupportedUsageFlags     ImageUsageFlags
}

func GetPhysicalDeviceSurfaceCapabilities(physicalDevice PhysicalDevice, surface Surface) (SurfaceCapabilities, error) {
	var capabilities SurfaceCapabilities
	result := Result(C.vkGetPhysicalDeviceSurfaceCapabilitiesKHR(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(C.VkSurfaceKHR)(unsafe.Pointer(surface)),
		(*C.VkSurfaceCapabilitiesKHR)(unsafe.Pointer(&capabilities)),
	))
	if result != Success {
		return SurfaceCapabilities{}, result
	}
	return capabilities, nil
}

type PhysicalDeviceVulkan11Features struct {
	Type                               StructureType
	Next                               uintptr
	StorageBuffer16BitAccess           bool
	_                                  [3]byte
	UniformAndStorageBuffer16BitAccess bool
	_                                  [3]byte
	StoragePushConstant16              bool
	_                                  [3]byte
	StorageInputOutput16               bool
	_                                  [3]byte
	Multiview                          bool
	_                                  [3]byte
	MultiviewGeometryShader            bool
	_                                  [3]byte
	MultiviewTessellationShader        bool
	_                                  [3]byte
	VariablePointersStorageBuffer      bool
	_                                  [3]byte
	VariablePointers                   bool
	_                                  [3]byte
	ProtectedMemory                    bool
	_                                  [3]byte
	SamplerYcbcrConversion             bool
	_                                  [3]byte
	ShaderDrawParameters               bool
	_                                  [3]byte
}

type PhysicalDeviceVulkan12Features struct {
	Type                                               StructureType
	Next                                               uintptr
	SamplerMirrorClampToEdge                           bool
	_                                                  [3]byte
	DrawIndirectCount                                  bool
	_                                                  [3]byte
	StorageBuffer8BitAccess                            bool
	_                                                  [3]byte
	UniformAndStorageBuffer8BitAccess                  bool
	_                                                  [3]byte
	StoragePushConstant8                               bool
	_                                                  [3]byte
	ShaderBufferInt64Atomics                           bool
	_                                                  [3]byte
	ShaderSharedInt64Atomics                           bool
	_                                                  [3]byte
	ShaderFloat16                                      bool
	_                                                  [3]byte
	ShaderInt8                                         bool
	_                                                  [3]byte
	DescriptorIndexing                                 bool
	_                                                  [3]byte
	ShaderInputAttachmentArrayDynamicIndexing          bool
	_                                                  [3]byte
	ShaderUniformTexelBufferArrayDynamicIndexing       bool
	_                                                  [3]byte
	ShaderStorageTexelBufferArrayDynamicIndexing       bool
	_                                                  [3]byte
	ShaderUniformBufferArrayNonUniformIndexing         bool
	_                                                  [3]byte
	ShaderSampledImagedArrayNonUniformIndexing         bool
	_                                                  [3]byte
	ShaderStorageBufferArrayNonUniformIndexing         bool
	_                                                  [3]byte
	ShaderStorageImageArrayNonUniformIndexing          bool
	_                                                  [3]byte
	ShaderInputAttachmentArrayNonUniformIndexing       bool
	_                                                  [3]byte
	ShaderUniformTexelBufferArrayNonUniformIndexing    bool
	_                                                  [3]byte
	ShaderStorageTexelBufferArrayNonUniformIndexing    bool
	_                                                  [3]byte
	DescriptorBindingUniformBufferUpdateAfterBind      bool
	_                                                  [3]byte
	DescriptorBindingSampledImageUpdateAfterBind       bool
	_                                                  [3]byte
	DescriptorBindingStorageImageUpdateAfterBind       bool
	_                                                  [3]byte
	DescriptorBindingStorageBufferUpdateAfterBind      bool
	_                                                  [3]byte
	DescriptorBindingUniformTexelBufferUpdateAfterBind bool
	_                                                  [3]byte
	DescriptorBindingStorageTexelBufferUpdateAfterBind bool
	_                                                  [3]byte
	DescriptorBindingUpdateUnusedWhilePending          bool
	_                                                  [3]byte
	DescriptorBindingPartiallyBound                    bool
	_                                                  [3]byte
	DescriptorBindingVariableDescriptorCount           bool
	_                                                  [3]byte
	RuntimeDescriptorArray                             bool
	_                                                  [3]byte
	SamplerFilterMinmax                                bool
	_                                                  [3]byte
	ScalarBlockLayout                                  bool
	_                                                  [3]byte
	ImagelessFramebuffer                               bool
	_                                                  [3]byte
	UniformBufferStandardLayout                        bool
	_                                                  [3]byte
	ShaderSubgroupExtendedTypes                        bool
	_                                                  [3]byte
	SeparateDepthStencilLayouts                        bool
	_                                                  [3]byte
	HostQueryReset                                     bool
	_                                                  [3]byte
	TimelineSemaphore                                  bool
	_                                                  [3]byte
	BufferDeviceAddress                                bool
	_                                                  [3]byte
	BufferDeviceAddressCaptureReplay                   bool
	_                                                  [3]byte
	BufferDeviceAddressMultiDevice                     bool
	_                                                  [3]byte
	VulkanMemoryModel                                  bool
	_                                                  [3]byte
	VulkanMemoryModelDeviceScope                       bool
	_                                                  [3]byte
	VulkanMemoryModelAvailabilityVisibilityChains      bool
	_                                                  [3]byte
	ShaderOutputViewportIndex                          bool
	_                                                  [3]byte
	ShaderOutputLayer                                  bool
	_                                                  [3]byte
	SubgroupBroadcastDynamicID                         bool
	_                                                  [3]byte
}

func GetPhysicalDeviceFeatures(physicalDevice PhysicalDevice) (PhysicalDeviceFeatures2, error) {
	var features PhysicalDeviceFeatures2
	var _features11 = (*PhysicalDeviceVulkan11Features)(C.calloc(1, (C.size_t)(unsafe.Sizeof(PhysicalDeviceVulkan11Features{}))))
	var _features12 = (*PhysicalDeviceVulkan12Features)(C.calloc(1, (C.size_t)(unsafe.Sizeof(PhysicalDeviceVulkan12Features{}))))
	defer C.free(unsafe.Pointer(_features12))
	defer C.free(unsafe.Pointer(_features11))
	features.Type = 1000059000
	features.Next = uintptr(unsafe.Pointer(_features11))
	_features11.Type = 49
	_features11.Next = uintptr(unsafe.Pointer(_features12))
	_features12.Type = C.VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_1_2_FEATURES // 51

	C.vkGetPhysicalDeviceFeatures2(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceFeatures2)(unsafe.Pointer(&features)),
	)
	features12 := &PhysicalDeviceVulkan12Features{}
	*features12 = *_features12
	features11 := &PhysicalDeviceVulkan11Features{}
	*features11 = *_features11
	features11.Next = uintptr(unsafe.Pointer(features12))
	features.Next = uintptr(unsafe.Pointer(features11))
	return features, nil
}

type FormatProperties struct {
	LinearTilingFeatures  FormatFeatureFlags
	OptimalTilingFeatures FormatFeatureFlags
	BufferFeatures        FormatFeatureFlags
}

type FormatFeatureFlags = FormatFeatureFlagBits

type FormatFeatureFlagBits uint32

const (
	FormatFeatureSampledImageBit FormatFeatureFlagBits = 1 << iota
	FormatFeatureStorageImageBit
	FormatFeatureStorageImageAtomicBit
	FormatFeatureUniformTexelBufferBit
	FormatFeatureStorageTexelBufferBit
	FormatFeatureStorageTexelBufferAtomicBit
	FormatFeatureVertexBufferBit
	FormatFeatureColorAttachmentBit
	FormatFeatureColorAttachmentBlendBit
	FormatFeatureDepthStencilAttachmentBit
	FormatFeatureBlitSrcBit
	FormatFeatureBlitDstBit
	FormatFeatureSampledImageFilterLinearBit
	FormatFeatureSampledImageFilterCubicBitImg
	FormatFeatureTransferSrcBit
	FormatFeatureTransferDstBit
	FormatFeatureSampledImageFilterMinMaxBit
	FormatFeatureMidpointChromaSamplesBit
	FormatFeatureSampledImageYCBCRConversionLinearFilterBit
	FormatFeatureSampledImageYCBCRConversionSeparateReconstructionFilterBit
	FormatFeatureSampledImageYCBCRConversionChromeReconstructionExplicitBit
	FormatFeatureSampledImageYCBCRConversionChromaReconstructionExplicitForceableBit
	FormatFeatureDisjointBit
	FormatFeatureCositedChromaSamplesBit
)

const (
	FormatFeatureFragmentDensityMapBitExt FormatFeatureFlagBits = 1 << (iota + 24)
	FormatFeatureVideoDecodeOutputBitKhr
	FormatFeatureVideoDecodeDPBBitKhr
	FormatFeatureVideoEncodeInputBitKhr
	FormatFeatureVideoEncodeDPBBitKhr
	FormatFeatureAccelerationStructureVertexBufferBitKhr
	FormatFeatureFragmentShadingRateAttachmentBitKhr
)

func GetPhysicalDeviceFormatProperties(physicalDevice PhysicalDevice, format Format) (properties FormatProperties) {
	C.vkGetPhysicalDeviceFormatProperties(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(C.VkFormat)(format),
		(*C.VkFormatProperties)(unsafe.Pointer(&properties)),
	)
	return properties
}
