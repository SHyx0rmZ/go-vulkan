package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type QueueFamilyProperties struct {
	QueueFlags                  C.VkQueueFlags
	QueueCount                  uint32
	TimestampValidBits          uint32
	MinImageTransferGranularity C.VkExtent3D
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
			for i := 0; i < int(info.QueueCount); i++ {
				*(*float32)(unsafe.Pointer(uintptr(p) + o)) = 1.0
				o += unsafe.Sizeof(float32(0))
			}
		}
		_info.QueueCreateInfos = (*DeviceQueueCreateInfo)(p)
		defer func() {
			C.free(unsafe.Pointer(_info.QueueCreateInfos))
		}()
	}
	defer fillNames(info.EnabledLayers, &_info.EnabledLayerCount, &_info.EnabledLayerNames).Free()
	defer fillNames(info.EnabledExtensions, &_info.EnabledExtensionCount, &_info.EnabledExtensionNames).Free()

	var count uint32
	C.vkGetPhysicalDeviceQueueFamilyProperties((C.VkPhysicalDevice)(unsafe.Pointer(d)), (*C.uint32_t)(unsafe.Pointer(&count)), nil)
	fmt.Println(count, "queue family properties")
	queueFamilyProperties := make([]QueueFamilyProperties, count)
	C.vkGetPhysicalDeviceQueueFamilyProperties((C.VkPhysicalDevice)(unsafe.Pointer(d)), (*C.uint32_t)(unsafe.Pointer(&count)), (*C.VkQueueFamilyProperties)(unsafe.Pointer(&queueFamilyProperties[0])))
	for _, p := range queueFamilyProperties {
		fmt.Printf("queue family properties: %+v\n", p)
	}

	result := C.vkCreateDevice((C.VkPhysicalDevice)(unsafe.Pointer(d)), (*C.VkDeviceCreateInfo)(unsafe.Pointer(&_info)), nil, (*C.VkDevice)(unsafe.Pointer(&device)))
	if result != C.VK_SUCCESS {
		return 0, Result(result)
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
	DeviceName        [MaxPhysicalDeviceNameSize]uint8
	PipelineCacheUUID [UUIDSize]uint8
	Limits            C.VkPhysicalDeviceLimits
	SparseProperties  C.VkPhysicalDeviceSparseProperties
}

type SurfaceFormat struct {
	Format     Format
	ColorSpace ColorSpace
}

func (d PhysicalDevice) GetSurfaceFormats(surface Surface) ([]SurfaceFormat, error) {
	var count uint32
	result := C.vkGetPhysicalDeviceSurfaceFormatsKHR((C.VkPhysicalDevice)(unsafe.Pointer(d)), (C.VkSurfaceKHR)(unsafe.Pointer(surface)), (*C.uint32_t)(unsafe.Pointer(&count)), nil)
	if result != C.VK_SUCCESS {
		return nil, fmt.Errorf("PhysicalDevice.GetSurfaceFormats")
	}
	formats := make([]SurfaceFormat, count)
	result = C.vkGetPhysicalDeviceSurfaceFormatsKHR((C.VkPhysicalDevice)(unsafe.Pointer(d)), (C.VkSurfaceKHR)(unsafe.Pointer(surface)), (*C.uint32_t)(unsafe.Pointer(&count)), (*C.VkSurfaceFormatKHR)(unsafe.Pointer(&formats[0])))
	if result != C.VK_SUCCESS {
		return nil, fmt.Errorf("PhysicalDevice.GetSurfaceFormats")
	}
	return formats, nil
}

type PresentMode C.VkPresentModeKHR

func (d PhysicalDevice) GetSurfacePresentModes(surface Surface) ([]PresentMode, error) {
	var count uint32
	result := Result(C.vkGetPhysicalDeviceSurfacePresentModesKHR(
		(C.VkPhysicalDevice)(unsafe.Pointer(d)),
		(C.VkSurfaceKHR)(unsafe.Pointer(surface)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		nil,
	))
	if result != Success {
		return nil, fmt.Errorf("previ: %s", result)
	}
	modes := make([]PresentMode, count)
	result = Result(C.vkGetPhysicalDeviceSurfacePresentModesKHR(
		(C.VkPhysicalDevice)(unsafe.Pointer(d)),
		(C.VkSurfaceKHR)(unsafe.Pointer(surface)),
		(*C.uint32_t)(unsafe.Pointer(&count)),
		(*C.VkPresentModeKHR)(unsafe.Pointer(&modes[0])),
	))
	if result != Success {
		return nil, result
	}
	return modes, nil
}

type SurfaceCapabilities struct {
	MinImageCount           uint32
	MaxImageCOunt           uint32
	CurrentExtent           Extent2D
	MinImageExtent          Extent2D
	MaxImageExtent          Extent2D
	MaxImageArrayLayers     uint32
	SupportedTransforms     C.VkSurfaceTransformFlagsKHR
	CurrentTransform        C.VkSurfaceTransformFlagBitsKHR
	SupportedCompositeAlpha C.VkCompositeAlphaFlagsKHR
	SupportedUsageFlags     C.VkImageUsageFlags
}

func (d PhysicalDevice) GetSurfaceCapabilities(surface Surface) (SurfaceCapabilities, error) {
	var capabilities SurfaceCapabilities
	result := C.vkGetPhysicalDeviceSurfaceCapabilitiesKHR((C.VkPhysicalDevice)(unsafe.Pointer(d)), (C.VkSurfaceKHR)(unsafe.Pointer(surface)), (*C.VkSurfaceCapabilitiesKHR)(unsafe.Pointer(&capabilities)))
	if result != C.VK_SUCCESS {
		return SurfaceCapabilities{}, fmt.Errorf("PhysicalDevice.GetSurfaceCapabilites")
	}
	return capabilities, nil
}
