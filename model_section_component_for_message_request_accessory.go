/** 
 * Discord HTTP API (Preview) - REST API Client
 * Preview of the Discord v10 HTTP API specification. See https://discord.com/developers/docs for more details.
 * 
 * ## Metadata
 *    * - **Copyright**: Copyright (c) 2025 Qntx
 *    * - **Author**: ΣX <gitctrlx@gmail.com>
 *    * - **Version**: 10
 *    * - **Modified**: 2025-07-01T10:17:22.835370245Z[Etc/UTC]
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

// SectionComponentForMessageRequestAccessory - struct for SectionComponentForMessageRequestAccessory
type SectionComponentForMessageRequestAccessory struct {
	ButtonComponentForMessageRequest *ButtonComponentForMessageRequest
	ThumbnailComponentForMessageRequest *ThumbnailComponentForMessageRequest
}

// ButtonComponentForMessageRequestAsSectionComponentForMessageRequestAccessory is a convenience function that returns ButtonComponentForMessageRequest wrapped in SectionComponentForMessageRequestAccessory
func ButtonComponentForMessageRequestAsSectionComponentForMessageRequestAccessory(v *ButtonComponentForMessageRequest) SectionComponentForMessageRequestAccessory {
	return SectionComponentForMessageRequestAccessory{
		ButtonComponentForMessageRequest: v,
	}
}

// ThumbnailComponentForMessageRequestAsSectionComponentForMessageRequestAccessory is a convenience function that returns ThumbnailComponentForMessageRequest wrapped in SectionComponentForMessageRequestAccessory
func ThumbnailComponentForMessageRequestAsSectionComponentForMessageRequestAccessory(v *ThumbnailComponentForMessageRequest) SectionComponentForMessageRequestAccessory {
	return SectionComponentForMessageRequestAccessory{
		ThumbnailComponentForMessageRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SectionComponentForMessageRequestAccessory) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ButtonComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.ButtonComponentForMessageRequest)
	if err == nil {
		jsonButtonComponentForMessageRequest, _ := json.Marshal(dst.ButtonComponentForMessageRequest)
		if string(jsonButtonComponentForMessageRequest) == "{}" { // empty struct
			dst.ButtonComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.ButtonComponentForMessageRequest); err != nil {
				dst.ButtonComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.ButtonComponentForMessageRequest = nil
	}

	// try to unmarshal data into ThumbnailComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.ThumbnailComponentForMessageRequest)
	if err == nil {
		jsonThumbnailComponentForMessageRequest, _ := json.Marshal(dst.ThumbnailComponentForMessageRequest)
		if string(jsonThumbnailComponentForMessageRequest) == "{}" { // empty struct
			dst.ThumbnailComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.ThumbnailComponentForMessageRequest); err != nil {
				dst.ThumbnailComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.ThumbnailComponentForMessageRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ButtonComponentForMessageRequest = nil
		dst.ThumbnailComponentForMessageRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SectionComponentForMessageRequestAccessory)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SectionComponentForMessageRequestAccessory)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SectionComponentForMessageRequestAccessory) MarshalJSON() ([]byte, error) {
	if src.ButtonComponentForMessageRequest != nil {
		return json.Marshal(&src.ButtonComponentForMessageRequest)
	}

	if src.ThumbnailComponentForMessageRequest != nil {
		return json.Marshal(&src.ThumbnailComponentForMessageRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SectionComponentForMessageRequestAccessory) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ButtonComponentForMessageRequest != nil {
		return obj.ButtonComponentForMessageRequest
	}

	if obj.ThumbnailComponentForMessageRequest != nil {
		return obj.ThumbnailComponentForMessageRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj SectionComponentForMessageRequestAccessory) GetActualInstanceValue() (interface{}) {
	if obj.ButtonComponentForMessageRequest != nil {
		return *obj.ButtonComponentForMessageRequest
	}

	if obj.ThumbnailComponentForMessageRequest != nil {
		return *obj.ThumbnailComponentForMessageRequest
	}

	// all schemas are nil
	return nil
}

type NullableSectionComponentForMessageRequestAccessory struct {
	value *SectionComponentForMessageRequestAccessory
	isSet bool
}

func (v NullableSectionComponentForMessageRequestAccessory) Get() *SectionComponentForMessageRequestAccessory {
	return v.value
}

func (v *NullableSectionComponentForMessageRequestAccessory) Set(val *SectionComponentForMessageRequestAccessory) {
	v.value = val
	v.isSet = true
}

func (v NullableSectionComponentForMessageRequestAccessory) IsSet() bool {
	return v.isSet
}

func (v *NullableSectionComponentForMessageRequestAccessory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSectionComponentForMessageRequestAccessory(val *SectionComponentForMessageRequestAccessory) *NullableSectionComponentForMessageRequestAccessory {
	return &NullableSectionComponentForMessageRequestAccessory{value: val, isSet: true}
}

func (v NullableSectionComponentForMessageRequestAccessory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSectionComponentForMessageRequestAccessory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


