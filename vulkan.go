package vulkan

// #include <vulkan/vulkan.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type Result C.VkResult

type AllocationCallbacks C.VkAllocationCallbacks

type ImageViewCreateInfo struct {
	Type             C.VkStructureType
	Next             uintptr
	Flags            C.VkImageViewCreateFlags
	Image            Image
	ViewType         ImageViewType
	Format           C.VkFormat
	Components       ComponentMapping
	SubresourceRange ImageSubresourceRange
}

type ImageView uintptr

type ImageViewType uint32

const (
	ImageViewType1D ImageViewType = iota
	ImageViewType2D
	ImageViewType3D
	ImageViewTypeCube
	ImageViewType1DArray
	ImageViewType2DArray
	ImageViewTypeCubeArray
)

type ComponentMapping struct {
	R ComponentSwizzle
	G ComponentSwizzle
	B ComponentSwizzle
	A ComponentSwizzle
}

type ComponentSwizzle uint32

const (
	ComponentSwizzleIdentity ComponentSwizzle = iota
	ComponentSwizzleZero
	ComponentSwizzleOne
	ComponentSwizzleR
	ComponentSwizzleG
	ComponentSwizzleB
	ComponentSwizzleA
)

type ImageSubresourceRange struct {
	AspectMask     ImageAspectFlags
	BaseMIPLevel   uint32
	LevelCount     uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

type ImageAspectFlags uint32

const (
	ImageAspectColorBit ImageAspectFlags = 1 << iota
	ImageAspectDepthBit
	ImageAspectStencilBit
	ImageAspectMetadataBit
	ImageAspectPlane0Bit
	ImageAspectPlane1Bit
	ImageAspectPlane2Bit
)

func CreateImageView(device Device, createInfo ImageViewCreateInfo, allocator *AllocationCallbacks) (ImageView, error) {
	var view ImageView
	result := C.vkCreateImageView(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkImageViewCreateInfo)(unsafe.Pointer(&createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkImageView)(unsafe.Pointer(&view)),
	)
	if result != C.VK_SUCCESS {
		return 0, fmt.Errorf("CreateImageView")
	}
	return view, nil
}

func DestroyImageView(device Device, imageView ImageView, allocator *AllocationCallbacks) {
	C.vkDestroyImageView(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkImageView)(unsafe.Pointer(imageView)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

type RenderPass uint64

type RenderPassCreateInfo struct {
	Type C.VkStructureType
}

func CreateRenderPass(device Device, createInfo RenderPassCreateInfo, allocator *AllocationCallbacks) (RenderPass, error) {

}
