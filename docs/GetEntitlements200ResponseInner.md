# GetEntitlements200ResponseInner

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

### NewGetEntitlements200ResponseInner

`func NewGetEntitlements200ResponseInner(id string, skuId string, applicationId string, userId string, deleted bool, type_ EntitlementTypes, ) *GetEntitlements200ResponseInner`

NewGetEntitlements200ResponseInner instantiates a new GetEntitlements200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEntitlements200ResponseInnerWithDefaults

`func NewGetEntitlements200ResponseInnerWithDefaults() *GetEntitlements200ResponseInner`

NewGetEntitlements200ResponseInnerWithDefaults instantiates a new GetEntitlements200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEntitlements200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEntitlements200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEntitlements200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetSkuId

`func (o *GetEntitlements200ResponseInner) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *GetEntitlements200ResponseInner) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *GetEntitlements200ResponseInner) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.


### GetApplicationId

`func (o *GetEntitlements200ResponseInner) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GetEntitlements200ResponseInner) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GetEntitlements200ResponseInner) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetUserId

`func (o *GetEntitlements200ResponseInner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetEntitlements200ResponseInner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetEntitlements200ResponseInner) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetGuildId

`func (o *GetEntitlements200ResponseInner) GetGuildId() GetEntitlementsSkuIdsParameterOneOfInner`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *GetEntitlements200ResponseInner) GetGuildIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *GetEntitlements200ResponseInner) SetGuildId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *GetEntitlements200ResponseInner) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetDeleted

`func (o *GetEntitlements200ResponseInner) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *GetEntitlements200ResponseInner) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *GetEntitlements200ResponseInner) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.


### GetStartsAt

`func (o *GetEntitlements200ResponseInner) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GetEntitlements200ResponseInner) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GetEntitlements200ResponseInner) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *GetEntitlements200ResponseInner) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetEndsAt

`func (o *GetEntitlements200ResponseInner) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *GetEntitlements200ResponseInner) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *GetEntitlements200ResponseInner) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *GetEntitlements200ResponseInner) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetType

`func (o *GetEntitlements200ResponseInner) GetType() EntitlementTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetEntitlements200ResponseInner) GetTypeOk() (*EntitlementTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetEntitlements200ResponseInner) SetType(v EntitlementTypes)`

SetType sets Type field to given value.


### GetFulfilledAt

`func (o *GetEntitlements200ResponseInner) GetFulfilledAt() time.Time`

GetFulfilledAt returns the FulfilledAt field if non-nil, zero value otherwise.

### GetFulfilledAtOk

`func (o *GetEntitlements200ResponseInner) GetFulfilledAtOk() (*time.Time, bool)`

GetFulfilledAtOk returns a tuple with the FulfilledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledAt

`func (o *GetEntitlements200ResponseInner) SetFulfilledAt(v time.Time)`

SetFulfilledAt sets FulfilledAt field to given value.

### HasFulfilledAt

`func (o *GetEntitlements200ResponseInner) HasFulfilledAt() bool`

HasFulfilledAt returns a boolean if a field has been set.

### GetFulfillmentStatus

`func (o *GetEntitlements200ResponseInner) GetFulfillmentStatus() EntitlementResponseFulfillmentStatus`

GetFulfillmentStatus returns the FulfillmentStatus field if non-nil, zero value otherwise.

### GetFulfillmentStatusOk

`func (o *GetEntitlements200ResponseInner) GetFulfillmentStatusOk() (*EntitlementResponseFulfillmentStatus, bool)`

GetFulfillmentStatusOk returns a tuple with the FulfillmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentStatus

`func (o *GetEntitlements200ResponseInner) SetFulfillmentStatus(v EntitlementResponseFulfillmentStatus)`

SetFulfillmentStatus sets FulfillmentStatus field to given value.

### HasFulfillmentStatus

`func (o *GetEntitlements200ResponseInner) HasFulfillmentStatus() bool`

HasFulfillmentStatus returns a boolean if a field has been set.

### GetConsumed

`func (o *GetEntitlements200ResponseInner) GetConsumed() bool`

GetConsumed returns the Consumed field if non-nil, zero value otherwise.

### GetConsumedOk

`func (o *GetEntitlements200ResponseInner) GetConsumedOk() (*bool, bool)`

GetConsumedOk returns a tuple with the Consumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumed

`func (o *GetEntitlements200ResponseInner) SetConsumed(v bool)`

SetConsumed sets Consumed field to given value.

### HasConsumed

`func (o *GetEntitlements200ResponseInner) HasConsumed() bool`

HasConsumed returns a boolean if a field has been set.

### GetGifterUserId

`func (o *GetEntitlements200ResponseInner) GetGifterUserId() GetEntitlementsSkuIdsParameterOneOfInner`

GetGifterUserId returns the GifterUserId field if non-nil, zero value otherwise.

### GetGifterUserIdOk

`func (o *GetEntitlements200ResponseInner) GetGifterUserIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetGifterUserIdOk returns a tuple with the GifterUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifterUserId

`func (o *GetEntitlements200ResponseInner) SetGifterUserId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetGifterUserId sets GifterUserId field to given value.

### HasGifterUserId

`func (o *GetEntitlements200ResponseInner) HasGifterUserId() bool`

HasGifterUserId returns a boolean if a field has been set.

### GetParentId

`func (o *GetEntitlements200ResponseInner) GetParentId() GetEntitlementsSkuIdsParameterOneOfInner`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GetEntitlements200ResponseInner) GetParentIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GetEntitlements200ResponseInner) SetParentId(v GetEntitlementsSkuIdsParameterOneOfInner)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GetEntitlements200ResponseInner) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


