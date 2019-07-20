package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type Queue uintptr

func (q Queue) Present(info PresentInfo) error {
	var _info presentInfo
	defer info.C(&_info).Free()
	result := C.vkQueuePresentKHR((C.VkQueue)(unsafe.Pointer(q)), (*C.VkPresentInfoKHR)(unsafe.Pointer(&_info)))
	if info.Results != nil {
		for i := range info.Results {
			result := *(*Result)(unsafe.Pointer(uintptr(unsafe.Pointer(_info.Results)) + uintptr(i)*unsafe.Sizeof(Result(0))))
			info.Results[i] = result
			fmt.Printf("swapchain #%d: %s", i, result)
		}
	}
	if result != C.VK_SUCCESS {
		return fmt.Errorf("present error")
	}
	return nil
}

//
//func (q Queue) Present(info PresentInfo, image uint32) error {
//	p := C.malloc(C.size_t(unsafe.Sizeof(C.uint32_t(0)) * 1))
//	*(*uint32)(p) = image
//	defer C.free(p)
//	info.ImageIndices = (*uint32)(p)
//	p2 := C.calloc(C.size_t(unsafe.Sizeof(Result(0))*uintptr(info.SwapchainCount)), 1)
//	defer C.free(p2)
//	info.Results = (*uint32)(p)
//	result := C.vkQueuePresentKHR((C.VkQueue)(unsafe.Pointer(q)), (*C.VkPresentInfoKHR)(unsafe.Pointer(&info)))
//	if result != C.VK_SUCCESS {
//		return fmt.Errorf("present error")
//	}
//	for i := 0; i < int(info.SwapchainCount); i++ {
//		fmt.Printf("swapchain #%d: %s", i, *(*Result)(unsafe.Pointer(uintptr(p2) + uintptr(i)*unsafe.Sizeof(Result(0)))))
//	}
//	return nil
//}
