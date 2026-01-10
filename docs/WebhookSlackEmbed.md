# WebhookSlackEmbed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**TitleLink** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Ts** | Pointer to **int32** |  | [optional] 
**Pretext** | Pointer to **string** |  | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**FooterIcon** | Pointer to **string** |  | [optional] 
**AuthorName** | Pointer to **string** |  | [optional] 
**AuthorLink** | Pointer to **string** |  | [optional] 
**AuthorIcon** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**ThumbUrl** | Pointer to **string** |  | [optional] 
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


