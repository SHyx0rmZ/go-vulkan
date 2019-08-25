package vulkan

type ColorSpace uint32

const (
	ColorSpaceSRGBNonLinearKHR ColorSpace = iota

	ColorSpaceSRGBNonLinear = ColorSpaceSRGBNonLinearKHR
)

const (
	ColorSpaceDisplayP3NonLinear ColorSpace = iota + 1000104001
	ColorSpaceExtendedSRGBLinear
	ColorSpaceDisplayP3Linear
	ColorSpaceDCIP3NonLinear
)
