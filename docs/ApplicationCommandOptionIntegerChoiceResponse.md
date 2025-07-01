# ApplicationCommandOptionIntegerChoiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**NameLocalized** | Pointer to **NullableString** |  | [optional] 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Value** | **int64** |  | 

## Methods

### NewApplicationCommandOptionIntegerChoiceResponse

`func NewApplicationCommandOptionIntegerChoiceResponse(name string, value int64, ) *ApplicationCommandOptionIntegerChoiceResponse`

NewApplicationCommandOptionIntegerChoiceResponse instantiates a new ApplicationCommandOptionIntegerChoiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandOptionIntegerChoiceResponseWithDefaults

`func NewApplicationCommandOptionIntegerChoiceResponseWithDefaults() *ApplicationCommandOptionIntegerChoiceResponse`

NewApplicationCommandOptionIntegerChoiceResponseWithDefaults instantiates a new ApplicationCommandOptionIntegerChoiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationCommandOptionIntegerChoiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCommandOptionIntegerChoiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCommandOptionIntegerChoiceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalized

`func (o *ApplicationCommandOptionIntegerChoiceResponse) GetNameLocalized() string`

GetNameLocalized returns the NameLocalized field if non-nil, zero value otherwise.

### GetNameLocalizedOk

`func (o *ApplicationCommandOptionIntegerChoiceResponse) GetNameLocalizedOk() (*string, bool)`

GetNameLocalizedOk returns a tuple with the NameLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalized

`func (o *ApplicationCommandOptionIntegerChoiceResponse) SetNameLocalized(v string)`

SetNameLocalized sets NameLocalized field to given value.

### HasNameLocalized

`func (o *ApplicationCommandOptionIntegerChoiceResponse) HasNameLocalized() bool`

HasNameLocalized returns a boolean if a field has been set.

### SetNameLocalizedNil

`func (o *ApplicationCommandOptionIntegerChoiceResponse) SetNameLocalizedNil(b bool)`

 SetNameLocalizedNil sets the value for NameLocalized to be an explicit nil

### UnsetNameLocalized
`func (o *ApplicationCommandOptionIntegerChoiceResponse) UnsetNameLocalized()`

UnsetNameLocalized ensures that no value is present for NameLocalized, not even an explicit nil
### GetNameLocalizations

`func (o *ApplicationCommandOptionIntegerChoiceResponse) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationCommandOptionIntegerChoiceResponse) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationCommandOptionIntegerChoiceResponse) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationCommandOptionIntegerChoiceResponse) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetValue

`func (o *ApplicationCommandOptionIntegerChoiceResponse) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApplicationCommandOptionIntegerChoiceResponse) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApplicationCommandOptionIntegerChoiceResponse) SetValue(v int64)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


