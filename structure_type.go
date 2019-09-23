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
	StructureTypePhysicalDeviceSubgroupProperties                      StructureType = 1000094000
	StructureTypeBindBufferMemoryInfo                                  StructureType = 1000157000
	StructureTypeBindImageMemoryInfo                                   StructureType = 1000157001
	StructureTypePhysicalDevice16BitStorageFeatures                    StructureType = 1000083000
	StructureTypeMemoryDedicatedRequirements                           StructureType = 1000127000
	StructureTypeMemoryDedicatedAllocateInfo                           StructureType = 1000127001
	StructureTypeMemoryAllocateFlagsInfo                               StructureType = 1000060000
	StructureTypeDeviceGroupRenderPassBeginInfo                        StructureType = 1000060003
	StructureTypeDeviceGroupCommandBufferBeginInfo                     StructureType = 1000060004
	StructureTypeDeviceGroupSubmitInfo                                 StructureType = 1000060005
	StructureTypeDeviceGroupBindSparseInfo                             StructureType = 1000060006
	StructureTypeBindBufferMemoryDeviceGroupInfo                       StructureType = 1000060013
	StructureTypeBindImageMemoryDeviceGroupInfo                        StructureType = 1000060014
	StructureTypePhysicalDeviceGroupProperties                         StructureType = 1000070000
	StructureTypeDeviceGroupDeviceCreateInfo                           StructureType = 1000070001
	StructureTypeBufferMemoryRequirementsInfo2                         StructureType = 1000146000
	StructureTypeImageMemoryRequirementsInfo2                          StructureType = 1000146001
	StructureTypeImageSparseMemoryRequirementsInfo2                    StructureType = 1000146002
	StructureTypeMemoryRequirements2                                   StructureType = 1000146003
	StructureTypeSparseImageMemoryRequirements2                        StructureType = 1000146004
	StructureTypePhysicalDeviceFeatures2                               StructureType = 1000059000
	StructureTypePhysicalDeviceProperties2                             StructureType = 1000059001
	StructureTypeFormatProperties2                                     StructureType = 1000059002
	StructureTypeImageFormatProperties2                                StructureType = 1000059003
	StructureTypePhysicalDeviceImageFormatInfo2                        StructureType = 1000059004
	StructureTypeQueueFamilyProperties2                                StructureType = 1000059005
	StructureTypePhysicalDeviceMemoryProperties2                       StructureType = 1000059006
	StructureTypeSparseImageFormatProperties2                          StructureType = 1000059007
	StructureTypePhysicalDeviceSparseImageFormatInfo2                  StructureType = 1000059008
	StructureTypePhysicalDevicePointClippingProperties                 StructureType = 1000117000
	StructureTypeRenderPassInputAttachmentAspectCreateInfo             StructureType = 1000117001
	StructureTypeImageViewUsageCreateInfo                              StructureType = 1000117002
	StructureTypePipelineTessellationDomainOriginStateCreateInfo       StructureType = 1000117003
	StructureTypeRenderPassMultiviewCreateInfo                         StructureType = 1000053000
	StructureTypePhysicalDeviceMultiviewFeatures                       StructureType = 1000053001
	StructureTypePhysicalDeviceMultiviewProperties                     StructureType = 1000053002
	StructureTypePhysicalDeviceVariablePointersFeatures                StructureType = 1000120000
	StructureTypeProtectedSubmitInfo                                   StructureType = 1000145000
	StructureTypePhysicalDeviceProtectedMemoryFeatures                 StructureType = 1000145001
	StructureTypePhysicalDeviceProtectedMemoryProperties               StructureType = 1000145002
	StructureTypeDeviceQueueInfo2                                      StructureType = 1000145003
	StructureTypeSamplerYCbCrConversionCreateInfo                      StructureType = 1000156000
	StructureTypeSamplerYCbCrConversionInfo                            StructureType = 1000156001
	StructureTypeBindImagePlaneMemoryInfo                              StructureType = 1000156002
	StructureTypeImagePlaneMemoryRequirementsInfo                      StructureType = 1000156003
	StructureTypePhysicalDeviceSamplerYCbCrConversionFeatures          StructureType = 1000156004
	StructureTypeSamplerYCbCrConversionImageFormatProperties           StructureType = 1000156005
	StructureTypeDescriptorUpdateTemplateCreateInfo                    StructureType = 1000085000
	StructureTypePhysicalDeviceExternalImageFormatInfo                 StructureType = 1000071000
	StructureTypeExternalImageFormatProperties                         StructureType = 1000071001
	StructureTypePhysicalDeviceExternalBufferInfo                      StructureType = 1000071002
	StructureTypeExternalBufferProperties                              StructureType = 1000071003
	StructureTypePhysicalDeviceIDProperties                            StructureType = 1000071004
	StructureTypeExternalMemoryBufferCreateInfo                        StructureType = 1000072000
	StructureTypeExternalMemoryImageCreateInfo                         StructureType = 1000072001
	StructureTypeExportMemoryAllocateInfo                              StructureType = 1000072002
	StructureTypePhysicalDeviceExternalFenceInfo                       StructureType = 1000112000
	StructureTypeExternalFenceProperties                               StructureType = 1000112001
	StructureTypeExportFenceCreateInfo                                 StructureType = 1000113000
	StructureTypeExportSemaphoreCreateInfo                             StructureType = 1000077000
	StructureTypePhysicalDeviceExternalSemaphoreInfo                   StructureType = 1000076000
	StructureTypeExternalSemaphoreProperties                           StructureType = 1000076001
	StructureTypePhysicalDeviceMaintenance3Properties                  StructureType = 1000168000
	StructureTypeDescriptorSetLayoutSupport                            StructureType = 1000168001
	StructureTypePhysicalDeviceShaderDrawParametersFeatures            StructureType = 1000063000
	StructureTypeSwapchainCreateInfoKHR                                StructureType = 1000001000
	StructureTypePresentInfoKHR                                        StructureType = 1000001001
	StructureTypeDeviceGroupPresentCapabilitiesKHR                     StructureType = 1000060007
	StructureTypeImageSwapchainCreateInfoKHR                           StructureType = 1000060008
	StructureTypeBindImageMemorySwapchainInfoKHR                       StructureType = 1000060009
	StructureTypeAcquireNextImageInfoKHR                               StructureType = 1000060010
	StructureTypeDeviceGroupPresentInfoKHR                             StructureType = 1000060011
	StructureTypeDeviceGroupSwapchainCreateInfoKHR                     StructureType = 1000060012
	StructureTypeDisplayModeCreateInfoKHR                              StructureType = 1000002000
	StructureTypeDisplaySurfaceCreateInfoKHR                           StructureType = 1000002001
	StructureTypeDisplayPresentInfoKHR                                 StructureType = 1000003000
	StructureTypeXlibSurfaceCreateInfoKHR                              StructureType = 1000004000
	StructureTypeXCBSurfaceCreateInfoKHR                               StructureType = 1000005000
	StructureTypeWaylandSurfaceCreateInfoKHR                           StructureType = 1000006000
	StructureTypeAndroidSurfaceCreateInfoKHR                           StructureType = 1000008000
	StructureTypeWin32SurfaceCreateInfoKHR                             StructureType = 1000009000
	StructureTypeDebugReportCallbackCreateInfoEXT                      StructureType = 1000011000
	StructureTypePipelineRasterizationStateRasterizationOrderAMD       StructureType = 1000018000
	StructureTypeDebugMarkerObjectNameInfoEXT                          StructureType = 1000022000
	StructureTypeDebugMarkerObjectTagInfoEXT                           StructureType = 1000022001
	StructureTypeDebugMarkerMarkerInfoEXT                              StructureType = 1000022002
	StructureTypeDedicatedAllocationImageCreateInfoNV                  StructureType = 1000026000
	StructureTypeDedicatedAllocationBufferCreateInfoNV                 StructureType = 1000026001
	StructureTypeDedicatedAllocationMemoryAllocateInfoNV               StructureType = 1000026002
	StructureTypePhysicalDeviceTransformFeedbackFeaturesEXT            StructureType = 1000028000
	StructureTypePhysicalDeviceTransformFeedbackPropertiesEXT          StructureType = 1000028001
	StructureTypePipelineRasterizationStateStreamCreateInfoEXT         StructureType = 1000028002
	StructureTypeImageViewHandleInfoNVX                                StructureType = 1000030000
	StructureTypeTextureLODGatherFormatPropertiesAMD                   StructureType = 1000041000
	StructureTypeStreamDescriptorSurfaceCreateInfoGGP                  StructureType = 1000049000
	StructureTypePhysicalDeviceCornerSampledImageFeaturesNV            StructureType = 1000050000
	StructureTypeExternalMemoryImageCreateInfoNV                       StructureType = 1000056000
	StructureTypeExportMemoryAllocateInfoNV                            StructureType = 1000056001
	StructureTypeImportMemoryWin32HandleInfoNV                         StructureType = 1000057000
	StructureTypeExportMemoryWin32HandleInfoNV                         StructureType = 1000057001
	StructureTypeWin32KeyedMutexAcquireReleaseInfoNV                   StructureType = 1000058000
	StructureTypeValidationFlagsEXT                                    StructureType = 1000061000
	StructureTypeVISurfaceCreateInfoNN                                 StructureType = 1000062000
	StructureTypePhysicalDeviceTextureCompressionASTCHDRFeaturesEXT    StructureType = 1000066000
	StructureTypeImageViewASTCDecodeModeEXT                            StructureType = 1000067000
	StructureTypePhysicalDeviceASTCDecodeFeaturesEXT                   StructureType = 1000067001
	StructureTypeImportMemoryWin32HandleInfoKHR                        StructureType = 1000073000
	StructureTypeExportMemoryWin32HandleInfoKHR                        StructureType = 1000073001
	StructureTypeMemoryWin32HandlePropertiesKHR                        StructureType = 1000073002
	StructureTypeMemoryGetWin32HandleInfoKHR                           StructureType = 1000073003
	StructureTypeImportMemoryFDInfoKHR                                 StructureType = 1000074000
	StructureTypeMemoryFDPropertiesKHR                                 StructureType = 1000074001
	StructureTypeMemoryGetFDInfoKHR                                    StructureType = 1000074002
	StructureTypeWin32KeyedMutexAcquireReleaseInfoKHR                  StructureType = 1000075000
	StructureTypeImportSemaphoreWin32HandleInfoKHR                     StructureType = 1000078000
	StructureTypeExportSemaphoreWin32HandleInfoKHR                     StructureType = 1000078001
	StructureTypeD3D12FenceSubmitInfoKHR                               StructureType = 1000078002
	StructureTypeSemaphoreGetWin32HandleInfoKHR                        StructureType = 1000078003
	StructureTypeImportSemaphoreFDInfoKHR                              StructureType = 1000079000
	StructureTypeSemaphoreGetFDInfoKHR                                 StructureType = 1000079001
	StructureTypePhysicalDevicePushDescriptorPropertiesKHR             StructureType = 1000080000
	StructureTypeCommandBufferInheritanceConditionalRenderingInfoEXT   StructureType = 1000081000
	StructureTypePhysicalDeviceConditionalRenderingFeaturesEXT         StructureType = 1000081001
	StructureTypeConditionalRenderingBeginInfoEXT                      StructureType = 1000081002
	StructureTypePhysicalDeviceShaderFloat16Int8FeaturesEXT            StructureType = 1000082000
	StructureTypePresentRegionsKHR                                     StructureType = 1000084000
	StructureTypeObjectTableCreateInfoNVX                              StructureType = 1000086000
	StructureTypeIndirectCommandsLayoutCreateInfoNVX                   StructureType = 1000086001
	StructureTypeCMDProcessCommandsInfoNVX                             StructureType = 1000086002
	StructureTypeCMDReserveSpaceForCommandsInfoNVX                     StructureType = 1000086003
	StructureTypeDeviceGeneratedCommandsLimitsNVX                      StructureType = 1000086004
	StructureTypeDeviceGeneratedCommandsFeaturesNVX                    StructureType = 1000086005
	StructureTypePipelineViewportWScalingStateCreateInfoNV             StructureType = 1000087000
	StructureTypeSurfaceCapabilities2EXT                               StructureType = 1000090000
	StructureTypeDisplayPowerInfoEXT                                   StructureType = 1000091000
	StructureTypeDeviceEventInfoEXT                                    StructureType = 1000091001
	StructureTypeDisplayEventInfoEXT                                   StructureType = 1000091002
	StructureTypeSwapchainCounterCreateInfoEXT                         StructureType = 1000091003
	StructureTypePresentTimesInfoGoogle                                StructureType = 1000092000
	StructureTypePhysicalDeviceMultiviewPerViewAttributesPropertiesNVX StructureType = 1000097000
	StructureTypePipelineViewportSwizzleStateCreateInfoNV              StructureType = 1000098000
	StructureTypePhysicalDeviceDiscardRectanglePropertiesEXT           StructureType = 1000099000
	StructureTypePipelineDiscardRectangleStateCreateInfoEXT            StructureType = 1000099001
	StructureTypePhysicalDeviceConservativeRasterizationPropertiesEXT  StructureType = 1000101000
	StructureTypePipelineRasterizationConservativeStateCreateInfoEXT   StructureType = 1000101001
	StructureTypePhysicalDeviceDepthClipEnableFeaturesEXT              StructureType = 1000102000
	StructureTypePipelineRasterizationDepthClipStateCreateInfoEXT      StructureType = 1000102001
	StructureTypeHDRMetadataEXT                                        StructureType = 1000105000
	StructureTypePhysicalDeviceImagelessFramebufferFeaturesKHR         StructureType = 1000108000
	StructureTypeFramebufferAttachmentsCreateInfoKHR                   StructureType = 1000108001
	StructureTypeFramebufferAttachmentImageInfoKHR                     StructureType = 1000108002
	StructureTypeRenderPassAttachmentBeginInfoKHR                      StructureType = 1000108003
	StructureTypeAttachmentDescription2KHR                             StructureType = 1000109000
	StructureTypeAttachmentReference2KHR                               StructureType = 1000109001
	StructureTypeSubpassDescription2KHR                                StructureType = 1000109002
	StructureTypeSubpassDependency2KHR                                 StructureType = 1000109003
	StructureTypeRenderPassCreateInfo2KHR                              StructureType = 1000109004
	StructureTypeSubpassBeginInfoKHR                                   StructureType = 1000109005
	StructureTypeSubpassEndInfoKHR                                     StructureType = 1000109006
	StructureTypeSharedPresentSurfaceCapabilitiesKHR                   StructureType = 1000111000
	StructureTypeImportFenceWin32HandleInfoKHR                         StructureType = 1000114000
	StructureTypeExportFenceWin32HandleInfoKHR                         StructureType = 1000114001
	StructureTypeFenceGetWin32HandleInfoKHR                            StructureType = 1000114002
	StructureTypeImportFenceFDInfoKHR                                  StructureType = 1000115000
	StructureTypeFenceGetFDInfoKHR                                     StructureType = 1000115001
	StructureTypePhysicalDeviceSurfaceInfo2KHR                         StructureType = 1000119000
	StructureTypeSurfaceCapabilities2KHR                               StructureType = 1000119001
	StructureTypeSurfaceFormat2KHR                                     StructureType = 1000119002
	StructureTypeDisplayProperties2KHR                                 StructureType = 1000121000
	StructureTypeDisplayPlaneProperties2KHR                            StructureType = 1000121001
	StructureTypeDisplayModeProperties2KHR                             StructureType = 1000121002
	StructureTypeDisplayPlaneInfo2KHR                                  StructureType = 1000121003
	StructureTypeDisplayPlaneCapabilities2KHR                          StructureType = 1000121004
	StructureTypeIOSSurfaceCreateInfoMVK                               StructureType = 1000122000
	StructureTypeMacOSSurfaceCreateInfoMVK                             StructureType = 1000123000
	StructureTypeDebugUtilsObjectNameInfoEXT                           StructureType = 1000128000
	StructureTypeDebugUtilsObjectTagInfoEXT                            StructureType = 1000128001
	StructureTypeDebugUtilsLabelEXT                                    StructureType = 1000128002
	StructureTypeDebugUtilsMessengerCallbackDataEXT                    StructureType = 1000128003
	StructureTypeDebugUtilsMessengerCreateInfoEXT                      StructureType = 1000128004
	StructureTypeAndroidHardwareBufferUsageAndroid                     StructureType = 1000129000
	StructureTypeAndroidHardwareBufferPropertiesAndroid                StructureType = 1000129001
	StructureTypeAndroidHardwareBufferFormatPropertiesAndroid          StructureType = 1000129002
	StructureTypeImportAndroidHardwareBufferInfoAndroid                StructureType = 1000129003
	StructureTypeMemoryGetAndroidHardwareBufferInfoAndroid             StructureType = 1000129004
	StructureTypeExternalFormatAndroid                                 StructureType = 1000129005
	StructureTypePhysicalDeviceSamplerFilterMinMaxPropertiesEXT        StructureType = 1000130000
	StructureTypeSamplerReductionModeCreateInfoEXT                     StructureType = 1000130001
	StructureTypePhysicalDeviceInlineUniformBlockFeaturesEXT           StructureType = 1000138000
	StructureTypePhysicalDeviceInlineUniformBlockPropertiesEXT         StructureType = 1000138001
	StructureTypeWriteDescriptorSetInlineUniformBlockEXT               StructureType = 1000138002
	StructureTypeDescriptorPoolInlineUniformBlockCreateInfoEXT         StructureType = 1000138003
	StructureTypeSampleLocationsInfoEXT                                StructureType = 1000143000
	StructureTypeRenderPassSampleLocationsBeginInfoEXT                 StructureType = 1000143001
	StructureTypePipelineSampleLocationsStateCreateInfoEXT             StructureType = 1000143002
	StructureTypePhysicalDeviceSampleLocationsPropertiesEXT            StructureType = 1000143003
	StructureTypeMultisamplePropertiesEXT                              StructureType = 1000143004
	StructureTypeImageFormatListCreateInfoKHR                          StructureType = 1000147000
	StructureTypePhysicalDeviceBlendOperationAdvancedFeaturesEXT       StructureType = 1000148000
	StructureTypePhysicalDeviceBlendOperationAdvancedPropertiesEXT     StructureType = 1000148001
	StructureTypePipelineColorBlendAdvancedStateCreateInfoEXT          StructureType = 1000148002
	StructureTypePipelineCoverageToColorStateCreateInfoNV              StructureType = 1000149000
	StructureTypePipelineCoverageModulationStateCreateInfoNV           StructureType = 1000152000
	StructureTypePhysicalDeviceShaderSMBuiltinsFeaturesNV              StructureType = 1000154000
	StructureTypePhysicalDeviceShaderSMBuiltinsPropertiesNV            StructureType = 1000154001
	StructureTypeDRMFormatModifierPropertiesListEXT                    StructureType = 1000158000
	StructureTypeDRMFormatModifierPropertiesEXT                        StructureType = 1000158001
	StructureTypePhysicalDeviceImageDRMFormatModifierInfoEXT           StructureType = 1000158002
	StructureTypeImageDRMFormatModifierListCreateInfoEXT               StructureType = 1000158003
	StructureTypeImageDRMFormatModifierExplicitCreateInfoEXT           StructureType = 1000158004
	StructureTypeImageDRMFormatModifierPropertiesEXT                   StructureType = 1000158005
	StructureTypeValidationCacheCreateInfoEXT                          StructureType = 1000160000
	StructureTypeShaderModuleValidationCacheCreateInfoEXT              StructureType = 1000160001
	StructureTypeDescriptorSetLayoutBindingFlagsCreateInfoExt          StructureType = 100161000

	StructureTypeImagePipeSurfaceCreateInfo StructureType = 1000214000
)
