# PartialDiscordIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**IntegrationTypes**](IntegrationTypes.md) |  | 
**Name** | Pointer to **string** |  | [optional] 
**Account** | [**AccountResponse**](AccountResponse.md) |  | 
**ApplicationId** | **string** |  | 

## Methods

### NewPartialDiscordIntegrationResponse

`func NewPartialDiscordIntegrationResponse(id string, type_ IntegrationTypes, account AccountResponse, applicationId string, ) *PartialDiscordIntegrationResponse`

NewPartialDiscordIntegrationResponse instantiates a new PartialDiscordIntegrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialDiscordIntegrationResponseWithDefaults

`func NewPartialDiscordIntegrationResponseWithDefaults() *PartialDiscordIntegrationResponse`

NewPartialDiscordIntegrationResponseWithDefaults instantiates a new PartialDiscordIntegrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartialDiscordIntegrationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartialDiscordIntegrationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartialDiscordIntegrationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PartialDiscordIntegrationResponse) GetType() IntegrationTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PartialDiscordIntegrationResponse) GetTypeOk() (*IntegrationTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PartialDiscordIntegrationResponse) SetType(v IntegrationTypes)`

SetType sets Type field to given value.


### GetName

`func (o *PartialDiscordIntegrationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartialDiscordIntegrationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartialDiscordIntegrationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PartialDiscordIntegrationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *PartialDiscordIntegrationResponse) GetAccount() AccountResponse`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PartialDiscordIntegrationResponse) GetAccountOk() (*AccountResponse, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PartialDiscordIntegrationResponse) SetAccount(v AccountResponse)`

SetAccount sets Account field to given value.


### GetApplicationId

`func (o *PartialDiscordIntegrationResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PartialDiscordIntegrationResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PartialDiscordIntegrationResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


