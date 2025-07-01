# GatewayBotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**SessionStartLimit** | [**GatewayBotSessionStartLimitResponse**](GatewayBotSessionStartLimitResponse.md) |  | 
**Shards** | **int32** |  | 

## Methods

### NewGatewayBotResponse

`func NewGatewayBotResponse(url string, sessionStartLimit GatewayBotSessionStartLimitResponse, shards int32, ) *GatewayBotResponse`

NewGatewayBotResponse instantiates a new GatewayBotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayBotResponseWithDefaults

`func NewGatewayBotResponseWithDefaults() *GatewayBotResponse`

NewGatewayBotResponseWithDefaults instantiates a new GatewayBotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *GatewayBotResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GatewayBotResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GatewayBotResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSessionStartLimit

`func (o *GatewayBotResponse) GetSessionStartLimit() GatewayBotSessionStartLimitResponse`

GetSessionStartLimit returns the SessionStartLimit field if non-nil, zero value otherwise.

### GetSessionStartLimitOk

`func (o *GatewayBotResponse) GetSessionStartLimitOk() (*GatewayBotSessionStartLimitResponse, bool)`

GetSessionStartLimitOk returns a tuple with the SessionStartLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStartLimit

`func (o *GatewayBotResponse) SetSessionStartLimit(v GatewayBotSessionStartLimitResponse)`

SetSessionStartLimit sets SessionStartLimit field to given value.


### GetShards

`func (o *GatewayBotResponse) GetShards() int32`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *GatewayBotResponse) GetShardsOk() (*int32, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *GatewayBotResponse) SetShards(v int32)`

SetShards sets Shards field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


