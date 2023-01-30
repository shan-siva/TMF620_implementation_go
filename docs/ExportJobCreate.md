# ExportJobCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionDate** | [**time.Time**](time.Time.md) | Data at which the job was completed | [optional] [default to null]
**ContentType** | **string** | The format of the exported data | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | Date at which the job was created | [optional] [default to null]
**ErrorLog** | **string** | Reason for failure | [optional] [default to null]
**Path** | **string** | URL of the root resource acting as the source for streaming content to the file specified by the export job | [optional] [default to null]
**Query** | **string** | Used to scope the exported data | [optional] [default to null]
**Url** | **string** | URL of the file containing the data to be exported | [default to null]
**Status** | [***JobStateType**](JobStateType.md) | Status of the export job (not started, running, succeeded, failed) | [optional] [default to null]
**BaseType** | **string** | When sub-classing, this defines the super-class | [optional] [default to null]
**SchemaLocation** | **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] [default to null]
**Type_** | **string** | When sub-classing, this defines the sub-class Extensible name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


