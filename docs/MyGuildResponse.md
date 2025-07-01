# MyGuildResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Banner** | Pointer to **NullableString** |  | [optional] 
**Owner** | **bool** |  | 
**Permissions** | **string** |  | 
**Features** | **[]string** |  | 
**ApproximateMemberCount** | Pointer to **NullableInt32** |  | [optional] 
**ApproximatePresenceCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewMyGuildResponse

`func NewMyGuildResponse(id string, name string, owner bool, permissions string, features []string, ) *MyGuildResponse`

NewMyGuildResponse instantiates a new MyGuildResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyGuildResponseWithDefaults

`func NewMyGuildResponseWithDefaults() *MyGuildResponse`

NewMyGuildResponseWithDefaults instantiates a new MyGuildResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MyGuildResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyGuildResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyGuildResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MyGuildResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MyGuildResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MyGuildResponse) SetName(v string)`

SetName sets Name field to given value.


### GetIcon

`func (o *MyGuildResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *MyGuildResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *MyGuildResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *MyGuildResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *MyGuildResponse) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *MyGuildResponse) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetBanner

`func (o *MyGuildResponse) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *MyGuildResponse) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *MyGuildResponse) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *MyGuildResponse) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### SetBannerNil

`func (o *MyGuildResponse) SetBannerNil(b bool)`

 SetBannerNil sets the value for Banner to be an explicit nil

### UnsetBanner
`func (o *MyGuildResponse) UnsetBanner()`

UnsetBanner ensures that no value is present for Banner, not even an explicit nil
### GetOwner

`func (o *MyGuildResponse) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *MyGuildResponse) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *MyGuildResponse) SetOwner(v bool)`

SetOwner sets Owner field to given value.


### GetPermissions

`func (o *MyGuildResponse) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MyGuildResponse) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *MyGuildResponse) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.


### GetFeatures

`func (o *MyGuildResponse) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *MyGuildResponse) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *MyGuildResponse) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetApproximateMemberCount

`func (o *MyGuildResponse) GetApproximateMemberCount() int32`

GetApproximateMemberCount returns the ApproximateMemberCount field if non-nil, zero value otherwise.

### GetApproximateMemberCountOk

`func (o *MyGuildResponse) GetApproximateMemberCountOk() (*int32, bool)`

GetApproximateMemberCountOk returns a tuple with the ApproximateMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateMemberCount

`func (o *MyGuildResponse) SetApproximateMemberCount(v int32)`

SetApproximateMemberCount sets ApproximateMemberCount field to given value.

### HasApproximateMemberCount

`func (o *MyGuildResponse) HasApproximateMemberCount() bool`

HasApproximateMemberCount returns a boolean if a field has been set.

### SetApproximateMemberCountNil

`func (o *MyGuildResponse) SetApproximateMemberCountNil(b bool)`

 SetApproximateMemberCountNil sets the value for ApproximateMemberCount to be an explicit nil

### UnsetApproximateMemberCount
`func (o *MyGuildResponse) UnsetApproximateMemberCount()`

UnsetApproximateMemberCount ensures that no value is present for ApproximateMemberCount, not even an explicit nil
### GetApproximatePresenceCount

`func (o *MyGuildResponse) GetApproximatePresenceCount() int32`

GetApproximatePresenceCount returns the ApproximatePresenceCount field if non-nil, zero value otherwise.

### GetApproximatePresenceCountOk

`func (o *MyGuildResponse) GetApproximatePresenceCountOk() (*int32, bool)`

GetApproximatePresenceCountOk returns a tuple with the ApproximatePresenceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximatePresenceCount

`func (o *MyGuildResponse) SetApproximatePresenceCount(v int32)`

SetApproximatePresenceCount sets ApproximatePresenceCount field to given value.

### HasApproximatePresenceCount

`func (o *MyGuildResponse) HasApproximatePresenceCount() bool`

HasApproximatePresenceCount returns a boolean if a field has been set.

### SetApproximatePresenceCountNil

`func (o *MyGuildResponse) SetApproximatePresenceCountNil(b bool)`

 SetApproximatePresenceCountNil sets the value for ApproximatePresenceCount to be an explicit nil

### UnsetApproximatePresenceCount
`func (o *MyGuildResponse) UnsetApproximatePresenceCount()`

UnsetApproximatePresenceCount ensures that no value is present for ApproximatePresenceCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


