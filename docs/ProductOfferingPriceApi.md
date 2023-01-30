# \ProductOfferingPriceApi

All URIs are relative to *https://serverRoot/tmf-api/productCatalogManagement/v4/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProductOfferingPrice**](ProductOfferingPriceApi.md#CreateProductOfferingPrice) | **Post** /productOfferingPrice | Creates a ProductOfferingPrice
[**DeleteProductOfferingPrice**](ProductOfferingPriceApi.md#DeleteProductOfferingPrice) | **Delete** /productOfferingPrice/{id} | Deletes a ProductOfferingPrice
[**ListProductOfferingPrice**](ProductOfferingPriceApi.md#ListProductOfferingPrice) | **Get** /productOfferingPrice | List or find ProductOfferingPrice objects
[**PatchProductOfferingPrice**](ProductOfferingPriceApi.md#PatchProductOfferingPrice) | **Patch** /productOfferingPrice/{id} | Updates partially a ProductOfferingPrice
[**RetrieveProductOfferingPrice**](ProductOfferingPriceApi.md#RetrieveProductOfferingPrice) | **Get** /productOfferingPrice/{id} | Retrieves a ProductOfferingPrice by ID


# **CreateProductOfferingPrice**
> ProductOfferingPrice CreateProductOfferingPrice(ctx, productOfferingPrice)
Creates a ProductOfferingPrice

This operation creates a ProductOfferingPrice entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **productOfferingPrice** | [**ProductOfferingPriceCreate**](ProductOfferingPriceCreate.md)| The ProductOfferingPrice to be created | 

### Return type

[**ProductOfferingPrice**](ProductOfferingPrice.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProductOfferingPrice**
> DeleteProductOfferingPrice(ctx, id)
Deletes a ProductOfferingPrice

This operation deletes a ProductOfferingPrice entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ProductOfferingPrice | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProductOfferingPrice**
> []ProductOfferingPrice ListProductOfferingPrice(ctx, optional)
List or find ProductOfferingPrice objects

This operation list or find ProductOfferingPrice entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductOfferingPriceApiListProductOfferingPriceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductOfferingPriceApiListProductOfferingPriceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Comma-separated properties to be provided in response | 
 **offset** | **optional.Int32**| Requested index for start of resources to be provided in response | 
 **limit** | **optional.Int32**| Requested number of resources to be provided in response | 

### Return type

[**[]ProductOfferingPrice**](ProductOfferingPrice.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchProductOfferingPrice**
> ProductOfferingPrice PatchProductOfferingPrice(ctx, id, productOfferingPrice)
Updates partially a ProductOfferingPrice

This operation updates partially a ProductOfferingPrice entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ProductOfferingPrice | 
  **productOfferingPrice** | [**ProductOfferingPriceUpdate**](ProductOfferingPriceUpdate.md)| The ProductOfferingPrice to be updated | 

### Return type

[**ProductOfferingPrice**](ProductOfferingPrice.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveProductOfferingPrice**
> ProductOfferingPrice RetrieveProductOfferingPrice(ctx, id, optional)
Retrieves a ProductOfferingPrice by ID

This operation retrieves a ProductOfferingPrice entity. Attribute selection is enabled for all first level attributes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ProductOfferingPrice | 
 **optional** | ***ProductOfferingPriceApiRetrieveProductOfferingPriceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductOfferingPriceApiRetrieveProductOfferingPriceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-separated properties to provide in response | 

### Return type

[**ProductOfferingPrice**](ProductOfferingPrice.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

