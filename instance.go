package vulkan

import "C"
import (
	"fmt"
	"unsafe"
)

// #cgo linux freebsd darwin LDFLAGS: -lvulkan
// #include <stdlib.h>
// #include <vulkan/vulkan.h>
// void (*_f)(VkPhysicalDevice device, VkPhysicalDeviceProperties2KHR *properties);
// void doInvoke(VkPhysicalDevice device, VkPhysicalDeviceProperties2KHR *properties) {
//   _f(device, properties);
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
	if _info.EnabledExtensionCount > 0 {
		var l int
		for _, ext := range info.EnabledExtensions {
			l += len(ext) + 1
		}
		p := C.malloc(C.size_t(len(info.EnabledExtensions)) * C.size_t(unsafe.Sizeof(uintptr(0))))
		fmt.Println(p)
		var o uintptr
		for _, ext := range info.EnabledExtensions {
			*(**C.char)(unsafe.Pointer(uintptr(p) + o)) = C.CString(ext)
			o += unsafe.Sizeof(uintptr(0))
			// for _, c := range []byte(ext) {
			// 	*(*uint8)(unsafe.Pointer(uintptr(p) + o)) = c
			// 	o++
			// }
			// *(*uint8)(unsafe.Pointer(uintptr(p) + o)) = 0
			// o++
		}
		_info.EnabledExtensionNames = (*C.char)(p)
		defer C.free(unsafe.Pointer(_info.EnabledExtensionNames))
	}
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

type PhysicalDevice uintptr

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
	C._f = C.vkGetInstanceProcAddr((C.VkInstance)(unsafe.Pointer(i)), C.CString("vkGetPhysicalDeviceProperties2KHR"))
	if C._f == nil {
		panic("empty function pointer")
	}
	for _, device := range devices {
		var properties PhysicalDeviceProperties2KHR
		properties.Type = 1000059001
		C.doInvoke((C.VkPhysicalDevice)(unsafe.Pointer(device)), (*C.VkPhysicalDeviceProperties2KHR)(unsafe.Pointer(&properties)))
	}
	return nil, nil
}
