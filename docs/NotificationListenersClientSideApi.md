# \NotificationListenersClientSideApi

All URIs are relative to *https://serverRoot/tmf-api/productCatalogManagement/v4/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListenToCatalogAttributeValueChangeEvent**](NotificationListenersClientSideApi.md#ListenToCatalogAttributeValueChangeEvent) | **Post** /listener/catalogAttributeValueChangeEvent | Client listener for entity CatalogAttributeValueChangeEvent
[**ListenToCatalogBatchEvent**](NotificationListenersClientSideApi.md#ListenToCatalogBatchEvent) | **Post** /listener/catalogBatchEvent | Client listener for entity CatalogBatchEvent
[**ListenToCatalogCreateEvent**](NotificationListenersClientSideApi.md#ListenToCatalogCreateEvent) | **Post** /listener/catalogCreateEvent | Client listener for entity CatalogCreateEvent
[**ListenToCatalogDeleteEvent**](NotificationListenersClientSideApi.md#ListenToCatalogDeleteEvent) | **Post** /listener/catalogDeleteEvent | Client listener for entity CatalogDeleteEvent
[**ListenToCatalogStateChangeEvent**](NotificationListenersClientSideApi.md#ListenToCatalogStateChangeEvent) | **Post** /listener/catalogStateChangeEvent | Client listener for entity CatalogStateChangeEvent
[**ListenToCategoryAttributeValueChangeEvent**](NotificationListenersClientSideApi.md#ListenToCategoryAttributeValueChangeEvent) | **Post** /listener/categoryAttributeValueChangeEvent | Client listener for entity CategoryAttributeValueChangeEvent
[**ListenToCategoryCreateEvent**](NotificationListenersClientSideApi.md#ListenToCategoryCreateEvent) | **Post** /listener/categoryCreateEvent | Client listener for entity CategoryCreateEvent
[**ListenToCategoryDeleteEvent**](NotificationListenersClientSideApi.md#ListenToCategoryDeleteEvent) | **Post** /listener/categoryDeleteEvent | Client listener for entity CategoryDeleteEvent
[**ListenToCategoryStateChangeEvent**](NotificationListenersClientSideApi.md#ListenToCategoryStateChangeEvent) | **Post** /listener/categoryStateChangeEvent | Client listener for entity CategoryStateChangeEvent
[**ListenToProductOfferingAttributeValueChangeEvent**](NotificationListenersClientSideApi.md#ListenToProductOfferingAttributeValueChangeEvent) | **Post** /listener/productOfferingAttributeValueChangeEvent | Client listener for entity ProductOfferingAttributeValueChangeEvent
[**ListenToProductOfferingCreateEvent**](NotificationListenersClientSideApi.md#ListenToProductOfferingCreateEvent) | **Post** /listener/productOfferingCreateEvent | Client listener for entity ProductOfferingCreateEvent
[**ListenToProductOfferingDeleteEvent**](NotificationListenersClientSideApi.md#ListenToProductOfferingDeleteEvent) | **Post** /listener/productOfferingDeleteEvent | Client listener for entity ProductOfferingDeleteEvent
[**ListenToProductOfferingPriceAttributeValueChangeEvent**](NotificationListenersClientSideApi.md#ListenToProductOfferingPriceAttributeValueChangeEvent) | **Post** /listener/productOfferingPriceAttributeValueChangeEvent | Client listener for entity ProductOfferingPriceAttributeValueChangeEvent
[**ListenToProductOfferingPriceCreateEvent**](NotificationListenersClientSideApi.md#ListenToProductOfferingPriceCreateEvent) | **Post** /listener/productOfferingPriceCreateEvent | Client listener for entity ProductOfferingPriceCreateEvent
[**ListenToProductOfferingPriceDeleteEvent**](NotificationListenersClientSideApi.md#ListenToProductOfferingPriceDeleteEvent) | **Post** /listener/productOfferingPriceDeleteEvent | Client listener for entity ProductOfferingPriceDeleteEvent
[**ListenToProductOfferingPriceStateChangeEvent**](NotificationListenersClientSideApi.md#ListenToProductOfferingPriceStateChangeEvent) | **Post** /listener/productOfferingPriceStateChangeEvent | Client listener for entity ProductOfferingPriceStateChangeEvent
[**ListenToProductOfferingStateChangeEvent**](NotificationListenersClientSideApi.md#ListenToProductOfferingStateChangeEvent) | **Post** /listener/productOfferingStateChangeEvent | Client listener for entity ProductOfferingStateChangeEvent
[**ListenToProductSpecificationAttributeValueChangeEvent**](NotificationListenersClientSideApi.md#ListenToProductSpecificationAttributeValueChangeEvent) | **Post** /listener/productSpecificationAttributeValueChangeEvent | Client listener for entity ProductSpecificationAttributeValueChangeEvent
[**ListenToProductSpecificationCreateEvent**](NotificationListenersClientSideApi.md#ListenToProductSpecificationCreateEvent) | **Post** /listener/productSpecificationCreateEvent | Client listener for entity ProductSpecificationCreateEvent
[**ListenToProductSpecificationDeleteEvent**](NotificationListenersClientSideApi.md#ListenToProductSpecificationDeleteEvent) | **Post** /listener/productSpecificationDeleteEvent | Client listener for entity ProductSpecificationDeleteEvent
[**ListenToProductSpecificationStateChangeEvent**](NotificationListenersClientSideApi.md#ListenToProductSpecificationStateChangeEvent) | **Post** /listener/productSpecificationStateChangeEvent | Client listener for entity ProductSpecificationStateChangeEvent


# **ListenToCatalogAttributeValueChangeEvent**
> EventSubscription ListenToCatalogAttributeValueChangeEvent(ctx, data)
Client listener for entity CatalogAttributeValueChangeEvent

Example of a client listener for receiving the notification CatalogAttributeValueChangeEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CatalogAttributeValueChangeEvent**](CatalogAttributeValueChangeEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToCatalogBatchEvent**
> EventSubscription ListenToCatalogBatchEvent(ctx, data)
Client listener for entity CatalogBatchEvent

Example of a client listener for receiving the notification CatalogBatchEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CatalogBatchEvent**](CatalogBatchEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToCatalogCreateEvent**
> EventSubscription ListenToCatalogCreateEvent(ctx, data)
Client listener for entity CatalogCreateEvent

Example of a client listener for receiving the notification CatalogCreateEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CatalogCreateEvent**](CatalogCreateEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToCatalogDeleteEvent**
> EventSubscription ListenToCatalogDeleteEvent(ctx, data)
Client listener for entity CatalogDeleteEvent

Example of a client listener for receiving the notification CatalogDeleteEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CatalogDeleteEvent**](CatalogDeleteEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToCatalogStateChangeEvent**
> EventSubscription ListenToCatalogStateChangeEvent(ctx, data)
Client listener for entity CatalogStateChangeEvent

Example of a client listener for receiving the notification CatalogStateChangeEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CatalogStateChangeEvent**](CatalogStateChangeEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToCategoryAttributeValueChangeEvent**
> EventSubscription ListenToCategoryAttributeValueChangeEvent(ctx, data)
Client listener for entity CategoryAttributeValueChangeEvent

Example of a client listener for receiving the notification CategoryAttributeValueChangeEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CategoryAttributeValueChangeEvent**](CategoryAttributeValueChangeEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToCategoryCreateEvent**
> EventSubscription ListenToCategoryCreateEvent(ctx, data)
Client listener for entity CategoryCreateEvent

Example of a client listener for receiving the notification CategoryCreateEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CategoryCreateEvent**](CategoryCreateEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToCategoryDeleteEvent**
> EventSubscription ListenToCategoryDeleteEvent(ctx, data)
Client listener for entity CategoryDeleteEvent

Example of a client listener for receiving the notification CategoryDeleteEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CategoryDeleteEvent**](CategoryDeleteEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToCategoryStateChangeEvent**
> EventSubscription ListenToCategoryStateChangeEvent(ctx, data)
Client listener for entity CategoryStateChangeEvent

Example of a client listener for receiving the notification CategoryStateChangeEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CategoryStateChangeEvent**](CategoryStateChangeEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToProductOfferingAttributeValueChangeEvent**
> EventSubscription ListenToProductOfferingAttributeValueChangeEvent(ctx, data)
Client listener for entity ProductOfferingAttributeValueChangeEvent

Example of a client listener for receiving the notification ProductOfferingAttributeValueChangeEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductOfferingAttributeValueChangeEvent**](ProductOfferingAttributeValueChangeEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToProductOfferingCreateEvent**
> EventSubscription ListenToProductOfferingCreateEvent(ctx, data)
Client listener for entity ProductOfferingCreateEvent

Example of a client listener for receiving the notification ProductOfferingCreateEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductOfferingCreateEvent**](ProductOfferingCreateEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToProductOfferingDeleteEvent**
> EventSubscription ListenToProductOfferingDeleteEvent(ctx, data)
Client listener for entity ProductOfferingDeleteEvent

Example of a client listener for receiving the notification ProductOfferingDeleteEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductOfferingDeleteEvent**](ProductOfferingDeleteEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToProductOfferingPriceAttributeValueChangeEvent**
> EventSubscription ListenToProductOfferingPriceAttributeValueChangeEvent(ctx, data)
Client listener for entity ProductOfferingPriceAttributeValueChangeEvent

Example of a client listener for receiving the notification ProductOfferingPriceAttributeValueChangeEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductOfferingPriceAttributeValueChangeEvent**](ProductOfferingPriceAttributeValueChangeEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToProductOfferingPriceCreateEvent**
> EventSubscription ListenToProductOfferingPriceCreateEvent(ctx, data)
Client listener for entity ProductOfferingPriceCreateEvent

Example of a client listener for receiving the notification ProductOfferingPriceCreateEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductOfferingPriceCreateEvent**](ProductOfferingPriceCreateEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToProductOfferingPriceDeleteEvent**
> EventSubscription ListenToProductOfferingPriceDeleteEvent(ctx, data)
Client listener for entity ProductOfferingPriceDeleteEvent

Example of a client listener for receiving the notification ProductOfferingPriceDeleteEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductOfferingPriceDeleteEvent**](ProductOfferingPriceDeleteEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToProductOfferingPriceStateChangeEvent**
> EventSubscription ListenToProductOfferingPriceStateChangeEvent(ctx, data)
Client listener for entity ProductOfferingPriceStateChangeEvent

Example of a client listener for receiving the notification ProductOfferingPriceStateChangeEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductOfferingPriceStateChangeEvent**](ProductOfferingPriceStateChangeEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToProductOfferingStateChangeEvent**
> EventSubscription ListenToProductOfferingStateChangeEvent(ctx, data)
Client listener for entity ProductOfferingStateChangeEvent

Example of a client listener for receiving the notification ProductOfferingStateChangeEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductOfferingStateChangeEvent**](ProductOfferingStateChangeEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToProductSpecificationAttributeValueChangeEvent**
> EventSubscription ListenToProductSpecificationAttributeValueChangeEvent(ctx, data)
Client listener for entity ProductSpecificationAttributeValueChangeEvent

Example of a client listener for receiving the notification ProductSpecificationAttributeValueChangeEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductSpecificationAttributeValueChangeEvent**](ProductSpecificationAttributeValueChangeEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToProductSpecificationCreateEvent**
> EventSubscription ListenToProductSpecificationCreateEvent(ctx, data)
Client listener for entity ProductSpecificationCreateEvent

Example of a client listener for receiving the notification ProductSpecificationCreateEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductSpecificationCreateEvent**](ProductSpecificationCreateEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToProductSpecificationDeleteEvent**
> EventSubscription ListenToProductSpecificationDeleteEvent(ctx, data)
Client listener for entity ProductSpecificationDeleteEvent

Example of a client listener for receiving the notification ProductSpecificationDeleteEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductSpecificationDeleteEvent**](ProductSpecificationDeleteEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListenToProductSpecificationStateChangeEvent**
> EventSubscription ListenToProductSpecificationStateChangeEvent(ctx, data)
Client listener for entity ProductSpecificationStateChangeEvent

Example of a client listener for receiving the notification ProductSpecificationStateChangeEvent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductSpecificationStateChangeEvent**](ProductSpecificationStateChangeEvent.md)| The event data | 

### Return type

[**EventSubscription**](EventSubscription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json;charset=utf-8
 - **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

