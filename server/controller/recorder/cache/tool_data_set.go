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

package cache

import (
	"errors"
	"fmt"

	cloudmodel "github.com/deepflowys/deepflow/server/controller/cloud/model"
	"github.com/deepflowys/deepflow/server/controller/common"
	"github.com/deepflowys/deepflow/server/controller/db/mysql"
	. "github.com/deepflowys/deepflow/server/controller/recorder/common"
)

// 各类资源的映射关系，用于刷新资源时，转换所需数据
type ToolDataSet struct {
	// 仅资源变更事件所需的数据
	EventToolDataSet

	AZLcuuidToID map[string]int

	RegionLcuuidToID map[string]int
	RegionIDToLcuuid map[int]string

	HostLcuuidToID map[string]int

	VMLcuuidToID map[string]int

	VPCLcuuidToID map[string]int
	VPCIDToLcuuid map[int]string

	PublicNetworkID   int
	NetworkLcuuidToID map[string]int
	NetworkIDToLcuuid map[int]string

	VRouterLcuuidToID map[string]int

	DHCPPortLcuuidToID map[string]int

	VInterfaceLcuuidToID         map[string]int
	VInterfaceLcuuidToType       map[string]int
	VInterfaceLcuuidToIndex      map[string]int
	VInterfaceLcuuidToNetworkID  map[string]int
	VInterfaceLcuuidToDeviceType map[string]int
	VInterfaceLcuuidToDeviceID   map[string]int

	HostIDToVinterfaceIndexes          map[int][]int
	VMIDToVinterfaceIndexes            map[int][]int
	VRouterIDToVinterfaceIndexes       map[int][]int
	DHCPPortIDToVinterfaceIndexes      map[int][]int
	NATGatewayIDToVinterfaceIndexes    map[int][]int
	LBIDToVinterfaceIndexes            map[int][]int
	RDSInstanceIDToVinterfaceIndexes   map[int][]int
	RedisInstanceIDToVinterfaceIndexes map[int][]int
	PodNodeIDToVinterfaceIndexes       map[int][]int
	PodServiceIDToVinterfaceIndexes    map[int][]int
	PodIDToVinterfaceIndexes           map[int][]int

	SecurityGroupLcuuidToID map[string]int

	NATGatewayLcuuidToID map[string]int

	LBLcuuidToID         map[string]int
	LBListenerLcuuidToID map[string]int

	RDSInstanceLcuuidToID map[string]int

	RedisInstanceLcuuidToID map[string]int

	PodClusterLcuuidToID map[string]int

	PodNodeLcuuidToID map[string]int
	PodNodeIDToLcuuid map[int]string

	PodNamespaceLcuuidToID map[string]int

	PodIngressLcuuidToID     map[string]int
	PodIngressIDToLcuuid     map[int]string
	PodIngressRuleLcuuidToID map[string]int

	PodServiceLcuuidToID map[string]int

	PodGroupLcuuidToID map[string]int

	PodReplicaSetLcuuidToID map[string]int
	PodReplicaSetIDToLcuuid map[int]string

	PodLcuuidToID map[string]int
}

func NewToolDataSet() ToolDataSet {
	return ToolDataSet{
		EventToolDataSet: NewEventToolDataSet(),

		AZLcuuidToID: make(map[string]int),

		RegionLcuuidToID: make(map[string]int),
		RegionIDToLcuuid: make(map[int]string),

		HostLcuuidToID: make(map[string]int),

		VMLcuuidToID: make(map[string]int),

		VPCLcuuidToID: make(map[string]int),
		VPCIDToLcuuid: make(map[int]string),

		NetworkLcuuidToID: make(map[string]int),
		NetworkIDToLcuuid: make(map[int]string),
		// NetworkIDToSubnetIndexes: make(map[int][]int),

		// SubnetLcuuidToIndex:     make(map[string]int),

		VRouterLcuuidToID: make(map[string]int),

		DHCPPortLcuuidToID: make(map[string]int),

		VInterfaceLcuuidToID:         make(map[string]int),
		VInterfaceLcuuidToType:       make(map[string]int),
		VInterfaceLcuuidToIndex:      make(map[string]int),
		VInterfaceLcuuidToNetworkID:  make(map[string]int),
		VInterfaceLcuuidToDeviceType: make(map[string]int),
		VInterfaceLcuuidToDeviceID:   make(map[string]int),

		HostIDToVinterfaceIndexes:          make(map[int][]int),
		VMIDToVinterfaceIndexes:            make(map[int][]int),
		VRouterIDToVinterfaceIndexes:       make(map[int][]int),
		DHCPPortIDToVinterfaceIndexes:      make(map[int][]int),
		NATGatewayIDToVinterfaceIndexes:    make(map[int][]int),
		LBIDToVinterfaceIndexes:            make(map[int][]int),
		RDSInstanceIDToVinterfaceIndexes:   make(map[int][]int),
		RedisInstanceIDToVinterfaceIndexes: make(map[int][]int),
		PodNodeIDToVinterfaceIndexes:       make(map[int][]int),
		PodServiceIDToVinterfaceIndexes:    make(map[int][]int),
		PodIDToVinterfaceIndexes:           make(map[int][]int),

		SecurityGroupLcuuidToID: make(map[string]int),

		NATGatewayLcuuidToID: make(map[string]int),

		LBLcuuidToID:         make(map[string]int),
		LBListenerLcuuidToID: make(map[string]int),

		RDSInstanceLcuuidToID: make(map[string]int),

		RedisInstanceLcuuidToID: make(map[string]int),

		PodClusterLcuuidToID: make(map[string]int),

		PodNodeLcuuidToID: make(map[string]int),
		PodNodeIDToLcuuid: make(map[int]string),

		PodNamespaceLcuuidToID: make(map[string]int),

		PodIngressLcuuidToID:     make(map[string]int),
		PodIngressIDToLcuuid:     make(map[int]string),
		PodIngressRuleLcuuidToID: make(map[string]int),

		PodServiceLcuuidToID: make(map[string]int),

		PodGroupLcuuidToID: make(map[string]int),

		PodReplicaSetLcuuidToID: make(map[string]int),
		PodReplicaSetIDToLcuuid: make(map[int]string),

		PodLcuuidToID: make(map[string]int),
	}
}

func (t *ToolDataSet) addAZ(item *mysql.AZ) {
	t.AZLcuuidToID[item.Lcuuid] = item.ID
	log.Info(addToToolMap(RESOURCE_TYPE_AZ_EN, item.Lcuuid))
}

func (t *ToolDataSet) deleteAZ(lcuuid string) {
	delete(t.AZLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_AZ_EN, lcuuid))
}

func (t *ToolDataSet) GetAZIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.AZLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_AZ_EN, lcuuid))
	var az mysql.AZ
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&az)
	if result.RowsAffected == 1 {
		t.addAZ(&az)
		return az.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_AZ_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) addRegion(item *mysql.Region) {
	t.RegionLcuuidToID[item.Lcuuid] = item.ID
	t.RegionIDToLcuuid[item.ID] = item.Lcuuid
	log.Info(addToToolMap(RESOURCE_TYPE_REGION_EN, item.Lcuuid))
}

func (t *ToolDataSet) deleteRegion(lcuuid string) {
	delete(t.RegionLcuuidToID, lcuuid)
	id, _ := t.GetRegionIDByLcuuid(lcuuid)
	delete(t.RegionIDToLcuuid, id)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_REGION_EN, lcuuid))
}

func (t *ToolDataSet) addHost(item *mysql.Host) {
	t.HostLcuuidToID[item.Lcuuid] = item.ID
	t.hostIPToID[item.IP] = item.ID
	t.hostIDtoInfo[item.ID] = &hostInfo{Name: item.Name}
	if regionID, ok := t.GetRegionIDByLcuuid(item.Region); ok {
		t.hostIDtoInfo[item.ID].RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(item.AZ); ok {
		t.hostIDtoInfo[item.ID].AZID = azID
	}
	log.Info(addToToolMap(RESOURCE_TYPE_HOST_EN, item.Lcuuid))
}

func (t *ToolDataSet) deleteHost(lcuuid string) {
	id, _ := t.GetHostIDByLcuuid(lcuuid)
	delete(t.hostIDtoInfo, id)
	delete(t.HostLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_HOST_EN, lcuuid))
}

func (t *ToolDataSet) updateHost(cloudItem *cloudmodel.Host) {
	defer log.Info(updateToolMap(RESOURCE_TYPE_HOST_EN, cloudItem.Lcuuid))
	id, _ := t.GetHostIDByLcuuid(cloudItem.Lcuuid)
	info, err := t.GetHostInfoByID(id)
	if err != nil {
		log.Error(err)
		return
	}
	info.Name = cloudItem.Name
	if regionID, ok := t.GetRegionIDByLcuuid(cloudItem.RegionLcuuid); ok {
		t.hostIDtoInfo[id].RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(cloudItem.AZLcuuid); ok {
		t.hostIDtoInfo[id].RegionID = azID
	}
}

func (t *ToolDataSet) addVM(item *mysql.VM) {
	t.VMLcuuidToID[item.Lcuuid] = item.ID
	t.vmIDToInfo[item.ID] = &vmInfo{
		Name:  item.Name,
		VPCID: item.VPCID,
	}
	if regionID, ok := t.GetRegionIDByLcuuid(item.Region); ok {
		t.vmIDToInfo[item.ID].RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(item.AZ); ok {
		t.vmIDToInfo[item.ID].AZID = azID
	}
	if hostID, ok := t.GetHostIDByIP(item.LaunchServer); ok {
		t.vmIDToInfo[item.ID].HostID = hostID
	}
	log.Info(addToToolMap(RESOURCE_TYPE_VM_EN, item.Lcuuid))
}

func (t *ToolDataSet) updateVM(cloudItem *cloudmodel.VM) {
	defer log.Info(updateToolMap(RESOURCE_TYPE_VM_EN, cloudItem.Lcuuid))
	id, _ := t.GetVMIDByLcuuid(cloudItem.Lcuuid)
	info, err := t.GetVMInfoByID(id)
	if err != nil {
		log.Error(err)
		return
	}
	if regionID, ok := t.GetRegionIDByLcuuid(cloudItem.RegionLcuuid); ok {
		info.RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(cloudItem.AZLcuuid); ok {
		info.AZID = azID
	}
	if vpcID, ok := t.VPCLcuuidToID[cloudItem.Lcuuid]; ok {
		info.VPCID = vpcID
	}
	if hostID, ok := t.GetHostIDByIP(cloudItem.LaunchServer); ok {
		info.HostID = hostID
	}
}

func (t *ToolDataSet) deleteVM(lcuuid string) {
	id, _ := t.GetVMIDByLcuuid(lcuuid)
	delete(t.VMIDToIPNetworkIDMap, id)
	delete(t.VMLcuuidToID, lcuuid)
	delete(t.vmIDToInfo, id)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_VM_EN, lcuuid))
}

func (t *ToolDataSet) addVPC(item *mysql.VPC) {
	t.VPCLcuuidToID[item.Lcuuid] = item.ID
	t.VPCIDToLcuuid[item.ID] = item.Lcuuid
	log.Info(addToToolMap(RESOURCE_TYPE_VPC_EN, item.Lcuuid))
}

func (t *ToolDataSet) deleteVPC(lcuuid string) {
	id, _ := t.GetVPCIDByLcuuid(lcuuid)
	delete(t.VPCIDToLcuuid, id)
	delete(t.VPCLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_VPC_EN, lcuuid))
}

func (t *ToolDataSet) addNetwork(item *mysql.Network) {
	t.NetworkLcuuidToID[item.Lcuuid] = item.ID
	t.NetworkIDToLcuuid[item.ID] = item.Lcuuid
	t.NetworkIDToName[item.ID] = item.Name
	log.Info(addToToolMap(RESOURCE_TYPE_NETWORK_EN, item.Lcuuid))
}

func (t *ToolDataSet) updateNetwork(cloudItem *cloudmodel.Network) {
	id, exists := t.GetNetworkIDByLcuuid(cloudItem.Lcuuid)
	if exists {
		t.NetworkIDToName[id] = cloudItem.Name
	}
	log.Info(updateToolMap(RESOURCE_TYPE_NETWORK_EN, cloudItem.Lcuuid))
}

func (t *ToolDataSet) deleteNetwork(lcuuid string) {
	id, _ := t.GetNetworkIDByLcuuid(lcuuid)
	delete(t.NetworkIDToLcuuid, id)
	delete(t.NetworkIDToName, id)
	delete(t.NetworkLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_NETWORK_EN, lcuuid))
}

func (t *ToolDataSet) addVRouter(item *mysql.VRouter) {
	t.VRouterLcuuidToID[item.Lcuuid] = item.ID
	t.vrouterIDToInfo[item.ID] = &vrouterInfo{
		Name:  item.Name,
		VPCID: item.VPCID,
	}
	if regionID, ok := t.GetRegionIDByLcuuid(item.Region); ok {
		t.vrouterIDToInfo[item.ID].RegionID = regionID
	}
	log.Info(addToToolMap(RESOURCE_TYPE_VROUTER_EN, item.Lcuuid))
}

func (t *ToolDataSet) updateVRouter(cloudItem *cloudmodel.VRouter) {
	defer log.Info(updateToolMap(RESOURCE_TYPE_VROUTER_EN, cloudItem.Lcuuid))
	id, _ := t.GetVRouterIDByLcuuid(cloudItem.Lcuuid)
	info, err := t.GetVRouterInfoByID(id)
	if err != nil {
		log.Error(err)
		return
	}
	info.Name = cloudItem.Name
	if regionID, ok := t.GetRegionIDByLcuuid(cloudItem.RegionLcuuid); ok {
		info.RegionID = regionID
	}
	if vpcID, ok := t.VPCLcuuidToID[cloudItem.Lcuuid]; ok {
		info.VPCID = vpcID
	}
}

func (t *ToolDataSet) deleteVRouter(lcuuid string) {
	id, _ := t.GetVRouterIDByLcuuid(lcuuid)
	delete(t.vrouterIDToInfo, id)
	delete(t.VRouterLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_VROUTER_EN, lcuuid))
}

func (t *ToolDataSet) addDHCPPort(item *mysql.DHCPPort) {
	t.DHCPPortLcuuidToID[item.Lcuuid] = item.ID
	t.dhcpPortIDToInfo[item.ID] = &dhcpPortInfo{
		Name:  item.Name,
		VPCID: item.VPCID,
	}
	if regionID, ok := t.GetRegionIDByLcuuid(item.Region); ok {
		t.dhcpPortIDToInfo[item.ID].RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(item.AZ); ok {
		t.dhcpPortIDToInfo[item.ID].AZID = azID
	}
	log.Info(addToToolMap(RESOURCE_TYPE_DHCP_PORT_EN, item.Lcuuid))
}

func (t *ToolDataSet) updateDHCPPort(cloudItem *cloudmodel.DHCPPort) {
	defer log.Info(updateToolMap(RESOURCE_TYPE_DHCP_PORT_EN, cloudItem.Lcuuid))
	id, _ := t.GetDHCPPortIDByLcuuid(cloudItem.Lcuuid)
	info, err := t.GetDHCPPortInfoByID(id)
	if err != nil {
		log.Error(err)
		return
	}
	info.Name = cloudItem.Name
	if regionID, ok := t.GetRegionIDByLcuuid(cloudItem.RegionLcuuid); ok {
		info.RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(cloudItem.AZLcuuid); ok {
		info.AZID = azID
	}
	if vpcID, ok := t.VPCLcuuidToID[cloudItem.Lcuuid]; ok {
		info.VPCID = vpcID
	}
}

func (t *ToolDataSet) deleteDHCPPort(lcuuid string) {
	id, _ := t.GetDHCPPortIDByLcuuid(lcuuid)
	delete(t.dhcpPortIDToInfo, id)
	delete(t.DHCPPortLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_DHCP_PORT_EN, lcuuid))
}

func (t *ToolDataSet) addVInterface(item *mysql.VInterface) {
	t.VInterfaceLcuuidToID[item.Lcuuid] = item.ID
	t.VInterfaceIDToLcuuid[item.ID] = item.Lcuuid
	t.VInterfaceLcuuidToNetworkID[item.Lcuuid] = item.NetworkID
	t.VInterfaceLcuuidToDeviceType[item.Lcuuid] = item.DeviceType
	t.VInterfaceLcuuidToDeviceID[item.Lcuuid] = item.DeviceID
	t.VInterfaceLcuuidToIndex[item.Lcuuid] = item.Index
	t.VInterfaceLcuuidToType[item.Lcuuid] = item.Type

	if item.DeviceType == common.VIF_DEVICE_TYPE_HOST {
		t.HostIDToVinterfaceIndexes[item.DeviceID] = append(
			t.HostIDToVinterfaceIndexes[item.DeviceID], item.ID,
		)
	} else if item.DeviceType == common.VIF_DEVICE_TYPE_VM {
		t.VMIDToVinterfaceIndexes[item.DeviceID] = append(
			t.VMIDToVinterfaceIndexes[item.DeviceID], item.ID,
		)
	} else if item.DeviceType == common.VIF_DEVICE_TYPE_VROUTER {
		t.VRouterIDToVinterfaceIndexes[item.DeviceID] = append(
			t.VRouterIDToVinterfaceIndexes[item.DeviceID], item.ID,
		)
	} else if item.DeviceType == common.VIF_DEVICE_TYPE_DHCP_PORT {
		t.DHCPPortIDToVinterfaceIndexes[item.DeviceID] = append(
			t.DHCPPortIDToVinterfaceIndexes[item.DeviceID], item.ID,
		)
	} else if item.DeviceType == common.VIF_DEVICE_TYPE_NAT_GATEWAY {
		t.NATGatewayIDToVinterfaceIndexes[item.DeviceID] = append(
			t.NATGatewayIDToVinterfaceIndexes[item.DeviceID], item.ID,
		)
	} else if item.DeviceType == common.VIF_DEVICE_TYPE_LB {
		t.LBIDToVinterfaceIndexes[item.DeviceID] = append(
			t.LBIDToVinterfaceIndexes[item.DeviceID], item.ID,
		)
	} else if item.DeviceType == common.VIF_DEVICE_TYPE_RDS_INSTANCE {
		t.RDSInstanceIDToVinterfaceIndexes[item.DeviceID] = append(
			t.RDSInstanceIDToVinterfaceIndexes[item.DeviceID], item.ID,
		)
	} else if item.DeviceType == common.VIF_DEVICE_TYPE_REDIS_INSTANCE {
		t.RedisInstanceIDToVinterfaceIndexes[item.DeviceID] = append(
			t.RedisInstanceIDToVinterfaceIndexes[item.DeviceID], item.ID,
		)
	} else if item.DeviceType == common.VIF_DEVICE_TYPE_POD_NODE {
		t.PodNodeIDToVinterfaceIndexes[item.DeviceID] = append(
			t.PodNodeIDToVinterfaceIndexes[item.DeviceID], item.ID,
		)
	} else if item.DeviceType == common.VIF_DEVICE_TYPE_POD_SERVICE {
		t.PodServiceIDToVinterfaceIndexes[item.DeviceID] = append(
			t.PodServiceIDToVinterfaceIndexes[item.DeviceID], item.ID,
		)
	} else if item.DeviceType == common.VIF_DEVICE_TYPE_POD {
		t.PodIDToVinterfaceIndexes[item.DeviceID] = append(
			t.PodIDToVinterfaceIndexes[item.DeviceID], item.ID,
		)
	}
	log.Info(addToToolMap(RESOURCE_TYPE_VINTERFACE_EN, item.Lcuuid))
}

func (t *ToolDataSet) updateVInterface(cloudItem *cloudmodel.VInterface) {
	t.VInterfaceLcuuidToType[cloudItem.Lcuuid] = cloudItem.Type
	log.Info(updateToolMap(RESOURCE_TYPE_VINTERFACE_EN, cloudItem.Lcuuid))
}

func (t *ToolDataSet) deleteVInterface(lcuuid string) {
	id, _ := t.VInterfaceLcuuidToID[lcuuid]
	delete(t.VInterfaceIDToLcuuid, id)
	delete(t.VInterfaceLcuuidToID, lcuuid)
	delete(t.VInterfaceLcuuidToNetworkID, lcuuid)
	delete(t.VInterfaceLcuuidToDeviceType, lcuuid)
	delete(t.VInterfaceLcuuidToDeviceID, lcuuid)
	delete(t.VInterfaceLcuuidToIndex, lcuuid)
	delete(t.VInterfaceLcuuidToType, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_VINTERFACE_EN, lcuuid))
}

func (t *ToolDataSet) addWANIP(item *mysql.WANIP) {
	t.WANIPLcuuidToVInterfaceID[item.Lcuuid] = item.VInterfaceID
	t.WANIPLcuuidToIP[item.Lcuuid] = item.IP
	vifLcuuid, _ := t.GetVInterfaceLcuuidByID(item.VInterfaceID)
	deviceType, _ := t.GetDeviceTypeByVInterfaceLcuuid(vifLcuuid)
	deviceID, _ := t.GetDeviceIDByVInterfaceLcuuid(vifLcuuid)
	networkID, _ := t.GetNetworkIDByVInterfaceLcuuid(vifLcuuid)
	t.setDeviceToIPNetworkMap(deviceType, deviceID, networkID, item.IP)
	log.Info(addToToolMap(RESOURCE_TYPE_WAN_IP_EN, item.Lcuuid))
}

func (t *ToolDataSet) deleteWANIP(lcuuid string) {
	vifID, _ := t.GetVInterfaceIDByWANIPLcuuid(lcuuid)
	vifLcuuid, _ := t.GetVInterfaceLcuuidByID(vifID)
	deviceType, _ := t.GetDeviceTypeByVInterfaceLcuuid(vifLcuuid)
	deviceID, _ := t.GetDeviceIDByVInterfaceLcuuid(vifLcuuid)
	networkID, _ := t.GetNetworkIDByVInterfaceLcuuid(vifLcuuid)
	ip, _ := t.GetWANIPByLcuuid(lcuuid)
	t.deleteDeviceToIPNetworkMapIP(deviceType, deviceID, networkID, ip)
	delete(t.WANIPLcuuidToVInterfaceID, lcuuid)
	delete(t.WANIPLcuuidToIP, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_WAN_IP_EN, lcuuid))
}

func (t *ToolDataSet) addLANIP(item *mysql.LANIP) {
	t.LANIPLcuuidToVInterfaceID[item.Lcuuid] = item.VInterfaceID
	t.LANIPLcuuidToIP[item.Lcuuid] = item.IP
	vifLcuuid, _ := t.GetVInterfaceLcuuidByID(item.VInterfaceID)
	deviceType, _ := t.GetDeviceTypeByVInterfaceLcuuid(vifLcuuid)
	deviceID, _ := t.GetDeviceIDByVInterfaceLcuuid(vifLcuuid)
	networkID, _ := t.GetNetworkIDByVInterfaceLcuuid(vifLcuuid)
	t.setDeviceToIPNetworkMap(deviceType, deviceID, networkID, item.IP)
	log.Info(addToToolMap(RESOURCE_TYPE_LAN_IP_EN, item.Lcuuid))
}

func (t *ToolDataSet) deleteLANIP(lcuuid string) {
	vifID, _ := t.GetVInterfaceIDByLANIPLcuuid(lcuuid)
	vifLcuuid, _ := t.GetVInterfaceLcuuidByID(vifID)
	deviceType, _ := t.GetDeviceTypeByVInterfaceLcuuid(vifLcuuid)
	deviceID, _ := t.GetDeviceIDByVInterfaceLcuuid(vifLcuuid)
	networkID, _ := t.GetNetworkIDByVInterfaceLcuuid(vifLcuuid)
	ip, _ := t.GetLANIPByLcuuid(lcuuid)
	t.deleteDeviceToIPNetworkMapIP(deviceType, deviceID, networkID, ip)
	delete(t.LANIPLcuuidToVInterfaceID, lcuuid)
	delete(t.LANIPLcuuidToIP, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_LAN_IP_EN, lcuuid))
}

func (t *ToolDataSet) addSecurityGroup(item *mysql.SecurityGroup) {
	t.SecurityGroupLcuuidToID[item.Lcuuid] = item.ID
	log.Info(addToToolMap(RESOURCE_TYPE_SECURITY_GROUP_EN, item.Lcuuid))
}

func (t *ToolDataSet) deleteSecurityGroup(lcuuid string) {
	delete(t.SecurityGroupLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_SECURITY_GROUP_EN, lcuuid))
}

func (t *ToolDataSet) addNATGateway(item *mysql.NATGateway) {
	t.NATGatewayLcuuidToID[item.Lcuuid] = item.ID
	t.natGatewayIDToInfo[item.ID] = &natGatewayInfo{
		Name:  item.Name,
		VPCID: item.VPCID,
	}
	if regionID, ok := t.GetRegionIDByLcuuid(item.Region); ok {
		t.natGatewayIDToInfo[item.ID].RegionID = regionID
	}
	if azID, ok := t.AZLcuuidToID[item.AZ]; ok {
		t.natGatewayIDToInfo[item.ID].AZID = azID
	}
	log.Info(addToToolMap(RESOURCE_TYPE_NAT_GATEWAY_EN, item.Lcuuid))
}

func (t *ToolDataSet) updateNATGateway(cloudItem *cloudmodel.NATGateway) {
	defer log.Info(updateToolMap(RESOURCE_TYPE_NAT_GATEWAY_EN, cloudItem.Lcuuid))
	id, _ := t.GetNATGatewayIDByLcuuid(cloudItem.Lcuuid)
	info, err := t.GetNATGatewayInfoByID(id)
	if err != nil {
		log.Error(err)
		return
	}
	info.Name = cloudItem.Name
	if regionID, ok := t.GetRegionIDByLcuuid(cloudItem.RegionLcuuid); ok {
		info.RegionID = regionID
	}
	if vpcID, ok := t.VPCLcuuidToID[cloudItem.Lcuuid]; ok {
		info.VPCID = vpcID
	}
}

func (t *ToolDataSet) deleteNATGateway(lcuuid string) {
	id, _ := t.GetNATGatewayIDByLcuuid(lcuuid)
	delete(t.natGatewayIDToInfo, id)
	delete(t.NATGatewayLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_NAT_GATEWAY_EN, lcuuid))
}

func (t *ToolDataSet) addLB(item *mysql.LB) {
	t.LBLcuuidToID[item.Lcuuid] = item.ID
	t.lbIDToInfo[item.ID] = &lbInfo{
		Name:  item.Name,
		VPCID: item.VPCID,
	}
	if regionID, ok := t.GetRegionIDByLcuuid(item.Region); ok {
		t.lbIDToInfo[item.ID].RegionID = regionID
	}
	log.Info(addToToolMap(RESOURCE_TYPE_LB_EN, item.Lcuuid))
}

func (t *ToolDataSet) updateLB(cloudItem *cloudmodel.LB) {
	defer log.Info(updateToolMap(RESOURCE_TYPE_LB_EN, cloudItem.Lcuuid))
	id, _ := t.GetLBIDByLcuuid(cloudItem.Lcuuid)
	info, err := t.GetLBInfoByID(id)
	if err != nil {
		log.Error(err)
		return
	}
	info.Name = cloudItem.Name
	if regionID, ok := t.GetRegionIDByLcuuid(cloudItem.RegionLcuuid); ok {
		info.RegionID = regionID
	}
	if vpcID, ok := t.VPCLcuuidToID[cloudItem.Lcuuid]; ok {
		info.VPCID = vpcID
	}
}

func (t *ToolDataSet) deleteLB(lcuuid string) {
	id, _ := t.GetLBIDByLcuuid(lcuuid)
	delete(t.lbIDToInfo, id)
	delete(t.LBLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_LB_EN, lcuuid))
}

func (t *ToolDataSet) addLBListener(item *mysql.LBListener) {
	t.LBListenerLcuuidToID[item.Lcuuid] = item.ID
	log.Info(addToToolMap(RESOURCE_TYPE_LB_LISTENER_EN, item.Lcuuid))
}

func (t *ToolDataSet) deleteLBListener(lcuuid string) {
	delete(t.LBListenerLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_LB_LISTENER_EN, lcuuid))
}

func (t *ToolDataSet) addRDSInstance(item *mysql.RDSInstance) {
	t.RDSInstanceLcuuidToID[item.Lcuuid] = item.ID
	t.rdsInstanceIDToInfo[item.ID] = &rdsInstanceInfo{
		Name:  item.Name,
		VPCID: item.VPCID,
	}
	if regionID, ok := t.GetRegionIDByLcuuid(item.Region); ok {
		t.rdsInstanceIDToInfo[item.ID].RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(item.AZ); ok {
		t.rdsInstanceIDToInfo[item.ID].AZID = azID
	}
	log.Info(addToToolMap(RESOURCE_TYPE_RDS_INSTANCE_EN, item.Lcuuid))
}

func (t *ToolDataSet) updateRDSInstance(cloudItem *cloudmodel.RDSInstance) {
	defer log.Info(updateToolMap(RESOURCE_TYPE_RDS_INSTANCE_EN, cloudItem.Lcuuid))
	id, _ := t.GetRDSInstanceIDByLcuuid(cloudItem.Lcuuid)
	info, err := t.GetRDSInstanceInfoByID(id)
	if err != nil {
		log.Error(err)
		return
	}
	info.Name = cloudItem.Name
	if regionID, ok := t.GetRegionIDByLcuuid(cloudItem.RegionLcuuid); ok {
		info.RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(cloudItem.AZLcuuid); ok {
		info.AZID = azID
	}
	if vpcID, ok := t.VPCLcuuidToID[cloudItem.Lcuuid]; ok {
		info.VPCID = vpcID
	}
}

func (t *ToolDataSet) deleteRDSInstance(lcuuid string) {
	id, _ := t.GetRDSInstanceIDByLcuuid(lcuuid)
	delete(t.rdsInstanceIDToInfo, id)
	delete(t.RDSInstanceLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_RDS_INSTANCE_EN, lcuuid))
}

func (t *ToolDataSet) addRedisInstance(item *mysql.RedisInstance) {
	t.RedisInstanceLcuuidToID[item.Lcuuid] = item.ID
	t.redisInstanceIDToInfo[item.ID] = &redisInstanceInfo{
		Name:  item.Name,
		VPCID: item.VPCID,
	}
	if regionID, ok := t.GetRegionIDByLcuuid(item.Region); ok {
		t.redisInstanceIDToInfo[item.ID].RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(item.AZ); ok {
		t.redisInstanceIDToInfo[item.ID].AZID = azID
	}
	log.Info(addToToolMap(RESOURCE_TYPE_REDIS_INSTANCE_EN, item.Lcuuid))
}

func (t *ToolDataSet) updateRedisInstance(cloudItem *cloudmodel.RedisInstance) {
	defer log.Info(updateToolMap(RESOURCE_TYPE_REDIS_INSTANCE_EN, cloudItem.Lcuuid))
	id, _ := t.GetRedisInstanceIDByLcuuid(cloudItem.Lcuuid)
	info, err := t.GetRedisInstanceInfoByID(id)
	if err != nil {
		log.Error(err)
		return
	}
	info.Name = cloudItem.Name
	if regionID, ok := t.GetRegionIDByLcuuid(cloudItem.RegionLcuuid); ok {
		info.RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(cloudItem.AZLcuuid); ok {
		info.AZID = azID
	}
	if vpcID, ok := t.VPCLcuuidToID[cloudItem.Lcuuid]; ok {
		info.VPCID = vpcID
	}
}

func (t *ToolDataSet) deleteRedisInstance(lcuuid string) {
	id, _ := t.GetRedisInstanceIDByLcuuid(lcuuid)
	delete(t.redisInstanceIDToInfo, id)
	delete(t.RedisInstanceLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_REDIS_INSTANCE_EN, lcuuid))
}

func (t *ToolDataSet) addPodCluster(item *mysql.PodCluster) {
	t.PodClusterLcuuidToID[item.Lcuuid] = item.ID
	log.Info(addToToolMap(RESOURCE_TYPE_POD_CLUSTER_EN, item.Lcuuid))
}

func (t *ToolDataSet) deletePodCluster(lcuuid string) {
	delete(t.PodClusterLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_POD_CLUSTER_EN, lcuuid))
}

func (t *ToolDataSet) addPodNode(item *mysql.PodNode) {
	t.PodNodeLcuuidToID[item.Lcuuid] = item.ID
	t.PodNodeIDToLcuuid[item.ID] = item.Lcuuid
	t.podNodeIDToInfo[item.ID] = &podNodeInfo{
		Name:         item.Name,
		VPCID:        item.VPCID,
		PodClusterID: item.PodClusterID,
	}
	if regionID, ok := t.GetRegionIDByLcuuid(item.Region); ok {
		t.podNodeIDToInfo[item.ID].RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(item.AZ); ok {
		t.podNodeIDToInfo[item.ID].AZID = azID
	}
	log.Info(addToToolMap(RESOURCE_TYPE_POD_NODE_EN, item.Lcuuid))
}

func (t *ToolDataSet) updatePodNode(cloudItem *cloudmodel.PodNode) {
	defer log.Info(updateToolMap(RESOURCE_TYPE_POD_NODE_EN, cloudItem.Lcuuid))
	id, _ := t.GetPodNodeIDByLcuuid(cloudItem.Lcuuid)
	info, err := t.GetPodNodeInfoByID(id)
	if err != nil {
		log.Error(err)
		return
	}
	info.Name = cloudItem.Name
	if regionID, ok := t.GetRegionIDByLcuuid(cloudItem.RegionLcuuid); ok {
		info.RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(cloudItem.AZLcuuid); ok {
		info.AZID = azID
	}
	if vpcID, ok := t.VPCLcuuidToID[cloudItem.Lcuuid]; ok {
		info.VPCID = vpcID
	}
	if podClusterLcuuid, ok := t.GetPodClusterIDByLcuuid(cloudItem.PodClusterLcuuid); ok {
		info.PodClusterID = podClusterLcuuid
	}
}

func (t *ToolDataSet) deletePodNode(lcuuid string) {
	id, _ := t.GetPodNodeIDByLcuuid(lcuuid)
	delete(t.PodNodeIDToLcuuid, id)
	delete(t.podNodeIDToInfo, id)
	delete(t.PodNodeLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_POD_NODE_EN, lcuuid))
}

func (t *ToolDataSet) addPodNamespace(item *mysql.PodNamespace) {
	t.PodNamespaceLcuuidToID[item.Lcuuid] = item.ID
	log.Info(addToToolMap(RESOURCE_TYPE_POD_NAMESPACE_EN, item.Lcuuid))
}

func (t *ToolDataSet) deletePodNamespace(lcuuid string) {
	delete(t.PodNamespaceLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_POD_NAMESPACE_EN, lcuuid))
}

func (t *ToolDataSet) addPodIngress(item *mysql.PodIngress) {
	t.PodIngressLcuuidToID[item.Lcuuid] = item.ID
	t.PodIngressIDToLcuuid[item.ID] = item.Lcuuid
	log.Info(addToToolMap(RESOURCE_TYPE_POD_INGRESS_EN, item.Lcuuid))
}

func (t *ToolDataSet) deletePodIngress(lcuuid string) {
	id, _ := t.GetPodIngressIDByLcuuid(lcuuid)
	delete(t.PodIngressIDToLcuuid, id)
	delete(t.PodIngressLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_POD_INGRESS_EN, lcuuid))
}

func (t *ToolDataSet) addPodIngressRule(item *mysql.PodIngressRule) {
	t.PodIngressRuleLcuuidToID[item.Lcuuid] = item.ID
	log.Info(addToToolMap(RESOURCE_TYPE_POD_INGRESS_RULE_EN, item.Lcuuid))
}

func (t *ToolDataSet) deletePodIngressRule(lcuuid string) {
	delete(t.PodIngressRuleLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_POD_INGRESS_RULE_EN, lcuuid))
}

func (t *ToolDataSet) addPodService(item *mysql.PodService) {
	t.PodServiceLcuuidToID[item.Lcuuid] = item.ID
	t.podServiceIDToInfo[item.ID] = &podServiceInfo{
		Name:         item.Name,
		VPCID:        item.VPCID,
		PodClusterID: item.PodClusterID,
		PodNSID:      item.PodNamespaceID,
	}
	if regionID, ok := t.GetRegionIDByLcuuid(item.Region); ok {
		t.podServiceIDToInfo[item.ID].RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(item.AZ); ok {
		t.podServiceIDToInfo[item.ID].AZID = azID
	}
	log.Info(addToToolMap(RESOURCE_TYPE_POD_SERVICE_EN, item.Lcuuid))
}

func (t *ToolDataSet) updatePodService(cloudItem *cloudmodel.PodService) {
	defer log.Info(updateToolMap(RESOURCE_TYPE_POD_SERVICE_EN, cloudItem.Lcuuid))
	id, _ := t.GetPodServiceIDByLcuuid(cloudItem.Lcuuid)
	info, err := t.GetPodServiceInfoByID(id)
	if err != nil {
		log.Error(err)
		return
	}
	info.Name = cloudItem.Name
	if regionID, ok := t.GetRegionIDByLcuuid(cloudItem.RegionLcuuid); ok {
		info.RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(cloudItem.AZLcuuid); ok {
		info.AZID = azID
	}
	if vpcID, ok := t.VPCLcuuidToID[cloudItem.Lcuuid]; ok {
		info.VPCID = vpcID
	}
	if podClusterLcuuid, ok := t.GetPodClusterIDByLcuuid(cloudItem.PodClusterLcuuid); ok {
		info.PodClusterID = podClusterLcuuid
	}
	if podNSID, ok := t.GetPodNamespaceIDByLcuuid(cloudItem.PodNamespaceLcuuid); ok {
		info.PodNSID = podNSID
	}
}

func (t *ToolDataSet) deletePodService(lcuuid string) {
	id, _ := t.GetPodServiceIDByLcuuid(lcuuid)
	delete(t.podServiceIDToInfo, id)
	delete(t.PodServiceLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_POD_SERVICE_EN, lcuuid))
}

func (t *ToolDataSet) addPodGroup(item *mysql.PodGroup) {
	t.PodGroupLcuuidToID[item.Lcuuid] = item.ID
	log.Info(addToToolMap(RESOURCE_TYPE_POD_GROUP_EN, item.Lcuuid))
}

func (t *ToolDataSet) deletePodGroup(lcuuid string) {
	delete(t.PodGroupLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_POD_GROUP_EN, lcuuid))
}

func (f *ToolDataSet) addPodReplicaSet(item *mysql.PodReplicaSet) {
	f.PodReplicaSetLcuuidToID[item.Lcuuid] = item.ID
	f.PodReplicaSetIDToLcuuid[item.ID] = item.Lcuuid
	log.Info(addToToolMap(RESOURCE_TYPE_POD_REPLICA_SET_EN, item.Lcuuid))
}

func (f *ToolDataSet) deletePodReplicaSet(lcuuid string) {
	id, exists := f.GetPodReplicaSetIDByLcuuid(lcuuid)
	if exists {
		delete(f.PodReplicaSetIDToLcuuid, id)
	}
	delete(f.PodReplicaSetLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_POD_REPLICA_SET_EN, lcuuid))
}

func (t *ToolDataSet) addPod(item *mysql.Pod) {
	t.PodLcuuidToID[item.Lcuuid] = item.ID
	t.podIDToInfo[item.ID] = &podInfo{
		Name:         item.Name,
		VPCID:        item.VPCID,
		PodClusterID: item.PodClusterID,
		PodNSID:      item.PodNamespaceID,
		PodGroupID:   item.PodGroupID,
		PodNodeID:    item.PodNodeID,
	}
	if regionID, ok := t.GetRegionIDByLcuuid(item.Region); ok {
		t.podIDToInfo[item.ID].RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(item.AZ); ok {
		t.podIDToInfo[item.ID].AZID = azID
	}
	log.Info(addToToolMap(RESOURCE_TYPE_POD_EN, item.Lcuuid))
}

func (t *ToolDataSet) updatePod(cloudItem *cloudmodel.Pod) {
	defer log.Info(updateToolMap(RESOURCE_TYPE_POD_EN, cloudItem.Lcuuid))
	id, _ := t.GetPodIDByLcuuid(cloudItem.Lcuuid)
	info, err := t.GetPodInfoeByID(id)
	if err != nil {
		log.Error(err)
		return
	}
	info.Name = cloudItem.Name
	if regionID, ok := t.GetRegionIDByLcuuid(cloudItem.RegionLcuuid); ok {
		info.RegionID = regionID
	}
	if azID, ok := t.GetAZIDByLcuuid(cloudItem.AZLcuuid); ok {
		info.AZID = azID
	}
	if id, ok := t.VPCLcuuidToID[cloudItem.Lcuuid]; ok {
		info.VPCID = id
	}
	if id, ok := t.GetPodClusterIDByLcuuid(cloudItem.PodClusterLcuuid); ok {
		info.PodClusterID = id
	}
	if id, ok := t.GetPodNamespaceIDByLcuuid(cloudItem.PodNamespaceLcuuid); ok {
		info.PodNSID = id
	}
	if id, ok := t.GetPodNodeIDByLcuuid(cloudItem.PodNodeLcuuid); ok {
		info.PodNodeID = id
	}
	if id, ok := t.GetPodGroupIDByLcuuid(cloudItem.PodGroupLcuuid); ok {
		info.PodGroupID = id
	}
}

func (t *ToolDataSet) deletePod(lcuuid string) {
	id, _ := t.GetPodIDByLcuuid(lcuuid)
	delete(t.podIDToInfo, id)
	delete(t.PodIDToIPNetworkIDMap, id)
	delete(t.PodLcuuidToID, lcuuid)
	log.Info(deleteFromToolMap(RESOURCE_TYPE_POD_EN, lcuuid))
}

func (t *ToolDataSet) GetRegionIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.RegionLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_REGION_EN, lcuuid))
	var region mysql.Region
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&region)
	if result.RowsAffected == 1 {
		t.addRegion(&region)
		return region.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_REGION_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetRegionLcuuidByID(id int) (string, bool) {
	lcuuid, exists := t.RegionIDToLcuuid[id]
	if exists {
		return lcuuid, true
	}
	log.Warning(cacheLcuuidByIDNotFound(RESOURCE_TYPE_REGION_EN, id))
	var region mysql.Region
	result := mysql.Db.Where("id = ?", id).Find(&region)
	if result.RowsAffected == 1 {
		t.addRegion(&region)
		return region.Lcuuid, true
	} else {
		log.Error(dbResourceByIDNotFound(RESOURCE_TYPE_REGION_EN, id))
		return lcuuid, false
	}
}

func (t *ToolDataSet) GetHostIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.HostLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_HOST_EN, lcuuid))
	var host mysql.Host
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&host)
	if result.RowsAffected == 1 {
		t.addHost(&host)
		return host.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_HOST_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetHostIDByIP(ip string) (int, bool) {
	id, exists := t.hostIPToID[ip]
	if exists {
		return id, true
	}
	log.Warningf("cache %s id (ip: %s) not found", RESOURCE_TYPE_HOST_EN, ip)
	var host mysql.Host
	result := mysql.Db.Where("ip = ?", ip).Find(&host)
	if result.RowsAffected == 1 {
		t.addHost(&host)
		return host.ID, true
	} else {
		log.Error("db %s (ip: %s) not found", RESOURCE_TYPE_HOST_EN, ip)
		return id, false
	}
}

func (t *ToolDataSet) GetVMIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.VMLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_VM_EN, lcuuid))
	var vm mysql.VM
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&vm)
	if result.RowsAffected == 1 {
		t.addVM(&vm)
		return vm.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_VM_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetVPCIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.VPCLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_VPC_EN, lcuuid))
	var vpc mysql.VPC
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&vpc)
	if result.RowsAffected == 1 {
		t.addVPC(&vpc)
		return vpc.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_VPC_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetVPCLcuuidByID(id int) (string, bool) {
	lcuuid, exists := t.VPCIDToLcuuid[id]
	if exists {
		return lcuuid, true
	}
	log.Warning(cacheLcuuidByIDNotFound(RESOURCE_TYPE_VPC_EN, id))
	var vpc mysql.VPC
	result := mysql.Db.Where("lcuuid = ?", id).Find(&vpc)
	if result.RowsAffected == 1 {
		t.addVPC(&vpc)
		return vpc.Lcuuid, true
	} else {
		log.Error(dbResourceByIDNotFound(RESOURCE_TYPE_VPC_EN, id))
		return lcuuid, false
	}
}

func (t *ToolDataSet) GetNetworkIDByLcuuid(lcuuid string) (int, bool) {
	if lcuuid == PUBLIC_NETWORK_LCUUID {
		return t.PublicNetworkID, true
	}
	id, exists := t.NetworkLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_NETWORK_EN, lcuuid))
	var network mysql.Network
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&network)
	if result.RowsAffected == 1 {
		t.addNetwork(&network)
		return network.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_NETWORK_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetNetworkIDByVInterfaceLcuuid(vifLcuuid string) (int, bool) {
	id, exists := t.VInterfaceLcuuidToNetworkID[vifLcuuid]
	if exists {
		return id, true
	}
	log.Warningf("cache %s id (%s lcuuid: %s) not found", RESOURCE_TYPE_NETWORK_EN, RESOURCE_TYPE_VINTERFACE_EN, vifLcuuid)
	var vif mysql.VInterface
	result := mysql.Db.Where("lcuuid = ?", vifLcuuid).Find(&vif)
	if result.RowsAffected == 1 {
		t.addVInterface(&vif)
		return vif.NetworkID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_VINTERFACE_EN, vifLcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetDeviceTypeByVInterfaceLcuuid(vifLcuuid string) (int, bool) {
	id, exists := t.VInterfaceLcuuidToDeviceType[vifLcuuid]
	if exists {
		return id, true
	}
	log.Warningf("cache device type (%s lcuuid: %s) not found", RESOURCE_TYPE_VINTERFACE_EN, vifLcuuid)
	var vif mysql.VInterface
	result := mysql.Db.Where("lcuuid = ?", vifLcuuid).Find(&vif)
	if result.RowsAffected == 1 {
		t.addVInterface(&vif)
		return vif.DeviceType, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_VINTERFACE_EN, vifLcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetDeviceIDByVInterfaceLcuuid(vifLcuuid string) (int, bool) {
	id, exists := t.VInterfaceLcuuidToDeviceID[vifLcuuid]
	if exists {
		return id, true
	}
	log.Warningf("cache device id (%s lcuuid: %s) not found", RESOURCE_TYPE_VINTERFACE_EN, vifLcuuid)
	var vif mysql.VInterface
	result := mysql.Db.Where("lcuuid = ?", vifLcuuid).Find(&vif)
	if result.RowsAffected == 1 {
		t.addVInterface(&vif)
		return vif.DeviceID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_VINTERFACE_EN, vifLcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetNetworkLcuuidByID(id int) (string, bool) {
	if id == t.PublicNetworkID {
		return PUBLIC_NETWORK_LCUUID, true
	}
	lcuuid, exists := t.NetworkIDToLcuuid[id]
	if exists {
		return lcuuid, true
	}
	log.Warning(cacheLcuuidByIDNotFound(RESOURCE_TYPE_NETWORK_EN, id))
	var network mysql.Network
	result := mysql.Db.Where("lcuuid = ?", id).Find(&network)
	if result.RowsAffected == 1 {
		t.addNetwork(&network)
		return network.Lcuuid, true
	} else {
		log.Error(dbResourceByIDNotFound(RESOURCE_TYPE_NETWORK_EN, id))
		return lcuuid, false
	}
}

func (t *ToolDataSet) GetVRouterIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.VRouterLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_VROUTER_EN, lcuuid))
	var vrouter mysql.VRouter
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&vrouter)
	if result.RowsAffected == 1 {
		t.addVRouter(&vrouter)
		return vrouter.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_VROUTER_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetDHCPPortIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.DHCPPortLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_DHCP_PORT_EN, lcuuid))
	var dhcpPort mysql.DHCPPort
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&dhcpPort)
	if result.RowsAffected == 1 {
		t.addDHCPPort(&dhcpPort)
		return dhcpPort.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_DHCP_PORT_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetVInterfaceIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.VInterfaceLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_VINTERFACE_EN, lcuuid))
	var vinterface mysql.VInterface
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&vinterface)
	if result.RowsAffected == 1 {
		t.addVInterface(&vinterface)
		return vinterface.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_VINTERFACE_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetVInterfaceTypeByLcuuid(lcuuid string) (int, bool) {
	vt, exists := t.VInterfaceLcuuidToType[lcuuid]
	if exists {
		return vt, true
	}
	log.Warningf("cache %s type (lcuuid: %s) not found", RESOURCE_TYPE_VINTERFACE_EN, lcuuid)
	var vinterface mysql.VInterface
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&vinterface)
	if result.RowsAffected == 1 {
		t.addVInterface(&vinterface)
		return vinterface.Type, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_VINTERFACE_EN, lcuuid))
		return vt, false
	}
}

func (t *ToolDataSet) GetDeviceIDByDeviceLcuuid(deviceType int, deviceLcuuid string) (int, bool) {
	if deviceType == common.VIF_DEVICE_TYPE_HOST {
		return t.GetHostIDByLcuuid(deviceLcuuid)
	} else if deviceType == common.VIF_DEVICE_TYPE_VM {
		return t.GetVMIDByLcuuid(deviceLcuuid)
	} else if deviceType == common.VIF_DEVICE_TYPE_VROUTER {
		return t.GetVRouterIDByLcuuid(deviceLcuuid)
	} else if deviceType == common.VIF_DEVICE_TYPE_DHCP_PORT {
		return t.GetDHCPPortIDByLcuuid(deviceLcuuid)
	} else if deviceType == common.VIF_DEVICE_TYPE_NAT_GATEWAY {
		return t.GetNATGatewayIDByLcuuid(deviceLcuuid)
	} else if deviceType == common.VIF_DEVICE_TYPE_LB {
		return t.GetLBIDByLcuuid(deviceLcuuid)
	} else if deviceType == common.VIF_DEVICE_TYPE_RDS_INSTANCE {
		return t.GetRDSInstanceIDByLcuuid(deviceLcuuid)
	} else if deviceType == common.VIF_DEVICE_TYPE_REDIS_INSTANCE {
		return t.GetRedisInstanceIDByLcuuid(deviceLcuuid)
	} else if deviceType == common.VIF_DEVICE_TYPE_POD_NODE {
		return t.GetPodNodeIDByLcuuid(deviceLcuuid)
	} else if deviceType == common.VIF_DEVICE_TYPE_POD_SERVICE {
		return t.GetPodServiceIDByLcuuid(deviceLcuuid)
	} else if deviceType == common.VIF_DEVICE_TYPE_POD {
		return t.GetPodIDByLcuuid(deviceLcuuid)
	} else {
		log.Errorf("device type %d not supported", deviceType)
		return 0, false
	}
}

func (t *ToolDataSet) GetDeviceNameByDeviceID(deviceType, deviceID int) (string, error) {
	if deviceType == common.VIF_DEVICE_TYPE_HOST {
		return t.GetHostNameByID(deviceID)
	} else if deviceType == common.VIF_DEVICE_TYPE_VM {
		return t.GetVMNameByID(deviceID)
	} else if deviceType == common.VIF_DEVICE_TYPE_VROUTER {
		return t.GetVRouterNameByID(deviceID)
	} else if deviceType == common.VIF_DEVICE_TYPE_DHCP_PORT {
		return t.GetDHCPPortNameByID(deviceID)
	} else if deviceType == common.VIF_DEVICE_TYPE_NAT_GATEWAY {
		return t.GetNATGatewayNameByID(deviceID)
	} else if deviceType == common.VIF_DEVICE_TYPE_LB {
		return t.GetLBNameByID(deviceID)
	} else if deviceType == common.VIF_DEVICE_TYPE_RDS_INSTANCE {
		return t.GetRDSInstanceNameByID(deviceID)
	} else if deviceType == common.VIF_DEVICE_TYPE_REDIS_INSTANCE {
		return t.GetRedisInstanceNameByID(deviceID)
	} else if deviceType == common.VIF_DEVICE_TYPE_POD_NODE {
		return t.GetPodNodeNameByID(deviceID)
	} else if deviceType == common.VIF_DEVICE_TYPE_POD_SERVICE {
		return t.GetPodServiceNameByID(deviceID)
	} else if deviceType == common.VIF_DEVICE_TYPE_POD {
		return t.GetPodNameByID(deviceID)
	} else {
		return "", fmt.Errorf("device type %d not supported", deviceType)
	}
}

func (t *ToolDataSet) GetSecurityGroupIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.SecurityGroupLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_SECURITY_GROUP_EN, lcuuid))
	var securityGroup mysql.SecurityGroup
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&securityGroup)
	if result.RowsAffected == 1 {
		t.addSecurityGroup(&securityGroup)
		return securityGroup.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_SECURITY_GROUP_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetNATGatewayIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.NATGatewayLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_NAT_GATEWAY_EN, lcuuid))
	var natGateway mysql.NATGateway
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&natGateway)
	if result.RowsAffected == 1 {
		t.addNATGateway(&natGateway)
		return natGateway.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_NAT_GATEWAY_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetLBIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.LBLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_LB_EN, lcuuid))
	var lb mysql.LB
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&lb)
	if result.RowsAffected == 1 {
		t.addLB(&lb)
		return lb.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_LB_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetLBListenerIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.LBListenerLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_LB_LISTENER_EN, lcuuid))
	var lbListener mysql.LBListener
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&lbListener)
	if result.RowsAffected == 1 {
		t.addLBListener(&lbListener)
		return lbListener.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_LB_LISTENER_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetRDSInstanceIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.RDSInstanceLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_RDS_INSTANCE_EN, lcuuid))
	var rdsInstance mysql.RDSInstance
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&rdsInstance)
	if result.RowsAffected == 1 {
		t.addRDSInstance(&rdsInstance)
		return rdsInstance.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_RDS_INSTANCE_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetRedisInstanceIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.RedisInstanceLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_REDIS_INSTANCE_EN, lcuuid))
	var redisInstance mysql.RedisInstance
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&redisInstance)
	if result.RowsAffected == 1 {
		t.addRedisInstance(&redisInstance)
		return redisInstance.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_REDIS_INSTANCE_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetPodClusterIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.PodClusterLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_POD_CLUSTER_EN, lcuuid))
	var podCluster mysql.PodCluster
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&podCluster)
	if result.RowsAffected == 1 {
		t.addPodCluster(&podCluster)
		return podCluster.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_POD_CLUSTER_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetPodNodeIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.PodNodeLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_POD_NODE_EN, lcuuid))
	var podNode mysql.PodNode
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&podNode)
	if result.RowsAffected == 1 {
		t.addPodNode(&podNode)
		return podNode.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_POD_NODE_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetPodNodeLcuuidByID(id int) (string, bool) {
	lcuuid, exists := t.PodNodeIDToLcuuid[id]
	if exists {
		return lcuuid, true
	}
	log.Warning(cacheLcuuidByIDNotFound(RESOURCE_TYPE_POD_NODE_EN, id))
	var podNode mysql.PodNode
	result := mysql.Db.Where("id = ?", id).Find(&podNode)
	if result.RowsAffected == 1 {
		t.addPodNode(&podNode)
		return podNode.Lcuuid, true
	} else {
		log.Error(dbResourceByIDNotFound(RESOURCE_TYPE_POD_NODE_EN, id))
		return lcuuid, false
	}
}

func (t *ToolDataSet) GetPodNamespaceIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.PodNamespaceLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_POD_NAMESPACE_EN, lcuuid))
	var podNamespace mysql.PodNamespace
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&podNamespace)
	if result.RowsAffected == 1 {
		t.addPodNamespace(&podNamespace)
		return podNamespace.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_POD_NAMESPACE_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetPodIngressIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.PodIngressLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_POD_INGRESS_EN, lcuuid))
	var podIngress mysql.PodIngress
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&podIngress)
	if result.RowsAffected == 1 {
		t.addPodIngress(&podIngress)
		return podIngress.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_POD_INGRESS_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetPodIngressLcuuidByID(id int) (string, bool) {
	lcuuid, exists := t.PodIngressIDToLcuuid[id]
	if exists {
		return lcuuid, true
	}
	log.Warning(cacheLcuuidByIDNotFound(RESOURCE_TYPE_POD_INGRESS_EN, id))
	var podIngress mysql.PodIngress
	result := mysql.Db.Where("id = ?", id).Find(&podIngress)
	if result.RowsAffected == 1 {
		t.addPodIngress(&podIngress)
		return podIngress.Lcuuid, true
	} else {
		log.Error(dbResourceByIDNotFound(RESOURCE_TYPE_POD_INGRESS_EN, id))
		return lcuuid, false
	}
}

func (t *ToolDataSet) GetPodIngressRuleIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.PodIngressRuleLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_POD_INGRESS_RULE_EN, lcuuid))
	var podIngressRule mysql.PodIngressRule
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&podIngressRule)
	if result.RowsAffected == 1 {
		t.addPodIngressRule(&podIngressRule)
		return podIngressRule.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_POD_INGRESS_RULE_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetPodServiceIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.PodServiceLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_POD_SERVICE_EN, lcuuid))
	var podService mysql.PodService
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&podService)
	if result.RowsAffected == 1 {
		t.addPodService(&podService)
		return podService.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_POD_SERVICE_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetPodGroupIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.PodGroupLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_POD_GROUP_EN, lcuuid))
	var podGroup mysql.PodGroup
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&podGroup)
	if result.RowsAffected == 1 {
		t.addPodGroup(&podGroup)
		return podGroup.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_POD_GROUP_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetPodReplicaSetIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.PodReplicaSetLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_POD_REPLICA_SET_EN, lcuuid))
	var podReplicaSet mysql.PodReplicaSet
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&podReplicaSet)
	if result.RowsAffected == 1 {
		t.addPodReplicaSet(&podReplicaSet)
		return podReplicaSet.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_POD_REPLICA_SET_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetPodReplicaSetLcuuidByID(id int) (string, bool) {
	lcuuid, exists := t.PodReplicaSetIDToLcuuid[id]
	if exists {
		return lcuuid, true
	}
	log.Warning(cacheLcuuidByIDNotFound(RESOURCE_TYPE_POD_REPLICA_SET_EN, id))
	var podReplicaSet mysql.PodReplicaSet
	result := mysql.Db.Where("id = ?", id).Find(&podReplicaSet)
	if result.RowsAffected == 1 {
		t.addPodReplicaSet(&podReplicaSet)
		return podReplicaSet.Lcuuid, true
	} else {
		log.Error(dbResourceByIDNotFound(RESOURCE_TYPE_POD_REPLICA_SET_EN, id))
		return lcuuid, false
	}
}

func (t *ToolDataSet) GetPodIDByLcuuid(lcuuid string) (int, bool) {
	id, exists := t.PodLcuuidToID[lcuuid]
	if exists {
		return id, true
	}
	log.Warning(cacheIDByLcuuidNotFound(RESOURCE_TYPE_POD_EN, lcuuid))
	var pod mysql.Pod
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&pod)
	if result.RowsAffected == 1 {
		t.addPod(&pod)
		return pod.ID, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_POD_EN, lcuuid))
		return id, false
	}
}

func (t *ToolDataSet) GetHostInfoByID(id int) (*hostInfo, error) {
	info, exists := t.hostIDtoInfo[id]
	if exists {
		return info, nil
	}

	log.Warning(cacheNameByIDNotFound(RESOURCE_TYPE_HOST_EN, id))
	var dbItem mysql.Host
	result := mysql.Db.Where("id = ?", id).Find(&dbItem)
	if result.RowsAffected == 1 {
		t.addHost(&dbItem)
		return t.hostIDtoInfo[dbItem.ID], nil
	}
	return nil, errors.New(dbResourceByIDNotFound(RESOURCE_TYPE_HOST_EN, id))
}

func (t *ToolDataSet) GetHostNameByID(id int) (string, error) {
	info, err := t.GetHostInfoByID(id)
	if err != nil {
		return "", err
	}
	return info.Name, nil
}

func (t *ToolDataSet) GetVMInfoByID(id int) (*vmInfo, error) {
	info, exists := t.vmIDToInfo[id]
	if exists {
		return info, nil
	}

	log.Warning(cacheNameByIDNotFound(RESOURCE_TYPE_VM_EN, id))
	var dbItem mysql.VM
	result := mysql.Db.Where("id = ?", id).Find(&dbItem)
	if result.RowsAffected == 1 {
		t.addVM(&dbItem)
		return t.vmIDToInfo[dbItem.ID], nil
	}
	return nil, errors.New(dbResourceByIDNotFound(RESOURCE_TYPE_VM_EN, id))
}

func (t *ToolDataSet) GetVMNameByID(id int) (string, error) {
	info, err := t.GetVMInfoByID(id)
	if err != nil {
		return "", err
	}
	return info.Name, nil
}

func (t *ToolDataSet) GetNetworkNameByID(id int) (string, bool) {
	name, exists := t.NetworkIDToName[id]
	if exists {
		return name, true
	}
	log.Warning(cacheNameByIDNotFound(RESOURCE_TYPE_NETWORK_EN, id))
	var network mysql.Network
	result := mysql.Db.Where("id = ?", id).Find(&network)
	if result.RowsAffected == 1 {
		t.addNetwork(&network)
		return network.Name, true
	} else {
		log.Error(dbResourceByIDNotFound(RESOURCE_TYPE_NETWORK_EN, id))
		return name, false
	}
}

func (t *ToolDataSet) GetVRouterInfoByID(id int) (*vrouterInfo, error) {
	info, exists := t.vrouterIDToInfo[id]
	if exists {
		return info, nil
	}
	log.Warning(cacheNameByIDNotFound(RESOURCE_TYPE_VROUTER_EN, id))

	var vRouter mysql.VRouter
	result := mysql.Db.Where("id = ?", id).Find(&vRouter)
	if result.RowsAffected == 1 {
		t.addVRouter(&vRouter)
		return t.vrouterIDToInfo[vRouter.ID], nil
	}
	return nil, errors.New(dbResourceByIDNotFound(RESOURCE_TYPE_VROUTER_EN, id))
}

func (t *ToolDataSet) GetVRouterNameByID(id int) (string, error) {
	info, err := t.GetVRouterInfoByID(id)
	if err != nil {
		return "", err
	}
	return info.Name, nil
}

func (t *ToolDataSet) GetDHCPPortInfoByID(id int) (*dhcpPortInfo, error) {
	info, exists := t.dhcpPortIDToInfo[id]
	if exists {
		return info, nil
	}
	log.Warning(cacheNameByIDNotFound(RESOURCE_TYPE_DHCP_PORT_EN, id))
	var dbItem mysql.DHCPPort
	result := mysql.Db.Where("id = ?", id).Find(&dbItem)
	if result.RowsAffected == 1 {
		t.addDHCPPort(&dbItem)
		return t.dhcpPortIDToInfo[id], nil
	}
	return nil, errors.New(dbResourceByIDNotFound(RESOURCE_TYPE_DHCP_PORT_EN, id))
}

func (t *ToolDataSet) GetDHCPPortNameByID(id int) (string, error) {
	info, err := t.GetDHCPPortInfoByID(id)
	if err != nil {
		return "", err
	}
	return info.Name, nil
}

func (t *ToolDataSet) GetLBInfoByID(id int) (*lbInfo, error) {
	info, exists := t.lbIDToInfo[id]
	if exists {
		return info, nil
	}

	log.Warning(cacheNameByIDNotFound(RESOURCE_TYPE_LB_EN, id))
	var dbItem mysql.LB
	result := mysql.Db.Where("id = ?", id).Find(&dbItem)
	if result.RowsAffected == 1 {
		t.addLB(&dbItem)
		return t.lbIDToInfo[dbItem.ID], nil
	}
	return nil, errors.New(dbResourceByIDNotFound(RESOURCE_TYPE_LB_EN, id))
}

func (t *ToolDataSet) GetLBNameByID(id int) (string, error) {
	info, err := t.GetLBInfoByID(id)
	if err != nil {
		return "", err
	}
	return info.Name, nil
}

func (t *ToolDataSet) GetNATGatewayInfoByID(id int) (*natGatewayInfo, error) {
	info, exists := t.natGatewayIDToInfo[id]
	if exists {
		return info, nil
	}
	log.Warning(cacheNameByIDNotFound(RESOURCE_TYPE_NAT_GATEWAY_EN, id))

	var dbItem mysql.NATGateway
	result := mysql.Db.Where("id = ?", id).Find(&dbItem)
	if result.RowsAffected == 1 {
		t.addNATGateway(&dbItem)
		return t.natGatewayIDToInfo[dbItem.ID], nil
	}
	return nil, errors.New(dbResourceByIDNotFound(RESOURCE_TYPE_NAT_GATEWAY_EN, id))
}

func (t *ToolDataSet) GetNATGatewayNameByID(id int) (string, error) {
	info, err := t.GetNATGatewayInfoByID(id)
	if err != nil {
		return "", err
	}
	return info.Name, nil
}

func (t *ToolDataSet) GetRDSInstanceInfoByID(id int) (*rdsInstanceInfo, error) {
	info, exists := t.rdsInstanceIDToInfo[id]
	if exists {
		return info, nil
	}

	log.Warning(cacheNameByIDNotFound(RESOURCE_TYPE_RDS_INSTANCE_EN, id))
	var dbItem mysql.RDSInstance
	result := mysql.Db.Where("id = ?", id).Find(&dbItem)
	if result.RowsAffected == 1 {
		t.addRDSInstance(&dbItem)
		return t.rdsInstanceIDToInfo[dbItem.ID], nil
	}
	return nil, errors.New(dbResourceByIDNotFound(RESOURCE_TYPE_RDS_INSTANCE_EN, id))
}

func (t *ToolDataSet) GetRDSInstanceNameByID(id int) (string, error) {
	info, err := t.GetRDSInstanceInfoByID(id)
	if err != nil {
		return "", nil
	}
	return info.Name, nil
}

func (t *ToolDataSet) GetRedisInstanceInfoByID(id int) (*redisInstanceInfo, error) {
	info, exists := t.redisInstanceIDToInfo[id]
	if exists {
		return info, nil
	}

	log.Warning(cacheNameByIDNotFound(RESOURCE_TYPE_REDIS_INSTANCE_EN, id))
	var dbItem mysql.RedisInstance
	result := mysql.Db.Where("id = ?", id).Find(&dbItem)
	if result.RowsAffected == 1 {
		t.addRedisInstance(&dbItem)
		return t.redisInstanceIDToInfo[dbItem.ID], nil
	}
	return nil, errors.New(dbResourceByIDNotFound(RESOURCE_TYPE_REDIS_INSTANCE_EN, id))
}

func (t *ToolDataSet) GetRedisInstanceNameByID(id int) (string, error) {
	info, err := t.GetRedisInstanceInfoByID(id)
	if err != nil {
		return "", err
	}
	return info.Name, nil
}

func (t *ToolDataSet) GetPodNodeInfoByID(id int) (*podNodeInfo, error) {
	info, exists := t.podNodeIDToInfo[id]
	if exists {
		return info, nil
	}

	log.Warning(cacheNameByIDNotFound(RESOURCE_TYPE_POD_NODE_EN, id))
	var dbItem mysql.PodNode
	result := mysql.Db.Where("id = ?", id).Find(&dbItem)
	if result.RowsAffected == 1 {
		t.addPodNode(&dbItem)
		return t.podNodeIDToInfo[dbItem.ID], nil
	}
	return nil, errors.New(dbResourceByIDNotFound(RESOURCE_TYPE_POD_NODE_EN, id))
}

func (t *ToolDataSet) GetPodNodeNameByID(id int) (string, error) {
	info, err := t.GetPodNodeInfoByID(id)
	if err != nil {
		return "", err
	}
	return info.Name, nil
}

func (t *ToolDataSet) GetPodServiceInfoByID(id int) (*podServiceInfo, error) {
	info, exists := t.podServiceIDToInfo[id]
	if exists {
		return info, nil
	}

	log.Warning(cacheNameByIDNotFound(RESOURCE_TYPE_POD_SERVICE_EN, id))
	var dbItem mysql.PodService
	result := mysql.Db.Where("id = ?", id).Find(&dbItem)
	if result.RowsAffected == 1 {
		t.addPodService(&dbItem)
		return t.podServiceIDToInfo[dbItem.ID], nil
	}
	return nil, errors.New(dbResourceByIDNotFound(RESOURCE_TYPE_POD_SERVICE_EN, id))
}

func (t *ToolDataSet) GetPodServiceNameByID(id int) (string, error) {
	info, err := t.GetPodServiceInfoByID(id)
	if err != nil {
		return "", nil
	}
	return info.Name, nil
}

func (t *ToolDataSet) GetPodInfoeByID(id int) (*podInfo, error) {
	info, exists := t.podIDToInfo[id]
	if exists {
		return info, nil
	}

	log.Warning(cacheNameByIDNotFound(RESOURCE_TYPE_POD_EN, id))
	var dbItem mysql.Pod
	result := mysql.Db.Where("id = ?", id).Find(&dbItem)
	if result.RowsAffected == 1 {
		t.addPod(&dbItem)
		return t.podIDToInfo[dbItem.ID], nil
	}
	return nil, errors.New(dbResourceByIDNotFound(RESOURCE_TYPE_POD_EN, id))
}

func (t *ToolDataSet) GetPodNameByID(id int) (string, error) {
	info, err := t.GetPodInfoeByID(id)
	if err != nil {
		return "", err
	}
	return info.Name, nil
}

func (t *ToolDataSet) GetVInterfaceLcuuidByID(id int) (string, bool) {
	lcuuid, exists := t.VInterfaceIDToLcuuid[id]
	if exists {
		return lcuuid, true
	}
	log.Warning(cacheLcuuidByIDNotFound(RESOURCE_TYPE_VINTERFACE_EN, id))
	var vif mysql.VInterface
	result := mysql.Db.Where("id = ?", id).Find(&vif)
	if result.RowsAffected == 1 {
		t.addVInterface(&vif)
		return vif.Lcuuid, true
	} else {
		log.Error(dbResourceByIDNotFound(RESOURCE_TYPE_VINTERFACE_EN, id))
		return lcuuid, false
	}
}

func (t *ToolDataSet) GetVInterfaceIDByWANIPLcuuid(lcuuid string) (int, bool) {
	vifID, exists := t.WANIPLcuuidToVInterfaceID[lcuuid]
	if exists {
		return vifID, true
	}
	log.Warningf("cache %s id (%s lcuuid: %s) not found", RESOURCE_TYPE_VINTERFACE_EN, RESOURCE_TYPE_WAN_IP_EN, lcuuid)
	var wanIP mysql.WANIP
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&wanIP)
	if result.RowsAffected == 1 {
		t.addWANIP(&wanIP)
		vifID, exists = t.WANIPLcuuidToVInterfaceID[lcuuid]
		return vifID, exists
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_WAN_IP_EN, lcuuid))
		return vifID, false
	}
}

func (t *ToolDataSet) GetWANIPByLcuuid(lcuuid string) (string, bool) {
	ip, exists := t.WANIPLcuuidToIP[lcuuid]
	if exists {
		return ip, true
	}
	log.Warning(cacheIPByLcuuidNotFound(RESOURCE_TYPE_WAN_IP_EN, lcuuid))
	var wanIP mysql.WANIP
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&wanIP)
	if result.RowsAffected == 1 {
		t.addWANIP(&wanIP)
		return wanIP.IP, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_WAN_IP_EN, lcuuid))
		return ip, false
	}
}

func (t *ToolDataSet) GetVInterfaceIDByLANIPLcuuid(lcuuid string) (int, bool) {
	vifID, exists := t.LANIPLcuuidToVInterfaceID[lcuuid]
	if exists {
		return vifID, true
	}
	log.Warningf("cache %s id (%s lcuuid: %s) not found", RESOURCE_TYPE_VINTERFACE_EN, RESOURCE_TYPE_LAN_IP_EN, lcuuid)
	var lanIP mysql.LANIP
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&lanIP)
	if result.RowsAffected == 1 {
		t.addLANIP(&lanIP)
		vifID, exists = t.LANIPLcuuidToVInterfaceID[lcuuid]
		return vifID, exists
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_LAN_IP_EN, lcuuid))
		return vifID, false
	}
}

func (t *ToolDataSet) GetLANIPByLcuuid(lcuuid string) (string, bool) {
	ip, exists := t.LANIPLcuuidToIP[lcuuid]
	if exists {
		return ip, true
	}
	log.Warning(cacheIPByLcuuidNotFound(RESOURCE_TYPE_LAN_IP_EN, lcuuid))
	var lanIP mysql.LANIP
	result := mysql.Db.Where("lcuuid = ?", lcuuid).Find(&lanIP)
	if result.RowsAffected == 1 {
		t.addLANIP(&lanIP)
		return lanIP.IP, true
	} else {
		log.Error(dbResourceByLcuuidNotFound(RESOURCE_TYPE_LAN_IP_EN, lcuuid))
		return ip, false
	}
}
