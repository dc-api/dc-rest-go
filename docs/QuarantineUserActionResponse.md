# QuarantineUserActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AutomodActionType**](AutomodActionType.md) |  | 
**Metadata** | **map[string]interface{}** |  | 

## Methods

### NewQuarantineUserActionResponse

`func NewQuarantineUserActionResponse(type_ AutomodActionType, metadata map[string]interface{}, ) *QuarantineUserActionResponse`

NewQuarantineUserActionResponse instantiates a new QuarantineUserActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuarantineUserActionResponseWithDefaults

`func NewQuarantineUserActionResponseWithDefaults() *QuarantineUserActionResponse`

NewQuarantineUserActionResponseWithDefaults instantiates a new QuarantineUserActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QuarantineUserActionResponse) GetType() AutomodActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuarantineUserActionResponse) GetTypeOk() (*AutomodActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuarantineUserActionResponse) SetType(v AutomodActionType)`

SetType sets Type field to given value.


### GetMetadata

`func (o *QuarantineUserActionResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QuarantineUserActionResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QuarantineUserActionResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


