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

func CreateSurface(instance Instance, info sdl.WMInfo) (Surface, error) {
	if info.Subsystem != sdl.SubsystemX11 {
		return NullHandle, fmt.Errorf("unexpected subsystem while expecting '%s': %s", sdl.SubsystemX11, info.Subsystem)
	}

	xlib := *(*sdl.WMInfoXlib)(unsafe.Pointer(&info))
	return instance.CreateXlibSurface(XlibSurfaceCreateInfo{
		Type:    1000004000,
		Display: xlib.Display,
		Window:  xlib.Window,
	})
}

type XlibSurfaceCreateFlags uint32

type XlibSurfaceCreateInfo struct {
	Type    StructureType
	Next    uintptr
	Flags   XlibSurfaceCreateFlags
	Display uintptr
	Window  uintptr
}

func (i Instance) CreateXlibSurface(info XlibSurfaceCreateInfo) (Surface, error) {
	var surface Surface
	result := Result(C.vkCreateXlibSurfaceKHR(
		(C.VkInstance)(unsafe.Pointer(i)),
		(*C.VkXlibSurfaceCreateInfoKHR)(unsafe.Pointer(&info)),
		nil,
		(*C.VkSurfaceKHR)(unsafe.Pointer(&surface)),
	))
	if result != Success {
		return 0, result
	}
	return surface, nil
}
