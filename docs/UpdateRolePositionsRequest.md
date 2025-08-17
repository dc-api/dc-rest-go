# UpdateRolePositionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewUpdateRolePositionsRequest

`func NewUpdateRolePositionsRequest() *UpdateRolePositionsRequest`

NewUpdateRolePositionsRequest instantiates a new UpdateRolePositionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRolePositionsRequestWithDefaults

`func NewUpdateRolePositionsRequestWithDefaults() *UpdateRolePositionsRequest`

NewUpdateRolePositionsRequestWithDefaults instantiates a new UpdateRolePositionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateRolePositionsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateRolePositionsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateRolePositionsRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateRolePositionsRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *UpdateRolePositionsRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UpdateRolePositionsRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UpdateRolePositionsRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *UpdateRolePositionsRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *UpdateRolePositionsRequest) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *UpdateRolePositionsRequest) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


