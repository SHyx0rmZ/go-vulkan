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
	Type StructureType
	Next *PhysicalDevicePropertiesInterface
	PhysicalDeviceProperties
}

type PhysicalDeviceProperties2 PhysicalDeviceProperties2KHR

type PhysicalDevicePropertiesInterface interface {
	init(*PhysicalDevicePropertiesInterface)
	alloc() (PhysicalDevicePropertiesInterface, unsafe.Pointer)
	copy(PhysicalDevicePropertiesInterface)
}

func (p *PhysicalDeviceProperties2) init(i *PhysicalDevicePropertiesInterface) {
	p.Type = StructureTypePhysicalDeviceProperties2
	if i != nil {
		p.Next = i
	}
}

func (p *PhysicalDeviceProperties2) alloc() (PhysicalDevicePropertiesInterface, unsafe.Pointer) {
	ptr := C.calloc(1, (C.size_t)(unsafe.Sizeof(*p)))
	return (*PhysicalDeviceProperties2)(ptr), ptr
}

func (p *PhysicalDeviceProperties2) copy(i PhysicalDevicePropertiesInterface) {
	*p = *(i.(*PhysicalDeviceProperties2))
}

type SubgroupFeatureFlags = SubgroupFeatureFlagBits
type SubgroupFeatureFlagBits uint32

const (
	SubgroupFeatureBasicBit SubgroupFeatureFlagBits = 1 << iota
	SubgroupFeatureVoteBit
	SubgroupFeatureArithmeticBit
	SubgroupFeatureBallotBit
	SubgroupFeatureShuffleBit
	SubgroupFeatureShuffleRelativeBit
	SubgroupFeatureClusteredBit
	SubgroupFeatureQuadBit
	SubgroupFeaturePartitionedBitNV
)

type PointClippingBehavior uint32

const (
	PointClippingBehaviorAllClipPanes PointClippingBehavior = iota
	PointClippingBehaviorUserClipPlanesOnly
)

type PhysicalDeviceVulkan11Properties struct {
	Type                              StructureType
	Next                              *PhysicalDevicePropertiesInterface
	DeviceUUID                        [UUIDSize]byte
	DriverUUID                        [UUIDSize]byte
	DeviceLUID                        [LUIDSize]byte
	DeviceNodeMask                    uint32
	DeviceLUIDValid                   bool
	_                                 [3]byte
	SubgroupSize                      uint32
	SubgroupSupportedStages           ShaderStageFlags
	SubgroupSupportedOperations       SubgroupFeatureFlags
	SubgroupQuadOperationsInAllStages bool
	_                                 [3]byte
	PointClippingBehavior             PointClippingBehavior
	MaxMultiviewViewCount             uint32
	MaxMultiviewInstanceIndex         uint32
	ProtectedNoFault                  bool
	_                                 [3]byte
	MaxPerSetDescriptors              uint32
	MaxMemoryAllocationSize           DeviceSize
}

func (p *PhysicalDeviceVulkan11Properties) init(i *PhysicalDevicePropertiesInterface) {
	p.Type = StructureTypePhysicalDeviceVulkan11Properties
	if i != nil {
		p.Next = i
	}
}

func (p *PhysicalDeviceVulkan11Properties) alloc() (PhysicalDevicePropertiesInterface, unsafe.Pointer) {
	ptr := C.calloc(1, (C.size_t)(unsafe.Sizeof(*p)))
	return (*PhysicalDeviceVulkan11Properties)(ptr), ptr
}

func (p *PhysicalDeviceVulkan11Properties) copy(i PhysicalDevicePropertiesInterface) {
	*p = *(i.(*PhysicalDeviceVulkan11Properties))
}

type DriverID uint32

const (
	DriverIDAMDProprietary DriverID = iota + 1
	DriverIDAMDOpenSource
	DriverIDMesaRADV
	DriverIDNVIDIAProprietary
	DriverIDIntelProprietaryWindows
	DriverIDIntelOpenSourceMesa
	DriverIDImaginationProprietary
	DriverIDQualcommProprietary
	DriverIDARMProprietary
	DriverIDGoogleSwiftShader
	DriverIDGGPProprietary
	DriverIDBroadcomProprietary
	DriverIDMesaLLVMPipe
	DriverIDMoltenVK
	DriverIDCoreAVIProprietary
	DriverIDJuiceProprietary
	DriverIDVeriSiliconProprietary
	DriverIDMesaTurnip
	DriverIDMesaV3DV
	DriverIDMesaPanVk
	DriverIDSamsungProprietary
	DriverIDMesaVenus
)

type ShaderFloatControlsIndependence uint32

const (
	ShaderFloatControlsIndependence32BitOnly ShaderFloatControlsIndependence = iota
	ShaderFloatControlsIndependenceAll
	ShaderFloatControlsIndependenceNone
)

type ResolveModeFlags = ResolveModeFlagBits
type ResolveModeFlagBits uint32

const (
	ResolveModeNone          ResolveModeFlagBits = 0
	ResolveModeSampleZeroBit ResolveModeFlagBits = 1 << (iota - 1)
	ResolveModeAverageBit
	ResolveModeMinBit
	ResolveModeMaxBit
)

type PhysicalDeviceVulkan12Properties struct {
	Type                                                 StructureType
	Next                                                 *PhysicalDevicePropertiesInterface
	DriverID                                             DriverID
	DriverName                                           [MaxDriverNameSize]byte
	DriverInfo                                           [MaxDriverInfoSize]byte
	ConformanceVersion                                   ConformanceVersion
	DenormBehaviorIndependence                           ShaderFloatControlsIndependence
	RoundingBehaviorIndependence                         ShaderFloatControlsIndependence
	ShaderSignedZeroInfNaNPreserveFloat16                bool
	_                                                    [3]byte
	ShaderSignedZeroInfNaNPreserveFloat32                bool
	_                                                    [3]byte
	ShaderSignedZeroInfNaNPreserveFloat64                bool
	_                                                    [3]byte
	ShaderDenormPreserveFloat16                          bool
	_                                                    [3]byte
	ShaderDenormPreserveFloat32                          bool
	_                                                    [3]byte
	ShaderDenormPreserveFloat64                          bool
	_                                                    [3]byte
	ShaderDenormFlushToZeroFloat16                       bool
	_                                                    [3]byte
	ShaderDenormFlushToZeroFloat32                       bool
	_                                                    [3]byte
	ShaderDenormFlushToZeroFloat64                       bool
	_                                                    [3]byte
	ShaderRoundingModeRTEFloat16                         bool
	_                                                    [3]byte
	ShaderRoundingModeRTEFloat32                         bool
	_                                                    [3]byte
	ShaderRoundingModeRTEFloat64                         bool
	_                                                    [3]byte
	ShaderRoundingModeRTZFloat16                         bool
	_                                                    [3]byte
	ShaderRoundingModeRTZFloat32                         bool
	_                                                    [3]byte
	ShaderRoundingModeRTZFloat64                         bool
	_                                                    [3]byte
	MaxUpdateAfterBindingDescriptorsInAllPools           uint32
	ShaderUniformBufferArrayNonUniformIndexingNative     bool
	_                                                    [3]byte
	ShaderSampledImageArrayNonUniformIndexingNative      bool
	_                                                    [3]byte
	ShaderStorageBufferArrayNonUniformIndexingNative     bool
	_                                                    [3]byte
	ShaderStorageImageArrayNonUniformIndexingNative      bool
	_                                                    [3]byte
	ShaderInputAttachmentArrayNonUniformIndexingNative   bool
	_                                                    [3]byte
	RobustBufferAccessUpdateAfterBind                    bool
	_                                                    [3]byte
	QuadDivergentImplicitLOD                             bool
	_                                                    [3]byte
	MaxPerStageDescriptorUpdateAfterBindSamplers         uint32
	MaxPerStageDescriptorUpdateAfterBindUniformBuffers   uint32
	MaxPerStageDescriptorUpdateAfterBindStorageBuffers   uint32
	MaxPerStageDescriptorUpdateAfterBindSampledImages    uint32
	MaxPerStageDescriptorUpdateAfterBindStorageImages    uint32
	MaxPerStageDescriptorUpdateAfterBindInputAttachments uint32
	MaxPerStageDescriptorUpdateAfterBindResources        uint32
	MaxDescriptorSetUpdateAfterBindSamplers              uint32
	MaxDescriptorSetUpdateAfterBindUniformBuffers        uint32
	MaxDescriptorSetUpdateAfterBindUniformBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindStorageBuffers        uint32
	MaxDescriptorSetUpdateAfterBindStorageBuffersDynamic uint32
	MaxDescriptorSetUpdateAfterBindSampledImages         uint32
	MaxDescriptorSetUpdateAfterBindStorageImages         uint32
	MaxDescriptorSetUpdateAfterBindInputAttachments      uint32
	SupportedDepthResolveModes                           ResolveModeFlags
	SupportedStencilResolveModes                         ResolveModeFlags
	IndependentResolveNone                               bool
	_                                                    [3]byte
	IndependentResolve                                   bool
	_                                                    [3]byte
	FilterMinmaxSingleComponentFormats                   bool
	_                                                    [3]byte
	FilterMimmaxImageComponentMapping                    bool
	_                                                    [3]byte
	MaxTimelineSemaphoreValueDifference                  uint64
	FramebufferIntegerColorSampleCounts                  SampleCountFlags
}

func (p *PhysicalDeviceVulkan12Properties) init(i *PhysicalDevicePropertiesInterface) {
	p.Type = StructureTypePhysicalDeviceVulkan12Properties
	if i != nil {
		p.Next = i
	}
}

func (p *PhysicalDeviceVulkan12Properties) alloc() (PhysicalDevicePropertiesInterface, unsafe.Pointer) {
	ptr := C.calloc(1, (C.size_t)(unsafe.Sizeof(*p)))
	return (*PhysicalDeviceVulkan12Properties)(ptr), ptr
}

func (p *PhysicalDeviceVulkan12Properties) copy(i PhysicalDevicePropertiesInterface) {
	*p = *(i.(*PhysicalDeviceVulkan12Properties))
}

type PhysicalDeviceVulkan13Properties struct {
	Type                                                                          StructureType
	Next                                                                          *PhysicalDevicePropertiesInterface
	MinSubgroupSize                                                               uint32
	MaxSubgroupSize                                                               uint32
	MaxComputeWorkgroupSubgroups                                                  uint32
	RequiredSubgroupSizeStages                                                    ShaderStageFlags
	MaxInlineUniformBlockSize                                                     uint32
	MaxPerStageDescriptorInlineUniformBlocks                                      uint32
	MaxPerStageDescriptorUpdateAfterBindInlineUniformBlocks                       uint32
	MaxDescriptorSetInlineUniformBlocks                                           uint32
	MaxDescriptorSetUpdateAfterBindInlineUniformBlocks                            uint32
	MaxInlineUniformTotalSize                                                     uint32
	IntegerDotProduct8BitUnsignedAccelerated                                      bool
	_                                                                             [3]byte
	IntegerDotProduct8BitSignedAccelerated                                        bool
	_                                                                             [3]byte
	IntegerDotProduct8BitMixedSignednessAccelerated                               bool
	_                                                                             [3]byte
	IntegerDotProduct4x8BitPackedUnsignedAccelerated                              bool
	_                                                                             [3]byte
	IntegerDotProduct4x8BitPackedSignedAccelerated                                bool
	_                                                                             [3]byte
	IntegerDotProduct4x8BitPackedMixedSignednessAccelerated                       bool
	_                                                                             [3]byte
	IntegerDotProduct16BitUnsignedAccelerated                                     bool
	_                                                                             [3]byte
	IntegerDotProduct16BitSignedAccelerated                                       bool
	_                                                                             [3]byte
	IntegerDotProduct16BitMixedSignednessAccelerated                              bool
	_                                                                             [3]byte
	IntegerDotProduct32BitUnsignedAccelerated                                     bool
	_                                                                             [3]byte
	IntegerDotProduct32BitSignedAccelerated                                       bool
	_                                                                             [3]byte
	IntegerDotProduct32BitMixedSignednessAccelerated                              bool
	_                                                                             [3]byte
	IntegerDotProduct64BitUnsignedAccelerated                                     bool
	_                                                                             [3]byte
	IntegerDotProduct64BitSignedAccelerated                                       bool
	_                                                                             [3]byte
	IntegerDotProduct64BitMixedSignednessAccelerated                              bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating8BitUnsignedAccelerated                bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating8BitSignedAccelerated                  bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating8BitMixedSignednessAccelerated         bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating4x8BitPackedUnsignedAccelerated        bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating4x8BitPackedSignedAccelerated          bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating4x8BitPackedMixedSignednessAccelerated bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating16BitUnsignedAccelerated               bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating16BitSignedAccelerated                 bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating16BitMixedSignednessAccelerated        bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating32BitUnsignedAccelerated               bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating32BitSignedAccelerated                 bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating32BitMixedSignednessAccelerated        bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating64BitUnsignedAccelerated               bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating64BitSignedAccelerated                 bool
	_                                                                             [3]byte
	IntegerDotProductAccumulatingSaturating64BitMixedSignednessAccelerated        bool
	_                                                                             [3]byte
	StorageTexelBufferOffsetAlignmentBytes                                        DeviceSize
	StorageTexelBufferOffsetSingleTexelAlignment                                  bool
	_                                                                             [3]byte
	UniformTexelBufferOffsetAlignmentBytes                                        DeviceSize
	UniformTexelBufferOffsetSingleTexelAlignment                                  bool
	_                                                                             [3]byte
	MaxBufferSize                                                                 DeviceSize
}

func (p *PhysicalDeviceVulkan13Properties) init(i *PhysicalDevicePropertiesInterface) {
	p.Type = StructureTypePhysicalDeviceVulkan13Properties
	if i != nil {
		p.Next = i
	}
}

func (p *PhysicalDeviceVulkan13Properties) alloc() (PhysicalDevicePropertiesInterface, unsafe.Pointer) {
	ptr := C.calloc(1, (C.size_t)(unsafe.Sizeof(*p)))
	return (*PhysicalDeviceVulkan13Properties)(ptr), ptr
}

func (p *PhysicalDeviceVulkan13Properties) copy(i PhysicalDevicePropertiesInterface) {
	*p = *(i.(*PhysicalDeviceVulkan13Properties))
}

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

func GetPhysicalDeviceProperties2(physicalDevice PhysicalDevice, next ...PhysicalDevicePropertiesInterface) PhysicalDeviceProperties2 {
	var properties PhysicalDeviceProperties2
	chain(func() {
		C.vkGetPhysicalDeviceProperties2(
			*(*C.VkPhysicalDevice)(unsafe.Pointer(&physicalDevice)),
			(*C.VkPhysicalDeviceProperties2)(unsafe.Pointer(&properties)),
		)
	}, append([]PhysicalDevicePropertiesInterface{&properties}, next...)...)
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

type PhysicalDeviceFeaturesInterface interface {
	init(*PhysicalDeviceFeaturesInterface)
	alloc() (PhysicalDeviceFeaturesInterface, unsafe.Pointer)
	copy(PhysicalDeviceFeaturesInterface)
}

type PhysicalDeviceVulkan11Features struct {
	Type                               StructureType
	Next                               *PhysicalDeviceFeaturesInterface
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

func (f *PhysicalDeviceVulkan11Features) init(i *PhysicalDeviceFeaturesInterface) {
	f.Type = StructureTypePhysicalDeviceVulkan11Features
	if i != nil {
		f.Next = i
	}
}

func (f *PhysicalDeviceVulkan11Features) alloc() (PhysicalDeviceFeaturesInterface, unsafe.Pointer) {
	ptr := C.calloc(1, (C.size_t)(unsafe.Sizeof(*f)))
	return (*PhysicalDeviceVulkan11Features)(ptr), ptr
}

func (f *PhysicalDeviceVulkan11Features) copy(i PhysicalDeviceFeaturesInterface) {
	*f = *(i.(*PhysicalDeviceVulkan11Features))
}

type PhysicalDeviceVulkan12Features struct {
	Type                                               StructureType
	Next                                               *PhysicalDeviceFeaturesInterface
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

func (f *PhysicalDeviceVulkan12Features) init(i *PhysicalDeviceFeaturesInterface) {
	f.Type = StructureTypePhysicalDeviceVulkan12Features
	if i != nil {
		f.Next = i
	}
}

func (f *PhysicalDeviceVulkan12Features) alloc() (PhysicalDeviceFeaturesInterface, unsafe.Pointer) {
	ptr := C.calloc(1, (C.size_t)(unsafe.Sizeof(*f)))
	return (*PhysicalDeviceVulkan12Features)(ptr), ptr
}

func (f *PhysicalDeviceVulkan12Features) copy(i PhysicalDeviceFeaturesInterface) {
	*f = *(i.(*PhysicalDeviceVulkan12Features))
}

type PhysicalDeviceVulkan13Features struct {
	Type                                               StructureType
	Next                                               *PhysicalDeviceFeaturesInterface
	RobustImageAccess                                  bool
	_                                                  [3]byte
	InlineUniformBlock                                 bool
	_                                                  [3]byte
	DescriptorBindingInlineUniformBlockUpdateAfterBind bool
	_                                                  [3]byte
	PipelineCreationCacheControl                       bool
	_                                                  [3]byte
	PrivateData                                        bool
	_                                                  [3]byte
	ShaderDemoteToHelperInvocation                     bool
	_                                                  [3]byte
	ShaderTerminateInvocation                          bool
	_                                                  [3]byte
	SubgroupSizeControl                                bool
	_                                                  [3]byte
	ComputeFullSubgroups                               bool
	_                                                  [3]byte
	Synchronization2                                   bool
	_                                                  [3]byte
	TextureCompressionASTCHDR                          bool
	_                                                  [3]byte
	ShaderZeroInitializeWorkgroupMemory                bool
	_                                                  [3]byte
	DynamicRendering                                   bool
	_                                                  [3]byte
	ShaderIntegerDotProduct                            bool
	_                                                  [3]byte
	Maintenance4                                       bool
	_                                                  [3]byte
}

func (f *PhysicalDeviceVulkan13Features) init(i *PhysicalDeviceFeaturesInterface) {
	f.Type = StructureTypePhysicalDeviceVulkan13Features
	if i != nil {
		f.Next = i
	}
}

func (f *PhysicalDeviceVulkan13Features) alloc() (PhysicalDeviceFeaturesInterface, unsafe.Pointer) {
	ptr := C.calloc(1, (C.size_t)(unsafe.Sizeof(*f)))
	return (*PhysicalDeviceVulkan13Features)(ptr), ptr
}

func (f *PhysicalDeviceVulkan13Features) copy(i PhysicalDeviceFeaturesInterface) {
	*f = *(i.(*PhysicalDeviceVulkan13Features))
}

func GetPhysicalDeviceFeatures(physicalDevice PhysicalDevice, next ...PhysicalDeviceFeaturesInterface) (PhysicalDeviceFeatures2, error) {
	var features PhysicalDeviceFeatures2
	chain(func() {
		C.vkGetPhysicalDeviceFeatures2(
			(C.VkPhysicalDevice)(*(*C.VkPhysicalDevice)(unsafe.Pointer(&physicalDevice))),
			(*C.VkPhysicalDeviceFeatures2)(unsafe.Pointer(&features)),
		)
	}, append([]PhysicalDeviceFeaturesInterface{&features}, next...)...)
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
