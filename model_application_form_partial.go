/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2026 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2026-01-08T01:38:41.954257418Z[Etc/UTC]
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
)

// checks if the ApplicationFormPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationFormPartial{}

// ApplicationFormPartial struct for ApplicationFormPartial
type ApplicationFormPartial struct {
	Description *ApplicationFormPartialDescription `json:"description,omitempty"`
	Icon *string `json:"icon,omitempty"`
	CoverImage *string `json:"cover_image,omitempty"`
	TeamId *GetEntitlementsSkuIdsParameterOneOfInner `json:"team_id,omitempty"`
	Flags *int32 `json:"flags,omitempty"`
	InteractionsEndpointUrl *string `json:"interactions_endpoint_url,omitempty"`
	ExplicitContentFilter *ApplicationFormPartialExplicitContentFilter `json:"explicit_content_filter,omitempty"`
	MaxParticipants *int32 `json:"max_participants,omitempty"`
	Type *ApplicationFormPartialType `json:"type,omitempty"`
	Tags []string `json:"tags,omitempty"`
	CustomInstallUrl *string `json:"custom_install_url,omitempty"`
	InstallParams *ApplicationFormPartialInstallParams `json:"install_params,omitempty"`
	RoleConnectionsVerificationUrl *string `json:"role_connections_verification_url,omitempty"`
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


// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *ApplicationFormPartial) SetIcon(v string) {
	o.Icon = &v
}


// GetCoverImage returns the CoverImage field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage) {
		var ret string
		return ret
	}
	return *o.CoverImage
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetCoverImageOk() (*string, bool) {
	if o == nil || IsNil(o.CoverImage) {
		return nil, false
	}
	return o.CoverImage, true
}

// HasCoverImage returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasCoverImage() bool {
	if o != nil && !IsNil(o.CoverImage) {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given string and assigns it to the CoverImage field.
func (o *ApplicationFormPartial) SetCoverImage(v string) {
	o.CoverImage = &v
}


// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetTeamId() GetEntitlementsSkuIdsParameterOneOfInner {
	if o == nil || IsNil(o.TeamId) {
		var ret GetEntitlementsSkuIdsParameterOneOfInner
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetTeamIdOk() (*GetEntitlementsSkuIdsParameterOneOfInner, bool) {
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

// SetTeamId gets a reference to the given GetEntitlementsSkuIdsParameterOneOfInner and assigns it to the TeamId field.
func (o *ApplicationFormPartial) SetTeamId(v GetEntitlementsSkuIdsParameterOneOfInner) {
	o.TeamId = &v
}


// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetFlags() int32 {
	if o == nil || IsNil(o.Flags) {
		var ret int32
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetFlagsOk() (*int32, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given int32 and assigns it to the Flags field.
func (o *ApplicationFormPartial) SetFlags(v int32) {
	o.Flags = &v
}


// GetInteractionsEndpointUrl returns the InteractionsEndpointUrl field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetInteractionsEndpointUrl() string {
	if o == nil || IsNil(o.InteractionsEndpointUrl) {
		var ret string
		return ret
	}
	return *o.InteractionsEndpointUrl
}

// GetInteractionsEndpointUrlOk returns a tuple with the InteractionsEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetInteractionsEndpointUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InteractionsEndpointUrl) {
		return nil, false
	}
	return o.InteractionsEndpointUrl, true
}

// HasInteractionsEndpointUrl returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasInteractionsEndpointUrl() bool {
	if o != nil && !IsNil(o.InteractionsEndpointUrl) {
		return true
	}

	return false
}

// SetInteractionsEndpointUrl gets a reference to the given string and assigns it to the InteractionsEndpointUrl field.
func (o *ApplicationFormPartial) SetInteractionsEndpointUrl(v string) {
	o.InteractionsEndpointUrl = &v
}


// GetExplicitContentFilter returns the ExplicitContentFilter field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetExplicitContentFilter() ApplicationFormPartialExplicitContentFilter {
	if o == nil || IsNil(o.ExplicitContentFilter) {
		var ret ApplicationFormPartialExplicitContentFilter
		return ret
	}
	return *o.ExplicitContentFilter
}

// GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetExplicitContentFilterOk() (*ApplicationFormPartialExplicitContentFilter, bool) {
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

// SetExplicitContentFilter gets a reference to the given ApplicationFormPartialExplicitContentFilter and assigns it to the ExplicitContentFilter field.
func (o *ApplicationFormPartial) SetExplicitContentFilter(v ApplicationFormPartialExplicitContentFilter) {
	o.ExplicitContentFilter = &v
}


// GetMaxParticipants returns the MaxParticipants field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetMaxParticipants() int32 {
	if o == nil || IsNil(o.MaxParticipants) {
		var ret int32
		return ret
	}
	return *o.MaxParticipants
}

// GetMaxParticipantsOk returns a tuple with the MaxParticipants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetMaxParticipantsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxParticipants) {
		return nil, false
	}
	return o.MaxParticipants, true
}

// HasMaxParticipants returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasMaxParticipants() bool {
	if o != nil && !IsNil(o.MaxParticipants) {
		return true
	}

	return false
}

// SetMaxParticipants gets a reference to the given int32 and assigns it to the MaxParticipants field.
func (o *ApplicationFormPartial) SetMaxParticipants(v int32) {
	o.MaxParticipants = &v
}


// GetType returns the Type field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetType() ApplicationFormPartialType {
	if o == nil || IsNil(o.Type) {
		var ret ApplicationFormPartialType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetTypeOk() (*ApplicationFormPartialType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ApplicationFormPartialType and assigns it to the Type field.
func (o *ApplicationFormPartial) SetType(v ApplicationFormPartialType) {
	o.Type = &v
}


// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
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


// GetCustomInstallUrl returns the CustomInstallUrl field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetCustomInstallUrl() string {
	if o == nil || IsNil(o.CustomInstallUrl) {
		var ret string
		return ret
	}
	return *o.CustomInstallUrl
}

// GetCustomInstallUrlOk returns a tuple with the CustomInstallUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetCustomInstallUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomInstallUrl) {
		return nil, false
	}
	return o.CustomInstallUrl, true
}

// HasCustomInstallUrl returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasCustomInstallUrl() bool {
	if o != nil && !IsNil(o.CustomInstallUrl) {
		return true
	}

	return false
}

// SetCustomInstallUrl gets a reference to the given string and assigns it to the CustomInstallUrl field.
func (o *ApplicationFormPartial) SetCustomInstallUrl(v string) {
	o.CustomInstallUrl = &v
}


// GetInstallParams returns the InstallParams field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetInstallParams() ApplicationFormPartialInstallParams {
	if o == nil || IsNil(o.InstallParams) {
		var ret ApplicationFormPartialInstallParams
		return ret
	}
	return *o.InstallParams
}

// GetInstallParamsOk returns a tuple with the InstallParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetInstallParamsOk() (*ApplicationFormPartialInstallParams, bool) {
	if o == nil || IsNil(o.InstallParams) {
		return nil, false
	}
	return o.InstallParams, true
}

// HasInstallParams returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasInstallParams() bool {
	if o != nil && !IsNil(o.InstallParams) {
		return true
	}

	return false
}

// SetInstallParams gets a reference to the given ApplicationFormPartialInstallParams and assigns it to the InstallParams field.
func (o *ApplicationFormPartial) SetInstallParams(v ApplicationFormPartialInstallParams) {
	o.InstallParams = &v
}


// GetRoleConnectionsVerificationUrl returns the RoleConnectionsVerificationUrl field value if set, zero value otherwise.
func (o *ApplicationFormPartial) GetRoleConnectionsVerificationUrl() string {
	if o == nil || IsNil(o.RoleConnectionsVerificationUrl) {
		var ret string
		return ret
	}
	return *o.RoleConnectionsVerificationUrl
}

// GetRoleConnectionsVerificationUrlOk returns a tuple with the RoleConnectionsVerificationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFormPartial) GetRoleConnectionsVerificationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RoleConnectionsVerificationUrl) {
		return nil, false
	}
	return o.RoleConnectionsVerificationUrl, true
}

// HasRoleConnectionsVerificationUrl returns a boolean if a field has been set.
func (o *ApplicationFormPartial) HasRoleConnectionsVerificationUrl() bool {
	if o != nil && !IsNil(o.RoleConnectionsVerificationUrl) {
		return true
	}

	return false
}

// SetRoleConnectionsVerificationUrl gets a reference to the given string and assigns it to the RoleConnectionsVerificationUrl field.
func (o *ApplicationFormPartial) SetRoleConnectionsVerificationUrl(v string) {
	o.RoleConnectionsVerificationUrl = &v
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
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.CoverImage) {
		toSerialize["cover_image"] = o.CoverImage
	}
	if !IsNil(o.TeamId) {
		toSerialize["team_id"] = o.TeamId
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.InteractionsEndpointUrl) {
		toSerialize["interactions_endpoint_url"] = o.InteractionsEndpointUrl
	}
	if !IsNil(o.ExplicitContentFilter) {
		toSerialize["explicit_content_filter"] = o.ExplicitContentFilter
	}
	if !IsNil(o.MaxParticipants) {
		toSerialize["max_participants"] = o.MaxParticipants
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomInstallUrl) {
		toSerialize["custom_install_url"] = o.CustomInstallUrl
	}
	if !IsNil(o.InstallParams) {
		toSerialize["install_params"] = o.InstallParams
	}
	if !IsNil(o.RoleConnectionsVerificationUrl) {
		toSerialize["role_connections_verification_url"] = o.RoleConnectionsVerificationUrl
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

