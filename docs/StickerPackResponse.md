# StickerPackResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SkuId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Stickers** | [**[]StandardStickerResponse**](StandardStickerResponse.md) |  | 
**CoverStickerId** | Pointer to **string** |  | [optional] 
**BannerAssetId** | Pointer to **string** |  | [optional] 

## Methods

### NewStickerPackResponse

`func NewStickerPackResponse(id string, skuId string, name string, stickers []StandardStickerResponse, ) *StickerPackResponse`

NewStickerPackResponse instantiates a new StickerPackResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStickerPackResponseWithDefaults

`func NewStickerPackResponseWithDefaults() *StickerPackResponse`

NewStickerPackResponseWithDefaults instantiates a new StickerPackResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StickerPackResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StickerPackResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StickerPackResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSkuId

`func (o *StickerPackResponse) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *StickerPackResponse) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *StickerPackResponse) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.


### GetName

`func (o *StickerPackResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StickerPackResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StickerPackResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *StickerPackResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StickerPackResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StickerPackResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StickerPackResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StickerPackResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StickerPackResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStickers

`func (o *StickerPackResponse) GetStickers() []StandardStickerResponse`

GetStickers returns the Stickers field if non-nil, zero value otherwise.

### GetStickersOk

`func (o *StickerPackResponse) GetStickersOk() (*[]StandardStickerResponse, bool)`

GetStickersOk returns a tuple with the Stickers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickers

`func (o *StickerPackResponse) SetStickers(v []StandardStickerResponse)`

SetStickers sets Stickers field to given value.


### GetCoverStickerId

`func (o *StickerPackResponse) GetCoverStickerId() string`

GetCoverStickerId returns the CoverStickerId field if non-nil, zero value otherwise.

### GetCoverStickerIdOk

`func (o *StickerPackResponse) GetCoverStickerIdOk() (*string, bool)`

GetCoverStickerIdOk returns a tuple with the CoverStickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverStickerId

`func (o *StickerPackResponse) SetCoverStickerId(v string)`

SetCoverStickerId sets CoverStickerId field to given value.

### HasCoverStickerId

`func (o *StickerPackResponse) HasCoverStickerId() bool`

HasCoverStickerId returns a boolean if a field has been set.

### GetBannerAssetId

`func (o *StickerPackResponse) GetBannerAssetId() string`

GetBannerAssetId returns the BannerAssetId field if non-nil, zero value otherwise.

### GetBannerAssetIdOk

`func (o *StickerPackResponse) GetBannerAssetIdOk() (*string, bool)`

GetBannerAssetIdOk returns a tuple with the BannerAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerAssetId

`func (o *StickerPackResponse) SetBannerAssetId(v string)`

SetBannerAssetId sets BannerAssetId field to given value.

### HasBannerAssetId

`func (o *StickerPackResponse) HasBannerAssetId() bool`

HasBannerAssetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


