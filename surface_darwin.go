package vulkan

// #cgo CFLAGS: -x objective-c
// #define VK_USE_PLATFORM_MACOS_MVK 1
// #include <vulkan/vulkan.h>
// #include <Foundation/Foundation.h>
// #include <AppKit/AppKit.h>
import "C"
import (
	"fmt"
	"unsafe"

	sdl "code.witches.io/go/sdl2"
)

const SurfaceExtension = "VK_MVK_macos_surface"

func CreateSurface(instance Instance, info sdl.WMInfo, allocator *AllocationCallbacks) (Surface, error) {
	if info.Subsystem != sdl.SubsystemCocoa {
		return NullHandle, fmt.Errorf("unexpected subsystem while expecting '%s': %s", sdl.SubsystemCocoa, info.Subsystem)
	}

	cocoa := *(*sdl.WMInfoCocoa)(unsafe.Pointer(&info))

	window := unsafe.Pointer(cocoa.Window)
	view := (*uintptr)(unsafe.Pointer(uintptr(window) + 32))

	return CreateMacOSSurface(instance, MacOSSurfaceCreateInfo{
		Type: StructureTypeMacOSSurfaceCreateInfoMVK,
		View: *view,
	}, allocator)
}

type MacOSSurfaceCreateFlags uint32

type MacOSSurfaceCreateInfo struct {
	Type  StructureType
	Next  uintptr
	Flags MacOSSurfaceCreateFlags
	View  uintptr
}

func CreateMacOSSurface(instance Instance, info MacOSSurfaceCreateInfo, allocator *AllocationCallbacks) (Surface, error) {
	var surface Surface
	result := Result(C.vkCreateMacOSSurfaceMVK(
		(C.VkInstance)(unsafe.Pointer(instance)),
		(*C.VkMacOSSurfaceCreateInfoMVK)(unsafe.Pointer(&info)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkSurfaceKHR)(unsafe.Pointer(&surface)),
	))
	if result != Success {
		return NullHandle, result
	}
	return surface, nil
}
