# GithubCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**Message** | **string** |  | 
**Author** | [**GithubAuthor**](GithubAuthor.md) |  | 

## Methods

### NewGithubCommit

`func NewGithubCommit(id string, url string, message string, author GithubAuthor, ) *GithubCommit`

NewGithubCommit instantiates a new GithubCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubCommitWithDefaults

`func NewGithubCommitWithDefaults() *GithubCommit`

NewGithubCommitWithDefaults instantiates a new GithubCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GithubCommit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubCommit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubCommit) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *GithubCommit) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GithubCommit) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GithubCommit) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMessage

`func (o *GithubCommit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GithubCommit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GithubCommit) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAuthor

`func (o *GithubCommit) GetAuthor() GithubAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GithubCommit) GetAuthorOk() (*GithubAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GithubCommit) SetAuthor(v GithubAuthor)`

SetAuthor sets Author field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


