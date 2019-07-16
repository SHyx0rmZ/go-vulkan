package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

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

const MaxPhysicalDeviceNameSize int = C.VK_MAX_PHYSICAL_DEVICE_NAME_SIZE
const UUIDSize int = C.VK_UUID_SIZE

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
