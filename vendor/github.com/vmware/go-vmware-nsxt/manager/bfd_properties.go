/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package manager

// BFD information
type BfdProperties struct {

	// True if tunnel is active in a gateway HA setup
	Active bool `json:"active,omitempty"`

	// A short message indicating what the BFD session thinks is wrong in case of a problem
	Diagnostic string `json:"diagnostic,omitempty"`

	// True if the BFD session believes this interface may be used to forward traffic
	Forwarding bool `json:"forwarding,omitempty"`

	// A short message indicating what the remote interface's BFD session thinks is wrong in case of a problem
	RemoteDiagnostic string `json:"remote_diagnostic,omitempty"`

	// State of the remote interface's BFD session
	RemoteState string `json:"remote_state,omitempty"`

	// State of the BFD session
	State string `json:"state,omitempty"`
}