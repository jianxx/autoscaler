// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateInstanceSourceViaImageDetails The details for updating the instance source from an image.
type UpdateInstanceSourceViaImageDetails struct {

	// The OCID of the image used to boot the instance.
	ImageId *string `mandatory:"true" json:"imageId"`

	// Whether to preserve the boot volume that was previously attached to the instance after a successful replacement of that boot volume.
	IsPreserveBootVolumeEnabled *bool `mandatory:"false" json:"isPreserveBootVolumeEnabled"`

	// The size of the boot volume in GBs. Minimum value is 50 GB and maximum value is 32,768 GB (32 TB).
	BootVolumeSizeInGBs *int64 `mandatory:"false" json:"bootVolumeSizeInGBs"`

	// The OCID of the Vault service key to assign as the master encryption key for the boot volume.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`
}

// GetIsPreserveBootVolumeEnabled returns IsPreserveBootVolumeEnabled
func (m UpdateInstanceSourceViaImageDetails) GetIsPreserveBootVolumeEnabled() *bool {
	return m.IsPreserveBootVolumeEnabled
}

func (m UpdateInstanceSourceViaImageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInstanceSourceViaImageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateInstanceSourceViaImageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateInstanceSourceViaImageDetails UpdateInstanceSourceViaImageDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeUpdateInstanceSourceViaImageDetails
	}{
		"image",
		(MarshalTypeUpdateInstanceSourceViaImageDetails)(m),
	}

	return json.Marshal(&s)
}
