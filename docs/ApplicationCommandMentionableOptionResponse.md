# ApplicationCommandMentionableOptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Name** | **string** |  | 
**NameLocalized** | Pointer to **NullableString** |  | [optional] 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Description** | **string** |  | 
**DescriptionLocalized** | Pointer to **NullableString** |  | [optional] 
**DescriptionLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Required** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewApplicationCommandMentionableOptionResponse

`func NewApplicationCommandMentionableOptionResponse(type_ int32, name string, description string, ) *ApplicationCommandMentionableOptionResponse`

NewApplicationCommandMentionableOptionResponse instantiates a new ApplicationCommandMentionableOptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandMentionableOptionResponseWithDefaults

`func NewApplicationCommandMentionableOptionResponseWithDefaults() *ApplicationCommandMentionableOptionResponse`

NewApplicationCommandMentionableOptionResponseWithDefaults instantiates a new ApplicationCommandMentionableOptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationCommandMentionableOptionResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandMentionableOptionResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandMentionableOptionResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetName

`func (o *ApplicationCommandMentionableOptionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandMentionableOptionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandMentionableOptionResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalized

`func (o *ApplicationCommandMentionableOptionResponse) GetNameLocalized() string`

GetNameLocalized returns the NameLocalized field if non-nil, zero value otherwise.

### GetNameLocalizedOk

`func (o *ApplicationCommandMentionableOptionResponse) GetNameLocalizedOk() (*string, bool)`

GetNameLocalizedOk returns a tuple with the NameLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalized

`func (o *ApplicationCommandMentionableOptionResponse) SetNameLocalized(v string)`

SetNameLocalized sets NameLocalized field to given value.

### HasNameLocalized

`func (o *ApplicationCommandMentionableOptionResponse) HasNameLocalized() bool`

HasNameLocalized returns a boolean if a field has been set.

### SetNameLocalizedNil

`func (o *ApplicationCommandMentionableOptionResponse) SetNameLocalizedNil(b bool)`

 SetNameLocalizedNil sets the value for NameLocalized to be an explicit nil

### UnsetNameLocalized
`func (o *ApplicationCommandMentionableOptionResponse) UnsetNameLocalized()`

UnsetNameLocalized ensures that no value is present for NameLocalized, not even an explicit nil
### GetNameLocalizations

`func (o *ApplicationCommandMentionableOptionResponse) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandMentionableOptionResponse) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandMentionableOptionResponse) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandMentionableOptionResponse) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationCommandMentionableOptionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCommandMentionableOptionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCommandMentionableOptionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionLocalized

`func (o *ApplicationCommandMentionableOptionResponse) GetDescriptionLocalized() string`

GetDescriptionLocalized returns the DescriptionLocalized field if non-nil, zero value otherwise.

### GetDescriptionLocalizedOk

`func (o *ApplicationCommandMentionableOptionResponse) GetDescriptionLocalizedOk() (*string, bool)`

GetDescriptionLocalizedOk returns a tuple with the DescriptionLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalized

`func (o *ApplicationCommandMentionableOptionResponse) SetDescriptionLocalized(v string)`

SetDescriptionLocalized sets DescriptionLocalized field to given value.

### HasDescriptionLocalized

`func (o *ApplicationCommandMentionableOptionResponse) HasDescriptionLocalized() bool`

HasDescriptionLocalized returns a boolean if a field has been set.

### SetDescriptionLocalizedNil

`func (o *ApplicationCommandMentionableOptionResponse) SetDescriptionLocalizedNil(b bool)`

 SetDescriptionLocalizedNil sets the value for DescriptionLocalized to be an explicit nil

### UnsetDescriptionLocalized
`func (o *ApplicationCommandMentionableOptionResponse) UnsetDescriptionLocalized()`

UnsetDescriptionLocalized ensures that no value is present for DescriptionLocalized, not even an explicit nil
### GetDescriptionLocalizations

`func (o *ApplicationCommandMentionableOptionResponse) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationCommandMentionableOptionResponse) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationCommandMentionableOptionResponse) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationCommandMentionableOptionResponse) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.

### GetRequired

`func (o *ApplicationCommandMentionableOptionResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ApplicationCommandMentionableOptionResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ApplicationCommandMentionableOptionResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ApplicationCommandMentionableOptionResponse) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *ApplicationCommandMentionableOptionResponse) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *ApplicationCommandMentionableOptionResponse) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


