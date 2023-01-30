# ProductOfferingPriceUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description of the productOfferingPrice | [optional] [default to null]
**IsBundle** | **bool** | A flag indicating if this ProductOfferingPrice is composite (bundle) or not | [optional] [default to null]
**LifecycleStatus** | **string** | the lifecycle status of this ProductOfferingPrice | [optional] [default to null]
**Name** | **string** | Name of the productOfferingPrice | [optional] [default to null]
**Percentage** | **float32** | Percentage to apply if this Product Offering Price is an Alteration (such as a Discount) | [optional] [default to null]
**PriceType** | **string** | A category that describes the price, such as recurring, discount, allowance, penalty, and so forth. | [optional] [default to null]
**RecurringChargePeriodLength** | **int32** | the period of the recurring charge:  1, 2, ... .It sets to zero if it is not applicable | [optional] [default to null]
**RecurringChargePeriodType** | **string** | The period to repeat the application of the price Could be month, week... | [optional] [default to null]
**Version** | **string** | ProductOfferingPrice version | [optional] [default to null]
**BundledPopRelationship** | [**[]BundledProductOfferingPriceRelationship**](BundledProductOfferingPriceRelationship.md) | this object represents a bundle relationship from a bundle product offering price (parent) to a simple product offering price (child). A simple product offering price may participate in more than one bundle relationship. | [optional] [default to null]
**Constraint** | [**[]ConstraintRef**](ConstraintRef.md) | The Constraint resource represents a policy/rule applied to ProductOfferingPrice. | [optional] [default to null]
**Place** | [**[]PlaceRef**](PlaceRef.md) | Place defines the places where the products are sold or delivered. | [optional] [default to null]
**PopRelationship** | [**[]ProductOfferingPriceRelationship**](ProductOfferingPriceRelationship.md) | Product Offering Prices related to this Product Offering Price, for example a price alteration such as allowance or discount | [optional] [default to null]
**Price** | [***Money**](Money.md) | The amount of money that characterizes the price. | [optional] [default to null]
**PricingLogicAlgorithm** | [**[]PricingLogicAlgorithm**](PricingLogicAlgorithm.md) | The PricingLogicAlgorithm entity represents an instantiation of an interface specification to external rating function (without a modeled behavior in SID). Some of the parameters of the interface definition may be already set (such as price per unit) and some may be gathered during the rating process from the event (such as call duration) or from ProductCharacteristicValues (such as assigned bandwidth). | [optional] [default to null]
**ProdSpecCharValueUse** | [**[]ProductSpecificationCharacteristicValueUse**](ProductSpecificationCharacteristicValueUse.md) | A use of the ProductSpecificationCharacteristicValue by a ProductOfferingPrice to which additional properties (attributes) apply or override the properties of similar properties contained in ProductSpecificationCharacteristicValue. It should be noted that characteristics which their value(s) addressed by this object must exist in corresponding product specification. The available characteristic values for a ProductSpecificationCharacteristic in a Product specification can be modified at the ProductOffering and ProcuctOfferingPrice level. The list of values in ProductSpecificationCharacteristicValueUse is a strict subset of the list of values as defined in the corresponding product specification characteristics. | [optional] [default to null]
**ProductOfferingTerm** | [**[]ProductOfferingTerm**](ProductOfferingTerm.md) | A list of conditions under which a ProductOfferingPrice is made available to Customers. For instance, a Product Offering Price can be offered with multiple commitment periods. | [optional] [default to null]
**Tax** | [**[]TaxItem**](TaxItem.md) | An amount of money levied on the price of a Product by a legislative body. | [optional] [default to null]
**UnitOfMeasure** | [***Quantity**](Quantity.md) | A number and unit representing how many (for instance 1 dozen) of an ProductOffering is available at the offered price. Its meaning depends on the priceType. It could be a price, a rate, or a discount. | [optional] [default to null]
**ValidFor** | [***TimePeriod**](TimePeriod.md) | The period for which the productOfferingPrice is valid | [optional] [default to null]
**SchemaLocation** | **string** | hyperlink reference to the schema describing this resource | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


