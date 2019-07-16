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
	Type              C.VkStructureType
	Next              *CreateInfo
	Flags             C.VkInstanceCreateFlags
	ApplicationInfo   *ApplicationInfo
	EnabledLayers     []string
	EnabledExtensions []string
}

type createInfo struct {
	Type                  C.VkStructureType
	Next                  *CreateInfo
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

func CreateInstance(info CreateInfo) (Instance, error) {
	var instance Instance
	_info := createInfo{
		Type:                  info.Type,
		Next:                  nil, // todo
		Flags:                 info.Flags,
		ApplicationInfo:       info.ApplicationInfo,
		EnabledLayerCount:     uint32(len(info.EnabledLayers)),
		EnabledLayerNames:     nil,
		EnabledExtensionCount: uint32(len(info.EnabledExtensions)),
		EnabledExtensionNames: nil,
	}
	defer fillNames(info.EnabledLayers, &_info.EnabledLayerCount, &_info.EnabledLayerNames).Free()
	defer fillNames(info.EnabledExtensions, &_info.EnabledExtensionCount, &_info.EnabledExtensionNames).Free()
	result := C.vkCreateInstance((*C.struct_VkInstanceCreateInfo)(unsafe.Pointer(&_info)), nil, (*C.VkInstance)(unsafe.Pointer(&instance)))
	if result != C.VK_SUCCESS {
		return 0, fmt.Errorf("vulkan error")
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
	Format          C.VkFormat
	ImageColorSpace C.VkColorSpaceKHR
	ImageExtent     struct {
		Width  uint32
		Height uint32
	}
	ImageArrayLayers      uint32
	ImageUsage            C.VkImageUsageFlags
	ImageSharingMode      C.VkSharingMode
	QueueFamilyIndexCount uint32
	QueueFamilyIndices    *uint32
	PreTransform          C.VkSurfaceTransformFlagBitsKHR
	CompositeAlpha        C.VkCompositeAlphaFlagBitsKHR
	PresentMode           C.VkPresentModeKHR
	Clipped               C.VkBool32
	OldSwapchain          C.VkSwapchainKHR
}

type Swapchain uintptr

type PresentInfo struct {
	Type               C.VkStructureType
	Next               uintptr
	WaitSemaphoreCount uint32
	WaitSemaphores     uintptr
	SwapchainCount     uint32
	Swapchains         uintptr
	ImageIndices       *uint32
	Results            *uint32
}

type DeviceCreateInfo struct {
	Type              C.VkStructureType
	Next              *DeviceCreateInfo
	Flags             C.VkDeviceCreateFlags
	QueueCreateInfos  []DeviceQueueCreateInfo
	EnabledLayers     []string
	EnabledExtensions []string
	EnabledFeatures   *C.VkPhysicalDeviceFeatures
}

type deviceCreateInfo struct {
	Type                  C.VkStructureType
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
	Type             C.VkStructureType
	Next             uintptr
	Flags            C.VkDeviceQueueCreateFlags
	QueueFamilyIndex uint32
	QueueCount       uint32
	QueuePriorities  uintptr // todo
}

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
		fmt.Println("  name:", string(properties.DeviceName[:bytes.IndexByte(properties.DeviceName[:], 0)]))
		fmt.Println("  uuid:", string(properties.PipelineCacheUUID[:bytes.IndexByte(properties.PipelineCacheUUID[:], 0)]))
	}
	return devices, nil
}
