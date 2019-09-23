package vulkan

//go:generate stringer -type=StructureType -output=structure_type_string.go
type StructureType uint32

// cat s.txt | head -n 49 | sed 's/[A-Z]/\L&/g;s/_\([a-z]\)/\U\1/g;s/^\s\+vk//;s/\s=\s[0-9]\+,$//'

const (
	StructureTypeApplicationInfo StructureType = iota
	StructureTypeInstanceCreateInfo
	StructureTypeDeviceQueueCreateInfo
	StructureTypeDeviceCreateInfo
	StructureTypeSubmitInfo
	StructureTypeMemoryAllocateInfo
	StructureTypeMappedMemoryRange
	StructureTypeBindSparseInfo
	StructureTypeFenceCreateInfo
	StructureTypeSemaphoreCreateInfo
	StructureTypeEventCreateInfo
	StructureTypeQueryPoolCreateInfo
	StructureTypeBufferCreateInfo
	StructureTypeBufferViewCreateInfo
	StructureTypeImageCreateInfo
	StructureTypeImageViewCreateInfo
	StructureTypeShaderModuleCreateInfo
	StructureTypePipelineCacheCreateInfo
	StructureTypePipelineShaderStageCreateInfo
	StructureTypePipelineVertexInputStateCreateInfo
	StructureTypePipelineInputAssemblyStateCreateInfo
	StructureTypePipelineTessellationStateCreateInfo
	StructureTypePipelineViewportStateCreateInfo
	StructureTypePipelineRasterizationStateCreateInfo
	StructureTypePipelineMultisampleStateCreateInfo
	StructureTypePipelineDepthStencilStateCreateInfo
	StructureTypePipelineColorBlendStateCreateInfo
	StructureTypePipelineDynamicStateCreateInfo
	StructureTypeGraphicsPipelineCreateInfo
	StructureTypeComputePipelineCreateInfo
	StructureTypePipelineLayoutCreateInfo
	StructureTypeSamplerCreateInfo
	StructureTypeDescriptorSetLayoutCreateInfo
	StructureTypeDescriptorPoolCreateInfo
	StructureTypeDescriptorSetAllocateInfo
	StructureTypeWriteDescriptorSet
	StructureTypeCopyDescriptorSet
	StructureTypeFramebufferCreateInfo
	StructureTypeRenderPassCreateInfo
	StructureTypeCommandPoolCreateInfo
	StructureTypeCommandBufferAllocateInfo
	StructureTypeCommandBufferInheritanceInfo
	StructureTypeCommandBufferBeginInfo
	StructureTypeRenderPassBeginInfo
	StructureTypeBufferMemoryBarrier
	StructureTypeImageMemoryBarrier
	StructureTypeMemoryBarrier
	StructureTypeLoaderInstanceCreateInfo
	StructureTypeLoaderDeviceCreateInfo
)

const (
	StructureTypeDisplayModeCreateInfo      StructureType = 1000002000
	StructureTypeDisplaySurfaceCreateInfo   StructureType = 1000002001
	StructureTypeDisplayPresentInfo         StructureType = 1000003000
	StructureTypeXlibSurfaceCreateInfo      StructureType = 1000004000
	StructureTypeXCBSurfaceCreateInfo       StructureType = 1000005000
	StructureTypeWaylandSurfaceCreateInfo   StructureType = 1000006000
	StructureTypeAndroidSurfaceCreateInfo   StructureType = 1000008000
	StructureTypeWin32SurfaceCreateInfo     StructureType = 1000009000
	StructureTypeVISurfaceCreateInfo        StructureType = 1000062000
	StructureTypeIOSSurfaceCreateInfo       StructureType = 1000122000
	StructureTypeMacOSSurfaceCreateInfo     StructureType = 1000123000
	StructureTypeImagePipeSurfaceCreateInfo StructureType = 1000214000
)
const StructureTypePhysicalDeviceGroupProperties StructureType = 1000070000
const StructureTypeQueueFamilyProperties2 StructureType = 1000059005
