// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linuxclient

import (
	vpp_clientv2 "go.ligato.io/vpp-agent/v3/clientv2/vpp"
	linux_interfaces "go.ligato.io/vpp-agent/v3/proto/ligato/linux/interfaces"
	linux_iptables "go.ligato.io/vpp-agent/v3/proto/ligato/linux/iptables"
	linux_l3 "go.ligato.io/vpp-agent/v3/proto/ligato/linux/l3"
	vpp_abf "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/abf"
	vpp_acl "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/acl"
	vpp_interfaces "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/interfaces"
	ipfix "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/ipfix"
	ipsec "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/ipsec"
	vpp_l2 "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/l2"
	vpp_l3 "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/l3"
	nat "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/nat"
	punt "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/punt"
	vpp_stn "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/stn"
)

// DataResyncDSL defines the Domain Specific Language (DSL) for data RESYNC
// of both Linux and VPP configuration.
// Use this interface to make your implementation independent of the local
// and any remote client.
// Each method (apart from Send) returns the receiver, allowing the calls
// to be chained together conveniently in a single statement.
type DataResyncDSL interface {
	// LinuxInterface adds Linux interface to the RESYNC request.
	LinuxInterface(intf *linux_interfaces.Interface) DataResyncDSL
	// LinuxInterface adds Linux ARP entry to the RESYNC request.
	LinuxArpEntry(arp *linux_l3.ARPEntry) DataResyncDSL
	// LinuxInterface adds Linux route to the RESYNC request.
	LinuxRoute(route *linux_l3.Route) DataResyncDSL
	// IptablesRuleChain adds iptables rule chain to the RESYNC request.
	IptablesRuleChain(val *linux_iptables.RuleChain) DataResyncDSL

	// VppInterface adds VPP interface to the RESYNC request.
	VppInterface(intf *vpp_interfaces.Interface) DataResyncDSL
	// Span adds VPP span to the RESYNC request.
	Span(span *vpp_interfaces.Span) DataResyncDSL
	// ACL adds VPP Access Control List to the RESYNC request.
	ACL(acl *vpp_acl.ACL) DataResyncDSL
	// ABF adds ACL-based forwarding to the RESYNC request.
	ABF(acl *vpp_abf.ABF) DataResyncDSL
	/*// BfdSession adds VPP bidirectional forwarding detection session
	// to the RESYNC request.
	BfdSession(val *vpp_bfd.SingleHopBFD_Session) DataResyncDSL
	// BfdAuthKeys adds VPP bidirectional forwarding detection key to the RESYNC
	// request.
	BfdAuthKeys(val *vpp_bfd.SingleHopBFD_Key) DataResyncDSL
	// BfdEchoFunction adds VPP bidirectional forwarding detection echo function
	// to the RESYNC request.
	BfdEchoFunction(val *vpp_bfd.SingleHopBFD_EchoFunction) DataResyncDSL*/
	// BD adds VPP Bridge Domain to the RESYNC request.
	BD(bd *vpp_l2.BridgeDomain) DataResyncDSL
	// BDFIB adds VPP L2 FIB to the RESYNC request.
	BDFIB(fib *vpp_l2.FIBEntry) DataResyncDSL
	// VrfTable adds VPP VRF table to the RESYNC request.
	VrfTable(val *vpp_l3.VrfTable) DataResyncDSL
	// XConnect adds VPP Cross Connect to the RESYNC request.
	XConnect(xcon *vpp_l2.XConnectPair) DataResyncDSL
	// StaticRoute adds VPP L3 Static Route to the RESYNC request.
	StaticRoute(staticRoute *vpp_l3.Route) DataResyncDSL
	// Arp adds VPP L3 ARP to the RESYNC request.
	Arp(arp *vpp_l3.ARPEntry) DataResyncDSL
	// ProxyArp adds L3 proxy ARP interfaces to the RESYNC request.
	ProxyArp(proxyArp *vpp_l3.ProxyARP) DataResyncDSL
	// IPScanNeighbor adds L3 IP Scan Neighbor to the RESYNC request.
	IPScanNeighbor(ipScanNeigh *vpp_l3.IPScanNeighbor) DataResyncDSL
	/*// L4Features adds L4 features to the RESYNC request
	L4Features(val *vpp_l4.L4Features) DataResyncDSL
	// AppNamespace adds VPP Application namespaces to the RESYNC request
	AppNamespace(appNs *vpp_l4.AppNamespaces_AppNamespace) DataResyncDSL*/
	// StnRule adds Stn rule to the RESYNC request.
	StnRule(stn *vpp_stn.Rule) DataResyncDSL
	// NAT44Global adds global NAT44 configuration to the RESYNC request.
	NAT44Global(nat *nat.Nat44Global) DataResyncDSL
	// DNAT44 adds DNAT44 configuration to the RESYNC request
	DNAT44(dnat *nat.DNat44) DataResyncDSL
	// NAT44Interface adds NAT44 interface configuration to the RESYNC request.
	NAT44Interface(natIf *nat.Nat44Interface) DataResyncDSL
	// NAT44AddressPool adds NAT44 address pool configuration to the RESYNC request.
	NAT44AddressPool(pool *nat.Nat44AddressPool) DataResyncDSL
	// IPSecSA adds request to RESYNC a new Security Association
	IPSecSA(sa *ipsec.SecurityAssociation) DataResyncDSL
	// IPSecSPD adds request to RESYNC a new Security Policy Database
	IPSecSPD(spd *ipsec.SecurityPolicyDatabase) DataResyncDSL
	// IPSecSP adds Security Policy to the RESYNC request
	IPSecSP(sp *ipsec.SecurityPolicy) DataResyncDSL
	// IPSecTunnelProtection adds request to RESYNC an IPSec tunnel protection
	IPSecTunnelProtection(tp *ipsec.TunnelProtection) DataResyncDSL
	// PuntIPRedirect adds request to RESYNC a rule used to punt L3 traffic via interface.
	PuntIPRedirect(val *punt.IPRedirect) DataResyncDSL
	// PuntToHost adds request to RESYNC a rule used to punt L4 traffic to a host.
	PuntToHost(val *punt.ToHost) DataResyncDSL
	// PuntException adds request to create or update exception to punt specific packets.
	PuntException(val *punt.Exception) DataResyncDSL
	// IPFIX adds IPFIX configuration to the RESYNC request.
	IPFIX(val *ipfix.IPFIX) DataResyncDSL
	// FlowprobeParams adds Flowprobe Params configuration to the RESYNC request.
	FlowprobeParams(val *ipfix.FlowProbeParams) DataResyncDSL
	// FlowprobeFeature adds Flowprobe Feature configuration to the RESYNC request.
	FlowprobeFeature(val *ipfix.FlowProbeFeature) DataResyncDSL

	// Send propagates the RESYNC request to the plugins.
	Send() vpp_clientv2.Reply
}
