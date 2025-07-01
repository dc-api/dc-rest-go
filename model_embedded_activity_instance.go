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

// checks if the EmbeddedActivityInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbeddedActivityInstance{}

// EmbeddedActivityInstance struct for EmbeddedActivityInstance
type EmbeddedActivityInstance struct {
	ApplicationId string `json:"application_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	InstanceId string `json:"instance_id"`
	LaunchId string `json:"launch_id"`
	Location NullableEmbeddedActivityInstanceLocation `json:"location,omitempty"`
	Users []string `json:"users"`
}

type _EmbeddedActivityInstance EmbeddedActivityInstance

// NewEmbeddedActivityInstance instantiates a new EmbeddedActivityInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddedActivityInstance(applicationId string, instanceId string, launchId string, users []string) *EmbeddedActivityInstance {
	this := EmbeddedActivityInstance{}
	this.ApplicationId = applicationId
	this.InstanceId = instanceId
	this.LaunchId = launchId
	this.Users = users
	return &this
}

// NewEmbeddedActivityInstanceWithDefaults instantiates a new EmbeddedActivityInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedActivityInstanceWithDefaults() *EmbeddedActivityInstance {
	this := EmbeddedActivityInstance{}
	return &this
}

// GetApplicationId returns the ApplicationId field value
func (o *EmbeddedActivityInstance) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *EmbeddedActivityInstance) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *EmbeddedActivityInstance) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetInstanceId returns the InstanceId field value
func (o *EmbeddedActivityInstance) GetInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceId
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
func (o *EmbeddedActivityInstance) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceId, true
}

// SetInstanceId sets field value
func (o *EmbeddedActivityInstance) SetInstanceId(v string) {
	o.InstanceId = v
}

// GetLaunchId returns the LaunchId field value
func (o *EmbeddedActivityInstance) GetLaunchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LaunchId
}

// GetLaunchIdOk returns a tuple with the LaunchId field value
// and a boolean to check if the value has been set.
func (o *EmbeddedActivityInstance) GetLaunchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LaunchId, true
}

// SetLaunchId sets field value
func (o *EmbeddedActivityInstance) SetLaunchId(v string) {
	o.LaunchId = v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmbeddedActivityInstance) GetLocation() EmbeddedActivityInstanceLocation {
	if o == nil || IsNil(o.Location.Get()) {
		var ret EmbeddedActivityInstanceLocation
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmbeddedActivityInstance) GetLocationOk() (*EmbeddedActivityInstanceLocation, bool) {
	if o == nil || IsNil(o.Location.Get()) {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *EmbeddedActivityInstance) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableEmbeddedActivityInstanceLocation and assigns it to the Location field.
func (o *EmbeddedActivityInstance) SetLocation(v EmbeddedActivityInstanceLocation) {
	o.Location.Set(&v)
}

// SetLocationNil sets the value for Location to be an explicit nil
func (o *EmbeddedActivityInstance) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *EmbeddedActivityInstance) UnsetLocation() {
	o.Location.Unset()
}

// GetUsers returns the Users field value
func (o *EmbeddedActivityInstance) GetUsers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *EmbeddedActivityInstance) GetUsersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *EmbeddedActivityInstance) SetUsers(v []string) {
	o.Users = v
}

func (o EmbeddedActivityInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbeddedActivityInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["application_id"] = o.ApplicationId
	toSerialize["instance_id"] = o.InstanceId
	toSerialize["launch_id"] = o.LaunchId
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *EmbeddedActivityInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"application_id",
		"instance_id",
		"launch_id",
		"users",
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

	varEmbeddedActivityInstance := _EmbeddedActivityInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmbeddedActivityInstance)

	if err != nil {
		return err
	}

	*o = EmbeddedActivityInstance(varEmbeddedActivityInstance)

	return err
}

type NullableEmbeddedActivityInstance struct {
	value *EmbeddedActivityInstance
	isSet bool
}

func (v NullableEmbeddedActivityInstance) Get() *EmbeddedActivityInstance {
	return v.value
}

func (v *NullableEmbeddedActivityInstance) Set(val *EmbeddedActivityInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedActivityInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedActivityInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedActivityInstance(val *EmbeddedActivityInstance) *NullableEmbeddedActivityInstance {
	return &NullableEmbeddedActivityInstance{value: val, isSet: true}
}

func (v NullableEmbeddedActivityInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedActivityInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


