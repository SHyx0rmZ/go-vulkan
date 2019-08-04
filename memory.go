package vulkan

type MemoryAllocateInfo struct {
	Type            StructureType
	Next            uintptr
	AllocationSize  DeviceSize
	MemoryTypeIndex uint32
}

type DeviceMemory uintptr

type MemoryMapFlags uint32

type MemoryRequirements struct {
	Size           DeviceSize
	Alignment      DeviceSize
	MemoryTypeBits uint32
}

type PhysicalDeviceMemoryProperties struct {
	MemoryTypeCount uint32
	MemoryTypes     [32]MemoryType
	MemoryHeapCount uint32
	MemoryHeaps     [16]MemoryHeap
}

type MemoryHeapFlagBits uint32
type MemoryHeapFlags = MemoryHeapFlagBits

const (
	MemoryHeapDeviceLocalBit MemoryHeapFlagBits = 1 << iota
	MemoryHeapMultiInstanceBit
)

type MemoryHeap struct {
	Size  DeviceSize
	Flags MemoryHeapFlags
}

type MemoryPropertyFlagBits uint32
type MemoryPropertyFlags = MemoryPropertyFlagBits

const (
	MemoryPropertyDeviceLocalBit MemoryPropertyFlagBits = 1 << iota
	MemoryPropertyHostVisibleBit
	MemoryPropertyHostCoherentBit
	MemoryPropertyHostCachedBit
	MemoryPropertyLazilyAllocatedBit
	MemoryPropertyProtectedBit
)

type MemoryType struct {
	PropertyFlags MemoryPropertyFlags
	HeapIndex     uint32
}
