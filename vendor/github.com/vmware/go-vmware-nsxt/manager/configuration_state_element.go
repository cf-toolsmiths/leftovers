/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

type ConfigurationStateElement struct {

	// Error code
	FailureCode int64 `json:"failure_code,omitempty"`

	// Error message in case of failure
	FailureMessage string `json:"failure_message,omitempty"`

	// State of configuration on this sub system
	State string `json:"state,omitempty"`

	// URI of backing resource on sub system
	SubSystemAddress string `json:"sub_system_address,omitempty"`

	// Identifier of backing resource on sub system
	SubSystemId string `json:"sub_system_id,omitempty"`

	// Type of backing resource on sub system
	SubSystemType string `json:"sub_system_type,omitempty"`
}