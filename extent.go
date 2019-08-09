package vulkan

// Extent2D - Structure specifying a two-dimensional extent
type Extent2D struct {
	// Width is the width of the extent.
	Width uint32

	// Height is the height of the extent.
	Height uint32
}

// Extent3D - Structure specifying a three-dimensional extent
type Extent3D struct {
	// Width  is the width of the extent.
	Width uint32

	// Height is the height of the extent.
	Height uint32

	// Depth is the depth of the extent.
	Depth uint32
}
