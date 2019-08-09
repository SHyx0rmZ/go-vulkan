package vulkan

type CommandBufferLevel uint32

const (
	CommandBufferLevelPrimary CommandBufferLevel = iota
	CommandBufferLevelSecondary
)

type CommandBufferAllocateInfo struct {
	Type               StructureType
	Next               uintptr
	CommandPool        CommandPool
	Level              CommandBufferLevel
	CommandBufferCount uint32
}

type CommandBufferUsageFlagBits uint32
type CommandBufferUsageFlags = CommandBufferUsageFlagBits

const (
	CommandBufferUsageOneTimeSubmitBit CommandBufferUsageFlagBits = 1 << iota
	CommandBufferUsageRenderPassContinueBit
	CommandBufferUsageSimultaneousUseBit
)

type CommandBufferBeginInfo struct {
	Type            StructureType
	Next            uintptr
	Flags           CommandBufferUsageFlags
	InheritanceInfo *CommandBufferInheritanceInfo
}

type CommandBufferInheritanceInfo struct{}

type SubmitInfo struct {
	Type             StructureType
	Next             uintptr
	WaitSemaphores   []Semaphore
	WaitDstStageMask []PipelineStageFlags
	CommandBuffers   []CommandBuffer
	SignalSemaphores []Semaphore
}

type CommandPoolCreateFlagBits uint32
type CommandPoolCreateFlags = CommandPoolCreateFlagBits

const (
	CommandPoolCreateTransient CommandPoolCreateFlagBits = 1 << iota
	CommandPoolCreateResetCommandBuffer
	CommandPoolCreateProtected
)

type CommandPoolCreateInfo struct {
	Type             StructureType
	Next             uintptr
	Flags            CommandPoolCreateFlags
	QueueFamilyIndex uint32
}

type CommandPool uint64

type CommandBuffer uint64

type CommandPoolTrimFlags uint32

type CommandPoolResetFlags uint32

type CommandBufferResetFlags uint32
