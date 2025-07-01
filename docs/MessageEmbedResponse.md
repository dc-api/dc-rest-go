# MessageEmbedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Url** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Color** | Pointer to **NullableInt32** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Fields** | Pointer to [**[]MessageEmbedFieldResponse**](MessageEmbedFieldResponse.md) |  | [optional] 
**Author** | Pointer to [**NullableMessageEmbedAuthorResponse**](MessageEmbedAuthorResponse.md) |  | [optional] 
**Provider** | Pointer to [**NullableMessageEmbedProviderResponse**](MessageEmbedProviderResponse.md) |  | [optional] 
**Image** | Pointer to [**NullableMessageEmbedImageResponse**](MessageEmbedImageResponse.md) |  | [optional] 
**Thumbnail** | Pointer to [**NullableMessageEmbedImageResponse**](MessageEmbedImageResponse.md) |  | [optional] 
**Video** | Pointer to [**NullableMessageEmbedVideoResponse**](MessageEmbedVideoResponse.md) |  | [optional] 
**Footer** | Pointer to [**NullableMessageEmbedFooterResponse**](MessageEmbedFooterResponse.md) |  | [optional] 

## Methods

### NewMessageEmbedResponse

`func NewMessageEmbedResponse(type_ string, ) *MessageEmbedResponse`

NewMessageEmbedResponse instantiates a new MessageEmbedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageEmbedResponseWithDefaults

`func NewMessageEmbedResponseWithDefaults() *MessageEmbedResponse`

NewMessageEmbedResponseWithDefaults instantiates a new MessageEmbedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageEmbedResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageEmbedResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageEmbedResponse) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *MessageEmbedResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MessageEmbedResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MessageEmbedResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MessageEmbedResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MessageEmbedResponse) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MessageEmbedResponse) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetTitle

`func (o *MessageEmbedResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MessageEmbedResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MessageEmbedResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MessageEmbedResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MessageEmbedResponse) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MessageEmbedResponse) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *MessageEmbedResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MessageEmbedResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MessageEmbedResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MessageEmbedResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MessageEmbedResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MessageEmbedResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetColor

`func (o *MessageEmbedResponse) GetColor() int32`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *MessageEmbedResponse) GetColorOk() (*int32, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *MessageEmbedResponse) SetColor(v int32)`

SetColor sets Color field to given value.

### HasColor

`func (o *MessageEmbedResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *MessageEmbedResponse) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *MessageEmbedResponse) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetTimestamp

`func (o *MessageEmbedResponse) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageEmbedResponse) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageEmbedResponse) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MessageEmbedResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *MessageEmbedResponse) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *MessageEmbedResponse) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetFields

`func (o *MessageEmbedResponse) GetFields() []MessageEmbedFieldResponse`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MessageEmbedResponse) GetFieldsOk() (*[]MessageEmbedFieldResponse, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MessageEmbedResponse) SetFields(v []MessageEmbedFieldResponse)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MessageEmbedResponse) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *MessageEmbedResponse) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *MessageEmbedResponse) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetAuthor

`func (o *MessageEmbedResponse) GetAuthor() MessageEmbedAuthorResponse`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *MessageEmbedResponse) GetAuthorOk() (*MessageEmbedAuthorResponse, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *MessageEmbedResponse) SetAuthor(v MessageEmbedAuthorResponse)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *MessageEmbedResponse) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *MessageEmbedResponse) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *MessageEmbedResponse) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetProvider

`func (o *MessageEmbedResponse) GetProvider() MessageEmbedProviderResponse`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *MessageEmbedResponse) GetProviderOk() (*MessageEmbedProviderResponse, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *MessageEmbedResponse) SetProvider(v MessageEmbedProviderResponse)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *MessageEmbedResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *MessageEmbedResponse) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *MessageEmbedResponse) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetImage

`func (o *MessageEmbedResponse) GetImage() MessageEmbedImageResponse`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *MessageEmbedResponse) GetImageOk() (*MessageEmbedImageResponse, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *MessageEmbedResponse) SetImage(v MessageEmbedImageResponse)`

SetImage sets Image field to given value.

### HasImage

`func (o *MessageEmbedResponse) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *MessageEmbedResponse) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *MessageEmbedResponse) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetThumbnail

`func (o *MessageEmbedResponse) GetThumbnail() MessageEmbedImageResponse`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *MessageEmbedResponse) GetThumbnailOk() (*MessageEmbedImageResponse, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *MessageEmbedResponse) SetThumbnail(v MessageEmbedImageResponse)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *MessageEmbedResponse) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### SetThumbnailNil

`func (o *MessageEmbedResponse) SetThumbnailNil(b bool)`

 SetThumbnailNil sets the value for Thumbnail to be an explicit nil

### UnsetThumbnail
`func (o *MessageEmbedResponse) UnsetThumbnail()`

UnsetThumbnail ensures that no value is present for Thumbnail, not even an explicit nil
### GetVideo

`func (o *MessageEmbedResponse) GetVideo() MessageEmbedVideoResponse`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *MessageEmbedResponse) GetVideoOk() (*MessageEmbedVideoResponse, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *MessageEmbedResponse) SetVideo(v MessageEmbedVideoResponse)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *MessageEmbedResponse) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### SetVideoNil

`func (o *MessageEmbedResponse) SetVideoNil(b bool)`

 SetVideoNil sets the value for Video to be an explicit nil

### UnsetVideo
`func (o *MessageEmbedResponse) UnsetVideo()`

UnsetVideo ensures that no value is present for Video, not even an explicit nil
### GetFooter

`func (o *MessageEmbedResponse) GetFooter() MessageEmbedFooterResponse`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *MessageEmbedResponse) GetFooterOk() (*MessageEmbedFooterResponse, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *MessageEmbedResponse) SetFooter(v MessageEmbedFooterResponse)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *MessageEmbedResponse) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### SetFooterNil

`func (o *MessageEmbedResponse) SetFooterNil(b bool)`

 SetFooterNil sets the value for Footer to be an explicit nil

### UnsetFooter
`func (o *MessageEmbedResponse) UnsetFooter()`

UnsetFooter ensures that no value is present for Footer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


