# CreateOrUpdateThreadTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EmojiId** | Pointer to **string** |  | [optional] 
**EmojiName** | Pointer to **NullableString** |  | [optional] 
**Moderated** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewCreateOrUpdateThreadTagRequest

`func NewCreateOrUpdateThreadTagRequest(name string, ) *CreateOrUpdateThreadTagRequest`

NewCreateOrUpdateThreadTagRequest instantiates a new CreateOrUpdateThreadTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateThreadTagRequestWithDefaults

`func NewCreateOrUpdateThreadTagRequestWithDefaults() *CreateOrUpdateThreadTagRequest`

NewCreateOrUpdateThreadTagRequestWithDefaults instantiates a new CreateOrUpdateThreadTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrUpdateThreadTagRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateThreadTagRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateThreadTagRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmojiId

`func (o *CreateOrUpdateThreadTagRequest) GetEmojiId() string`

GetEmojiId returns the EmojiId field if non-nil, zero value otherwise.

### GetEmojiIdOk

`func (o *CreateOrUpdateThreadTagRequest) GetEmojiIdOk() (*string, bool)`

GetEmojiIdOk returns a tuple with the EmojiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiId

`func (o *CreateOrUpdateThreadTagRequest) SetEmojiId(v string)`

SetEmojiId sets EmojiId field to given value.

### HasEmojiId

`func (o *CreateOrUpdateThreadTagRequest) HasEmojiId() bool`

HasEmojiId returns a boolean if a field has been set.

### GetEmojiName

`func (o *CreateOrUpdateThreadTagRequest) GetEmojiName() string`

GetEmojiName returns the EmojiName field if non-nil, zero value otherwise.

### GetEmojiNameOk

`func (o *CreateOrUpdateThreadTagRequest) GetEmojiNameOk() (*string, bool)`

GetEmojiNameOk returns a tuple with the EmojiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmojiName

`func (o *CreateOrUpdateThreadTagRequest) SetEmojiName(v string)`

SetEmojiName sets EmojiName field to given value.

### HasEmojiName

`func (o *CreateOrUpdateThreadTagRequest) HasEmojiName() bool`

HasEmojiName returns a boolean if a field has been set.

### SetEmojiNameNil

`func (o *CreateOrUpdateThreadTagRequest) SetEmojiNameNil(b bool)`

 SetEmojiNameNil sets the value for EmojiName to be an explicit nil

### UnsetEmojiName
`func (o *CreateOrUpdateThreadTagRequest) UnsetEmojiName()`

UnsetEmojiName ensures that no value is present for EmojiName, not even an explicit nil
### GetModerated

`func (o *CreateOrUpdateThreadTagRequest) GetModerated() bool`

GetModerated returns the Moderated field if non-nil, zero value otherwise.

### GetModeratedOk

`func (o *CreateOrUpdateThreadTagRequest) GetModeratedOk() (*bool, bool)`

GetModeratedOk returns a tuple with the Moderated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerated

`func (o *CreateOrUpdateThreadTagRequest) SetModerated(v bool)`

SetModerated sets Moderated field to given value.

### HasModerated

`func (o *CreateOrUpdateThreadTagRequest) HasModerated() bool`

HasModerated returns a boolean if a field has been set.

### SetModeratedNil

`func (o *CreateOrUpdateThreadTagRequest) SetModeratedNil(b bool)`

 SetModeratedNil sets the value for Moderated to be an explicit nil

### UnsetModerated
`func (o *CreateOrUpdateThreadTagRequest) UnsetModerated()`

UnsetModerated ensures that no value is present for Moderated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


