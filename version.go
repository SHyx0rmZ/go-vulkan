package vulkan

type Version uint32

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

	HeaderVersion = 114
)
