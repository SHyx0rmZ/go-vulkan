package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
// VkResult (*_ptr_vkCreateSwapchainKHR)(VkDevice device, const VkSwapchainCreateInfoKHR *info, const VkAllocationCallbacks *allocator, VkSwapchainKHR *swapchain);
// VkResult _vkCreateSwapchainKHR(VkDevice device, const VkSwapchainCreateInfoKHR *info, const VkAllocationCallbacks *allocator, VkSwapchainKHR *swapchain) {
//   return _ptr_vkCreateSwapchainKHR(device, info, allocator, swapchain);
// }
// void (*_ptr_vkDestroySwapchainKHR)(VkDevice device, VkSwapchainKHR swapchain, const VkAllocationCallbacks *allocator);
// void _vkDestroySwapchainKHR(VkDevice device, VkSwapchainKHR swapchain, const VkAllocationCallbacks *allocator) {
//   return _ptr_vkDestroySwapchainKHR(device, swapchain, allocator);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

type Device uintptr

func (d Device) Destroy() {
	C.vkDestroyDevice((C.VkDevice)(unsafe.Pointer(d)), nil)
}
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
func (d Device) GetQueue(queueFamilyIndex, queueIndex uint32) Queue {
	var queue Queue
	C.vkGetDeviceQueue((C.VkDevice)(unsafe.Pointer(d)), C.uint32_t(queueFamilyIndex), C.uint32_t(queueIndex), (*C.VkQueue)(unsafe.Pointer(&queue)))
	return queue
}
