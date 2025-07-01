# PurchaseNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **NullableInt32** |  | 
**GuildProductPurchase** | Pointer to [**NullableGuildProductPurchaseResponse**](GuildProductPurchaseResponse.md) |  | [optional] 

## Methods

### NewPurchaseNotificationResponse

`func NewPurchaseNotificationResponse(type_ NullableInt32, ) *PurchaseNotificationResponse`

NewPurchaseNotificationResponse instantiates a new PurchaseNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseNotificationResponseWithDefaults

`func NewPurchaseNotificationResponseWithDefaults() *PurchaseNotificationResponse`

NewPurchaseNotificationResponseWithDefaults instantiates a new PurchaseNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PurchaseNotificationResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PurchaseNotificationResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PurchaseNotificationResponse) SetType(v int32)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PurchaseNotificationResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PurchaseNotificationResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetGuildProductPurchase

`func (o *PurchaseNotificationResponse) GetGuildProductPurchase() GuildProductPurchaseResponse`

GetGuildProductPurchase returns the GuildProductPurchase field if non-nil, zero value otherwise.

### GetGuildProductPurchaseOk

`func (o *PurchaseNotificationResponse) GetGuildProductPurchaseOk() (*GuildProductPurchaseResponse, bool)`

GetGuildProductPurchaseOk returns a tuple with the GuildProductPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildProductPurchase

`func (o *PurchaseNotificationResponse) SetGuildProductPurchase(v GuildProductPurchaseResponse)`

SetGuildProductPurchase sets GuildProductPurchase field to given value.

### HasGuildProductPurchase

`func (o *PurchaseNotificationResponse) HasGuildProductPurchase() bool`

HasGuildProductPurchase returns a boolean if a field has been set.

### SetGuildProductPurchaseNil

`func (o *PurchaseNotificationResponse) SetGuildProductPurchaseNil(b bool)`

 SetGuildProductPurchaseNil sets the value for GuildProductPurchase to be an explicit nil

### UnsetGuildProductPurchase
`func (o *PurchaseNotificationResponse) UnsetGuildProductPurchase()`

UnsetGuildProductPurchase ensures that no value is present for GuildProductPurchase, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


