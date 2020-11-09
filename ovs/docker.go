package ovs

import "fmt"

// Note that following functions should be executed
// through the version of ovs-docker located under scripts/ovs-docker

// A DockerService is used in a Client to execute 'ovs-docker' commands.
type DockerService struct {
	// Wrapped Client for ExecFunc and debugging.
	c *Client
}

// AddPort attaches a port to a bridge on Open vSwitch.
func (v *DockerService) AddPort(bridge string, iface string, cid string, options DockerOptions) error {
	args := []string{"add-port", bridge, iface, cid}
	args = append(args, options.slice()...)

	_, err := v.exec(args...)
	return err
}

// DeletePort removes the interface from the container and detaches the port form Open vSwitch
func (v *DockerService) DeletePort(bridge string, iface string, cid string) error {
	_, err := v.exec("del-port", bridge, iface, cid)
	return err
}

// DeletePorts removes all Open vSwitch interfaces from the container.
func (v *DockerService) DeletePorts(bridge string, cid string) error {
	_, err := v.exec("del-ports", bridge, cid)
	return err
}

// SetVlan configures the interface of the container .
func (v *DockerService) SetVlan(bridge string, iface string, cid string, vlan string) error {
	_, err := v.exec("set-vlan", bridge, iface, cid, vlan)
	return err
}

// exec executes an ExecFunc using 'ovs-docker'.
func (v *DockerService) exec(args ...string) ([]byte, error) {
	return v.c.exec("ovs-docker", args...)
}

// A DockerOptions struct enables configuration of the interface added to the docker container
type DockerOptions struct {
	// IPAddress defines the ipaddress that the interface should be given in the CIDR format, example 192.168.1.6/24
	IPAddress string
	// MACAddress defines the MACAddress of the interface added to the docker container
	MACAddress string
	// Gateway is used to add a route to the docker container
	Gateway string
	MTU     string
	// DHCP true or not
	DHCP bool
	// VLAN
	VlanTag string
}

func (d DockerOptions) slice() []string {
	var s []string

	if d.IPAddress != "" {
		s = append(s, fmt.Sprintf("--ipaddress=%s", d.IPAddress))
	}
	if d.MACAddress != "" {
		s = append(s, fmt.Sprintf("--macaddress=%s", d.MACAddress))
	}
	if d.Gateway != "" {
		s = append(s, fmt.Sprintf("--gateway=%s", d.Gateway))
	}
	if d.MTU != "" {
		s = append(s, fmt.Sprintf("--mtu=%s", d.MTU))
	}
	if d.DHCP {
		s = append(s, fmt.Sprintf("--dhcp=%v", d.DHCP))
	}

	if d.VlanTag != "" {
		s = append(s, fmt.Sprintf("--vlan=%s", d.VlanTag))
	}
	return s
}
