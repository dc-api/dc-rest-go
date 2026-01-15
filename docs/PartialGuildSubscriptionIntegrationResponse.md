# PartialGuildSubscriptionIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**IntegrationTypes**](IntegrationTypes.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Account** | [**AccountResponse**](AccountResponse.md) |  | 

## Methods

### NewPartialGuildSubscriptionIntegrationResponse

`func NewPartialGuildSubscriptionIntegrationResponse(id string, type_ IntegrationTypes, account AccountResponse, ) *PartialGuildSubscriptionIntegrationResponse`

NewPartialGuildSubscriptionIntegrationResponse instantiates a new PartialGuildSubscriptionIntegrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialGuildSubscriptionIntegrationResponseWithDefaults

`func NewPartialGuildSubscriptionIntegrationResponseWithDefaults() *PartialGuildSubscriptionIntegrationResponse`

NewPartialGuildSubscriptionIntegrationResponseWithDefaults instantiates a new PartialGuildSubscriptionIntegrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartialGuildSubscriptionIntegrationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartialGuildSubscriptionIntegrationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartialGuildSubscriptionIntegrationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PartialGuildSubscriptionIntegrationResponse) GetType() IntegrationTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartialGuildSubscriptionIntegrationResponse) GetTypeOk() (*IntegrationTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartialGuildSubscriptionIntegrationResponse) SetType(v IntegrationTypes)`

SetType sets Type field to given value.


### GetName

`func (o *PartialGuildSubscriptionIntegrationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartialGuildSubscriptionIntegrationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartialGuildSubscriptionIntegrationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartialGuildSubscriptionIntegrationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *PartialGuildSubscriptionIntegrationResponse) GetAccount() AccountResponse`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PartialGuildSubscriptionIntegrationResponse) GetAccountOk() (*AccountResponse, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PartialGuildSubscriptionIntegrationResponse) SetAccount(v AccountResponse)`

SetAccount sets Account field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


