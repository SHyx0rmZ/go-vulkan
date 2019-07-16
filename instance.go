package vulkan

import "C"
import (
	"bytes"
	"fmt"
	"unsafe"
)

// #cgo linux freebsd darwin LDFLAGS: -lvulkan
// #include <stdlib.h>
// #include <X11/X.h>
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
// VkResult (*_ptr_vkCreateSwapchainKHR)(VkDevice device, const VkSwapchainCreateInfoKHR *info, const VkAllocationCallbacks *allocator, VkSwapchainKHR *swapchain);
// VkResult _vkCreateSwapchainKHR(VkDevice device, const VkSwapchainCreateInfoKHR *info, const VkAllocationCallbacks *allocator, VkSwapchainKHR *swapchain) {
//   return _ptr_vkCreateSwapchainKHR(device, info, allocator, swapchain);
// }
// void (*_ptr_vkDestroySwapchainKHR)(VkDevice device, VkSwapchainKHR swapchain, const VkAllocationCallbacks *allocator);
// void _vkDestroySwapchainKHR(VkDevice device, VkSwapchainKHR swapchain, const VkAllocationCallbacks *allocator) {
//   return _ptr_vkDestroySwapchainKHR(device, swapchain, allocator);
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

type Device uintptr

func (d Device) Destroy() {
	C.vkDestroyDevice((C.VkDevice)(unsafe.Pointer(d)), nil)
}

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

func (d Device) CreateSwapchain(info SwapchainCreateInfo, surface Surface) (Swapchain, error) {
	str := C.CString("vkCreateSwapchainKHR")
	defer C.free(unsafe.Pointer(str))
	C._ptr_vkCreateSwapchainKHR = C.vkGetDeviceProcAddr((C.VkDevice)(unsafe.Pointer(d)), str)
	var swapchain Swapchain
	fmt.Println("vkCreateSwapchainKHR", unsafe.Pointer(C._ptr_vkCreateSwapchainKHR))
	info = SwapchainCreateInfo{
		Type:            1000001000,
		Surface:         (C.VkSurfaceKHR)(unsafe.Pointer(surface)),
		MinImageCount:   2,
		Format:          27,
		ImageColorSpace: 0,
		ImageExtent: struct {
			Width  uint32
			Height uint32
		}{
			Width:  1280,
			Height: 800,
		},
		ImageArrayLayers:      1,
		ImageUsage:            0,
		ImageSharingMode:      C.VK_SHARING_MODE_EXCLUSIVE,
		QueueFamilyIndexCount: 0,
		QueueFamilyIndices:    nil,
		PreTransform:          0x100,
		CompositeAlpha:        1,
		PresentMode:           C.VK_PRESENT_MODE_IMMEDIATE_KHR,
		Clipped:               C.VK_TRUE,
		OldSwapchain:          nil,
	}
	fmt.Println("internal ", unsafe.Pointer(C.vkCreateSwapchainKHR))
	result := C.vkCreateSwapchainKHR((C.VkDevice)(unsafe.Pointer(d)), (*C.VkSwapchainCreateInfoKHR)(unsafe.Pointer(&info)), nil, (*C.VkSwapchainKHR)(unsafe.Pointer(&swapchain)))
	if result != C.VK_SUCCESS {
		return 0, fmt.Errorf("swapchain error")
	}
	return swapchain, nil
}

func (d Device) DestroySwapchain(swapchain Swapchain) {
	str := C.CString("vkDestroySwapchainKHR")
	defer C.free(unsafe.Pointer(str))
	C._ptr_vkDestroySwapchainKHR = C.vkGetDeviceProcAddr((C.VkDevice)(unsafe.Pointer(d)), str)
	C._vkDestroySwapchainKHR((C.VkDevice)(unsafe.Pointer(d)), (C.VkSwapchainKHR)(unsafe.Pointer(swapchain)), nil)
}

type Queue uintptr

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

func (q Queue) Present(info PresentInfo) error {
	p := C.malloc(C.size_t(unsafe.Sizeof(C.uint32_t(0)) * 1))
	*(*uint32)(p) = 0
	defer C.free(p)
	info.ImageIndices = (*uint32)(p)
	result := C.vkQueuePresentKHR((C.VkQueue)(unsafe.Pointer(q)), (*C.VkPresentInfoKHR)(unsafe.Pointer(&info)))
	if result != C.VK_SUCCESS {
		return fmt.Errorf("present error")
	}
	return nil
}

func (d Device) GetQueue(queueFamilyIndex, queueIndex uint32) Queue {
	var queue Queue
	C.vkGetDeviceQueue((C.VkDevice)(unsafe.Pointer(d)), C.uint32_t(queueFamilyIndex), C.uint32_t(queueIndex), (*C.VkQueue)(unsafe.Pointer(&queue)))
	return queue
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

func (d PhysicalDevice) CreateDevice(info DeviceCreateInfo) (Device, error) {
	var device Device
	_info := deviceCreateInfo{
		Type:                  info.Type,
		Next:                  0, // todo
		Flags:                 info.Flags,
		QueueCreateInfoCount:  uint32(len(info.QueueCreateInfos)),
		EnabledLayerCount:     uint32(len(info.EnabledLayers)),
		EnabledExtensionCount: uint32(len(info.EnabledExtensions)),
		EnabledFeatures:       info.EnabledFeatures,
	}
	if _info.QueueCreateInfoCount > 0 {
		sz := unsafe.Sizeof(DeviceQueueCreateInfo{})
		var l uintptr
		for range info.QueueCreateInfos {
			l += 1 * unsafe.Sizeof(float32(0))
		}
		p := C.malloc(C.size_t(len(info.QueueCreateInfos))*C.size_t(sz) + C.size_t(l))
		var o uintptr
		for _, info := range info.QueueCreateInfos {
			*(*DeviceQueueCreateInfo)(unsafe.Pointer(uintptr(p) + o)) = DeviceQueueCreateInfo{
				Type:             info.Type,
				Next:             info.Next,
				Flags:            info.Flags,
				QueueFamilyIndex: info.QueueFamilyIndex,
				QueueCount:       info.QueueCount,
				QueuePriorities:  uintptr(p) + o + sz,
			}
			o += sz
			*(*float32)(unsafe.Pointer(uintptr(p) + o)) = 1.0
			o += unsafe.Sizeof(float32(0))
		}
		_info.QueueCreateInfos = (*DeviceQueueCreateInfo)(p)
		defer func() {
			C.free(unsafe.Pointer(_info.QueueCreateInfos))
		}()
	}
	defer fillNames(info.EnabledLayers, &_info.EnabledLayerCount, &_info.EnabledLayerNames).Free()
	defer fillNames(info.EnabledExtensions, &_info.EnabledExtensionCount, &_info.EnabledExtensionNames).Free()
	result := C.vkCreateDevice((C.VkPhysicalDevice)(unsafe.Pointer(d)), (*C.VkDeviceCreateInfo)(unsafe.Pointer(&_info)), nil, (*C.VkDevice)(unsafe.Pointer(&device)))
	if result != C.VK_SUCCESS {
		return 0, fmt.Errorf("device error")
	}
	return device, nil
}

type PhysicalDeviceProperties2KHR struct {
	Type C.VkStructureType
	Next *PhysicalDeviceProperties2KHR
	PhysicalDeviceProperties
}

type PhysicalDeviceProperties struct {
	APIVersion        uint32
	DriverVersion     uint32
	VendorID          uint32
	DeviceID          uint32
	DeviceType        C.VkPhysicalDeviceType
	DeviceName        [C.VK_MAX_PHYSICAL_DEVICE_NAME_SIZE]uint8
	PipelineCacheUUID [C.VK_UUID_SIZE]uint8
	Limits            C.VkPhysicalDeviceLimits
	SparseProperties  C.VkPhysicalDeviceSparseProperties
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
