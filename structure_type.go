package vulkan

//go:generate go run stringer.go -type=StructureType -output=structure_type_string.go
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
	StructureTypePhysicalDeviceSubgroupProperties                   StructureType = 1000094000
	StructureTypeBindBufferMemoryInfo                               StructureType = 1000157000
	StructureTypeBindImageMemoryInfo                                StructureType = 1000157001
	StructureTypePhysicalDevice16BitStorageFeatures                 StructureType = 1000083000
	StructureTypeMemoryDedicatedRequirements                        StructureType = 1000127000
	StructureTypeMemoryDedicatedAllocateInfo                        StructureType = 1000127001
	StructureTypeMemoryAllocateFlagsInfo                            StructureType = 1000060000
	StructureTypeDeviceGroupRenderPassBeginInfo                     StructureType = 1000060003
	StructureTypeDeviceGroupCommandBufferBeginInfo                  StructureType = 1000060004
	StructureTypeDeviceGroupSubmitInfo                              StructureType = 1000060005
	StructureTypeDeviceGroupBindSparseInfo                          StructureType = 1000060006
	StructureTypeBindBufferMemoryDeviceGroupInfo                    StructureType = 1000060013
	StructureTypeBindImageMemoryDeviceGroupInfo                     StructureType = 1000060014
	StructureTypePhysicalDeviceGroupProperties                      StructureType = 1000070000
	StructureTypeDeviceGroupDeviceCreateInfo                        StructureType = 1000070001
	StructureTypeBufferMemoryRequirementsInfo2                      StructureType = 1000146000
	StructureTypeImageMemoryRequirementsInfo2                       StructureType = 1000146001
	StructureTypeImageSparseMemoryRequirementsInfo2                 StructureType = 1000146002
	StructureTypeMemoryRequirements2                                StructureType = 1000146003
	StructureTypeSparseImageMemoryRequirements2                     StructureType = 1000146004
	StructureTypePhysicalDeviceFeatures2                            StructureType = 1000059000
	StructureTypePhysicalDeviceProperties2                          StructureType = 1000059001
	StructureTypeFormatProperties2                                  StructureType = 1000059002
	StructureTypeImageFormatProperties2                             StructureType = 1000059003
	StructureTypePhysicalDeviceImageFormatInfo2                     StructureType = 1000059004
	StructureTypeQueueFamilyProperties2                             StructureType = 1000059005
	StructureTypePhysicalDeviceMemoryProperties2                    StructureType = 1000059006
	StructureTypeSparseImageFormatProperties2                       StructureType = 1000059007
	StructureTypePhysicalDeviceSparseImageFormatInfo2               StructureType = 1000059008
	StructureTypePhysicalDevicePointClippingProperties              StructureType = 1000117000
	StructureTypeRenderPassInputAttachmentAspectCreateInfo          StructureType = 1000117001
	StructureTypeImageViewUsageCreateInfo                           StructureType = 1000117002
	StructureTypePipelineTessellationDomainOriginStateCreateInfo    StructureType = 1000117003
	StructureTypeRenderPassMultiviewCreateInfo                      StructureType = 1000053000
	StructureTypePhysicalDeviceMultiviewFeatures                    StructureType = 1000053001
	StructureTypePhysicalDeviceMultiviewProperties                  StructureType = 1000053002
	StructureTypePhysicalDeviceVariablePointersFeatures             StructureType = 1000120000
	StructureTypeProtectedSubmitInfo                                StructureType = 1000145000
	StructureTypePhysicalDeviceProtectedMemoryFeatures              StructureType = 1000145001
	StructureTypePhysicalDeviceProtectedMemoryProperties            StructureType = 1000145002
	StructureTypeDeviceQueueInfo2                                   StructureType = 1000145003
	StructureTypeSamplerYCbCrConversionCreateInfo                   StructureType = 1000156000
	StructureTypeSamplerYCbCrConversionInfo                         StructureType = 1000156001
	StructureTypeBindImagePlaneMemoryInfo                           StructureType = 1000156002
	StructureTypeImagePlaneMemoryRequirementsInfo                   StructureType = 1000156003
	StructureTypePhysicalDeviceSamplerYCbCrConversionFeatures       StructureType = 1000156004
	StructureTypeSamplerYCbCrConversionImageFormatProperties        StructureType = 1000156005
	StructureTypeDescriptorUpdateTemplateCreateInfo                 StructureType = 1000085000
	StructureTypePhysicalDeviceExternalImageFormatInfo              StructureType = 1000071000
	StructureTypeExternalImageFormatProperties                      StructureType = 1000071001
	StructureTypePhysicalDeviceExternalBufferInfo                   StructureType = 1000071002
	StructureTypeExternalBufferProperties                           StructureType = 1000071003
	StructureTypePhysicalDeviceIDProperties                         StructureType = 1000071004
	StructureTypeExternalMemoryBufferCreateInfo                     StructureType = 1000072000
	StructureTypeExternalMemoryImageCreateInfo                      StructureType = 1000072001
	StructureTypeExportMemoryAllocateInfo                           StructureType = 1000072002
	StructureTypePhysicalDeviceExternalFenceInfo                    StructureType = 1000112000
	StructureTypeExternalFenceProperties                            StructureType = 1000112001
	StructureTypeExportFenceCreateInfo                              StructureType = 1000113000
	StructureTypeExportSemaphoreCreateInfo                          StructureType = 1000077000
	StructureTypePhysicalDeviceExternalSemaphoreInfo                StructureType = 1000076000
	StructureTypeExternalSemaphoreProperties                        StructureType = 1000076001
	StructureTypePhysicalDeviceMaintenance3Properties               StructureType = 1000168000
	StructureTypeDescriptorSetLayoutSupport                         StructureType = 1000168001
	StructureTypePhysicalDeviceShaderDrawParametersFeatures         StructureType = 1000063000
	StructureTypeSwapchainCreateInfoKHR                             StructureType = 1000001000
	StructureTypePresentInfoKHR                                     StructureType = 1000001001
	StructureTypeDeviceGroupPresentCapabilitiesKHR                  StructureType = 1000060007
	StructureTypeImageSwapchainCreateInfoKHR                        StructureType = 1000060008
	StructureTypeBindImageMemorySwapchainInfoKHR                    StructureType = 1000060009
	StructureTypeAcquireNextImageInfoKHR                            StructureType = 1000060010
	StructureTypeDeviceGroupPresentInfoKHR                          StructureType = 1000060011
	StructureTypeDeviceGroupSwapchainCreateInfoKHR                  StructureType = 1000060012
	StructureTypeDisplayModeCreateInfoKHR                           StructureType = 1000002000
	StructureTypeDisplaySurfaceCreateInfoKHR                        StructureType = 1000002001
	StructureTypeDisplayPresentInfoKHR                              StructureType = 1000003000
	StructureTypeXlibSurfaceCreateInfoKHR                           StructureType = 1000004000
	StructureTypeXCBSurfaceCreateInfoKHR                            StructureType = 1000005000
	StructureTypeWaylandSurfaceCreateInfoKHR                        StructureType = 1000006000
	StructureTypeAndroidSurfaceCreateInfoKHR                        StructureType = 1000008000
	StructureTypeWin32SurfaceCreateInfoKHR                          StructureType = 1000009000
	StructureTypeDebugReportCallbackCreateInfoEXT                   StructureType = 1000011000
	StructureTypePipelineRasterizationStateRasterizationOrderAMD    StructureType = 1000018000
	StructureTypeDebugMarkerObjectNameInfoEXT                       StructureType = 1000022000
	StructureTypeDebugMarkerObjectTagInfoEXT                        StructureType = 1000022001
	StructureTypeDebugMarkerMarkerInfoEXT                           StructureType = 1000022002
	StructureTypeDedicatedAllocationImageCreateInfoNV               StructureType = 1000026000
	StructureTypeDedicatedAllocationBufferCreateInfoNV              StructureType = 1000026001
	StructureTypeDedicatedAllocationMemoryAllocateInfoNV            StructureType = 1000026002
	StructureTypePhysicalDeviceTransformFeedbackFeaturesEXT         StructureType = 1000028000
	StructureTypePhysicalDeviceTransformFeedbackPropertiesEXT       StructureType = 1000028001
	StructureTypePipelineRasterizationStateStreamCreateInfoEXT      StructureType = 1000028002
	StructureTypeImageViewHandleInfoNVX                             StructureType = 1000030000
	StructureTypeTextureLODGatherFormatPropertiesAMD                StructureType = 1000041000
	StructureTypeStreamDescriptorSurfaceCreateInfoGGP               StructureType = 1000049000
	StructureTypePhysicalDeviceCornerSampledImageFeaturesNV         StructureType = 1000050000
	StructureTypeExternalMemoryImageCreateInfoNV                    StructureType = 1000056000
	StructureTypeExportMemoryAllocateInfoNV                         StructureType = 1000056001
	StructureTypeImportMemoryWin32HandleInfoNV                      StructureType = 1000057000
	StructureTypeExportMemoryWin32HandleInfoNV                      StructureType = 1000057001
	StructureTypeWin32KeyedMutexAcquireReleaseInfoNV                StructureType = 1000058000
	StructureTypeValidationFlagsEXT                                 StructureType = 1000061000
	StructureTypeVISurfaceCreateInfoNN                              StructureType = 1000062000
	StructureTypePhysicalDeviceTextureCompressionASTCHDRFeaturesEXT StructureType = 1000066000
	StructureTypeImageViewASTCDecodeModeEXT                         StructureType = 1000067000
	StructureTypePhysicalDeviceASTCDecodeFeaturesEXT                StructureType = 1000067001
	StructureTypeImportMemoryWin32HandleInfoKHR                     StructureType = 1000073000
	StructureTypeExportMemoryWin32HandleInfoKHR                     StructureType = 1000073001
	StructureTypeMemoryWin32HandlePropertiesKHR                     StructureType = 1000073002
	StructureTypeMemoryGetWin32HandleInfoKHR                        StructureType = 1000073003
	StructureTypeImportMemoryFDInfoKHR                              StructureType = 1000074000
	StructureTypeMemoryFDPropertiesKHR                              StructureType = 1000074001
	StructureTypeMemoryGetFDInfoKHR                                 StructureType = 1000074002
	StructureTypeWin32KeyedMutexAcquireReleaseInfoKHR               StructureType = 1000075000
	StructureTypeImportSemaphoreWin32HandleInfoKHR                  StructureType = 1000078000
	StructureTypeExportSemaphoreWin32HandleInfoKHR                  StructureType = 1000078001
	StructureTypeD3D12FenceSubmitInfoKHR                            StructureType = 1000078002
	StructureTypeSemaphoreGetWin32HandleInfoKHR                     StructureType = 1000078003
	StructureTypeImportSemaphoreFDInfoKHR                           StructureType = 1000079000
	StructureTypeSemaphoreGetFDInfoKHR                              StructureType = 1000079001

	StructureTypeVISurfaceCreateInfo        StructureType = 1000062000
	StructureTypeIOSSurfaceCreateInfo       StructureType = 1000122000
	StructureTypeMacOSSurfaceCreateInfo     StructureType = 1000123000
	StructureTypeImagePipeSurfaceCreateInfo StructureType = 1000214000
)
