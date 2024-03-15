package datalink

import (
	"github.com/ibuilding-x/driver-box/internal/plugins/bacnet/bacnet/btypes"
)

type DataLink interface {
	GetMyAddress() *btypes.Address
	GetBroadcastAddress() *btypes.Address
	Send(data []byte, npdu *btypes.NPDU, dest *btypes.Address) (int, error)
	Receive(data []byte) (*btypes.Address, int, error)
	Close() error
}
