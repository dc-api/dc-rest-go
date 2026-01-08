# EntitlementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SkuId** | **string** |  | 
**ApplicationId** | **string** |  | 
**UserId** | **string** |  | 
**GuildId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**Deleted** | **bool** |  | 
**StartsAt** | Pointer to **time.Time** |  | [optional] 
**EndsAt** | Pointer to **time.Time** |  | [optional] 
**Type** | [**EntitlementTypes**](EntitlementTypes.md) |  | 
**FulfilledAt** | Pointer to **time.Time** |  | [optional] 
**FulfillmentStatus** | Pointer to [**EntitlementResponseFulfillmentStatus**](EntitlementResponseFulfillmentStatus.md) |  | [optional] 
**Consumed** | Pointer to **bool** |  | [optional] 
**GifterUserId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 
**ParentId** | Pointer to [**GetEntitlementsSkuIdsParameterOneOfInner**](GetEntitlementsSkuIdsParameterOneOfInner.md) |  | [optional] 

## Methods

### NewEntitlementResponse

`func NewEntitlementResponse(id string, skuId string, applicationId string, userId string, deleted bool, type_ EntitlementTypes, ) *EntitlementResponse`

NewEntitlementResponse instantiates a new EntitlementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementResponseWithDefaults

`func NewEntitlementResponseWithDefaults() *EntitlementResponse`

NewEntitlementResponseWithDefaults instantiates a new EntitlementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSkuId

`func (o *EntitlementResponse) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *EntitlementResponse) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *EntitlementResponse) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.


### GetApplicationId

`func (o *EntitlementResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *EntitlementResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *EntitlementResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetUserId

`func (o *EntitlementResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EntitlementResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EntitlementResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetGuildId

`func (o *EntitlementResponse) GetGuildId() GetEntitlementsSkuIdsParameterOneOfInner`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *EntitlementResponse) GetGuildIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *EntitlementResponse) SetGuildId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *EntitlementResponse) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetDeleted

`func (o *EntitlementResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *EntitlementResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *EntitlementResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetStartsAt

`func (o *EntitlementResponse) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *EntitlementResponse) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *EntitlementResponse) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *EntitlementResponse) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetEndsAt

`func (o *EntitlementResponse) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *EntitlementResponse) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *EntitlementResponse) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *EntitlementResponse) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetType

`func (o *EntitlementResponse) GetType() EntitlementTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitlementResponse) GetTypeOk() (*EntitlementTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitlementResponse) SetType(v EntitlementTypes)`

SetType sets Type field to given value.


### GetFulfilledAt

`func (o *EntitlementResponse) GetFulfilledAt() time.Time`

GetFulfilledAt returns the FulfilledAt field if non-nil, zero value otherwise.

### GetFulfilledAtOk

`func (o *EntitlementResponse) GetFulfilledAtOk() (*time.Time, bool)`

GetFulfilledAtOk returns a tuple with the FulfilledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledAt

`func (o *EntitlementResponse) SetFulfilledAt(v time.Time)`

SetFulfilledAt sets FulfilledAt field to given value.

### HasFulfilledAt

`func (o *EntitlementResponse) HasFulfilledAt() bool`

HasFulfilledAt returns a boolean if a field has been set.

### GetFulfillmentStatus

`func (o *EntitlementResponse) GetFulfillmentStatus() EntitlementResponseFulfillmentStatus`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *EntitlementResponse) GetFulfillmentStatusOk() (*EntitlementResponseFulfillmentStatus, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *EntitlementResponse) SetFulfillmentStatus(v EntitlementResponseFulfillmentStatus)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.

### HasFulfillmentStatus

`func (o *EntitlementResponse) HasFulfillmentStatus() bool`

HasFulfillmentStatus returns a boolean if a field has been set.

### GetConsumed

`func (o *EntitlementResponse) GetConsumed() bool`

GetConsumed returns the Consumed field if non-nil, zero value otherwise.

### GetConsumedOk

`func (o *EntitlementResponse) GetConsumedOk() (*bool, bool)`

GetConsumedOk returns a tuple with the Consumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumed

`func (o *EntitlementResponse) SetConsumed(v bool)`

SetConsumed sets Consumed field to given value.

### HasConsumed

`func (o *EntitlementResponse) HasConsumed() bool`

HasConsumed returns a boolean if a field has been set.

### GetGifterUserId

`func (o *EntitlementResponse) GetGifterUserId() GetEntitlementsSkuIdsParameterOneOfInner`

GetGifterUserId returns the GifterUserId field if non-nil, zero value otherwise.

### GetGifterUserIdOk

`func (o *EntitlementResponse) GetGifterUserIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetGifterUserIdOk returns a tuple with the GifterUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifterUserId

`func (o *EntitlementResponse) SetGifterUserId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetGifterUserId sets GifterUserId field to given value.

### HasGifterUserId

`func (o *EntitlementResponse) HasGifterUserId() bool`

HasGifterUserId returns a boolean if a field has been set.

### GetParentId

`func (o *EntitlementResponse) GetParentId() GetEntitlementsSkuIdsParameterOneOfInner`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *EntitlementResponse) GetParentIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *EntitlementResponse) SetParentId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *EntitlementResponse) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


