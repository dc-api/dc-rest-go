# SettingsEmojiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Animated** | **bool** |  | 

## Methods

### NewSettingsEmojiResponse

`func NewSettingsEmojiResponse(animated bool, ) *SettingsEmojiResponse`

NewSettingsEmojiResponse instantiates a new SettingsEmojiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsEmojiResponseWithDefaults

`func NewSettingsEmojiResponseWithDefaults() *SettingsEmojiResponse`

NewSettingsEmojiResponseWithDefaults instantiates a new SettingsEmojiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SettingsEmojiResponse) GetId() GetEntitlementsSkuIdsParameterOneOfInner`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SettingsEmojiResponse) GetIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SettingsEmojiResponse) SetId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetId sets Id field to given value.

### HasId

`func (o *SettingsEmojiResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SettingsEmojiResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SettingsEmojiResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SettingsEmojiResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SettingsEmojiResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnimated

`func (o *SettingsEmojiResponse) GetAnimated() bool`

GetAnimated returns the Animated field if non-nil, zero value otherwise.

### GetAnimatedOk

`func (o *SettingsEmojiResponse) GetAnimatedOk() (*bool, bool)`

GetAnimatedOk returns a tuple with the Animated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnimated

`func (o *SettingsEmojiResponse) SetAnimated(v bool)`

SetAnimated sets Animated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


