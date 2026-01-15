# RichEmbed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Author** | Pointer to [**RichEmbedAuthor**](RichEmbedAuthor.md) |  | [optional] 
**Image** | Pointer to [**RichEmbedImage**](RichEmbedImage.md) |  | [optional] 
**Thumbnail** | Pointer to [**RichEmbedThumbnail**](RichEmbedThumbnail.md) |  | [optional] 
**Footer** | Pointer to [**RichEmbedFooter**](RichEmbedFooter.md) |  | [optional] 
**Fields** | Pointer to [**[]RichEmbedField**](RichEmbedField.md) |  | [optional] 
**Provider** | Pointer to [**RichEmbedProvider**](RichEmbedProvider.md) |  | [optional] 
**Video** | Pointer to [**RichEmbedVideo**](RichEmbedVideo.md) |  | [optional] 

## Methods

### NewRichEmbed

`func NewRichEmbed() *RichEmbed`

NewRichEmbed instantiates a new RichEmbed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRichEmbedWithDefaults

`func NewRichEmbedWithDefaults() *RichEmbed`

NewRichEmbedWithDefaults instantiates a new RichEmbed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RichEmbed) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RichEmbed) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RichEmbed) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RichEmbed) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *RichEmbed) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RichEmbed) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RichEmbed) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RichEmbed) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTitle

`func (o *RichEmbed) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RichEmbed) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RichEmbed) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RichEmbed) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetColor

`func (o *RichEmbed) GetColor() int32`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *RichEmbed) GetColorOk() (*int32, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *RichEmbed) SetColor(v int32)`

SetColor sets Color field to given value.

### HasColor

`func (o *RichEmbed) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetTimestamp

`func (o *RichEmbed) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RichEmbed) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RichEmbed) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RichEmbed) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDescription

`func (o *RichEmbed) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RichEmbed) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RichEmbed) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RichEmbed) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAuthor

`func (o *RichEmbed) GetAuthor() RichEmbedAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *RichEmbed) GetAuthorOk() (*RichEmbedAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *RichEmbed) SetAuthor(v RichEmbedAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *RichEmbed) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetImage

`func (o *RichEmbed) GetImage() RichEmbedImage`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *RichEmbed) GetImageOk() (*RichEmbedImage, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *RichEmbed) SetImage(v RichEmbedImage)`

SetImage sets Image field to given value.

### HasImage

`func (o *RichEmbed) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetThumbnail

`func (o *RichEmbed) GetThumbnail() RichEmbedThumbnail`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *RichEmbed) GetThumbnailOk() (*RichEmbedThumbnail, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *RichEmbed) SetThumbnail(v RichEmbedThumbnail)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *RichEmbed) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetFooter

`func (o *RichEmbed) GetFooter() RichEmbedFooter`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *RichEmbed) GetFooterOk() (*RichEmbedFooter, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *RichEmbed) SetFooter(v RichEmbedFooter)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *RichEmbed) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetFields

`func (o *RichEmbed) GetFields() []RichEmbedField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RichEmbed) GetFieldsOk() (*[]RichEmbedField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RichEmbed) SetFields(v []RichEmbedField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RichEmbed) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetProvider

`func (o *RichEmbed) GetProvider() RichEmbedProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *RichEmbed) GetProviderOk() (*RichEmbedProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *RichEmbed) SetProvider(v RichEmbedProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *RichEmbed) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetVideo

`func (o *RichEmbed) GetVideo() RichEmbedVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *RichEmbed) GetVideoOk() (*RichEmbedVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *RichEmbed) SetVideo(v RichEmbedVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *RichEmbed) HasVideo() bool`

HasVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


