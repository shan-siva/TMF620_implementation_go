# \ProductOfferingApi

All URIs are relative to *https://serverRoot/tmf-api/productCatalogManagement/v4/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProductOffering**](ProductOfferingApi.md#CreateProductOffering) | **Post** /productOffering | Creates a ProductOffering
[**DeleteProductOffering**](ProductOfferingApi.md#DeleteProductOffering) | **Delete** /productOffering/{id} | Deletes a ProductOffering
[**ListProductOffering**](ProductOfferingApi.md#ListProductOffering) | **Get** /productOffering | List or find ProductOffering objects
[**PatchProductOffering**](ProductOfferingApi.md#PatchProductOffering) | **Patch** /productOffering/{id} | Updates partially a ProductOffering
[**RetrieveProductOffering**](ProductOfferingApi.md#RetrieveProductOffering) | **Get** /productOffering/{id} | Retrieves a ProductOffering by ID


# **CreateProductOffering**
> ProductOffering CreateProductOffering(ctx, productOffering)
Creates a ProductOffering

This operation creates a ProductOffering entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productOffering** | [**ProductOfferingCreate**](ProductOfferingCreate.md)| The ProductOffering to be created | 

### Return type

[**ProductOffering**](ProductOffering.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProductOffering**
> DeleteProductOffering(ctx, id)
Deletes a ProductOffering

This operation deletes a ProductOffering entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ProductOffering | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProductOffering**
> []ProductOffering ListProductOffering(ctx, optional)
List or find ProductOffering objects

This operation list or find ProductOffering entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductOfferingApiListProductOfferingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductOfferingApiListProductOfferingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Comma-separated properties to be provided in response | 
 **offset** | **optional.Int32**| Requested index for start of resources to be provided in response | 
 **limit** | **optional.Int32**| Requested number of resources to be provided in response | 

### Return type

[**[]ProductOffering**](ProductOffering.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchProductOffering**
> ProductOffering PatchProductOffering(ctx, id, productOffering)
Updates partially a ProductOffering

This operation updates partially a ProductOffering entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ProductOffering | 
  **productOffering** | [**ProductOfferingUpdate**](ProductOfferingUpdate.md)| The ProductOffering to be updated | 

### Return type

[**ProductOffering**](ProductOffering.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveProductOffering**
> ProductOffering RetrieveProductOffering(ctx, id, optional)
Retrieves a ProductOffering by ID

This operation retrieves a ProductOffering entity. Attribute selection is enabled for all first level attributes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ProductOffering | 
 **optional** | ***ProductOfferingApiRetrieveProductOfferingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductOfferingApiRetrieveProductOfferingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-separated properties to provide in response | 

### Return type

[**ProductOffering**](ProductOffering.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

