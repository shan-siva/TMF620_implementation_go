# \CategoryApi

All URIs are relative to *https://serverRoot/tmf-api/productCatalogManagement/v4/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCategory**](CategoryApi.md#CreateCategory) | **Post** /category | Creates a Category
[**DeleteCategory**](CategoryApi.md#DeleteCategory) | **Delete** /category/{id} | Deletes a Category
[**ListCategory**](CategoryApi.md#ListCategory) | **Get** /category | List or find Category objects
[**PatchCategory**](CategoryApi.md#PatchCategory) | **Patch** /category/{id} | Updates partially a Category
[**RetrieveCategory**](CategoryApi.md#RetrieveCategory) | **Get** /category/{id} | Retrieves a Category by ID


# **CreateCategory**
> Category CreateCategory(ctx, category)
Creates a Category

This operation creates a Category entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **category** | [**CategoryCreate**](CategoryCreate.md)| The Category to be created | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCategory**
> DeleteCategory(ctx, id)
Deletes a Category

This operation deletes a Category entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the Category | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCategory**
> []Category ListCategory(ctx, optional)
List or find Category objects

This operation list or find Category entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CategoryApiListCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoryApiListCategoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Comma-separated properties to be provided in response | 
 **offset** | **optional.Int32**| Requested index for start of resources to be provided in response | 
 **limit** | **optional.Int32**| Requested number of resources to be provided in response | 

### Return type

[**[]Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCategory**
> Category PatchCategory(ctx, id, category)
Updates partially a Category

This operation updates partially a Category entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the Category | 
  **category** | [**CategoryUpdate**](CategoryUpdate.md)| The Category to be updated | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCategory**
> Category RetrieveCategory(ctx, id, optional)
Retrieves a Category by ID

This operation retrieves a Category entity. Attribute selection is enabled for all first level attributes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the Category | 
 **optional** | ***CategoryApiRetrieveCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoryApiRetrieveCategoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-separated properties to provide in response | 

### Return type

[**Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

