# \ExportJobApi

All URIs are relative to *https://serverRoot/tmf-api/productCatalogManagement/v4/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExportJob**](ExportJobApi.md#CreateExportJob) | **Post** /exportJob | Creates a ExportJob
[**DeleteExportJob**](ExportJobApi.md#DeleteExportJob) | **Delete** /exportJob/{id} | Deletes a ExportJob
[**ListExportJob**](ExportJobApi.md#ListExportJob) | **Get** /exportJob | List or find ExportJob objects
[**RetrieveExportJob**](ExportJobApi.md#RetrieveExportJob) | **Get** /exportJob/{id} | Retrieves a ExportJob by ID


# **CreateExportJob**
> ExportJob CreateExportJob(ctx, exportJob)
Creates a ExportJob

This operation creates a ExportJob entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **exportJob** | [**ExportJobCreate**](ExportJobCreate.md)| The ExportJob to be created | 

### Return type

[**ExportJob**](ExportJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExportJob**
> DeleteExportJob(ctx, id)
Deletes a ExportJob

This operation deletes a ExportJob entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ExportJob | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExportJob**
> []ExportJob ListExportJob(ctx, optional)
List or find ExportJob objects

This operation list or find ExportJob entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExportJobApiListExportJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportJobApiListExportJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Comma-separated properties to be provided in response | 
 **offset** | **optional.Int32**| Requested index for start of resources to be provided in response | 
 **limit** | **optional.Int32**| Requested number of resources to be provided in response | 

### Return type

[**[]ExportJob**](ExportJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveExportJob**
> ExportJob RetrieveExportJob(ctx, id, optional)
Retrieves a ExportJob by ID

This operation retrieves a ExportJob entity. Attribute selection is enabled for all first level attributes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the ExportJob | 
 **optional** | ***ExportJobApiRetrieveExportJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExportJobApiRetrieveExportJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-separated properties to provide in response | 

### Return type

[**ExportJob**](ExportJob.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

