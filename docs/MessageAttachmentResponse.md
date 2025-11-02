# MessageAttachmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Filename** | **string** |  | 
**Size** | **int32** |  | 
**Url** | **string** |  | 
**ProxyUrl** | **string** |  | 
**Width** | Pointer to **int32** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**DurationSecs** | Pointer to **float64** |  | [optional] 
**Waveform** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Ephemeral** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Application** | Pointer to [**ApplicationResponse**](ApplicationResponse.md) |  | [optional] 
**ClipCreatedAt** | Pointer to **time.Time** |  | [optional] 
**ClipParticipants** | Pointer to [**[]UserResponse**](UserResponse.md) |  | [optional] 

## Methods

### NewMessageAttachmentResponse

`func NewMessageAttachmentResponse(id string, filename string, size int32, url string, proxyUrl string, ) *MessageAttachmentResponse`

NewMessageAttachmentResponse instantiates a new MessageAttachmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageAttachmentResponseWithDefaults

`func NewMessageAttachmentResponseWithDefaults() *MessageAttachmentResponse`

NewMessageAttachmentResponseWithDefaults instantiates a new MessageAttachmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageAttachmentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageAttachmentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageAttachmentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetFilename

`func (o *MessageAttachmentResponse) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *MessageAttachmentResponse) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *MessageAttachmentResponse) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetSize

`func (o *MessageAttachmentResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MessageAttachmentResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MessageAttachmentResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetUrl

`func (o *MessageAttachmentResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MessageAttachmentResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MessageAttachmentResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetProxyUrl

`func (o *MessageAttachmentResponse) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *MessageAttachmentResponse) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *MessageAttachmentResponse) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.


### GetWidth

`func (o *MessageAttachmentResponse) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *MessageAttachmentResponse) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *MessageAttachmentResponse) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *MessageAttachmentResponse) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *MessageAttachmentResponse) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *MessageAttachmentResponse) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *MessageAttachmentResponse) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *MessageAttachmentResponse) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetDurationSecs

`func (o *MessageAttachmentResponse) GetDurationSecs() float64`

GetDurationSecs returns the DurationSecs field if non-nil, zero value otherwise.

### GetDurationSecsOk

`func (o *MessageAttachmentResponse) GetDurationSecsOk() (*float64, bool)`

GetDurationSecsOk returns a tuple with the DurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSecs

`func (o *MessageAttachmentResponse) SetDurationSecs(v float64)`

SetDurationSecs sets DurationSecs field to given value.

### HasDurationSecs

`func (o *MessageAttachmentResponse) HasDurationSecs() bool`

HasDurationSecs returns a boolean if a field has been set.

### GetWaveform

`func (o *MessageAttachmentResponse) GetWaveform() string`

GetWaveform returns the Waveform field if non-nil, zero value otherwise.

### GetWaveformOk

`func (o *MessageAttachmentResponse) GetWaveformOk() (*string, bool)`

GetWaveformOk returns a tuple with the Waveform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaveform

`func (o *MessageAttachmentResponse) SetWaveform(v string)`

SetWaveform sets Waveform field to given value.

### HasWaveform

`func (o *MessageAttachmentResponse) HasWaveform() bool`

HasWaveform returns a boolean if a field has been set.

### GetDescription

`func (o *MessageAttachmentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MessageAttachmentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MessageAttachmentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MessageAttachmentResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContentType

`func (o *MessageAttachmentResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MessageAttachmentResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MessageAttachmentResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MessageAttachmentResponse) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEphemeral

`func (o *MessageAttachmentResponse) GetEphemeral() bool`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *MessageAttachmentResponse) GetEphemeralOk() (*bool, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *MessageAttachmentResponse) SetEphemeral(v bool)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *MessageAttachmentResponse) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetTitle

`func (o *MessageAttachmentResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MessageAttachmentResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MessageAttachmentResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MessageAttachmentResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MessageAttachmentResponse) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MessageAttachmentResponse) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetApplication

`func (o *MessageAttachmentResponse) GetApplication() ApplicationResponse`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *MessageAttachmentResponse) GetApplicationOk() (*ApplicationResponse, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *MessageAttachmentResponse) SetApplication(v ApplicationResponse)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *MessageAttachmentResponse) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetClipCreatedAt

`func (o *MessageAttachmentResponse) GetClipCreatedAt() time.Time`

GetClipCreatedAt returns the ClipCreatedAt field if non-nil, zero value otherwise.

### GetClipCreatedAtOk

`func (o *MessageAttachmentResponse) GetClipCreatedAtOk() (*time.Time, bool)`

GetClipCreatedAtOk returns a tuple with the ClipCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClipCreatedAt

`func (o *MessageAttachmentResponse) SetClipCreatedAt(v time.Time)`

SetClipCreatedAt sets ClipCreatedAt field to given value.

### HasClipCreatedAt

`func (o *MessageAttachmentResponse) HasClipCreatedAt() bool`

HasClipCreatedAt returns a boolean if a field has been set.

### GetClipParticipants

`func (o *MessageAttachmentResponse) GetClipParticipants() []UserResponse`

GetClipParticipants returns the ClipParticipants field if non-nil, zero value otherwise.

### GetClipParticipantsOk

`func (o *MessageAttachmentResponse) GetClipParticipantsOk() (*[]UserResponse, bool)`

GetClipParticipantsOk returns a tuple with the ClipParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClipParticipants

`func (o *MessageAttachmentResponse) SetClipParticipants(v []UserResponse)`

SetClipParticipants sets ClipParticipants field to given value.

### HasClipParticipants

`func (o *MessageAttachmentResponse) HasClipParticipants() bool`

HasClipParticipants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


