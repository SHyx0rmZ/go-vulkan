package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
// void *allocationCallback(void *, size_t, size_t, VkSystemAllocationScope);
// void *reallocationCallback(void *, void *, size_t, size_t, VkSystemAllocationScope);
// void freeCallback(void *, void *);
// void internalAllocationCallback(void *, size_t, VkInternalAllocationType, VkSystemAllocationScope);
// void internalFreeCallback(void *, size_t, VkInternalAllocationType, VkSystemAllocationScope);
import "C"
import (
	"sync"
	"unsafe"
)

//export allocationCallbackGo
func allocationCallbackGo(userData unsafe.Pointer, size, alignment int, allocationScope SystemAllocationScope) unsafe.Pointer {
	var id uint64
	if ptrSize == 8 {
		id = *(*uint64)(unsafe.Pointer(&userData))
	} else {
		id = *(*uint64)(userData)
	}
	return allocationCallbackMap.mapping[id].allocation(size, alignment, allocationScope)
}

//export reallocationCallbackGo
func reallocationCallbackGo(userData, original unsafe.Pointer, size, alignment int, allocationScope SystemAllocationScope) unsafe.Pointer {
	var id uint64
	if ptrSize == 8 {
		id = *(*uint64)(unsafe.Pointer(&userData))
	} else {
		id = *(*uint64)(userData)
	}
	return allocationCallbackMap.mapping[id].reallocation(original, size, alignment, allocationScope)
}

//export freeCallbackGo
func freeCallbackGo(userData, memory unsafe.Pointer) {
	var id uint64
	if ptrSize == 8 {
		id = *(*uint64)(unsafe.Pointer(&userData))
	} else {
		id = *(*uint64)(userData)
	}
	allocationCallbackMap.mapping[id].free(memory)
}

//export internalAllocationCallbackGo
func internalAllocationCallbackGo(userData unsafe.Pointer, size int, allocationType InternalAllocationType, allocationScope SystemAllocationScope) {
	var id uint64
	if ptrSize == 8 {
		id = *(*uint64)(unsafe.Pointer(&userData))
	} else {
		id = *(*uint64)(userData)
	}
	allocationCallbackMap.mapping[id].internalAllocation(size, allocationType, allocationScope)
}

//export internalFreeCallbackGo
func internalFreeCallbackGo(userData unsafe.Pointer, size int, allocationType InternalAllocationType, allocationScope SystemAllocationScope) {
	var id uint64
	if ptrSize == 8 {
		id = *(*uint64)(unsafe.Pointer(&userData))
	} else {
		id = *(*uint64)(userData)
	}
	allocationCallbackMap.mapping[id].internalFree(size, allocationType, allocationScope)
}

var allocationCallbackMap = struct {
	mapping map[uint64]*wrapperGo
	counter uint64
	sync.Mutex
}{
	mapping: make(map[uint64]*wrapperGo),
}

type wrapper struct {
	id                 unsafe.Pointer
	allocation         unsafe.Pointer
	reallocation       unsafe.Pointer
	free               unsafe.Pointer
	internalAllocation unsafe.Pointer
	internalFree       unsafe.Pointer
}

type wrapperGo struct {
	allocation         AllocationFunction
	reallocation       ReallocationFunction
	free               FreeFunction
	internalAllocation InternalAllocationNotification
	internalFree       InternalFreeNotification
}

type StdAllocator struct{}

func (StdAllocator) Allocate(size, alignment int, scope SystemAllocationScope) unsafe.Pointer {
	ptr := C.aligned_alloc(C.size_t(alignment), C.size_t(size))
	return ptr
}

func (a StdAllocator) Reallocate(original unsafe.Pointer, size, alignment int, scope SystemAllocationScope) unsafe.Pointer {
	ptr := C.realloc(original, C.size_t(size))
	if ptr == nil {
		return nil
	}
	if uintptr(ptr)%uintptr(alignment) == 0 {
		return ptr
	}
	ptr2 := a.Allocate(size, alignment, scope)
	if ptr2 == nil {
		return nil
	}
	copy(unsafe.Slice((*byte)(ptr2), size), unsafe.Slice((*byte)(ptr), size))
	return ptr2
}

func (StdAllocator) Free(memory unsafe.Pointer) {
	C.free(memory)
}

type AllocationFunction func(size, alignment int, allocationScope SystemAllocationScope) unsafe.Pointer
type ReallocationFunction func(original unsafe.Pointer, size, alignment int, allocationScope SystemAllocationScope) unsafe.Pointer
type FreeFunction func(memory unsafe.Pointer)
type InternalAllocationNotification func(size int, allocationType InternalAllocationType, allocationScope SystemAllocationScope)
type InternalFreeNotification func(size int, allocationType InternalAllocationType, allocationScope SystemAllocationScope)

func AllocationCallbacksWrapper(
	allocation AllocationFunction,
	reallocation ReallocationFunction,
	free FreeFunction,
	internalAllocation InternalAllocationNotification,
	internalFree InternalFreeNotification,
) *wrapper {
	w := (*wrapper)(C.calloc(1, C.size_t(unsafe.Sizeof(wrapper{}))))
	if allocation == nil {
		panic("must not be nil")
	}
	if reallocation == nil {
		panic("must not be nil")
	}
	if free == nil {
		panic("must not be nil")
	}
	w.allocation = unsafe.Pointer(C.allocationCallback)
	w.reallocation = unsafe.Pointer(C.reallocationCallback)
	w.free = unsafe.Pointer(C.freeCallback)
	switch {
	case internalAllocation == nil && internalFree == nil:
	case internalAllocation != nil && internalFree != nil:
		w.internalAllocation = unsafe.Pointer(C.internalAllocationCallback)
		w.internalFree = unsafe.Pointer(C.internalFreeCallback)
	default:
		panic("must not be nil")
	}
	c := &wrapperGo{
		allocation:         allocation,
		reallocation:       reallocation,
		free:               free,
		internalAllocation: internalAllocation,
		internalFree:       internalFree,
	}
	allocationCallbackMap.Lock()
	defer allocationCallbackMap.Unlock()
	if _, ok := allocationCallbackMap.mapping[allocationCallbackMap.counter]; ok {
		panic("too many allocation callbacks") // todo
	}
	allocationCallbackMap.mapping[allocationCallbackMap.counter] = c
	if ptrSize == 8 {
		w.id = *(*unsafe.Pointer)(unsafe.Pointer(&allocationCallbackMap.counter))
	} else {
		w.id = C.malloc(8)
		*(*uint64)(w.id) = allocationCallbackMap.counter
	}
	allocationCallbackMap.counter++
	return w
}

func (w *wrapper) Free() {
	if ptrSize != 8 {
		C.free(w.id)
	}
	C.free(unsafe.Pointer(w))
}
