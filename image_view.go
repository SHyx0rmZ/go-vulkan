package vulkan

type ImageViewCreateInfo struct {
	Type             StructureType
	Next             uintptr
	Flags            ImageViewCreateFlags
	Image            Image
	ViewType         ImageViewType
	Format           Format
	Components       ComponentMapping
	SubresourceRange ImageSubresourceRange
}

type ImageViewCreateFlagBits uint32
type ImageViewCreateFlags = ImageViewCreateFlagBits

type ImageView uint64

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
