package vulkan

type SampleCountFlagBits uint32
type SampleCountFlags SampleCountFlagBits

const (
	SampleCount1Bit SampleCountFlagBits = 1 << iota
	SampleCount2Bit
	SampleCount4Bit
	SampleCount8Bit
	SampleCount16Bit
	SampleCount32Bit
	SampleCount64Bit
)
