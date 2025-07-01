# CreateGuildFromTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Icon** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateGuildFromTemplateRequest

`func NewCreateGuildFromTemplateRequest(name string, ) *CreateGuildFromTemplateRequest`

NewCreateGuildFromTemplateRequest instantiates a new CreateGuildFromTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGuildFromTemplateRequestWithDefaults

`func NewCreateGuildFromTemplateRequestWithDefaults() *CreateGuildFromTemplateRequest`

NewCreateGuildFromTemplateRequestWithDefaults instantiates a new CreateGuildFromTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateGuildFromTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGuildFromTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGuildFromTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIcon

`func (o *CreateGuildFromTemplateRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateGuildFromTemplateRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateGuildFromTemplateRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateGuildFromTemplateRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *CreateGuildFromTemplateRequest) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *CreateGuildFromTemplateRequest) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


