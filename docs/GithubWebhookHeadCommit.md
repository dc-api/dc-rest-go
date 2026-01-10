# GithubWebhookHeadCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**Message** | **string** |  | 
**Author** | [**GithubAuthor**](GithubAuthor.md) |  | 

## Methods

### NewGithubWebhookHeadCommit

`func NewGithubWebhookHeadCommit(id string, url string, message string, author GithubAuthor, ) *GithubWebhookHeadCommit`

NewGithubWebhookHeadCommit instantiates a new GithubWebhookHeadCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubWebhookHeadCommitWithDefaults

`func NewGithubWebhookHeadCommitWithDefaults() *GithubWebhookHeadCommit`

NewGithubWebhookHeadCommitWithDefaults instantiates a new GithubWebhookHeadCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GithubWebhookHeadCommit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubWebhookHeadCommit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubWebhookHeadCommit) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *GithubWebhookHeadCommit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GithubWebhookHeadCommit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GithubWebhookHeadCommit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMessage

`func (o *GithubWebhookHeadCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GithubWebhookHeadCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GithubWebhookHeadCommit) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAuthor

`func (o *GithubWebhookHeadCommit) GetAuthor() GithubAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GithubWebhookHeadCommit) GetAuthorOk() (*GithubAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GithubWebhookHeadCommit) SetAuthor(v GithubAuthor)`

SetAuthor sets Author field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


