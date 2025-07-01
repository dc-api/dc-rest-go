# WebhookSlackEmbed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**TitleLink** | Pointer to **NullableString** |  | [optional] 
**Text** | Pointer to **NullableString** |  | [optional] 
**Color** | Pointer to **NullableString** |  | [optional] 
**Ts** | Pointer to **NullableInt32** |  | [optional] 
**Pretext** | Pointer to **NullableString** |  | [optional] 
**Footer** | Pointer to **NullableString** |  | [optional] 
**FooterIcon** | Pointer to **NullableString** |  | [optional] 
**AuthorName** | Pointer to **NullableString** |  | [optional] 
**AuthorLink** | Pointer to **NullableString** |  | [optional] 
**AuthorIcon** | Pointer to **NullableString** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**ThumbUrl** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to [**[]WebhookSlackEmbedField**](WebhookSlackEmbedField.md) |  | [optional] 

## Methods

### NewWebhookSlackEmbed

`func NewWebhookSlackEmbed() *WebhookSlackEmbed`

NewWebhookSlackEmbed instantiates a new WebhookSlackEmbed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSlackEmbedWithDefaults

`func NewWebhookSlackEmbedWithDefaults() *WebhookSlackEmbed`

NewWebhookSlackEmbedWithDefaults instantiates a new WebhookSlackEmbed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *WebhookSlackEmbed) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebhookSlackEmbed) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebhookSlackEmbed) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebhookSlackEmbed) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *WebhookSlackEmbed) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *WebhookSlackEmbed) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTitleLink

`func (o *WebhookSlackEmbed) GetTitleLink() string`

GetTitleLink returns the TitleLink field if non-nil, zero value otherwise.

### GetTitleLinkOk

`func (o *WebhookSlackEmbed) GetTitleLinkOk() (*string, bool)`

GetTitleLinkOk returns a tuple with the TitleLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleLink

`func (o *WebhookSlackEmbed) SetTitleLink(v string)`

SetTitleLink sets TitleLink field to given value.

### HasTitleLink

`func (o *WebhookSlackEmbed) HasTitleLink() bool`

HasTitleLink returns a boolean if a field has been set.

### SetTitleLinkNil

`func (o *WebhookSlackEmbed) SetTitleLinkNil(b bool)`

 SetTitleLinkNil sets the value for TitleLink to be an explicit nil

### UnsetTitleLink
`func (o *WebhookSlackEmbed) UnsetTitleLink()`

UnsetTitleLink ensures that no value is present for TitleLink, not even an explicit nil
### GetText

`func (o *WebhookSlackEmbed) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WebhookSlackEmbed) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WebhookSlackEmbed) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WebhookSlackEmbed) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *WebhookSlackEmbed) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *WebhookSlackEmbed) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetColor

`func (o *WebhookSlackEmbed) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *WebhookSlackEmbed) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *WebhookSlackEmbed) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *WebhookSlackEmbed) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *WebhookSlackEmbed) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *WebhookSlackEmbed) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetTs

`func (o *WebhookSlackEmbed) GetTs() int32`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *WebhookSlackEmbed) GetTsOk() (*int32, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *WebhookSlackEmbed) SetTs(v int32)`

SetTs sets Ts field to given value.

### HasTs

`func (o *WebhookSlackEmbed) HasTs() bool`

HasTs returns a boolean if a field has been set.

### SetTsNil

`func (o *WebhookSlackEmbed) SetTsNil(b bool)`

 SetTsNil sets the value for Ts to be an explicit nil

### UnsetTs
`func (o *WebhookSlackEmbed) UnsetTs()`

UnsetTs ensures that no value is present for Ts, not even an explicit nil
### GetPretext

`func (o *WebhookSlackEmbed) GetPretext() string`

GetPretext returns the Pretext field if non-nil, zero value otherwise.

### GetPretextOk

`func (o *WebhookSlackEmbed) GetPretextOk() (*string, bool)`

GetPretextOk returns a tuple with the Pretext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPretext

`func (o *WebhookSlackEmbed) SetPretext(v string)`

SetPretext sets Pretext field to given value.

### HasPretext

`func (o *WebhookSlackEmbed) HasPretext() bool`

HasPretext returns a boolean if a field has been set.

### SetPretextNil

`func (o *WebhookSlackEmbed) SetPretextNil(b bool)`

 SetPretextNil sets the value for Pretext to be an explicit nil

### UnsetPretext
`func (o *WebhookSlackEmbed) UnsetPretext()`

UnsetPretext ensures that no value is present for Pretext, not even an explicit nil
### GetFooter

`func (o *WebhookSlackEmbed) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *WebhookSlackEmbed) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *WebhookSlackEmbed) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *WebhookSlackEmbed) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### SetFooterNil

`func (o *WebhookSlackEmbed) SetFooterNil(b bool)`

 SetFooterNil sets the value for Footer to be an explicit nil

### UnsetFooter
`func (o *WebhookSlackEmbed) UnsetFooter()`

UnsetFooter ensures that no value is present for Footer, not even an explicit nil
### GetFooterIcon

`func (o *WebhookSlackEmbed) GetFooterIcon() string`

GetFooterIcon returns the FooterIcon field if non-nil, zero value otherwise.

### GetFooterIconOk

`func (o *WebhookSlackEmbed) GetFooterIconOk() (*string, bool)`

GetFooterIconOk returns a tuple with the FooterIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterIcon

`func (o *WebhookSlackEmbed) SetFooterIcon(v string)`

SetFooterIcon sets FooterIcon field to given value.

### HasFooterIcon

`func (o *WebhookSlackEmbed) HasFooterIcon() bool`

HasFooterIcon returns a boolean if a field has been set.

### SetFooterIconNil

`func (o *WebhookSlackEmbed) SetFooterIconNil(b bool)`

 SetFooterIconNil sets the value for FooterIcon to be an explicit nil

### UnsetFooterIcon
`func (o *WebhookSlackEmbed) UnsetFooterIcon()`

UnsetFooterIcon ensures that no value is present for FooterIcon, not even an explicit nil
### GetAuthorName

`func (o *WebhookSlackEmbed) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *WebhookSlackEmbed) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *WebhookSlackEmbed) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *WebhookSlackEmbed) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### SetAuthorNameNil

`func (o *WebhookSlackEmbed) SetAuthorNameNil(b bool)`

 SetAuthorNameNil sets the value for AuthorName to be an explicit nil

### UnsetAuthorName
`func (o *WebhookSlackEmbed) UnsetAuthorName()`

UnsetAuthorName ensures that no value is present for AuthorName, not even an explicit nil
### GetAuthorLink

`func (o *WebhookSlackEmbed) GetAuthorLink() string`

GetAuthorLink returns the AuthorLink field if non-nil, zero value otherwise.

### GetAuthorLinkOk

`func (o *WebhookSlackEmbed) GetAuthorLinkOk() (*string, bool)`

GetAuthorLinkOk returns a tuple with the AuthorLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorLink

`func (o *WebhookSlackEmbed) SetAuthorLink(v string)`

SetAuthorLink sets AuthorLink field to given value.

### HasAuthorLink

`func (o *WebhookSlackEmbed) HasAuthorLink() bool`

HasAuthorLink returns a boolean if a field has been set.

### SetAuthorLinkNil

`func (o *WebhookSlackEmbed) SetAuthorLinkNil(b bool)`

 SetAuthorLinkNil sets the value for AuthorLink to be an explicit nil

### UnsetAuthorLink
`func (o *WebhookSlackEmbed) UnsetAuthorLink()`

UnsetAuthorLink ensures that no value is present for AuthorLink, not even an explicit nil
### GetAuthorIcon

`func (o *WebhookSlackEmbed) GetAuthorIcon() string`

GetAuthorIcon returns the AuthorIcon field if non-nil, zero value otherwise.

### GetAuthorIconOk

`func (o *WebhookSlackEmbed) GetAuthorIconOk() (*string, bool)`

GetAuthorIconOk returns a tuple with the AuthorIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorIcon

`func (o *WebhookSlackEmbed) SetAuthorIcon(v string)`

SetAuthorIcon sets AuthorIcon field to given value.

### HasAuthorIcon

`func (o *WebhookSlackEmbed) HasAuthorIcon() bool`

HasAuthorIcon returns a boolean if a field has been set.

### SetAuthorIconNil

`func (o *WebhookSlackEmbed) SetAuthorIconNil(b bool)`

 SetAuthorIconNil sets the value for AuthorIcon to be an explicit nil

### UnsetAuthorIcon
`func (o *WebhookSlackEmbed) UnsetAuthorIcon()`

UnsetAuthorIcon ensures that no value is present for AuthorIcon, not even an explicit nil
### GetImageUrl

`func (o *WebhookSlackEmbed) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *WebhookSlackEmbed) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *WebhookSlackEmbed) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *WebhookSlackEmbed) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *WebhookSlackEmbed) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *WebhookSlackEmbed) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetThumbUrl

`func (o *WebhookSlackEmbed) GetThumbUrl() string`

GetThumbUrl returns the ThumbUrl field if non-nil, zero value otherwise.

### GetThumbUrlOk

`func (o *WebhookSlackEmbed) GetThumbUrlOk() (*string, bool)`

GetThumbUrlOk returns a tuple with the ThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbUrl

`func (o *WebhookSlackEmbed) SetThumbUrl(v string)`

SetThumbUrl sets ThumbUrl field to given value.

### HasThumbUrl

`func (o *WebhookSlackEmbed) HasThumbUrl() bool`

HasThumbUrl returns a boolean if a field has been set.

### SetThumbUrlNil

`func (o *WebhookSlackEmbed) SetThumbUrlNil(b bool)`

 SetThumbUrlNil sets the value for ThumbUrl to be an explicit nil

### UnsetThumbUrl
`func (o *WebhookSlackEmbed) UnsetThumbUrl()`

UnsetThumbUrl ensures that no value is present for ThumbUrl, not even an explicit nil
### GetFields

`func (o *WebhookSlackEmbed) GetFields() []WebhookSlackEmbedField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WebhookSlackEmbed) GetFieldsOk() (*[]WebhookSlackEmbedField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WebhookSlackEmbed) SetFields(v []WebhookSlackEmbedField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *WebhookSlackEmbed) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *WebhookSlackEmbed) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *WebhookSlackEmbed) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


