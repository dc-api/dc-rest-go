# GithubWebhookRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**TagName** | **string** |  | 
**HtmlUrl** | **string** |  | 
**Author** | [**GithubUser**](GithubUser.md) |  | 

## Methods

### NewGithubWebhookRelease

`func NewGithubWebhookRelease(id int32, tagName string, htmlUrl string, author GithubUser, ) *GithubWebhookRelease`

NewGithubWebhookRelease instantiates a new GithubWebhookRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubWebhookReleaseWithDefaults

`func NewGithubWebhookReleaseWithDefaults() *GithubWebhookRelease`

NewGithubWebhookReleaseWithDefaults instantiates a new GithubWebhookRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GithubWebhookRelease) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubWebhookRelease) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubWebhookRelease) SetId(v int32)`

SetId sets Id field to given value.


### GetTagName

`func (o *GithubWebhookRelease) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *GithubWebhookRelease) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *GithubWebhookRelease) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetHtmlUrl

`func (o *GithubWebhookRelease) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubWebhookRelease) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubWebhookRelease) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetAuthor

`func (o *GithubWebhookRelease) GetAuthor() GithubUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GithubWebhookRelease) GetAuthorOk() (*GithubUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GithubWebhookRelease) SetAuthor(v GithubUser)`

SetAuthor sets Author field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


