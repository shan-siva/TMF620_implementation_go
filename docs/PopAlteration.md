# PopAlteration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique identifier | [optional] [default to null]
**Href** | **string** | Hyperlink reference | [optional] [default to null]
**Description** | **string** | A narrative that explains in detail the semantics of this order item price alteration | [optional] [default to null]
**Name** | **string** | Name given to this price alteration | [optional] [default to null]
**PriceType** | **string** | A category that describes the price such as recurring, one time and usage. | [default to null]
**Priority** | **int32** | Priority level for applying this alteration among all the defined alterations on the order item price | [optional] [default to null]
**RecurringChargePeriod** | **string** | Could be month, week... | [optional] [default to null]
**ApplicationDuration** | [***Duration**](Duration.md) | The period for which the productOfferingPriceAlteration is applicable | [optional] [default to null]
**Price** | [***ProductPriceValue**](ProductPriceValue.md) |  | [default to null]
**UnitOfMeasure** | [***Quantity**](Quantity.md) | A number and unit representing denominator of an alteration rate. For example, for a data discount rate of $1 per 20 GB usage, the amount of unitOfMeasure will be 20 with units as GB. | [optional] [default to null]
**ValidFor** | [***TimePeriod**](TimePeriod.md) | The period for which this productOfferingPriceAlteration is valid | [optional] [default to null]
**BaseType** | **string** | When sub-classing, this defines the super-class | [optional] [default to null]
**SchemaLocation** | **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] [default to null]
**Type_** | **string** | When sub-classing, this defines the sub-class Extensible name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


