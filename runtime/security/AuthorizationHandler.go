/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package security

import "github.com/zhengxiexie/vsphere-automation-sdk-go/runtime/core"

//The AuthorizationHandler interface is used to verify the authentication
//data provided in the security context against an identity source.

type AuthorizationHandler interface {
	// returns (true, nil) if auth is successful
	// returns (false, err) if auth fails
	// returns (false, nil) otherwise
	Authorize(serviceID string, operationID string, ctx core.SecurityContext) (bool, error)
}
