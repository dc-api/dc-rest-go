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
	"fmt"
	"gopkg.in/validator.v2"
)

// DefaultKeywordListUpsertRequestActionsInner - struct for DefaultKeywordListUpsertRequestActionsInner
type DefaultKeywordListUpsertRequestActionsInner struct {
	BlockMessageAction *BlockMessageAction
	FlagToChannelAction *FlagToChannelAction
	QuarantineUserAction *QuarantineUserAction
	UserCommunicationDisabledAction *UserCommunicationDisabledAction
}

// BlockMessageActionAsDefaultKeywordListUpsertRequestActionsInner is a convenience function that returns BlockMessageAction wrapped in DefaultKeywordListUpsertRequestActionsInner
func BlockMessageActionAsDefaultKeywordListUpsertRequestActionsInner(v *BlockMessageAction) DefaultKeywordListUpsertRequestActionsInner {
	return DefaultKeywordListUpsertRequestActionsInner{
		BlockMessageAction: v,
	}
}

// FlagToChannelActionAsDefaultKeywordListUpsertRequestActionsInner is a convenience function that returns FlagToChannelAction wrapped in DefaultKeywordListUpsertRequestActionsInner
func FlagToChannelActionAsDefaultKeywordListUpsertRequestActionsInner(v *FlagToChannelAction) DefaultKeywordListUpsertRequestActionsInner {
	return DefaultKeywordListUpsertRequestActionsInner{
		FlagToChannelAction: v,
	}
}

// QuarantineUserActionAsDefaultKeywordListUpsertRequestActionsInner is a convenience function that returns QuarantineUserAction wrapped in DefaultKeywordListUpsertRequestActionsInner
func QuarantineUserActionAsDefaultKeywordListUpsertRequestActionsInner(v *QuarantineUserAction) DefaultKeywordListUpsertRequestActionsInner {
	return DefaultKeywordListUpsertRequestActionsInner{
		QuarantineUserAction: v,
	}
}

// UserCommunicationDisabledActionAsDefaultKeywordListUpsertRequestActionsInner is a convenience function that returns UserCommunicationDisabledAction wrapped in DefaultKeywordListUpsertRequestActionsInner
func UserCommunicationDisabledActionAsDefaultKeywordListUpsertRequestActionsInner(v *UserCommunicationDisabledAction) DefaultKeywordListUpsertRequestActionsInner {
	return DefaultKeywordListUpsertRequestActionsInner{
		UserCommunicationDisabledAction: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DefaultKeywordListUpsertRequestActionsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into BlockMessageAction
	err = newStrictDecoder(data).Decode(&dst.BlockMessageAction)
	if err == nil {
		jsonBlockMessageAction, _ := json.Marshal(dst.BlockMessageAction)
		if string(jsonBlockMessageAction) == "{}" { // empty struct
			dst.BlockMessageAction = nil
		} else {
			if err = validator.Validate(dst.BlockMessageAction); err != nil {
				dst.BlockMessageAction = nil
			} else {
				match++
			}
		}
	} else {
		dst.BlockMessageAction = nil
	}

	// try to unmarshal data into FlagToChannelAction
	err = newStrictDecoder(data).Decode(&dst.FlagToChannelAction)
	if err == nil {
		jsonFlagToChannelAction, _ := json.Marshal(dst.FlagToChannelAction)
		if string(jsonFlagToChannelAction) == "{}" { // empty struct
			dst.FlagToChannelAction = nil
		} else {
			if err = validator.Validate(dst.FlagToChannelAction); err != nil {
				dst.FlagToChannelAction = nil
			} else {
				match++
			}
		}
	} else {
		dst.FlagToChannelAction = nil
	}

	// try to unmarshal data into QuarantineUserAction
	err = newStrictDecoder(data).Decode(&dst.QuarantineUserAction)
	if err == nil {
		jsonQuarantineUserAction, _ := json.Marshal(dst.QuarantineUserAction)
		if string(jsonQuarantineUserAction) == "{}" { // empty struct
			dst.QuarantineUserAction = nil
		} else {
			if err = validator.Validate(dst.QuarantineUserAction); err != nil {
				dst.QuarantineUserAction = nil
			} else {
				match++
			}
		}
	} else {
		dst.QuarantineUserAction = nil
	}

	// try to unmarshal data into UserCommunicationDisabledAction
	err = newStrictDecoder(data).Decode(&dst.UserCommunicationDisabledAction)
	if err == nil {
		jsonUserCommunicationDisabledAction, _ := json.Marshal(dst.UserCommunicationDisabledAction)
		if string(jsonUserCommunicationDisabledAction) == "{}" { // empty struct
			dst.UserCommunicationDisabledAction = nil
		} else {
			if err = validator.Validate(dst.UserCommunicationDisabledAction); err != nil {
				dst.UserCommunicationDisabledAction = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserCommunicationDisabledAction = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.BlockMessageAction = nil
		dst.FlagToChannelAction = nil
		dst.QuarantineUserAction = nil
		dst.UserCommunicationDisabledAction = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DefaultKeywordListUpsertRequestActionsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DefaultKeywordListUpsertRequestActionsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DefaultKeywordListUpsertRequestActionsInner) MarshalJSON() ([]byte, error) {
	if src.BlockMessageAction != nil {
		return json.Marshal(&src.BlockMessageAction)
	}

	if src.FlagToChannelAction != nil {
		return json.Marshal(&src.FlagToChannelAction)
	}

	if src.QuarantineUserAction != nil {
		return json.Marshal(&src.QuarantineUserAction)
	}

	if src.UserCommunicationDisabledAction != nil {
		return json.Marshal(&src.UserCommunicationDisabledAction)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DefaultKeywordListUpsertRequestActionsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.BlockMessageAction != nil {
		return obj.BlockMessageAction
	}

	if obj.FlagToChannelAction != nil {
		return obj.FlagToChannelAction
	}

	if obj.QuarantineUserAction != nil {
		return obj.QuarantineUserAction
	}

	if obj.UserCommunicationDisabledAction != nil {
		return obj.UserCommunicationDisabledAction
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj DefaultKeywordListUpsertRequestActionsInner) GetActualInstanceValue() (interface{}) {
	if obj.BlockMessageAction != nil {
		return *obj.BlockMessageAction
	}

	if obj.FlagToChannelAction != nil {
		return *obj.FlagToChannelAction
	}

	if obj.QuarantineUserAction != nil {
		return *obj.QuarantineUserAction
	}

	if obj.UserCommunicationDisabledAction != nil {
		return *obj.UserCommunicationDisabledAction
	}

	// all schemas are nil
	return nil
}

type NullableDefaultKeywordListUpsertRequestActionsInner struct {
	value *DefaultKeywordListUpsertRequestActionsInner
	isSet bool
}

func (v NullableDefaultKeywordListUpsertRequestActionsInner) Get() *DefaultKeywordListUpsertRequestActionsInner {
	return v.value
}

func (v *NullableDefaultKeywordListUpsertRequestActionsInner) Set(val *DefaultKeywordListUpsertRequestActionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultKeywordListUpsertRequestActionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultKeywordListUpsertRequestActionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultKeywordListUpsertRequestActionsInner(val *DefaultKeywordListUpsertRequestActionsInner) *NullableDefaultKeywordListUpsertRequestActionsInner {
	return &NullableDefaultKeywordListUpsertRequestActionsInner{value: val, isSet: true}
}

func (v NullableDefaultKeywordListUpsertRequestActionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultKeywordListUpsertRequestActionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


