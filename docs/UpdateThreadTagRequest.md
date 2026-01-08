# UpdateThreadTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EmojiId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**EmojiName** | Pointer to **string** |  | [optional] 
**Moderated** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 

## Methods

### NewUpdateThreadTagRequest

`func NewUpdateThreadTagRequest(name string, ) *UpdateThreadTagRequest`

NewUpdateThreadTagRequest instantiates a new UpdateThreadTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateThreadTagRequestWithDefaults

`func NewUpdateThreadTagRequestWithDefaults() *UpdateThreadTagRequest`

NewUpdateThreadTagRequestWithDefaults instantiates a new UpdateThreadTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateThreadTagRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateThreadTagRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateThreadTagRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmojiId

`func (o *UpdateThreadTagRequest) GetEmojiId() GetEntitlementsSkuIdsParameterOneOfInner`

GetEmojiId returns the EmojiId field if non-nil, zero value otherwise.

### GetEmojiIdOk

`func (o *UpdateThreadTagRequest) GetEmojiIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetEmojiIdOk returns a tuple with the EmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiId

`func (o *UpdateThreadTagRequest) SetEmojiId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetEmojiId sets EmojiId field to given value.

### HasEmojiId

`func (o *UpdateThreadTagRequest) HasEmojiId() bool`

HasEmojiId returns a boolean if a field has been set.

### GetEmojiName

`func (o *UpdateThreadTagRequest) GetEmojiName() string`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *UpdateThreadTagRequest) GetEmojiNameOk() (*string, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *UpdateThreadTagRequest) SetEmojiName(v string)`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *UpdateThreadTagRequest) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### GetModerated

`func (o *UpdateThreadTagRequest) GetModerated() bool`

GetModerated returns the Moderated field if non-nil, zero value otherwise.

### GetModeratedOk

`func (o *UpdateThreadTagRequest) GetModeratedOk() (*bool, bool)`

GetModeratedOk returns a tuple with the Moderated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerated

`func (o *UpdateThreadTagRequest) SetModerated(v bool)`

SetModerated sets Moderated field to given value.

### HasModerated

`func (o *UpdateThreadTagRequest) HasModerated() bool`

HasModerated returns a boolean if a field has been set.

### GetId

`func (o *UpdateThreadTagRequest) GetId() GetEntitlementsSkuIdsParameterOneOfInner`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateThreadTagRequest) GetIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateThreadTagRequest) SetId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateThreadTagRequest) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


