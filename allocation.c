#include <vulkan/vulkan.h>
#include "_cgo_export.h"

void *allocationCallback(void *userdata, size_t size, size_t alignment, VkSystemAllocationScope scope) {
    return allocationCallbackGo(userdata, size, alignment, scope);
}

void *reallocationCallback(void *userdata, void *original, size_t size, size_t alignment, VkSystemAllocationScope scope) {
    return reallocationCallbackGo(userdata, original, size, alignment, scope);
}

void freeCallback(void *userdata, void *memory) {
    freeCallbackGo(userdata, memory);
}

void internalAllocationCallback(void *userdata, size_t size, VkInternalAllocationType type, VkSystemAllocationScope scope) {
    internalAllocationCallbackGo(userdata, size, type, scope);
}

void internalFreeCallback(void *userdata, size_t size, VkInternalAllocationType type, VkSystemAllocationScope scope) {
    internalFreeCallbackGo(userdata, size, type, scope);
}