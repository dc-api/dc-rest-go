# BlockMessageAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Metadata** | Pointer to [**NullableBlockMessageActionMetadata**](BlockMessageActionMetadata.md) |  | [optional] 

## Methods

### NewBlockMessageAction

`func NewBlockMessageAction(type_ int32, ) *BlockMessageAction`

NewBlockMessageAction instantiates a new BlockMessageAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockMessageActionWithDefaults

`func NewBlockMessageActionWithDefaults() *BlockMessageAction`

NewBlockMessageActionWithDefaults instantiates a new BlockMessageAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlockMessageAction) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlockMessageAction) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlockMessageAction) SetType(v int32)`

SetType sets Type field to given value.


### GetMetadata

`func (o *BlockMessageAction) GetMetadata() BlockMessageActionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlockMessageAction) GetMetadataOk() (*BlockMessageActionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlockMessageAction) SetMetadata(v BlockMessageActionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BlockMessageAction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BlockMessageAction) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BlockMessageAction) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


