# GithubReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**GithubUser**](GithubUser.md) |  | 
**Body** | Pointer to **NullableString** |  | [optional] 
**HtmlUrl** | **string** |  | 
**State** | **string** |  | 

## Methods

### NewGithubReview

`func NewGithubReview(user GithubUser, htmlUrl string, state string, ) *GithubReview`

NewGithubReview instantiates a new GithubReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubReviewWithDefaults

`func NewGithubReviewWithDefaults() *GithubReview`

NewGithubReviewWithDefaults instantiates a new GithubReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *GithubReview) GetUser() GithubUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GithubReview) GetUserOk() (*GithubUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GithubReview) SetUser(v GithubUser)`

SetUser sets User field to given value.


### GetBody

`func (o *GithubReview) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GithubReview) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GithubReview) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GithubReview) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *GithubReview) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *GithubReview) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetHtmlUrl

`func (o *GithubReview) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubReview) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubReview) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetState

`func (o *GithubReview) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GithubReview) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GithubReview) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


