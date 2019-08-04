package vulkan

type SamplerCreateFlags uint32

type Filter uint32

const (
	FilterNearest Filter = iota
	FilterLinear
)

type SamplerMipMapMode uint32

const (
	SamplerMipMapModeNearest SamplerMipMapMode = iota
	SamplerMipMapModeLinear
)

type SamplerAddressMode uint32

const (
	SamplerAddressModeRepeat SamplerAddressMode = iota
	SamplerAddressModeMirroredRepeat
	SamplerAddressModeClampToEdge
	SamplerAddressModeClampToBorder
	SamplerAddressModeMirrorClampToEdge
)

type BorderColor uint32

const (
	BorderColorFloatTransparentBlack BorderColor = iota
	BorderColorIntTransparentBlack
	BorderColorFloatOpaqueBlack
	BorderColorIntOpaqueBlack
	BorderColorFloatOpaqueWhite
	BorderColorIntOpaqueWhite
)

type SamplerCreateInfo struct {
	Type                    StructureType
	Next                    uintptr
	Flags                   SamplerCreateFlags
	MagFilter               Filter
	MinFilter               Filter
	MipMapMode              SamplerMipMapMode
	AddressModeU            SamplerAddressMode
	AddressModeV            SamplerAddressMode
	AddressModeW            SamplerAddressMode
	MipLodBias              float32
	AnisotropyEnable        bool
	_                       [3]byte
	MaxAnisotropy           float32
	CompareEnable           bool
	_                       [3]byte
	CompareOp               CompareOp
	MinLod                  float32
	MaxLod                  float32
	BorderColor             BorderColor
	UnnormalizedCoordinates bool
	_                       [3]byte
}
