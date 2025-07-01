# MessageCallResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndedTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Participants** | **[]string** |  | 

## Methods

### NewMessageCallResponse

`func NewMessageCallResponse(participants []string, ) *MessageCallResponse`

NewMessageCallResponse instantiates a new MessageCallResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageCallResponseWithDefaults

`func NewMessageCallResponseWithDefaults() *MessageCallResponse`

NewMessageCallResponseWithDefaults instantiates a new MessageCallResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndedTimestamp

`func (o *MessageCallResponse) GetEndedTimestamp() time.Time`

GetEndedTimestamp returns the EndedTimestamp field if non-nil, zero value otherwise.

### GetEndedTimestampOk

`func (o *MessageCallResponse) GetEndedTimestampOk() (*time.Time, bool)`

GetEndedTimestampOk returns a tuple with the EndedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedTimestamp

`func (o *MessageCallResponse) SetEndedTimestamp(v time.Time)`

SetEndedTimestamp sets EndedTimestamp field to given value.

### HasEndedTimestamp

`func (o *MessageCallResponse) HasEndedTimestamp() bool`

HasEndedTimestamp returns a boolean if a field has been set.

### SetEndedTimestampNil

`func (o *MessageCallResponse) SetEndedTimestampNil(b bool)`

 SetEndedTimestampNil sets the value for EndedTimestamp to be an explicit nil

### UnsetEndedTimestamp
`func (o *MessageCallResponse) UnsetEndedTimestamp()`

UnsetEndedTimestamp ensures that no value is present for EndedTimestamp, not even an explicit nil
### GetParticipants

`func (o *MessageCallResponse) GetParticipants() []string`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *MessageCallResponse) GetParticipantsOk() (*[]string, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *MessageCallResponse) SetParticipants(v []string)`

SetParticipants sets Participants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


