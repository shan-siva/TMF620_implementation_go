# \ProductSpecificationApi

All URIs are relative to *https://serverRoot/tmf-api/productCatalogManagement/v4/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProductSpecification**](ProductSpecificationApi.md#CreateProductSpecification) | **Post** /productSpecification | Creates a ProductSpecification
[**DeleteProductSpecification**](ProductSpecificationApi.md#DeleteProductSpecification) | **Delete** /productSpecification/{id} | Deletes a ProductSpecification
[**ListProductSpecification**](ProductSpecificationApi.md#ListProductSpecification) | **Get** /productSpecification | List or find ProductSpecification objects
[**PatchProductSpecification**](ProductSpecificationApi.md#PatchProductSpecification) | **Patch** /productSpecification/{id} | Updates partially a ProductSpecification
[**RetrieveProductSpecification**](ProductSpecificationApi.md#RetrieveProductSpecification) | **Get** /productSpecification/{id} | Retrieves a ProductSpecification by ID


# **CreateProductSpecification**
> ProductSpecification CreateProductSpecification(ctx, productSpecification)
Creates a ProductSpecification

This operation creates a ProductSpecification entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productSpecification** | [**ProductSpecificationCreate**](ProductSpecificationCreate.md)| The ProductSpecification to be created | 

### Return type

[**ProductSpecification**](ProductSpecification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProductSpecification**
> DeleteProductSpecification(ctx, id)
Deletes a ProductSpecification

This operation deletes a ProductSpecification entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ProductSpecification | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProductSpecification**
> []ProductSpecification ListProductSpecification(ctx, optional)
List or find ProductSpecification objects

This operation list or find ProductSpecification entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductSpecificationApiListProductSpecificationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductSpecificationApiListProductSpecificationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Comma-separated properties to be provided in response | 
 **offset** | **optional.Int32**| Requested index for start of resources to be provided in response | 
 **limit** | **optional.Int32**| Requested number of resources to be provided in response | 

### Return type

[**[]ProductSpecification**](ProductSpecification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchProductSpecification**
> ProductSpecification PatchProductSpecification(ctx, id, productSpecification)
Updates partially a ProductSpecification

This operation updates partially a ProductSpecification entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ProductSpecification | 
  **productSpecification** | [**ProductSpecificationUpdate**](ProductSpecificationUpdate.md)| The ProductSpecification to be updated | 

### Return type

[**ProductSpecification**](ProductSpecification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveProductSpecification**
> ProductSpecification RetrieveProductSpecification(ctx, id, optional)
Retrieves a ProductSpecification by ID

This operation retrieves a ProductSpecification entity. Attribute selection is enabled for all first level attributes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ProductSpecification | 
 **optional** | ***ProductSpecificationApiRetrieveProductSpecificationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductSpecificationApiRetrieveProductSpecificationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-separated properties to provide in response | 

### Return type

[**ProductSpecification**](ProductSpecification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

