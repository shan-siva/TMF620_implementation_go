# ImportJobCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionDate** | [**time.Time**](time.Time.md) | Date at which the job was completed | [optional] [default to null]
**ContentType** | **string** | Indicates the format of the imported data | [optional] [default to null]
**CreationDate** | [**time.Time**](time.Time.md) | Date at which the job was created | [optional] [default to null]
**ErrorLog** | **string** | Reason for failure if status is failed | [optional] [default to null]
**Path** | **string** | URL of the root resource where the content of the file specified by the import job must be applied | [optional] [default to null]
**Url** | **string** | URL of the file containing the data to be imported | [default to null]
**Status** | [***JobStateType**](JobStateType.md) | Status of the import job (not started, running, succeeded, failed) | [optional] [default to null]
**BaseType** | **string** | When sub-classing, this defines the super-class | [optional] [default to null]
**SchemaLocation** | **string** | A URI to a JSON-Schema file that defines additional attributes and relationships | [optional] [default to null]
**Type_** | **string** | When sub-classing, this defines the sub-class Extensible name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


