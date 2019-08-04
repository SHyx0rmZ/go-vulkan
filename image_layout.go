package vulkan

type ImageLayout uint32

const (
	ImageLayoutUndefined ImageLayout = iota
	ImageLayoutGeneral
	ImageLayoutColorAttachmentOptimal
	ImageLayoutDepthStencilAttachmentOptimal
	ImageLayoutDepthStencilReadOnlyOptimal
	ImageLayoutShaderReadOnlyOptimal
	ImageLayoutTranserSrcOptimal
	ImageLayoutTransferDstOptimal
	ImageLayoutPreinitialized
)

const (
	ImageLayoutDepthReadOnlyStencilAttachmentOptimal ImageLayout = 1000117000 + iota
	ImageLayoutDepthAttachmentStencilReadOnlyOptimal
)

const (
	ImageLayoutPresentSrcKHR    ImageLayout = 1000001002
	ImageLayoutSharedPresentKHR ImageLayout = 1000111000
)
