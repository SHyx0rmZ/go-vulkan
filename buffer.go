package vulkan

type Buffer uintptr

type BufferCreateFlagBits uint32
type BufferCreateFlags = BufferCreateFlagBits

const (
	BufferCreateSparseBindingsBit BufferCreateFlagBits = 1 << iota
	BufferCreateSparseResidencyBit
	BufferCreateSparseAliasedBit
	BufferCreateProtectedBit
)

type BufferCreateInfo struct {
	Type               StructureType
	Next               uintptr
	Flags              BufferCreateFlags
	Size               DeviceSize
	Usage              BufferUsageFlags
	SharingMode        SharingMode
	QueueFamilyIndices []uint32
}

type BufferUsageFlagBits uint32
type BufferUsageFlags = BufferUsageFlagBits

const (
	BufferUsageTransferSrcBit BufferUsageFlagBits = 1 << iota
	BufferUsageTransferDstBit
	BufferUsageUniformTexelBufferBit
	BufferUsageStorageTexelBufferBit
	BufferUsageUniformBufferBit
	BufferUsageStorageBufferBit
	BufferUsageIndexBufferBit
	BufferUsageVertexBufferBit
	BufferUsageIndirectBufferBit
)
