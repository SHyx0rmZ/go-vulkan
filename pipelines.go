package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type PipelineCache uintptr

type Pipeline uintptr

type PipelineCreateFlagBits uint32
type PipelineCreateFlags = PipelineCreateFlagBits

const (
	PipelineCreateDisableOptimizationBit PipelineCreateFlagBits = 1 << iota
	PipelineCreateAllowDerivativeBit
	PipelineCreateDerivativeBit
	PipelineCreateViewIndexFromDeviceIndexBit
	PipelineCreateDispatchBase
)

type PipelineShaderStageCreateFlags uint32

type ShaderStageFlagBits uint32
type ShaderStageFlags = ShaderStageFlagBits

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

type ComputePipelineCreateInfo struct{
	Type StructureType
	Next uintptr
	Flags PipelineCreateFlags
	Stage PipelineShaderStageCreateInfo
	Layout PipelineLayout
	BasePipelineHandle Pipeline
	BasePipelineIndex int32
}

func (info *ComputePipelineCreateInfo) C(_info *computePipelineCreateInfo) freeFunc {
	*_info = computePipelineCreateInfo{
		Type:   info.Type,
		Next:   info.Next,
		Flags:  info.Flags,
		Layout: info.Layout,
		BasePipelineHandle: info.BasePipelineHandle,
		BasePipelineIndex: info.BasePipelineIndex,
	}
	return info.Stage.C(&_info.Stage)
}

type computePipelineCreateInfo struct {
	Type StructureType
	Next uintptr
	Flags PipelineCreateFlags
	Stage pipelineShaderStageCreateInfo
	Layout PipelineLayout
	BasePipelineHandle Pipeline
	BasePipelineIndex int32
}

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
	fmt.Printf("%q @ %#v\n", info.Name, _info.Name)
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
type CullModeFlags = CullModeFlagBits

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

type PipelineVertexInputStateCreateFlags uint32

type PipelineVertexInputStateCreateInfo struct {
	Type                        StructureType
	Next                        uintptr
	Flags                       PipelineVertexInputStateCreateFlags
	VertexBindingDescriptions   []VertexInputBindingDescription
	VertexAttributeDescriptions []VertexInputAttributeDescription
}

func (info *PipelineVertexInputStateCreateInfo) C(_info *pipelineVertexInputStateCreateInfo) freeFunc {
	*_info = pipelineVertexInputStateCreateInfo{
		Type:                            info.Type,
		Next:                            info.Next,
		Flags:                           info.Flags,
		VertexBindingDescriptionCount:   uint32(len(info.VertexBindingDescriptions)),
		VertexAttributeDescriptionCount: uint32(len(info.VertexAttributeDescriptions)),
	}
	var ps []unsafe.Pointer
	if _info.VertexBindingDescriptionCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.VertexBindingDescriptionCount) * unsafe.Sizeof(VertexInputBindingDescription{})))
		ps = append(ps, p)
		for i, description := range info.VertexBindingDescriptions {
			*(*VertexInputBindingDescription)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(VertexInputBindingDescription{}))) = description
		}
		_info.VertexBindingDescriptions = (*VertexInputBindingDescription)(p)
	}
	if _info.VertexAttributeDescriptionCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.VertexAttributeDescriptionCount) * unsafe.Sizeof(VertexInputAttributeDescription{})))
		ps = append(ps, p)
		for i, description := range info.VertexAttributeDescriptions {
			*(*VertexInputAttributeDescription)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(VertexInputAttributeDescription{}))) = description
		}
		_info.VertexAttributeDescriptions = (*VertexInputAttributeDescription)(p)
	}
	return freeFunc(func() {
		for i := len(ps); i > 0; i-- {
			C.free(ps[i-1])
		}
	})
}

type pipelineVertexInputStateCreateInfo struct {
	Type                            StructureType
	Next                            uintptr
	Flags                           PipelineVertexInputStateCreateFlags
	VertexBindingDescriptionCount   uint32
	VertexBindingDescriptions       *VertexInputBindingDescription
	VertexAttributeDescriptionCount uint32
	VertexAttributeDescriptions     *VertexInputAttributeDescription
}

type VertexInputBindingDescription struct {
	Binding   uint32
	Stride    uint32
	InputRate VertexInputRate
}

type VertexInputRate uint32

const (
	VertexInputRateVertex VertexInputRate = iota
	VertexInputRateInstance
)

type VertexInputAttributeDescription struct {
	Location uint32
	Binding  uint32
	Format   Format
	Offset   uint32
}

type PipelineInputAssemblyStateCreateFlags uint32

type PipelineInputAssemblyStateCreateInfo struct {
	Type                   StructureType
	Next                   uintptr
	Flags                  PipelineInputAssemblyStateCreateFlags
	Topology               PrimitiveTopology
	PrimitiveRestartEnable bool
	_                      [3]byte
}

type PrimitiveTopology uint32

const (
	PrimitiveTopologyPointList PrimitiveTopology = iota
	PrimitiveTopologyLineList
	PrimitiveTopologyLineStrip
	PrimitiveTopologyTriangleList
	PrimitiveTopologyTriangleStrip
	PrimitiveTopologyTriangleFan
	PrimitiveTopologyLineListWithAdjacency
	PrimitiveTopologyLineStripWithAdjacency
	PrimitiveTopologyTriangleListWithAdjacency
	PrimitiveTopologyTriangleStripWithAdjacency
	PrimitiveTopologyPatchList
)

type PipelineTessellationStateCreateInfo struct{}
type PipelineViewportStateCreateFlags uint32
type PipelineViewportStateCreateInfo struct {
	Type      StructureType
	Next      uintptr
	Flags     PipelineViewportStateCreateFlags
	Viewports []Viewport
	Scissors  []Rect2D
}

func (info *PipelineViewportStateCreateInfo) C(_info *pipelineViewportStateCreateInfo) freeFunc {
	*_info = pipelineViewportStateCreateInfo{
		Type:          info.Type,
		Next:          info.Next,
		Flags:         info.Flags,
		ViewportCount: uint32(len(info.Viewports)),
		Viewports:     nil,
		ScissorCount:  uint32(len(info.Scissors)),
		Scissors:      nil,
	}
	var ps []unsafe.Pointer
	if _info.ViewportCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.ViewportCount) * unsafe.Sizeof(Viewport{})))
		ps = append(ps, p)
		for i, viewport := range info.Viewports {
			*(*Viewport)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(Viewport{}))) = viewport
		}
		_info.Viewports = (*Viewport)(p)
	}
	if _info.ScissorCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.ScissorCount) * unsafe.Sizeof(Rect2D{})))
		ps = append(ps, p)
		for i, scissor := range info.Scissors {
			*(*Rect2D)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(Rect2D{}))) = scissor
		}
		_info.Scissors = (*Rect2D)(p)
	}
	return freeFunc(func() {
		for i := len(ps); i > 0; i-- {
			C.free(ps[i-1])
		}
	})
}

type pipelineViewportStateCreateInfo struct {
	Type          StructureType
	Next          uintptr
	Flags         PipelineViewportStateCreateFlags
	ViewportCount uint32
	Viewports     *Viewport
	ScissorCount  uint32
	Scissors      *Rect2D
}

type Viewport struct {
	X        float32
	Y        float32
	Width    float32
	Height   float32
	MinDepth float32
	MaxDepth float32
}

type PipelineRasterizationStateCreateInfo struct {
	Type                    StructureType
	Next                    uintptr
	Flags                   PipelineRasterizationStateCreateFlags
	DepthClampEnable        bool
	_                       [3]byte
	RasterizerDiscardEnable bool
	_                       [3]byte
	PolygonMode             PolygonMode
	CullMode                CullModeFlags
	FrontFace               FrontFace
	DepthBiasEnable         bool
	_                       [3]byte
	DepthBiasConstantFactor float32
	DepthBiasClamp          float32
	DepthBiasSlopeFactor    float32
	LineWidth               float32
}

type PipelineMultisampleStateCreateFlags uint32
type PipelineMultisampleStateCreateInfo struct {
	Type                  StructureType
	Next                  uintptr
	Flags                 PipelineMultisampleStateCreateFlags
	RasterizationSamples  SampleCountFlagBits
	SampleShadingEnable   bool
	_                     [3]byte
	MinSampleShading      float32
	SampleMask            *SampleMask
	AlphaToCoverageEnable bool
	_                     [3]byte
	AlphaToOneEnable      bool
	_                     [3]byte
}

type SampleMask uint32

type PipelineDepthStencilStateCreateInfo struct {
	Type                  StructureType
	Next                  uintptr
	Flags                 PipelineDepthStencilStateCreateFlags
	DepthTestEnable       bool
	_                     [3]byte
	DepthWriteEnable      bool
	_                     [3]byte
	DepthCompareOp        CompareOp
	DepthBoundsTestEnable bool
	_                     [3]byte
	StencilTestEnable     bool
	_                     [3]byte
	Front                 StencilOpState
	Back                  StencilOpState
	MinDepthBounds        float32
	MaxDeptBounds         float32
}

type PipelineDepthStencilStateCreateFlags uint32

type StencilOpState struct {
	FailOp      StencilOp
	PassOp      StencilOp
	DepthFailOp StencilOp
	CompareOp   CompareOp
	CompareMask uint32
	WriteMask   uint32
	Reference   uint32
}

type StencilOp uint32

const (
	StencilOpKeep StencilOp = iota
	StencilOpZero
	StencilOpReplace
	StencilOpIncrementAndClamp
	StencilOpIncrementAndWrap
	StencilOpInvert
	StencilOpDecrementAndClamp
	StencilOpDecrementAndWrap
)

type CompareOp uint32

const (
	CompareOpNever CompareOp = iota
	CompareOpLess
	CompareOpEqual
	CompareOpLessOrEqual
	CompareOpGreater
	CompareOpNotEqual
	CompareOpGreaterOrEqual
	CompareOpAlways
)

type PipelineColorBlendStateCreateFlags uint32

type PipelineColorBlendAttachmentState struct {
	BlendEnable         bool
	_                   [3]byte
	SrcColorBlendFactor BlendFactor
	DstColorBlendFactor BlendFactor
	ColorBlendOp        BlendOp
	SrcAlphaBlendFactor BlendFactor
	DstAlphaBlendFactor BlendFactor
	AlphaBlendOp        BlendOp
	ColorWriteMask      ColorComponentFlags
}

type BlendFactor uint32

const (
	BlendFactorZero BlendFactor = iota
	BlendFactorOne
	BlendFactorSrcAlphaSaturate
	BlendFactorSrcColor
	BlendFactorOneMinusSrcColor
	BlendFactorDstColor
	BlendFactorOneMinusDstColor
	BlendFactorSrcAlpha
	BlendFactorOneMinusSrcAlpha
	BlendFactorDstAlpha
	BlendFactorOneMinusDstAlpha
	BlendFactorConstantColor
	BlendFactorOneMinusConstantColor
	BlendFactorConstantAlpha
	BlendFactorOneMinusConstantAlpha
	BlendFactorSrc1Color
	BlendFactorOneMinusSrc1Color
	BlendFactorSrc1Alpha
	BlendFactorOneMinusSrc1Alpha
)

type PipelineColorBlendStateCreateInfo struct {
	Type           StructureType
	Next           uintptr
	Flags          PipelineColorBlendStateCreateFlags
	LogicOpEnable  bool
	LogicOp        LogicOp
	Attachments    []PipelineColorBlendAttachmentState
	BlendConstants [4]float32
}

func (info *PipelineColorBlendStateCreateInfo) C(_info *pipelineColorBlendStateCreateInfo) freeFunc {
	*_info = pipelineColorBlendStateCreateInfo{
		Type:            info.Type,
		Next:            info.Next,
		Flags:           info.Flags,
		LogicOpEnable:   C.VK_FALSE,
		LogicOp:         info.LogicOp,
		AttachmentCount: uint32(len(info.Attachments)),
		Attachments:     nil,
		BlendConstants:  info.BlendConstants,
	}
	if info.LogicOpEnable {
		_info.LogicOpEnable = C.VK_TRUE
	}
	if _info.AttachmentCount > 0 {
		p := C.malloc(C.size_t(uintptr(_info.AttachmentCount) * unsafe.Sizeof(PipelineColorBlendAttachmentState{})))
		for i, attachment := range info.Attachments {
			*(*PipelineColorBlendAttachmentState)(unsafe.Pointer(uintptr(p) + uintptr(i)*unsafe.Sizeof(PipelineColorBlendAttachmentState{}))) = attachment
		}
		_info.Attachments = (*PipelineColorBlendAttachmentState)(p)
		return freeFunc(func() {
			C.free(p)
		})
	}
	return freeFunc(nil)
}

type pipelineColorBlendStateCreateInfo struct {
	Type            StructureType
	Next            uintptr
	Flags           PipelineColorBlendStateCreateFlags
	LogicOpEnable   C.VkBool32
	LogicOp         LogicOp
	AttachmentCount uint32
	Attachments     *PipelineColorBlendAttachmentState
	BlendConstants  [4]float32
}

type LogicOp uint32

const (
	LogicOpClear LogicOp = iota
	LogicOpAnd
	LogicOpAndReverse
	LogicOpCopy
	LogicOpAndInverted
	LogicOpNoOp
	LogicOpXor
	LogicOpOr
	LogicOpNor
	LogicOpEquivalent
	LogicOpInvert
	LogicOpOrReverse
	LogicOpCopyInverted
	LogicOpOrInverted
	LogicOpNand
	LogicOpSet
)

type BlendOp uint32

const (
	BlendOpAdd BlendOp = iota
	BlendOpSubtract
	BlendOpReverseSubtract
	BlendOpMin
	BlendOpMax
)

type ColorComponentFlagBits uint32
type ColorComponentFlags = ColorComponentFlagBits

const (
	ColorComponentRBit ColorComponentFlagBits = 1 << iota
	ColorComponentGBit
	ColorComponentBBit
	ColorComponentABit
)

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
		VertexInputState:   nil,
		InputAssemblyState: nil,
		TessellationState:  info.TessellationState,
		ViewportState:      nil,
		RasterizationState: nil,
		MultisampleState:   nil,
		DepthStencilState:  nil,
		ColorBlendState:    nil,
		DynamicState:       info.DynamicState,
		Layout:             info.Layout,
		RenderPass:         info.RenderPass,
		Subpass:            info.Subpass,
		BasePipelineHandle: info.BasePipelineHandle,
		BasePipelineIndex:  info.BasePipelineIndex,
	}
	var fs []freeFunc
	if len(info.Stages) > 0 {
		p := C.malloc(C.size_t(uintptr(_info.StageCount) * unsafe.Sizeof(pipelineShaderStageCreateInfo{})))
		fs = append(fs, freeFunc(func() {
			C.free(p)
		}))
		for i, stage := range info.Stages {
			fs = append(fs, stage.C((*pipelineShaderStageCreateInfo)(unsafe.Pointer(uintptr(p)+uintptr(i)*unsafe.Sizeof(pipelineShaderStageCreateInfo{})))))
		}
		_info.Stages = (*pipelineShaderStageCreateInfo)(p)
	}
	if info.VertexInputState != nil {
		p := C.malloc(C.size_t(unsafe.Sizeof(pipelineVertexInputStateCreateInfo{})))
		fs = append(fs, freeFunc(func() {
			C.free(p)
		}))
		_info.VertexInputState = (*pipelineVertexInputStateCreateInfo)(p)
		fs = append(fs, info.VertexInputState.C(_info.VertexInputState))
	}
	if info.InputAssemblyState != nil {
		p := C.malloc(C.size_t(unsafe.Sizeof(PipelineInputAssemblyStateCreateInfo{})))
		fs = append(fs, freeFunc(func() {
			C.free(p)
		}))
		_info.InputAssemblyState = (*PipelineInputAssemblyStateCreateInfo)(p)
		*_info.InputAssemblyState = *info.InputAssemblyState
	}
	if info.ViewportState != nil {
		p := C.malloc(C.size_t(unsafe.Sizeof(pipelineViewportStateCreateInfo{})))
		fs = append(fs, freeFunc(func() {
			C.free(p)
		}))
		_info.ViewportState = (*pipelineViewportStateCreateInfo)(p)
		fs = append(fs, info.ViewportState.C(_info.ViewportState))
	}
	if info.RasterizationState != nil {
		p := C.malloc(C.size_t(unsafe.Sizeof(PipelineRasterizationStateCreateInfo{})))
		fs = append(fs, freeFunc(func() {
			C.free(p)
		}))
		_info.RasterizationState = (*PipelineRasterizationStateCreateInfo)(p)
		*_info.RasterizationState = *info.RasterizationState
	}
	if info.MultisampleState != nil {
		p := C.malloc(C.size_t(unsafe.Sizeof(PipelineMultisampleStateCreateInfo{})))
		fs = append(fs, freeFunc(func() {
			C.free(p)
		}))
		_info.MultisampleState = (*PipelineMultisampleStateCreateInfo)(p)
		*_info.MultisampleState = *info.MultisampleState
	}
	if info.DepthStencilState != nil {
		p := C.malloc(C.size_t(unsafe.Sizeof(PipelineDepthStencilStateCreateInfo{})))
		fs = append(fs, freeFunc(func() {
			C.free(p)
		}))
		_info.DepthStencilState = (*PipelineDepthStencilStateCreateInfo)(p)
		*_info.DepthStencilState = *info.DepthStencilState
	}
	if info.ColorBlendState != nil {
		p := C.malloc(C.size_t(unsafe.Sizeof(pipelineColorBlendStateCreateInfo{})))
		fs = append(fs, freeFunc(func() {
			C.free(p)
		}))
		_info.ColorBlendState = (*pipelineColorBlendStateCreateInfo)(p)
		fs = append(fs, info.ColorBlendState.C(_info.ColorBlendState))
	}
	return freeFunc(func() {
		for i := len(fs); i > 0; i-- {
			fs[i-1]()
		}
	})
}

type graphicsPipelineCreateInfo struct {
	Type               StructureType
	Next               uintptr
	Flags              PipelineCreateFlags
	StageCount         uint32
	Stages             *pipelineShaderStageCreateInfo
	VertexInputState   *pipelineVertexInputStateCreateInfo
	InputAssemblyState *PipelineInputAssemblyStateCreateInfo
	TessellationState  *PipelineTessellationStateCreateInfo
	ViewportState      *pipelineViewportStateCreateInfo
	RasterizationState *PipelineRasterizationStateCreateInfo
	MultisampleState   *PipelineMultisampleStateCreateInfo
	DepthStencilState  *PipelineDepthStencilStateCreateInfo
	ColorBlendState    *pipelineColorBlendStateCreateInfo
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
	_createInfos := make([]computePipelineCreateInfo, len(createInfos))
	var ps []freeFunc
	for i, createInfo := range createInfos {
		ps = append(ps, createInfo.C(&_createInfos[i]))
		// ps = append(ps, createInfo.C((*computePipelineCreateInfo{})(unsafe.Pointer(&_createInfos[i]))))
	}
	defer func() {
		for i := len(ps); i > 0; i-- {
			ps[i-1].Free()
		}
	}()
	result := Result(C.vkCreateComputePipelines(
		(C.VkDevice)(unsafe.Pointer(device)),
		(C.VkPipelineCache)(unsafe.Pointer(pipelineCache)),
		(C.uint32_t)(len(_createInfos)),
		(*C.VkComputePipelineCreateInfo)(unsafe.Pointer(&_createInfos[0])),
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
		for i := len(ps); i > 0; i-- {
			ps[i-1].Free()
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
