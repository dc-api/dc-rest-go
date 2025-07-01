# PollAnswerDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]UserResponse**](UserResponse.md) |  | [optional] 

## Methods

### NewPollAnswerDetailsResponse

`func NewPollAnswerDetailsResponse() *PollAnswerDetailsResponse`

NewPollAnswerDetailsResponse instantiates a new PollAnswerDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollAnswerDetailsResponseWithDefaults

`func NewPollAnswerDetailsResponseWithDefaults() *PollAnswerDetailsResponse`

NewPollAnswerDetailsResponseWithDefaults instantiates a new PollAnswerDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *PollAnswerDetailsResponse) GetUsers() []UserResponse`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PollAnswerDetailsResponse) GetUsersOk() (*[]UserResponse, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PollAnswerDetailsResponse) SetUsers(v []UserResponse)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PollAnswerDetailsResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *PollAnswerDetailsResponse) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *PollAnswerDetailsResponse) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


