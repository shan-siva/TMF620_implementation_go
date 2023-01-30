# ProductOfferingRelationship

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique identifier | [optional] [default to null]
**Href** | **string** | Hyperlink reference | [optional] [default to null]
**Name** | **string** | Name of the related entity. | [optional] [default to null]
**RelationshipType** | **string** | Type of relationship between product offerings such as requires, exchangableTo, optionalFor | [optional] [default to null]
**Role** | **string** | The association role for the source product offering | [optional] [default to null]
**ValidFor** | [***TimePeriod**](TimePeriod.md) | The period for which the Relationship is valid | [optional] [default to null]
**BaseType** | **string** | When sub-classing, this defines the super-class | [optional] [default to null]
**SchemaLocation** | **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] [default to null]
**Type_** | **string** | When sub-classing, this defines the sub-class Extensible name | [optional] [default to null]
**ReferredType** | **string** | The actual type of the target instance when needed for disambiguation. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


