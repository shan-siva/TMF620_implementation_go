# ProductSpecificationCharacteristicRelationship

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | the identifier of the associated product specification | [optional] [default to null]
**Href** | **string** | Hyperlink reference to the target product specification | [optional] [default to null]
**CharSpecSeq** | **int32** | The order in which a CharacteristicSpecification appears within another CharacteristicSpecification that defines a grouping of CharacteristicSpecifications.  For example, a grouping may represent the name of an individual. The given name is first, the middle name is second, and the last name is third. | [optional] [default to null]
**Name** | **string** | Name of the target product specification characteristic | [optional] [default to null]
**RelationshipType** | **string** | Type of relationship such as aggregation, migration, substitution, dependency, exclusivity | [optional] [default to null]
**ValidFor** | [***TimePeriod**](TimePeriod.md) | The period for which the productSpecificationCharacteristicRelationship is valid | [optional] [default to null]
**BaseType** | **string** | When sub-classing, this defines the super-class | [optional] [default to null]
**SchemaLocation** | **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] [default to null]
**Type_** | **string** | When sub-classing, this defines the sub-class Extensible name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


