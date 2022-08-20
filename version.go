package vulkan

import "fmt"

type Version uint32

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", VersionMajor(v), VersionMinor(v), VersionPatch(v))
}

func MakeVersion(major, minor, patch uint) Version {
	return Version(major<<22 | minor<<12 | patch)
}

func VersionMajor(v Version) uint {
	return uint(v >> 22)
}

func VersionMinor(v Version) uint {
	return uint(v >> 12 & 0x3ff)
}

func VersionPatch(v Version) uint {
	return uint(v & 0xfff)
}

const (
	APIVersion10 Version = 1<<22 | 0<<12 | 0
	APIVersion11 Version = 1<<22 | 1<<12 | 0
	APIVersion12 Version = 1<<22 | 2<<12 | 0
	APIVersion13 Version = 1<<22 | 3<<12 | 0

	HeaderVersion = 114
)

type ConformanceVersion struct {
	Major    uint8
	Minor    uint8
	SubMinor uint8
	Patch    uint8
}
