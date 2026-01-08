# QuarantineUserAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AutomodActionType**](AutomodActionType.md) |  | 
**Metadata** | Pointer to [**QuarantineUserActionMetadata**](QuarantineUserActionMetadata.md) |  | [optional] 

## Methods

### NewQuarantineUserAction

`func NewQuarantineUserAction(type_ AutomodActionType, ) *QuarantineUserAction`

NewQuarantineUserAction instantiates a new QuarantineUserAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuarantineUserActionWithDefaults

`func NewQuarantineUserActionWithDefaults() *QuarantineUserAction`

NewQuarantineUserActionWithDefaults instantiates a new QuarantineUserAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QuarantineUserAction) GetType() AutomodActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QuarantineUserAction) GetTypeOk() (*AutomodActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QuarantineUserAction) SetType(v AutomodActionType)`

SetType sets Type field to given value.


### GetMetadata

`func (o *QuarantineUserAction) GetMetadata() QuarantineUserActionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *QuarantineUserAction) GetMetadataOk() (*QuarantineUserActionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *QuarantineUserAction) SetMetadata(v QuarantineUserActionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *QuarantineUserAction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


