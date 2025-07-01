# CommandPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ApplicationId** | **string** |  | 
**GuildId** | **string** |  | 
**Permissions** | [**[]CommandPermissionResponse**](CommandPermissionResponse.md) |  | 

## Methods

### NewCommandPermissionsResponse

`func NewCommandPermissionsResponse(id string, applicationId string, guildId string, permissions []CommandPermissionResponse, ) *CommandPermissionsResponse`

NewCommandPermissionsResponse instantiates a new CommandPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandPermissionsResponseWithDefaults

`func NewCommandPermissionsResponseWithDefaults() *CommandPermissionsResponse`

NewCommandPermissionsResponseWithDefaults instantiates a new CommandPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommandPermissionsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommandPermissionsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommandPermissionsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetApplicationId

`func (o *CommandPermissionsResponse) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CommandPermissionsResponse) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CommandPermissionsResponse) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetGuildId

`func (o *CommandPermissionsResponse) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *CommandPermissionsResponse) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *CommandPermissionsResponse) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.


### GetPermissions

`func (o *CommandPermissionsResponse) GetPermissions() []CommandPermissionResponse`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CommandPermissionsResponse) GetPermissionsOk() (*[]CommandPermissionResponse, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CommandPermissionsResponse) SetPermissions(v []CommandPermissionResponse)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


