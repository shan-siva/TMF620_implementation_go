# CatalogUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogType** | **string** | Indicates if the catalog is a product, service or resource catalog | [optional] [default to null]
**Description** | **string** | Description of this catalog | [optional] [default to null]
**LifecycleStatus** | **string** | Used to indicate the current lifecycle status | [optional] [default to null]
**Name** | **string** | Name of the catalog | [optional] [default to null]
**Version** | **string** | Catalog version | [optional] [default to null]
**Category** | [**[]CategoryRef**](CategoryRef.md) | List of root categories contained in this catalog | [optional] [default to null]
**RelatedParty** | [**[]RelatedParty**](RelatedParty.md) | List of parties involved in this catalog | [optional] [default to null]
**ValidFor** | [***TimePeriod**](TimePeriod.md) | The period for which the catalog is valid | [optional] [default to null]
**SchemaLocation** | **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


