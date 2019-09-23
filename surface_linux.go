package vulkan

// #define VK_USE_PLATFORM_XLIB_KHR 1
// #include <vulkan/vulkan.h>
import "C"
import (
	"fmt"
	"unsafe"

	sdl "code.witches.io/go/sdl2"
)

const SurfaceExtension = "VK_KHR_xlib_surface"

func CreateSurface(instance Instance, info sdl.WMInfo, allocator *AllocationCallbacks) (Surface, error) {
	if info.Subsystem != sdl.SubsystemX11 {
		return NullHandle, fmt.Errorf("unexpected subsystem while expecting '%s': %s", sdl.SubsystemX11, info.Subsystem)
	}

	xlib := *(*sdl.WMInfoXlib)(unsafe.Pointer(&info))
	return CreateXlibSurface(instance, XlibSurfaceCreateInfo{
		Type:    StructureTypeXlibSurfaceCreateInfo,
		Display: xlib.Display,
		Window:  xlib.Window,
	}, allocator)
}

type XlibSurfaceCreateFlags uint32

type XlibSurfaceCreateInfo struct {
	Type    StructureType
	Next    uintptr
	Flags   XlibSurfaceCreateFlags
	Display uintptr
	Window  uintptr
}

func CreateXlibSurface(instance Instance, info XlibSurfaceCreateInfo, allocator *AllocationCallbacks) (Surface, error) {
	var surface Surface
	result := Result(C.vkCreateXlibSurfaceKHR(
		(C.VkInstance)(unsafe.Pointer(instance)),
		(*C.VkXlibSurfaceCreateInfoKHR)(unsafe.Pointer(&info)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkSurfaceKHR)(unsafe.Pointer(&surface)),
	))
	if result != Success {
		return NullHandle, result
	}
	return surface, nil
}
