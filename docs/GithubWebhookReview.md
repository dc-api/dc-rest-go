# GithubWebhookReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**GithubUser**](GithubUser.md) |  | 
**Body** | Pointer to **string** |  | [optional] 
**HtmlUrl** | **string** |  | 
**State** | **string** |  | 

## Methods

### NewGithubWebhookReview

`func NewGithubWebhookReview(user GithubUser, htmlUrl string, state string, ) *GithubWebhookReview`

NewGithubWebhookReview instantiates a new GithubWebhookReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubWebhookReviewWithDefaults

`func NewGithubWebhookReviewWithDefaults() *GithubWebhookReview`

NewGithubWebhookReviewWithDefaults instantiates a new GithubWebhookReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *GithubWebhookReview) GetUser() GithubUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GithubWebhookReview) GetUserOk() (*GithubUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GithubWebhookReview) SetUser(v GithubUser)`

SetUser sets User field to given value.


### GetBody

`func (o *GithubWebhookReview) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GithubWebhookReview) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GithubWebhookReview) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GithubWebhookReview) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *GithubWebhookReview) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubWebhookReview) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubWebhookReview) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetState

`func (o *GithubWebhookReview) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GithubWebhookReview) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GithubWebhookReview) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


