/* Copyright © 2020 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

package bindings

import "github.com/wenqiq/vsphere-automation-sdk-go/runtime/data"

// Structure Base interface that represents a Go language bindings structure.
type Structure interface {
	GetType__() BindingType
	GetDataValue__() (data.DataValue, []error)
}
