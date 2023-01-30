/*
 * Product Catalog Management
 *
 * ## TMF API Reference: TMF620 - Product Catalog Management  ### Release : 20.5 - March 2021  Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.   ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 4.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A use of the ProductSpecificationCharacteristicValue by a ProductOffering to which additional properties (attributes) apply or override the properties of similar properties contained in ProductSpecificationCharacteristicValue. It should be noted that characteristics which their value(s) addressed by this object must exist in corresponding product specification. The available characteristic values for a ProductSpecificationCharacteristic in a Product specification can be modified at the ProductOffering level. For example, a characteristic 'Color' might have values White, Blue, Green, and Red. But, the list of values can be restricted to e.g. White and Blue in an associated product offering. It should be noted that the list of values in 'ProductSpecificationCharacteristicValueUse' is a strict subset of the list of values as defined in the corresponding product specification characteristics.
type ProductSpecificationCharacteristicValueUse struct {
	// Unique ID for the characteristic
	Id string `json:"id,omitempty"`
	// A narrative that explains in detail what the productSpecificationCharacteristic is
	Description string `json:"description,omitempty"`
	// The maximum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where five is the value for the maxCardinality.
	MaxCardinality int32 `json:"maxCardinality,omitempty"`
	// The minimum number of instances a CharacteristicValue can take on. For example, zero to five phone numbers in a group calling plan, where zero is the value for the minCardinality.
	MinCardinality int32 `json:"minCardinality,omitempty"`
	// Name of the associated productSpecificationCharacteristic
	Name string `json:"name,omitempty"`
	// A kind of value that the characteristic can take on, such as numeric, text and so forth
	ValueType string `json:"valueType,omitempty"`
	// A number or text that can be assigned to a ProductSpecificationCharacteristic.
	ProductSpecCharacteristicValue []CharacteristicValueSpecification `json:"productSpecCharacteristicValue,omitempty"`
	// A ProductSpecification is a detailed description of a tangible or intangible object made available externally in the form of a ProductOffering to customers or other parties playing a party role.
	ProductSpecification *ProductSpecificationRef `json:"productSpecification,omitempty"`
	// The period for which the productSpecificationCharacteristic is valid
	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	// When sub-classing, this defines the sub-class Extensible name
	Type_ string `json:"@type,omitempty"`
}
