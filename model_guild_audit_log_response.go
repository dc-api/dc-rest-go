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

// checks if the GuildAuditLogResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GuildAuditLogResponse{}

// GuildAuditLogResponse struct for GuildAuditLogResponse
type GuildAuditLogResponse struct {
	AuditLogEntries []AuditLogEntryResponse `json:"audit_log_entries"`
	Users []UserResponse `json:"users"`
	Integrations []GuildAuditLogResponseIntegrationsInner `json:"integrations"`
	Webhooks []ListChannelWebhooks200ResponseInner `json:"webhooks"`
	GuildScheduledEvents []ListGuildScheduledEvents200ResponseInner `json:"guild_scheduled_events"`
	Threads []ThreadResponse `json:"threads"`
	ApplicationCommands []ApplicationCommandResponse `json:"application_commands"`
	AutoModerationRules []ListAutoModerationRules200ResponseInner `json:"auto_moderation_rules"`
}

type _GuildAuditLogResponse GuildAuditLogResponse

// NewGuildAuditLogResponse instantiates a new GuildAuditLogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuildAuditLogResponse(auditLogEntries []AuditLogEntryResponse, users []UserResponse, integrations []GuildAuditLogResponseIntegrationsInner, webhooks []ListChannelWebhooks200ResponseInner, guildScheduledEvents []ListGuildScheduledEvents200ResponseInner, threads []ThreadResponse, applicationCommands []ApplicationCommandResponse, autoModerationRules []ListAutoModerationRules200ResponseInner) *GuildAuditLogResponse {
	this := GuildAuditLogResponse{}
	this.AuditLogEntries = auditLogEntries
	this.Users = users
	this.Integrations = integrations
	this.Webhooks = webhooks
	this.GuildScheduledEvents = guildScheduledEvents
	this.Threads = threads
	this.ApplicationCommands = applicationCommands
	this.AutoModerationRules = autoModerationRules
	return &this
}

// NewGuildAuditLogResponseWithDefaults instantiates a new GuildAuditLogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuildAuditLogResponseWithDefaults() *GuildAuditLogResponse {
	this := GuildAuditLogResponse{}
	return &this
}

// GetAuditLogEntries returns the AuditLogEntries field value
func (o *GuildAuditLogResponse) GetAuditLogEntries() []AuditLogEntryResponse {
	if o == nil {
		var ret []AuditLogEntryResponse
		return ret
	}

	return o.AuditLogEntries
}

// GetAuditLogEntriesOk returns a tuple with the AuditLogEntries field value
// and a boolean to check if the value has been set.
func (o *GuildAuditLogResponse) GetAuditLogEntriesOk() ([]AuditLogEntryResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuditLogEntries, true
}

// SetAuditLogEntries sets field value
func (o *GuildAuditLogResponse) SetAuditLogEntries(v []AuditLogEntryResponse) {
	o.AuditLogEntries = v
}

// GetUsers returns the Users field value
func (o *GuildAuditLogResponse) GetUsers() []UserResponse {
	if o == nil {
		var ret []UserResponse
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *GuildAuditLogResponse) GetUsersOk() ([]UserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *GuildAuditLogResponse) SetUsers(v []UserResponse) {
	o.Users = v
}

// GetIntegrations returns the Integrations field value
func (o *GuildAuditLogResponse) GetIntegrations() []GuildAuditLogResponseIntegrationsInner {
	if o == nil {
		var ret []GuildAuditLogResponseIntegrationsInner
		return ret
	}

	return o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value
// and a boolean to check if the value has been set.
func (o *GuildAuditLogResponse) GetIntegrationsOk() ([]GuildAuditLogResponseIntegrationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Integrations, true
}

// SetIntegrations sets field value
func (o *GuildAuditLogResponse) SetIntegrations(v []GuildAuditLogResponseIntegrationsInner) {
	o.Integrations = v
}

// GetWebhooks returns the Webhooks field value
func (o *GuildAuditLogResponse) GetWebhooks() []ListChannelWebhooks200ResponseInner {
	if o == nil {
		var ret []ListChannelWebhooks200ResponseInner
		return ret
	}

	return o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value
// and a boolean to check if the value has been set.
func (o *GuildAuditLogResponse) GetWebhooksOk() ([]ListChannelWebhooks200ResponseInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Webhooks, true
}

// SetWebhooks sets field value
func (o *GuildAuditLogResponse) SetWebhooks(v []ListChannelWebhooks200ResponseInner) {
	o.Webhooks = v
}

// GetGuildScheduledEvents returns the GuildScheduledEvents field value
func (o *GuildAuditLogResponse) GetGuildScheduledEvents() []ListGuildScheduledEvents200ResponseInner {
	if o == nil {
		var ret []ListGuildScheduledEvents200ResponseInner
		return ret
	}

	return o.GuildScheduledEvents
}

// GetGuildScheduledEventsOk returns a tuple with the GuildScheduledEvents field value
// and a boolean to check if the value has been set.
func (o *GuildAuditLogResponse) GetGuildScheduledEventsOk() ([]ListGuildScheduledEvents200ResponseInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.GuildScheduledEvents, true
}

// SetGuildScheduledEvents sets field value
func (o *GuildAuditLogResponse) SetGuildScheduledEvents(v []ListGuildScheduledEvents200ResponseInner) {
	o.GuildScheduledEvents = v
}

// GetThreads returns the Threads field value
func (o *GuildAuditLogResponse) GetThreads() []ThreadResponse {
	if o == nil {
		var ret []ThreadResponse
		return ret
	}

	return o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value
// and a boolean to check if the value has been set.
func (o *GuildAuditLogResponse) GetThreadsOk() ([]ThreadResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Threads, true
}

// SetThreads sets field value
func (o *GuildAuditLogResponse) SetThreads(v []ThreadResponse) {
	o.Threads = v
}

// GetApplicationCommands returns the ApplicationCommands field value
func (o *GuildAuditLogResponse) GetApplicationCommands() []ApplicationCommandResponse {
	if o == nil {
		var ret []ApplicationCommandResponse
		return ret
	}

	return o.ApplicationCommands
}

// GetApplicationCommandsOk returns a tuple with the ApplicationCommands field value
// and a boolean to check if the value has been set.
func (o *GuildAuditLogResponse) GetApplicationCommandsOk() ([]ApplicationCommandResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApplicationCommands, true
}

// SetApplicationCommands sets field value
func (o *GuildAuditLogResponse) SetApplicationCommands(v []ApplicationCommandResponse) {
	o.ApplicationCommands = v
}

// GetAutoModerationRules returns the AutoModerationRules field value
func (o *GuildAuditLogResponse) GetAutoModerationRules() []ListAutoModerationRules200ResponseInner {
	if o == nil {
		var ret []ListAutoModerationRules200ResponseInner
		return ret
	}

	return o.AutoModerationRules
}

// GetAutoModerationRulesOk returns a tuple with the AutoModerationRules field value
// and a boolean to check if the value has been set.
func (o *GuildAuditLogResponse) GetAutoModerationRulesOk() ([]ListAutoModerationRules200ResponseInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoModerationRules, true
}

// SetAutoModerationRules sets field value
func (o *GuildAuditLogResponse) SetAutoModerationRules(v []ListAutoModerationRules200ResponseInner) {
	o.AutoModerationRules = v
}

func (o GuildAuditLogResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GuildAuditLogResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["audit_log_entries"] = o.AuditLogEntries
	toSerialize["users"] = o.Users
	toSerialize["integrations"] = o.Integrations
	toSerialize["webhooks"] = o.Webhooks
	toSerialize["guild_scheduled_events"] = o.GuildScheduledEvents
	toSerialize["threads"] = o.Threads
	toSerialize["application_commands"] = o.ApplicationCommands
	toSerialize["auto_moderation_rules"] = o.AutoModerationRules
	return toSerialize, nil
}

func (o *GuildAuditLogResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"audit_log_entries",
		"users",
		"integrations",
		"webhooks",
		"guild_scheduled_events",
		"threads",
		"application_commands",
		"auto_moderation_rules",
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

	varGuildAuditLogResponse := _GuildAuditLogResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGuildAuditLogResponse)

	if err != nil {
		return err
	}

	*o = GuildAuditLogResponse(varGuildAuditLogResponse)

	return err
}

type NullableGuildAuditLogResponse struct {
	value *GuildAuditLogResponse
	isSet bool
}

func (v NullableGuildAuditLogResponse) Get() *GuildAuditLogResponse {
	return v.value
}

func (v *NullableGuildAuditLogResponse) Set(val *GuildAuditLogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGuildAuditLogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGuildAuditLogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuildAuditLogResponse(val *GuildAuditLogResponse) *NullableGuildAuditLogResponse {
	return &NullableGuildAuditLogResponse{value: val, isSet: true}
}

func (v NullableGuildAuditLogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuildAuditLogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


