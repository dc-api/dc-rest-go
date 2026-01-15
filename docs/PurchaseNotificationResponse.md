# PurchaseNotificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PurchaseType**](PurchaseType.md) |  | 
**GuildProductPurchase** | Pointer to [**GuildProductPurchaseResponse**](GuildProductPurchaseResponse.md) |  | [optional] 

## Methods

### NewPurchaseNotificationResponse

`func NewPurchaseNotificationResponse(type_ PurchaseType, ) *PurchaseNotificationResponse`

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

`func (o *PurchaseNotificationResponse) GetType() PurchaseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PurchaseNotificationResponse) GetTypeOk() (*PurchaseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PurchaseNotificationResponse) SetType(v PurchaseType)`

SetType sets Type field to given value.


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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


