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

package model

type StaticConfig struct {
	ProxyControllerPort              *uint16                            `yaml:"proxy-controller-port,omitempty"`
	LogLevel                         *string                            `yaml:"log-level,omitempty"`
	Profiler                         *bool                              `yaml:"profiler,omitempty"`
	AfpacketBlocksEnabled            *bool                              `yaml:"afpacket-blocks-enabled,omitempty"`
	AfpacketBlocks                   *int                               `yaml:"afpacket-blocks,omitempty"`
	EnableDebugStats                 *bool                              `yaml:"enable-debug-stats,omitempty"`
	AnalyzerDedupDisabled            *bool                              `yaml:"analyzer-dedup-disabled,omitempty"`
	DefaultTapType                   *uint32                            `yaml:"default-tap-type,omitempty"`
	DebugListenPort                  *uint16                            `yaml:"debug-listen-port,omitempty"`
	EnableQosBypass                  *bool                              `yaml:"enable-qos-bypass,omitempty"`
	FastPathMapSize                  *int                               `yaml:"fast-path-map-size,omitempty"`
	FirstPathLevel                   *int                               `yaml:"first-path-level,omitempty"`
	SrcInterfaces                    []string                           `yaml:"src-interfaces,omitempty"`
	CloudGatewayTraffic              *bool                              `yaml:"cloud-gateway-traffic,omitempty"`
	MirrorTrafficPcp                 *uint16                            `yaml:"mirror-traffic-pcp,omitempty"`
	PCap                             *PCapConfig                        `yaml:"pcap,omitempty"`
	Flow                             *FlowGeneratorConfig               `yaml:"flow,omitempty"`
	FlowQueueSize                    *int                               `yaml:"flow-queue-size,omitempty"`
	QuadrupleQueueSize               *int                               `yaml:"quadruple-queue-size,omitempty"`
	AnalyzerQueueSize                *int                               `yaml:"analyzer-queue-size,omitempty"`
	OvsDpdkEnable                    *bool                              `yaml:"ovs-dpdk-enable,omitempty"`
	DpdkPmdCoreId                    *uint32                            `yaml:"dpdk-pmd-core-id,omitempty"`
	DpdkRingPort                     *string                            `yaml:"dpdk-ring-port,omitempty"`
	XflowCollector                   *XflowCollectorConfig              `yaml:"xflow-collector,omitempty"`
	VxlanPort                        *uint16                            `yaml:"vxlan-port,omitempty"`
	VxlanFlags                       *uint8                             `yaml:"vxlan-flags,omitempty"`
	CollectorSenderQueueSize         *int                               `yaml:"collector-sender-queue-size,omitempty"`
	CollectorSenderQueueCount        *int                               `yaml:"collector-sender-queue-count,omitempty"`
	FlowSenderQueueSize              *int                               `yaml:"flow-sender-queue-size,omitempty"`
	FlowSenderQueueCount             *int                               `yaml:"flow-sender-queue-count,omitempty"`
	SecondFlowExtraDelaySecond       *int                               `yaml:"second-flow-extra-delay-second,omitempty"`
	PacketDelay                      *int                               `yaml:"packet-delay,omitempty"`
	Triple                           *TripleMapConfig                   `yaml:"triple,omitempty"`
	KubernetesPollerType             *string                            `yaml:"kubernetes-poller-type,omitempty"`
	DecapErspan                      *bool                              `yaml:"decap-erspan,omitempty"`
	AnalyzerIp                       *string                            `yaml:"analyzer-ip,omitempty"`
	AnalyzerPort                     *uint16                            `yaml:"analyzer-port,omitempty"`
	KubernetesNamespace              *string                            `yaml:"kubernetes-namespace,omitempty"`
	IngressFlavour                   *string                            `yaml:"ingress-flavour,omitempty"`
	GrpcBufferSize                   *int                               `yaml:"grpc-buffer-size,omitempty"`            // 单位：M
	L7LogSessionAggrTimeout          *int                               `yaml:"l7-log-session-aggr-timeout,omitempty"` // 单位: s
	TapMacScript                     *string                            `yaml:"tap-mac-script,omitempty"`
	BpfDisabled                      *bool                              `yaml:"bpf-disabled,omitempty"`
	L7ProtocolInferenceMaxFailCount  *uint64                            `yaml:"l7-protocol-inference-max-fail-count,omitempty"`
	L7ProtocolInferenceTtl           *uint64                            `yaml:"l7-protocol-inference-ttl,omitempty"`
	PacketSequenceBlockSize          *int                               `yaml:"packet-sequence-block-size,omitempty"`
	PacketSequenceQueueSize          *int                               `yaml:"packet-sequence-queue-size,omitempty"`
	PacketSequenceQueueCount         *int                               `yaml:"packet-sequence-queue-count,omitempty"`
	PacketSequenceFlag               *uint8                             `yaml:"packet-sequence-flag,omitempty"`
	EbpfDisabled                     *bool                              `yaml:"ebpf-disabled,omitempty"`
	EbpfUprobeGolangSymbolEnabled    *bool                              `yaml:"ebpf-uprobe-golang-symbol-enabled,omitempty"`
	L7ProtocolEnabled                []string                           `yaml:"l7-protocol-enabled,omitempty"`
	EbpfUprobeProcessNameRegexs      *EbpfUprobeProcessNameRegexsConfig `yaml:"ebpf-uprobe-process-name-regexs,omitempty"`
	StandaloneDataFileSize           *uint64                            `yaml:"standalone-data-file-size,omitempty"`
	StandaloneDataFileDir            *string                            `yaml:"standalone-data-file-dir,omitempty"`
	LogFile                          *string                            `yaml:"log-file,omitempty"`
	ExternalAgentHttpProxyCompressed *bool                              `yaml:"external-agent-http-proxy-compressed,omitempty"`
	FeatureFlags                     []string                           `yaml:"feature-flags,omitempty"`
	L7ProtocolPorts                  map[string]string                  `yaml:"l7-protocol-ports,omitempty"`
}

type XflowCollectorConfig struct {
	SflowPorts   []string `yaml:"sflow-ports,omitempty"`
	NetflowPorts []string `yaml:"netflow-ports,omitempty"`
}

type PCapConfig struct {
	Enabled               *bool   `yaml:"enabled,omitempty"`
	QueueSize             *int    `yaml:"queue-size,omitempty"`
	QueueCount            *int    `yaml:"queue-count,omitempty"`
	TCPIPChecksum         *bool   `yaml:"tcpip-checksum,omitempty"`
	BlockSizeKB           *int    `yaml:"block-size-kb,omitempty"`
	MaxConcurrentFiles    *int    `yaml:"max-concurrent-files,omitempty"`
	MaxFileSizeMB         *int    `yaml:"max-file-size-mb,omitempty"`
	MaxDirectorySizeGb    *int    `yaml:"max-directory-size-gb,omitempty"`
	DiskFreeSpaceMarginGb *int    `yaml:"disk-free-space-margin-gb,omitempty"`
	MaxFilePeriodSecond   *int    `yaml:"max-file-period-second,omitempty"`
	FileDirectory         *string `yaml:"file-directory,omitempty"`
	ServerPort            *int    `yaml:"server-port,omitempty"`
}

type TripleMapConfig struct {
	HashSlots *int `yaml:"hash-slots-size,omitempty"`
	Capacity  *int `yaml:"capacity,omitempty"`
}

type TcpTimeoutConfig struct {
	EstablishedTimeout *int `yaml:"established-timeout,omitempty"`
	ClosingRstTimeout  *int `yaml:"closing-rst-timeout,omitempty"`
	OthersTimeout      *int `yaml:"others-timeout,omitempty"`
}

type FlowGeneratorConfig struct {
	TcpTimeoutConfig `yaml:",inline"`
	HashSlots        *int `yaml:"flow-slots-size,omitempty"`
	Capacity         *int `yaml:"flow-count-limit,omitempty"`
	FlushInterval    *int `yaml:"flush-interval,omitempty"`
	SenderThrottle   *int `yaml:"flow-sender-throttle,omitempty"`
	AggrQueueSize    *int `yaml:"flow-aggr-queue-size,omitempty"`

	IgnoreTorMac *bool `yaml:"ignore-tor-mac,omitempty"`
	IgnoreL2End  *bool `yaml:"ignore-l2-end,omitempty"`
}

type EbpfUprobeProcessNameRegexsConfig struct {
	GolangSymbol *string `yaml:"golang-symbol,omitempty"`
	Golang       *string `yaml:"golang,omitempty"`
	Openssl      *string `yaml:"openssl,omitempty"`
}
