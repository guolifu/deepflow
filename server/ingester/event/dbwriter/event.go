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

package dbwriter

import (
	"github.com/deepflowys/deepflow/server/ingester/event/common"
	"github.com/deepflowys/deepflow/server/libs/ckdb"
	"github.com/deepflowys/deepflow/server/libs/pool"
	"github.com/deepflowys/deepflow/server/libs/zerodoc"
)

const (
	DefaultPartition = ckdb.TimeFuncTwelveHour
)

type EventStore struct {
	Time   uint32 // s
	Source string
	Tagged uint8

	InstanceType uint32 // l3_device_type
	InstanceID   uint32
	InstanceName string

	EventType        string
	EventDescription string
	SubnetIDs        []uint32
	IPs              []string

	RegionID     uint16
	AZID         uint16
	L3EpcID      int32
	HostID       uint16
	PodID        uint32
	PodNodeID    uint32
	PodNSID      uint16
	PodClusterID uint16
	PodGroupID   uint32
	L3DeviceType uint8
	L3DeviceID   uint32
	ServiceID    uint32
}

func (e *EventStore) WriteBlock(block *ckdb.Block) error {
	if err := block.WriteDateTime(e.Time); err != nil {
		return err
	}
	if err := block.WriteString(e.Source); err != nil {
		return err
	}
	if err := block.WriteUInt8(e.Tagged); err != nil {
		return err
	}
	if err := block.WriteUInt32(e.InstanceType); err != nil {
		return err
	}
	if err := block.WriteUInt32(e.InstanceID); err != nil {
		return err
	}
	if err := block.WriteString(e.InstanceName); err != nil {
		return err
	}
	if err := block.WriteString(e.EventType); err != nil {
		return err
	}
	if err := block.WriteString(e.EventDescription); err != nil {
		return err
	}
	if err := block.WriteArrayUInt32(e.SubnetIDs); err != nil {
		return err
	}
	if err := block.WriteArrayString(e.IPs); err != nil {
		return err
	}

	if err := block.WriteUInt16(e.RegionID); err != nil {
		return err
	}
	if err := block.WriteUInt16(e.AZID); err != nil {
		return err
	}
	if err := block.WriteInt32(zerodoc.MarshalInt32WithSpecialID(int16(e.L3EpcID))); err != nil {
		return err
	}
	if err := block.WriteUInt16(e.HostID); err != nil {
		return err
	}

	if err := block.WriteUInt32(e.PodID); err != nil {
		return err
	}
	if err := block.WriteUInt32(e.PodNodeID); err != nil {
		return err
	}
	if err := block.WriteUInt16(e.PodNSID); err != nil {
		return err
	}
	if err := block.WriteUInt16(e.PodClusterID); err != nil {
		return err
	}
	if err := block.WriteUInt32(e.PodGroupID); err != nil {
		return err
	}

	if err := block.WriteUInt8(e.L3DeviceType); err != nil {
		return err
	}
	if err := block.WriteUInt32(e.L3DeviceID); err != nil {
		return err
	}
	if err := block.WriteUInt32(e.ServiceID); err != nil {
		return err
	}

	return nil
}

func (e *EventStore) Release() {
	ReleaseEventStore(e)
}

func EventColumns() []*ckdb.Column {
	return []*ckdb.Column{
		ckdb.NewColumn("time", ckdb.DateTime),
		ckdb.NewColumn("source", ckdb.LowCardinalityString).SetComment("事件来源"),
		ckdb.NewColumn("tagged", ckdb.UInt8).SetComment("标签是否为填充"),
		ckdb.NewColumn("instance_type", ckdb.UInt32).SetComment("资源类型"),
		ckdb.NewColumn("instance_id", ckdb.UInt32).SetComment("资源ID"),
		ckdb.NewColumn("instance_name", ckdb.LowCardinalityString).SetComment("资源名称"),
		ckdb.NewColumn("event_type", ckdb.LowCardinalityString).SetComment("事件类型"),
		ckdb.NewColumn("event_desc", ckdb.String).SetComment("事件信息"),
		ckdb.NewColumn("subnet_ids", ckdb.ArrayUInt32).SetComment("子网IDs"),
		ckdb.NewColumn("ips", ckdb.ArrayString).SetComment("IPs"),

		ckdb.NewColumn("region_id", ckdb.UInt16).SetComment("云平台区域ID"),
		ckdb.NewColumn("az_id", ckdb.UInt16).SetComment("可用区ID"),
		ckdb.NewColumn("l3_epc_id", ckdb.Int32).SetComment("ip对应的EPC ID"),
		ckdb.NewColumn("host_id", ckdb.UInt16).SetComment("宿主机ID"),
		ckdb.NewColumn("pod_id", ckdb.UInt32).SetComment("容器ID"),
		ckdb.NewColumn("pod_node_id", ckdb.UInt32).SetComment("容器节点ID"),
		ckdb.NewColumn("pod_ns_id", ckdb.UInt16).SetComment("容器命名空间ID"),
		ckdb.NewColumn("pod_cluster_id", ckdb.UInt16).SetComment("容器集群ID"),
		ckdb.NewColumn("pod_group_id", ckdb.UInt32).SetComment("容器组ID"),

		ckdb.NewColumn("l3_device_type", ckdb.UInt8).SetComment("资源类型"),
		ckdb.NewColumn("l3_device_id", ckdb.UInt32).SetComment("资源ID"),
		ckdb.NewColumn("service_id", ckdb.UInt32).SetComment("服务ID"),
	}
}

func GenEventCKTable(eventType common.EventType, cluster, storagePolicy string, ttl int, coldStorage *ckdb.ColdStorage) *ckdb.Table {
	timeKey := "time"
	engine := ckdb.MergeTree
	orderKeys := []string{"l3_device_type", "l3_device_id", timeKey}

	var columns []*ckdb.Column
	switch eventType {
	case common.RESOURCE_EVENT:
		columns = EventColumns()
	default:
		return nil
	}

	return &ckdb.Table{
		Database:        EVENT_DB,
		LocalName:       eventType.TableName() + ckdb.LOCAL_SUBFFIX,
		GlobalName:      eventType.TableName(),
		Columns:         columns,
		TimeKey:         timeKey,
		TTL:             ttl,
		PartitionFunc:   DefaultPartition,
		Engine:          engine,
		Cluster:         cluster,
		StoragePolicy:   storagePolicy,
		ColdStorage:     *coldStorage,
		OrderKeys:       orderKeys,
		PrimaryKeyCount: len(orderKeys),
	}
}

var eventPool = pool.NewLockFreePool(func() interface{} {
	return &EventStore{
		SubnetIDs: []uint32{},
		IPs:       []string{},
	}
})

func AcquireEventStore() *EventStore {
	return eventPool.Get().(*EventStore)
}

func ReleaseEventStore(e *EventStore) {
	if e == nil {
		return
	}
	subnetIDs := e.SubnetIDs[:0]
	ips := e.IPs[:0]
	*e = EventStore{}
	e.SubnetIDs = subnetIDs
	e.IPs = ips
	eventPool.Put(e)
}
