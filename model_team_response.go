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

// checks if the TeamResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TeamResponse{}

// TeamResponse struct for TeamResponse
type TeamResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Icon NullableString `json:"icon,omitempty"`
	Name string `json:"name"`
	OwnerUserId string `json:"owner_user_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Members []TeamMemberResponse `json:"members"`
}

type _TeamResponse TeamResponse

// NewTeamResponse instantiates a new TeamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamResponse(id string, name string, ownerUserId string, members []TeamMemberResponse) *TeamResponse {
	this := TeamResponse{}
	this.Id = id
	this.Name = name
	this.OwnerUserId = ownerUserId
	this.Members = members
	return &this
}

// NewTeamResponseWithDefaults instantiates a new TeamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamResponseWithDefaults() *TeamResponse {
	this := TeamResponse{}
	return &this
}

// GetId returns the Id field value
func (o *TeamResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TeamResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TeamResponse) SetId(v string) {
	o.Id = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TeamResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TeamResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *TeamResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *TeamResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *TeamResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *TeamResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetName returns the Name field value
func (o *TeamResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TeamResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TeamResponse) SetName(v string) {
	o.Name = v
}

// GetOwnerUserId returns the OwnerUserId field value
func (o *TeamResponse) GetOwnerUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerUserId
}

// GetOwnerUserIdOk returns a tuple with the OwnerUserId field value
// and a boolean to check if the value has been set.
func (o *TeamResponse) GetOwnerUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerUserId, true
}

// SetOwnerUserId sets field value
func (o *TeamResponse) SetOwnerUserId(v string) {
	o.OwnerUserId = v
}

// GetMembers returns the Members field value
func (o *TeamResponse) GetMembers() []TeamMemberResponse {
	if o == nil {
		var ret []TeamMemberResponse
		return ret
	}

	return o.Members
}

// GetMembersOk returns a tuple with the Members field value
// and a boolean to check if the value has been set.
func (o *TeamResponse) GetMembersOk() ([]TeamMemberResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Members, true
}

// SetMembers sets field value
func (o *TeamResponse) SetMembers(v []TeamMemberResponse) {
	o.Members = v
}

func (o TeamResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["owner_user_id"] = o.OwnerUserId
	toSerialize["members"] = o.Members
	return toSerialize, nil
}

func (o *TeamResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"owner_user_id",
		"members",
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

	varTeamResponse := _TeamResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeamResponse)

	if err != nil {
		return err
	}

	*o = TeamResponse(varTeamResponse)

	return err
}

type NullableTeamResponse struct {
	value *TeamResponse
	isSet bool
}

func (v NullableTeamResponse) Get() *TeamResponse {
	return v.value
}

func (v *NullableTeamResponse) Set(val *TeamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamResponse(val *TeamResponse) *NullableTeamResponse {
	return &NullableTeamResponse{value: val, isSet: true}
}

func (v NullableTeamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


