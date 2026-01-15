/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-15T01:01:02.545594866Z[Etc/UTC]
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
	"time"
	"bytes"
	"fmt"
)

// checks if the ExternalConnectionIntegrationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalConnectionIntegrationResponse{}

// ExternalConnectionIntegrationResponse struct for ExternalConnectionIntegrationResponse
type ExternalConnectionIntegrationResponse struct {
	Type IntegrationTypes `json:"type"`
	Name *string `json:"name,omitempty"`
	Account AccountResponse `json:"account"`
	Enabled bool `json:"enabled"`
	Id string `json:"id"`
	User UserResponse `json:"user"`
	Revoked *bool `json:"revoked,omitempty"`
	ExpireBehavior *IntegrationExpireBehaviorTypes `json:"expire_behavior,omitempty"`
	ExpireGracePeriod *IntegrationExpireGracePeriodTypes `json:"expire_grace_period,omitempty"`
	SubscriberCount *int32 `json:"subscriber_count,omitempty"`
	SyncedAt *time.Time `json:"synced_at,omitempty"`
	RoleId *GetEntitlementsSkuIdsParameterOneOfInner `json:"role_id,omitempty"`
	Syncing *bool `json:"syncing,omitempty"`
	EnableEmoticons *bool `json:"enable_emoticons,omitempty"`
}

type _ExternalConnectionIntegrationResponse ExternalConnectionIntegrationResponse

// NewExternalConnectionIntegrationResponse instantiates a new ExternalConnectionIntegrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalConnectionIntegrationResponse(type_ IntegrationTypes, account AccountResponse, enabled bool, id string, user UserResponse) *ExternalConnectionIntegrationResponse {
	this := ExternalConnectionIntegrationResponse{}
	this.Type = type_
	this.Account = account
	this.Enabled = enabled
	this.Id = id
	this.User = user
	return &this
}

// NewExternalConnectionIntegrationResponseWithDefaults instantiates a new ExternalConnectionIntegrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalConnectionIntegrationResponseWithDefaults() *ExternalConnectionIntegrationResponse {
	this := ExternalConnectionIntegrationResponse{}
	return &this
}

// GetType returns the Type field value
func (o *ExternalConnectionIntegrationResponse) GetType() IntegrationTypes {
	if o == nil {
		var ret IntegrationTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetTypeOk() (*IntegrationTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExternalConnectionIntegrationResponse) SetType(v IntegrationTypes) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExternalConnectionIntegrationResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExternalConnectionIntegrationResponse) SetName(v string) {
	o.Name = &v
}


// GetAccount returns the Account field value
func (o *ExternalConnectionIntegrationResponse) GetAccount() AccountResponse {
	if o == nil {
		var ret AccountResponse
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetAccountOk() (*AccountResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *ExternalConnectionIntegrationResponse) SetAccount(v AccountResponse) {
	o.Account = v
}

// GetEnabled returns the Enabled field value
func (o *ExternalConnectionIntegrationResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ExternalConnectionIntegrationResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value
func (o *ExternalConnectionIntegrationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalConnectionIntegrationResponse) SetId(v string) {
	o.Id = v
}

// GetUser returns the User field value
func (o *ExternalConnectionIntegrationResponse) GetUser() UserResponse {
	if o == nil {
		var ret UserResponse
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetUserOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ExternalConnectionIntegrationResponse) SetUser(v UserResponse) {
	o.User = v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *ExternalConnectionIntegrationResponse) GetRevoked() bool {
	if o == nil || IsNil(o.Revoked) {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetRevokedOk() (*bool, bool) {
	if o == nil || IsNil(o.Revoked) {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasRevoked() bool {
	if o != nil && !IsNil(o.Revoked) {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *ExternalConnectionIntegrationResponse) SetRevoked(v bool) {
	o.Revoked = &v
}


// GetExpireBehavior returns the ExpireBehavior field value if set, zero value otherwise.
func (o *ExternalConnectionIntegrationResponse) GetExpireBehavior() IntegrationExpireBehaviorTypes {
	if o == nil || IsNil(o.ExpireBehavior) {
		var ret IntegrationExpireBehaviorTypes
		return ret
	}
	return *o.ExpireBehavior
}

// GetExpireBehaviorOk returns a tuple with the ExpireBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetExpireBehaviorOk() (*IntegrationExpireBehaviorTypes, bool) {
	if o == nil || IsNil(o.ExpireBehavior) {
		return nil, false
	}
	return o.ExpireBehavior, true
}

// HasExpireBehavior returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasExpireBehavior() bool {
	if o != nil && !IsNil(o.ExpireBehavior) {
		return true
	}

	return false
}

// SetExpireBehavior gets a reference to the given IntegrationExpireBehaviorTypes and assigns it to the ExpireBehavior field.
func (o *ExternalConnectionIntegrationResponse) SetExpireBehavior(v IntegrationExpireBehaviorTypes) {
	o.ExpireBehavior = &v
}


// GetExpireGracePeriod returns the ExpireGracePeriod field value if set, zero value otherwise.
func (o *ExternalConnectionIntegrationResponse) GetExpireGracePeriod() IntegrationExpireGracePeriodTypes {
	if o == nil || IsNil(o.ExpireGracePeriod) {
		var ret IntegrationExpireGracePeriodTypes
		return ret
	}
	return *o.ExpireGracePeriod
}

// GetExpireGracePeriodOk returns a tuple with the ExpireGracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetExpireGracePeriodOk() (*IntegrationExpireGracePeriodTypes, bool) {
	if o == nil || IsNil(o.ExpireGracePeriod) {
		return nil, false
	}
	return o.ExpireGracePeriod, true
}

// HasExpireGracePeriod returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasExpireGracePeriod() bool {
	if o != nil && !IsNil(o.ExpireGracePeriod) {
		return true
	}

	return false
}

// SetExpireGracePeriod gets a reference to the given IntegrationExpireGracePeriodTypes and assigns it to the ExpireGracePeriod field.
func (o *ExternalConnectionIntegrationResponse) SetExpireGracePeriod(v IntegrationExpireGracePeriodTypes) {
	o.ExpireGracePeriod = &v
}


// GetSubscriberCount returns the SubscriberCount field value if set, zero value otherwise.
func (o *ExternalConnectionIntegrationResponse) GetSubscriberCount() int32 {
	if o == nil || IsNil(o.SubscriberCount) {
		var ret int32
		return ret
	}
	return *o.SubscriberCount
}

// GetSubscriberCountOk returns a tuple with the SubscriberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetSubscriberCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SubscriberCount) {
		return nil, false
	}
	return o.SubscriberCount, true
}

// HasSubscriberCount returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasSubscriberCount() bool {
	if o != nil && !IsNil(o.SubscriberCount) {
		return true
	}

	return false
}

// SetSubscriberCount gets a reference to the given int32 and assigns it to the SubscriberCount field.
func (o *ExternalConnectionIntegrationResponse) SetSubscriberCount(v int32) {
	o.SubscriberCount = &v
}


// GetSyncedAt returns the SyncedAt field value if set, zero value otherwise.
func (o *ExternalConnectionIntegrationResponse) GetSyncedAt() time.Time {
	if o == nil || IsNil(o.SyncedAt) {
		var ret time.Time
		return ret
	}
	return *o.SyncedAt
}

// GetSyncedAtOk returns a tuple with the SyncedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetSyncedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SyncedAt) {
		return nil, false
	}
	return o.SyncedAt, true
}

// HasSyncedAt returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasSyncedAt() bool {
	if o != nil && !IsNil(o.SyncedAt) {
		return true
	}

	return false
}

// SetSyncedAt gets a reference to the given time.Time and assigns it to the SyncedAt field.
func (o *ExternalConnectionIntegrationResponse) SetSyncedAt(v time.Time) {
	o.SyncedAt = &v
}


// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *ExternalConnectionIntegrationResponse) GetRoleId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.RoleId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetRoleIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the RoleId field.
func (o *ExternalConnectionIntegrationResponse) SetRoleId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.RoleId = &v
}


// GetSyncing returns the Syncing field value if set, zero value otherwise.
func (o *ExternalConnectionIntegrationResponse) GetSyncing() bool {
	if o == nil || IsNil(o.Syncing) {
		var ret bool
		return ret
	}
	return *o.Syncing
}

// GetSyncingOk returns a tuple with the Syncing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetSyncingOk() (*bool, bool) {
	if o == nil || IsNil(o.Syncing) {
		return nil, false
	}
	return o.Syncing, true
}

// HasSyncing returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasSyncing() bool {
	if o != nil && !IsNil(o.Syncing) {
		return true
	}

	return false
}

// SetSyncing gets a reference to the given bool and assigns it to the Syncing field.
func (o *ExternalConnectionIntegrationResponse) SetSyncing(v bool) {
	o.Syncing = &v
}


// GetEnableEmoticons returns the EnableEmoticons field value if set, zero value otherwise.
func (o *ExternalConnectionIntegrationResponse) GetEnableEmoticons() bool {
	if o == nil || IsNil(o.EnableEmoticons) {
		var ret bool
		return ret
	}
	return *o.EnableEmoticons
}

// GetEnableEmoticonsOk returns a tuple with the EnableEmoticons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalConnectionIntegrationResponse) GetEnableEmoticonsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableEmoticons) {
		return nil, false
	}
	return o.EnableEmoticons, true
}

// HasEnableEmoticons returns a boolean if a field has been set.
func (o *ExternalConnectionIntegrationResponse) HasEnableEmoticons() bool {
	if o != nil && !IsNil(o.EnableEmoticons) {
		return true
	}

	return false
}

// SetEnableEmoticons gets a reference to the given bool and assigns it to the EnableEmoticons field.
func (o *ExternalConnectionIntegrationResponse) SetEnableEmoticons(v bool) {
	o.EnableEmoticons = &v
}


func (o ExternalConnectionIntegrationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalConnectionIntegrationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["account"] = o.Account
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	toSerialize["user"] = o.User
	if !IsNil(o.Revoked) {
		toSerialize["revoked"] = o.Revoked
	}
	if !IsNil(o.ExpireBehavior) {
		toSerialize["expire_behavior"] = o.ExpireBehavior
	}
	if !IsNil(o.ExpireGracePeriod) {
		toSerialize["expire_grace_period"] = o.ExpireGracePeriod
	}
	if !IsNil(o.SubscriberCount) {
		toSerialize["subscriber_count"] = o.SubscriberCount
	}
	if !IsNil(o.SyncedAt) {
		toSerialize["synced_at"] = o.SyncedAt
	}
	if !IsNil(o.RoleId) {
		toSerialize["role_id"] = o.RoleId
	}
	if !IsNil(o.Syncing) {
		toSerialize["syncing"] = o.Syncing
	}
	if !IsNil(o.EnableEmoticons) {
		toSerialize["enable_emoticons"] = o.EnableEmoticons
	}
	return toSerialize, nil
}

func (o *ExternalConnectionIntegrationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"account",
		"enabled",
		"id",
		"user",
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

	varExternalConnectionIntegrationResponse := _ExternalConnectionIntegrationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalConnectionIntegrationResponse)

	if err != nil {
		return err
	}

	*o = ExternalConnectionIntegrationResponse(varExternalConnectionIntegrationResponse)

	return err
}

type NullableExternalConnectionIntegrationResponse struct {
	value *ExternalConnectionIntegrationResponse
	isSet bool
}

func (v NullableExternalConnectionIntegrationResponse) Get() *ExternalConnectionIntegrationResponse {
	return v.value
}

func (v *NullableExternalConnectionIntegrationResponse) Set(val *ExternalConnectionIntegrationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalConnectionIntegrationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalConnectionIntegrationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalConnectionIntegrationResponse(val *ExternalConnectionIntegrationResponse) *NullableExternalConnectionIntegrationResponse {
	return &NullableExternalConnectionIntegrationResponse{value: val, isSet: true}
}

func (v NullableExternalConnectionIntegrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalConnectionIntegrationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

