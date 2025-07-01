/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:17:22.835370245Z[Etc/UTC]
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

// checks if the ChannelFollowerWebhookResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelFollowerWebhookResponse{}

// ChannelFollowerWebhookResponse struct for ChannelFollowerWebhookResponse
type ChannelFollowerWebhookResponse struct {
	ApplicationId *string `json:"application_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Avatar NullableString `json:"avatar,omitempty"`
	ChannelId *string `json:"channel_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	GuildId *string `json:"guild_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Type int32 `json:"type"`
	User NullableUserResponse `json:"user,omitempty"`
	SourceGuild NullableWebhookSourceGuildResponse `json:"source_guild,omitempty"`
	SourceChannel NullableWebhookSourceChannelResponse `json:"source_channel,omitempty"`
}

type _ChannelFollowerWebhookResponse ChannelFollowerWebhookResponse

// NewChannelFollowerWebhookResponse instantiates a new ChannelFollowerWebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelFollowerWebhookResponse(id string, name string, type_ int32) *ChannelFollowerWebhookResponse {
	this := ChannelFollowerWebhookResponse{}
	this.Id = id
	this.Name = name
	this.Type = type_
	return &this
}

// NewChannelFollowerWebhookResponseWithDefaults instantiates a new ChannelFollowerWebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelFollowerWebhookResponseWithDefaults() *ChannelFollowerWebhookResponse {
	this := ChannelFollowerWebhookResponse{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ChannelFollowerWebhookResponse) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelFollowerWebhookResponse) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ChannelFollowerWebhookResponse) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ChannelFollowerWebhookResponse) SetApplicationId(v string) {
	o.ApplicationId = &v
}


// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelFollowerWebhookResponse) GetAvatar() string {
	if o == nil || IsNil(o.Avatar.Get()) {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelFollowerWebhookResponse) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar.Get()) {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *ChannelFollowerWebhookResponse) HasAvatar() bool {
	if o != nil && o.Avatar.IsSet() {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given NullableString and assigns it to the Avatar field.
func (o *ChannelFollowerWebhookResponse) SetAvatar(v string) {
	o.Avatar.Set(&v)
}

// SetAvatarNil sets the value for Avatar to be an explicit nil
func (o *ChannelFollowerWebhookResponse) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
func (o *ChannelFollowerWebhookResponse) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *ChannelFollowerWebhookResponse) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelFollowerWebhookResponse) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *ChannelFollowerWebhookResponse) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *ChannelFollowerWebhookResponse) SetChannelId(v string) {
	o.ChannelId = &v
}


// GetGuildId returns the GuildId field value if set, zero value otherwise.
func (o *ChannelFollowerWebhookResponse) GetGuildId() string {
	if o == nil || IsNil(o.GuildId) {
		var ret string
		return ret
	}
	return *o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelFollowerWebhookResponse) GetGuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.GuildId) {
		return nil, false
	}
	return o.GuildId, true
}

// HasGuildId returns a boolean if a field has been set.
func (o *ChannelFollowerWebhookResponse) HasGuildId() bool {
	if o != nil && !IsNil(o.GuildId) {
		return true
	}

	return false
}

// SetGuildId gets a reference to the given string and assigns it to the GuildId field.
func (o *ChannelFollowerWebhookResponse) SetGuildId(v string) {
	o.GuildId = &v
}


// GetId returns the Id field value
func (o *ChannelFollowerWebhookResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChannelFollowerWebhookResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChannelFollowerWebhookResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ChannelFollowerWebhookResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChannelFollowerWebhookResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ChannelFollowerWebhookResponse) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ChannelFollowerWebhookResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChannelFollowerWebhookResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ChannelFollowerWebhookResponse) SetType(v int32) {
	o.Type = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelFollowerWebhookResponse) GetUser() UserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelFollowerWebhookResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.User.Get()) {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *ChannelFollowerWebhookResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserResponse and assigns it to the User field.
func (o *ChannelFollowerWebhookResponse) SetUser(v UserResponse) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *ChannelFollowerWebhookResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *ChannelFollowerWebhookResponse) UnsetUser() {
	o.User.Unset()
}

// GetSourceGuild returns the SourceGuild field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelFollowerWebhookResponse) GetSourceGuild() WebhookSourceGuildResponse {
	if o == nil || IsNil(o.SourceGuild.Get()) {
		var ret WebhookSourceGuildResponse
		return ret
	}
	return *o.SourceGuild.Get()
}

// GetSourceGuildOk returns a tuple with the SourceGuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelFollowerWebhookResponse) GetSourceGuildOk() (*WebhookSourceGuildResponse, bool) {
	if o == nil || IsNil(o.SourceGuild.Get()) {
		return nil, false
	}
	return o.SourceGuild.Get(), o.SourceGuild.IsSet()
}

// HasSourceGuild returns a boolean if a field has been set.
func (o *ChannelFollowerWebhookResponse) HasSourceGuild() bool {
	if o != nil && o.SourceGuild.IsSet() {
		return true
	}

	return false
}

// SetSourceGuild gets a reference to the given NullableWebhookSourceGuildResponse and assigns it to the SourceGuild field.
func (o *ChannelFollowerWebhookResponse) SetSourceGuild(v WebhookSourceGuildResponse) {
	o.SourceGuild.Set(&v)
}

// SetSourceGuildNil sets the value for SourceGuild to be an explicit nil
func (o *ChannelFollowerWebhookResponse) SetSourceGuildNil() {
	o.SourceGuild.Set(nil)
}

// UnsetSourceGuild ensures that no value is present for SourceGuild, not even an explicit nil
func (o *ChannelFollowerWebhookResponse) UnsetSourceGuild() {
	o.SourceGuild.Unset()
}

// GetSourceChannel returns the SourceChannel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelFollowerWebhookResponse) GetSourceChannel() WebhookSourceChannelResponse {
	if o == nil || IsNil(o.SourceChannel.Get()) {
		var ret WebhookSourceChannelResponse
		return ret
	}
	return *o.SourceChannel.Get()
}

// GetSourceChannelOk returns a tuple with the SourceChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelFollowerWebhookResponse) GetSourceChannelOk() (*WebhookSourceChannelResponse, bool) {
	if o == nil || IsNil(o.SourceChannel.Get()) {
		return nil, false
	}
	return o.SourceChannel.Get(), o.SourceChannel.IsSet()
}

// HasSourceChannel returns a boolean if a field has been set.
func (o *ChannelFollowerWebhookResponse) HasSourceChannel() bool {
	if o != nil && o.SourceChannel.IsSet() {
		return true
	}

	return false
}

// SetSourceChannel gets a reference to the given NullableWebhookSourceChannelResponse and assigns it to the SourceChannel field.
func (o *ChannelFollowerWebhookResponse) SetSourceChannel(v WebhookSourceChannelResponse) {
	o.SourceChannel.Set(&v)
}

// SetSourceChannelNil sets the value for SourceChannel to be an explicit nil
func (o *ChannelFollowerWebhookResponse) SetSourceChannelNil() {
	o.SourceChannel.Set(nil)
}

// UnsetSourceChannel ensures that no value is present for SourceChannel, not even an explicit nil
func (o *ChannelFollowerWebhookResponse) UnsetSourceChannel() {
	o.SourceChannel.Unset()
}

func (o ChannelFollowerWebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelFollowerWebhookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["application_id"] = o.ApplicationId
	}
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.GuildId) {
		toSerialize["guild_id"] = o.GuildId
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.SourceGuild.IsSet() {
		toSerialize["source_guild"] = o.SourceGuild.Get()
	}
	if o.SourceChannel.IsSet() {
		toSerialize["source_channel"] = o.SourceChannel.Get()
	}
	return toSerialize, nil
}

func (o *ChannelFollowerWebhookResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
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

	varChannelFollowerWebhookResponse := _ChannelFollowerWebhookResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChannelFollowerWebhookResponse)

	if err != nil {
		return err
	}

	*o = ChannelFollowerWebhookResponse(varChannelFollowerWebhookResponse)

	return err
}

type NullableChannelFollowerWebhookResponse struct {
	value *ChannelFollowerWebhookResponse
	isSet bool
}

func (v NullableChannelFollowerWebhookResponse) Get() *ChannelFollowerWebhookResponse {
	return v.value
}

func (v *NullableChannelFollowerWebhookResponse) Set(val *ChannelFollowerWebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelFollowerWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelFollowerWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelFollowerWebhookResponse(val *ChannelFollowerWebhookResponse) *NullableChannelFollowerWebhookResponse {
	return &NullableChannelFollowerWebhookResponse{value: val, isSet: true}
}

func (v NullableChannelFollowerWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelFollowerWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


