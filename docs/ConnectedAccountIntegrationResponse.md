# ConnectedAccountIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**IntegrationTypes**](IntegrationTypes.md) |  | 
**Account** | [**AccountResponse**](AccountResponse.md) |  | 
**Guild** | [**ConnectedAccountGuildResponse**](ConnectedAccountGuildResponse.md) |  | 

## Methods

### NewConnectedAccountIntegrationResponse

`func NewConnectedAccountIntegrationResponse(id string, type_ IntegrationTypes, account AccountResponse, guild ConnectedAccountGuildResponse, ) *ConnectedAccountIntegrationResponse`

NewConnectedAccountIntegrationResponse instantiates a new ConnectedAccountIntegrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedAccountIntegrationResponseWithDefaults

`func NewConnectedAccountIntegrationResponseWithDefaults() *ConnectedAccountIntegrationResponse`

NewConnectedAccountIntegrationResponseWithDefaults instantiates a new ConnectedAccountIntegrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectedAccountIntegrationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectedAccountIntegrationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectedAccountIntegrationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ConnectedAccountIntegrationResponse) GetType() IntegrationTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectedAccountIntegrationResponse) GetTypeOk() (*IntegrationTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectedAccountIntegrationResponse) SetType(v IntegrationTypes)`

SetType sets Type field to given value.


### GetAccount

`func (o *ConnectedAccountIntegrationResponse) GetAccount() AccountResponse`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ConnectedAccountIntegrationResponse) GetAccountOk() (*AccountResponse, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ConnectedAccountIntegrationResponse) SetAccount(v AccountResponse)`

SetAccount sets Account field to given value.


### GetGuild

`func (o *ConnectedAccountIntegrationResponse) GetGuild() ConnectedAccountGuildResponse`

GetGuild returns the Guild field if non-nil, zero value otherwise.

### GetGuildOk

`func (o *ConnectedAccountIntegrationResponse) GetGuildOk() (*ConnectedAccountGuildResponse, bool)`

GetGuildOk returns a tuple with the Guild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuild

`func (o *ConnectedAccountIntegrationResponse) SetGuild(v ConnectedAccountGuildResponse)`

SetGuild sets Guild field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


