/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha3

import (
	infrav1alpha4 "sigs.k8s.io/cluster-api-provider-gcp/api/v1alpha4"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this GCPMachineTemplate to the Hub version (v1alpha4).
func (src *GCPMachineTemplate) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*infrav1alpha4.GCPMachineTemplate)

	if err := Convert_v1alpha3_GCPMachineTemplate_To_v1alpha4_GCPMachineTemplate(src, dst, nil); err != nil {
		return err
	}

	return nil
}

// ConvertFrom converts from the Hub version (v1alpha4) to this version.
func (dst *GCPMachineTemplate) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*infrav1alpha4.GCPMachineTemplate)
	if err := Convert_v1alpha4_GCPMachineTemplate_To_v1alpha3_GCPMachineTemplate(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion.
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}

	return nil
}

// ConvertTo converts this GCPMachineTemplateList to the Hub version (v1alpha4).
func (src *GCPMachineTemplateList) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*infrav1alpha4.GCPMachineTemplateList)
	return Convert_v1alpha3_GCPMachineTemplateList_To_v1alpha4_GCPMachineTemplateList(src, dst, nil)
}

// ConvertFrom converts from the Hub version (v1alpha4) to this version.
func (dst *GCPMachineTemplateList) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*infrav1alpha4.GCPMachineTemplateList)
	return Convert_v1alpha4_GCPMachineTemplateList_To_v1alpha3_GCPMachineTemplateList(src, dst, nil)
}