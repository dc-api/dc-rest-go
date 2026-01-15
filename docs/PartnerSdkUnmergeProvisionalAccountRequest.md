# PartnerSdkUnmergeProvisionalAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**ClientSecret** | Pointer to **string** |  | [optional] 
**ExternalAuthToken** | **string** |  | 
**ExternalAuthType** | [**ApplicationIdentityProviderAuthType**](ApplicationIdentityProviderAuthType.md) |  | 

## Methods

### NewPartnerSdkUnmergeProvisionalAccountRequest

`func NewPartnerSdkUnmergeProvisionalAccountRequest(clientId string, externalAuthToken string, externalAuthType ApplicationIdentityProviderAuthType, ) *PartnerSdkUnmergeProvisionalAccountRequest`

NewPartnerSdkUnmergeProvisionalAccountRequest instantiates a new PartnerSdkUnmergeProvisionalAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerSdkUnmergeProvisionalAccountRequestWithDefaults

`func NewPartnerSdkUnmergeProvisionalAccountRequestWithDefaults() *PartnerSdkUnmergeProvisionalAccountRequest`

NewPartnerSdkUnmergeProvisionalAccountRequestWithDefaults instantiates a new PartnerSdkUnmergeProvisionalAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PartnerSdkUnmergeProvisionalAccountRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PartnerSdkUnmergeProvisionalAccountRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PartnerSdkUnmergeProvisionalAccountRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetExternalAuthToken

`func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetExternalAuthToken() string`

GetExternalAuthToken returns the ExternalAuthToken field if non-nil, zero value otherwise.

### GetExternalAuthTokenOk

`func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetExternalAuthTokenOk() (*string, bool)`

GetExternalAuthTokenOk returns a tuple with the ExternalAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAuthToken

`func (o *PartnerSdkUnmergeProvisionalAccountRequest) SetExternalAuthToken(v string)`

SetExternalAuthToken sets ExternalAuthToken field to given value.


### GetExternalAuthType

`func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetExternalAuthType() ApplicationIdentityProviderAuthType`

GetExternalAuthType returns the ExternalAuthType field if non-nil, zero value otherwise.

### GetExternalAuthTypeOk

`func (o *PartnerSdkUnmergeProvisionalAccountRequest) GetExternalAuthTypeOk() (*ApplicationIdentityProviderAuthType, bool)`

GetExternalAuthTypeOk returns a tuple with the ExternalAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAuthType

`func (o *PartnerSdkUnmergeProvisionalAccountRequest) SetExternalAuthType(v ApplicationIdentityProviderAuthType)`

SetExternalAuthType sets ExternalAuthType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


