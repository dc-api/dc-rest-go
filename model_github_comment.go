/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-19T09:30:49.800547817Z[Etc/UTC]
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

// checks if the GithubComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubComment{}

// GithubComment struct for GithubComment
type GithubComment struct {
	Id int32 `json:"id"`
	HtmlUrl string `json:"html_url"`
	User GithubUser `json:"user"`
	CommitId NullableString `json:"commit_id,omitempty"`
	Body string `json:"body"`
}

type _GithubComment GithubComment

// NewGithubComment instantiates a new GithubComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComment(id int32, htmlUrl string, user GithubUser, body string) *GithubComment {
	this := GithubComment{}
	this.Id = id
	this.HtmlUrl = htmlUrl
	this.User = user
	this.Body = body
	return &this
}

// NewGithubCommentWithDefaults instantiates a new GithubComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubCommentWithDefaults() *GithubComment {
	this := GithubComment{}
	return &this
}

// GetId returns the Id field value
func (o *GithubComment) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GithubComment) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GithubComment) SetId(v int32) {
	o.Id = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *GithubComment) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *GithubComment) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *GithubComment) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetUser returns the User field value
func (o *GithubComment) GetUser() GithubUser {
	if o == nil {
		var ret GithubUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GithubComment) GetUserOk() (*GithubUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GithubComment) SetUser(v GithubUser) {
	o.User = v
}

// GetCommitId returns the CommitId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubComment) GetCommitId() string {
	if o == nil || IsNil(o.CommitId.Get()) {
		var ret string
		return ret
	}
	return *o.CommitId.Get()
}

// GetCommitIdOk returns a tuple with the CommitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubComment) GetCommitIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommitId.Get()) {
		return nil, false
	}
	return o.CommitId.Get(), o.CommitId.IsSet()
}

// HasCommitId returns a boolean if a field has been set.
func (o *GithubComment) HasCommitId() bool {
	if o != nil && o.CommitId.IsSet() {
		return true
	}

	return false
}

// SetCommitId gets a reference to the given NullableString and assigns it to the CommitId field.
func (o *GithubComment) SetCommitId(v string) {
	o.CommitId.Set(&v)
}

// SetCommitIdNil sets the value for CommitId to be an explicit nil
func (o *GithubComment) SetCommitIdNil() {
	o.CommitId.Set(nil)
}

// UnsetCommitId ensures that no value is present for CommitId, not even an explicit nil
func (o *GithubComment) UnsetCommitId() {
	o.CommitId.Unset()
}

// GetBody returns the Body field value
func (o *GithubComment) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *GithubComment) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *GithubComment) SetBody(v string) {
	o.Body = v
}

func (o GithubComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["html_url"] = o.HtmlUrl
	toSerialize["user"] = o.User
	if o.CommitId.IsSet() {
		toSerialize["commit_id"] = o.CommitId.Get()
	}
	toSerialize["body"] = o.Body
	return toSerialize, nil
}

func (o *GithubComment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"html_url",
		"user",
		"body",
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

	varGithubComment := _GithubComment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGithubComment)

	if err != nil {
		return err
	}

	*o = GithubComment(varGithubComment)

	return err
}

type NullableGithubComment struct {
	value *GithubComment
	isSet bool
}

func (v NullableGithubComment) Get() *GithubComment {
	return v.value
}

func (v *NullableGithubComment) Set(val *GithubComment) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComment) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComment(val *GithubComment) *NullableGithubComment {
	return &NullableGithubComment{value: val, isSet: true}
}

func (v NullableGithubComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


