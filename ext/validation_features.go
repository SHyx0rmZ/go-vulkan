package ext

// #include <stdlib.h>
import "C"
import (
	"unsafe"

	"code.witches.io/go/vulkan/base"
)

type InstanceCreateInfoNext interface {
	C() (unsafe.Pointer, func())
	instanceCreateInfoNext()
}

type ValidationFeaturesEXT struct {
	Next                       InstanceCreateInfoNext
	EnabledValidationFeatures  []ValidationFeatureEnableEXT
	DisabledValidationFeatures []ValidationFeatureDisableEXT
}

func (vf ValidationFeaturesEXT) instanceCreateInfoNext() {}

func allocUndef[T any](n int, x T) *T {
	return (*T)(C.malloc(C.size_t(uintptr(n) * unsafe.Sizeof(x))))
}

func allocZero[T any](n int, x T) *T {
	return (*T)(C.calloc(C.size_t(n), C.size_t(unsafe.Sizeof(x))))
}

func allocCopy[T any](x []T) *T {
	var _x T
	if len(x) == 0 { // todo: might need to remove this
		return nil
	}
	ptr := allocUndef(len(x), _x)
	copy(unsafe.Slice(ptr, len(x)), x)
	return ptr
}

func (vf ValidationFeaturesEXT) C() (unsafe.Pointer, func()) {
	_validationFeatures := (*validationFeaturesEXT)(C.malloc(C.size_t(unsafe.Sizeof(validationFeaturesEXT{}))))
	_enabled := &_validationFeatures.EnabledValidationFeatures
	_disabled := &_validationFeatures.DisabledValidationFeatures

	var nextPtr unsafe.Pointer
	if vf.Next != nil {
		var nextPtrFree func()
		nextPtr, nextPtrFree = vf.Next.C()
		defer nextPtrFree()
	}

	*_validationFeatures = validationFeaturesEXT{
		Type:                           base.StructureTypeValidationFeaturesEXT,
		Next:                           nextPtr,
		EnabledValidationFeatureCount:  uint32(len(vf.EnabledValidationFeatures)),
		DisabledValidationFeatureCount: uint32(len(vf.DisabledValidationFeatures)),
	}
	if c := len(vf.EnabledValidationFeatures); c > 0 {
		*_enabled = allocUndef(c, ValidationFeatureEnableEXT(0))
		copy(unsafe.Slice(*_enabled, c), vf.EnabledValidationFeatures)
	}
	if c := len(vf.DisabledValidationFeatures); c > 0 {
		*_disabled = allocCopy(vf.DisabledValidationFeatures)
	}
	return unsafe.Pointer(_validationFeatures), func() {
		if *_enabled != nil {
			C.free(unsafe.Pointer(*_enabled))
		}
		if *_disabled != nil {
			C.free(unsafe.Pointer(*_disabled))
		}
		C.free(unsafe.Pointer(_validationFeatures))
	}
}

type validationFeaturesEXT struct {
	Type                           base.StructureType
	Next                           unsafe.Pointer
	EnabledValidationFeatureCount  uint32
	EnabledValidationFeatures      *ValidationFeatureEnableEXT
	DisabledValidationFeatureCount uint32
	DisabledValidationFeatures     *ValidationFeatureDisableEXT
}

type ValidationFeatureEnableEXT uint32

const (
	ValidationFeatureEnableGPUAssistedEXT ValidationFeatureEnableEXT = iota
	ValidationFeatureEnableGPUAssistedReserveBindingSlotEXT
	ValidationFeatureEnableBestPracticesEXT
	ValidationFeatureEnableDebugPrintfEXT
	ValidationFeatureEnableSynchronizationValidationEXT
)

type ValidationFeatureDisableEXT uint32

const (
	ValidationFeatureDisableAllEXT ValidationFeatureDisableEXT = iota
	ValidationFeatureDisableShadersEXT
	ValidationFeatureDisableThreadSafetyEXT
	ValidationFeatureDisableAPIParametersEXT
	ValidationFeatureDisableObjectLifetimesEXT
	ValidationFeatureDisableCoreChecksEXT
	ValidationFeatureDisableUniqueHandlesEXT
	ValidationFeatureDisableShaderValidationCacheEXT
)
