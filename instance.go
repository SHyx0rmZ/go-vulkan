package vulkan

import "C"
import (
	"bytes"
	"fmt"
	"unsafe"
)

// #cgo linux freebsd darwin LDFLAGS: -lvulkan
// #include <stdlib.h>
// #define VK_USE_PLATFORM_XLIB_KHR 1
// #define VK_KHR_SURFACE 1
// #define VK_KHR_SWAPCHAIN 1
// #include <vulkan/vulkan.h>
// void (*_f)(VkPhysicalDevice device, VkPhysicalDeviceProperties2KHR *properties);
// void doInvoke(VkPhysicalDevice device, VkPhysicalDeviceProperties2KHR *properties) {
//   _f(device, properties);
// }
// VkResult (*_ptr_vkCreateXlibSurfaceKHR)(VkInstance instance, const VkXlibSurfaceCreateInfoKHR *info, const VkAllocationCallbacks *allocator, VkSurfaceKHR *surface);
// VkResult _vkCreateXlibSurfaceKHR(VkInstance instance, const VkXlibSurfaceCreateInfoKHR *info, const VkAllocationCallbacks *allocator, VkSurfaceKHR *surface) {
//   return _ptr_vkCreateXlibSurfaceKHR(instance, info, allocator, surface);
// }
import "C"

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
	ApplicationInfo       *ApplicationInfo
	EnabledLayerCount     uint32
	EnabledLayerNames     *C.char
	EnabledExtensionCount uint32
	EnabledExtensionNames *C.char
}

type ApplicationInfo struct {
	Type               C.VkStructureType
	Next               *ApplicationInfo
	ApplicationName    *[]uint8
	ApplicationVersion uint32
	EngineName         *[]uint8
	EngineVersion      uint32
	APIVersion         uint32
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
		*(**C.char)(unsafe.Pointer(uintptr(p) + uintptr(i*ptrSize))) = C.CString(name)
	}
	*count = uint32(len(slice))
	*names = (*C.char)(p)
	return freeFunc(func() {
		for i := uint32(0); i < *count; i++ {
			C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(uintptr(p) + uintptr(i*ptrSize)))))
		}
		C.free(p)
	})
}

const MaxExtensionNameSize int = C.VK_MAX_EXTENSION_NAME_SIZE
const MaxDescriptionSize int = C.VK_MAX_DESCRIPTION_SIZE

type Layer struct {
	LayerName             [MaxExtensionNameSize]uint8
	SpecVersion           uint32
	ImplementationVersion uint32
	Description           [MaxDescriptionSize]uint8
}

func CreateInstance(info CreateInfo) (Instance, error) {
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

	var instance Instance
	_info := createInfo{
		Type: info.Type,
		//Next:                  nil, // todo
		Flags:                 info.Flags,
		ApplicationInfo:       info.ApplicationInfo,
		EnabledLayerCount:     uint32(len(info.EnabledLayers)),
		EnabledLayerNames:     nil,
		EnabledExtensionCount: uint32(len(info.EnabledExtensions)),
		EnabledExtensionNames: nil,
	}
	defer fillNames(info.EnabledLayers, &_info.EnabledLayerCount, &_info.EnabledLayerNames).Free()
	defer fillNames(info.EnabledExtensions, &_info.EnabledExtensionCount, &_info.EnabledExtensionNames).Free()
	_info.Next = (*createInfo)(unsafe.Pointer(uintptr(0)))
	result = C.vkCreateInstance((*C.VkInstanceCreateInfo)(unsafe.Pointer(&_info)), nil, (*C.VkInstance)(unsafe.Pointer(&instance)))
	if result != C.VK_SUCCESS {
		return 0, Result(result)
	}
	return instance, nil
}

type Instance uintptr

func (i Instance) Destroy() {
	C.vkDestroyInstance((C.VkInstance)(unsafe.Pointer(i)), nil)
}

type Surface uintptr

type XlibSurfaceCreateInfo struct {
	Type    C.VkStructureType
	Next    *XlibSurfaceCreateInfo
	Flags   C.VkFlags
	Display uintptr
	Window  uintptr
}

func (i Instance) CreateXlibSurface(info XlibSurfaceCreateInfo) (Surface, error) {
	str := C.CString("vkCreateXlibSurfaceKHR")
	defer C.free(unsafe.Pointer(str))
	C._ptr_vkCreateXlibSurfaceKHR = C.vkGetInstanceProcAddr((C.VkInstance)(unsafe.Pointer(i)), str)
	var surface Surface
	info.Next = nil
	result := C.vkCreateXlibSurfaceKHR((C.VkInstance)(unsafe.Pointer(i)), (*C.struct_VkXlibSurfaceCreateInfoKHR)(unsafe.Pointer(&info)), nil, (*C.VkSurfaceKHR)(unsafe.Pointer(&surface)))
	if result != C.VK_SUCCESS {
		return 0, fmt.Errorf("surface error")
	}
	return surface, nil
}

func (i Instance) DestroySurface(surface Surface) {
	C.vkDestroySurfaceKHR((C.VkInstance)(unsafe.Pointer(i)), (C.VkSurfaceKHR)(unsafe.Pointer(surface)), nil)
}

type PhysicalDevice uintptr

type SwapchainCreateInfo struct {
	Type            C.VkStructureType
	Next            uintptr
	Flags           C.VkSwapchainCreateFlagsKHR
	Surface         C.VkSurfaceKHR
	MinImageCount   uint32
	Format          Format
	ImageColorSpace C.VkColorSpaceKHR
	ImageExtent     struct {
		Width  uint32
		Height uint32
	}
	ImageArrayLayers      uint32
	ImageUsage            C.VkImageUsageFlags
	ImageSharingMode      SharingMode
	QueueFamilyIndexCount uint32
	QueueFamilyIndices    *uint32
	PreTransform          C.VkSurfaceTransformFlagBitsKHR
	CompositeAlpha        C.VkCompositeAlphaFlagBitsKHR
	PresentMode           PresentMode
	Clipped               C.VkBool32
	OldSwapchain          C.VkSwapchainKHR
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
		p := C.malloc(C.size_t(uintptr(_info.SwapchainCount) * unsafe.Sizeof(Result(0))))
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

type DeviceCreateInfo struct {
	Type              StructureType
	Next              *DeviceCreateInfo
	Flags             C.VkDeviceCreateFlags
	QueueCreateInfos  []DeviceQueueCreateInfo
	EnabledLayers     []string
	EnabledExtensions []string
	EnabledFeatures   *C.VkPhysicalDeviceFeatures
}

type deviceCreateInfo struct {
	Type                  StructureType
	Next                  uintptr
	Flags                 C.VkDeviceCreateFlags
	QueueCreateInfoCount  uint32
	QueueCreateInfos      *DeviceQueueCreateInfo
	EnabledLayerCount     uint32
	EnabledLayerNames     *C.char
	EnabledExtensionCount uint32
	EnabledExtensionNames *C.char
	EnabledFeatures       *C.VkPhysicalDeviceFeatures
}

type DeviceQueueCreateInfo struct {
	Type             StructureType
	Next             uintptr
	Flags            C.VkDeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueueCount       uint32
	QueuePriorities  uintptr // todo
}

//type deviceQueueCreateInfo struct {
//	Type C.VkStructureType
//	Next uintptr
//	Flags C.VkDeviceQueueCreateFlags
//}

func (i Instance) EnumeratePhysicalDevices() ([]PhysicalDevice, error) {
	var count C.uint32_t
	// var devices uintptr
	// (*C.VkPhysicalDevice)(unsafe.Pointer(&devices))
	fmt.Println(unsafe.Sizeof(PhysicalDeviceProperties2KHR{}))
	result := C.vkEnumeratePhysicalDevices((C.VkInstance)(unsafe.Pointer(i)), &count, nil)
	if result != C.VK_SUCCESS {
		return nil, fmt.Errorf("vulkan error")
	}
	devices := make([]PhysicalDevice, count)
	result = C.vkEnumeratePhysicalDevices((C.VkInstance)(unsafe.Pointer(i)), &count, (*C.VkPhysicalDevice)(unsafe.Pointer(&devices[0])))
	// C._f = C.vkGetInstanceProcAddr((C.VkInstance)(unsafe.Pointer(i)), C.CString("vkGetPhysicalDeviceProperties2KHR"))
	// if C._f == nil {
	// 	panic("empty function pointer")
	// }
	for _, device := range devices {
		var properties PhysicalDeviceProperties
		// properties.Type = 1000059001
		// C.doInvoke((C.VkPhysicalDevice)(unsafe.Pointer(device)), (*C.VkPhysicalDeviceProperties2KHR)(unsafe.Pointer(&properties)))
		C.vkGetPhysicalDeviceProperties((C.VkPhysicalDevice)(unsafe.Pointer(device)), (*C.VkPhysicalDeviceProperties)(unsafe.Pointer(&properties)))
		fmt.Println("- physical device found:")
		// fmt.Println("  name:", string(propertieis.DeviceName[:bytes.IndexByte(properties.DeviceName[:], 0)]))
		// fmt.Println("  uuid:", string(properties.PipelineCacheUUID[:bytes.IndexByte(properties.PipelineCacheUUID[:], 0)]))
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

func (d PhysicalDevice) GetSurfaceSupport(queueFamilyIndex uint32, surface Surface) (bool, error) {
	var supported uint32
	result := C.vkGetPhysicalDeviceSurfaceSupportKHR((C.VkPhysicalDevice)(unsafe.Pointer(d)), (C.uint32_t)(queueFamilyIndex), (C.VkSurfaceKHR)(unsafe.Pointer(surface)), (*C.VkBool32)(unsafe.Pointer(&supported)))
	if result != C.VK_SUCCESS {
		return false, fmt.Errorf("surface support error")
	}
	return supported == C.VK_TRUE, nil
}
