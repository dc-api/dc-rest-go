/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-19T09:30:49.800547817Z[Etc/UTC]
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

// checks if the ApplicationCommandResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCommandResponse{}

// ApplicationCommandResponse struct for ApplicationCommandResponse
type ApplicationCommandResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	ApplicationId string `json:"application_id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Version string `json:"version" validate:"regexp=^(0|[1-9][0-9]*)$"`
	DefaultMemberPermissions NullableString `json:"default_member_permissions,omitempty"`
	Type int32 `json:"type"`
	Name string `json:"name"`
	NameLocalized NullableString `json:"name_localized,omitempty"`
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`
	Description string `json:"description"`
	DescriptionLocalized NullableString `json:"description_localized,omitempty"`
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`
	GuildId *string `json:"guild_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	DmPermission NullableBool `json:"dm_permission,omitempty"`
	Contexts []int32 `json:"contexts,omitempty"`
	IntegrationTypes []int32 `json:"integration_types,omitempty"`
	Options []ApplicationCommandResponseOptionsInner `json:"options,omitempty"`
	Nsfw NullableBool `json:"nsfw,omitempty"`
}

type _ApplicationCommandResponse ApplicationCommandResponse

// NewApplicationCommandResponse instantiates a new ApplicationCommandResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCommandResponse(id string, applicationId string, version string, type_ int32, name string, description string) *ApplicationCommandResponse {
	this := ApplicationCommandResponse{}
	this.Id = id
	this.ApplicationId = applicationId
	this.Version = version
	this.Type = type_
	this.Name = name
	this.Description = description
	return &this
}

// NewApplicationCommandResponseWithDefaults instantiates a new ApplicationCommandResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCommandResponseWithDefaults() *ApplicationCommandResponse {
	this := ApplicationCommandResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationCommandResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationCommandResponse) SetId(v string) {
	o.Id = v
}

// GetApplicationId returns the ApplicationId field value
func (o *ApplicationCommandResponse) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandResponse) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *ApplicationCommandResponse) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetVersion returns the Version field value
func (o *ApplicationCommandResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ApplicationCommandResponse) SetVersion(v string) {
	o.Version = v
}

// GetDefaultMemberPermissions returns the DefaultMemberPermissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandResponse) GetDefaultMemberPermissions() string {
	if o == nil || IsNil(o.DefaultMemberPermissions.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultMemberPermissions.Get()
}

// GetDefaultMemberPermissionsOk returns a tuple with the DefaultMemberPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandResponse) GetDefaultMemberPermissionsOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultMemberPermissions.Get()) {
		return nil, false
	}
	return o.DefaultMemberPermissions.Get(), o.DefaultMemberPermissions.IsSet()
}

// HasDefaultMemberPermissions returns a boolean if a field has been set.
func (o *ApplicationCommandResponse) HasDefaultMemberPermissions() bool {
	if o != nil && o.DefaultMemberPermissions.IsSet() {
		return true
	}

	return false
}

// SetDefaultMemberPermissions gets a reference to the given NullableString and assigns it to the DefaultMemberPermissions field.
func (o *ApplicationCommandResponse) SetDefaultMemberPermissions(v string) {
	o.DefaultMemberPermissions.Set(&v)
}

// SetDefaultMemberPermissionsNil sets the value for DefaultMemberPermissions to be an explicit nil
func (o *ApplicationCommandResponse) SetDefaultMemberPermissionsNil() {
	o.DefaultMemberPermissions.Set(nil)
}

// UnsetDefaultMemberPermissions ensures that no value is present for DefaultMemberPermissions, not even an explicit nil
func (o *ApplicationCommandResponse) UnsetDefaultMemberPermissions() {
	o.DefaultMemberPermissions.Unset()
}

// GetType returns the Type field value
func (o *ApplicationCommandResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationCommandResponse) SetType(v int32) {
	o.Type = v
}

// GetName returns the Name field value
func (o *ApplicationCommandResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCommandResponse) SetName(v string) {
	o.Name = v
}

// GetNameLocalized returns the NameLocalized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandResponse) GetNameLocalized() string {
	if o == nil || IsNil(o.NameLocalized.Get()) {
		var ret string
		return ret
	}
	return *o.NameLocalized.Get()
}

// GetNameLocalizedOk returns a tuple with the NameLocalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandResponse) GetNameLocalizedOk() (*string, bool) {
	if o == nil || IsNil(o.NameLocalized.Get()) {
		return nil, false
	}
	return o.NameLocalized.Get(), o.NameLocalized.IsSet()
}

// HasNameLocalized returns a boolean if a field has been set.
func (o *ApplicationCommandResponse) HasNameLocalized() bool {
	if o != nil && o.NameLocalized.IsSet() {
		return true
	}

	return false
}

// SetNameLocalized gets a reference to the given NullableString and assigns it to the NameLocalized field.
func (o *ApplicationCommandResponse) SetNameLocalized(v string) {
	o.NameLocalized.Set(&v)
}

// SetNameLocalizedNil sets the value for NameLocalized to be an explicit nil
func (o *ApplicationCommandResponse) SetNameLocalizedNil() {
	o.NameLocalized.Set(nil)
}

// UnsetNameLocalized ensures that no value is present for NameLocalized, not even an explicit nil
func (o *ApplicationCommandResponse) UnsetNameLocalized() {
	o.NameLocalized.Unset()
}

// GetNameLocalizations returns the NameLocalizations field value if set, zero value otherwise.
func (o *ApplicationCommandResponse) GetNameLocalizations() map[string]string {
	if o == nil || IsNil(o.NameLocalizations) {
		var ret map[string]string
		return ret
	}
	return o.NameLocalizations
}

// GetNameLocalizationsOk returns a tuple with the NameLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandResponse) GetNameLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.NameLocalizations) {
		return map[string]string{}, false
	}
	return o.NameLocalizations, true
}

// HasNameLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandResponse) HasNameLocalizations() bool {
	if o != nil && !IsNil(o.NameLocalizations) {
		return true
	}

	return false
}

// SetNameLocalizations gets a reference to the given map[string]string and assigns it to the NameLocalizations field.
func (o *ApplicationCommandResponse) SetNameLocalizations(v map[string]string) {
	o.NameLocalizations = v
}


// GetDescription returns the Description field value
func (o *ApplicationCommandResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApplicationCommandResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ApplicationCommandResponse) SetDescription(v string) {
	o.Description = v
}

// GetDescriptionLocalized returns the DescriptionLocalized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandResponse) GetDescriptionLocalized() string {
	if o == nil || IsNil(o.DescriptionLocalized.Get()) {
		var ret string
		return ret
	}
	return *o.DescriptionLocalized.Get()
}

// GetDescriptionLocalizedOk returns a tuple with the DescriptionLocalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandResponse) GetDescriptionLocalizedOk() (*string, bool) {
	if o == nil || IsNil(o.DescriptionLocalized.Get()) {
		return nil, false
	}
	return o.DescriptionLocalized.Get(), o.DescriptionLocalized.IsSet()
}

// HasDescriptionLocalized returns a boolean if a field has been set.
func (o *ApplicationCommandResponse) HasDescriptionLocalized() bool {
	if o != nil && o.DescriptionLocalized.IsSet() {
		return true
	}

	return false
}

// SetDescriptionLocalized gets a reference to the given NullableString and assigns it to the DescriptionLocalized field.
func (o *ApplicationCommandResponse) SetDescriptionLocalized(v string) {
	o.DescriptionLocalized.Set(&v)
}

// SetDescriptionLocalizedNil sets the value for DescriptionLocalized to be an explicit nil
func (o *ApplicationCommandResponse) SetDescriptionLocalizedNil() {
	o.DescriptionLocalized.Set(nil)
}

// UnsetDescriptionLocalized ensures that no value is present for DescriptionLocalized, not even an explicit nil
func (o *ApplicationCommandResponse) UnsetDescriptionLocalized() {
	o.DescriptionLocalized.Unset()
}

// GetDescriptionLocalizations returns the DescriptionLocalizations field value if set, zero value otherwise.
func (o *ApplicationCommandResponse) GetDescriptionLocalizations() map[string]string {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		var ret map[string]string
		return ret
	}
	return o.DescriptionLocalizations
}

// GetDescriptionLocalizationsOk returns a tuple with the DescriptionLocalizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandResponse) GetDescriptionLocalizationsOk() (map[string]string, bool) {
	if o == nil || IsNil(o.DescriptionLocalizations) {
		return map[string]string{}, false
	}
	return o.DescriptionLocalizations, true
}

// HasDescriptionLocalizations returns a boolean if a field has been set.
func (o *ApplicationCommandResponse) HasDescriptionLocalizations() bool {
	if o != nil && !IsNil(o.DescriptionLocalizations) {
		return true
	}

	return false
}

// SetDescriptionLocalizations gets a reference to the given map[string]string and assigns it to the DescriptionLocalizations field.
func (o *ApplicationCommandResponse) SetDescriptionLocalizations(v map[string]string) {
	o.DescriptionLocalizations = v
}


// GetGuildId returns the GuildId field value if set, zero value otherwise.
func (o *ApplicationCommandResponse) GetGuildId() string {
	if o == nil || IsNil(o.GuildId) {
		var ret string
		return ret
	}
	return *o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCommandResponse) GetGuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.GuildId) {
		return nil, false
	}
	return o.GuildId, true
}

// HasGuildId returns a boolean if a field has been set.
func (o *ApplicationCommandResponse) HasGuildId() bool {
	if o != nil && !IsNil(o.GuildId) {
		return true
	}

	return false
}

// SetGuildId gets a reference to the given string and assigns it to the GuildId field.
func (o *ApplicationCommandResponse) SetGuildId(v string) {
	o.GuildId = &v
}


// GetDmPermission returns the DmPermission field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandResponse) GetDmPermission() bool {
	if o == nil || IsNil(o.DmPermission.Get()) {
		var ret bool
		return ret
	}
	return *o.DmPermission.Get()
}

// GetDmPermissionOk returns a tuple with the DmPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandResponse) GetDmPermissionOk() (*bool, bool) {
	if o == nil || IsNil(o.DmPermission.Get()) {
		return nil, false
	}
	return o.DmPermission.Get(), o.DmPermission.IsSet()
}

// HasDmPermission returns a boolean if a field has been set.
func (o *ApplicationCommandResponse) HasDmPermission() bool {
	if o != nil && o.DmPermission.IsSet() {
		return true
	}

	return false
}

// SetDmPermission gets a reference to the given NullableBool and assigns it to the DmPermission field.
func (o *ApplicationCommandResponse) SetDmPermission(v bool) {
	o.DmPermission.Set(&v)
}

// SetDmPermissionNil sets the value for DmPermission to be an explicit nil
func (o *ApplicationCommandResponse) SetDmPermissionNil() {
	o.DmPermission.Set(nil)
}

// UnsetDmPermission ensures that no value is present for DmPermission, not even an explicit nil
func (o *ApplicationCommandResponse) UnsetDmPermission() {
	o.DmPermission.Unset()
}

// GetContexts returns the Contexts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandResponse) GetContexts() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Contexts
}

// GetContextsOk returns a tuple with the Contexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandResponse) GetContextsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contexts, true
}

// HasContexts returns a boolean if a field has been set.
func (o *ApplicationCommandResponse) HasContexts() bool {
	if o != nil && !IsNil(o.Contexts) {
		return true
	}

	return false
}

// SetContexts gets a reference to the given []int32 and assigns it to the Contexts field.
func (o *ApplicationCommandResponse) SetContexts(v []int32) {
	o.Contexts = v
}


// GetIntegrationTypes returns the IntegrationTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandResponse) GetIntegrationTypes() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.IntegrationTypes
}

// GetIntegrationTypesOk returns a tuple with the IntegrationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandResponse) GetIntegrationTypesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntegrationTypes, true
}

// HasIntegrationTypes returns a boolean if a field has been set.
func (o *ApplicationCommandResponse) HasIntegrationTypes() bool {
	if o != nil && !IsNil(o.IntegrationTypes) {
		return true
	}

	return false
}

// SetIntegrationTypes gets a reference to the given []int32 and assigns it to the IntegrationTypes field.
func (o *ApplicationCommandResponse) SetIntegrationTypes(v []int32) {
	o.IntegrationTypes = v
}


// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandResponse) GetOptions() []ApplicationCommandResponseOptionsInner {
	if o == nil {
		var ret []ApplicationCommandResponseOptionsInner
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandResponse) GetOptionsOk() ([]ApplicationCommandResponseOptionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ApplicationCommandResponse) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ApplicationCommandResponseOptionsInner and assigns it to the Options field.
func (o *ApplicationCommandResponse) SetOptions(v []ApplicationCommandResponseOptionsInner) {
	o.Options = v
}


// GetNsfw returns the Nsfw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationCommandResponse) GetNsfw() bool {
	if o == nil || IsNil(o.Nsfw.Get()) {
		var ret bool
		return ret
	}
	return *o.Nsfw.Get()
}

// GetNsfwOk returns a tuple with the Nsfw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationCommandResponse) GetNsfwOk() (*bool, bool) {
	if o == nil || IsNil(o.Nsfw.Get()) {
		return nil, false
	}
	return o.Nsfw.Get(), o.Nsfw.IsSet()
}

// HasNsfw returns a boolean if a field has been set.
func (o *ApplicationCommandResponse) HasNsfw() bool {
	if o != nil && o.Nsfw.IsSet() {
		return true
	}

	return false
}

// SetNsfw gets a reference to the given NullableBool and assigns it to the Nsfw field.
func (o *ApplicationCommandResponse) SetNsfw(v bool) {
	o.Nsfw.Set(&v)
}

// SetNsfwNil sets the value for Nsfw to be an explicit nil
func (o *ApplicationCommandResponse) SetNsfwNil() {
	o.Nsfw.Set(nil)
}

// UnsetNsfw ensures that no value is present for Nsfw, not even an explicit nil
func (o *ApplicationCommandResponse) UnsetNsfw() {
	o.Nsfw.Unset()
}

func (o ApplicationCommandResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCommandResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["application_id"] = o.ApplicationId
	toSerialize["version"] = o.Version
	if o.DefaultMemberPermissions.IsSet() {
		toSerialize["default_member_permissions"] = o.DefaultMemberPermissions.Get()
	}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if o.NameLocalized.IsSet() {
		toSerialize["name_localized"] = o.NameLocalized.Get()
	}
	if !IsNil(o.NameLocalizations) {
		toSerialize["name_localizations"] = o.NameLocalizations
	}
	toSerialize["description"] = o.Description
	if o.DescriptionLocalized.IsSet() {
		toSerialize["description_localized"] = o.DescriptionLocalized.Get()
	}
	if !IsNil(o.DescriptionLocalizations) {
		toSerialize["description_localizations"] = o.DescriptionLocalizations
	}
	if !IsNil(o.GuildId) {
		toSerialize["guild_id"] = o.GuildId
	}
	if o.DmPermission.IsSet() {
		toSerialize["dm_permission"] = o.DmPermission.Get()
	}
	if o.Contexts != nil {
		toSerialize["contexts"] = o.Contexts
	}
	if o.IntegrationTypes != nil {
		toSerialize["integration_types"] = o.IntegrationTypes
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Nsfw.IsSet() {
		toSerialize["nsfw"] = o.Nsfw.Get()
	}
	return toSerialize, nil
}

func (o *ApplicationCommandResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"application_id",
		"version",
		"type",
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

	varApplicationCommandResponse := _ApplicationCommandResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCommandResponse)

	if err != nil {
		return err
	}

	*o = ApplicationCommandResponse(varApplicationCommandResponse)

	return err
}

type NullableApplicationCommandResponse struct {
	value *ApplicationCommandResponse
	isSet bool
}

func (v NullableApplicationCommandResponse) Get() *ApplicationCommandResponse {
	return v.value
}

func (v *NullableApplicationCommandResponse) Set(val *ApplicationCommandResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCommandResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCommandResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCommandResponse(val *ApplicationCommandResponse) *NullableApplicationCommandResponse {
	return &NullableApplicationCommandResponse{value: val, isSet: true}
}

func (v NullableApplicationCommandResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCommandResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


