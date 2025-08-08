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

// checks if the EmojiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmojiResponse{}

// EmojiResponse struct for EmojiResponse
type EmojiResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	User NullableUserResponse `json:"user,omitempty"`
	Roles []string `json:"roles"`
	RequireColons bool `json:"require_colons"`
	Managed bool `json:"managed"`
	Animated bool `json:"animated"`
	Available bool `json:"available"`
}

type _EmojiResponse EmojiResponse

// NewEmojiResponse instantiates a new EmojiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmojiResponse(id string, name string, roles []string, requireColons bool, managed bool, animated bool, available bool) *EmojiResponse {
	this := EmojiResponse{}
	this.Id = id
	this.Name = name
	this.Roles = roles
	this.RequireColons = requireColons
	this.Managed = managed
	this.Animated = animated
	this.Available = available
	return &this
}

// NewEmojiResponseWithDefaults instantiates a new EmojiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmojiResponseWithDefaults() *EmojiResponse {
	this := EmojiResponse{}
	return &this
}

// GetId returns the Id field value
func (o *EmojiResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EmojiResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EmojiResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EmojiResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EmojiResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EmojiResponse) SetName(v string) {
	o.Name = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmojiResponse) GetUser() UserResponse {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmojiResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.User.Get()) {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *EmojiResponse) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserResponse and assigns it to the User field.
func (o *EmojiResponse) SetUser(v UserResponse) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *EmojiResponse) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *EmojiResponse) UnsetUser() {
	o.User.Unset()
}

// GetRoles returns the Roles field value
func (o *EmojiResponse) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *EmojiResponse) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *EmojiResponse) SetRoles(v []string) {
	o.Roles = v
}

// GetRequireColons returns the RequireColons field value
func (o *EmojiResponse) GetRequireColons() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireColons
}

// GetRequireColonsOk returns a tuple with the RequireColons field value
// and a boolean to check if the value has been set.
func (o *EmojiResponse) GetRequireColonsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireColons, true
}

// SetRequireColons sets field value
func (o *EmojiResponse) SetRequireColons(v bool) {
	o.RequireColons = v
}

// GetManaged returns the Managed field value
func (o *EmojiResponse) GetManaged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Managed
}

// GetManagedOk returns a tuple with the Managed field value
// and a boolean to check if the value has been set.
func (o *EmojiResponse) GetManagedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Managed, true
}

// SetManaged sets field value
func (o *EmojiResponse) SetManaged(v bool) {
	o.Managed = v
}

// GetAnimated returns the Animated field value
func (o *EmojiResponse) GetAnimated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Animated
}

// GetAnimatedOk returns a tuple with the Animated field value
// and a boolean to check if the value has been set.
func (o *EmojiResponse) GetAnimatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Animated, true
}

// SetAnimated sets field value
func (o *EmojiResponse) SetAnimated(v bool) {
	o.Animated = v
}

// GetAvailable returns the Available field value
func (o *EmojiResponse) GetAvailable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *EmojiResponse) GetAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *EmojiResponse) SetAvailable(v bool) {
	o.Available = v
}

func (o EmojiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmojiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	toSerialize["roles"] = o.Roles
	toSerialize["require_colons"] = o.RequireColons
	toSerialize["managed"] = o.Managed
	toSerialize["animated"] = o.Animated
	toSerialize["available"] = o.Available
	return toSerialize, nil
}

func (o *EmojiResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"roles",
		"require_colons",
		"managed",
		"animated",
		"available",
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

	varEmojiResponse := _EmojiResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmojiResponse)

	if err != nil {
		return err
	}

	*o = EmojiResponse(varEmojiResponse)

	return err
}

type NullableEmojiResponse struct {
	value *EmojiResponse
	isSet bool
}

func (v NullableEmojiResponse) Get() *EmojiResponse {
	return v.value
}

func (v *NullableEmojiResponse) Set(val *EmojiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmojiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmojiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmojiResponse(val *EmojiResponse) *NullableEmojiResponse {
	return &NullableEmojiResponse{value: val, isSet: true}
}

func (v NullableEmojiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmojiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


