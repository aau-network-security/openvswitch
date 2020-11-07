package ovs

import (
	"reflect"
	"testing"
)

func TestClientDockerAddPortOK(t *testing.T) {
	bridge := "br0"
	iface := "eth0"
	container := "ubuntu"
	options := DockerOptions{
		IPAddress:  "192.168.1.2",
		Gateway:    "192.168.1.1",
		MACAddress: "c1:9c:64:9a:bb:c3",
		MTU:        "1500",
	}

	// Apply Timeout option to verify arguments
	c := testClient([]OptionFunc{Timeout(1)}, func(cmd string, args ...string) ([]byte, error) {
		// Verify correct command and arguments passed, including option flags
		if want, got := "ovs-docker", cmd; want != got {
			t.Fatalf("incorrect command:\n- want: %v\n-  got: %v",
				want, got)
		}

		wantArgs := []string{"--timeout=1", "add-port", string(bridge), string(iface), string(container), "--ipaddress=192.168.1.2", "--macaddress=c1:9c:64:9a:bb:c3", "--gateway=192.168.1.1", "--mtu=1500"}
		if want, got := wantArgs, args; !reflect.DeepEqual(want, got) {
			t.Fatalf("incorrect arguments\n- want: %v\n-  got: %v",
				want, got)
		}

		return nil, nil
	})

	if err := c.Docker.AddPort(bridge, iface, container, options); err != nil {
		t.Fatalf("unexpected error for Client.VSwitch.AddPort: %v", err)
	}
}

func TestClientDockerSetVlanOK(t *testing.T) {
	bridge := "br0"
	iface := "eth0"
	container := "ubuntu"
	vlan := "5"

	// Apply Timeout option to verify arguments
	c := testClient([]OptionFunc{Timeout(1)}, func(cmd string, args ...string) ([]byte, error) {
		// Verify correct command and arguments passed, including option flags
		if want, got := "ovs-docker", cmd; want != got {
			t.Fatalf("incorrect command:\n- want: %v\n-  got: %v",
				want, got)
		}

		wantArgs := []string{"--timeout=1", "set-vlan", string(bridge), string(iface), string(container), string(vlan)}
		if want, got := wantArgs, args; !reflect.DeepEqual(want, got) {
			t.Fatalf("incorrect arguments\n- want: %v\n-  got: %v",
				want, got)
		}

		return nil, nil
	})

	if err := c.Docker.SetVlan(bridge, iface, container, vlan); err != nil {
		t.Fatalf("unexpected error for Client.VSwitch.AddPort: %v", err)
	}
}

func TestClientDockerDeletePortOK(t *testing.T) {
	bridge := "br0"
	iface := "eth1"
	container := "ubuntu"

	// Apply Timeout option to verify arguments
	c := testClient([]OptionFunc{Timeout(1)}, func(cmd string, args ...string) ([]byte, error) {
		// Verify correct command and arguments passed, including option flags
		if want, got := "ovs-docker", cmd; want != got {
			t.Fatalf("incorrect command:\n- want: %v\n-  got: %v",
				want, got)
		}

		wantArgs := []string{"--timeout=1", "del-port", string(bridge), string(iface), string(container)}
		if want, got := wantArgs, args; !reflect.DeepEqual(want, got) {
			t.Fatalf("incorrect arguments\n- want: %v\n-  got: %v",
				want, got)
		}

		return nil, nil
	})
	if err := c.Docker.DeletePort(bridge, iface, container); err != nil {
		t.Fatalf("unexpected error for Client.Docker.DeletePort: %v", err)
	}
}

func TestClientDockerDeletePortsOK(t *testing.T) {
	bridge := "br0"
	container := "ubuntu"

	// Apply Timeout option to verify arguments
	c := testClient([]OptionFunc{Timeout(1)}, func(cmd string, args ...string) ([]byte, error) {
		// Verify correct command and arguments passed, including option flags
		if want, got := "ovs-docker", cmd; want != got {
			t.Fatalf("incorrect command:\n- want: %v\n-  got: %v",
				want, got)
		}

		wantArgs := []string{"--timeout=1", "del-ports", string(bridge), string(container)}
		if want, got := wantArgs, args; !reflect.DeepEqual(want, got) {
			t.Fatalf("incorrect arguments\n- want: %v\n-  got: %v",
				want, got)
		}

		return nil, nil
	})
	if err := c.Docker.DeletePorts(bridge, container); err != nil {
		t.Fatalf("unexpected error for Client.Docker.DeletePort: %v", err)
	}
}

func TestDockerOptions_slice(t *testing.T) {
	var tests = []struct {
		desc string
		i    DockerOptions
		out  []string
	}{
		{
			desc: "no options",
		},
		{
			desc: "only ipaddress",
			i: DockerOptions{
				IPAddress: "192.168.1.2/24",
			},
			out: []string{
				"--ipaddress=192.168.1.2/24",
			},
		},
		{
			desc: "only MACAddress",
			i: DockerOptions{
				MACAddress: "c1:9c:64:9a:bb:c3",
			},
			out: []string{
				"--macaddress=c1:9c:64:9a:bb:c3",
			},
		},
		{
			desc: "only gateway",
			i: DockerOptions{
				Gateway: "192.168.1.1",
			},
			out: []string{
				"--gateway=192.168.1.1",
			},
		},
		{
			desc: "only MTU",
			i: DockerOptions{
				MTU: "1500",
			},
			out: []string{
				"--mtu=1500",
			},
		},
		{
			desc: "IP and mac",
			i: DockerOptions{
				IPAddress:  "192.168.1.2",
				MACAddress: "c1:9c:64:9a:bb:c3",
			},
			out: []string{
				"--ipaddress=192.168.1.2",
				"--macaddress=c1:9c:64:9a:bb:c3",
			},
		},
		{
			desc: "IP and gateway",
			i: DockerOptions{
				IPAddress: "192.168.1.2",
				Gateway:   "192.168.1.1",
			},
			out: []string{
				"--ipaddress=192.168.1.2",
				"--gateway=192.168.1.1",
			},
		},
		{
			desc: "IP and MTU",
			i: DockerOptions{
				IPAddress: "192.168.1.2",
				MTU:       "1500",
			},
			out: []string{
				"--ipaddress=192.168.1.2",
				"--mtu=1500",
			},
		},
		{
			desc: "gateway and mtu",
			i: DockerOptions{
				Gateway: "192.168.1.1",
				MTU:     "1500",
			},
			out: []string{
				"--gateway=192.168.1.1",
				"--mtu=1500",
			},
		},
		{
			desc: "all options",
			i: DockerOptions{
				IPAddress:  "192.168.1.2",
				Gateway:    "192.168.1.1",
				MACAddress: "c1:9c:64:9a:bb:c3",
				MTU:        "1500",
			},
			out: []string{
				"--ipaddress=192.168.1.2",
				"--macaddress=c1:9c:64:9a:bb:c3",
				"--gateway=192.168.1.1",
				"--mtu=1500",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if want, got := tt.out, tt.i.slice(); !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected slices:\n- want: %v\n-  got: %v",
					want, got)
			}
		})
	}
}
