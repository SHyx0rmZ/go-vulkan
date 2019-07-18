package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

type PipelineCache uintptr

type Pipeline uintptr

type PipelineCreateFlags uint32
type PipelineShaderStageCreateFlags uint32

type ShaderStageFlagBits uint32

const (
	ShaderStageVertexBit ShaderStageFlagBits = 1 << iota
	ShaderStageTessellationControlBit
	ShaderStageTessellationEvaluationBit
	ShaderStageGeometryBit
	ShaderStageFragmentBit
	ShaderStageComputeBit
	ShaderStageAllGraphics = ShaderStageComputeBit - 1
	ShaderStageAll         = ^ShaderStageFlagBits(0x80000000)
)

type ComputePipelineCreateInfo struct{}

type PipelineShaderStageCreateInfo struct {
	Type               StructureType
	Next               uintptr
	Flags              PipelineShaderStageCreateFlags
	Stage              ShaderStageFlagBits
	Module             ShaderModule
	Name               string
	SpecializationInfo SpecializationInfo
}

type pipelineShaderStageCreateInfo struct {
	Type               StructureType
	Next               uintptr
	Flags              PipelineShaderStageCreateFlags
	Stage              ShaderStageFlagBits
	Module             ShaderModule
	Name               *C.char
	SpecializationInfo *specializationInfo
}

type SpecializationInfo struct {
	MapEntries []SpecializationMapEntry
	Data       []byte
}

type specializationInfo struct {
	MapEntryCount uint32
	MapEntries    *SpecializationMapEntry
	DataSize      C.size_t
	Data          *byte
}

type SpecializationMapEntry struct {
	ConstantID uint32
	Offset     uint32
	Size       int
}

type PipelineVertexInputStateCreateInfo struct{}
type PipelineInputAssemblyStateCreateInfo struct{}
type PipelineTessellationStateCreateInfo struct{}
type PipelineViewportStateCreateInfo struct{}
type PipelineRasterizationStateCreateInfo struct{}
type PipelineMultisampleStateCreateInfo struct{}
type PipelineDepthStencilStateCreateInfo struct{}
type PipelineColorBlendStateCreateInfo struct{}
type PipelineDynamicStateCreateInfo struct{}
type PipelineLayout struct{}
type GraphicsPipelineCreateInfo struct {
	Type               StructureType
	Next               uintptr
	Flags              PipelineCreateFlags
	Stages             []PipelineShaderStageCreateInfo
	VertexInputState   *PipelineVertexInputStateCreateInfo
	InputAssemblyState *PipelineInputAssemblyStateCreateInfo
	TessellationState  *PipelineTessellationStateCreateInfo
	ViewportState      *PipelineViewportStateCreateInfo
	RasterizationState *PipelineRasterizationStateCreateInfo
	MultisampleState   *PipelineMultisampleStateCreateInfo
	DepthStencilState  *PipelineDepthStencilStateCreateInfo
	ColorBlendState    *PipelineColorBlendStateCreateInfo
	DynamicState       *PipelineDynamicStateCreateInfo
	Layout             *PipelineLayout
	RenderPass         RenderPass
	Subpass            uint32
	BasePipelineHandle Pipeline
	BasePipelineIndex  int32
}

type graphicsPipelineCreateInfo struct {
	Type               StructureType
	Next               uintptr
	Flags              PipelineCreateFlags
	StageCount         uint32
	Stages             *PipelineShaderStageCreateInfo
	VertexInputState   *PipelineVertexInputStateCreateInfo
	InputAssemblyState *PipelineInputAssemblyStateCreateInfo
	TessellationState  *PipelineTessellationStateCreateInfo
	ViewportState      *PipelineViewportStateCreateInfo
	RasterizationState *PipelineRasterizationStateCreateInfo
	MultisampleState   *PipelineMultisampleStateCreateInfo
	DepthStencilState  *PipelineDepthStencilStateCreateInfo
	ColorBlendState    *PipelineColorBlendStateCreateInfo
	DynamicState       *PipelineDynamicStateCreateInfo
	Layout             *PipelineLayout
	RenderPass         RenderPass
	Subpass            uint32
	BasePipelineHandle Pipeline
	BasePipelineIndex  int32
}

type PipelineCacheCreateInfo struct{}

func CreateComputePipelines(device Device, pipelineCache PipelineCache, createInfos []ComputePipelineCreateInfo, allocator *AllocationCallbacks) ([]Pipeline, error) {
	pipelines := make([]Pipeline, len(createInfos))
	result := Result(C.vkCreateComputePipelines(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineCache)(unsafe.Pointer(pipelineCache)),
		(C.uint32_t)(len(createInfos)),
		(*C.VkComputePipelineCreateInfo)(unsafe.Pointer(&createInfos[0])),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkPipeline)(unsafe.Pointer(&pipelines[0])),
	))
	if result != Success {
		return nil, result
	}
	return pipelines, nil
}

func CreateGraphicsPipelines(device Device, pipelineCache PipelineCache, createInfos []GraphicsPipelineCreateInfo, allocator *AllocationCallbacks) ([]Pipeline, error) {
	pipelines := make([]Pipeline, len(createInfos))
	result := Result(C.vkCreateGraphicsPipelines(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineCache)(unsafe.Pointer(pipelineCache)),
		(C.uint32_t)(len(createInfos)),
		(*C.VkGraphicsPipelineCreateInfo)(unsafe.Pointer(&createInfos[0])),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkPipeline)(unsafe.Pointer(&pipelines[0])),
	))
	if result != Success {
		return nil, result
	}
	return pipelines, nil
}

func DestroyPipeline(device Device, pipeline Pipeline, allocator *AllocationCallbacks) {
	C.vkDestroyPipeline(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipeline(unsafe.Pointer(pipeline))),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

func CreatePipelineCache(device Device, createInfo PipelineCacheCreateInfo, allocator *AllocationCallbacks) (PipelineCache, error) {
	var pipelineCache PipelineCache
	result := Result(C.vkCreatePipelineCache(
		(C.VkDevice)(unsafe.Pointer(device)),
		(*C.VkPipelineCacheCreateInfo)(unsafe.Pointer(&createInfo)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
		(*C.VkPipelineCache)(unsafe.Pointer(&pipelineCache)),
	))
	if result != Success {
		return 0, result
	}
	return pipelineCache, nil
}

func MergePipelineCaches(device Device, dstCache PipelineCache, srcCaches []PipelineCache) error {
	result := Result(C.vkMergePipelineCaches(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineCache)(unsafe.Pointer(dstCache)),
		(C.uint32_t)(len(srcCaches)),
		(*C.VkPipelineCache)(unsafe.Pointer(&srcCaches[0])),
	))
	if result != Success {
		return result
	}
	return nil
}

func GetPipelineCacheData(device Device, pipelineCache PipelineCache, data []byte) (uint, error) {
	var size uint32
	result := Result(C.vkGetPipelineCacheData(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineCache)(unsafe.Pointer(pipelineCache)),
		(*C.size_t)(unsafe.Pointer(&size)),
		unsafe.Pointer(&data[0]),
	))
	if result != Success {
		return 0, result
	}
	return uint(size), nil
}

func DestroyPipelineCache(device Device, pipelineCache PipelineCache, allocator *AllocationCallbacks) {
	C.vkDestroyPipelineCache(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineCache)(unsafe.Pointer(pipelineCache)),
		(*C.VkAllocationCallbacks)(unsafe.Pointer(allocator)),
	)
}

func CmdBindPipeline(commandBuffer CommandBuffer, pipelineBindPoint PipelineBindPoint, pipeline Pipeline) {
	C.vkCmdBindPipeline(
		(C.VkCommandBuffer)(unsafe.Pointer(commandBuffer)),
		(C.VkPipelineBindPoint)(pipelineBindPoint),
		(C.VkPipeline)(unsafe.Pointer(pipeline)),
	)
}
