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
