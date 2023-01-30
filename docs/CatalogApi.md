# \CatalogApi

All URIs are relative to *https://serverRoot/tmf-api/productCatalogManagement/v4/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCatalog**](CatalogApi.md#CreateCatalog) | **Post** /catalog | Creates a Catalog
[**DeleteCatalog**](CatalogApi.md#DeleteCatalog) | **Delete** /catalog/{id} | Deletes a Catalog
[**ListCatalog**](CatalogApi.md#ListCatalog) | **Get** /catalog | List or find Catalog objects
[**PatchCatalog**](CatalogApi.md#PatchCatalog) | **Patch** /catalog/{id} | Updates partially a Catalog
[**RetrieveCatalog**](CatalogApi.md#RetrieveCatalog) | **Get** /catalog/{id} | Retrieves a Catalog by ID


# **CreateCatalog**
> Catalog CreateCatalog(ctx, catalog)
Creates a Catalog

This operation creates a Catalog entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **catalog** | [**CatalogCreate**](CatalogCreate.md)| The Catalog to be created | 

### Return type

[**Catalog**](Catalog.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCatalog**
> DeleteCatalog(ctx, id)
Deletes a Catalog

This operation deletes a Catalog entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the Catalog | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCatalog**
> []Catalog ListCatalog(ctx, optional)
List or find Catalog objects

This operation list or find Catalog entities

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CatalogApiListCatalogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiListCatalogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Comma-separated properties to be provided in response | 
 **offset** | **optional.Int32**| Requested index for start of resources to be provided in response | 
 **limit** | **optional.Int32**| Requested number of resources to be provided in response | 

### Return type

[**[]Catalog**](Catalog.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCatalog**
> Catalog PatchCatalog(ctx, id, catalog)
Updates partially a Catalog

This operation updates partially a Catalog entity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the Catalog | 
  **catalog** | [**CatalogUpdate**](CatalogUpdate.md)| The Catalog to be updated | 

### Return type

[**Catalog**](Catalog.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCatalog**
> Catalog RetrieveCatalog(ctx, id, optional)
Retrieves a Catalog by ID

This operation retrieves a Catalog entity. Attribute selection is enabled for all first level attributes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Identifier of the Catalog | 
 **optional** | ***CatalogApiRetrieveCatalogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiRetrieveCatalogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **optional.String**| Comma-separated properties to provide in response | 

### Return type

[**Catalog**](Catalog.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

