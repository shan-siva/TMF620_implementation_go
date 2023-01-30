# ProductSpecificationCharacteristicValueUse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the characteristic | [optional] [default to null]
**Description** | **string** | A narrative that explains in detail what the productSpecificationCharacteristic is | [optional] [default to null]
**MaxCardinality** | **int32** | The maximum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where five is the value for the maxCardinality. | [optional] [default to null]
**MinCardinality** | **int32** | The minimum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where zero is the value for the minCardinality. | [optional] [default to null]
**Name** | **string** | Name of the associated productSpecificationCharacteristic | [optional] [default to null]
**ValueType** | **string** | A kind of value that the characteristic can take on, such as numeric, text and so forth | [optional] [default to null]
**ProductSpecCharacteristicValue** | [**[]CharacteristicValueSpecification**](CharacteristicValueSpecification.md) | A number or text that can be assigned to a ProductSpecificationCharacteristic. | [optional] [default to null]
**ProductSpecification** | [***ProductSpecificationRef**](ProductSpecificationRef.md) | A ProductSpecification is a detailed description of a tangible or intangible object made available externally in the form of a ProductOffering to customers or other parties playing a party role. | [optional] [default to null]
**ValidFor** | [***TimePeriod**](TimePeriod.md) | The period for which the productSpecificationCharacteristic is valid | [optional] [default to null]
**BaseType** | **string** | When sub-classing, this defines the super-class | [optional] [default to null]
**SchemaLocation** | **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] [default to null]
**Type_** | **string** | When sub-classing, this defines the sub-class Extensible name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


