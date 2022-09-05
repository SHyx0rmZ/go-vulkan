package vulkan

// #include <vulkan/vulkan.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

type Flags64 uint64

type BufferMemoryBarrier2 struct {
	Type                StructureType
	Next                unsafe.Pointer
	SrcStageMask        PipelineStageFlags2
	SrcAccessMask       AccessFlags2
	DstStageMask        PipelineStageFlags2
	DstAccessMask       AccessFlags2
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Buffer              Buffer
	Offset              DeviceSize
	Size                DeviceSize
}

type CommandBufferSubmitInfo struct {
	Type          StructureType
	Next          unsafe.Pointer
	CommandBuffer CommandBuffer
	DeviceMask    uint32
}

type DependencyInfo struct {
	Type                 StructureType
	Next                 unsafe.Pointer
	DependencyFlags      DependencyFlags
	MemoryBarriers       []MemoryBarrier2
	BufferMemoryBarriers []BufferMemoryBarrier2
	ImageMemoryBarriers  []ImageMemoryBarrier2
}

func (info DependencyInfo) C(_info *dependencyInfo) freeFunc {
	var ps []unsafe.Pointer
	*_info = dependencyInfo{
		Type:                     info.Type,
		Next:                     info.Next,
		DependencyFlags:          info.DependencyFlags,
		MemoryBarrierCount:       uint32(len(info.MemoryBarriers)),
		BufferMemoryBarrierCount: uint32(len(info.BufferMemoryBarriers)),
		ImageMemoryBarrierCount:  uint32(len(info.ImageMemoryBarriers)),
	}
	if len(info.MemoryBarriers) > 0 {
		p := C.malloc(C.size_t(uintptr(len(info.MemoryBarriers)) * unsafe.Sizeof(MemoryBarrier2{})))
		ps = append(ps, p)
		copy(unsafe.Slice((*MemoryBarrier2)(p), len(info.MemoryBarriers)), info.MemoryBarriers)
		_info.MemoryBarrierPtr = (*MemoryBarrier2)(p)
	}
	if len(info.BufferMemoryBarriers) > 0 {
		ps = append(ps, copySliceToC(&_info.BufferMemoryBarrierPtr, info.BufferMemoryBarriers))
	}
	if len(info.ImageMemoryBarriers) > 0 {
		ps = append(ps, copySliceToC(&_info.ImageMemoryBarrierPtr, info.ImageMemoryBarriers))
	}
	return func() {
		for i := len(ps); i > 0; i-- {
			C.free(ps[i-1])
		}
	}
}

type dependencyInfo struct {
	Type                     StructureType
	Next                     unsafe.Pointer
	DependencyFlags          DependencyFlags
	MemoryBarrierCount       uint32
	MemoryBarrierPtr         *MemoryBarrier2
	BufferMemoryBarrierCount uint32
	BufferMemoryBarrierPtr   *BufferMemoryBarrier2
	ImageMemoryBarrierCount  uint32
	ImageMemoryBarrierPtr    *ImageMemoryBarrier2
}

type ImageMemoryBarrier2 struct {
	Type                StructureType
	Next                unsafe.Pointer
	SrcStageMask        PipelineStageFlags2
	SrcAccessMask       AccessFlags2
	DstStageMask        PipelineStageFlags2
	DstAccessMask       AccessFlags2
	OldLayout           ImageLayout
	NewLayout           ImageLayout
	SrcQueueFamilyIndex uint32
	DstQueueFamilyIndex uint32
	Image               Image
	SubresourceRange    ImageSubresourceRange
}

type SemaphoreSubmitInfo struct {
	Type        StructureType
	Next        unsafe.Pointer
	Semaphore   Semaphore
	Value       uint64
	StageMask   PipelineStageFlags2
	DeviceIndex uint32
}

type SubmitInfo2 struct {
	Type                 StructureType
	Next                 unsafe.Pointer
	Flags                SubmitFlags
	WaitSemaphoreInfos   []SemaphoreSubmitInfo
	CommandBufferInfos   []CommandBufferSubmitInfo
	SignalSemaphoreInfos []SemaphoreSubmitInfo
}

type A = SubmitInfo

func (info SubmitInfo2) C(_info *submitInfo2) {
	*_info = submitInfo2{
		Type:                     info.Type,
		Next:                     info.Next,
		Flags:                    info.Flags,
		WaitSemaphoreInfoCount:   uint32(len(info.WaitSemaphoreInfos)),
		CommandBufferInfoCount:   uint32(len(info.CommandBufferInfos)),
		SignalSemaphoreInfoCount: uint32(len(info.SignalSemaphoreInfos)),
	}
	if len(info.WaitSemaphoreInfos) > 0 {
		_info.WaitSemaphoreInfoPtr = &info.WaitSemaphoreInfos[0]
	}
	if len(info.CommandBufferInfos) > 0 {
		_info.CommandBufferInfoPtr = &info.CommandBufferInfos[0]
	}
	if len(info.SignalSemaphoreInfos) > 0 {
		_info.SignalSemaphoreInfoPtr = &info.SignalSemaphoreInfos[0]
	}
}

type submitInfo2 struct {
	Type                     StructureType
	Next                     unsafe.Pointer
	Flags                    SubmitFlags
	WaitSemaphoreInfoCount   uint32
	WaitSemaphoreInfoPtr     *SemaphoreSubmitInfo
	CommandBufferInfoCount   uint32
	CommandBufferInfoPtr     *CommandBufferSubmitInfo
	SignalSemaphoreInfoCount uint32
	SignalSemaphoreInfoPtr   *SemaphoreSubmitInfo
}

type PhysicalDeviceSynchronization2Features struct {
	Type             StructureType
	Next             *any
	Synchronization2 bool
	_                [3]byte
}

func (f *PhysicalDeviceSynchronization2Features) dciiInit(i *DeviceCreateInfoInterface) {
	f.Type = StructureTypePhysicalDeviceSynchronization2Features
	if i != nil {
		f.Next = (*any)(unsafe.Pointer(i))
	}
}

func (f *PhysicalDeviceSynchronization2Features) pdfiInit(i *PhysicalDeviceFeaturesInterface) {
	f.Type = StructureTypePhysicalDeviceSynchronization2Features
	if i != nil {
		f.Next = (*any)(unsafe.Pointer(i))
	}
}

func (f *PhysicalDeviceSynchronization2Features) dciiAlloc() (DeviceCreateInfoInterface, unsafe.Pointer) {
	ptr := C.calloc(1, (C.size_t)(unsafe.Sizeof(*f)))
	return (*PhysicalDeviceSynchronization2Features)(ptr), ptr
}

func (f *PhysicalDeviceSynchronization2Features) pdfiAlloc() (PhysicalDeviceFeaturesInterface, unsafe.Pointer) {
	ptr := C.calloc(1, (C.size_t)(unsafe.Sizeof(*f)))
	return (*PhysicalDeviceSynchronization2Features)(ptr), ptr
}

func (f *PhysicalDeviceSynchronization2Features) dciiCopy(i DeviceCreateInfoInterface) {
	*f = *(i.(*PhysicalDeviceSynchronization2Features))
}

func (f *PhysicalDeviceSynchronization2Features) pdfiCopy(i PhysicalDeviceFeaturesInterface) {
	*f = *(i.(*PhysicalDeviceSynchronization2Features))
}

type MemoryBarrier2 struct {
	Type          StructureType
	Next          unsafe.Pointer
	SrcStageMask  PipelineStageFlags2
	SrcAccessMask AccessFlags2
	DstStageMask  PipelineStageFlags2
	DstAccessMask AccessFlags2
}

type CheckpointData2NV struct {
	Type             StructureType
	Next             unsafe.Pointer
	Stage            PipelineStageFlags2
	CheckpointMarker unsafe.Pointer
}

type QueueFamilyCheckpointProperties2NV struct {
	Type                         StructureType
	Next                         unsafe.Pointer
	CheckpointExecutionStageMask PipelineStageFlags2
}

type AccessFlagBits2 Flags64
type AccessFlags2 = AccessFlagBits2

const (
	Access2None AccessFlagBits2 = 0
)

const (
	Access2IndirectCommandReadBit AccessFlagBits2 = 1 << iota
	Access2IndexReadBit
	Access2VertexAttributeReadBit
	Access2UniformReadBit
	Access2InputAttachmentReadBit
	Access2ShaderReadBit
	Access2ShaderWriteBit
	Access2ColorAttachmentReadBit
	Access2ColorAttachmentWriteBit
	Access2DepthStencilAttachmentReadBit
	Access2DepthStencilAttachmentWriteBit
	Access2TransferReadBit
	Access2TransferWriteBit
	Access2HostReadBit
	Access2HostWriteBit
	Access2MemoryReadBit
	Access2MemoryWriteBit
	Access2CommandPreprocessReadBitNV
	Access2CommandPreprocessWriteBitNV
	Access2ColorAttachmentReadNonCoherentBit
	Access2ConditionalRenderingReadBitEXT
	Access2AccelerationStructureReadBitKHR
	Access2AccelerationStructureWriteBitKHR
	Access2FragmentShadingRateAttachmentReadBitKHR
	Access2FragmentDensityMapReadBitEXT
	Access2TransformFeedbackWriteBitEXT
	Access2TransformFeedbackCounterReadBitEXT
	Access2TransformFeedbackCounterWriteBitEXT
)

const (
	Access2ShaderSampledReadBit AccessFlagBits2 = 1 << (iota + 32)
	Access2ShaderStorageReadBit
	Access2ShaderStorageWriteBit
	Access2VideoDecodeReadBit
	Access2VideoDecodeWriteBit
	Access2VideoEncodeReadBit
	Access2VideoEncodeWriteBit
	Access2InvocationMaskReadBitHUAWEI
	Access2ShaderBindingTableReadBitKHR
)

type PipelineStageFlagBits2 Flags64
type PipelineStageFlags2 = PipelineStageFlagBits2

const (
	PipelineStage2None PipelineStageFlagBits2 = 0
)

const (
	PipelineStage2TopOfPipeBit PipelineStageFlagBits2 = 1 << iota
	PipelineStage2DrawIndirectBit
	PipelineStage2VertexInputBit
	PipelineStage2VertexShaderBit
	PipelineStage2TessellationControlShaderBit
	PipelineStage2TessellationEvaluationShaderBit
	PipelineStage2GeometryShaderBit
	PipelineStage2FragmentShaderBit
	PipelineStage2EarlyFragmentTestsBit
	PipelineStage2LateFragmentTestsBit
	PipelineStage2ColorAttachmentOutputBit
	PipelineStage2ComputeShaderBit
	PipelineStage2TransferBit
	PipelineStage2BottomOfPipe
	PipelineStage2HostBit
	PipelineStage2AllGraphicsBit
	PipelineStage2AllCommandsBit
	PipelineStage2CommandPreprocessBitNV
	PipelineStage2ConditionalRenderingBitEXT
	PipelineStage2TaskShaderBitEXT
	PipelineStage2MeshShaderBitEXT
	PipelineStage2RayTracingShaderBitKHR
	PipelineStage2FragmentShadingRateAttachmentBitKHR
	PipelineStage2FragmentDensityProcessBitEXT
	PipelineStage2TransformFeedbackBitEXT
	PipelineStage2AccelerationStructureBuildBitKHR
	PipelineStage2VideoDecodeBitKHR
	PipelineStage2VideoEncodeBitKHR
	PipelineStage2AccelerationStructureCopyBitKHR
)

const (
	PipelineStage2CopyBit PipelineStageFlagBits2 = 1 << (iota + 32)
	PipelineStage2ResolveBit
	PipelineStage2BlitBit
	PipelineStage2ClearBit
	PipelineStage2IndexInputBit
	PipelineStage2VertexAttributeInputBit
	PipelineStage2PreRasterizationShadersBit
	PipelineStage2SubpassShadingBitHUAWEI
	PipelineStage2InvocationMaskBitHUAWEI
)

const (
	PipelineStage2AllTransferBit = PipelineStage2TransferBit
)

type SubmitFlagBits uint32
type SubmitFlags = SubmitFlagBits

const (
	SubmitProtectedBit SubmitFlagBits = 1 << iota
)

const (
	AccessNone                                          AccessFlagBits        = 0
	EventCreateDeviceOnlyBit                            EventCreateFlags      = 1
	ImageLayoutReadOnlyOptimal                          ImageLayout           = 1000314000
	ImageLayoutAttachmentOptimal                        ImageLayout           = 1000314001
	PipelineStageNone                                   PipelineStageFlagBits = 0
	StructureTypeBufferMemoryBarrier2                   StructureType         = 1000314001
	StructureTypeCommandBufferSubmitInfo                StructureType         = 1000314006
	StructureTypeDependencyInfo                         StructureType         = 1000314003
	StructureTypeImageMemoryBarrier2                    StructureType         = 1000314002
	StructureTypeMemoryBarrier2                         StructureType         = 1000314000
	StructureTypePhysicalDeviceSynchronization2Features StructureType         = 1000314007
	StructureTypeSemaphoreSubmitInfo                    StructureType         = 1000314005
	StructureTypeSubmitInfo2                            StructureType         = 1000314004
	StructureTypeQueueFamilyCheckpointProperties2NV     StructureType         = 1000314008
	StructureTypeCheckpointData2NV                      StructureType         = 1000314009
)

func CmdPipelineBarrier2(commandBuffer CommandBuffer, info DependencyInfo) {
	var _info dependencyInfo
	defer info.C(&_info)()
	C.vkCmdPipelineBarrier2(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		(*C.VkDependencyInfo)(unsafe.Pointer(&_info)),
	)
}

func CmdResetEvent2(commandBuffer CommandBuffer, event Event, stageMask PipelineStageFlags2) {
	C.vkCmdResetEvent2(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		*(*C.VkEvent)(unsafe.Pointer(&event)),
		(C.VkPipelineStageFlags2)(stageMask),
	)
}

func CmdSetEvent2(commandBuffer CommandBuffer, event Event, info DependencyInfo) {
	var _info dependencyInfo
	defer info.C(&_info)()
	C.vkCmdSetEvent2(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		*(*C.VkEvent)(unsafe.Pointer(&event)),
		(*C.VkDependencyInfo)(unsafe.Pointer(&_info)),
	)
}

func CmdWaitEvents2(commandBuffer CommandBuffer, events []Event, info DependencyInfo) {
	var _info dependencyInfo
	defer info.C(&_info)()
	C.vkCmdWaitEvents2(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		(C.uint32_t)(len(events)),
		(*C.VkEvent)(unsafe.Pointer(&events[0])),
		(*C.VkDependencyInfo)(unsafe.Pointer(&_info)),
	)
}

func CmdWriteTimestamp2(commandBuffer CommandBuffer, stage PipelineStageFlags2, queryPool QueryPool, query uint32) {
	C.vkCmdWriteTimestamp2(
		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
		(C.VkPipelineStageFlags2)(stage),
		*(*C.VkQueryPool)(unsafe.Pointer(&queryPool)),
		(C.uint32_t)(query),
	)
}

func QueueSubmit2(queue Queue, submits []SubmitInfo2, fence Fence) error {
	var submitPtr unsafe.Pointer
	if len(submits) > 0 {
		submitPtr = unsafe.Pointer(&submits[0])
	}
	result := Result(C.vkQueueSubmit2(
		*(*C.VkQueue)(unsafe.Pointer(&queue)),
		(C.uint32_t)(len(submits)),
		(*C.VkSubmitInfo2)(submitPtr),
		*(*C.VkFence)(unsafe.Pointer(&fence)),
	))
	if result != Success {
		return result
	}
	return nil
}

//func CmdWriteBufferMarker2AMD(commandBuffer CommandBuffer, stage PipelineStageFlags2, dstBuffer Buffer, dstOffset DeviceSize, marker uint32) {
//	C.vkCmdWriteBufferMarker2AMD(
//		*(*C.VkCommandBuffer)(unsafe.Pointer(&commandBuffer)),
//		(C.VkPipelineStageFlags2)(stage),
//		*(*C.VkBuffer)(unsafe.Pointer(&dstBuffer)),
//		(C.VkDeviceSize)(dstOffset),
//		(C.uint32_t)(marker),
//	)
//}
//
//func GetQueueCheckpointData2NV(queue Queue) []CheckpointData2NV {
//	var count uint32
//	C.vkGetQueueCheckpointData2NV(
//		*(*C.VkQueue)(unsafe.Pointer(&queue)),
//		(*C.uint32_t)(unsafe.Pointer(&count)),
//		nil,
//	)
//	data := make([]CheckpointData2NV, count)
//	for i := range data {
//		data[i].Type = StructureTypeCheckpointData2NV
//	}
//	C.vkGetQueueCheckpointData2NV(
//		*(*C.VkQueue)(unsafe.Pointer(&queue)),
//		(*C.uint32_t)(unsafe.Pointer(&count)),
//		(*C.VkCheckpointData2NV)(unsafe.Pointer(&data[0])),
//	)
//	return data
//}
