# BlockMessageActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AutomodActionType**](AutomodActionType.md) |  | 
**Metadata** | [**BlockMessageActionMetadataResponse**](BlockMessageActionMetadataResponse.md) |  | 

## Methods

### NewBlockMessageActionResponse

`func NewBlockMessageActionResponse(type_ AutomodActionType, metadata BlockMessageActionMetadataResponse, ) *BlockMessageActionResponse`

NewBlockMessageActionResponse instantiates a new BlockMessageActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockMessageActionResponseWithDefaults

`func NewBlockMessageActionResponseWithDefaults() *BlockMessageActionResponse`

NewBlockMessageActionResponseWithDefaults instantiates a new BlockMessageActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlockMessageActionResponse) GetType() AutomodActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlockMessageActionResponse) GetTypeOk() (*AutomodActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlockMessageActionResponse) SetType(v AutomodActionType)`

SetType sets Type field to given value.


### GetMetadata

`func (o *BlockMessageActionResponse) GetMetadata() BlockMessageActionMetadataResponse`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BlockMessageActionResponse) GetMetadataOk() (*BlockMessageActionMetadataResponse, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BlockMessageActionResponse) SetMetadata(v BlockMessageActionMetadataResponse)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


