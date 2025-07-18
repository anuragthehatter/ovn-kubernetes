// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package nbdb

import "github.com/ovn-kubernetes/libovsdb/model"

const LogicalRouterPortTable = "Logical_Router_Port"

// LogicalRouterPort defines an object in Logical_Router_Port table
type LogicalRouterPort struct {
	UUID           string            `ovsdb:"_uuid"`
	DhcpRelay      *string           `ovsdb:"dhcp_relay"`
	Enabled        *bool             `ovsdb:"enabled"`
	ExternalIDs    map[string]string `ovsdb:"external_ids"`
	GatewayChassis []string          `ovsdb:"gateway_chassis"`
	HaChassisGroup *string           `ovsdb:"ha_chassis_group"`
	Ipv6Prefix     []string          `ovsdb:"ipv6_prefix"`
	Ipv6RaConfigs  map[string]string `ovsdb:"ipv6_ra_configs"`
	MAC            string            `ovsdb:"mac"`
	Name           string            `ovsdb:"name"`
	Networks       []string          `ovsdb:"networks"`
	Options        map[string]string `ovsdb:"options"`
	Peer           *string           `ovsdb:"peer"`
	Status         map[string]string `ovsdb:"status"`
}

func (a *LogicalRouterPort) GetUUID() string {
	return a.UUID
}

func (a *LogicalRouterPort) GetDhcpRelay() *string {
	return a.DhcpRelay
}

func copyLogicalRouterPortDhcpRelay(a *string) *string {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalLogicalRouterPortDhcpRelay(a, b *string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *LogicalRouterPort) GetEnabled() *bool {
	return a.Enabled
}

func copyLogicalRouterPortEnabled(a *bool) *bool {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalLogicalRouterPortEnabled(a, b *bool) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *LogicalRouterPort) GetExternalIDs() map[string]string {
	return a.ExternalIDs
}

func copyLogicalRouterPortExternalIDs(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalLogicalRouterPortExternalIDs(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func (a *LogicalRouterPort) GetGatewayChassis() []string {
	return a.GatewayChassis
}

func copyLogicalRouterPortGatewayChassis(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalLogicalRouterPortGatewayChassis(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func (a *LogicalRouterPort) GetHaChassisGroup() *string {
	return a.HaChassisGroup
}

func copyLogicalRouterPortHaChassisGroup(a *string) *string {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalLogicalRouterPortHaChassisGroup(a, b *string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *LogicalRouterPort) GetIpv6Prefix() []string {
	return a.Ipv6Prefix
}

func copyLogicalRouterPortIpv6Prefix(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalLogicalRouterPortIpv6Prefix(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func (a *LogicalRouterPort) GetIpv6RaConfigs() map[string]string {
	return a.Ipv6RaConfigs
}

func copyLogicalRouterPortIpv6RaConfigs(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalLogicalRouterPortIpv6RaConfigs(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func (a *LogicalRouterPort) GetMAC() string {
	return a.MAC
}

func (a *LogicalRouterPort) GetName() string {
	return a.Name
}

func (a *LogicalRouterPort) GetNetworks() []string {
	return a.Networks
}

func copyLogicalRouterPortNetworks(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalLogicalRouterPortNetworks(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func (a *LogicalRouterPort) GetOptions() map[string]string {
	return a.Options
}

func copyLogicalRouterPortOptions(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalLogicalRouterPortOptions(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func (a *LogicalRouterPort) GetPeer() *string {
	return a.Peer
}

func copyLogicalRouterPortPeer(a *string) *string {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalLogicalRouterPortPeer(a, b *string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *LogicalRouterPort) GetStatus() map[string]string {
	return a.Status
}

func copyLogicalRouterPortStatus(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalLogicalRouterPortStatus(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func (a *LogicalRouterPort) DeepCopyInto(b *LogicalRouterPort) {
	*b = *a
	b.DhcpRelay = copyLogicalRouterPortDhcpRelay(a.DhcpRelay)
	b.Enabled = copyLogicalRouterPortEnabled(a.Enabled)
	b.ExternalIDs = copyLogicalRouterPortExternalIDs(a.ExternalIDs)
	b.GatewayChassis = copyLogicalRouterPortGatewayChassis(a.GatewayChassis)
	b.HaChassisGroup = copyLogicalRouterPortHaChassisGroup(a.HaChassisGroup)
	b.Ipv6Prefix = copyLogicalRouterPortIpv6Prefix(a.Ipv6Prefix)
	b.Ipv6RaConfigs = copyLogicalRouterPortIpv6RaConfigs(a.Ipv6RaConfigs)
	b.Networks = copyLogicalRouterPortNetworks(a.Networks)
	b.Options = copyLogicalRouterPortOptions(a.Options)
	b.Peer = copyLogicalRouterPortPeer(a.Peer)
	b.Status = copyLogicalRouterPortStatus(a.Status)
}

func (a *LogicalRouterPort) DeepCopy() *LogicalRouterPort {
	b := new(LogicalRouterPort)
	a.DeepCopyInto(b)
	return b
}

func (a *LogicalRouterPort) CloneModelInto(b model.Model) {
	c := b.(*LogicalRouterPort)
	a.DeepCopyInto(c)
}

func (a *LogicalRouterPort) CloneModel() model.Model {
	return a.DeepCopy()
}

func (a *LogicalRouterPort) Equals(b *LogicalRouterPort) bool {
	return a.UUID == b.UUID &&
		equalLogicalRouterPortDhcpRelay(a.DhcpRelay, b.DhcpRelay) &&
		equalLogicalRouterPortEnabled(a.Enabled, b.Enabled) &&
		equalLogicalRouterPortExternalIDs(a.ExternalIDs, b.ExternalIDs) &&
		equalLogicalRouterPortGatewayChassis(a.GatewayChassis, b.GatewayChassis) &&
		equalLogicalRouterPortHaChassisGroup(a.HaChassisGroup, b.HaChassisGroup) &&
		equalLogicalRouterPortIpv6Prefix(a.Ipv6Prefix, b.Ipv6Prefix) &&
		equalLogicalRouterPortIpv6RaConfigs(a.Ipv6RaConfigs, b.Ipv6RaConfigs) &&
		a.MAC == b.MAC &&
		a.Name == b.Name &&
		equalLogicalRouterPortNetworks(a.Networks, b.Networks) &&
		equalLogicalRouterPortOptions(a.Options, b.Options) &&
		equalLogicalRouterPortPeer(a.Peer, b.Peer) &&
		equalLogicalRouterPortStatus(a.Status, b.Status)
}

func (a *LogicalRouterPort) EqualsModel(b model.Model) bool {
	c := b.(*LogicalRouterPort)
	return a.Equals(c)
}

var _ model.CloneableModel = &LogicalRouterPort{}
var _ model.ComparableModel = &LogicalRouterPort{}
