/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-08-08T14:09:23.736426080Z[Etc/UTC]
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

// checks if the NewMemberActionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewMemberActionResponse{}

// NewMemberActionResponse struct for NewMemberActionResponse
type NewMemberActionResponse struct {
	ChannelId string `json:"channel_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ActionType NullableInt32 `json:"action_type"`
	Title string `json:"title"`
	Description string `json:"description"`
	Emoji NullableSettingsEmojiResponse `json:"emoji,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
}

type _NewMemberActionResponse NewMemberActionResponse

// NewNewMemberActionResponse instantiates a new NewMemberActionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewMemberActionResponse(channelId string, actionType NullableInt32, title string, description string) *NewMemberActionResponse {
	this := NewMemberActionResponse{}
	this.ChannelId = channelId
	this.ActionType = actionType
	this.Title = title
	this.Description = description
	return &this
}

// NewNewMemberActionResponseWithDefaults instantiates a new NewMemberActionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewMemberActionResponseWithDefaults() *NewMemberActionResponse {
	this := NewMemberActionResponse{}
	return &this
}

// GetChannelId returns the ChannelId field value
func (o *NewMemberActionResponse) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *NewMemberActionResponse) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *NewMemberActionResponse) SetChannelId(v string) {
	o.ChannelId = v
}

// GetActionType returns the ActionType field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *NewMemberActionResponse) GetActionType() int32 {
	if o == nil || o.ActionType.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ActionType.Get()
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewMemberActionResponse) GetActionTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionType.Get(), o.ActionType.IsSet()
}

// SetActionType sets field value
func (o *NewMemberActionResponse) SetActionType(v int32) {
	o.ActionType.Set(&v)
}

// GetTitle returns the Title field value
func (o *NewMemberActionResponse) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *NewMemberActionResponse) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *NewMemberActionResponse) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *NewMemberActionResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *NewMemberActionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *NewMemberActionResponse) SetDescription(v string) {
	o.Description = v
}

// GetEmoji returns the Emoji field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NewMemberActionResponse) GetEmoji() SettingsEmojiResponse {
	if o == nil || IsNil(o.Emoji.Get()) {
		var ret SettingsEmojiResponse
		return ret
	}
	return *o.Emoji.Get()
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewMemberActionResponse) GetEmojiOk() (*SettingsEmojiResponse, bool) {
	if o == nil || IsNil(o.Emoji.Get()) {
		return nil, false
	}
	return o.Emoji.Get(), o.Emoji.IsSet()
}

// HasEmoji returns a boolean if a field has been set.
func (o *NewMemberActionResponse) HasEmoji() bool {
	if o != nil && o.Emoji.IsSet() {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given NullableSettingsEmojiResponse and assigns it to the Emoji field.
func (o *NewMemberActionResponse) SetEmoji(v SettingsEmojiResponse) {
	o.Emoji.Set(&v)
}

// SetEmojiNil sets the value for Emoji to be an explicit nil
func (o *NewMemberActionResponse) SetEmojiNil() {
	o.Emoji.Set(nil)
}

// UnsetEmoji ensures that no value is present for Emoji, not even an explicit nil
func (o *NewMemberActionResponse) UnsetEmoji() {
	o.Emoji.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NewMemberActionResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewMemberActionResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *NewMemberActionResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *NewMemberActionResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *NewMemberActionResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *NewMemberActionResponse) UnsetIcon() {
	o.Icon.Unset()
}

func (o NewMemberActionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewMemberActionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channel_id"] = o.ChannelId
	toSerialize["action_type"] = o.ActionType.Get()
	toSerialize["title"] = o.Title
	toSerialize["description"] = o.Description
	if o.Emoji.IsSet() {
		toSerialize["emoji"] = o.Emoji.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	return toSerialize, nil
}

func (o *NewMemberActionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channel_id",
		"action_type",
		"title",
		"description",
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

	varNewMemberActionResponse := _NewMemberActionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNewMemberActionResponse)

	if err != nil {
		return err
	}

	*o = NewMemberActionResponse(varNewMemberActionResponse)

	return err
}

type NullableNewMemberActionResponse struct {
	value *NewMemberActionResponse
	isSet bool
}

func (v NullableNewMemberActionResponse) Get() *NewMemberActionResponse {
	return v.value
}

func (v *NullableNewMemberActionResponse) Set(val *NewMemberActionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNewMemberActionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNewMemberActionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewMemberActionResponse(val *NewMemberActionResponse) *NullableNewMemberActionResponse {
	return &NullableNewMemberActionResponse{value: val, isSet: true}
}

func (v NullableNewMemberActionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewMemberActionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


