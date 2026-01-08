# UpdateOnboardingPromptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Options** | [**[]OnboardingPromptOptionRequest**](OnboardingPromptOptionRequest.md) |  | 
**SingleSelect** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**InOnboarding** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**UpdateOnboardingPromptRequestType**](UpdateOnboardingPromptRequestType.md) |  | [optional] 
**Id** | **string** |  | 

## Methods

### NewUpdateOnboardingPromptRequest

`func NewUpdateOnboardingPromptRequest(title string, options []OnboardingPromptOptionRequest, id string, ) *UpdateOnboardingPromptRequest`

NewUpdateOnboardingPromptRequest instantiates a new UpdateOnboardingPromptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOnboardingPromptRequestWithDefaults

`func NewUpdateOnboardingPromptRequestWithDefaults() *UpdateOnboardingPromptRequest`

NewUpdateOnboardingPromptRequestWithDefaults instantiates a new UpdateOnboardingPromptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateOnboardingPromptRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateOnboardingPromptRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateOnboardingPromptRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetOptions

`func (o *UpdateOnboardingPromptRequest) GetOptions() []OnboardingPromptOptionRequest`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateOnboardingPromptRequest) GetOptionsOk() (*[]OnboardingPromptOptionRequest, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateOnboardingPromptRequest) SetOptions(v []OnboardingPromptOptionRequest)`

SetOptions sets Options field to given value.


### GetSingleSelect

`func (o *UpdateOnboardingPromptRequest) GetSingleSelect() bool`

GetSingleSelect returns the SingleSelect field if non-nil, zero value otherwise.

### GetSingleSelectOk

`func (o *UpdateOnboardingPromptRequest) GetSingleSelectOk() (*bool, bool)`

GetSingleSelectOk returns a tuple with the SingleSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSelect

`func (o *UpdateOnboardingPromptRequest) SetSingleSelect(v bool)`

SetSingleSelect sets SingleSelect field to given value.

### HasSingleSelect

`func (o *UpdateOnboardingPromptRequest) HasSingleSelect() bool`

HasSingleSelect returns a boolean if a field has been set.

### GetRequired

`func (o *UpdateOnboardingPromptRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *UpdateOnboardingPromptRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *UpdateOnboardingPromptRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *UpdateOnboardingPromptRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetInOnboarding

`func (o *UpdateOnboardingPromptRequest) GetInOnboarding() bool`

GetInOnboarding returns the InOnboarding field if non-nil, zero value otherwise.

### GetInOnboardingOk

`func (o *UpdateOnboardingPromptRequest) GetInOnboardingOk() (*bool, bool)`

GetInOnboardingOk returns a tuple with the InOnboarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInOnboarding

`func (o *UpdateOnboardingPromptRequest) SetInOnboarding(v bool)`

SetInOnboarding sets InOnboarding field to given value.

### HasInOnboarding

`func (o *UpdateOnboardingPromptRequest) HasInOnboarding() bool`

HasInOnboarding returns a boolean if a field has been set.

### GetType

`func (o *UpdateOnboardingPromptRequest) GetType() UpdateOnboardingPromptRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateOnboardingPromptRequest) GetTypeOk() (*UpdateOnboardingPromptRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateOnboardingPromptRequest) SetType(v UpdateOnboardingPromptRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateOnboardingPromptRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *UpdateOnboardingPromptRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOnboardingPromptRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOnboardingPromptRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


