package vulkan

import (
	"unsafe"
)

//go:generate go run stringer.go -type=Format -output=format_string.go
type Format uint32

const (
	FormatUndefined Format = iota
	FormatR4G4UNormPack8
	FormatR4G4B4A4UNormPack16
	FormatB4G4R4A4UNormPack16
	FormatR5G6B5UNormPack16
	FormatB5G6R5UNormPack16
	FormatR5G5B5A1UNormPack16
	FormatB5G5R5A1UNormPack16
	FormatA1R5G5B5UNormPack16
	FormatR8UNorm
	FormatR8SNorm
	FormatR8UScaled
	FormatR8SScaled
	FormatR8UInt
	FormatR8SInt
	FormatR8SRGB
	FormatR8G8UNorm
	FormatR8G8SNorm
	FormatR8G8UScaled
	FormatR8G8SScaled
	FormatR8G8UInt
	FormatR8G8SInt
	FormatR8G8SRGB
	FormatR8G8B8UNorm
	FormatR8G8B8SNorm
	FormatR8G8B8UScaled
	FormatR8G8B8SScaled
	FormatR8G8B8UInt
	FormatR8G8B8SInt
	FormatR8G8B8SRGB
	FormatB8G8R8UNorm
	FormatB8G8R8SNorm
	FormatB8G8R8UScaled
	FormatB8G8R8SScaled
	FormatB8G8R8UInt
	FormatB8G8R8SInt
	FormatB8G8R8SRGB
	FormatR8G8B8A8UNorm
	FormatR8G8B8A8SNorm
	FormatR8G8B8A8UScaled
	FormatR8G8B8A8SScaled
	FormatR8G8B8A8UInt
	FormatR8G8B8A8SInt
	FormatR8G8B8A8SRGB
	FormatB8G8R8A8UNorm
	FormatB8G8R8A8SNorm
	FormatB8G8R8A8UScaled
	FormatB8G8R8A8SScaled
	FormatB8G8R8A8UInt
	FormatB8G8R8A8SInt
	FormatB8G8R8A8SRGB
	FormatA8B8G8R8UNormPack32
	FormatA8B8G8R8SNormPack32
	FormatA8B8G8R8UScaledPack32
	FormatA8B8G8R8SScaledPack32
	FormatA8B8G8R8UIntPack32
	FormatA8B8G8R8SIntPack32
	FormatA8B8G8R8SRGBPack32
	FormatA2R10G10B10UNormPack32
	FormatA2R10G10B10SNormPack32
	FormatA2R10G10B10UScaledPack32
	FormatA2R10G10B10SScaledPack32
	FormatA2R10G10B10UIntPack32
	FormatA2R10G10B10SIntPack32
	FormatA2B10G10R10UNormPack32
	FormatA2B10G10R10SNormPack32
	FormatA2B10G10R10UScaledPack32
	FormatA2B10G10R10SScaledPack32
	FormatA2B10G10R10UIntPack32
	FormatA2B10G10R10SIntPack32
	FormatR16UNorm
	FormatR16SNorm
	FormatR16UScaled
	FormatR16SScaled
	FormatR16UInt
	FormatR16SInt
	FormatR16SFloat
	FormatR16G16UNorm
	FormatR16G16SNorm
	FormatR16G16UScaled
	FormatR16G16SScaled
	FormatR16G16UInt
	FormatR16G16SInt
	FormatR16G16SFloat
	FormatR16G16B16UNorm
	FormatR16G16B16SNorm
	FormatR16G16B16UScaled
	FormatR16G16B16SScaled
	FormatR16G16B16UInt
	FormatR16G16B16SInt
	FormatR16G16B16SFloat
	FormatR16G16B16A16UNorm
	FormatR16G16B16A16SNorm
	FormatR16G16B16A16UScaled
	FormatR16G16B16A16SScaled
	FormatR16G16B16A16UInt
	FormatR16G16B16A16SInt
	FormatR16G16B16A16SFloat
	FormatR32UInt
	FormatR32SInt
	FormatR32SFloat
	FormatR32G32UInt
	FormatR32G32SInt
	FormatR32G32SFloat
	FormatR32G32B32UInt
	FormatR32G32B32SInt
	FormatR32G32B32SFloat
	FormatR32G32B32A32UInt
	FormatR32G32B32A32SInt
	FormatR32G32B32A32SFloat
	FormatR64UInt
	FormatR64SInt
	FormatR64SFloat
	FormatR64G64UInt
	FormatR64G64SInt
	FormatR64G64SFloat
	FormatR64G64B64UInt
	FormatR64G64B64SInt
	FormatR64G64B64SFloat
	FormatR64G64B64A64UInt
	FormatR64G64B64A64SInt
	FormatR64G64B64A64SFloat
	FormatB10G11R11UFloatPack32
	FormatE5B9G9R9UFloatPack32
	FormatD16UNorm
	FormatX8D24UNormPack32
	FormatD32SFloat
	FormatS8UInt
	FormatD16UNormS8UInt
	FormatD24UNormS8UInt
	FormatD32SFloatS8UInt
	FormatBC1RGBUNormBlock
	FormatBC1RGBSRGBBlock
	FormatBC1RGBAUNormBlock
	FormatBC1RGBASRGBBlock
	FormatBC2UNormBlock
	FormatBC2SRGBBlock
	FormatBC3UNormBlock
	FormatBC3SRGBBlock
	FormatBC4UNormBlock
	FormatBC4SNormBlock
	FormatBC5UNormBlock
	FormatBC5SNormBlock
	FormatBC6HUFloatBlock
	FormatBC6HSFloatBlock
	FormatBC7UNormBlock
	FormatBC7SRGBBlock
	FormatETC2R8G8B8UNormBlock
	FormatETC2R8G8B8SRGBBlock
	FormatETC2R8G8B8A1UNormBlock
	FormatETC2R8G8B8A1SRGBBlock
	FormatETC2R8G8B8A8UNormBlock
	FormatETC2R8G8B8A8SRGBBlock
	FormatEACR11UNormBlock
	FormatEACR11SRGBBlock
	FormatEACR11G11UNormBlock
	FormatEACR11G11SRGBBlock
	FormatASTC4x4UNormBlock
	FormatASTC4x4SRGBBlock
	FormatASTC5x4UNormBlock
	FormatASTC5x4SRGBBlock
	FormatASTC5x5UNormBlock
	FormatASTC5x5SRGBBlock
	FormatASTC6x5UNormBlock
	FormatASTC6x5SRGBBlock
	FormatASTC6x6UNormBlock
	FormatASTC6x6SRGBBlock
	FormatASTC8x5UNormBlock
	FormatASTC8x5SRGBBlock
	FormatASTC8x6UNormBlock
	FormatASTC8x6SRGBBlock
	FormatASTC8x8UNormBlock
	FormatASTC8x8SRGBBlock
	FormatASTC10x5UNormBlock
	FormatASTC10x5SRGBBlock
	FormatASTC10x6UNormBlock
	FormatASTC10x6SRGBBlock
	FormatASTC10x8UNormBlock
	FormatASTC10x8SRGBBlock
	FormatASTC10x10UNormBlock
	FormatASTC10x10SRGBBlock
	FormatASTC12x10UNormBlock
	FormatASTC12x10SRGBBlock
	FormatASTC12x12UNormBlock
	FormatASTC12x12SRGBBlock
)

const (
	FormatG8B8G8R8422UNorm Format = iota + 1000156000
	FormatB8G8R8G8422UNorm
)

type ColorR16G16B16A16UInt uint64

func (c *ColorR16G16B16A16UInt) RGBA() (r, g, b, a uint32) {
	if c == nil {
		return
	}
	const m = (1 << 16) - 1
	pr := (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + 0))
	pg := (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + 2))
	pb := (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + 4))
	pa := (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(c)) + 8))
	r = uint32(*pr) * uint32(*pa) / m
	g = uint32(*pg) * uint32(*pa) / m
	b = uint32(*pb) * uint32(*pa) / m
	a = uint32(*pa)
	return r, g, b, a
}
