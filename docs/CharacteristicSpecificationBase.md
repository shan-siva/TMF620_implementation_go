# CharacteristicSpecificationBase

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the characteristic | [optional] [default to null]
**Configurable** | **bool** | If true, the Boolean indicates that the target Characteristic is configurable | [optional] [default to null]
**Description** | **string** | A narrative that explains the CharacteristicSpecification. | [optional] [default to null]
**Extensible** | **bool** | An indicator that specifies that the values for the characteristic can be extended by adding new values when instantiating a characteristic for a resource. | [optional] [default to null]
**IsUnique** | **bool** | An indicator that specifies if a value is unique for the specification. Possible values are; \&quot;unique while value is in effect\&quot; and \&quot;unique whether value is in effect or not\&quot; | [optional] [default to null]
**MaxCardinality** | **int32** | The maximum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where five is the value for the maxCardinality. | [optional] [default to null]
**MinCardinality** | **int32** | The minimum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where zero is the value for the minCardinality. | [optional] [default to null]
**Name** | **string** | A word, term, or phrase by which this characteristic specification is known and distinguished from other characteristic specifications. | [optional] [default to null]
**Regex** | **string** | A rule or principle represented in regular expression used to derive the value of a characteristic value. | [optional] [default to null]
**ValueType** | **string** | A kind of value that the characteristic can take on, such as numeric, text and so forth | [optional] [default to null]
**ValidFor** | [***TimePeriod**](TimePeriod.md) | The period of time for which a characteristic is applicable. | [optional] [default to null]
**BaseType** | **string** | When sub-classing, this defines the super-class | [optional] [default to null]
**SchemaLocation** | **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] [default to null]
**Type_** | **string** | When sub-classing, this defines the sub-class Extensible name | [optional] [default to null]
**ValueSchemaLocation** | **string** | This (optional) field provides a link to the schema describing the value type. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


