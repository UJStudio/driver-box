package network

import (
	"github.com/ibuilding-x/driver-box/internal/plugins/bacnet/bacnet"
	"github.com/ibuilding-x/driver-box/internal/plugins/bacnet/bacnet/btypes"
)

func (device *Device) Whois(options *bacnet.WhoIsOpts) ([]btypes.Device, error) {
	// go device.network.ClientRun()
	resp, err := device.network.WhoIs(options)
	return resp, err
}

func (net *Network) Whois(options *bacnet.WhoIsOpts) ([]btypes.Device, error) {
	// go net.NetworkRun()
	resp, err := net.Client.WhoIs(options)
	return resp, err
}
