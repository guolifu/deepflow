/*
 * Copyright (c) 2022 Yunshan Networks
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package event

import (
	"errors"
	"fmt"

	cloudmodel "github.com/deepflowys/deepflow/server/controller/cloud/model"
	"github.com/deepflowys/deepflow/server/controller/common"
	"github.com/deepflowys/deepflow/server/controller/db/mysql"
	"github.com/deepflowys/deepflow/server/controller/recorder/cache"
	. "github.com/deepflowys/deepflow/server/controller/recorder/common"
	"github.com/deepflowys/deepflow/server/libs/eventapi"
	"github.com/deepflowys/deepflow/server/libs/queue"
)

var (
	VMStateToString = map[int]string{
		common.VM_STATE_RUNNING:   "running",
		common.VM_STATE_STOPPED:   "stopped",
		common.VM_STATE_EXCEPTION: "exception",
	}
)

type VM struct {
	EventManager[cloudmodel.VM, mysql.VM, *cache.VM]
	deviceType int
}

func NewVM(toolDS *cache.ToolDataSet, eq *queue.OverwriteQueue) *VM {
	mng := &VM{
		EventManager[cloudmodel.VM, mysql.VM, *cache.VM]{
			resourceType: RESOURCE_TYPE_VM_EN,
			ToolDataSet:  toolDS,
			Queue:        eq,
		},
		common.VIF_DEVICE_TYPE_VM,
	}
	return mng
}

func (v *VM) ProduceByAdd(items []*mysql.VM) {
	for _, item := range items {
		var opts []eventapi.TagFieldOption
		info, err := v.ToolDataSet.GetVMInfoByID(item.ID)
		if err != nil {
			log.Error(err)
		} else {
			opts = append(opts, []eventapi.TagFieldOption{
				eventapi.TagAZID(info.AZID),
				eventapi.TagRegionID(info.RegionID),
				eventapi.TagHostID(info.HostID),
			}...)
		}
		opts = append(opts, []eventapi.TagFieldOption{
			eventapi.TagL3DeviceType(v.deviceType),
			eventapi.TagL3DeviceID(item.ID),
			eventapi.TagVPCID(item.VPCID),
		}...)

		v.createAndPutEvent(
			eventapi.RESOURCE_EVENT_TYPE_CREATE,
			item.Name,
			v.deviceType,
			item.ID,
			opts...,
		)
	}
}

func (v *VM) ProduceByUpdate(cloudItem *cloudmodel.VM, diffBase *cache.VM) {
	id, name, err := v.getVMIDAndNameByLcuuid(cloudItem.Lcuuid)
	if err != nil {
		log.Error(err)
	}
	var eType string
	var description string
	if diffBase.LaunchServer != cloudItem.LaunchServer {
		eType = eventapi.RESOURCE_EVENT_TYPE_MIGRATE
		description = fmt.Sprintf("%s,%s", diffBase.LaunchServer, cloudItem.LaunchServer)
	}
	if diffBase.State != cloudItem.State {
		eType = eventapi.RESOURCE_EVENT_TYPE_UPDATE_STATE
		description = fmt.Sprintf("%s,%s", VMStateToString[diffBase.State], VMStateToString[cloudItem.State])
	}

	nIDs, ips := v.getIPNetworksByID(id)
	v.createAndPutEvent(
		eType,
		name,
		common.VIF_DEVICE_TYPE_VM,
		id,
		eventapi.TagDescription(description),
		eventapi.TagSubnetIDs(nIDs),
		eventapi.TagIPs(ips),
	)
}

func (v *VM) ProduceByDelete(lcuuids []string) {
	for _, lcuuid := range lcuuids {
		id, name, err := v.getVMIDAndNameByLcuuid(lcuuid)
		if err != nil {
			log.Errorf("%v, %v", idByLcuuidNotFound(v.resourceType, lcuuid), err)
		}

		v.createAndPutEvent(eventapi.RESOURCE_EVENT_TYPE_DELETE, name, v.deviceType, id)
	}
}

func (v *VM) getIPNetworksByID(id int) (networkIDs []uint32, ips []string) {
	ipNetworkMap, _ := v.ToolDataSet.EventToolDataSet.GetVMIPNetworkMapByID(id)
	for ip, nID := range ipNetworkMap {
		networkIDs = append(networkIDs, nID)
		ips = append(ips, ip)
	}
	return
}

func (v *VM) getVMIDAndNameByLcuuid(lcuuid string) (int, string, error) {
	id, ok := v.ToolDataSet.GetVMIDByLcuuid(lcuuid)
	if !ok {
		return 0, "", errors.New(nameByIDNotFound(v.resourceType, id))
	}
	name, err := v.ToolDataSet.GetVMNameByID(id)
	if !ok {
		return 0, "", err
	}

	return id, name, nil
}
