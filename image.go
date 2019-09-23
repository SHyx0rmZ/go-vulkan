package vulkan

type ImageType uint32

const (
	ImageType1D ImageType = iota
	ImageType2D
	ImageType3D
)

type ImageTiling uint32

const (
	ImageTilingOptimal ImageTiling = iota
	ImageTilingLinear
)

//go:generate stringer -type ImageUsageFlagBits -output image_usage_flags_string.go
type ImageUsageFlagBits uint32
type ImageUsageFlags = ImageUsageFlagBits

const (
	ImageUsageTransferSrcBit ImageUsageFlagBits = 1 << iota
	ImageUsageTransferDstBit
	ImageUsageSampledBit
	ImageUsageStorageBit
	ImageUsageColorAttachmentBit
	ImageUsageDepthStencilAttachmentBit
	ImageUsageTransientAttachmentBit
	ImageUsageInputAttachmentBit
)

type ImageSubresourceLayers struct {
	AspectMask     ImageAspectFlags
	MipLevel       uint32
	BaseArrayLayer uint32
	LayerCount     uint32
}

type BufferImageCopy struct {
	BufferOffset      DeviceSize
	BufferRowLength   uint32
	BufferImageHeight uint32
	ImageSubresource  ImageSubresourceLayers
	ImageOffset       Offset3D
	ImageExtent       Extent3D
}
