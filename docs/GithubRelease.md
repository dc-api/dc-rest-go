# GithubRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**TagName** | **string** |  | 
**HtmlUrl** | **string** |  | 
**Author** | [**GithubUser**](GithubUser.md) |  | 

## Methods

### NewGithubRelease

`func NewGithubRelease(id int32, tagName string, htmlUrl string, author GithubUser, ) *GithubRelease`

NewGithubRelease instantiates a new GithubRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubReleaseWithDefaults

`func NewGithubReleaseWithDefaults() *GithubRelease`

NewGithubReleaseWithDefaults instantiates a new GithubRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GithubRelease) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubRelease) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubRelease) SetId(v int32)`

SetId sets Id field to given value.


### GetTagName

`func (o *GithubRelease) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *GithubRelease) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *GithubRelease) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetHtmlUrl

`func (o *GithubRelease) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *GithubRelease) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *GithubRelease) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetAuthor

`func (o *GithubRelease) GetAuthor() GithubUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GithubRelease) GetAuthorOk() (*GithubUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GithubRelease) SetAuthor(v GithubUser)`

SetAuthor sets Author field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


