package vulkan

type ClearValue struct {
	Color ClearColorValue
	//DepthStencil ClearDepthStencilValue
}

type clearValue struct {
	Color ClearColorValueFloat
	//DepthStencil ClearDepthStencilValue
}

type ClearColorValue interface {
	clearColorValue()
}

type ClearColorValueFloat [4]float32
type ClearColorValueInt [4]int32
type ClearColorValueUint [4]uint32

func (ClearColorValueFloat) clearColorValue() {}
func (ClearColorValueInt) clearColorValue()   {}
func (ClearColorValueUint) clearColorValue()  {}

type ClearDepthStencilValue struct {
	Depth   float32
	Stencil uint32
}
