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
	"fmt"
	"gopkg.in/validator.v2"
)

// ContainerComponentForMessageRequestComponentsInner - struct for ContainerComponentForMessageRequestComponentsInner
type ContainerComponentForMessageRequestComponentsInner struct {
	ActionRowComponentForMessageRequest *ActionRowComponentForMessageRequest
	FileComponentForMessageRequest *FileComponentForMessageRequest
	MediaGalleryComponentForMessageRequest *MediaGalleryComponentForMessageRequest
	SectionComponentForMessageRequest *SectionComponentForMessageRequest
	SeparatorComponentForMessageRequest *SeparatorComponentForMessageRequest
	TextDisplayComponentForMessageRequest *TextDisplayComponentForMessageRequest
}

// ActionRowComponentForMessageRequestAsContainerComponentForMessageRequestComponentsInner is a convenience function that returns ActionRowComponentForMessageRequest wrapped in ContainerComponentForMessageRequestComponentsInner
func ActionRowComponentForMessageRequestAsContainerComponentForMessageRequestComponentsInner(v *ActionRowComponentForMessageRequest) ContainerComponentForMessageRequestComponentsInner {
	return ContainerComponentForMessageRequestComponentsInner{
		ActionRowComponentForMessageRequest: v,
	}
}

// FileComponentForMessageRequestAsContainerComponentForMessageRequestComponentsInner is a convenience function that returns FileComponentForMessageRequest wrapped in ContainerComponentForMessageRequestComponentsInner
func FileComponentForMessageRequestAsContainerComponentForMessageRequestComponentsInner(v *FileComponentForMessageRequest) ContainerComponentForMessageRequestComponentsInner {
	return ContainerComponentForMessageRequestComponentsInner{
		FileComponentForMessageRequest: v,
	}
}

// MediaGalleryComponentForMessageRequestAsContainerComponentForMessageRequestComponentsInner is a convenience function that returns MediaGalleryComponentForMessageRequest wrapped in ContainerComponentForMessageRequestComponentsInner
func MediaGalleryComponentForMessageRequestAsContainerComponentForMessageRequestComponentsInner(v *MediaGalleryComponentForMessageRequest) ContainerComponentForMessageRequestComponentsInner {
	return ContainerComponentForMessageRequestComponentsInner{
		MediaGalleryComponentForMessageRequest: v,
	}
}

// SectionComponentForMessageRequestAsContainerComponentForMessageRequestComponentsInner is a convenience function that returns SectionComponentForMessageRequest wrapped in ContainerComponentForMessageRequestComponentsInner
func SectionComponentForMessageRequestAsContainerComponentForMessageRequestComponentsInner(v *SectionComponentForMessageRequest) ContainerComponentForMessageRequestComponentsInner {
	return ContainerComponentForMessageRequestComponentsInner{
		SectionComponentForMessageRequest: v,
	}
}

// SeparatorComponentForMessageRequestAsContainerComponentForMessageRequestComponentsInner is a convenience function that returns SeparatorComponentForMessageRequest wrapped in ContainerComponentForMessageRequestComponentsInner
func SeparatorComponentForMessageRequestAsContainerComponentForMessageRequestComponentsInner(v *SeparatorComponentForMessageRequest) ContainerComponentForMessageRequestComponentsInner {
	return ContainerComponentForMessageRequestComponentsInner{
		SeparatorComponentForMessageRequest: v,
	}
}

// TextDisplayComponentForMessageRequestAsContainerComponentForMessageRequestComponentsInner is a convenience function that returns TextDisplayComponentForMessageRequest wrapped in ContainerComponentForMessageRequestComponentsInner
func TextDisplayComponentForMessageRequestAsContainerComponentForMessageRequestComponentsInner(v *TextDisplayComponentForMessageRequest) ContainerComponentForMessageRequestComponentsInner {
	return ContainerComponentForMessageRequestComponentsInner{
		TextDisplayComponentForMessageRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContainerComponentForMessageRequestComponentsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ActionRowComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.ActionRowComponentForMessageRequest)
	if err == nil {
		jsonActionRowComponentForMessageRequest, _ := json.Marshal(dst.ActionRowComponentForMessageRequest)
		if string(jsonActionRowComponentForMessageRequest) == "{}" { // empty struct
			dst.ActionRowComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.ActionRowComponentForMessageRequest); err != nil {
				dst.ActionRowComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActionRowComponentForMessageRequest = nil
	}

	// try to unmarshal data into FileComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.FileComponentForMessageRequest)
	if err == nil {
		jsonFileComponentForMessageRequest, _ := json.Marshal(dst.FileComponentForMessageRequest)
		if string(jsonFileComponentForMessageRequest) == "{}" { // empty struct
			dst.FileComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.FileComponentForMessageRequest); err != nil {
				dst.FileComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.FileComponentForMessageRequest = nil
	}

	// try to unmarshal data into MediaGalleryComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.MediaGalleryComponentForMessageRequest)
	if err == nil {
		jsonMediaGalleryComponentForMessageRequest, _ := json.Marshal(dst.MediaGalleryComponentForMessageRequest)
		if string(jsonMediaGalleryComponentForMessageRequest) == "{}" { // empty struct
			dst.MediaGalleryComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.MediaGalleryComponentForMessageRequest); err != nil {
				dst.MediaGalleryComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.MediaGalleryComponentForMessageRequest = nil
	}

	// try to unmarshal data into SectionComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.SectionComponentForMessageRequest)
	if err == nil {
		jsonSectionComponentForMessageRequest, _ := json.Marshal(dst.SectionComponentForMessageRequest)
		if string(jsonSectionComponentForMessageRequest) == "{}" { // empty struct
			dst.SectionComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.SectionComponentForMessageRequest); err != nil {
				dst.SectionComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.SectionComponentForMessageRequest = nil
	}

	// try to unmarshal data into SeparatorComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.SeparatorComponentForMessageRequest)
	if err == nil {
		jsonSeparatorComponentForMessageRequest, _ := json.Marshal(dst.SeparatorComponentForMessageRequest)
		if string(jsonSeparatorComponentForMessageRequest) == "{}" { // empty struct
			dst.SeparatorComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.SeparatorComponentForMessageRequest); err != nil {
				dst.SeparatorComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.SeparatorComponentForMessageRequest = nil
	}

	// try to unmarshal data into TextDisplayComponentForMessageRequest
	err = newStrictDecoder(data).Decode(&dst.TextDisplayComponentForMessageRequest)
	if err == nil {
		jsonTextDisplayComponentForMessageRequest, _ := json.Marshal(dst.TextDisplayComponentForMessageRequest)
		if string(jsonTextDisplayComponentForMessageRequest) == "{}" { // empty struct
			dst.TextDisplayComponentForMessageRequest = nil
		} else {
			if err = validator.Validate(dst.TextDisplayComponentForMessageRequest); err != nil {
				dst.TextDisplayComponentForMessageRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.TextDisplayComponentForMessageRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ActionRowComponentForMessageRequest = nil
		dst.FileComponentForMessageRequest = nil
		dst.MediaGalleryComponentForMessageRequest = nil
		dst.SectionComponentForMessageRequest = nil
		dst.SeparatorComponentForMessageRequest = nil
		dst.TextDisplayComponentForMessageRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ContainerComponentForMessageRequestComponentsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ContainerComponentForMessageRequestComponentsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContainerComponentForMessageRequestComponentsInner) MarshalJSON() ([]byte, error) {
	if src.ActionRowComponentForMessageRequest != nil {
		return json.Marshal(&src.ActionRowComponentForMessageRequest)
	}

	if src.FileComponentForMessageRequest != nil {
		return json.Marshal(&src.FileComponentForMessageRequest)
	}

	if src.MediaGalleryComponentForMessageRequest != nil {
		return json.Marshal(&src.MediaGalleryComponentForMessageRequest)
	}

	if src.SectionComponentForMessageRequest != nil {
		return json.Marshal(&src.SectionComponentForMessageRequest)
	}

	if src.SeparatorComponentForMessageRequest != nil {
		return json.Marshal(&src.SeparatorComponentForMessageRequest)
	}

	if src.TextDisplayComponentForMessageRequest != nil {
		return json.Marshal(&src.TextDisplayComponentForMessageRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContainerComponentForMessageRequestComponentsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ActionRowComponentForMessageRequest != nil {
		return obj.ActionRowComponentForMessageRequest
	}

	if obj.FileComponentForMessageRequest != nil {
		return obj.FileComponentForMessageRequest
	}

	if obj.MediaGalleryComponentForMessageRequest != nil {
		return obj.MediaGalleryComponentForMessageRequest
	}

	if obj.SectionComponentForMessageRequest != nil {
		return obj.SectionComponentForMessageRequest
	}

	if obj.SeparatorComponentForMessageRequest != nil {
		return obj.SeparatorComponentForMessageRequest
	}

	if obj.TextDisplayComponentForMessageRequest != nil {
		return obj.TextDisplayComponentForMessageRequest
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ContainerComponentForMessageRequestComponentsInner) GetActualInstanceValue() (interface{}) {
	if obj.ActionRowComponentForMessageRequest != nil {
		return *obj.ActionRowComponentForMessageRequest
	}

	if obj.FileComponentForMessageRequest != nil {
		return *obj.FileComponentForMessageRequest
	}

	if obj.MediaGalleryComponentForMessageRequest != nil {
		return *obj.MediaGalleryComponentForMessageRequest
	}

	if obj.SectionComponentForMessageRequest != nil {
		return *obj.SectionComponentForMessageRequest
	}

	if obj.SeparatorComponentForMessageRequest != nil {
		return *obj.SeparatorComponentForMessageRequest
	}

	if obj.TextDisplayComponentForMessageRequest != nil {
		return *obj.TextDisplayComponentForMessageRequest
	}

	// all schemas are nil
	return nil
}

type NullableContainerComponentForMessageRequestComponentsInner struct {
	value *ContainerComponentForMessageRequestComponentsInner
	isSet bool
}

func (v NullableContainerComponentForMessageRequestComponentsInner) Get() *ContainerComponentForMessageRequestComponentsInner {
	return v.value
}

func (v *NullableContainerComponentForMessageRequestComponentsInner) Set(val *ContainerComponentForMessageRequestComponentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerComponentForMessageRequestComponentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerComponentForMessageRequestComponentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerComponentForMessageRequestComponentsInner(val *ContainerComponentForMessageRequestComponentsInner) *NullableContainerComponentForMessageRequestComponentsInner {
	return &NullableContainerComponentForMessageRequestComponentsInner{value: val, isSet: true}
}

func (v NullableContainerComponentForMessageRequestComponentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerComponentForMessageRequestComponentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


