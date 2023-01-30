/*
 * Product Catalog Management
 *
 * ## TMF API Reference: TMF620 - Product Catalog Management  ### Release : 20.5 - March 2021  Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.   ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 4.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Is based on both the basic cost to develop and produce products and the enterprises policy on revenue targets. This price may be further revised through discounting (a Product Offering Price that reflects an alteration). The price, applied for a productOffering may also be influenced by the productOfferingTerm, the customer selected, eg: a productOffering can be offered with multiple terms, like commitment periods for the contract. The price may be influenced by this productOfferingTerm. A productOffering may be cheaper with a 24 month commitment than with a 12 month commitment. Skipped properties: id,href,lastUpdate,@type,@baseType
type ProductOfferingPriceUpdate struct {
	// Description of the productOfferingPrice
	Description string `json:"description,omitempty"`
	// A flag indicating if this ProductOfferingPrice is composite (bundle) or not
	IsBundle bool `json:"isBundle,omitempty"`
	// the lifecycle status of this ProductOfferingPrice
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`
	// Name of the productOfferingPrice
	Name string `json:"name,omitempty"`
	// Percentage to apply if this Product Offering Price is an Alteration (such as a Discount)
	Percentage float32 `json:"percentage,omitempty"`
	// A category that describes the price, such as recurring, discount, allowance, penalty, and so forth.
	PriceType string `json:"priceType,omitempty"`
	// the period of the recurring charge:  1, 2, ... .It sets to zero if it is not applicable
	RecurringChargePeriodLength int32 `json:"recurringChargePeriodLength,omitempty"`
	// The period to repeat the application of the price Could be month, week...
	RecurringChargePeriodType string `json:"recurringChargePeriodType,omitempty"`
	// ProductOfferingPrice version
	Version string `json:"version,omitempty"`
	// this object represents a bundle relationship from a bundle product offering price (parent) to a simple product offering price (child). A simple product offering price may participate in more than one bundle relationship.
	BundledPopRelationship []BundledProductOfferingPriceRelationship `json:"bundledPopRelationship,omitempty"`
	// The Constraint resource represents a policy/rule applied to ProductOfferingPrice.
	Constraint []ConstraintRef `json:"constraint,omitempty"`
	// Place defines the places where the products are sold or delivered.
	Place []PlaceRef `json:"place,omitempty"`
	// Product Offering Prices related to this Product Offering Price, for example a price alteration such as allowance or discount
	PopRelationship []ProductOfferingPriceRelationship `json:"popRelationship,omitempty"`
	// The amount of money that characterizes the price.
	Price *Money `json:"price,omitempty"`
	// The PricingLogicAlgorithm entity represents an instantiation of an interface specification to external rating function (without a modeled behavior in SID). Some of the parameters of the interface definition may be already set (such as price per unit) and some may be gathered during the rating process from the event (such as call duration) or from ProductCharacteristicValues (such as assigned bandwidth).
	PricingLogicAlgorithm []PricingLogicAlgorithm `json:"pricingLogicAlgorithm,omitempty"`
	// A use of the ProductSpecificationCharacteristicValue by a ProductOfferingPrice to which additional properties (attributes) apply or override the properties of similar properties contained in ProductSpecificationCharacteristicValue. It should be noted that characteristics which their value(s) addressed by this object must exist in corresponding product specification. The available characteristic values for a ProductSpecificationCharacteristic in a Product specification can be modified at the ProductOffering and ProcuctOfferingPrice level. The list of values in ProductSpecificationCharacteristicValueUse is a strict subset of the list of values as defined in the corresponding product specification characteristics.
	ProdSpecCharValueUse []ProductSpecificationCharacteristicValueUse `json:"prodSpecCharValueUse,omitempty"`
	// A list of conditions under which a ProductOfferingPrice is made available to Customers. For instance, a Product Offering Price can be offered with multiple commitment periods.
	ProductOfferingTerm []ProductOfferingTerm `json:"productOfferingTerm,omitempty"`
	// An amount of money levied on the price of a Product by a legislative body.
	Tax []TaxItem `json:"tax,omitempty"`
	// A number and unit representing how many (for instance 1 dozen) of an ProductOffering is available at the offered price. Its meaning depends on the priceType. It could be a price, a rate, or a discount.
	UnitOfMeasure *Quantity `json:"unitOfMeasure,omitempty"`
	// The period for which the productOfferingPrice is valid
	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// hyperlink reference to the schema describing this resource
	SchemaLocation string `json:"@schemaLocation,omitempty"`
}
