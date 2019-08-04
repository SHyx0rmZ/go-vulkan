package vulkan

type DependencyFlagBits uint32
type DependencyFlags = DependencyFlagBits

const (
	DependencyByRegionBit DependencyFlagBits = 1 << iota
	DependencyDeviceGroupBit
	DependencyViewLocalBit
)
