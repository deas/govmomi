/*
Copyright (c) 2014-2022 VMware, Inc. All Rights Reserved.

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

package types

import "reflect"

type QueryVirtualDiskInfo_Task QueryVirtualDiskInfoRequestType

type QueryVirtualDiskInfoRequestType struct {
	This           ManagedObjectReference  `xml:"_this" json:"_this"`
	Name           string                  `xml:"name" json:"name"`
	Datacenter     *ManagedObjectReference `xml:"datacenter,omitempty" json:"datacenter,omitempty"`
	IncludeParents bool                    `xml:"includeParents" json:"includeParents"`
}

/*
type queryVirtualDiskInfoTaskRequest struct {
	This           types.ManagedObjectReference  `xml:"_this"`
	Name           string                        `xml:"name"`
	Datacenter     *types.ManagedObjectReference `xml:"datacenter,omitempty"`
	IncludeParents bool                          `xml:"includeParents"`
}
*/

func init() {
	t["QueryVirtualDiskInfoRequestType"] = reflect.TypeOf((*QueryVirtualDiskInfoRequestType)(nil)).Elem()
}

func init() {
	t["QueryVirtualDiskInfo_Task"] = reflect.TypeOf((*QueryVirtualDiskInfo_Task)(nil)).Elem()
}
