/*
 * Product Catalog Management
 *
 * ## TMF API Reference: TMF620 - Product Catalog Management  ### Release : 20.5 - March 2021  Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.   ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 4.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The PricingLogicAlgorithm entity represents an instantiation of an interface specification to external rating function (without a modeled bahavior in SID). Some of the parameters of the interface definiition may be already set (such as price per unit) and some may be gathered during the rating process from the event (such as call duration) or from ProductCharacteristicValues (such as assigned bandwidth)
type PricingLogicAlgorithm struct {
	// unique identifier
	Id string `json:"id,omitempty"`
	// Hyperlink reference
	Href string `json:"href,omitempty"`
	// Description of the PricingLogicAlgorithm
	Description string `json:"description,omitempty"`
	// Name given to the PricingLogicAlgorithm
	Name string `json:"name,omitempty"`
	// id of corresponding PricingLogicAlgorithm specification
	PlaSpecId string `json:"plaSpecId,omitempty"`
	// The period for which the PricingLogicAlgorithm is valid
	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	// When sub-classing, this defines the sub-class Extensible name
	Type_ string `json:"@type,omitempty"`
}
