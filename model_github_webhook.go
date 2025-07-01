/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:27:32.119933921Z[Etc/UTC]
 *    * - **Generator Version**: 7.14.0
 * 
 * <details>
 * <summary><strong>⚠️ Important Disclaimer & Limitation of Liability</strong></summary>
 * <br>
 * > **IMPORTANT**: This software is provided "as is" without any warranties, express or implied, including but not limited
 * > to warranties of merchantability, fitness for a particular purpose, or non-infringement. The developers, contributors,
 * > and licensors (collectively, "Developers") make no representations regarding the accuracy, completeness, or reliability
 * > of this software or its outputs.
 * > 
 * > This client is not intended to provide financial, investment, tax, or legal advice. It facilitates interaction with the
 * > Discord HTTP API (Preview) service but does not endorse or recommend any financial actions, including the purchase, sale, or holding of
 * > financial instruments (e.g., stocks, bonds, derivatives, cryptocurrencies). Users must consult qualified financial or
 * > legal professionals before making decisions based on this software's outputs.
 * > 
 * > Financial markets are inherently speculative and carry significant risks. Using this software in trading, analysis, or
 * > other financial activities may result in substantial losses, including total loss of capital. The Developers are not
 * > liable for any losses or damages arising from such use. Users assume full responsibility for validating the software's
 * > outputs and ensuring their suitability for intended purposes.
 * > 
 * > This client may rely on third-party data or services (e.g., market feeds, APIs). The Developers do not control or verify
 * > the accuracy of these services and are not liable for any errors, delays, or losses resulting from their use. Users must
 * > comply with third-party terms and conditions.
 * > 
 * > Users are solely responsible for ensuring compliance with all applicable financial, tax, and regulatory requirements in
 * > their jurisdiction. This includes obtaining necessary licenses or approvals for trading or investment activities. The
 * > Developers disclaim liability for any legal consequences arising from non-compliance.
 * > 
 * > To the fullest extent permitted by law, the Developers shall not be liable for any direct, indirect, incidental,
 * > consequential, or punitive damages arising from the use or inability to use this software, including but not limited to
 * > loss of profits, data, or business opportunities.
 * 
 * </details>
 */

package dc_rest

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GithubWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubWebhook{}

// GithubWebhook struct for GithubWebhook
type GithubWebhook struct {
	Action NullableString `json:"action,omitempty"`
	Ref NullableString `json:"ref,omitempty"`
	RefType NullableString `json:"ref_type,omitempty"`
	Comment NullableGithubComment `json:"comment,omitempty"`
	Issue NullableGithubIssue `json:"issue,omitempty"`
	PullRequest NullableGithubIssue `json:"pull_request,omitempty"`
	Repository NullableGithubRepository `json:"repository,omitempty"`
	Forkee NullableGithubRepository `json:"forkee,omitempty"`
	Sender GithubUser `json:"sender"`
	Member NullableGithubUser `json:"member,omitempty"`
	Release NullableGithubRelease `json:"release,omitempty"`
	HeadCommit NullableGithubCommit `json:"head_commit,omitempty"`
	Commits []GithubCommit `json:"commits,omitempty"`
	Forced NullableBool `json:"forced,omitempty"`
	Compare NullableString `json:"compare,omitempty"`
	Review NullableGithubReview `json:"review,omitempty"`
	CheckRun NullableGithubCheckRun `json:"check_run,omitempty"`
	CheckSuite NullableGithubCheckSuite `json:"check_suite,omitempty"`
	Discussion NullableGithubDiscussion `json:"discussion,omitempty"`
	Answer NullableGithubComment `json:"answer,omitempty"`
}

type _GithubWebhook GithubWebhook

// NewGithubWebhook instantiates a new GithubWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubWebhook(sender GithubUser) *GithubWebhook {
	this := GithubWebhook{}
	this.Sender = sender
	return &this
}

// NewGithubWebhookWithDefaults instantiates a new GithubWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubWebhookWithDefaults() *GithubWebhook {
	this := GithubWebhook{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action.Get()) {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *GithubWebhook) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *GithubWebhook) SetAction(v string) {
	o.Action.Set(&v)
}

// SetActionNil sets the value for Action to be an explicit nil
func (o *GithubWebhook) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *GithubWebhook) UnsetAction() {
	o.Action.Unset()
}

// GetRef returns the Ref field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetRef() string {
	if o == nil || IsNil(o.Ref.Get()) {
		var ret string
		return ret
	}
	return *o.Ref.Get()
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref.Get()) {
		return nil, false
	}
	return o.Ref.Get(), o.Ref.IsSet()
}

// HasRef returns a boolean if a field has been set.
func (o *GithubWebhook) HasRef() bool {
	if o != nil && o.Ref.IsSet() {
		return true
	}

	return false
}

// SetRef gets a reference to the given NullableString and assigns it to the Ref field.
func (o *GithubWebhook) SetRef(v string) {
	o.Ref.Set(&v)
}

// SetRefNil sets the value for Ref to be an explicit nil
func (o *GithubWebhook) SetRefNil() {
	o.Ref.Set(nil)
}

// UnsetRef ensures that no value is present for Ref, not even an explicit nil
func (o *GithubWebhook) UnsetRef() {
	o.Ref.Unset()
}

// GetRefType returns the RefType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetRefType() string {
	if o == nil || IsNil(o.RefType.Get()) {
		var ret string
		return ret
	}
	return *o.RefType.Get()
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetRefTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RefType.Get()) {
		return nil, false
	}
	return o.RefType.Get(), o.RefType.IsSet()
}

// HasRefType returns a boolean if a field has been set.
func (o *GithubWebhook) HasRefType() bool {
	if o != nil && o.RefType.IsSet() {
		return true
	}

	return false
}

// SetRefType gets a reference to the given NullableString and assigns it to the RefType field.
func (o *GithubWebhook) SetRefType(v string) {
	o.RefType.Set(&v)
}

// SetRefTypeNil sets the value for RefType to be an explicit nil
func (o *GithubWebhook) SetRefTypeNil() {
	o.RefType.Set(nil)
}

// UnsetRefType ensures that no value is present for RefType, not even an explicit nil
func (o *GithubWebhook) UnsetRefType() {
	o.RefType.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetComment() GithubComment {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret GithubComment
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetCommentOk() (*GithubComment, bool) {
	if o == nil || IsNil(o.Comment.Get()) {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *GithubWebhook) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableGithubComment and assigns it to the Comment field.
func (o *GithubWebhook) SetComment(v GithubComment) {
	o.Comment.Set(&v)
}

// SetCommentNil sets the value for Comment to be an explicit nil
func (o *GithubWebhook) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *GithubWebhook) UnsetComment() {
	o.Comment.Unset()
}

// GetIssue returns the Issue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetIssue() GithubIssue {
	if o == nil || IsNil(o.Issue.Get()) {
		var ret GithubIssue
		return ret
	}
	return *o.Issue.Get()
}

// GetIssueOk returns a tuple with the Issue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetIssueOk() (*GithubIssue, bool) {
	if o == nil || IsNil(o.Issue.Get()) {
		return nil, false
	}
	return o.Issue.Get(), o.Issue.IsSet()
}

// HasIssue returns a boolean if a field has been set.
func (o *GithubWebhook) HasIssue() bool {
	if o != nil && o.Issue.IsSet() {
		return true
	}

	return false
}

// SetIssue gets a reference to the given NullableGithubIssue and assigns it to the Issue field.
func (o *GithubWebhook) SetIssue(v GithubIssue) {
	o.Issue.Set(&v)
}

// SetIssueNil sets the value for Issue to be an explicit nil
func (o *GithubWebhook) SetIssueNil() {
	o.Issue.Set(nil)
}

// UnsetIssue ensures that no value is present for Issue, not even an explicit nil
func (o *GithubWebhook) UnsetIssue() {
	o.Issue.Unset()
}

// GetPullRequest returns the PullRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetPullRequest() GithubIssue {
	if o == nil || IsNil(o.PullRequest.Get()) {
		var ret GithubIssue
		return ret
	}
	return *o.PullRequest.Get()
}

// GetPullRequestOk returns a tuple with the PullRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetPullRequestOk() (*GithubIssue, bool) {
	if o == nil || IsNil(o.PullRequest.Get()) {
		return nil, false
	}
	return o.PullRequest.Get(), o.PullRequest.IsSet()
}

// HasPullRequest returns a boolean if a field has been set.
func (o *GithubWebhook) HasPullRequest() bool {
	if o != nil && o.PullRequest.IsSet() {
		return true
	}

	return false
}

// SetPullRequest gets a reference to the given NullableGithubIssue and assigns it to the PullRequest field.
func (o *GithubWebhook) SetPullRequest(v GithubIssue) {
	o.PullRequest.Set(&v)
}

// SetPullRequestNil sets the value for PullRequest to be an explicit nil
func (o *GithubWebhook) SetPullRequestNil() {
	o.PullRequest.Set(nil)
}

// UnsetPullRequest ensures that no value is present for PullRequest, not even an explicit nil
func (o *GithubWebhook) UnsetPullRequest() {
	o.PullRequest.Unset()
}

// GetRepository returns the Repository field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetRepository() GithubRepository {
	if o == nil || IsNil(o.Repository.Get()) {
		var ret GithubRepository
		return ret
	}
	return *o.Repository.Get()
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetRepositoryOk() (*GithubRepository, bool) {
	if o == nil || IsNil(o.Repository.Get()) {
		return nil, false
	}
	return o.Repository.Get(), o.Repository.IsSet()
}

// HasRepository returns a boolean if a field has been set.
func (o *GithubWebhook) HasRepository() bool {
	if o != nil && o.Repository.IsSet() {
		return true
	}

	return false
}

// SetRepository gets a reference to the given NullableGithubRepository and assigns it to the Repository field.
func (o *GithubWebhook) SetRepository(v GithubRepository) {
	o.Repository.Set(&v)
}

// SetRepositoryNil sets the value for Repository to be an explicit nil
func (o *GithubWebhook) SetRepositoryNil() {
	o.Repository.Set(nil)
}

// UnsetRepository ensures that no value is present for Repository, not even an explicit nil
func (o *GithubWebhook) UnsetRepository() {
	o.Repository.Unset()
}

// GetForkee returns the Forkee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetForkee() GithubRepository {
	if o == nil || IsNil(o.Forkee.Get()) {
		var ret GithubRepository
		return ret
	}
	return *o.Forkee.Get()
}

// GetForkeeOk returns a tuple with the Forkee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetForkeeOk() (*GithubRepository, bool) {
	if o == nil || IsNil(o.Forkee.Get()) {
		return nil, false
	}
	return o.Forkee.Get(), o.Forkee.IsSet()
}

// HasForkee returns a boolean if a field has been set.
func (o *GithubWebhook) HasForkee() bool {
	if o != nil && o.Forkee.IsSet() {
		return true
	}

	return false
}

// SetForkee gets a reference to the given NullableGithubRepository and assigns it to the Forkee field.
func (o *GithubWebhook) SetForkee(v GithubRepository) {
	o.Forkee.Set(&v)
}

// SetForkeeNil sets the value for Forkee to be an explicit nil
func (o *GithubWebhook) SetForkeeNil() {
	o.Forkee.Set(nil)
}

// UnsetForkee ensures that no value is present for Forkee, not even an explicit nil
func (o *GithubWebhook) UnsetForkee() {
	o.Forkee.Unset()
}

// GetSender returns the Sender field value
func (o *GithubWebhook) GetSender() GithubUser {
	if o == nil {
		var ret GithubUser
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *GithubWebhook) GetSenderOk() (*GithubUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *GithubWebhook) SetSender(v GithubUser) {
	o.Sender = v
}

// GetMember returns the Member field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetMember() GithubUser {
	if o == nil || IsNil(o.Member.Get()) {
		var ret GithubUser
		return ret
	}
	return *o.Member.Get()
}

// GetMemberOk returns a tuple with the Member field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetMemberOk() (*GithubUser, bool) {
	if o == nil || IsNil(o.Member.Get()) {
		return nil, false
	}
	return o.Member.Get(), o.Member.IsSet()
}

// HasMember returns a boolean if a field has been set.
func (o *GithubWebhook) HasMember() bool {
	if o != nil && o.Member.IsSet() {
		return true
	}

	return false
}

// SetMember gets a reference to the given NullableGithubUser and assigns it to the Member field.
func (o *GithubWebhook) SetMember(v GithubUser) {
	o.Member.Set(&v)
}

// SetMemberNil sets the value for Member to be an explicit nil
func (o *GithubWebhook) SetMemberNil() {
	o.Member.Set(nil)
}

// UnsetMember ensures that no value is present for Member, not even an explicit nil
func (o *GithubWebhook) UnsetMember() {
	o.Member.Unset()
}

// GetRelease returns the Release field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetRelease() GithubRelease {
	if o == nil || IsNil(o.Release.Get()) {
		var ret GithubRelease
		return ret
	}
	return *o.Release.Get()
}

// GetReleaseOk returns a tuple with the Release field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetReleaseOk() (*GithubRelease, bool) {
	if o == nil || IsNil(o.Release.Get()) {
		return nil, false
	}
	return o.Release.Get(), o.Release.IsSet()
}

// HasRelease returns a boolean if a field has been set.
func (o *GithubWebhook) HasRelease() bool {
	if o != nil && o.Release.IsSet() {
		return true
	}

	return false
}

// SetRelease gets a reference to the given NullableGithubRelease and assigns it to the Release field.
func (o *GithubWebhook) SetRelease(v GithubRelease) {
	o.Release.Set(&v)
}

// SetReleaseNil sets the value for Release to be an explicit nil
func (o *GithubWebhook) SetReleaseNil() {
	o.Release.Set(nil)
}

// UnsetRelease ensures that no value is present for Release, not even an explicit nil
func (o *GithubWebhook) UnsetRelease() {
	o.Release.Unset()
}

// GetHeadCommit returns the HeadCommit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetHeadCommit() GithubCommit {
	if o == nil || IsNil(o.HeadCommit.Get()) {
		var ret GithubCommit
		return ret
	}
	return *o.HeadCommit.Get()
}

// GetHeadCommitOk returns a tuple with the HeadCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetHeadCommitOk() (*GithubCommit, bool) {
	if o == nil || IsNil(o.HeadCommit.Get()) {
		return nil, false
	}
	return o.HeadCommit.Get(), o.HeadCommit.IsSet()
}

// HasHeadCommit returns a boolean if a field has been set.
func (o *GithubWebhook) HasHeadCommit() bool {
	if o != nil && o.HeadCommit.IsSet() {
		return true
	}

	return false
}

// SetHeadCommit gets a reference to the given NullableGithubCommit and assigns it to the HeadCommit field.
func (o *GithubWebhook) SetHeadCommit(v GithubCommit) {
	o.HeadCommit.Set(&v)
}

// SetHeadCommitNil sets the value for HeadCommit to be an explicit nil
func (o *GithubWebhook) SetHeadCommitNil() {
	o.HeadCommit.Set(nil)
}

// UnsetHeadCommit ensures that no value is present for HeadCommit, not even an explicit nil
func (o *GithubWebhook) UnsetHeadCommit() {
	o.HeadCommit.Unset()
}

// GetCommits returns the Commits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetCommits() []GithubCommit {
	if o == nil {
		var ret []GithubCommit
		return ret
	}
	return o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetCommitsOk() ([]GithubCommit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Commits, true
}

// HasCommits returns a boolean if a field has been set.
func (o *GithubWebhook) HasCommits() bool {
	if o != nil && !IsNil(o.Commits) {
		return true
	}

	return false
}

// SetCommits gets a reference to the given []GithubCommit and assigns it to the Commits field.
func (o *GithubWebhook) SetCommits(v []GithubCommit) {
	o.Commits = v
}


// GetForced returns the Forced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetForced() bool {
	if o == nil || IsNil(o.Forced.Get()) {
		var ret bool
		return ret
	}
	return *o.Forced.Get()
}

// GetForcedOk returns a tuple with the Forced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetForcedOk() (*bool, bool) {
	if o == nil || IsNil(o.Forced.Get()) {
		return nil, false
	}
	return o.Forced.Get(), o.Forced.IsSet()
}

// HasForced returns a boolean if a field has been set.
func (o *GithubWebhook) HasForced() bool {
	if o != nil && o.Forced.IsSet() {
		return true
	}

	return false
}

// SetForced gets a reference to the given NullableBool and assigns it to the Forced field.
func (o *GithubWebhook) SetForced(v bool) {
	o.Forced.Set(&v)
}

// SetForcedNil sets the value for Forced to be an explicit nil
func (o *GithubWebhook) SetForcedNil() {
	o.Forced.Set(nil)
}

// UnsetForced ensures that no value is present for Forced, not even an explicit nil
func (o *GithubWebhook) UnsetForced() {
	o.Forced.Unset()
}

// GetCompare returns the Compare field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetCompare() string {
	if o == nil || IsNil(o.Compare.Get()) {
		var ret string
		return ret
	}
	return *o.Compare.Get()
}

// GetCompareOk returns a tuple with the Compare field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetCompareOk() (*string, bool) {
	if o == nil || IsNil(o.Compare.Get()) {
		return nil, false
	}
	return o.Compare.Get(), o.Compare.IsSet()
}

// HasCompare returns a boolean if a field has been set.
func (o *GithubWebhook) HasCompare() bool {
	if o != nil && o.Compare.IsSet() {
		return true
	}

	return false
}

// SetCompare gets a reference to the given NullableString and assigns it to the Compare field.
func (o *GithubWebhook) SetCompare(v string) {
	o.Compare.Set(&v)
}

// SetCompareNil sets the value for Compare to be an explicit nil
func (o *GithubWebhook) SetCompareNil() {
	o.Compare.Set(nil)
}

// UnsetCompare ensures that no value is present for Compare, not even an explicit nil
func (o *GithubWebhook) UnsetCompare() {
	o.Compare.Unset()
}

// GetReview returns the Review field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetReview() GithubReview {
	if o == nil || IsNil(o.Review.Get()) {
		var ret GithubReview
		return ret
	}
	return *o.Review.Get()
}

// GetReviewOk returns a tuple with the Review field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetReviewOk() (*GithubReview, bool) {
	if o == nil || IsNil(o.Review.Get()) {
		return nil, false
	}
	return o.Review.Get(), o.Review.IsSet()
}

// HasReview returns a boolean if a field has been set.
func (o *GithubWebhook) HasReview() bool {
	if o != nil && o.Review.IsSet() {
		return true
	}

	return false
}

// SetReview gets a reference to the given NullableGithubReview and assigns it to the Review field.
func (o *GithubWebhook) SetReview(v GithubReview) {
	o.Review.Set(&v)
}

// SetReviewNil sets the value for Review to be an explicit nil
func (o *GithubWebhook) SetReviewNil() {
	o.Review.Set(nil)
}

// UnsetReview ensures that no value is present for Review, not even an explicit nil
func (o *GithubWebhook) UnsetReview() {
	o.Review.Unset()
}

// GetCheckRun returns the CheckRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetCheckRun() GithubCheckRun {
	if o == nil || IsNil(o.CheckRun.Get()) {
		var ret GithubCheckRun
		return ret
	}
	return *o.CheckRun.Get()
}

// GetCheckRunOk returns a tuple with the CheckRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetCheckRunOk() (*GithubCheckRun, bool) {
	if o == nil || IsNil(o.CheckRun.Get()) {
		return nil, false
	}
	return o.CheckRun.Get(), o.CheckRun.IsSet()
}

// HasCheckRun returns a boolean if a field has been set.
func (o *GithubWebhook) HasCheckRun() bool {
	if o != nil && o.CheckRun.IsSet() {
		return true
	}

	return false
}

// SetCheckRun gets a reference to the given NullableGithubCheckRun and assigns it to the CheckRun field.
func (o *GithubWebhook) SetCheckRun(v GithubCheckRun) {
	o.CheckRun.Set(&v)
}

// SetCheckRunNil sets the value for CheckRun to be an explicit nil
func (o *GithubWebhook) SetCheckRunNil() {
	o.CheckRun.Set(nil)
}

// UnsetCheckRun ensures that no value is present for CheckRun, not even an explicit nil
func (o *GithubWebhook) UnsetCheckRun() {
	o.CheckRun.Unset()
}

// GetCheckSuite returns the CheckSuite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetCheckSuite() GithubCheckSuite {
	if o == nil || IsNil(o.CheckSuite.Get()) {
		var ret GithubCheckSuite
		return ret
	}
	return *o.CheckSuite.Get()
}

// GetCheckSuiteOk returns a tuple with the CheckSuite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetCheckSuiteOk() (*GithubCheckSuite, bool) {
	if o == nil || IsNil(o.CheckSuite.Get()) {
		return nil, false
	}
	return o.CheckSuite.Get(), o.CheckSuite.IsSet()
}

// HasCheckSuite returns a boolean if a field has been set.
func (o *GithubWebhook) HasCheckSuite() bool {
	if o != nil && o.CheckSuite.IsSet() {
		return true
	}

	return false
}

// SetCheckSuite gets a reference to the given NullableGithubCheckSuite and assigns it to the CheckSuite field.
func (o *GithubWebhook) SetCheckSuite(v GithubCheckSuite) {
	o.CheckSuite.Set(&v)
}

// SetCheckSuiteNil sets the value for CheckSuite to be an explicit nil
func (o *GithubWebhook) SetCheckSuiteNil() {
	o.CheckSuite.Set(nil)
}

// UnsetCheckSuite ensures that no value is present for CheckSuite, not even an explicit nil
func (o *GithubWebhook) UnsetCheckSuite() {
	o.CheckSuite.Unset()
}

// GetDiscussion returns the Discussion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetDiscussion() GithubDiscussion {
	if o == nil || IsNil(o.Discussion.Get()) {
		var ret GithubDiscussion
		return ret
	}
	return *o.Discussion.Get()
}

// GetDiscussionOk returns a tuple with the Discussion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetDiscussionOk() (*GithubDiscussion, bool) {
	if o == nil || IsNil(o.Discussion.Get()) {
		return nil, false
	}
	return o.Discussion.Get(), o.Discussion.IsSet()
}

// HasDiscussion returns a boolean if a field has been set.
func (o *GithubWebhook) HasDiscussion() bool {
	if o != nil && o.Discussion.IsSet() {
		return true
	}

	return false
}

// SetDiscussion gets a reference to the given NullableGithubDiscussion and assigns it to the Discussion field.
func (o *GithubWebhook) SetDiscussion(v GithubDiscussion) {
	o.Discussion.Set(&v)
}

// SetDiscussionNil sets the value for Discussion to be an explicit nil
func (o *GithubWebhook) SetDiscussionNil() {
	o.Discussion.Set(nil)
}

// UnsetDiscussion ensures that no value is present for Discussion, not even an explicit nil
func (o *GithubWebhook) UnsetDiscussion() {
	o.Discussion.Unset()
}

// GetAnswer returns the Answer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubWebhook) GetAnswer() GithubComment {
	if o == nil || IsNil(o.Answer.Get()) {
		var ret GithubComment
		return ret
	}
	return *o.Answer.Get()
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubWebhook) GetAnswerOk() (*GithubComment, bool) {
	if o == nil || IsNil(o.Answer.Get()) {
		return nil, false
	}
	return o.Answer.Get(), o.Answer.IsSet()
}

// HasAnswer returns a boolean if a field has been set.
func (o *GithubWebhook) HasAnswer() bool {
	if o != nil && o.Answer.IsSet() {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given NullableGithubComment and assigns it to the Answer field.
func (o *GithubWebhook) SetAnswer(v GithubComment) {
	o.Answer.Set(&v)
}

// SetAnswerNil sets the value for Answer to be an explicit nil
func (o *GithubWebhook) SetAnswerNil() {
	o.Answer.Set(nil)
}

// UnsetAnswer ensures that no value is present for Answer, not even an explicit nil
func (o *GithubWebhook) UnsetAnswer() {
	o.Answer.Unset()
}

func (o GithubWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.Ref.IsSet() {
		toSerialize["ref"] = o.Ref.Get()
	}
	if o.RefType.IsSet() {
		toSerialize["ref_type"] = o.RefType.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Issue.IsSet() {
		toSerialize["issue"] = o.Issue.Get()
	}
	if o.PullRequest.IsSet() {
		toSerialize["pull_request"] = o.PullRequest.Get()
	}
	if o.Repository.IsSet() {
		toSerialize["repository"] = o.Repository.Get()
	}
	if o.Forkee.IsSet() {
		toSerialize["forkee"] = o.Forkee.Get()
	}
	toSerialize["sender"] = o.Sender
	if o.Member.IsSet() {
		toSerialize["member"] = o.Member.Get()
	}
	if o.Release.IsSet() {
		toSerialize["release"] = o.Release.Get()
	}
	if o.HeadCommit.IsSet() {
		toSerialize["head_commit"] = o.HeadCommit.Get()
	}
	if o.Commits != nil {
		toSerialize["commits"] = o.Commits
	}
	if o.Forced.IsSet() {
		toSerialize["forced"] = o.Forced.Get()
	}
	if o.Compare.IsSet() {
		toSerialize["compare"] = o.Compare.Get()
	}
	if o.Review.IsSet() {
		toSerialize["review"] = o.Review.Get()
	}
	if o.CheckRun.IsSet() {
		toSerialize["check_run"] = o.CheckRun.Get()
	}
	if o.CheckSuite.IsSet() {
		toSerialize["check_suite"] = o.CheckSuite.Get()
	}
	if o.Discussion.IsSet() {
		toSerialize["discussion"] = o.Discussion.Get()
	}
	if o.Answer.IsSet() {
		toSerialize["answer"] = o.Answer.Get()
	}
	return toSerialize, nil
}

func (o *GithubWebhook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sender",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGithubWebhook := _GithubWebhook{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGithubWebhook)

	if err != nil {
		return err
	}

	*o = GithubWebhook(varGithubWebhook)

	return err
}

type NullableGithubWebhook struct {
	value *GithubWebhook
	isSet bool
}

func (v NullableGithubWebhook) Get() *GithubWebhook {
	return v.value
}

func (v *NullableGithubWebhook) Set(val *GithubWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubWebhook(val *GithubWebhook) *NullableGithubWebhook {
	return &NullableGithubWebhook{value: val, isSet: true}
}

func (v NullableGithubWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


