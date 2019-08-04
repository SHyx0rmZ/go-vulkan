package vulkan

type AccessFlagBits uint32
type AccessFlags = AccessFlagBits

const (
	AccessIndirectCommandReadBit AccessFlagBits = 1 << iota
	AccessIndexReadBit
	AccessVertexAttributeReadBit
	AccessUniformReadBit
	AccessInputAttachmentReadBit
	AccessShaderReadBit
	AccessShaderWriteBit
	AccessColorAttachmentReadBit
	AccessColorAttachmentWriteBit
	AccessDepthStencilAttachmentReadBit
	AccessDepthStencilAttachmentWriteBit
	AccessTransferReadBit
	AccessTransferWriteBit
	AccessHostReadBit
	AccessHostWriteBit
	AccessMemoryReadBit
	AccessMemoryWriteBit
)
