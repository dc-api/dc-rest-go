/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-09-06T05:32:22.317073992Z[Etc/UTC]
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

// checks if the InviteApplicationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteApplicationResponse{}

// InviteApplicationResponse struct for InviteApplicationResponse
type InviteApplicationResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Icon NullableString `json:"icon,omitempty"`
	Description string `json:"description"`
	Type NullableInt32 `json:"type,omitempty"`
	CoverImage *string `json:"cover_image,omitempty"`
	PrimarySkuId *string `json:"primary_sku_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Bot *UserResponse `json:"bot,omitempty"`
	Slug *string `json:"slug,omitempty"`
	GuildId *string `json:"guild_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	RpcOrigins []*string `json:"rpc_origins,omitempty"`
	BotPublic *bool `json:"bot_public,omitempty"`
	BotRequireCodeGrant *bool `json:"bot_require_code_grant,omitempty"`
	TermsOfServiceUrl *string `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyUrl *string `json:"privacy_policy_url,omitempty"`
	CustomInstallUrl *string `json:"custom_install_url,omitempty"`
	InstallParams *ApplicationOAuth2InstallParamsResponse `json:"install_params,omitempty"`
	IntegrationTypesConfig map[string]ApplicationIntegrationTypeConfigurationResponse `json:"integration_types_config,omitempty"`
	VerifyKey string `json:"verify_key"`
	Flags int32 `json:"flags"`
	MaxParticipants NullableInt32 `json:"max_participants,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

type _InviteApplicationResponse InviteApplicationResponse

// NewInviteApplicationResponse instantiates a new InviteApplicationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteApplicationResponse(id string, name string, description string, verifyKey string, flags int32) *InviteApplicationResponse {
	this := InviteApplicationResponse{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.VerifyKey = verifyKey
	this.Flags = flags
	return &this
}

// NewInviteApplicationResponseWithDefaults instantiates a new InviteApplicationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteApplicationResponseWithDefaults() *InviteApplicationResponse {
	this := InviteApplicationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *InviteApplicationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InviteApplicationResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *InviteApplicationResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InviteApplicationResponse) SetName(v string) {
	o.Name = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteApplicationResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteApplicationResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *InviteApplicationResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *InviteApplicationResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *InviteApplicationResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetDescription returns the Description field value
func (o *InviteApplicationResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *InviteApplicationResponse) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteApplicationResponse) GetType() int32 {
	if o == nil || IsNil(o.Type.Get()) {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteApplicationResponse) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type.Get()) {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableInt32 and assigns it to the Type field.
func (o *InviteApplicationResponse) SetType(v int32) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *InviteApplicationResponse) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *InviteApplicationResponse) UnsetType() {
	o.Type.Unset()
}

// GetCoverImage returns the CoverImage field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage) {
		var ret string
		return ret
	}
	return *o.CoverImage
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetCoverImageOk() (*string, bool) {
	if o == nil || IsNil(o.CoverImage) {
		return nil, false
	}
	return o.CoverImage, true
}

// HasCoverImage returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasCoverImage() bool {
	if o != nil && !IsNil(o.CoverImage) {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given string and assigns it to the CoverImage field.
func (o *InviteApplicationResponse) SetCoverImage(v string) {
	o.CoverImage = &v
}


// GetPrimarySkuId returns the PrimarySkuId field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetPrimarySkuId() string {
	if o == nil || IsNil(o.PrimarySkuId) {
		var ret string
		return ret
	}
	return *o.PrimarySkuId
}

// GetPrimarySkuIdOk returns a tuple with the PrimarySkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetPrimarySkuIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrimarySkuId) {
		return nil, false
	}
	return o.PrimarySkuId, true
}

// HasPrimarySkuId returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasPrimarySkuId() bool {
	if o != nil && !IsNil(o.PrimarySkuId) {
		return true
	}

	return false
}

// SetPrimarySkuId gets a reference to the given string and assigns it to the PrimarySkuId field.
func (o *InviteApplicationResponse) SetPrimarySkuId(v string) {
	o.PrimarySkuId = &v
}


// GetBot returns the Bot field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetBot() UserResponse {
	if o == nil || IsNil(o.Bot) {
		var ret UserResponse
		return ret
	}
	return *o.Bot
}

// GetBotOk returns a tuple with the Bot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetBotOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.Bot) {
		return nil, false
	}
	return o.Bot, true
}

// HasBot returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasBot() bool {
	if o != nil && !IsNil(o.Bot) {
		return true
	}

	return false
}

// SetBot gets a reference to the given UserResponse and assigns it to the Bot field.
func (o *InviteApplicationResponse) SetBot(v UserResponse) {
	o.Bot = &v
}


// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *InviteApplicationResponse) SetSlug(v string) {
	o.Slug = &v
}


// GetGuildId returns the GuildId field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetGuildId() string {
	if o == nil || IsNil(o.GuildId) {
		var ret string
		return ret
	}
	return *o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetGuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.GuildId) {
		return nil, false
	}
	return o.GuildId, true
}

// HasGuildId returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasGuildId() bool {
	if o != nil && !IsNil(o.GuildId) {
		return true
	}

	return false
}

// SetGuildId gets a reference to the given string and assigns it to the GuildId field.
func (o *InviteApplicationResponse) SetGuildId(v string) {
	o.GuildId = &v
}


// GetRpcOrigins returns the RpcOrigins field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetRpcOrigins() []*string {
	if o == nil || IsNil(o.RpcOrigins) {
		var ret []*string
		return ret
	}
	return o.RpcOrigins
}

// GetRpcOriginsOk returns a tuple with the RpcOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetRpcOriginsOk() ([]*string, bool) {
	if o == nil || IsNil(o.RpcOrigins) {
		return nil, false
	}
	return o.RpcOrigins, true
}

// HasRpcOrigins returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasRpcOrigins() bool {
	if o != nil && !IsNil(o.RpcOrigins) {
		return true
	}

	return false
}

// SetRpcOrigins gets a reference to the given []*string and assigns it to the RpcOrigins field.
func (o *InviteApplicationResponse) SetRpcOrigins(v []*string) {
	o.RpcOrigins = v
}


// GetBotPublic returns the BotPublic field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetBotPublic() bool {
	if o == nil || IsNil(o.BotPublic) {
		var ret bool
		return ret
	}
	return *o.BotPublic
}

// GetBotPublicOk returns a tuple with the BotPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetBotPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.BotPublic) {
		return nil, false
	}
	return o.BotPublic, true
}

// HasBotPublic returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasBotPublic() bool {
	if o != nil && !IsNil(o.BotPublic) {
		return true
	}

	return false
}

// SetBotPublic gets a reference to the given bool and assigns it to the BotPublic field.
func (o *InviteApplicationResponse) SetBotPublic(v bool) {
	o.BotPublic = &v
}


// GetBotRequireCodeGrant returns the BotRequireCodeGrant field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetBotRequireCodeGrant() bool {
	if o == nil || IsNil(o.BotRequireCodeGrant) {
		var ret bool
		return ret
	}
	return *o.BotRequireCodeGrant
}

// GetBotRequireCodeGrantOk returns a tuple with the BotRequireCodeGrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetBotRequireCodeGrantOk() (*bool, bool) {
	if o == nil || IsNil(o.BotRequireCodeGrant) {
		return nil, false
	}
	return o.BotRequireCodeGrant, true
}

// HasBotRequireCodeGrant returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasBotRequireCodeGrant() bool {
	if o != nil && !IsNil(o.BotRequireCodeGrant) {
		return true
	}

	return false
}

// SetBotRequireCodeGrant gets a reference to the given bool and assigns it to the BotRequireCodeGrant field.
func (o *InviteApplicationResponse) SetBotRequireCodeGrant(v bool) {
	o.BotRequireCodeGrant = &v
}


// GetTermsOfServiceUrl returns the TermsOfServiceUrl field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetTermsOfServiceUrl() string {
	if o == nil || IsNil(o.TermsOfServiceUrl) {
		var ret string
		return ret
	}
	return *o.TermsOfServiceUrl
}

// GetTermsOfServiceUrlOk returns a tuple with the TermsOfServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetTermsOfServiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TermsOfServiceUrl) {
		return nil, false
	}
	return o.TermsOfServiceUrl, true
}

// HasTermsOfServiceUrl returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasTermsOfServiceUrl() bool {
	if o != nil && !IsNil(o.TermsOfServiceUrl) {
		return true
	}

	return false
}

// SetTermsOfServiceUrl gets a reference to the given string and assigns it to the TermsOfServiceUrl field.
func (o *InviteApplicationResponse) SetTermsOfServiceUrl(v string) {
	o.TermsOfServiceUrl = &v
}


// GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetPrivacyPolicyUrl() string {
	if o == nil || IsNil(o.PrivacyPolicyUrl) {
		var ret string
		return ret
	}
	return *o.PrivacyPolicyUrl
}

// GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetPrivacyPolicyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyPolicyUrl) {
		return nil, false
	}
	return o.PrivacyPolicyUrl, true
}

// HasPrivacyPolicyUrl returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasPrivacyPolicyUrl() bool {
	if o != nil && !IsNil(o.PrivacyPolicyUrl) {
		return true
	}

	return false
}

// SetPrivacyPolicyUrl gets a reference to the given string and assigns it to the PrivacyPolicyUrl field.
func (o *InviteApplicationResponse) SetPrivacyPolicyUrl(v string) {
	o.PrivacyPolicyUrl = &v
}


// GetCustomInstallUrl returns the CustomInstallUrl field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetCustomInstallUrl() string {
	if o == nil || IsNil(o.CustomInstallUrl) {
		var ret string
		return ret
	}
	return *o.CustomInstallUrl
}

// GetCustomInstallUrlOk returns a tuple with the CustomInstallUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetCustomInstallUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomInstallUrl) {
		return nil, false
	}
	return o.CustomInstallUrl, true
}

// HasCustomInstallUrl returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasCustomInstallUrl() bool {
	if o != nil && !IsNil(o.CustomInstallUrl) {
		return true
	}

	return false
}

// SetCustomInstallUrl gets a reference to the given string and assigns it to the CustomInstallUrl field.
func (o *InviteApplicationResponse) SetCustomInstallUrl(v string) {
	o.CustomInstallUrl = &v
}


// GetInstallParams returns the InstallParams field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetInstallParams() ApplicationOAuth2InstallParamsResponse {
	if o == nil || IsNil(o.InstallParams) {
		var ret ApplicationOAuth2InstallParamsResponse
		return ret
	}
	return *o.InstallParams
}

// GetInstallParamsOk returns a tuple with the InstallParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetInstallParamsOk() (*ApplicationOAuth2InstallParamsResponse, bool) {
	if o == nil || IsNil(o.InstallParams) {
		return nil, false
	}
	return o.InstallParams, true
}

// HasInstallParams returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasInstallParams() bool {
	if o != nil && !IsNil(o.InstallParams) {
		return true
	}

	return false
}

// SetInstallParams gets a reference to the given ApplicationOAuth2InstallParamsResponse and assigns it to the InstallParams field.
func (o *InviteApplicationResponse) SetInstallParams(v ApplicationOAuth2InstallParamsResponse) {
	o.InstallParams = &v
}


// GetIntegrationTypesConfig returns the IntegrationTypesConfig field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetIntegrationTypesConfig() map[string]ApplicationIntegrationTypeConfigurationResponse {
	if o == nil || IsNil(o.IntegrationTypesConfig) {
		var ret map[string]ApplicationIntegrationTypeConfigurationResponse
		return ret
	}
	return o.IntegrationTypesConfig
}

// GetIntegrationTypesConfigOk returns a tuple with the IntegrationTypesConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetIntegrationTypesConfigOk() (map[string]ApplicationIntegrationTypeConfigurationResponse, bool) {
	if o == nil || IsNil(o.IntegrationTypesConfig) {
		return map[string]ApplicationIntegrationTypeConfigurationResponse{}, false
	}
	return o.IntegrationTypesConfig, true
}

// HasIntegrationTypesConfig returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasIntegrationTypesConfig() bool {
	if o != nil && !IsNil(o.IntegrationTypesConfig) {
		return true
	}

	return false
}

// SetIntegrationTypesConfig gets a reference to the given map[string]ApplicationIntegrationTypeConfigurationResponse and assigns it to the IntegrationTypesConfig field.
func (o *InviteApplicationResponse) SetIntegrationTypesConfig(v map[string]ApplicationIntegrationTypeConfigurationResponse) {
	o.IntegrationTypesConfig = v
}


// GetVerifyKey returns the VerifyKey field value
func (o *InviteApplicationResponse) GetVerifyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerifyKey
}

// GetVerifyKeyOk returns a tuple with the VerifyKey field value
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetVerifyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifyKey, true
}

// SetVerifyKey sets field value
func (o *InviteApplicationResponse) SetVerifyKey(v string) {
	o.VerifyKey = v
}

// GetFlags returns the Flags field value
func (o *InviteApplicationResponse) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *InviteApplicationResponse) SetFlags(v int32) {
	o.Flags = v
}

// GetMaxParticipants returns the MaxParticipants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InviteApplicationResponse) GetMaxParticipants() int32 {
	if o == nil || IsNil(o.MaxParticipants.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxParticipants.Get()
}

// GetMaxParticipantsOk returns a tuple with the MaxParticipants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InviteApplicationResponse) GetMaxParticipantsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxParticipants.Get()) {
		return nil, false
	}
	return o.MaxParticipants.Get(), o.MaxParticipants.IsSet()
}

// HasMaxParticipants returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasMaxParticipants() bool {
	if o != nil && o.MaxParticipants.IsSet() {
		return true
	}

	return false
}

// SetMaxParticipants gets a reference to the given NullableInt32 and assigns it to the MaxParticipants field.
func (o *InviteApplicationResponse) SetMaxParticipants(v int32) {
	o.MaxParticipants.Set(&v)
}

// SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil
func (o *InviteApplicationResponse) SetMaxParticipantsNil() {
	o.MaxParticipants.Set(nil)
}

// UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
func (o *InviteApplicationResponse) UnsetMaxParticipants() {
	o.MaxParticipants.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InviteApplicationResponse) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteApplicationResponse) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InviteApplicationResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InviteApplicationResponse) SetTags(v []string) {
	o.Tags = v
}


func (o InviteApplicationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteApplicationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	toSerialize["description"] = o.Description
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.CoverImage) {
		toSerialize["cover_image"] = o.CoverImage
	}
	if !IsNil(o.PrimarySkuId) {
		toSerialize["primary_sku_id"] = o.PrimarySkuId
	}
	if !IsNil(o.Bot) {
		toSerialize["bot"] = o.Bot
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.GuildId) {
		toSerialize["guild_id"] = o.GuildId
	}
	if !IsNil(o.RpcOrigins) {
		toSerialize["rpc_origins"] = o.RpcOrigins
	}
	if !IsNil(o.BotPublic) {
		toSerialize["bot_public"] = o.BotPublic
	}
	if !IsNil(o.BotRequireCodeGrant) {
		toSerialize["bot_require_code_grant"] = o.BotRequireCodeGrant
	}
	if !IsNil(o.TermsOfServiceUrl) {
		toSerialize["terms_of_service_url"] = o.TermsOfServiceUrl
	}
	if !IsNil(o.PrivacyPolicyUrl) {
		toSerialize["privacy_policy_url"] = o.PrivacyPolicyUrl
	}
	if !IsNil(o.CustomInstallUrl) {
		toSerialize["custom_install_url"] = o.CustomInstallUrl
	}
	if !IsNil(o.InstallParams) {
		toSerialize["install_params"] = o.InstallParams
	}
	if !IsNil(o.IntegrationTypesConfig) {
		toSerialize["integration_types_config"] = o.IntegrationTypesConfig
	}
	toSerialize["verify_key"] = o.VerifyKey
	toSerialize["flags"] = o.Flags
	if o.MaxParticipants.IsSet() {
		toSerialize["max_participants"] = o.MaxParticipants.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *InviteApplicationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"verify_key",
		"flags",
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

	varInviteApplicationResponse := _InviteApplicationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInviteApplicationResponse)

	if err != nil {
		return err
	}

	*o = InviteApplicationResponse(varInviteApplicationResponse)

	return err
}

type NullableInviteApplicationResponse struct {
	value *InviteApplicationResponse
	isSet bool
}

func (v NullableInviteApplicationResponse) Get() *InviteApplicationResponse {
	return v.value
}

func (v *NullableInviteApplicationResponse) Set(val *InviteApplicationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteApplicationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteApplicationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteApplicationResponse(val *InviteApplicationResponse) *NullableInviteApplicationResponse {
	return &NullableInviteApplicationResponse{value: val, isSet: true}
}

func (v NullableInviteApplicationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteApplicationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


