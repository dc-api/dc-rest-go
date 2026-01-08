# GithubWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to [**GithubWebhookComment**](GithubWebhookComment.md) |  | [optional] 
**Issue** | Pointer to [**GithubWebhookIssue**](GithubWebhookIssue.md) |  | [optional] 
**PullRequest** | Pointer to [**GithubWebhookIssue**](GithubWebhookIssue.md) |  | [optional] 
**Repository** | Pointer to [**GithubWebhookRepository**](GithubWebhookRepository.md) |  | [optional] 
**Forkee** | Pointer to [**GithubWebhookRepository**](GithubWebhookRepository.md) |  | [optional] 
**Sender** | [**GithubUser**](GithubUser.md) |  | 
**Member** | Pointer to [**GithubWebhookMember**](GithubWebhookMember.md) |  | [optional] 
**Release** | Pointer to [**GithubWebhookRelease**](GithubWebhookRelease.md) |  | [optional] 
**HeadCommit** | Pointer to [**GithubWebhookHeadCommit**](GithubWebhookHeadCommit.md) |  | [optional] 
**Commits** | Pointer to [**[]GithubCommit**](GithubCommit.md) |  | [optional] 
**Forced** | Pointer to **bool** |  | [optional] 
**Compare** | Pointer to **string** |  | [optional] 
**Review** | Pointer to [**GithubWebhookReview**](GithubWebhookReview.md) |  | [optional] 
**CheckRun** | Pointer to [**GithubWebhookCheckRun**](GithubWebhookCheckRun.md) |  | [optional] 
**CheckSuite** | Pointer to [**GithubWebhookCheckSuite**](GithubWebhookCheckSuite.md) |  | [optional] 
**Discussion** | Pointer to [**GithubWebhookDiscussion**](GithubWebhookDiscussion.md) |  | [optional] 
**Answer** | Pointer to [**GithubWebhookComment**](GithubWebhookComment.md) |  | [optional] 

## Methods

### NewGithubWebhook

`func NewGithubWebhook(sender GithubUser, ) *GithubWebhook`

NewGithubWebhook instantiates a new GithubWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubWebhookWithDefaults

`func NewGithubWebhookWithDefaults() *GithubWebhook`

NewGithubWebhookWithDefaults instantiates a new GithubWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GithubWebhook) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GithubWebhook) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GithubWebhook) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GithubWebhook) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRef

`func (o *GithubWebhook) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GithubWebhook) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GithubWebhook) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GithubWebhook) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRefType

`func (o *GithubWebhook) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GithubWebhook) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GithubWebhook) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GithubWebhook) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetComment

`func (o *GithubWebhook) GetComment() GithubWebhookComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GithubWebhook) GetCommentOk() (*GithubWebhookComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GithubWebhook) SetComment(v GithubWebhookComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GithubWebhook) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetIssue

`func (o *GithubWebhook) GetIssue() GithubWebhookIssue`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *GithubWebhook) GetIssueOk() (*GithubWebhookIssue, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *GithubWebhook) SetIssue(v GithubWebhookIssue)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *GithubWebhook) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetPullRequest

`func (o *GithubWebhook) GetPullRequest() GithubWebhookIssue`

GetPullRequest returns the PullRequest field if non-nil, zero value otherwise.

### GetPullRequestOk

`func (o *GithubWebhook) GetPullRequestOk() (*GithubWebhookIssue, bool)`

GetPullRequestOk returns a tuple with the PullRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequest

`func (o *GithubWebhook) SetPullRequest(v GithubWebhookIssue)`

SetPullRequest sets PullRequest field to given value.

### HasPullRequest

`func (o *GithubWebhook) HasPullRequest() bool`

HasPullRequest returns a boolean if a field has been set.

### GetRepository

`func (o *GithubWebhook) GetRepository() GithubWebhookRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *GithubWebhook) GetRepositoryOk() (*GithubWebhookRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *GithubWebhook) SetRepository(v GithubWebhookRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *GithubWebhook) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetForkee

`func (o *GithubWebhook) GetForkee() GithubWebhookRepository`

GetForkee returns the Forkee field if non-nil, zero value otherwise.

### GetForkeeOk

`func (o *GithubWebhook) GetForkeeOk() (*GithubWebhookRepository, bool)`

GetForkeeOk returns a tuple with the Forkee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForkee

`func (o *GithubWebhook) SetForkee(v GithubWebhookRepository)`

SetForkee sets Forkee field to given value.

### HasForkee

`func (o *GithubWebhook) HasForkee() bool`

HasForkee returns a boolean if a field has been set.

### GetSender

`func (o *GithubWebhook) GetSender() GithubUser`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *GithubWebhook) GetSenderOk() (*GithubUser, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *GithubWebhook) SetSender(v GithubUser)`

SetSender sets Sender field to given value.


### GetMember

`func (o *GithubWebhook) GetMember() GithubWebhookMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GithubWebhook) GetMemberOk() (*GithubWebhookMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GithubWebhook) SetMember(v GithubWebhookMember)`

SetMember sets Member field to given value.

### HasMember

`func (o *GithubWebhook) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetRelease

`func (o *GithubWebhook) GetRelease() GithubWebhookRelease`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *GithubWebhook) GetReleaseOk() (*GithubWebhookRelease, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *GithubWebhook) SetRelease(v GithubWebhookRelease)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *GithubWebhook) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetHeadCommit

`func (o *GithubWebhook) GetHeadCommit() GithubWebhookHeadCommit`

GetHeadCommit returns the HeadCommit field if non-nil, zero value otherwise.

### GetHeadCommitOk

`func (o *GithubWebhook) GetHeadCommitOk() (*GithubWebhookHeadCommit, bool)`

GetHeadCommitOk returns a tuple with the HeadCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadCommit

`func (o *GithubWebhook) SetHeadCommit(v GithubWebhookHeadCommit)`

SetHeadCommit sets HeadCommit field to given value.

### HasHeadCommit

`func (o *GithubWebhook) HasHeadCommit() bool`

HasHeadCommit returns a boolean if a field has been set.

### GetCommits

`func (o *GithubWebhook) GetCommits() []GithubCommit`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *GithubWebhook) GetCommitsOk() (*[]GithubCommit, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *GithubWebhook) SetCommits(v []GithubCommit)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *GithubWebhook) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetForced

`func (o *GithubWebhook) GetForced() bool`

GetForced returns the Forced field if non-nil, zero value otherwise.

### GetForcedOk

`func (o *GithubWebhook) GetForcedOk() (*bool, bool)`

GetForcedOk returns a tuple with the Forced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForced

`func (o *GithubWebhook) SetForced(v bool)`

SetForced sets Forced field to given value.

### HasForced

`func (o *GithubWebhook) HasForced() bool`

HasForced returns a boolean if a field has been set.

### GetCompare

`func (o *GithubWebhook) GetCompare() string`

GetCompare returns the Compare field if non-nil, zero value otherwise.

### GetCompareOk

`func (o *GithubWebhook) GetCompareOk() (*string, bool)`

GetCompareOk returns a tuple with the Compare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompare

`func (o *GithubWebhook) SetCompare(v string)`

SetCompare sets Compare field to given value.

### HasCompare

`func (o *GithubWebhook) HasCompare() bool`

HasCompare returns a boolean if a field has been set.

### GetReview

`func (o *GithubWebhook) GetReview() GithubWebhookReview`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *GithubWebhook) GetReviewOk() (*GithubWebhookReview, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *GithubWebhook) SetReview(v GithubWebhookReview)`

SetReview sets Review field to given value.

### HasReview

`func (o *GithubWebhook) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetCheckRun

`func (o *GithubWebhook) GetCheckRun() GithubWebhookCheckRun`

GetCheckRun returns the CheckRun field if non-nil, zero value otherwise.

### GetCheckRunOk

`func (o *GithubWebhook) GetCheckRunOk() (*GithubWebhookCheckRun, bool)`

GetCheckRunOk returns a tuple with the CheckRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckRun

`func (o *GithubWebhook) SetCheckRun(v GithubWebhookCheckRun)`

SetCheckRun sets CheckRun field to given value.

### HasCheckRun

`func (o *GithubWebhook) HasCheckRun() bool`

HasCheckRun returns a boolean if a field has been set.

### GetCheckSuite

`func (o *GithubWebhook) GetCheckSuite() GithubWebhookCheckSuite`

GetCheckSuite returns the CheckSuite field if non-nil, zero value otherwise.

### GetCheckSuiteOk

`func (o *GithubWebhook) GetCheckSuiteOk() (*GithubWebhookCheckSuite, bool)`

GetCheckSuiteOk returns a tuple with the CheckSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSuite

`func (o *GithubWebhook) SetCheckSuite(v GithubWebhookCheckSuite)`

SetCheckSuite sets CheckSuite field to given value.

### HasCheckSuite

`func (o *GithubWebhook) HasCheckSuite() bool`

HasCheckSuite returns a boolean if a field has been set.

### GetDiscussion

`func (o *GithubWebhook) GetDiscussion() GithubWebhookDiscussion`

GetDiscussion returns the Discussion field if non-nil, zero value otherwise.

### GetDiscussionOk

`func (o *GithubWebhook) GetDiscussionOk() (*GithubWebhookDiscussion, bool)`

GetDiscussionOk returns a tuple with the Discussion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscussion

`func (o *GithubWebhook) SetDiscussion(v GithubWebhookDiscussion)`

SetDiscussion sets Discussion field to given value.

### HasDiscussion

`func (o *GithubWebhook) HasDiscussion() bool`

HasDiscussion returns a boolean if a field has been set.

### GetAnswer

`func (o *GithubWebhook) GetAnswer() GithubWebhookComment`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *GithubWebhook) GetAnswerOk() (*GithubWebhookComment, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *GithubWebhook) SetAnswer(v GithubWebhookComment)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *GithubWebhook) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


