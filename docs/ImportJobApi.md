# \ImportJobApi

All URIs are relative to *https://serverRoot/tmf-api/productCatalogManagement/v4/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImportJob**](ImportJobApi.md#CreateImportJob) | **Post** /importJob | Creates a ImportJob
[**DeleteImportJob**](ImportJobApi.md#DeleteImportJob) | **Delete** /importJob/{id} | Deletes a ImportJob
[**ListImportJob**](ImportJobApi.md#ListImportJob) | **Get** /importJob | List or find ImportJob objects
[**RetrieveImportJob**](ImportJobApi.md#RetrieveImportJob) | **Get** /importJob/{id} | Retrieves a ImportJob by ID


# **CreateImportJob**
> ImportJob CreateImportJob(ctx, importJob)
Creates a ImportJob

This operation creates a ImportJob entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **importJob** | [**ImportJobCreate**](ImportJobCreate.md)| The ImportJob to be created | 

### Return type

[**ImportJob**](ImportJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteImportJob**
> DeleteImportJob(ctx, id)
Deletes a ImportJob

This operation deletes a ImportJob entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ImportJob | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListImportJob**
> []ImportJob ListImportJob(ctx, optional)
List or find ImportJob objects

This operation list or find ImportJob entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImportJobApiListImportJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportJobApiListImportJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Comma-separated properties to be provided in response | 
 **offset** | **optional.Int32**| Requested index for start of resources to be provided in response | 
 **limit** | **optional.Int32**| Requested number of resources to be provided in response | 

### Return type

[**[]ImportJob**](ImportJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveImportJob**
> ImportJob RetrieveImportJob(ctx, id, optional)
Retrieves a ImportJob by ID

This operation retrieves a ImportJob entity. Attribute selection is enabled for all first level attributes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ImportJob | 
 **optional** | ***ImportJobApiRetrieveImportJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImportJobApiRetrieveImportJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-separated properties to provide in response | 

### Return type

[**ImportJob**](ImportJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

