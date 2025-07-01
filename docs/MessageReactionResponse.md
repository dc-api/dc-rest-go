# MessageReactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emoji** | [**MessageReactionEmojiResponse**](MessageReactionEmojiResponse.md) |  | 
**Count** | **int32** |  | 
**CountDetails** | [**MessageReactionCountDetailsResponse**](MessageReactionCountDetailsResponse.md) |  | 
**BurstColors** | **[]string** |  | 
**MeBurst** | **bool** |  | 
**Me** | **bool** |  | 

## Methods

### NewMessageReactionResponse

`func NewMessageReactionResponse(emoji MessageReactionEmojiResponse, count int32, countDetails MessageReactionCountDetailsResponse, burstColors []string, meBurst bool, me bool, ) *MessageReactionResponse`

NewMessageReactionResponse instantiates a new MessageReactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageReactionResponseWithDefaults

`func NewMessageReactionResponseWithDefaults() *MessageReactionResponse`

NewMessageReactionResponseWithDefaults instantiates a new MessageReactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmoji

`func (o *MessageReactionResponse) GetEmoji() MessageReactionEmojiResponse`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *MessageReactionResponse) GetEmojiOk() (*MessageReactionEmojiResponse, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *MessageReactionResponse) SetEmoji(v MessageReactionEmojiResponse)`

SetEmoji sets Emoji field to given value.


### GetCount

`func (o *MessageReactionResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MessageReactionResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MessageReactionResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetCountDetails

`func (o *MessageReactionResponse) GetCountDetails() MessageReactionCountDetailsResponse`

GetCountDetails returns the CountDetails field if non-nil, zero value otherwise.

### GetCountDetailsOk

`func (o *MessageReactionResponse) GetCountDetailsOk() (*MessageReactionCountDetailsResponse, bool)`

GetCountDetailsOk returns a tuple with the CountDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountDetails

`func (o *MessageReactionResponse) SetCountDetails(v MessageReactionCountDetailsResponse)`

SetCountDetails sets CountDetails field to given value.


### GetBurstColors

`func (o *MessageReactionResponse) GetBurstColors() []string`

GetBurstColors returns the BurstColors field if non-nil, zero value otherwise.

### GetBurstColorsOk

`func (o *MessageReactionResponse) GetBurstColorsOk() (*[]string, bool)`

GetBurstColorsOk returns a tuple with the BurstColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstColors

`func (o *MessageReactionResponse) SetBurstColors(v []string)`

SetBurstColors sets BurstColors field to given value.


### GetMeBurst

`func (o *MessageReactionResponse) GetMeBurst() bool`

GetMeBurst returns the MeBurst field if non-nil, zero value otherwise.

### GetMeBurstOk

`func (o *MessageReactionResponse) GetMeBurstOk() (*bool, bool)`

GetMeBurstOk returns a tuple with the MeBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeBurst

`func (o *MessageReactionResponse) SetMeBurst(v bool)`

SetMeBurst sets MeBurst field to given value.


### GetMe

`func (o *MessageReactionResponse) GetMe() bool`

GetMe returns the Me field if non-nil, zero value otherwise.

### GetMeOk

`func (o *MessageReactionResponse) GetMeOk() (*bool, bool)`

GetMeOk returns a tuple with the Me field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMe

`func (o *MessageReactionResponse) SetMe(v bool)`

SetMe sets Me field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


