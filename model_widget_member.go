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

// checks if the WidgetMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WidgetMember{}

// WidgetMember struct for WidgetMember
type WidgetMember struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Discriminator NullableString `json:"discriminator"`
	Avatar interface{} `json:"avatar,omitempty"`
	Status string `json:"status"`
	AvatarUrl string `json:"avatar_url"`
	Activity NullableWidgetActivity `json:"activity,omitempty"`
	Deaf NullableBool `json:"deaf,omitempty"`
	Mute NullableBool `json:"mute,omitempty"`
	SelfDeaf NullableBool `json:"self_deaf,omitempty"`
	SelfMute NullableBool `json:"self_mute,omitempty"`
	Suppress NullableBool `json:"suppress,omitempty"`
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _WidgetMember WidgetMember

// NewWidgetMember instantiates a new WidgetMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetMember(id string, username string, discriminator NullableString, status string, avatarUrl string) *WidgetMember {
	this := WidgetMember{}
	this.Id = id
	this.Username = username
	this.Discriminator = discriminator
	this.Status = status
	this.AvatarUrl = avatarUrl
	return &this
}

// NewWidgetMemberWithDefaults instantiates a new WidgetMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetMemberWithDefaults() *WidgetMember {
	this := WidgetMember{}
	return &this
}

// GetId returns the Id field value
func (o *WidgetMember) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WidgetMember) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WidgetMember) SetId(v string) {
	o.Id = v
}

// GetUsername returns the Username field value
func (o *WidgetMember) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *WidgetMember) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *WidgetMember) SetUsername(v string) {
	o.Username = v
}

// GetDiscriminator returns the Discriminator field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WidgetMember) GetDiscriminator() string {
	if o == nil || o.Discriminator.Get() == nil {
		var ret string
		return ret
	}

	return *o.Discriminator.Get()
}

// GetDiscriminatorOk returns a tuple with the Discriminator field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetMember) GetDiscriminatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Discriminator.Get(), o.Discriminator.IsSet()
}

// SetDiscriminator sets field value
func (o *WidgetMember) SetDiscriminator(v string) {
	o.Discriminator.Set(&v)
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetMember) GetAvatar() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetMember) GetAvatarOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *WidgetMember) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given interface{} and assigns it to the Avatar field.
func (o *WidgetMember) SetAvatar(v interface{}) {
	o.Avatar = v
}


// GetStatus returns the Status field value
func (o *WidgetMember) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WidgetMember) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WidgetMember) SetStatus(v string) {
	o.Status = v
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *WidgetMember) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *WidgetMember) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *WidgetMember) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

// GetActivity returns the Activity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetMember) GetActivity() WidgetActivity {
	if o == nil || IsNil(o.Activity.Get()) {
		var ret WidgetActivity
		return ret
	}
	return *o.Activity.Get()
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetMember) GetActivityOk() (*WidgetActivity, bool) {
	if o == nil || IsNil(o.Activity.Get()) {
		return nil, false
	}
	return o.Activity.Get(), o.Activity.IsSet()
}

// HasActivity returns a boolean if a field has been set.
func (o *WidgetMember) HasActivity() bool {
	if o != nil && o.Activity.IsSet() {
		return true
	}

	return false
}

// SetActivity gets a reference to the given NullableWidgetActivity and assigns it to the Activity field.
func (o *WidgetMember) SetActivity(v WidgetActivity) {
	o.Activity.Set(&v)
}

// SetActivityNil sets the value for Activity to be an explicit nil
func (o *WidgetMember) SetActivityNil() {
	o.Activity.Set(nil)
}

// UnsetActivity ensures that no value is present for Activity, not even an explicit nil
func (o *WidgetMember) UnsetActivity() {
	o.Activity.Unset()
}

// GetDeaf returns the Deaf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetMember) GetDeaf() bool {
	if o == nil || IsNil(o.Deaf.Get()) {
		var ret bool
		return ret
	}
	return *o.Deaf.Get()
}

// GetDeafOk returns a tuple with the Deaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetMember) GetDeafOk() (*bool, bool) {
	if o == nil || IsNil(o.Deaf.Get()) {
		return nil, false
	}
	return o.Deaf.Get(), o.Deaf.IsSet()
}

// HasDeaf returns a boolean if a field has been set.
func (o *WidgetMember) HasDeaf() bool {
	if o != nil && o.Deaf.IsSet() {
		return true
	}

	return false
}

// SetDeaf gets a reference to the given NullableBool and assigns it to the Deaf field.
func (o *WidgetMember) SetDeaf(v bool) {
	o.Deaf.Set(&v)
}

// SetDeafNil sets the value for Deaf to be an explicit nil
func (o *WidgetMember) SetDeafNil() {
	o.Deaf.Set(nil)
}

// UnsetDeaf ensures that no value is present for Deaf, not even an explicit nil
func (o *WidgetMember) UnsetDeaf() {
	o.Deaf.Unset()
}

// GetMute returns the Mute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetMember) GetMute() bool {
	if o == nil || IsNil(o.Mute.Get()) {
		var ret bool
		return ret
	}
	return *o.Mute.Get()
}

// GetMuteOk returns a tuple with the Mute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetMember) GetMuteOk() (*bool, bool) {
	if o == nil || IsNil(o.Mute.Get()) {
		return nil, false
	}
	return o.Mute.Get(), o.Mute.IsSet()
}

// HasMute returns a boolean if a field has been set.
func (o *WidgetMember) HasMute() bool {
	if o != nil && o.Mute.IsSet() {
		return true
	}

	return false
}

// SetMute gets a reference to the given NullableBool and assigns it to the Mute field.
func (o *WidgetMember) SetMute(v bool) {
	o.Mute.Set(&v)
}

// SetMuteNil sets the value for Mute to be an explicit nil
func (o *WidgetMember) SetMuteNil() {
	o.Mute.Set(nil)
}

// UnsetMute ensures that no value is present for Mute, not even an explicit nil
func (o *WidgetMember) UnsetMute() {
	o.Mute.Unset()
}

// GetSelfDeaf returns the SelfDeaf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetMember) GetSelfDeaf() bool {
	if o == nil || IsNil(o.SelfDeaf.Get()) {
		var ret bool
		return ret
	}
	return *o.SelfDeaf.Get()
}

// GetSelfDeafOk returns a tuple with the SelfDeaf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetMember) GetSelfDeafOk() (*bool, bool) {
	if o == nil || IsNil(o.SelfDeaf.Get()) {
		return nil, false
	}
	return o.SelfDeaf.Get(), o.SelfDeaf.IsSet()
}

// HasSelfDeaf returns a boolean if a field has been set.
func (o *WidgetMember) HasSelfDeaf() bool {
	if o != nil && o.SelfDeaf.IsSet() {
		return true
	}

	return false
}

// SetSelfDeaf gets a reference to the given NullableBool and assigns it to the SelfDeaf field.
func (o *WidgetMember) SetSelfDeaf(v bool) {
	o.SelfDeaf.Set(&v)
}

// SetSelfDeafNil sets the value for SelfDeaf to be an explicit nil
func (o *WidgetMember) SetSelfDeafNil() {
	o.SelfDeaf.Set(nil)
}

// UnsetSelfDeaf ensures that no value is present for SelfDeaf, not even an explicit nil
func (o *WidgetMember) UnsetSelfDeaf() {
	o.SelfDeaf.Unset()
}

// GetSelfMute returns the SelfMute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetMember) GetSelfMute() bool {
	if o == nil || IsNil(o.SelfMute.Get()) {
		var ret bool
		return ret
	}
	return *o.SelfMute.Get()
}

// GetSelfMuteOk returns a tuple with the SelfMute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetMember) GetSelfMuteOk() (*bool, bool) {
	if o == nil || IsNil(o.SelfMute.Get()) {
		return nil, false
	}
	return o.SelfMute.Get(), o.SelfMute.IsSet()
}

// HasSelfMute returns a boolean if a field has been set.
func (o *WidgetMember) HasSelfMute() bool {
	if o != nil && o.SelfMute.IsSet() {
		return true
	}

	return false
}

// SetSelfMute gets a reference to the given NullableBool and assigns it to the SelfMute field.
func (o *WidgetMember) SetSelfMute(v bool) {
	o.SelfMute.Set(&v)
}

// SetSelfMuteNil sets the value for SelfMute to be an explicit nil
func (o *WidgetMember) SetSelfMuteNil() {
	o.SelfMute.Set(nil)
}

// UnsetSelfMute ensures that no value is present for SelfMute, not even an explicit nil
func (o *WidgetMember) UnsetSelfMute() {
	o.SelfMute.Unset()
}

// GetSuppress returns the Suppress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WidgetMember) GetSuppress() bool {
	if o == nil || IsNil(o.Suppress.Get()) {
		var ret bool
		return ret
	}
	return *o.Suppress.Get()
}

// GetSuppressOk returns a tuple with the Suppress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WidgetMember) GetSuppressOk() (*bool, bool) {
	if o == nil || IsNil(o.Suppress.Get()) {
		return nil, false
	}
	return o.Suppress.Get(), o.Suppress.IsSet()
}

// HasSuppress returns a boolean if a field has been set.
func (o *WidgetMember) HasSuppress() bool {
	if o != nil && o.Suppress.IsSet() {
		return true
	}

	return false
}

// SetSuppress gets a reference to the given NullableBool and assigns it to the Suppress field.
func (o *WidgetMember) SetSuppress(v bool) {
	o.Suppress.Set(&v)
}

// SetSuppressNil sets the value for Suppress to be an explicit nil
func (o *WidgetMember) SetSuppressNil() {
	o.Suppress.Set(nil)
}

// UnsetSuppress ensures that no value is present for Suppress, not even an explicit nil
func (o *WidgetMember) UnsetSuppress() {
	o.Suppress.Unset()
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *WidgetMember) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetMember) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *WidgetMember) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *WidgetMember) SetChannelId(v string) {
	o.ChannelId = &v
}


func (o WidgetMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WidgetMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["username"] = o.Username
	toSerialize["discriminator"] = o.Discriminator.Get()
	if o.Avatar != nil {
		toSerialize["avatar"] = o.Avatar
	}
	toSerialize["status"] = o.Status
	toSerialize["avatar_url"] = o.AvatarUrl
	if o.Activity.IsSet() {
		toSerialize["activity"] = o.Activity.Get()
	}
	if o.Deaf.IsSet() {
		toSerialize["deaf"] = o.Deaf.Get()
	}
	if o.Mute.IsSet() {
		toSerialize["mute"] = o.Mute.Get()
	}
	if o.SelfDeaf.IsSet() {
		toSerialize["self_deaf"] = o.SelfDeaf.Get()
	}
	if o.SelfMute.IsSet() {
		toSerialize["self_mute"] = o.SelfMute.Get()
	}
	if o.Suppress.IsSet() {
		toSerialize["suppress"] = o.Suppress.Get()
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	return toSerialize, nil
}

func (o *WidgetMember) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"username",
		"discriminator",
		"status",
		"avatar_url",
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

	varWidgetMember := _WidgetMember{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWidgetMember)

	if err != nil {
		return err
	}

	*o = WidgetMember(varWidgetMember)

	return err
}

type NullableWidgetMember struct {
	value *WidgetMember
	isSet bool
}

func (v NullableWidgetMember) Get() *WidgetMember {
	return v.value
}

func (v *NullableWidgetMember) Set(val *WidgetMember) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetMember) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetMember(val *WidgetMember) *NullableWidgetMember {
	return &NullableWidgetMember{value: val, isSet: true}
}

func (v NullableWidgetMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


