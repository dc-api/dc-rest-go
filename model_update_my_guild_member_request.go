/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-08T01:38:41.954257418Z[Etc/UTC]
 *    * - **Generator Version**: 7.18.0
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
)

// checks if the UpdateMyGuildMemberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMyGuildMemberRequest{}

// UpdateMyGuildMemberRequest struct for UpdateMyGuildMemberRequest
type UpdateMyGuildMemberRequest struct {
	Nick *string `json:"nick,omitempty"`
	Avatar *string `json:"avatar,omitempty"`
	Bio *string `json:"bio,omitempty"`
	Banner *string `json:"banner,omitempty"`
}

// NewUpdateMyGuildMemberRequest instantiates a new UpdateMyGuildMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMyGuildMemberRequest() *UpdateMyGuildMemberRequest {
	this := UpdateMyGuildMemberRequest{}
	return &this
}

// NewUpdateMyGuildMemberRequestWithDefaults instantiates a new UpdateMyGuildMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMyGuildMemberRequestWithDefaults() *UpdateMyGuildMemberRequest {
	this := UpdateMyGuildMemberRequest{}
	return &this
}

// GetNick returns the Nick field value if set, zero value otherwise.
func (o *UpdateMyGuildMemberRequest) GetNick() string {
	if o == nil || IsNil(o.Nick) {
		var ret string
		return ret
	}
	return *o.Nick
}

// GetNickOk returns a tuple with the Nick field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMyGuildMemberRequest) GetNickOk() (*string, bool) {
	if o == nil || IsNil(o.Nick) {
		return nil, false
	}
	return o.Nick, true
}

// HasNick returns a boolean if a field has been set.
func (o *UpdateMyGuildMemberRequest) HasNick() bool {
	if o != nil && !IsNil(o.Nick) {
		return true
	}

	return false
}

// SetNick gets a reference to the given string and assigns it to the Nick field.
func (o *UpdateMyGuildMemberRequest) SetNick(v string) {
	o.Nick = &v
}


// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *UpdateMyGuildMemberRequest) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMyGuildMemberRequest) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *UpdateMyGuildMemberRequest) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *UpdateMyGuildMemberRequest) SetAvatar(v string) {
	o.Avatar = &v
}


// GetBio returns the Bio field value if set, zero value otherwise.
func (o *UpdateMyGuildMemberRequest) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMyGuildMemberRequest) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *UpdateMyGuildMemberRequest) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *UpdateMyGuildMemberRequest) SetBio(v string) {
	o.Bio = &v
}


// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *UpdateMyGuildMemberRequest) GetBanner() string {
	if o == nil || IsNil(o.Banner) {
		var ret string
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMyGuildMemberRequest) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner) {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *UpdateMyGuildMemberRequest) HasBanner() bool {
	if o != nil && !IsNil(o.Banner) {
		return true
	}

	return false
}

// SetBanner gets a reference to the given string and assigns it to the Banner field.
func (o *UpdateMyGuildMemberRequest) SetBanner(v string) {
	o.Banner = &v
}


func (o UpdateMyGuildMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMyGuildMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nick) {
		toSerialize["nick"] = o.Nick
	}
	if !IsNil(o.Avatar) {
		toSerialize["avatar"] = o.Avatar
	}
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	if !IsNil(o.Banner) {
		toSerialize["banner"] = o.Banner
	}
	return toSerialize, nil
}

type NullableUpdateMyGuildMemberRequest struct {
	value *UpdateMyGuildMemberRequest
	isSet bool
}

func (v NullableUpdateMyGuildMemberRequest) Get() *UpdateMyGuildMemberRequest {
	return v.value
}

func (v *NullableUpdateMyGuildMemberRequest) Set(val *UpdateMyGuildMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMyGuildMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMyGuildMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMyGuildMemberRequest(val *UpdateMyGuildMemberRequest) *NullableUpdateMyGuildMemberRequest {
	return &NullableUpdateMyGuildMemberRequest{value: val, isSet: true}
}

func (v NullableUpdateMyGuildMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMyGuildMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

