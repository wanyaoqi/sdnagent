// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"yunion.io/x/jsonutils"
)

type guestDesc struct {
	NICs               []*GuestNIC `json:"nics"`
	SecurityRules      string      `json:"security_rules"`
	AdminSecurityRules string      `json:"admin_security_rules"`
	Name               string

	IsMaster bool   `json:"is_master"`
	HostId   string `json:"host_id"`

	SrcIpCheck  bool `json:"src_ip_check"`
	SrcMacCheck bool `json:"src_mac_check"`
}

func newGuestDesc() *guestDesc {
	desc := &guestDesc{
		IsMaster:    true,
		SrcIpCheck:  true,
		SrcMacCheck: true,
	}
	return desc
}

type GuestNIC struct {
	Bridge     string
	Bw         int
	Dns        string
	Domain     string
	Driver     string
	Gateway    string
	IfnameHost string `json:"ifname"`
	Index      int
	IfnameVM   string   `json:"interface"`
	IP         string   `json:"ip"`
	VirtualIps []string `json:"virtual_ips"`
	MAC        string
	Masklen    int
	Net        string
	NetId      string `json:"net_id"`
	Virtual    bool
	VLAN       int
	WireId     string      `json:"wire_id"`
	HostId     string      `json:"host_id"`
	Vpc        GuestNICVpc `json:"vpc"`

	CtZoneId    uint16 `json:"-"`
	CtZoneIdSet bool   `json:"-"`
	PortNo      int    `json:"-"`

	NetworkAddresses []GuestNICNetworkAddress `json:"networkaddresses"`
}

type GuestNICNetworkAddress struct {
	Type    string `json:"type"`
	IpAddr  string `json:"ip_addr"`
	Masklen int    `json:"masklen"`
	Gateway string `json:"gateway"`
}

type GuestNICVpc struct {
	Id           string
	Provider     string `json:"provider"`
	MappedIpAddr string `json:"mapped_ip_addr"`
}

func (n *GuestNIC) TcData() *TcData {
	return &TcData{
		Type:        TC_DATA_TYPE_GUEST,
		Ifname:      n.IfnameHost,
		IngressMbps: uint64(n.Bw),
	}
}

func (n *GuestNIC) Map() map[string]interface{} {
	m := map[string]interface{}{
		"IP":      n.IP,
		"SubIPs":  n.SubIPs(),
		"MAC":     n.MAC,
		"VLAN":    n.VLAN & 0xfff,
		"CT_ZONE": n.CtZoneId,
		"PortNo":  n.PortNo,
	}
	vlanTci := n.VLAN & 0xfff
	if n.VLAN > 1 {
		// 802.1Q vlan header present
		vlanTci |= 0x1000
	} else {
		vlanTci = 0
	}
	m["VLANTci"] = fmt.Sprintf("0x%04x/0x1fff", vlanTci)
	return m
}

func (n *GuestNIC) SubIPs() []string {
	var (
		ipAddrs []string
		nas     = n.NetworkAddresses
	)
	for i := range nas {
		na := &nas[i]
		if na.Type == "sub_ip" {
			ipAddrs = append(ipAddrs, na.IpAddr)
		}
	}
	for _, ip := range n.VirtualIps {
		ipAddrs = append(ipAddrs, ip)
	}
	return ipAddrs
}

type Guest struct {
	Id         string
	Path       string
	HostConfig *HostConfig

	Name          string
	SecurityRules *SecurityRules
	NICs          []*GuestNIC
	VpcNICs       []*GuestNIC
	HostId        string

	srcIpCheck  bool
	srcMacCheck bool

	isSlave bool
}

func (g *Guest) IsVM() bool {
	startvmPath := path.Join(g.Path, "startvm")
	_, err := os.Stat(startvmPath)
	if err == nil {
		return true
	}
	return false
}

func (g *Guest) Running() bool {
	pidPath := path.Join(g.Path, "pid")
	pidData, err := ioutil.ReadFile(pidPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		// NOTE just to be reservative
		return true
	}
	if len(pidData) == 0 {
		return false
	}
	// NOTE check /proc/<pid>/exe links a qemu executable
	pidStr := strings.TrimSpace(string(pidData))
	procPath := path.Join("/proc", pidStr)
	fileInfo, err := os.Stat(procPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		// NOTE just to be reservative
		return true
	}
	if fileInfo.IsDir() {
		return true
	}
	return false
}

func (g *Guest) IsSlave() bool {
	return g.isSlave
}

func (g *Guest) GetJSONObjectDesc() (jsonutils.JSONObject, error) {
	descPath := path.Join(g.Path, "desc")
	data, err := ioutil.ReadFile(descPath)
	if err != nil {
		return nil, err
	}
	obj, err := jsonutils.Parse(data)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (g *Guest) LoadDesc() error {
	descPath := path.Join(g.Path, "desc")
	descFile, err := os.Open(descPath)
	if err != nil {
		return err
	}
	defer descFile.Close()
	dec := json.NewDecoder(descFile)
	desc := newGuestDesc()
	err = dec.Decode(&desc)
	if err != nil {
		return err
	}
	g.Name = desc.Name
	g.HostId = desc.HostId
	g.NICs = desc.NICs

	g.VpcNICs = nil
	for i := len(g.NICs) - 1; i >= 0; i-- {
		nic := g.NICs[i]
		if nic.Vpc.Provider != "" {
			g.VpcNICs = append(g.VpcNICs, nic)
			g.NICs = append(g.NICs[:i], g.NICs[i+1:]...)
		}
	}
	g.isSlave = !desc.IsMaster

	g.srcIpCheck = desc.SrcIpCheck
	g.srcMacCheck = desc.SrcMacCheck
	if !g.srcMacCheck && g.srcIpCheck {
		g.srcIpCheck = false
	}

	{
		rstr := desc.AdminSecurityRules + "; " + desc.SecurityRules
		rs, err := NewSecurityRules(rstr)
		if err != nil {
			return err
		}
		g.SecurityRules = rs
	}
	return nil
}

func (g *Guest) NeedsSync() bool {
	if g.HostId == "" {
		if len(g.VpcNICs) > 0 {
			return true
		}
	}
	return false
}

func (g *Guest) SrcIpCheck() bool {
	if !g.HostConfig.AllowRouterVMs {
		return true
	}
	return g.srcIpCheck
}

func (g *Guest) SrcMacCheck() bool {
	if !g.HostConfig.AllowSwitchVMs {
		return true
	}
	return g.srcMacCheck
}

func (g *Guest) FindNicByNetIdIP(netId, ip string) *GuestNIC {
	var searchNic = func(nics []*GuestNIC) *GuestNIC {
		for _, nic := range nics {
			if nic.NetId == netId && nic.IP == ip {
				return nic
			}
		}
		return nil
	}
	if nic := searchNic(g.VpcNICs); nic != nil {
		return nic
	}
	if nic := searchNic(g.NICs); nic != nil {
		return nic
	}
	return nil
}
