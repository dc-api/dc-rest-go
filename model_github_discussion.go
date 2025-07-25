/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-25T15:01:17.719932686Z[Etc/UTC]
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

// checks if the GithubDiscussion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubDiscussion{}

// GithubDiscussion struct for GithubDiscussion
type GithubDiscussion struct {
	Title string `json:"title"`
	Number int32 `json:"number"`
	HtmlUrl string `json:"html_url"`
	AnswerHtmlUrl NullableString `json:"answer_html_url,omitempty"`
	Body NullableString `json:"body,omitempty"`
	User GithubUser `json:"user"`
}

type _GithubDiscussion GithubDiscussion

// NewGithubDiscussion instantiates a new GithubDiscussion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubDiscussion(title string, number int32, htmlUrl string, user GithubUser) *GithubDiscussion {
	this := GithubDiscussion{}
	this.Title = title
	this.Number = number
	this.HtmlUrl = htmlUrl
	this.User = user
	return &this
}

// NewGithubDiscussionWithDefaults instantiates a new GithubDiscussion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubDiscussionWithDefaults() *GithubDiscussion {
	this := GithubDiscussion{}
	return &this
}

// GetTitle returns the Title field value
func (o *GithubDiscussion) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *GithubDiscussion) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *GithubDiscussion) SetTitle(v string) {
	o.Title = v
}

// GetNumber returns the Number field value
func (o *GithubDiscussion) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *GithubDiscussion) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *GithubDiscussion) SetNumber(v int32) {
	o.Number = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *GithubDiscussion) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *GithubDiscussion) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *GithubDiscussion) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetAnswerHtmlUrl returns the AnswerHtmlUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubDiscussion) GetAnswerHtmlUrl() string {
	if o == nil || IsNil(o.AnswerHtmlUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AnswerHtmlUrl.Get()
}

// GetAnswerHtmlUrlOk returns a tuple with the AnswerHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubDiscussion) GetAnswerHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AnswerHtmlUrl.Get()) {
		return nil, false
	}
	return o.AnswerHtmlUrl.Get(), o.AnswerHtmlUrl.IsSet()
}

// HasAnswerHtmlUrl returns a boolean if a field has been set.
func (o *GithubDiscussion) HasAnswerHtmlUrl() bool {
	if o != nil && o.AnswerHtmlUrl.IsSet() {
		return true
	}

	return false
}

// SetAnswerHtmlUrl gets a reference to the given NullableString and assigns it to the AnswerHtmlUrl field.
func (o *GithubDiscussion) SetAnswerHtmlUrl(v string) {
	o.AnswerHtmlUrl.Set(&v)
}

// SetAnswerHtmlUrlNil sets the value for AnswerHtmlUrl to be an explicit nil
func (o *GithubDiscussion) SetAnswerHtmlUrlNil() {
	o.AnswerHtmlUrl.Set(nil)
}

// UnsetAnswerHtmlUrl ensures that no value is present for AnswerHtmlUrl, not even an explicit nil
func (o *GithubDiscussion) UnsetAnswerHtmlUrl() {
	o.AnswerHtmlUrl.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GithubDiscussion) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GithubDiscussion) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body.Get()) {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *GithubDiscussion) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *GithubDiscussion) SetBody(v string) {
	o.Body.Set(&v)
}

// SetBodyNil sets the value for Body to be an explicit nil
func (o *GithubDiscussion) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *GithubDiscussion) UnsetBody() {
	o.Body.Unset()
}

// GetUser returns the User field value
func (o *GithubDiscussion) GetUser() GithubUser {
	if o == nil {
		var ret GithubUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GithubDiscussion) GetUserOk() (*GithubUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GithubDiscussion) SetUser(v GithubUser) {
	o.User = v
}

func (o GithubDiscussion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubDiscussion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["number"] = o.Number
	toSerialize["html_url"] = o.HtmlUrl
	if o.AnswerHtmlUrl.IsSet() {
		toSerialize["answer_html_url"] = o.AnswerHtmlUrl.Get()
	}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *GithubDiscussion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"number",
		"html_url",
		"user",
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

	varGithubDiscussion := _GithubDiscussion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGithubDiscussion)

	if err != nil {
		return err
	}

	*o = GithubDiscussion(varGithubDiscussion)

	return err
}

type NullableGithubDiscussion struct {
	value *GithubDiscussion
	isSet bool
}

func (v NullableGithubDiscussion) Get() *GithubDiscussion {
	return v.value
}

func (v *NullableGithubDiscussion) Set(val *GithubDiscussion) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubDiscussion) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubDiscussion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubDiscussion(val *GithubDiscussion) *NullableGithubDiscussion {
	return &NullableGithubDiscussion{value: val, isSet: true}
}

func (v NullableGithubDiscussion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubDiscussion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


