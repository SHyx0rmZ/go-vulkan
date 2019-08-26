package vulkan

// #define VK_USE_PLATFORM_WIN32_KHR 1
// #include <vulkan/vulkan.h>
import "C"
import (
	"fmt"
	"unsafe"

	sdl "code.witches.io/go/sdl2"
)

const SurfaceExtension = "VK_KHR_win32_surface"

func CreateSurface(instance Instance, info sdl.WMInfo) (Surface, error) {
	if info.Subsystem != sdl.SubsystemWindows {
		return NullHandle, fmt.Errorf("unexpected subsystem while expecting '%s': %s", sdl.SubsystemWindows, info.Subsystem)
	}

	win32 := *(*sdl.WMInfoWin32)(unsafe.Pointer(&info))
	return instance.CreateWin32Surface(Win32SurfaceCreateInfo{
		Type:     1000009000,
		Instance: win32.Instance,
		Window:   win32.Window,
	})
}


type Win32SurfaceCreateFlags uint32

type Win32SurfaceCreateInfo struct {
	Type    StructureType
	Next    uintptr
	Flags   Win32SurfaceCreateFlags
	Instance uintptr
	Window  uintptr
}

func (i Instance) CreateWin32Surface(info Win32SurfaceCreateInfo) (Surface, error) {
	var surface Surface
	result := Result(C.vkCreateWin32SurfaceKHR(
		(C.VkInstance)(unsafe.Pointer(i)),
		(*C.VkWin32SurfaceCreateInfoKHR)(unsafe.Pointer(&info)),
		nil,
		(*C.VkSurfaceKHR)(unsafe.Pointer(&surface)),
	))
	if result != Success {
		return 0, result
	}
	return surface, nil
}
