# ProductSpecificationUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | **string** | The manufacturer or trademark of the specification | [optional] [default to null]
**Description** | **string** | A narrative that explains in detail what the product specification is | [optional] [default to null]
**IsBundle** | **bool** | isBundle determines whether a productSpecification represents a single productSpecification (false), or a bundle of productSpecification (true). | [optional] [default to null]
**LifecycleStatus** | **string** | Used to indicate the current lifecycle status | [optional] [default to null]
**Name** | **string** | Name of the product specification | [optional] [default to null]
**ProductNumber** | **string** | An identification number assigned to uniquely identity the specification | [optional] [default to null]
**Version** | **string** | Product specification version | [optional] [default to null]
**Attachment** | [**[]AttachmentRefOrValue**](AttachmentRefOrValue.md) | Complements the description of an element (for instance a product) through video, pictures... | [optional] [default to null]
**BundledProductSpecification** | [**[]BundledProductSpecification**](BundledProductSpecification.md) | A type of ProductSpecification that belongs to a grouping of ProductSpecifications made available to the market. It inherits of all attributes of ProductSpecification. | [optional] [default to null]
**ProductSpecCharacteristic** | [**[]ProductSpecificationCharacteristic**](ProductSpecificationCharacteristic.md) | A characteristic quality or distinctive feature of a ProductSpecification.  The characteristic can be take on a discrete value, such as color, can take on a range of values, (for example, sensitivity of 100-240 mV), or can be derived from a formula (for example, usage time (hrs) &#x3D; 30 - talk time *3). Certain characteristics, such as color, may be configured during the ordering or some other process. | [optional] [default to null]
**ProductSpecificationRelationship** | [**[]ProductSpecificationRelationship**](ProductSpecificationRelationship.md) | A migration, substitution, dependency or exclusivity relationship between/among product specifications. | [optional] [default to null]
**RelatedParty** | [**[]RelatedParty**](RelatedParty.md) | A related party defines party or party role linked to a specific entity. | [optional] [default to null]
**ResourceSpecification** | [**[]ResourceSpecificationRef**](ResourceSpecificationRef.md) | The ResourceSpecification is required to realize a ProductSpecification. | [optional] [default to null]
**ServiceSpecification** | [**[]ServiceSpecificationRef**](ServiceSpecificationRef.md) | ServiceSpecification(s) required to realize a ProductSpecification. | [optional] [default to null]
**TargetProductSchema** | [***TargetProductSchema**](TargetProductSchema.md) | A target product schema reference. The reference object to the schema and type of target product which is described by product specification. | [optional] [default to null]
**ValidFor** | [***TimePeriod**](TimePeriod.md) | The period for which the product specification is valid | [optional] [default to null]
**SchemaLocation** | **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


