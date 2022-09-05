package vulkan

// #cgo linux freebsd darwin LDFLAGS: -lvulkan
// #cgo windows LDFLAGS: -lvulkan-1
// #include <stdlib.h>
// #include <vulkan/vulkan.h>
// void (*_f)(VkPhysicalDevice device, VkPhysicalDeviceProperties2KHR *properties);
// void doInvoke(VkPhysicalDevice device, VkPhysicalDeviceProperties2KHR *properties) {
//   _f(device, properties);
// }
import "C"
import (
	"bytes"
	"fmt"
	"unsafe"
)

type CreateInfo struct {
	Type              StructureType
	Next              *CreateInfo
	Flags             C.VkInstanceCreateFlags
	ApplicationInfo   *ApplicationInfo
	EnabledLayers     []string
	EnabledExtensions []string
}

type createInfo struct {
	Type                  StructureType
	Next                  *createInfo
	Flags                 C.VkInstanceCreateFlags
	ApplicationInfo       *applicationInfo
	EnabledLayerCount     uint32
	EnabledLayerNames     *C.char
	EnabledExtensionCount uint32
	EnabledExtensionNames *C.char
}

type ApplicationInfo struct {
	Type               StructureType
	Next               uintptr
	ApplicationName    string
	ApplicationVersion uint32
	EngineName         string
	EngineVersion      uint32
	APIVersion         Version
}

func (info *ApplicationInfo) C(_info *applicationInfo) freeFunc {
	*_info = applicationInfo{
		Type:               info.Type,
		Next:               info.Next,
		ApplicationName:    nil,
		ApplicationVersion: info.ApplicationVersion,
		EngineName:         nil,
		EngineVersion:      info.EngineVersion,
		APIVersion:         info.APIVersion,
	}
	var application *C.char
	var engine *C.char
	application = C.CString(info.ApplicationName)
	engine = C.CString(info.EngineName)
	_info.ApplicationName = application
	_info.EngineName = engine
	return func() {
		C.free(unsafe.Pointer(engine))
		C.free(unsafe.Pointer(application))
	}
}

type applicationInfo struct {
	Type               StructureType
	Next               uintptr
	ApplicationName    *C.char
	ApplicationVersion uint32
	EngineName         *C.char
	EngineVersion      uint32
	APIVersion         Version
}

const ptrSize = 4 << (^uintptr(0) >> 63)

type freeFunc func()

func (f freeFunc) Free() {
	if f == nil {
		return
	}
	f()
}

func fillNames(slice []string, count *uint32, names **C.char) interface{ Free() } {
	if len(slice) == 0 {
		return freeFunc(nil)
	}

	p := C.malloc(C.size_t(len(slice)) * ptrSize)
	for i, name := range slice {
		*(**C.char)(unsafe.Add(p, i*ptrSize)) = C.CString(name)
	}
	*count = uint32(len(slice))
	*names = (*C.char)(p)
	return freeFunc(func() {
		for i := uint32(0); i < *count; i++ {
			C.free(unsafe.Pointer(*(**C.char)(unsafe.Add(p, i*ptrSize))))
		}
		C.free(p)
	})
}

type Layer struct {
	LayerName             [MaxExtensionNameSize]uint8
	SpecVersion           uint32
	ImplementationVersion uint32
	Description           [MaxDescriptionSize]uint8
}

func EnumerateInstanceVersion() (Version, error) {
	var version Version
	result := Result(C.vkEnumerateInstanceVersion(
		(*C.uint32_t)(&version),
	))
	if result != Success {
		return 0, result
	}
	return version, nil
}

func CreateInstance(info CreateInfo, allocator *AllocationCallbacks) (Instance, error) {
	var count uint32
	result := C.vkEnumerateInstanceLayerProperties((*C.uint32_t)(unsafe.Pointer(&count)), nil)
	if result != C.VK_SUCCESS {
		panic("enum")
	}
	fmt.Println(count, "layers")
	layers := make([]Layer, count)
	result = C.vkEnumerateInstanceLayerProperties((*C.uint32_t)(unsafe.Pointer(&count)), (*C.VkLayerProperties)(unsafe.Pointer(&layers[0])))
	if result != C.VK_SUCCESS {
		panic("enum2")
	}
	for _, layer := range layers {
		name := string(layer.LayerName[:])
		if off := bytes.IndexByte(layer.LayerName[:], 0); off != -1 {
			name = string(layer.LayerName[:off])
		}
		description := string(layer.Description[:])
		if off := bytes.IndexByte(layer.Description[:], 0); off != -1 {
			description = string(layer.Description[:off])
		}

		fmt.Println(name)
		fmt.Println(description)
		fmt.Println()
	}

	_appInfo := (*applicationInfo)(C.malloc(C.size_t(unsafe.Sizeof(applicationInfo{}))))
	defer C.free(unsafe.Pointer(_appInfo))
	defer info.ApplicationInfo.C(_appInfo).Free()
	var instance Instance
	_info := createInfo{
		Type: info.Type,
		//Next:                  nil, // todo
		Flags:                 info.Flags,
		ApplicationInfo:       _appInfo,
		EnabledLayerCount:     uint32(len(info.EnabledLayers)),
		EnabledLayerNames:     nil,
		EnabledExtensionCount: uint32(len(info.EnabledExtensions)),
		EnabledExtensionNames: nil,
	}
	defer fillNames(info.EnabledLayers, &_info.EnabledLayerCount, &_info.EnabledLayerNames).Free()
	defer fillNames(info.EnabledExtensions, &_info.EnabledExtensionCount, &_info.EnabledExtensionNames).Free()
	_info.Next = (*createInfo)(unsafe.Pointer(uintptr(0)))
	result = C.vkCreateInstance(
		(*C.VkInstanceCreateInfo)(unsafe.Pointer(&_info)),
		(*C.VkAllocationCallbacks)(allocator),
		(*C.VkInstance)(unsafe.Pointer(&instance)),
	)
	if result != C.VK_SUCCESS {
		return 0, Result(result)
	}
	return instance, nil
}

type Instance uintptr

func (i Instance) Destroy() {
	DestroyInstance(i)
}

func DestroyInstance(instance Instance) {
	C.vkDestroyInstance((C.VkInstance)(unsafe.Pointer(instance)), nil)
}

type Surface uintptr

func (i Instance) DestroySurface(surface Surface) {
	C.vkDestroySurfaceKHR((C.VkInstance)(unsafe.Pointer(i)), (C.VkSurfaceKHR)(unsafe.Pointer(surface)), nil)
}

type PhysicalDevice uintptr

type SwapchainCreateFlagBits uint32
type SwapchainCreateFlags = SwapchainCreateFlagBits

const (
	SwapchainCreateSplitInstanceBindRegionsBit SwapchainCreateFlagBits = 1 << iota
	SwapchainCreateProtectedBit
	SwapchainCreateMutableFormatBit
)

type CompositeAlphaFlagBits uint32
type CompositeAlphaFlags = CompositeAlphaFlagBits

const (
	CompositeAlphaOpaqueBit CompositeAlphaFlagBits = 1 << iota
	CompositeAlphaPreMultipliedBit
	CompositeAlphaPostMultipliedBit
	CompositeAlphaInheritBit
)

type SwapchainCreateInfo struct {
	Type                  StructureType
	Next                  uintptr
	Flags                 SwapchainCreateFlags
	Surface               Surface
	MinImageCount         uint32
	Format                Format
	ImageColorSpace       ColorSpace
	ImageExtent           Extent2D
	ImageArrayLayers      uint32
	ImageUsage            ImageUsageFlags
	ImageSharingMode      SharingMode
	QueueFamilyIndexCount uint32
	QueueFamilyIndices    *uint32
	PreTransform          C.VkSurfaceTransformFlagBitsKHR
	CompositeAlpha        CompositeAlphaFlags
	PresentMode           PresentMode
	Clipped               bool
	OldSwapchain          Swapchain
}

type Swapchain uintptr

type PresentInfo struct {
	Type           StructureType
	Next           uintptr
	WaitSemaphores []Semaphore
	Swapchains     []Swapchain
	ImageIndices   []uint32
	Results        []Result
}

func (info *PresentInfo) C(_info *presentInfo) freeFunc {
	*_info = presentInfo{
		Type:               info.Type,
		Next:               info.Next,
		WaitSemaphoreCount: uint32(len(info.WaitSemaphores)),
		WaitSemaphores:     nil,
		SwapchainCount:     uint32(len(info.Swapchains)),
		Swapchains:         nil,
		ImageIndices:       nil,
		Results:            nil,
	}
	var ps []unsafe.Pointer
	if _info.WaitSemaphoreCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.WaitSemaphoreCount) * unsafe.Sizeof(Semaphore(0))))
		ps = append(ps, p)
		for i, semaphore := range info.WaitSemaphores {
			*(*Semaphore)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(Semaphore(0)))) = semaphore
		}
		_info.WaitSemaphores = (*Semaphore)(p)
	}
	if _info.SwapchainCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.SwapchainCount) * unsafe.Sizeof(Swapchain(0))))
		ps = append(ps, p)
		for i, swapchain := range info.Swapchains {
			*(*Swapchain)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(Swapchain(0)))) = swapchain
		}
		_info.Swapchains = (*Swapchain)(p)
	}
	if _info.SwapchainCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.SwapchainCount) * unsafe.Sizeof(uint32(0))))
		ps = append(ps, p)
		for i, imageIndex := range info.ImageIndices {
			*(*uint32)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(uint32(0)))) = imageIndex
		}
		_info.ImageIndices = (*uint32)(p)
	}
	if info.Results != nil {
		p := C.calloc(C.size_t(uintptr(_info.SwapchainCount)*unsafe.Sizeof(Result(0))), 1)
		ps = append(ps, p)
		_info.Results = (*Result)(p)
	}
	return freeFunc(func() {
		for i := len(ps); i > 0; i-- {
			C.free(ps[i-1])
		}
	})
}

type presentInfo struct {
	Type               StructureType
	Next               uintptr
	WaitSemaphoreCount uint32
	WaitSemaphores     *Semaphore
	SwapchainCount     uint32
	Swapchains         *Swapchain
	ImageIndices       *uint32
	Results            *Result
}

type DeviceCreateFlags uint32

type PhysicalDeviceFeatures struct {
	RobustBufferAccess                      bool
	_                                       [3]byte
	FullDrawIndexUint32                     bool
	_                                       [3]byte
	ImageCubeArray                          bool
	_                                       [3]byte
	IndependentBlend                        bool
	_                                       [3]byte
	GeometryShader                          bool
	_                                       [3]byte
	TessellationShader                      bool
	_                                       [3]byte
	SampleRateShading                       bool
	_                                       [3]byte
	DualSrcBlend                            bool
	_                                       [3]byte
	LogicOp                                 bool
	_                                       [3]byte
	MultiDrawIndirect                       bool
	_                                       [3]byte
	DrawIndirectFirstInstance               bool
	_                                       [3]byte
	DepthClamp                              bool
	_                                       [3]byte
	DepthBiasClamp                          bool
	_                                       [3]byte
	FillModeNonSolid                        bool
	_                                       [3]byte
	DepthBounds                             bool
	_                                       [3]byte
	WideLines                               bool
	_                                       [3]byte
	LargePoints                             bool
	_                                       [3]byte
	AlphaToOne                              bool
	_                                       [3]byte
	MultiViewport                           bool
	_                                       [3]byte
	SamplerAnisotropy                       bool
	_                                       [3]byte
	TextureCompressionETC2                  bool
	_                                       [3]byte
	TextureCompressionASTCLDR               bool
	_                                       [3]byte
	TextureCompressionBC                    bool
	_                                       [3]byte
	OcclusionQueryPrecise                   bool
	_                                       [3]byte
	PipelineStatisticsQuery                 bool
	_                                       [3]byte
	VertexPipelineStoresAndAtomics          bool
	_                                       [3]byte
	FragmentStoresAndAtomics                bool
	_                                       [3]byte
	ShaderTessellationAndGeometryPointSize  bool
	_                                       [3]byte
	ShaderImageGatherExtended               bool
	_                                       [3]byte
	ShaderStorageImageExtendedFormats       bool
	_                                       [3]byte
	ShaderStorageImageMultisample           bool
	_                                       [3]byte
	ShaderStorageImageReadWithoutFormat     bool
	_                                       [3]byte
	ShaderStorageImageWriteWithoutFormat    bool
	_                                       [3]byte
	ShaderUniformBufferArrayDynamicIndexing bool
	_                                       [3]byte
	ShaderSampledImageArrayDynamicIndexing  bool
	_                                       [3]byte
	ShaderStorageBufferArrayDynamicIndexing bool
	_                                       [3]byte
	ShaderStorageImageArrayDynamicIndexing  bool
	_                                       [3]byte
	ShaderClipDistance                      bool
	_                                       [3]byte
	ShaderCullDistance                      bool
	_                                       [3]byte
	ShaderFloat64                           bool
	_                                       [3]byte
	ShaderInt64                             bool
	_                                       [3]byte
	ShaderInt16                             bool
	_                                       [3]byte
	ShaderResourceResidency                 bool
	_                                       [3]byte
	ShaderResourceMinLod                    bool
	_                                       [3]byte
	SparseBinding                           bool
	_                                       [3]byte
	SparseResidencyBuffer                   bool
	_                                       [3]byte
	SparseResidencyImage2D                  bool
	_                                       [3]byte
	SparseResidencyImage3D                  bool
	_                                       [3]byte
	SparseResidency2Samples                 bool
	_                                       [3]byte
	SparseResidency4Samples                 bool
	_                                       [3]byte
	SparseResidency8Samples                 bool
	_                                       [3]byte
	SparseResidency16Samples                bool
	_                                       [3]byte
	SparseResidencyAliased                  bool
	_                                       [3]byte
	VariableMultisampleRate                 bool
	_                                       [3]byte
	InheritedQueries                        bool
	_                                       [3]byte
}

func (f *PhysicalDeviceFeatures2) pdfiInit(i *PhysicalDeviceFeaturesInterface) {
	f.Type = StructureTypePhysicalDeviceFeatures2
	if i != nil {
		f.Next = i
	}
}

func (f *PhysicalDeviceFeatures2) pdfiAlloc() (PhysicalDeviceFeaturesInterface, unsafe.Pointer) {
	ptr := C.calloc(1, (C.size_t)(unsafe.Sizeof(*f)))
	return (*PhysicalDeviceFeatures2)(ptr), ptr
}

func (f *PhysicalDeviceFeatures2) pdfiCopy(i PhysicalDeviceFeaturesInterface) {
	*f = *(i.(*PhysicalDeviceFeatures2))
}

type PhysicalDeviceFeatures2 struct {
	Type StructureType
	Next *PhysicalDeviceFeaturesInterface
	PhysicalDeviceFeatures
}

type DeviceCreateInfo struct {
	Type              StructureType
	Next              *DeviceCreateInfoInterface
	Flags             DeviceCreateFlags
	QueueCreateInfos  []DeviceQueueCreateInfo
	EnabledLayers     []string
	EnabledExtensions []string
	EnabledFeatures   *PhysicalDeviceFeatures
}

type deviceCreateInfo struct {
	Type                  StructureType
	Next                  *DeviceCreateInfoInterface
	Flags                 DeviceCreateFlags
	QueueCreateInfoCount  uint32
	QueueCreateInfos      *deviceQueueCreateInfo
	EnabledLayerCount     uint32
	EnabledLayerNames     *C.char
	EnabledExtensionCount uint32
	EnabledExtensionNames *C.char
	EnabledFeatures       *PhysicalDeviceFeatures
}

func (info *deviceCreateInfo) dciiInit(i *DeviceCreateInfoInterface) {
	info.Type = StructureTypeDeviceCreateInfo
	if i != nil {
		info.Next = i
	}
}

func (info *deviceCreateInfo) dciiAlloc() (DeviceCreateInfoInterface, unsafe.Pointer) {
	ptr := C.calloc(1, (C.size_t)(unsafe.Sizeof(*info)))
	return (*deviceCreateInfo)(ptr), ptr
}

func (info *deviceCreateInfo) dciiCopy(i DeviceCreateInfoInterface) {
	*info = *(i.(*deviceCreateInfo))
}

type DeviceGroupDeviceCreateInfo struct {
	Type            StructureType
	Next            *DeviceCreateInfoInterface
	PhysicalDevices []PhysicalDevice

	free freeFunc
}

func (info *DeviceGroupDeviceCreateInfo) C(_info *deviceGroupDeviceCreateInfo) freeFunc {
	var physicalDevicePtr unsafe.Pointer
	if len(info.PhysicalDevices) > 0 {
		physicalDevicePtr = copySliceToC(nil, info.PhysicalDevices)
	}
	*_info = deviceGroupDeviceCreateInfo{
		Type:                info.Type,
		Next:                info.Next,
		PhysicalDeviceCount: uint32(len(info.PhysicalDevices)),
		PhysicalDevicePtr:   (*PhysicalDevice)(physicalDevicePtr),
	}
	f := func() {
		C.free(physicalDevicePtr)
	}
	info.free = f
	return f
}

func (info *DeviceGroupDeviceCreateInfo) dciiInit(i *DeviceCreateInfoInterface) {
	info.Type = StructureTypeDeviceGroupDeviceCreateInfo
	if i != nil {
		info.Next = i
	}
}

func (info *DeviceGroupDeviceCreateInfo) dciiAlloc() (DeviceCreateInfoInterface, unsafe.Pointer) {
	ptr := C.calloc(1, (C.size_t)(unsafe.Sizeof(deviceGroupDeviceCreateInfo{})))
	return (*deviceGroupDeviceCreateInfo)(ptr), ptr
}

func (info *DeviceGroupDeviceCreateInfo) dciiCopy(i DeviceCreateInfoInterface) {
	_info := i.(*deviceGroupDeviceCreateInfo)
	physicalDevices := make([]PhysicalDevice, _info.PhysicalDeviceCount)
	copy(physicalDevices, unsafe.Slice(_info.PhysicalDevicePtr, _info.PhysicalDeviceCount))
	defer info.free()
	*info = DeviceGroupDeviceCreateInfo{
		Type:            _info.Type,
		Next:            _info.Next,
		PhysicalDevices: physicalDevices,
	}
}

type deviceGroupDeviceCreateInfo struct {
	Type                StructureType
	Next                *DeviceCreateInfoInterface
	PhysicalDeviceCount uint32
	PhysicalDevicePtr   *PhysicalDevice
}

func (info *deviceGroupDeviceCreateInfo) dciiInit(i *DeviceCreateInfoInterface) {
	info.Type = StructureTypeDeviceGroupDeviceCreateInfo
	if i != nil {
		info.Next = i
	}
}

func (info *deviceGroupDeviceCreateInfo) dciiAlloc() (DeviceCreateInfoInterface, unsafe.Pointer) {
	ptr := C.calloc(1, (C.size_t)(unsafe.Sizeof(*info)))
	return (*deviceGroupDeviceCreateInfo)(ptr), ptr
}

func (info *deviceGroupDeviceCreateInfo) dciiCopy(i DeviceCreateInfoInterface) {
	i.(*DeviceGroupDeviceCreateInfo).C(info)
}

type DeviceQueueCreateInfo struct {
	Type             StructureType
	Next             uintptr
	Flags            DeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueuePriorities  []float32
}

type DeviceQueueCreateFlagBits uint32
type DeviceQueueCreateFlags = DeviceQueueCreateFlagBits

const (
	DeviceQueueCreateProtectedBit DeviceQueueCreateFlagBits = 1 << iota
)

type deviceQueueCreateInfo struct {
	Type             StructureType
	Next             uintptr
	Flags            DeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueueCount       uint32
	QueuePriorities  *float32
}

//type deviceQueueCreateInfo struct {
//	Type C.VkStructureType
//	Next uintptr
//	Flags C.VkDeviceQueueCreateFlags
//}

func (i Instance) EnumeratePhysicalDevices() ([]PhysicalDevice, error) {
	return EnumeratePhysicalDevices(i)
}

type PhysicalDeviceGroupProperties struct {
	Type             StructureType
	Next             uintptr
	PhysicalDevices  []PhysicalDevice
	SubsetAllocation bool
}

type physicalDeviceGroupProperties struct {
	Type                StructureType
	Next                uintptr
	PhysicalDeviceCount uint32
	PhysicalDevices     [MaxDeviceGroupSize]PhysicalDevice
	SubsetAllocation    bool
	_                   [3]byte
}

func EnumeratePhysicalDeviceGroups(instance Instance) ([]PhysicalDeviceGroupProperties, error) {
	//var count uint32
	count := uint32(0)
	result := Result(C.vkEnumeratePhysicalDeviceGroups(
		(C.VkInstance)(unsafe.Pointer(instance)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		nil,
	))
	if result != Success {
		return nil, result
	}
	_groups := make([]physicalDeviceGroupProperties, count)
	for i := range _groups {
		_groups[i].Type = StructureTypePhysicalDeviceGroupProperties
	}
	result = Result(C.vkEnumeratePhysicalDeviceGroups(
		(C.VkInstance)(unsafe.Pointer(instance)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		(*C.VkPhysicalDeviceGroupProperties)(unsafe.Pointer(&_groups[0])),
	))
	if result != Success {
		return nil, result
	}
	groups := make([]PhysicalDeviceGroupProperties, count)
	for i := range _groups {
		groups[i] = PhysicalDeviceGroupProperties{
			Type:             _groups[i].Type,
			Next:             _groups[i].Next,
			PhysicalDevices:  _groups[i].PhysicalDevices[:_groups[i].PhysicalDeviceCount:_groups[i].PhysicalDeviceCount],
			SubsetAllocation: _groups[i].SubsetAllocation,
		}
	}
	return groups, nil
}

func EnumeratePhysicalDevices(instance Instance) ([]PhysicalDevice, error) {
	var count C.uint32_t //asd768687
	fmt.Println(unsafe.Sizeof(PhysicalDeviceProperties2KHR{}))
	result := C.vkEnumeratePhysicalDevices(
		(C.VkInstance)(unsafe.Pointer(instance)),
		&count,
		nil,
	)
	if result != C.VK_SUCCESS {
		return nil, Result(result)
	}
	devices := make([]PhysicalDevice, count)
	result = C.vkEnumeratePhysicalDevices(
		(C.VkInstance)(unsafe.Pointer(instance)),
		&count,
		(*C.VkPhysicalDevice)(unsafe.Pointer(&devices[0])),
	)
	for _, device := range devices {
		properties := GetPhysicalDeviceProperties2(device)
		fmt.Println("- physical device found:")
		name := string(properties.DeviceName[:])
		if off := bytes.IndexByte(properties.DeviceName[:], 0); off != -1 {
			name = string(properties.DeviceName[:off])
		}
		uuid := fmt.Sprintf(
			"%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
			properties.PipelineCacheUUID[0],
			properties.PipelineCacheUUID[1],
			properties.PipelineCacheUUID[2],
			properties.PipelineCacheUUID[3],
			properties.PipelineCacheUUID[4],
			properties.PipelineCacheUUID[5],
			properties.PipelineCacheUUID[6],
			properties.PipelineCacheUUID[7],
			properties.PipelineCacheUUID[8],
			properties.PipelineCacheUUID[9],
			properties.PipelineCacheUUID[10],
			properties.PipelineCacheUUID[11],
			properties.PipelineCacheUUID[12],
			properties.PipelineCacheUUID[13],
			properties.PipelineCacheUUID[14],
			properties.PipelineCacheUUID[15],
		)
		fmt.Println("  name:", name)
		fmt.Println("  uuid:", uuid)
	}
	return devices, nil
}

func GetPhysicalDeviceProperties(physicalDevice PhysicalDevice) PhysicalDeviceProperties {
	var properties PhysicalDeviceProperties
	C.vkGetPhysicalDeviceProperties(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(*C.VkPhysicalDeviceProperties)(unsafe.Pointer(&properties)),
	)
	return properties
}

func GetPhysicalDeviceSurfaceSupport(physicalDevice PhysicalDevice, queueFamilyIndex uint32, surface Surface) (bool, error) {
	var supported struct {
		bool
		_ [3]byte
	}
	result := Result(C.vkGetPhysicalDeviceSurfaceSupportKHR(
		(C.VkPhysicalDevice)(unsafe.Pointer(physicalDevice)),
		(C.uint32_t)(queueFamilyIndex),
		(C.VkSurfaceKHR)(unsafe.Pointer(surface)),
		(*C.VkBool32)(unsafe.Pointer(&supported.bool)),
	))
	if result != Success {
		return false, result
	}
	return supported.bool, nil
}

func (d PhysicalDevice) GetSurfaceSupport(queueFamilyIndex uint32, surface Surface) (bool, error) {
	var supported uint32
	result := C.vkGetPhysicalDeviceSurfaceSupportKHR((C.VkPhysicalDevice)(unsafe.Pointer(d)), (C.uint32_t)(queueFamilyIndex), (C.VkSurfaceKHR)(unsafe.Pointer(surface)), (*C.VkBool32)(unsafe.Pointer(&supported)))
	if result != C.VK_SUCCESS {
		return false, fmt.Errorf("surface support error")
	}
	return supported == C.VK_TRUE, nil
}

func DestroySurface(instance Instance, surface Surface, allocator *AllocationCallbacks) {
	C.vkDestroySurfaceKHR(
		(C.VkInstance)(unsafe.Pointer(instance)),
		(C.VkSurfaceKHR)(unsafe.Pointer(surface)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}
