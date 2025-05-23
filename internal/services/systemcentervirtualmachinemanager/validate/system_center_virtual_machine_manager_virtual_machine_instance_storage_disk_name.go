// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package validate

import (
	"fmt"
	"regexp"
)

func SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskName(i interface{}, k string) (warnings []string, errors []error) {
	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return
	}

	if !regexp.MustCompile("^[a-zA-Z]([-_a-zA-Z0-9]{0,78}[a-zA-Z0-9])?$").MatchString(v) {
		errors = append(errors, fmt.Errorf("%q must start and end with a letter, may contain alphanumeric characters, dashes or underscores and must be between 1 and 80 characters long", k))
	}

	return warnings, errors
}
