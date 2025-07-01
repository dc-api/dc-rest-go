/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:27:32.119933921Z[Etc/UTC]
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

// checks if the ApplicationRoleConnectionsMetadataItemRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationRoleConnectionsMetadataItemRequest{}

// ApplicationRoleConnectionsMetadataItemRequest struct for ApplicationRoleConnectionsMetadataItemRequest
type ApplicationRoleConnectionsMetadataItemRequest struct {
	Type int32 `json:"type"`
	Key string `json:"key"`
	Name string `json:"name"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Description string `json:"description"`
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
}

type _ApplicationRoleConnectionsMetadataItemRequest ApplicationRoleConnectionsMetadataItemRequest

// NewApplicationRoleConnectionsMetadataItemRequest instantiates a new ApplicationRoleConnectionsMetadataItemRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationRoleConnectionsMetadataItemRequest(type_ int32, key string, name string, description string) *ApplicationRoleConnectionsMetadataItemRequest {
	this := ApplicationRoleConnectionsMetadataItemRequest{}
	this.Type = type_
	this.Key = key
	this.Name = name
	this.Description = description
	return &this
}

// NewApplicationRoleConnectionsMetadataItemRequestWithDefaults instantiates a new ApplicationRoleConnectionsMetadataItemRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationRoleConnectionsMetadataItemRequestWithDefaults() *ApplicationRoleConnectionsMetadataItemRequest {
	this := ApplicationRoleConnectionsMetadataItemRequest{}
	return &this
}

// GetType returns the Type field value
func (o *ApplicationRoleConnectionsMetadataItemRequest) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationRoleConnectionsMetadataItemRequest) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationRoleConnectionsMetadataItemRequest) SetType(v int32) {
	o.Type = v
}

// GetKey returns the Key field value
func (o *ApplicationRoleConnectionsMetadataItemRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ApplicationRoleConnectionsMetadataItemRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ApplicationRoleConnectionsMetadataItemRequest) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *ApplicationRoleConnectionsMetadataItemRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationRoleConnectionsMetadataItemRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationRoleConnectionsMetadataItemRequest) SetName(v string) {
	o.Name = v
}

// GetNameLocalizations returns the NameLocalizations field value if set, zero value otherwise.
func (o *ApplicationRoleConnectionsMetadataItemRequest) GetNameLocalizations() map[string]string {
	if o == nil || IsNil(o.NameLocalizations) {
		var ret map[string]string
		return ret
	}
	return o.NameLocalizations
}

// GetNameLocalizationsOk returns a tuple with the NameLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRoleConnectionsMetadataItemRequest) GetNameLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.NameLocalizations) {
		return map[string]string{}, false
	}
	return o.NameLocalizations, true
}

// HasNameLocalizations returns a boolean if a field has been set.
func (o *ApplicationRoleConnectionsMetadataItemRequest) HasNameLocalizations() bool {
	if o != nil && !IsNil(o.NameLocalizations) {
		return true
	}

	return false
}

// SetNameLocalizations gets a reference to the given map[string]string and assigns it to the NameLocalizations field.
func (o *ApplicationRoleConnectionsMetadataItemRequest) SetNameLocalizations(v map[string]string) {
	o.NameLocalizations = v
}


// GetDescription returns the Description field value
func (o *ApplicationRoleConnectionsMetadataItemRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApplicationRoleConnectionsMetadataItemRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ApplicationRoleConnectionsMetadataItemRequest) SetDescription(v string) {
	o.Description = v
}

// GetDescriptionLocalizations returns the DescriptionLocalizations field value if set, zero value otherwise.
func (o *ApplicationRoleConnectionsMetadataItemRequest) GetDescriptionLocalizations() map[string]string {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		var ret map[string]string
		return ret
	}
	return o.DescriptionLocalizations
}

// GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRoleConnectionsMetadataItemRequest) GetDescriptionLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		return map[string]string{}, false
	}
	return o.DescriptionLocalizations, true
}

// HasDescriptionLocalizations returns a boolean if a field has been set.
func (o *ApplicationRoleConnectionsMetadataItemRequest) HasDescriptionLocalizations() bool {
	if o != nil && !IsNil(o.DescriptionLocalizations) {
		return true
	}

	return false
}

// SetDescriptionLocalizations gets a reference to the given map[string]string and assigns it to the DescriptionLocalizations field.
func (o *ApplicationRoleConnectionsMetadataItemRequest) SetDescriptionLocalizations(v map[string]string) {
	o.DescriptionLocalizations = v
}


func (o ApplicationRoleConnectionsMetadataItemRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationRoleConnectionsMetadataItemRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	if !IsNil(o.NameLocalizations) {
		toSerialize["name_localizations"] = o.NameLocalizations
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.DescriptionLocalizations) {
		toSerialize["description_localizations"] = o.DescriptionLocalizations
	}
	return toSerialize, nil
}

func (o *ApplicationRoleConnectionsMetadataItemRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"key",
		"name",
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

	varApplicationRoleConnectionsMetadataItemRequest := _ApplicationRoleConnectionsMetadataItemRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationRoleConnectionsMetadataItemRequest)

	if err != nil {
		return err
	}

	*o = ApplicationRoleConnectionsMetadataItemRequest(varApplicationRoleConnectionsMetadataItemRequest)

	return err
}

type NullableApplicationRoleConnectionsMetadataItemRequest struct {
	value *ApplicationRoleConnectionsMetadataItemRequest
	isSet bool
}

func (v NullableApplicationRoleConnectionsMetadataItemRequest) Get() *ApplicationRoleConnectionsMetadataItemRequest {
	return v.value
}

func (v *NullableApplicationRoleConnectionsMetadataItemRequest) Set(val *ApplicationRoleConnectionsMetadataItemRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationRoleConnectionsMetadataItemRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationRoleConnectionsMetadataItemRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationRoleConnectionsMetadataItemRequest(val *ApplicationRoleConnectionsMetadataItemRequest) *NullableApplicationRoleConnectionsMetadataItemRequest {
	return &NullableApplicationRoleConnectionsMetadataItemRequest{value: val, isSet: true}
}

func (v NullableApplicationRoleConnectionsMetadataItemRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationRoleConnectionsMetadataItemRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


