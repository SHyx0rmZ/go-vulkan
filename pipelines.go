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
type ShaderStageFlags ShaderStageFlagBits

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

func (info *PipelineShaderStageCreateInfo) C(_info *pipelineShaderStageCreateInfo) freeFunc {
	*_info = pipelineShaderStageCreateInfo{
		Type:   info.Type,
		Next:   info.Next,
		Flags:  info.Flags,
		Stage:  info.Stage,
		Module: info.Module,
		Name:   C.CString(info.Name),
	}
	//info.SpecializationInfo
	if len(info.SpecializationInfo.Data) > 0 || len(info.SpecializationInfo.MapEntries) > 0 {
		panic("ikohasdoa ")
	}
	return freeFunc(func() {
		C.free(unsafe.Pointer(_info.Name))
	})
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

type PipelineRasterizationStateCreateFlagBits uint32
type PipelineRasterizationStateCreateFlags PipelineRasterizationStateCreateFlagBits

type PolygonMode uint32

const (
	PolygonModeFill PolygonMode = iota
	PolygonModeLine
	PolygonModePoint
)

type CullModeFlagBits uint32
type CullModeFlags CullModeFlagBits

const (
	CullModeNoneBit         CullModeFlagBits = 0
	CullModeFrontBit        CullModeFlagBits = 1
	CullModeBackBit         CullModeFlagBits = 2
	CullModeFrontAndBackBit                  = CullModeFrontBit | CullModeBackBit
)

type FrontFace uint32

const (
	FrontFaceCounterClockwise FrontFace = iota
	FrontFaceClockwise
)

type PipelineVertexInputStateCreateInfo struct{}
type PipelineInputAssemblyStateCreateInfo struct{}
type PipelineTessellationStateCreateInfo struct{}
type PipelineViewportStateCreateInfo struct{}
type PipelineRasterizationStateCreateInfo struct {
	Type                    StructureType
	Next                    uintptr
	Flags                   PipelineRasterizationStateCreateFlags
	DepthClampEnable        bool
	RasterizerDiscardEnable bool
	PolygonMode             PolygonMode
	CullMode                CullModeFlags
	FrontFace               FrontFace
	DepthBiasEnable         bool
	DepthBiasConstantFactor float32
	DepthBiasClamp          float32
	DepthBiasSlopeFactor    float32
	LineWidth               float32
}

func (info *PipelineRasterizationStateCreateInfo) C(_info *pipelineRasterizationStateCreateInfo) {
	if info == nil {
		return
	}
	*_info = pipelineRasterizationStateCreateInfo{
		Type:                    info.Type,
		Next:                    info.Next,
		Flags:                   info.Flags,
		DepthClampEnable:        C.VK_FALSE,
		RasterizerDiscardEnable: C.VK_FALSE,
		PolygonMode:             info.PolygonMode,
		CullMode:                info.CullMode,
		FrontFace:               info.FrontFace,
		DepthBiasEnable:         C.VK_FALSE,
		DepthBiasConstantFactor: info.DepthBiasConstantFactor,
		DepthBiasClamp:          info.DepthBiasClamp,
		DepthBiasSlopeFactor:    info.DepthBiasSlopeFactor,
		LineWidth:               info.LineWidth,
	}
	if info.DepthClampEnable {
		_info.DepthClampEnable = C.VK_TRUE
	}
	if info.RasterizerDiscardEnable {
		_info.RasterizerDiscardEnable = C.VK_TRUE
	}
	if info.DepthBiasEnable {
		_info.DepthBiasEnable = C.VK_TRUE
	}
}

type pipelineRasterizationStateCreateInfo struct {
	Type                    StructureType
	Next                    uintptr
	Flags                   PipelineRasterizationStateCreateFlags
	DepthClampEnable        C.VkBool32
	RasterizerDiscardEnable C.VkBool32
	PolygonMode             PolygonMode
	CullMode                CullModeFlags
	FrontFace               FrontFace
	DepthBiasEnable         C.VkBool32
	DepthBiasConstantFactor float32
	DepthBiasClamp          float32
	DepthBiasSlopeFactor    float32
	LineWidth               float32
}
type PipelineMultisampleStateCreateInfo struct{}
type PipelineDepthStencilStateCreateInfo struct{}
type PipelineColorBlendStateCreateInfo struct{}
type PipelineDynamicStateCreateInfo struct{}
type PipelineLayout uintptr
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
	Layout             PipelineLayout
	RenderPass         RenderPass
	Subpass            uint32
	BasePipelineHandle Pipeline
	BasePipelineIndex  int32
}

func (info *GraphicsPipelineCreateInfo) C(_info *graphicsPipelineCreateInfo) freeFunc {
	if info == nil {
		return freeFunc(nil)
	}
	*_info = graphicsPipelineCreateInfo{
		Type:               info.Type,
		Next:               info.Next,
		Flags:              info.Flags,
		StageCount:         uint32(len(info.Stages)),
		Stages:             nil,
		VertexInputState:   info.VertexInputState,
		InputAssemblyState: info.InputAssemblyState,
		TessellationState:  info.TessellationState,
		ViewportState:      info.ViewportState,
		MultisampleState:   info.MultisampleState,
		DepthStencilState:  info.DepthStencilState,
		ColorBlendState:    info.ColorBlendState,
		DynamicState:       info.DynamicState,
		Layout:             info.Layout,
		RenderPass:         info.RenderPass,
		Subpass:            info.Subpass,
		BasePipelineHandle: info.BasePipelineHandle,
		BasePipelineIndex:  info.BasePipelineIndex,
	}
	var fs []freeFunc
	if info.RasterizationState != nil {
		p := C.malloc(C.size_t(unsafe.Sizeof(pipelineRasterizationStateCreateInfo{})))
		fs = append(fs, freeFunc(func() {
			C.free(p)
		}))
		_info.RasterizationState = (*pipelineRasterizationStateCreateInfo)(p)
		info.RasterizationState.C(_info.RasterizationState)
	}
	if len(info.Stages) > 0 {
		p := C.malloc(C.size_t(uintptr(_info.StageCount) * unsafe.Sizeof(pipelineShaderStageCreateInfo{})))
		fs = append(fs, freeFunc(func() {
			C.free(p)
		}))
		for i, stage := range info.Stages {
			//*(*pipelineShaderStageCreateInfo)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(pipelineShaderStageCreateInfo{}))) =
			fs = append(fs, stage.C((*pipelineShaderStageCreateInfo)(unsafe.Pointer(uintptr(p)+uintptr(i)*unsafe.Sizeof(pipelineShaderStageCreateInfo{})))))
		}
		_info.Stages = (*pipelineShaderStageCreateInfo)(p)
	}
	return freeFunc(func() {
		for _, f := range fs {
			f()
		}
	})
}

type graphicsPipelineCreateInfo struct {
	Type               StructureType
	Next               uintptr
	Flags              PipelineCreateFlags
	StageCount         uint32
	Stages             *pipelineShaderStageCreateInfo
	VertexInputState   *PipelineVertexInputStateCreateInfo
	InputAssemblyState *PipelineInputAssemblyStateCreateInfo
	TessellationState  *PipelineTessellationStateCreateInfo
	ViewportState      *PipelineViewportStateCreateInfo
	RasterizationState *pipelineRasterizationStateCreateInfo
	MultisampleState   *PipelineMultisampleStateCreateInfo
	DepthStencilState  *PipelineDepthStencilStateCreateInfo
	ColorBlendState    *PipelineColorBlendStateCreateInfo
	DynamicState       *PipelineDynamicStateCreateInfo
	Layout             PipelineLayout
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
	_createInfos := make([]graphicsPipelineCreateInfo, len(createInfos))
	var ps []freeFunc
	for i, createInfo := range createInfos {
		ps = append(ps, createInfo.C((*graphicsPipelineCreateInfo)(unsafe.Pointer(&_createInfos[i]))))
	}
	defer func() {
		for _, p := range ps {
			p.Free()
		}
	}()
	result := Result(C.vkCreateGraphicsPipelines(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineCache)(unsafe.Pointer(pipelineCache)),
		(C.uint32_t)(len(_createInfos)),
		(*C.VkGraphicsPipelineCreateInfo)(unsafe.Pointer(&_createInfos[0])),
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
