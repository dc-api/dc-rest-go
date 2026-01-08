# OnboardingPromptResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Options** | [**[]OnboardingPromptOptionResponse**](OnboardingPromptOptionResponse.md) |  | 
**SingleSelect** | **bool** |  | 
**Required** | **bool** |  | 
**InOnboarding** | **bool** |  | 
**Type** | [**OnboardingPromptType**](OnboardingPromptType.md) |  | 

## Methods

### NewOnboardingPromptResponse

`func NewOnboardingPromptResponse(id string, title string, options []OnboardingPromptOptionResponse, singleSelect bool, required bool, inOnboarding bool, type_ OnboardingPromptType, ) *OnboardingPromptResponse`

NewOnboardingPromptResponse instantiates a new OnboardingPromptResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingPromptResponseWithDefaults

`func NewOnboardingPromptResponseWithDefaults() *OnboardingPromptResponse`

NewOnboardingPromptResponseWithDefaults instantiates a new OnboardingPromptResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnboardingPromptResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnboardingPromptResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnboardingPromptResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *OnboardingPromptResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OnboardingPromptResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OnboardingPromptResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetOptions

`func (o *OnboardingPromptResponse) GetOptions() []OnboardingPromptOptionResponse`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OnboardingPromptResponse) GetOptionsOk() (*[]OnboardingPromptOptionResponse, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OnboardingPromptResponse) SetOptions(v []OnboardingPromptOptionResponse)`

SetOptions sets Options field to given value.


### GetSingleSelect

`func (o *OnboardingPromptResponse) GetSingleSelect() bool`

GetSingleSelect returns the SingleSelect field if non-nil, zero value otherwise.

### GetSingleSelectOk

`func (o *OnboardingPromptResponse) GetSingleSelectOk() (*bool, bool)`

GetSingleSelectOk returns a tuple with the SingleSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSelect

`func (o *OnboardingPromptResponse) SetSingleSelect(v bool)`

SetSingleSelect sets SingleSelect field to given value.


### GetRequired

`func (o *OnboardingPromptResponse) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *OnboardingPromptResponse) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *OnboardingPromptResponse) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetInOnboarding

`func (o *OnboardingPromptResponse) GetInOnboarding() bool`

GetInOnboarding returns the InOnboarding field if non-nil, zero value otherwise.

### GetInOnboardingOk

`func (o *OnboardingPromptResponse) GetInOnboardingOk() (*bool, bool)`

GetInOnboardingOk returns a tuple with the InOnboarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInOnboarding

`func (o *OnboardingPromptResponse) SetInOnboarding(v bool)`

SetInOnboarding sets InOnboarding field to given value.


### GetType

`func (o *OnboardingPromptResponse) GetType() OnboardingPromptType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OnboardingPromptResponse) GetTypeOk() (*OnboardingPromptType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OnboardingPromptResponse) SetType(v OnboardingPromptType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


