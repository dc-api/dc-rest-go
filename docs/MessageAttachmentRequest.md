# MessageAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Filename** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DurationSecs** | Pointer to **NullableFloat64** |  | [optional] 
**Waveform** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**IsRemix** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMessageAttachmentRequest

`func NewMessageAttachmentRequest(id string, ) *MessageAttachmentRequest`

NewMessageAttachmentRequest instantiates a new MessageAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageAttachmentRequestWithDefaults

`func NewMessageAttachmentRequestWithDefaults() *MessageAttachmentRequest`

NewMessageAttachmentRequestWithDefaults instantiates a new MessageAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageAttachmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageAttachmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageAttachmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetFilename

`func (o *MessageAttachmentRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *MessageAttachmentRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *MessageAttachmentRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *MessageAttachmentRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### SetFilenameNil

`func (o *MessageAttachmentRequest) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *MessageAttachmentRequest) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetDescription

`func (o *MessageAttachmentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MessageAttachmentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MessageAttachmentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MessageAttachmentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MessageAttachmentRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MessageAttachmentRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDurationSecs

`func (o *MessageAttachmentRequest) GetDurationSecs() float64`

GetDurationSecs returns the DurationSecs field if non-nil, zero value otherwise.

### GetDurationSecsOk

`func (o *MessageAttachmentRequest) GetDurationSecsOk() (*float64, bool)`

GetDurationSecsOk returns a tuple with the DurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSecs

`func (o *MessageAttachmentRequest) SetDurationSecs(v float64)`

SetDurationSecs sets DurationSecs field to given value.

### HasDurationSecs

`func (o *MessageAttachmentRequest) HasDurationSecs() bool`

HasDurationSecs returns a boolean if a field has been set.

### SetDurationSecsNil

`func (o *MessageAttachmentRequest) SetDurationSecsNil(b bool)`

 SetDurationSecsNil sets the value for DurationSecs to be an explicit nil

### UnsetDurationSecs
`func (o *MessageAttachmentRequest) UnsetDurationSecs()`

UnsetDurationSecs ensures that no value is present for DurationSecs, not even an explicit nil
### GetWaveform

`func (o *MessageAttachmentRequest) GetWaveform() string`

GetWaveform returns the Waveform field if non-nil, zero value otherwise.

### GetWaveformOk

`func (o *MessageAttachmentRequest) GetWaveformOk() (*string, bool)`

GetWaveformOk returns a tuple with the Waveform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaveform

`func (o *MessageAttachmentRequest) SetWaveform(v string)`

SetWaveform sets Waveform field to given value.

### HasWaveform

`func (o *MessageAttachmentRequest) HasWaveform() bool`

HasWaveform returns a boolean if a field has been set.

### SetWaveformNil

`func (o *MessageAttachmentRequest) SetWaveformNil(b bool)`

 SetWaveformNil sets the value for Waveform to be an explicit nil

### UnsetWaveform
`func (o *MessageAttachmentRequest) UnsetWaveform()`

UnsetWaveform ensures that no value is present for Waveform, not even an explicit nil
### GetTitle

`func (o *MessageAttachmentRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MessageAttachmentRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MessageAttachmentRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MessageAttachmentRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MessageAttachmentRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MessageAttachmentRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetIsRemix

`func (o *MessageAttachmentRequest) GetIsRemix() bool`

GetIsRemix returns the IsRemix field if non-nil, zero value otherwise.

### GetIsRemixOk

`func (o *MessageAttachmentRequest) GetIsRemixOk() (*bool, bool)`

GetIsRemixOk returns a tuple with the IsRemix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemix

`func (o *MessageAttachmentRequest) SetIsRemix(v bool)`

SetIsRemix sets IsRemix field to given value.

### HasIsRemix

`func (o *MessageAttachmentRequest) HasIsRemix() bool`

HasIsRemix returns a boolean if a field has been set.

### SetIsRemixNil

`func (o *MessageAttachmentRequest) SetIsRemixNil(b bool)`

 SetIsRemixNil sets the value for IsRemix to be an explicit nil

### UnsetIsRemix
`func (o *MessageAttachmentRequest) UnsetIsRemix()`

UnsetIsRemix ensures that no value is present for IsRemix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


