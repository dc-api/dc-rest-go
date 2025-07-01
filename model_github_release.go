/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T06:33:06.733235362Z[Etc/UTC]
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

// checks if the GithubRelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubRelease{}

// GithubRelease struct for GithubRelease
type GithubRelease struct {
	Id int32 `json:"id"`
	TagName string `json:"tag_name"`
	HtmlUrl string `json:"html_url"`
	Author GithubUser `json:"author"`
}

type _GithubRelease GithubRelease

// NewGithubRelease instantiates a new GithubRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubRelease(id int32, tagName string, htmlUrl string, author GithubUser) *GithubRelease {
	this := GithubRelease{}
	this.Id = id
	this.TagName = tagName
	this.HtmlUrl = htmlUrl
	this.Author = author
	return &this
}

// NewGithubReleaseWithDefaults instantiates a new GithubRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubReleaseWithDefaults() *GithubRelease {
	this := GithubRelease{}
	return &this
}

// GetId returns the Id field value
func (o *GithubRelease) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GithubRelease) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GithubRelease) SetId(v int32) {
	o.Id = v
}

// GetTagName returns the TagName field value
func (o *GithubRelease) GetTagName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TagName
}

// GetTagNameOk returns a tuple with the TagName field value
// and a boolean to check if the value has been set.
func (o *GithubRelease) GetTagNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagName, true
}

// SetTagName sets field value
func (o *GithubRelease) SetTagName(v string) {
	o.TagName = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *GithubRelease) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *GithubRelease) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *GithubRelease) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetAuthor returns the Author field value
func (o *GithubRelease) GetAuthor() GithubUser {
	if o == nil {
		var ret GithubUser
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *GithubRelease) GetAuthorOk() (*GithubUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *GithubRelease) SetAuthor(v GithubUser) {
	o.Author = v
}

func (o GithubRelease) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubRelease) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["tag_name"] = o.TagName
	toSerialize["html_url"] = o.HtmlUrl
	toSerialize["author"] = o.Author
	return toSerialize, nil
}

func (o *GithubRelease) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"tag_name",
		"html_url",
		"author",
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

	varGithubRelease := _GithubRelease{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGithubRelease)

	if err != nil {
		return err
	}

	*o = GithubRelease(varGithubRelease)

	return err
}

type NullableGithubRelease struct {
	value *GithubRelease
	isSet bool
}

func (v NullableGithubRelease) Get() *GithubRelease {
	return v.value
}

func (v *NullableGithubRelease) Set(val *GithubRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubRelease(val *GithubRelease) *NullableGithubRelease {
	return &NullableGithubRelease{value: val, isSet: true}
}

func (v NullableGithubRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


