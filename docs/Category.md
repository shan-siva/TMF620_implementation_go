# Category

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the category | [optional] [default to null]
**Href** | **string** | Reference of the category | [optional] [default to null]
**Description** | **string** | Description of the category | [optional] [default to null]
**IsRoot** | **bool** | If true, this Boolean indicates that the category is a root of categories | [optional] [default to null]
**LastUpdate** | [**time.Time**](time.Time.md) | Date and time of the last update | [optional] [default to null]
**LifecycleStatus** | **string** | Used to indicate the current lifecycle status | [optional] [default to null]
**Name** | **string** | Name of the category | [optional] [default to null]
**ParentId** | **string** | Unique identifier of the parent category | [optional] [default to null]
**Version** | **string** | Category version | [optional] [default to null]
**ProductOffering** | [**[]ProductOfferingRef**](ProductOfferingRef.md) | A product offering represents entities that are orderable from the provider of the catalog, this resource includes pricing information. | [optional] [default to null]
**SubCategory** | [**[]CategoryRef**](CategoryRef.md) | The category resource is used to group product offerings, service and resource candidates in logical containers. Categories can contain other (sub-)categories and/or product offerings. | [optional] [default to null]
**ValidFor** | [***TimePeriod**](TimePeriod.md) | The period for which the category is valid | [optional] [default to null]
**BaseType** | **string** | When sub-classing, this defines the super-class | [optional] [default to null]
**SchemaLocation** | **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] [default to null]
**Type_** | **string** | When sub-classing, this defines the sub-class Extensible name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


