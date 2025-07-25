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
)

// checks if the ApplicationFormPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationFormPartial{}

// ApplicationFormPartial struct for ApplicationFormPartial
type ApplicationFormPartial struct {
	Description *ApplicationFormPartialDescription `json:"description,omitempty"`
	Icon NullableString `json:"icon,omitempty"`
	CoverImage NullableString `json:"cover_image,omitempty"`
	TeamId *string `json:"team_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Flags NullableInt32 `json:"flags,omitempty"`
	InteractionsEndpointUrl NullableString `json:"interactions_endpoint_url,omitempty"`
	ExplicitContentFilter *int32 `json:"explicit_content_filter,omitempty"`
	MaxParticipants NullableInt32 `json:"max_participants,omitempty"`
	Type NullableInt32 `json:"type,omitempty"`
	Tags []string `json:"tags,omitempty"`
	CustomInstallUrl NullableString `json:"custom_install_url,omitempty"`
	InstallParams NullableApplicationOAuth2InstallParams `json:"install_params,omitempty"`
	RoleConnectionsVerificationUrl NullableString `json:"role_connections_verification_url,omitempty"`
	IntegrationTypesConfig *map[string]ApplicationFormPartialIntegrationTypesConfigValue `json:"integration_types_config,omitempty"`
}

// NewApplicationFormPartial instantiates a new ApplicationFormPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationFormPartial() *ApplicationFormPartial {
	this := ApplicationFormPartial{}
	return &this
}

// NewApplicationFormPartialWithDefaults instantiates a new ApplicationFormPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationFormPartialWithDefaults() *ApplicationFormPartial {
	this := ApplicationFormPartial{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetDescription() ApplicationFormPartialDescription {
	if o == nil || IsNil(o.Description) {
		var ret ApplicationFormPartialDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetDescriptionOk() (*ApplicationFormPartialDescription, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given ApplicationFormPartialDescription and assigns it to the Description field.
func (o *ApplicationFormPartial) SetDescription(v ApplicationFormPartialDescription) {
	o.Description = &v
}


// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationFormPartial) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationFormPartial) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *ApplicationFormPartial) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *ApplicationFormPartial) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *ApplicationFormPartial) UnsetIcon() {
	o.Icon.Unset()
}

// GetCoverImage returns the CoverImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationFormPartial) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImage.Get()
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationFormPartial) GetCoverImageOk() (*string, bool) {
	if o == nil || IsNil(o.CoverImage.Get()) {
		return nil, false
	}
	return o.CoverImage.Get(), o.CoverImage.IsSet()
}

// HasCoverImage returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasCoverImage() bool {
	if o != nil && o.CoverImage.IsSet() {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given NullableString and assigns it to the CoverImage field.
func (o *ApplicationFormPartial) SetCoverImage(v string) {
	o.CoverImage.Set(&v)
}

// SetCoverImageNil sets the value for CoverImage to be an explicit nil
func (o *ApplicationFormPartial) SetCoverImageNil() {
	o.CoverImage.Set(nil)
}

// UnsetCoverImage ensures that no value is present for CoverImage, not even an explicit nil
func (o *ApplicationFormPartial) UnsetCoverImage() {
	o.CoverImage.Unset()
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetTeamId() string {
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *ApplicationFormPartial) SetTeamId(v string) {
	o.TeamId = &v
}


// GetFlags returns the Flags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationFormPartial) GetFlags() int32 {
	if o == nil || IsNil(o.Flags.Get()) {
		var ret int32
		return ret
	}
	return *o.Flags.Get()
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationFormPartial) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags.Get()) {
		return nil, false
	}
	return o.Flags.Get(), o.Flags.IsSet()
}

// HasFlags returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasFlags() bool {
	if o != nil && o.Flags.IsSet() {
		return true
	}

	return false
}

// SetFlags gets a reference to the given NullableInt32 and assigns it to the Flags field.
func (o *ApplicationFormPartial) SetFlags(v int32) {
	o.Flags.Set(&v)
}

// SetFlagsNil sets the value for Flags to be an explicit nil
func (o *ApplicationFormPartial) SetFlagsNil() {
	o.Flags.Set(nil)
}

// UnsetFlags ensures that no value is present for Flags, not even an explicit nil
func (o *ApplicationFormPartial) UnsetFlags() {
	o.Flags.Unset()
}

// GetInteractionsEndpointUrl returns the InteractionsEndpointUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationFormPartial) GetInteractionsEndpointUrl() string {
	if o == nil || IsNil(o.InteractionsEndpointUrl.Get()) {
		var ret string
		return ret
	}
	return *o.InteractionsEndpointUrl.Get()
}

// GetInteractionsEndpointUrlOk returns a tuple with the InteractionsEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationFormPartial) GetInteractionsEndpointUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InteractionsEndpointUrl.Get()) {
		return nil, false
	}
	return o.InteractionsEndpointUrl.Get(), o.InteractionsEndpointUrl.IsSet()
}

// HasInteractionsEndpointUrl returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasInteractionsEndpointUrl() bool {
	if o != nil && o.InteractionsEndpointUrl.IsSet() {
		return true
	}

	return false
}

// SetInteractionsEndpointUrl gets a reference to the given NullableString and assigns it to the InteractionsEndpointUrl field.
func (o *ApplicationFormPartial) SetInteractionsEndpointUrl(v string) {
	o.InteractionsEndpointUrl.Set(&v)
}

// SetInteractionsEndpointUrlNil sets the value for InteractionsEndpointUrl to be an explicit nil
func (o *ApplicationFormPartial) SetInteractionsEndpointUrlNil() {
	o.InteractionsEndpointUrl.Set(nil)
}

// UnsetInteractionsEndpointUrl ensures that no value is present for InteractionsEndpointUrl, not even an explicit nil
func (o *ApplicationFormPartial) UnsetInteractionsEndpointUrl() {
	o.InteractionsEndpointUrl.Unset()
}

// GetExplicitContentFilter returns the ExplicitContentFilter field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetExplicitContentFilter() int32 {
	if o == nil || IsNil(o.ExplicitContentFilter) {
		var ret int32
		return ret
	}
	return *o.ExplicitContentFilter
}

// GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetExplicitContentFilterOk() (*int32, bool) {
	if o == nil || IsNil(o.ExplicitContentFilter) {
		return nil, false
	}
	return o.ExplicitContentFilter, true
}

// HasExplicitContentFilter returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasExplicitContentFilter() bool {
	if o != nil && !IsNil(o.ExplicitContentFilter) {
		return true
	}

	return false
}

// SetExplicitContentFilter gets a reference to the given int32 and assigns it to the ExplicitContentFilter field.
func (o *ApplicationFormPartial) SetExplicitContentFilter(v int32) {
	o.ExplicitContentFilter = &v
}


// GetMaxParticipants returns the MaxParticipants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationFormPartial) GetMaxParticipants() int32 {
	if o == nil || IsNil(o.MaxParticipants.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxParticipants.Get()
}

// GetMaxParticipantsOk returns a tuple with the MaxParticipants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationFormPartial) GetMaxParticipantsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxParticipants.Get()) {
		return nil, false
	}
	return o.MaxParticipants.Get(), o.MaxParticipants.IsSet()
}

// HasMaxParticipants returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasMaxParticipants() bool {
	if o != nil && o.MaxParticipants.IsSet() {
		return true
	}

	return false
}

// SetMaxParticipants gets a reference to the given NullableInt32 and assigns it to the MaxParticipants field.
func (o *ApplicationFormPartial) SetMaxParticipants(v int32) {
	o.MaxParticipants.Set(&v)
}

// SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil
func (o *ApplicationFormPartial) SetMaxParticipantsNil() {
	o.MaxParticipants.Set(nil)
}

// UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
func (o *ApplicationFormPartial) UnsetMaxParticipants() {
	o.MaxParticipants.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationFormPartial) GetType() int32 {
	if o == nil || IsNil(o.Type.Get()) {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationFormPartial) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type.Get()) {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableInt32 and assigns it to the Type field.
func (o *ApplicationFormPartial) SetType(v int32) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *ApplicationFormPartial) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ApplicationFormPartial) UnsetType() {
	o.Type.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationFormPartial) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationFormPartial) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ApplicationFormPartial) SetTags(v []string) {
	o.Tags = v
}


// GetCustomInstallUrl returns the CustomInstallUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationFormPartial) GetCustomInstallUrl() string {
	if o == nil || IsNil(o.CustomInstallUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CustomInstallUrl.Get()
}

// GetCustomInstallUrlOk returns a tuple with the CustomInstallUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationFormPartial) GetCustomInstallUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomInstallUrl.Get()) {
		return nil, false
	}
	return o.CustomInstallUrl.Get(), o.CustomInstallUrl.IsSet()
}

// HasCustomInstallUrl returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasCustomInstallUrl() bool {
	if o != nil && o.CustomInstallUrl.IsSet() {
		return true
	}

	return false
}

// SetCustomInstallUrl gets a reference to the given NullableString and assigns it to the CustomInstallUrl field.
func (o *ApplicationFormPartial) SetCustomInstallUrl(v string) {
	o.CustomInstallUrl.Set(&v)
}

// SetCustomInstallUrlNil sets the value for CustomInstallUrl to be an explicit nil
func (o *ApplicationFormPartial) SetCustomInstallUrlNil() {
	o.CustomInstallUrl.Set(nil)
}

// UnsetCustomInstallUrl ensures that no value is present for CustomInstallUrl, not even an explicit nil
func (o *ApplicationFormPartial) UnsetCustomInstallUrl() {
	o.CustomInstallUrl.Unset()
}

// GetInstallParams returns the InstallParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationFormPartial) GetInstallParams() ApplicationOAuth2InstallParams {
	if o == nil || IsNil(o.InstallParams.Get()) {
		var ret ApplicationOAuth2InstallParams
		return ret
	}
	return *o.InstallParams.Get()
}

// GetInstallParamsOk returns a tuple with the InstallParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationFormPartial) GetInstallParamsOk() (*ApplicationOAuth2InstallParams, bool) {
	if o == nil || IsNil(o.InstallParams.Get()) {
		return nil, false
	}
	return o.InstallParams.Get(), o.InstallParams.IsSet()
}

// HasInstallParams returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasInstallParams() bool {
	if o != nil && o.InstallParams.IsSet() {
		return true
	}

	return false
}

// SetInstallParams gets a reference to the given NullableApplicationOAuth2InstallParams and assigns it to the InstallParams field.
func (o *ApplicationFormPartial) SetInstallParams(v ApplicationOAuth2InstallParams) {
	o.InstallParams.Set(&v)
}

// SetInstallParamsNil sets the value for InstallParams to be an explicit nil
func (o *ApplicationFormPartial) SetInstallParamsNil() {
	o.InstallParams.Set(nil)
}

// UnsetInstallParams ensures that no value is present for InstallParams, not even an explicit nil
func (o *ApplicationFormPartial) UnsetInstallParams() {
	o.InstallParams.Unset()
}

// GetRoleConnectionsVerificationUrl returns the RoleConnectionsVerificationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationFormPartial) GetRoleConnectionsVerificationUrl() string {
	if o == nil || IsNil(o.RoleConnectionsVerificationUrl.Get()) {
		var ret string
		return ret
	}
	return *o.RoleConnectionsVerificationUrl.Get()
}

// GetRoleConnectionsVerificationUrlOk returns a tuple with the RoleConnectionsVerificationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationFormPartial) GetRoleConnectionsVerificationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RoleConnectionsVerificationUrl.Get()) {
		return nil, false
	}
	return o.RoleConnectionsVerificationUrl.Get(), o.RoleConnectionsVerificationUrl.IsSet()
}

// HasRoleConnectionsVerificationUrl returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasRoleConnectionsVerificationUrl() bool {
	if o != nil && o.RoleConnectionsVerificationUrl.IsSet() {
		return true
	}

	return false
}

// SetRoleConnectionsVerificationUrl gets a reference to the given NullableString and assigns it to the RoleConnectionsVerificationUrl field.
func (o *ApplicationFormPartial) SetRoleConnectionsVerificationUrl(v string) {
	o.RoleConnectionsVerificationUrl.Set(&v)
}

// SetRoleConnectionsVerificationUrlNil sets the value for RoleConnectionsVerificationUrl to be an explicit nil
func (o *ApplicationFormPartial) SetRoleConnectionsVerificationUrlNil() {
	o.RoleConnectionsVerificationUrl.Set(nil)
}

// UnsetRoleConnectionsVerificationUrl ensures that no value is present for RoleConnectionsVerificationUrl, not even an explicit nil
func (o *ApplicationFormPartial) UnsetRoleConnectionsVerificationUrl() {
	o.RoleConnectionsVerificationUrl.Unset()
}

// GetIntegrationTypesConfig returns the IntegrationTypesConfig field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetIntegrationTypesConfig() map[string]ApplicationFormPartialIntegrationTypesConfigValue {
	if o == nil || IsNil(o.IntegrationTypesConfig) {
		var ret map[string]ApplicationFormPartialIntegrationTypesConfigValue
		return ret
	}
	return *o.IntegrationTypesConfig
}

// GetIntegrationTypesConfigOk returns a tuple with the IntegrationTypesConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetIntegrationTypesConfigOk() (*map[string]ApplicationFormPartialIntegrationTypesConfigValue, bool) {
	if o == nil || IsNil(o.IntegrationTypesConfig) {
		return nil, false
	}
	return o.IntegrationTypesConfig, true
}

// HasIntegrationTypesConfig returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasIntegrationTypesConfig() bool {
	if o != nil && !IsNil(o.IntegrationTypesConfig) {
		return true
	}

	return false
}

// SetIntegrationTypesConfig gets a reference to the given map[string]ApplicationFormPartialIntegrationTypesConfigValue and assigns it to the IntegrationTypesConfig field.
func (o *ApplicationFormPartial) SetIntegrationTypesConfig(v map[string]ApplicationFormPartialIntegrationTypesConfigValue) {
	o.IntegrationTypesConfig = &v
}


func (o ApplicationFormPartial) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationFormPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.CoverImage.IsSet() {
		toSerialize["cover_image"] = o.CoverImage.Get()
	}
	if !IsNil(o.TeamId) {
		toSerialize["team_id"] = o.TeamId
	}
	if o.Flags.IsSet() {
		toSerialize["flags"] = o.Flags.Get()
	}
	if o.InteractionsEndpointUrl.IsSet() {
		toSerialize["interactions_endpoint_url"] = o.InteractionsEndpointUrl.Get()
	}
	if !IsNil(o.ExplicitContentFilter) {
		toSerialize["explicit_content_filter"] = o.ExplicitContentFilter
	}
	if o.MaxParticipants.IsSet() {
		toSerialize["max_participants"] = o.MaxParticipants.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.CustomInstallUrl.IsSet() {
		toSerialize["custom_install_url"] = o.CustomInstallUrl.Get()
	}
	if o.InstallParams.IsSet() {
		toSerialize["install_params"] = o.InstallParams.Get()
	}
	if o.RoleConnectionsVerificationUrl.IsSet() {
		toSerialize["role_connections_verification_url"] = o.RoleConnectionsVerificationUrl.Get()
	}
	if !IsNil(o.IntegrationTypesConfig) {
		toSerialize["integration_types_config"] = o.IntegrationTypesConfig
	}
	return toSerialize, nil
}

type NullableApplicationFormPartial struct {
	value *ApplicationFormPartial
	isSet bool
}

func (v NullableApplicationFormPartial) Get() *ApplicationFormPartial {
	return v.value
}

func (v *NullableApplicationFormPartial) Set(val *ApplicationFormPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationFormPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationFormPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationFormPartial(val *ApplicationFormPartial) *NullableApplicationFormPartial {
	return &NullableApplicationFormPartial{value: val, isSet: true}
}

func (v NullableApplicationFormPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationFormPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


