/*
 * Product Catalog Management
 *
 * ## TMF API Reference: TMF620 - Product Catalog Management  ### Release : 20.5 - March 2021  Product Catalog API is one of Catalog Management API Family. Product Catalog API goal is to provide a catalog of products.   ### Operations Product Catalog API performs the following operations on the resources : - Retrieve an entity or a collection of entities depending on filter criteria - Partial update of an entity (including updating rules) - Create an entity (including default values and creation rules) - Delete an entity - Manage notification of events
 *
 * API version: 4.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// Is a detailed description of a tangible or intangible object made available externally in the form of a ProductOffering to customers or other parties playing a party role.
type ProductSpecification struct {
	// Unique identifier of the product specification
	Id string `json:"id,omitempty"`
	// Reference of the product specification
	Href string `json:"href,omitempty"`
	// The manufacturer or trademark of the specification
	Brand string `json:"brand,omitempty"`
	// A narrative that explains in detail what the product specification is
	Description string `json:"description,omitempty"`
	// isBundle determines whether a productSpecification represents a single productSpecification (false), or a bundle of productSpecification (true).
	IsBundle bool `json:"isBundle,omitempty"`
	// Date and time of the last update
	LastUpdate time.Time `json:"lastUpdate,omitempty"`
	// Used to indicate the current lifecycle status
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`
	// Name of the product specification
	Name string `json:"name,omitempty"`
	// An identification number assigned to uniquely identity the specification
	ProductNumber string `json:"productNumber,omitempty"`
	// Product specification version
	Version string `json:"version,omitempty"`
	// Complements the description of an element (for instance a product) through video, pictures...
	Attachment []AttachmentRefOrValue `json:"attachment,omitempty"`
	// A type of ProductSpecification that belongs to a grouping of ProductSpecifications made available to the market. It inherits of all attributes of ProductSpecification.
	BundledProductSpecification []BundledProductSpecification `json:"bundledProductSpecification,omitempty"`
	// A characteristic quality or distinctive feature of a ProductSpecification.  The characteristic can be take on a discrete value, such as color, can take on a range of values, (for example, sensitivity of 100-240 mV), or can be derived from a formula (for example, usage time (hrs) = 30 - talk time *3). Certain characteristics, such as color, may be configured during the ordering or some other process.
	ProductSpecCharacteristic []ProductSpecificationCharacteristic `json:"productSpecCharacteristic,omitempty"`
	// A migration, substitution, dependency or exclusivity relationship between/among product specifications.
	ProductSpecificationRelationship []ProductSpecificationRelationship `json:"productSpecificationRelationship,omitempty"`
	// A related party defines party or party role linked to a specific entity.
	RelatedParty []RelatedParty `json:"relatedParty,omitempty"`
	// The ResourceSpecification is required to realize a ProductSpecification.
	ResourceSpecification []ResourceSpecificationRef `json:"resourceSpecification,omitempty"`
	// ServiceSpecification(s) required to realize a ProductSpecification.
	ServiceSpecification []ServiceSpecificationRef `json:"serviceSpecification,omitempty"`
	// A target product schema reference. The reference object to the schema and type of target product which is described by product specification.
	TargetProductSchema *TargetProductSchema `json:"targetProductSchema,omitempty"`
	// The period for which the product specification is valid
	ValidFor *TimePeriod `json:"validFor,omitempty"`
	// When sub-classing, this defines the super-class
	BaseType string `json:"@baseType,omitempty"`
	// A URI to a JSON-Schema file that defines additional attributes and relationships
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	// When sub-classing, this defines the sub-class Extensible name
	Type_ string `json:"@type,omitempty"`
}
