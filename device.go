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
		ImageUsage:            C.VK_IMAGE_USAGE_TRANSFER_SRC_BIT | C.VK_IMAGE_USAGE_TRANSFER_DST_BIT,
		ImageSharingMode:      C.VK_SHARING_MODE_EXCLUSIVE,
		QueueFamilyIndexCount: 0,
		QueueFamilyIndices:    nil,
		PreTransform:          0x100,
		CompositeAlpha:        1,
		PresentMode:           C.VK_PRESENT_MODE_IMMEDIATE_KHR,
		Clipped:               C.VK_TRUE,
		OldSwapchain:          nil,
	}
	p := C.malloc(4)
	*(*uint32)(p) = 0
	info.QueueFamilyIndices = (*uint32)(p)
	defer C.free(p)
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

type Image uintptr

type ImageCreateInfo struct {
	Type      C.VkStructureType
	Next      uintptr
	Flags     C.VkImageCreateFlags
	ImageType C.VkImageType
	Format    C.VkFormat
	Extent    struct {
		Width  uint32
		Height uint32
		Depth  uint32
	}
	MipLevels             uint32
	ArrayLayers           uint32
	Samples               C.VkSampleCountFlagBits
	Tiling                C.VkImageTiling
	Usage                 C.VkImageUsageFlags
	SharingMode           C.VkSharingMode
	QueueFamilyIndexCount uint32
	QueueFamilyIndices    *uint32
	InitialLayout         C.VkImageLayout
}

func (d Device) GetSwapchainImages(swapchain Swapchain) {
	var count uint32
	result := C.vkGetSwapchainImagesKHR((C.VkDevice)(unsafe.Pointer(d)), (C.VkSwapchainKHR)(unsafe.Pointer(swapchain)), (*C.uint32_t)(unsafe.Pointer(&count)), nil)
	if result != C.VK_SUCCESS {
		panic("asdds")
	}
	images := make([]Image, count)
	result = C.vkGetSwapchainImagesKHR((C.VkDevice)(unsafe.Pointer(d)), (C.VkSwapchainKHR)(unsafe.Pointer(swapchain)), (*C.uint32_t)(unsafe.Pointer(&count)), (*C.VkImage)(unsafe.Pointer(&images[0])))
	if result != C.VK_SUCCESS {
		panic("asd98")
	}
	for _, image := range images {
		fmt.Printf("image: %#v\n", image)
	}
}

func (d Device) CreateImage() (Image, error) {
	var image Image
	info := ImageCreateInfo{
		//Type:      14,
		//Next:      0,
		//Flags:     C.VK_IMAGE_CREATE_,
		//ImageType: nil,
		//Format:    nil,
		//Extent: struct {
		//	Width  uint32
		//	Height uint32
		//	Depth  uint32
		//}{},
		//MipLevels:             0,
		//ArrayLayers:           0,
		//Samples:               nil,
		//Tiling:                nil,
		//Usage:                 nil,
		//SharingMode:           nil,
		//QueueFamilyIndexCount: 0,
		//QueueFamilyIndices:    nil,
		//InitialLayout:         nil,
	}
	result := C.vkCreateImage((C.VkDevice)(unsafe.Pointer(d)), (*C.VkImageCreateInfo)(unsafe.Pointer(&info)), nil, (*C.VkImage)(unsafe.Pointer(&image)))
	if result != C.VK_SUCCESS {
		return 0, fmt.Errorf("image error")
	}
	return image, nil
}

func (d Device) AcquireNextImage(swapchain Swapchain) (uint32, error) {
	var image uint32
	result := C.vkAcquireNextImageKHR((C.VkDevice)(unsafe.Pointer(d)), (C.VkSwapchainKHR)(unsafe.Pointer(swapchain)), C.uint64_t(^uint64(0)), (C.VkSemaphore)(unsafe.Pointer(nil)), (C.VkFence)(unsafe.Pointer(nil)), (*C.uint32_t)(unsafe.Pointer(&image)))
	if result != C.VK_SUCCESS {
		return 0, fmt.Errorf("image error")
	}
	return image, nil
}
