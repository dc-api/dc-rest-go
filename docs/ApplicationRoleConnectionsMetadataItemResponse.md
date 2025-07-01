# ApplicationRoleConnectionsMetadataItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** |  | 
**Key** | **string** |  | 
**Name** | **string** |  | 
**NameLocalizations** | Pointer to **map[string]string** |  | [optional] 
**Description** | **string** |  | 
**DescriptionLocalizations** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewApplicationRoleConnectionsMetadataItemResponse

`func NewApplicationRoleConnectionsMetadataItemResponse(type_ int32, key string, name string, description string, ) *ApplicationRoleConnectionsMetadataItemResponse`

NewApplicationRoleConnectionsMetadataItemResponse instantiates a new ApplicationRoleConnectionsMetadataItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRoleConnectionsMetadataItemResponseWithDefaults

`func NewApplicationRoleConnectionsMetadataItemResponseWithDefaults() *ApplicationRoleConnectionsMetadataItemResponse`

NewApplicationRoleConnectionsMetadataItemResponseWithDefaults instantiates a new ApplicationRoleConnectionsMetadataItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationRoleConnectionsMetadataItemResponse) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationRoleConnectionsMetadataItemResponse) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationRoleConnectionsMetadataItemResponse) SetType(v int32)`

SetType sets Type field to given value.


### GetKey

`func (o *ApplicationRoleConnectionsMetadataItemResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApplicationRoleConnectionsMetadataItemResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApplicationRoleConnectionsMetadataItemResponse) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ApplicationRoleConnectionsMetadataItemResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRoleConnectionsMetadataItemResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRoleConnectionsMetadataItemResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNameLocalizations

`func (o *ApplicationRoleConnectionsMetadataItemResponse) GetNameLocalizations() map[string]string`

GetNameLocalizations returns the NameLocalizations field if non-nil, zero value otherwise.

### GetNameLocalizationsOk

`func (o *ApplicationRoleConnectionsMetadataItemResponse) GetNameLocalizationsOk() (*map[string]string, bool)`

GetNameLocalizationsOk returns a tuple with the NameLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocalizations

`func (o *ApplicationRoleConnectionsMetadataItemResponse) SetNameLocalizations(v map[string]string)`

SetNameLocalizations sets NameLocalizations field to given value.

### HasNameLocalizations

`func (o *ApplicationRoleConnectionsMetadataItemResponse) HasNameLocalizations() bool`

HasNameLocalizations returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationRoleConnectionsMetadataItemResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationRoleConnectionsMetadataItemResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationRoleConnectionsMetadataItemResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDescriptionLocalizations

`func (o *ApplicationRoleConnectionsMetadataItemResponse) GetDescriptionLocalizations() map[string]string`

GetDescriptionLocalizations returns the DescriptionLocalizations field if non-nil, zero value otherwise.

### GetDescriptionLocalizationsOk

`func (o *ApplicationRoleConnectionsMetadataItemResponse) GetDescriptionLocalizationsOk() (*map[string]string, bool)`

GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocalizations

`func (o *ApplicationRoleConnectionsMetadataItemResponse) SetDescriptionLocalizations(v map[string]string)`

SetDescriptionLocalizations sets DescriptionLocalizations field to given value.

### HasDescriptionLocalizations

`func (o *ApplicationRoleConnectionsMetadataItemResponse) HasDescriptionLocalizations() bool`

HasDescriptionLocalizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


