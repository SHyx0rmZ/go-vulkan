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

func CreateSurface(instance Instance, info sdl.WMInfo, allocator *AllocationCallbacks) (Surface, error) {
	if info.Subsystem != sdl.SubsystemWindows {
		return NullHandle, fmt.Errorf("unexpected subsystem while expecting '%s': %s", sdl.SubsystemWindows, info.Subsystem)
	}

	win32 := *(*sdl.WMInfoWin32)(unsafe.Pointer(&info))
	return CreateWin32Surface(instance, Win32SurfaceCreateInfo{
		Type:     StructureTypeWin32SurfaceCreateInfo,
		Instance: win32.Instance,
		Window:   win32.Window,
	}, allocator)
}

type Win32SurfaceCreateFlags uint32

type Win32SurfaceCreateInfo struct {
	Type     StructureType
	Next     uintptr
	Flags    Win32SurfaceCreateFlags
	Instance uintptr
	Window   uintptr
}

func CreateWin32Surface(instance Instance, info Win32SurfaceCreateInfo, allocator *AllocationCallbacks) (Surface, error) {
	var surface Surface
	result := Result(C.vkCreateWin32SurfaceKHR(
		(C.VkInstance)(unsafe.Pointer(i)),
		(*C.VkWin32SurfaceCreateInfoKHR)(unsafe.Pointer(&info)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkSurfaceKHR)(unsafe.Pointer(&surface)),
	))
	if result != Success {
		return NullHandle, result
	}
	return surface, nil
}
