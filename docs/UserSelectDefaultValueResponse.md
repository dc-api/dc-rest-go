# UserSelectDefaultValueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SnowflakeSelectDefaultValueTypes**](SnowflakeSelectDefaultValueTypes.md) |  | 
**Id** | **string** |  | 

## Methods

### NewUserSelectDefaultValueResponse

`func NewUserSelectDefaultValueResponse(type_ SnowflakeSelectDefaultValueTypes, id string, ) *UserSelectDefaultValueResponse`

NewUserSelectDefaultValueResponse instantiates a new UserSelectDefaultValueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSelectDefaultValueResponseWithDefaults

`func NewUserSelectDefaultValueResponseWithDefaults() *UserSelectDefaultValueResponse`

NewUserSelectDefaultValueResponseWithDefaults instantiates a new UserSelectDefaultValueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserSelectDefaultValueResponse) GetType() SnowflakeSelectDefaultValueTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSelectDefaultValueResponse) GetTypeOk() (*SnowflakeSelectDefaultValueTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSelectDefaultValueResponse) SetType(v SnowflakeSelectDefaultValueTypes)`

SetType sets Type field to given value.


### GetId

`func (o *UserSelectDefaultValueResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSelectDefaultValueResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSelectDefaultValueResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


