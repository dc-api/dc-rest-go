/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-05T02:42:25.742582151Z[Etc/UTC]
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

// checks if the PrivateApplicationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateApplicationResponse{}

// PrivateApplicationResponse struct for PrivateApplicationResponse
type PrivateApplicationResponse struct {
	Id string `json:"id" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Name string `json:"name"`
	Icon NullableString `json:"icon,omitempty"`
	Description string `json:"description"`
	Type NullableInt32 `json:"type,omitempty"`
	CoverImage NullableString `json:"cover_image,omitempty"`
	PrimarySkuId *string `json:"primary_sku_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	Bot NullableUserResponse `json:"bot,omitempty"`
	Slug NullableString `json:"slug,omitempty"`
	GuildId *string `json:"guild_id,omitempty" validate:"regexp=^(0|[1-9][0-9]*)$"`
	RpcOrigins []*string `json:"rpc_origins,omitempty"`
	BotPublic NullableBool `json:"bot_public,omitempty"`
	BotRequireCodeGrant NullableBool `json:"bot_require_code_grant,omitempty"`
	TermsOfServiceUrl NullableString `json:"terms_of_service_url,omitempty"`
	PrivacyPolicyUrl NullableString `json:"privacy_policy_url,omitempty"`
	CustomInstallUrl NullableString `json:"custom_install_url,omitempty"`
	InstallParams NullableApplicationOAuth2InstallParamsResponse `json:"install_params,omitempty"`
	IntegrationTypesConfig map[string]ApplicationIntegrationTypeConfigurationResponse `json:"integration_types_config,omitempty"`
	VerifyKey string `json:"verify_key"`
	Flags int32 `json:"flags"`
	MaxParticipants NullableInt32 `json:"max_participants,omitempty"`
	Tags []string `json:"tags,omitempty"`
	RedirectUris []*string `json:"redirect_uris"`
	InteractionsEndpointUrl NullableString `json:"interactions_endpoint_url,omitempty"`
	RoleConnectionsVerificationUrl NullableString `json:"role_connections_verification_url,omitempty"`
	Owner UserResponse `json:"owner"`
	ApproximateGuildCount NullableInt32 `json:"approximate_guild_count,omitempty"`
	ApproximateUserInstallCount int32 `json:"approximate_user_install_count"`
	ApproximateUserAuthorizationCount int32 `json:"approximate_user_authorization_count"`
	ExplicitContentFilter int32 `json:"explicit_content_filter"`
	Team NullableTeamResponse `json:"team,omitempty"`
}

type _PrivateApplicationResponse PrivateApplicationResponse

// NewPrivateApplicationResponse instantiates a new PrivateApplicationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateApplicationResponse(id string, name string, description string, verifyKey string, flags int32, redirectUris []*string, owner UserResponse, approximateUserInstallCount int32, approximateUserAuthorizationCount int32, explicitContentFilter int32) *PrivateApplicationResponse {
	this := PrivateApplicationResponse{}
	this.Id = id
	this.Name = name
	this.Description = description
	this.VerifyKey = verifyKey
	this.Flags = flags
	this.RedirectUris = redirectUris
	this.Owner = owner
	this.ApproximateUserInstallCount = approximateUserInstallCount
	this.ApproximateUserAuthorizationCount = approximateUserAuthorizationCount
	this.ExplicitContentFilter = explicitContentFilter
	return &this
}

// NewPrivateApplicationResponseWithDefaults instantiates a new PrivateApplicationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateApplicationResponseWithDefaults() *PrivateApplicationResponse {
	this := PrivateApplicationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *PrivateApplicationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrivateApplicationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrivateApplicationResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PrivateApplicationResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrivateApplicationResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrivateApplicationResponse) SetName(v string) {
	o.Name = v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon.Get()) {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *PrivateApplicationResponse) SetIcon(v string) {
	o.Icon.Set(&v)
}

// SetIconNil sets the value for Icon to be an explicit nil
func (o *PrivateApplicationResponse) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetIcon() {
	o.Icon.Unset()
}

// GetDescription returns the Description field value
func (o *PrivateApplicationResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PrivateApplicationResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PrivateApplicationResponse) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetType() int32 {
	if o == nil || IsNil(o.Type.Get()) {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type.Get()) {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableInt32 and assigns it to the Type field.
func (o *PrivateApplicationResponse) SetType(v int32) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *PrivateApplicationResponse) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetType() {
	o.Type.Unset()
}

// GetCoverImage returns the CoverImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImage.Get()
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetCoverImageOk() (*string, bool) {
	if o == nil || IsNil(o.CoverImage.Get()) {
		return nil, false
	}
	return o.CoverImage.Get(), o.CoverImage.IsSet()
}

// HasCoverImage returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasCoverImage() bool {
	if o != nil && o.CoverImage.IsSet() {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given NullableString and assigns it to the CoverImage field.
func (o *PrivateApplicationResponse) SetCoverImage(v string) {
	o.CoverImage.Set(&v)
}

// SetCoverImageNil sets the value for CoverImage to be an explicit nil
func (o *PrivateApplicationResponse) SetCoverImageNil() {
	o.CoverImage.Set(nil)
}

// UnsetCoverImage ensures that no value is present for CoverImage, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetCoverImage() {
	o.CoverImage.Unset()
}

// GetPrimarySkuId returns the PrimarySkuId field value if set, zero value otherwise.
func (o *PrivateApplicationResponse) GetPrimarySkuId() string {
	if o == nil || IsNil(o.PrimarySkuId) {
		var ret string
		return ret
	}
	return *o.PrimarySkuId
}

// GetPrimarySkuIdOk returns a tuple with the PrimarySkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateApplicationResponse) GetPrimarySkuIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrimarySkuId) {
		return nil, false
	}
	return o.PrimarySkuId, true
}

// HasPrimarySkuId returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasPrimarySkuId() bool {
	if o != nil && !IsNil(o.PrimarySkuId) {
		return true
	}

	return false
}

// SetPrimarySkuId gets a reference to the given string and assigns it to the PrimarySkuId field.
func (o *PrivateApplicationResponse) SetPrimarySkuId(v string) {
	o.PrimarySkuId = &v
}


// GetBot returns the Bot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetBot() UserResponse {
	if o == nil || IsNil(o.Bot.Get()) {
		var ret UserResponse
		return ret
	}
	return *o.Bot.Get()
}

// GetBotOk returns a tuple with the Bot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetBotOk() (*UserResponse, bool) {
	if o == nil || IsNil(o.Bot.Get()) {
		return nil, false
	}
	return o.Bot.Get(), o.Bot.IsSet()
}

// HasBot returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasBot() bool {
	if o != nil && o.Bot.IsSet() {
		return true
	}

	return false
}

// SetBot gets a reference to the given NullableUserResponse and assigns it to the Bot field.
func (o *PrivateApplicationResponse) SetBot(v UserResponse) {
	o.Bot.Set(&v)
}

// SetBotNil sets the value for Bot to be an explicit nil
func (o *PrivateApplicationResponse) SetBotNil() {
	o.Bot.Set(nil)
}

// UnsetBot ensures that no value is present for Bot, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetBot() {
	o.Bot.Unset()
}

// GetSlug returns the Slug field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetSlug() string {
	if o == nil || IsNil(o.Slug.Get()) {
		var ret string
		return ret
	}
	return *o.Slug.Get()
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug.Get()) {
		return nil, false
	}
	return o.Slug.Get(), o.Slug.IsSet()
}

// HasSlug returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasSlug() bool {
	if o != nil && o.Slug.IsSet() {
		return true
	}

	return false
}

// SetSlug gets a reference to the given NullableString and assigns it to the Slug field.
func (o *PrivateApplicationResponse) SetSlug(v string) {
	o.Slug.Set(&v)
}

// SetSlugNil sets the value for Slug to be an explicit nil
func (o *PrivateApplicationResponse) SetSlugNil() {
	o.Slug.Set(nil)
}

// UnsetSlug ensures that no value is present for Slug, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetSlug() {
	o.Slug.Unset()
}

// GetGuildId returns the GuildId field value if set, zero value otherwise.
func (o *PrivateApplicationResponse) GetGuildId() string {
	if o == nil || IsNil(o.GuildId) {
		var ret string
		return ret
	}
	return *o.GuildId
}

// GetGuildIdOk returns a tuple with the GuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateApplicationResponse) GetGuildIdOk() (*string, bool) {
	if o == nil || IsNil(o.GuildId) {
		return nil, false
	}
	return o.GuildId, true
}

// HasGuildId returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasGuildId() bool {
	if o != nil && !IsNil(o.GuildId) {
		return true
	}

	return false
}

// SetGuildId gets a reference to the given string and assigns it to the GuildId field.
func (o *PrivateApplicationResponse) SetGuildId(v string) {
	o.GuildId = &v
}


// GetRpcOrigins returns the RpcOrigins field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetRpcOrigins() []*string {
	if o == nil {
		var ret []*string
		return ret
	}
	return o.RpcOrigins
}

// GetRpcOriginsOk returns a tuple with the RpcOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetRpcOriginsOk() ([]*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RpcOrigins, true
}

// HasRpcOrigins returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasRpcOrigins() bool {
	if o != nil && !IsNil(o.RpcOrigins) {
		return true
	}

	return false
}

// SetRpcOrigins gets a reference to the given []*string and assigns it to the RpcOrigins field.
func (o *PrivateApplicationResponse) SetRpcOrigins(v []*string) {
	o.RpcOrigins = v
}


// GetBotPublic returns the BotPublic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetBotPublic() bool {
	if o == nil || IsNil(o.BotPublic.Get()) {
		var ret bool
		return ret
	}
	return *o.BotPublic.Get()
}

// GetBotPublicOk returns a tuple with the BotPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetBotPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.BotPublic.Get()) {
		return nil, false
	}
	return o.BotPublic.Get(), o.BotPublic.IsSet()
}

// HasBotPublic returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasBotPublic() bool {
	if o != nil && o.BotPublic.IsSet() {
		return true
	}

	return false
}

// SetBotPublic gets a reference to the given NullableBool and assigns it to the BotPublic field.
func (o *PrivateApplicationResponse) SetBotPublic(v bool) {
	o.BotPublic.Set(&v)
}

// SetBotPublicNil sets the value for BotPublic to be an explicit nil
func (o *PrivateApplicationResponse) SetBotPublicNil() {
	o.BotPublic.Set(nil)
}

// UnsetBotPublic ensures that no value is present for BotPublic, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetBotPublic() {
	o.BotPublic.Unset()
}

// GetBotRequireCodeGrant returns the BotRequireCodeGrant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetBotRequireCodeGrant() bool {
	if o == nil || IsNil(o.BotRequireCodeGrant.Get()) {
		var ret bool
		return ret
	}
	return *o.BotRequireCodeGrant.Get()
}

// GetBotRequireCodeGrantOk returns a tuple with the BotRequireCodeGrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetBotRequireCodeGrantOk() (*bool, bool) {
	if o == nil || IsNil(o.BotRequireCodeGrant.Get()) {
		return nil, false
	}
	return o.BotRequireCodeGrant.Get(), o.BotRequireCodeGrant.IsSet()
}

// HasBotRequireCodeGrant returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasBotRequireCodeGrant() bool {
	if o != nil && o.BotRequireCodeGrant.IsSet() {
		return true
	}

	return false
}

// SetBotRequireCodeGrant gets a reference to the given NullableBool and assigns it to the BotRequireCodeGrant field.
func (o *PrivateApplicationResponse) SetBotRequireCodeGrant(v bool) {
	o.BotRequireCodeGrant.Set(&v)
}

// SetBotRequireCodeGrantNil sets the value for BotRequireCodeGrant to be an explicit nil
func (o *PrivateApplicationResponse) SetBotRequireCodeGrantNil() {
	o.BotRequireCodeGrant.Set(nil)
}

// UnsetBotRequireCodeGrant ensures that no value is present for BotRequireCodeGrant, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetBotRequireCodeGrant() {
	o.BotRequireCodeGrant.Unset()
}

// GetTermsOfServiceUrl returns the TermsOfServiceUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetTermsOfServiceUrl() string {
	if o == nil || IsNil(o.TermsOfServiceUrl.Get()) {
		var ret string
		return ret
	}
	return *o.TermsOfServiceUrl.Get()
}

// GetTermsOfServiceUrlOk returns a tuple with the TermsOfServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetTermsOfServiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TermsOfServiceUrl.Get()) {
		return nil, false
	}
	return o.TermsOfServiceUrl.Get(), o.TermsOfServiceUrl.IsSet()
}

// HasTermsOfServiceUrl returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasTermsOfServiceUrl() bool {
	if o != nil && o.TermsOfServiceUrl.IsSet() {
		return true
	}

	return false
}

// SetTermsOfServiceUrl gets a reference to the given NullableString and assigns it to the TermsOfServiceUrl field.
func (o *PrivateApplicationResponse) SetTermsOfServiceUrl(v string) {
	o.TermsOfServiceUrl.Set(&v)
}

// SetTermsOfServiceUrlNil sets the value for TermsOfServiceUrl to be an explicit nil
func (o *PrivateApplicationResponse) SetTermsOfServiceUrlNil() {
	o.TermsOfServiceUrl.Set(nil)
}

// UnsetTermsOfServiceUrl ensures that no value is present for TermsOfServiceUrl, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetTermsOfServiceUrl() {
	o.TermsOfServiceUrl.Unset()
}

// GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetPrivacyPolicyUrl() string {
	if o == nil || IsNil(o.PrivacyPolicyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PrivacyPolicyUrl.Get()
}

// GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetPrivacyPolicyUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PrivacyPolicyUrl.Get()) {
		return nil, false
	}
	return o.PrivacyPolicyUrl.Get(), o.PrivacyPolicyUrl.IsSet()
}

// HasPrivacyPolicyUrl returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasPrivacyPolicyUrl() bool {
	if o != nil && o.PrivacyPolicyUrl.IsSet() {
		return true
	}

	return false
}

// SetPrivacyPolicyUrl gets a reference to the given NullableString and assigns it to the PrivacyPolicyUrl field.
func (o *PrivateApplicationResponse) SetPrivacyPolicyUrl(v string) {
	o.PrivacyPolicyUrl.Set(&v)
}

// SetPrivacyPolicyUrlNil sets the value for PrivacyPolicyUrl to be an explicit nil
func (o *PrivateApplicationResponse) SetPrivacyPolicyUrlNil() {
	o.PrivacyPolicyUrl.Set(nil)
}

// UnsetPrivacyPolicyUrl ensures that no value is present for PrivacyPolicyUrl, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetPrivacyPolicyUrl() {
	o.PrivacyPolicyUrl.Unset()
}

// GetCustomInstallUrl returns the CustomInstallUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetCustomInstallUrl() string {
	if o == nil || IsNil(o.CustomInstallUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CustomInstallUrl.Get()
}

// GetCustomInstallUrlOk returns a tuple with the CustomInstallUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetCustomInstallUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CustomInstallUrl.Get()) {
		return nil, false
	}
	return o.CustomInstallUrl.Get(), o.CustomInstallUrl.IsSet()
}

// HasCustomInstallUrl returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasCustomInstallUrl() bool {
	if o != nil && o.CustomInstallUrl.IsSet() {
		return true
	}

	return false
}

// SetCustomInstallUrl gets a reference to the given NullableString and assigns it to the CustomInstallUrl field.
func (o *PrivateApplicationResponse) SetCustomInstallUrl(v string) {
	o.CustomInstallUrl.Set(&v)
}

// SetCustomInstallUrlNil sets the value for CustomInstallUrl to be an explicit nil
func (o *PrivateApplicationResponse) SetCustomInstallUrlNil() {
	o.CustomInstallUrl.Set(nil)
}

// UnsetCustomInstallUrl ensures that no value is present for CustomInstallUrl, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetCustomInstallUrl() {
	o.CustomInstallUrl.Unset()
}

// GetInstallParams returns the InstallParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetInstallParams() ApplicationOAuth2InstallParamsResponse {
	if o == nil || IsNil(o.InstallParams.Get()) {
		var ret ApplicationOAuth2InstallParamsResponse
		return ret
	}
	return *o.InstallParams.Get()
}

// GetInstallParamsOk returns a tuple with the InstallParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetInstallParamsOk() (*ApplicationOAuth2InstallParamsResponse, bool) {
	if o == nil || IsNil(o.InstallParams.Get()) {
		return nil, false
	}
	return o.InstallParams.Get(), o.InstallParams.IsSet()
}

// HasInstallParams returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasInstallParams() bool {
	if o != nil && o.InstallParams.IsSet() {
		return true
	}

	return false
}

// SetInstallParams gets a reference to the given NullableApplicationOAuth2InstallParamsResponse and assigns it to the InstallParams field.
func (o *PrivateApplicationResponse) SetInstallParams(v ApplicationOAuth2InstallParamsResponse) {
	o.InstallParams.Set(&v)
}

// SetInstallParamsNil sets the value for InstallParams to be an explicit nil
func (o *PrivateApplicationResponse) SetInstallParamsNil() {
	o.InstallParams.Set(nil)
}

// UnsetInstallParams ensures that no value is present for InstallParams, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetInstallParams() {
	o.InstallParams.Unset()
}

// GetIntegrationTypesConfig returns the IntegrationTypesConfig field value if set, zero value otherwise.
func (o *PrivateApplicationResponse) GetIntegrationTypesConfig() map[string]ApplicationIntegrationTypeConfigurationResponse {
	if o == nil || IsNil(o.IntegrationTypesConfig) {
		var ret map[string]ApplicationIntegrationTypeConfigurationResponse
		return ret
	}
	return o.IntegrationTypesConfig
}

// GetIntegrationTypesConfigOk returns a tuple with the IntegrationTypesConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateApplicationResponse) GetIntegrationTypesConfigOk() (map[string]ApplicationIntegrationTypeConfigurationResponse, bool) {
	if o == nil || IsNil(o.IntegrationTypesConfig) {
		return map[string]ApplicationIntegrationTypeConfigurationResponse{}, false
	}
	return o.IntegrationTypesConfig, true
}

// HasIntegrationTypesConfig returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasIntegrationTypesConfig() bool {
	if o != nil && !IsNil(o.IntegrationTypesConfig) {
		return true
	}

	return false
}

// SetIntegrationTypesConfig gets a reference to the given map[string]ApplicationIntegrationTypeConfigurationResponse and assigns it to the IntegrationTypesConfig field.
func (o *PrivateApplicationResponse) SetIntegrationTypesConfig(v map[string]ApplicationIntegrationTypeConfigurationResponse) {
	o.IntegrationTypesConfig = v
}


// GetVerifyKey returns the VerifyKey field value
func (o *PrivateApplicationResponse) GetVerifyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerifyKey
}

// GetVerifyKeyOk returns a tuple with the VerifyKey field value
// and a boolean to check if the value has been set.
func (o *PrivateApplicationResponse) GetVerifyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifyKey, true
}

// SetVerifyKey sets field value
func (o *PrivateApplicationResponse) SetVerifyKey(v string) {
	o.VerifyKey = v
}

// GetFlags returns the Flags field value
func (o *PrivateApplicationResponse) GetFlags() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value
// and a boolean to check if the value has been set.
func (o *PrivateApplicationResponse) GetFlagsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Flags, true
}

// SetFlags sets field value
func (o *PrivateApplicationResponse) SetFlags(v int32) {
	o.Flags = v
}

// GetMaxParticipants returns the MaxParticipants field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetMaxParticipants() int32 {
	if o == nil || IsNil(o.MaxParticipants.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxParticipants.Get()
}

// GetMaxParticipantsOk returns a tuple with the MaxParticipants field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetMaxParticipantsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxParticipants.Get()) {
		return nil, false
	}
	return o.MaxParticipants.Get(), o.MaxParticipants.IsSet()
}

// HasMaxParticipants returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasMaxParticipants() bool {
	if o != nil && o.MaxParticipants.IsSet() {
		return true
	}

	return false
}

// SetMaxParticipants gets a reference to the given NullableInt32 and assigns it to the MaxParticipants field.
func (o *PrivateApplicationResponse) SetMaxParticipants(v int32) {
	o.MaxParticipants.Set(&v)
}

// SetMaxParticipantsNil sets the value for MaxParticipants to be an explicit nil
func (o *PrivateApplicationResponse) SetMaxParticipantsNil() {
	o.MaxParticipants.Set(nil)
}

// UnsetMaxParticipants ensures that no value is present for MaxParticipants, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetMaxParticipants() {
	o.MaxParticipants.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PrivateApplicationResponse) SetTags(v []string) {
	o.Tags = v
}


// GetRedirectUris returns the RedirectUris field value
func (o *PrivateApplicationResponse) GetRedirectUris() []*string {
	if o == nil {
		var ret []*string
		return ret
	}

	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value
// and a boolean to check if the value has been set.
func (o *PrivateApplicationResponse) GetRedirectUrisOk() ([]*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectUris, true
}

// SetRedirectUris sets field value
func (o *PrivateApplicationResponse) SetRedirectUris(v []*string) {
	o.RedirectUris = v
}

// GetInteractionsEndpointUrl returns the InteractionsEndpointUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetInteractionsEndpointUrl() string {
	if o == nil || IsNil(o.InteractionsEndpointUrl.Get()) {
		var ret string
		return ret
	}
	return *o.InteractionsEndpointUrl.Get()
}

// GetInteractionsEndpointUrlOk returns a tuple with the InteractionsEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetInteractionsEndpointUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InteractionsEndpointUrl.Get()) {
		return nil, false
	}
	return o.InteractionsEndpointUrl.Get(), o.InteractionsEndpointUrl.IsSet()
}

// HasInteractionsEndpointUrl returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasInteractionsEndpointUrl() bool {
	if o != nil && o.InteractionsEndpointUrl.IsSet() {
		return true
	}

	return false
}

// SetInteractionsEndpointUrl gets a reference to the given NullableString and assigns it to the InteractionsEndpointUrl field.
func (o *PrivateApplicationResponse) SetInteractionsEndpointUrl(v string) {
	o.InteractionsEndpointUrl.Set(&v)
}

// SetInteractionsEndpointUrlNil sets the value for InteractionsEndpointUrl to be an explicit nil
func (o *PrivateApplicationResponse) SetInteractionsEndpointUrlNil() {
	o.InteractionsEndpointUrl.Set(nil)
}

// UnsetInteractionsEndpointUrl ensures that no value is present for InteractionsEndpointUrl, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetInteractionsEndpointUrl() {
	o.InteractionsEndpointUrl.Unset()
}

// GetRoleConnectionsVerificationUrl returns the RoleConnectionsVerificationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetRoleConnectionsVerificationUrl() string {
	if o == nil || IsNil(o.RoleConnectionsVerificationUrl.Get()) {
		var ret string
		return ret
	}
	return *o.RoleConnectionsVerificationUrl.Get()
}

// GetRoleConnectionsVerificationUrlOk returns a tuple with the RoleConnectionsVerificationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetRoleConnectionsVerificationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RoleConnectionsVerificationUrl.Get()) {
		return nil, false
	}
	return o.RoleConnectionsVerificationUrl.Get(), o.RoleConnectionsVerificationUrl.IsSet()
}

// HasRoleConnectionsVerificationUrl returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasRoleConnectionsVerificationUrl() bool {
	if o != nil && o.RoleConnectionsVerificationUrl.IsSet() {
		return true
	}

	return false
}

// SetRoleConnectionsVerificationUrl gets a reference to the given NullableString and assigns it to the RoleConnectionsVerificationUrl field.
func (o *PrivateApplicationResponse) SetRoleConnectionsVerificationUrl(v string) {
	o.RoleConnectionsVerificationUrl.Set(&v)
}

// SetRoleConnectionsVerificationUrlNil sets the value for RoleConnectionsVerificationUrl to be an explicit nil
func (o *PrivateApplicationResponse) SetRoleConnectionsVerificationUrlNil() {
	o.RoleConnectionsVerificationUrl.Set(nil)
}

// UnsetRoleConnectionsVerificationUrl ensures that no value is present for RoleConnectionsVerificationUrl, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetRoleConnectionsVerificationUrl() {
	o.RoleConnectionsVerificationUrl.Unset()
}

// GetOwner returns the Owner field value
func (o *PrivateApplicationResponse) GetOwner() UserResponse {
	if o == nil {
		var ret UserResponse
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *PrivateApplicationResponse) GetOwnerOk() (*UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *PrivateApplicationResponse) SetOwner(v UserResponse) {
	o.Owner = v
}

// GetApproximateGuildCount returns the ApproximateGuildCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetApproximateGuildCount() int32 {
	if o == nil || IsNil(o.ApproximateGuildCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ApproximateGuildCount.Get()
}

// GetApproximateGuildCountOk returns a tuple with the ApproximateGuildCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetApproximateGuildCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ApproximateGuildCount.Get()) {
		return nil, false
	}
	return o.ApproximateGuildCount.Get(), o.ApproximateGuildCount.IsSet()
}

// HasApproximateGuildCount returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasApproximateGuildCount() bool {
	if o != nil && o.ApproximateGuildCount.IsSet() {
		return true
	}

	return false
}

// SetApproximateGuildCount gets a reference to the given NullableInt32 and assigns it to the ApproximateGuildCount field.
func (o *PrivateApplicationResponse) SetApproximateGuildCount(v int32) {
	o.ApproximateGuildCount.Set(&v)
}

// SetApproximateGuildCountNil sets the value for ApproximateGuildCount to be an explicit nil
func (o *PrivateApplicationResponse) SetApproximateGuildCountNil() {
	o.ApproximateGuildCount.Set(nil)
}

// UnsetApproximateGuildCount ensures that no value is present for ApproximateGuildCount, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetApproximateGuildCount() {
	o.ApproximateGuildCount.Unset()
}

// GetApproximateUserInstallCount returns the ApproximateUserInstallCount field value
func (o *PrivateApplicationResponse) GetApproximateUserInstallCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApproximateUserInstallCount
}

// GetApproximateUserInstallCountOk returns a tuple with the ApproximateUserInstallCount field value
// and a boolean to check if the value has been set.
func (o *PrivateApplicationResponse) GetApproximateUserInstallCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproximateUserInstallCount, true
}

// SetApproximateUserInstallCount sets field value
func (o *PrivateApplicationResponse) SetApproximateUserInstallCount(v int32) {
	o.ApproximateUserInstallCount = v
}

// GetApproximateUserAuthorizationCount returns the ApproximateUserAuthorizationCount field value
func (o *PrivateApplicationResponse) GetApproximateUserAuthorizationCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApproximateUserAuthorizationCount
}

// GetApproximateUserAuthorizationCountOk returns a tuple with the ApproximateUserAuthorizationCount field value
// and a boolean to check if the value has been set.
func (o *PrivateApplicationResponse) GetApproximateUserAuthorizationCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproximateUserAuthorizationCount, true
}

// SetApproximateUserAuthorizationCount sets field value
func (o *PrivateApplicationResponse) SetApproximateUserAuthorizationCount(v int32) {
	o.ApproximateUserAuthorizationCount = v
}

// GetExplicitContentFilter returns the ExplicitContentFilter field value
func (o *PrivateApplicationResponse) GetExplicitContentFilter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExplicitContentFilter
}

// GetExplicitContentFilterOk returns a tuple with the ExplicitContentFilter field value
// and a boolean to check if the value has been set.
func (o *PrivateApplicationResponse) GetExplicitContentFilterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExplicitContentFilter, true
}

// SetExplicitContentFilter sets field value
func (o *PrivateApplicationResponse) SetExplicitContentFilter(v int32) {
	o.ExplicitContentFilter = v
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrivateApplicationResponse) GetTeam() TeamResponse {
	if o == nil || IsNil(o.Team.Get()) {
		var ret TeamResponse
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrivateApplicationResponse) GetTeamOk() (*TeamResponse, bool) {
	if o == nil || IsNil(o.Team.Get()) {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *PrivateApplicationResponse) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableTeamResponse and assigns it to the Team field.
func (o *PrivateApplicationResponse) SetTeam(v TeamResponse) {
	o.Team.Set(&v)
}

// SetTeamNil sets the value for Team to be an explicit nil
func (o *PrivateApplicationResponse) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *PrivateApplicationResponse) UnsetTeam() {
	o.Team.Unset()
}

func (o PrivateApplicationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateApplicationResponse) ToMap() (map[string]interface{}, error) {
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
	if o.CoverImage.IsSet() {
		toSerialize["cover_image"] = o.CoverImage.Get()
	}
	if !IsNil(o.PrimarySkuId) {
		toSerialize["primary_sku_id"] = o.PrimarySkuId
	}
	if o.Bot.IsSet() {
		toSerialize["bot"] = o.Bot.Get()
	}
	if o.Slug.IsSet() {
		toSerialize["slug"] = o.Slug.Get()
	}
	if !IsNil(o.GuildId) {
		toSerialize["guild_id"] = o.GuildId
	}
	if o.RpcOrigins != nil {
		toSerialize["rpc_origins"] = o.RpcOrigins
	}
	if o.BotPublic.IsSet() {
		toSerialize["bot_public"] = o.BotPublic.Get()
	}
	if o.BotRequireCodeGrant.IsSet() {
		toSerialize["bot_require_code_grant"] = o.BotRequireCodeGrant.Get()
	}
	if o.TermsOfServiceUrl.IsSet() {
		toSerialize["terms_of_service_url"] = o.TermsOfServiceUrl.Get()
	}
	if o.PrivacyPolicyUrl.IsSet() {
		toSerialize["privacy_policy_url"] = o.PrivacyPolicyUrl.Get()
	}
	if o.CustomInstallUrl.IsSet() {
		toSerialize["custom_install_url"] = o.CustomInstallUrl.Get()
	}
	if o.InstallParams.IsSet() {
		toSerialize["install_params"] = o.InstallParams.Get()
	}
	if !IsNil(o.IntegrationTypesConfig) {
		toSerialize["integration_types_config"] = o.IntegrationTypesConfig
	}
	toSerialize["verify_key"] = o.VerifyKey
	toSerialize["flags"] = o.Flags
	if o.MaxParticipants.IsSet() {
		toSerialize["max_participants"] = o.MaxParticipants.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["redirect_uris"] = o.RedirectUris
	if o.InteractionsEndpointUrl.IsSet() {
		toSerialize["interactions_endpoint_url"] = o.InteractionsEndpointUrl.Get()
	}
	if o.RoleConnectionsVerificationUrl.IsSet() {
		toSerialize["role_connections_verification_url"] = o.RoleConnectionsVerificationUrl.Get()
	}
	toSerialize["owner"] = o.Owner
	if o.ApproximateGuildCount.IsSet() {
		toSerialize["approximate_guild_count"] = o.ApproximateGuildCount.Get()
	}
	toSerialize["approximate_user_install_count"] = o.ApproximateUserInstallCount
	toSerialize["approximate_user_authorization_count"] = o.ApproximateUserAuthorizationCount
	toSerialize["explicit_content_filter"] = o.ExplicitContentFilter
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}
	return toSerialize, nil
}

func (o *PrivateApplicationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"description",
		"verify_key",
		"flags",
		"redirect_uris",
		"owner",
		"approximate_user_install_count",
		"approximate_user_authorization_count",
		"explicit_content_filter",
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

	varPrivateApplicationResponse := _PrivateApplicationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPrivateApplicationResponse)

	if err != nil {
		return err
	}

	*o = PrivateApplicationResponse(varPrivateApplicationResponse)

	return err
}

type NullablePrivateApplicationResponse struct {
	value *PrivateApplicationResponse
	isSet bool
}

func (v NullablePrivateApplicationResponse) Get() *PrivateApplicationResponse {
	return v.value
}

func (v *NullablePrivateApplicationResponse) Set(val *PrivateApplicationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateApplicationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateApplicationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateApplicationResponse(val *PrivateApplicationResponse) *NullablePrivateApplicationResponse {
	return &NullablePrivateApplicationResponse{value: val, isSet: true}
}

func (v NullablePrivateApplicationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateApplicationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


