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

// checks if the GuildChannelLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildChannelLocation{}

// GuildChannelLocation struct for GuildChannelLocation
type GuildChannelLocation struct {
	Id string `json:"id"`
	Kind string `json:"kind"`
	ChannelId string `json:"channel_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	GuildId string `json:"guild_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
}

type _GuildChannelLocation GuildChannelLocation

// NewGuildChannelLocation instantiates a new GuildChannelLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildChannelLocation(id string, kind string, channelId string, guildId string) *GuildChannelLocation {
	this := GuildChannelLocation{}
	this.Id = id
	this.Kind = kind
	this.ChannelId = channelId
	this.GuildId = guildId
	return &this
}

// NewGuildChannelLocationWithDefaults instantiates a new GuildChannelLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildChannelLocationWithDefaults() *GuildChannelLocation {
	this := GuildChannelLocation{}
	return &this
}

// GetId returns the Id field value
func (o *GuildChannelLocation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GuildChannelLocation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GuildChannelLocation) SetId(v string) {
	o.Id = v
}

// GetKind returns the Kind field value
func (o *GuildChannelLocation) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *GuildChannelLocation) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *GuildChannelLocation) SetKind(v string) {
	o.Kind = v
}

// GetChannelId returns the ChannelId field value
func (o *GuildChannelLocation) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *GuildChannelLocation) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *GuildChannelLocation) SetChannelId(v string) {
	o.ChannelId = v
}

// GetGuildId returns the GuildId field value
func (o *GuildChannelLocation) GetGuildId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value
// and a boolean to check if the value has been set.
func (o *GuildChannelLocation) GetGuildIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GuildId, true
}

// SetGuildId sets field value
func (o *GuildChannelLocation) SetGuildId(v string) {
	o.GuildId = v
}

func (o GuildChannelLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildChannelLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["kind"] = o.Kind
	toSerialize["channel_id"] = o.ChannelId
	toSerialize["guild_id"] = o.GuildId
	return toSerialize, nil
}

func (o *GuildChannelLocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"kind",
		"channel_id",
		"guild_id",
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

	varGuildChannelLocation := _GuildChannelLocation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildChannelLocation)

	if err != nil {
		return err
	}

	*o = GuildChannelLocation(varGuildChannelLocation)

	return err
}

type NullableGuildChannelLocation struct {
	value *GuildChannelLocation
	isSet bool
}

func (v NullableGuildChannelLocation) Get() *GuildChannelLocation {
	return v.value
}

func (v *NullableGuildChannelLocation) Set(val *GuildChannelLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildChannelLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildChannelLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildChannelLocation(val *GuildChannelLocation) *NullableGuildChannelLocation {
	return &NullableGuildChannelLocation{value: val, isSet: true}
}

func (v NullableGuildChannelLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildChannelLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


