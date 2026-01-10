/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-10T02:53:50.800755799Z[Etc/UTC]
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
	"bytes"
	"fmt"
)

// checks if the ConnectedAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectedAccountResponse{}

// ConnectedAccountResponse struct for ConnectedAccountResponse
type ConnectedAccountResponse struct {
	Id string `json:"id"`
	Name *string `json:"name,omitempty"`
	Type ConnectedAccountProviders `json:"type"`
	FriendSync bool `json:"friend_sync"`
	Integrations []ConnectedAccountIntegrationResponse `json:"integrations,omitempty"`
	ShowActivity bool `json:"show_activity"`
	TwoWayLink bool `json:"two_way_link"`
	Verified bool `json:"verified"`
	Visibility ConnectedAccountVisibility `json:"visibility"`
	Revoked *bool `json:"revoked,omitempty"`
}

type _ConnectedAccountResponse ConnectedAccountResponse

// NewConnectedAccountResponse instantiates a new ConnectedAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedAccountResponse(id string, type_ ConnectedAccountProviders, friendSync bool, showActivity bool, twoWayLink bool, verified bool, visibility ConnectedAccountVisibility) *ConnectedAccountResponse {
	this := ConnectedAccountResponse{}
	this.Id = id
	this.Type = type_
	this.FriendSync = friendSync
	this.ShowActivity = showActivity
	this.TwoWayLink = twoWayLink
	this.Verified = verified
	this.Visibility = visibility
	return &this
}

// NewConnectedAccountResponseWithDefaults instantiates a new ConnectedAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedAccountResponseWithDefaults() *ConnectedAccountResponse {
	this := ConnectedAccountResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ConnectedAccountResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConnectedAccountResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConnectedAccountResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectedAccountResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAccountResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectedAccountResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectedAccountResponse) SetName(v string) {
	o.Name = &v
}


// GetType returns the Type field value
func (o *ConnectedAccountResponse) GetType() ConnectedAccountProviders {
	if o == nil {
		var ret ConnectedAccountProviders
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConnectedAccountResponse) GetTypeOk() (*ConnectedAccountProviders, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConnectedAccountResponse) SetType(v ConnectedAccountProviders) {
	o.Type = v
}

// GetFriendSync returns the FriendSync field value
func (o *ConnectedAccountResponse) GetFriendSync() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FriendSync
}

// GetFriendSyncOk returns a tuple with the FriendSync field value
// and a boolean to check if the value has been set.
func (o *ConnectedAccountResponse) GetFriendSyncOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FriendSync, true
}

// SetFriendSync sets field value
func (o *ConnectedAccountResponse) SetFriendSync(v bool) {
	o.FriendSync = v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *ConnectedAccountResponse) GetIntegrations() []ConnectedAccountIntegrationResponse {
	if o == nil || IsNil(o.Integrations) {
		var ret []ConnectedAccountIntegrationResponse
		return ret
	}
	return o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAccountResponse) GetIntegrationsOk() ([]ConnectedAccountIntegrationResponse, bool) {
	if o == nil || IsNil(o.Integrations) {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *ConnectedAccountResponse) HasIntegrations() bool {
	if o != nil && !IsNil(o.Integrations) {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given []ConnectedAccountIntegrationResponse and assigns it to the Integrations field.
func (o *ConnectedAccountResponse) SetIntegrations(v []ConnectedAccountIntegrationResponse) {
	o.Integrations = v
}


// GetShowActivity returns the ShowActivity field value
func (o *ConnectedAccountResponse) GetShowActivity() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShowActivity
}

// GetShowActivityOk returns a tuple with the ShowActivity field value
// and a boolean to check if the value has been set.
func (o *ConnectedAccountResponse) GetShowActivityOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShowActivity, true
}

// SetShowActivity sets field value
func (o *ConnectedAccountResponse) SetShowActivity(v bool) {
	o.ShowActivity = v
}

// GetTwoWayLink returns the TwoWayLink field value
func (o *ConnectedAccountResponse) GetTwoWayLink() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TwoWayLink
}

// GetTwoWayLinkOk returns a tuple with the TwoWayLink field value
// and a boolean to check if the value has been set.
func (o *ConnectedAccountResponse) GetTwoWayLinkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TwoWayLink, true
}

// SetTwoWayLink sets field value
func (o *ConnectedAccountResponse) SetTwoWayLink(v bool) {
	o.TwoWayLink = v
}

// GetVerified returns the Verified field value
func (o *ConnectedAccountResponse) GetVerified() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value
// and a boolean to check if the value has been set.
func (o *ConnectedAccountResponse) GetVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verified, true
}

// SetVerified sets field value
func (o *ConnectedAccountResponse) SetVerified(v bool) {
	o.Verified = v
}

// GetVisibility returns the Visibility field value
func (o *ConnectedAccountResponse) GetVisibility() ConnectedAccountVisibility {
	if o == nil {
		var ret ConnectedAccountVisibility
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *ConnectedAccountResponse) GetVisibilityOk() (*ConnectedAccountVisibility, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *ConnectedAccountResponse) SetVisibility(v ConnectedAccountVisibility) {
	o.Visibility = v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *ConnectedAccountResponse) GetRevoked() bool {
	if o == nil || IsNil(o.Revoked) {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedAccountResponse) GetRevokedOk() (*bool, bool) {
	if o == nil || IsNil(o.Revoked) {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *ConnectedAccountResponse) HasRevoked() bool {
	if o != nil && !IsNil(o.Revoked) {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *ConnectedAccountResponse) SetRevoked(v bool) {
	o.Revoked = &v
}


func (o ConnectedAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectedAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["type"] = o.Type
	toSerialize["friend_sync"] = o.FriendSync
	if !IsNil(o.Integrations) {
		toSerialize["integrations"] = o.Integrations
	}
	toSerialize["show_activity"] = o.ShowActivity
	toSerialize["two_way_link"] = o.TwoWayLink
	toSerialize["verified"] = o.Verified
	toSerialize["visibility"] = o.Visibility
	if !IsNil(o.Revoked) {
		toSerialize["revoked"] = o.Revoked
	}
	return toSerialize, nil
}

func (o *ConnectedAccountResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"friend_sync",
		"show_activity",
		"two_way_link",
		"verified",
		"visibility",
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

	varConnectedAccountResponse := _ConnectedAccountResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnectedAccountResponse)

	if err != nil {
		return err
	}

	*o = ConnectedAccountResponse(varConnectedAccountResponse)

	return err
}

type NullableConnectedAccountResponse struct {
	value *ConnectedAccountResponse
	isSet bool
}

func (v NullableConnectedAccountResponse) Get() *ConnectedAccountResponse {
	return v.value
}

func (v *NullableConnectedAccountResponse) Set(val *ConnectedAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedAccountResponse(val *ConnectedAccountResponse) *NullableConnectedAccountResponse {
	return &NullableConnectedAccountResponse{value: val, isSet: true}
}

func (v NullableConnectedAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

