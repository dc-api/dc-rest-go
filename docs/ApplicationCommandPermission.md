# ApplicationCommandPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**ApplicationCommandPermissionType**](ApplicationCommandPermissionType.md) |  | 
**Permission** | **bool** |  | 

## Methods

### NewApplicationCommandPermission

`func NewApplicationCommandPermission(id string, type_ ApplicationCommandPermissionType, permission bool, ) *ApplicationCommandPermission`

NewApplicationCommandPermission instantiates a new ApplicationCommandPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCommandPermissionWithDefaults

`func NewApplicationCommandPermissionWithDefaults() *ApplicationCommandPermission`

NewApplicationCommandPermissionWithDefaults instantiates a new ApplicationCommandPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationCommandPermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationCommandPermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationCommandPermission) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ApplicationCommandPermission) GetType() ApplicationCommandPermissionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationCommandPermission) GetTypeOk() (*ApplicationCommandPermissionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationCommandPermission) SetType(v ApplicationCommandPermissionType)`

SetType sets Type field to given value.


### GetPermission

`func (o *ApplicationCommandPermission) GetPermission() bool`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ApplicationCommandPermission) GetPermissionOk() (*bool, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ApplicationCommandPermission) SetPermission(v bool)`

SetPermission sets Permission field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


