# PollEmojiCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Animated** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewPollEmojiCreateRequest

`func NewPollEmojiCreateRequest() *PollEmojiCreateRequest`

NewPollEmojiCreateRequest instantiates a new PollEmojiCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollEmojiCreateRequestWithDefaults

`func NewPollEmojiCreateRequestWithDefaults() *PollEmojiCreateRequest`

NewPollEmojiCreateRequestWithDefaults instantiates a new PollEmojiCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PollEmojiCreateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PollEmojiCreateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PollEmojiCreateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PollEmojiCreateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PollEmojiCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PollEmojiCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PollEmojiCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PollEmojiCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PollEmojiCreateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PollEmojiCreateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAnimated

`func (o *PollEmojiCreateRequest) GetAnimated() bool`

GetAnimated returns the Animated field if non-nil, zero value otherwise.

### GetAnimatedOk

`func (o *PollEmojiCreateRequest) GetAnimatedOk() (*bool, bool)`

GetAnimatedOk returns a tuple with the Animated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimated

`func (o *PollEmojiCreateRequest) SetAnimated(v bool)`

SetAnimated sets Animated field to given value.

### HasAnimated

`func (o *PollEmojiCreateRequest) HasAnimated() bool`

HasAnimated returns a boolean if a field has been set.

### SetAnimatedNil

`func (o *PollEmojiCreateRequest) SetAnimatedNil(b bool)`

 SetAnimatedNil sets the value for Animated to be an explicit nil

### UnsetAnimated
`func (o *PollEmojiCreateRequest) UnsetAnimated()`

UnsetAnimated ensures that no value is present for Animated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


