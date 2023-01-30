# PopCharge

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique identifier | [optional] [default to null]
**Href** | **string** | Hyperlink reference | [optional] [default to null]
**Description** | **string** | Description of the productOfferingPrice | [optional] [default to null]
**LastUpdate** | [**time.Time**](time.Time.md) | the last update time of this ProductOfferingPrice | [optional] [default to null]
**LifecycleStatus** | **string** | the lifecycle status of this ProductOfferingPrice | [optional] [default to null]
**Name** | **string** | Name of the productOfferingPrice | [optional] [default to null]
**PriceType** | **string** | A category that describes the price charge, such as recurring, penalty, One time fee and so forth. | [optional] [default to null]
**RecurringChargePeriod** | **string** | The period type to repeat the application of the price Could be month, week... | [optional] [default to null]
**RecurringChargePeriodLength** | **int32** | the period of the recurring charge:  1, 2, ... .It sets to zero if it is not applicable | [optional] [default to null]
**Version** | **string** | ProductOffering version | [optional] [default to null]
**Constraint** | [**[]ConstraintRef**](ConstraintRef.md) | The Constraint resource represents a policy/rule applied to ProductOfferingPrice. | [optional] [default to null]
**Price** | [***ProductPriceValue**](ProductPriceValue.md) |  | [optional] [default to null]
**PriceAlteration** | [**[]PopAlteration**](POPAlteration.md) |  | [optional] [default to null]
**UnitOfMeasure** | [***Quantity**](Quantity.md) | A number and unit representing denominator of a rate. For example, for a data charge rate of $2 per 5 GB usage, the amount of unitOfMeasure will be 5 with units as GB. | [optional] [default to null]
**ValidFor** | [***TimePeriod**](TimePeriod.md) | The period for which the productOfferingPrice is valid | [optional] [default to null]
**BaseType** | **string** | When sub-classing, this defines the super-class | [optional] [default to null]
**SchemaLocation** | **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] [default to null]
**Type_** | **string** | When sub-classing, this defines the sub-class Extensible name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


