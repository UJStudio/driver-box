package dptwrap

import (
	"github.com/ibuilding-x/driver-box/internal/plugins/knx/knx/dpt"
	"github.com/spf13/cast"
)

// DPT_17001 represents DPT 17.001 / Scene Number.
// type DPT_17001 uint8
type DPT_17001_Wrap struct {
	src dpt.DPT_17001
}

func (d *DPT_17001_Wrap) Name() string {
	return "Scene Number"
}

func (d *DPT_17001_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_17001_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_17001_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_17001_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_17001_Wrap) Set(i interface{}) {
	b := cast.ToUint8(i)
	d.src = dpt.DPT_17001(b)
}

func (d *DPT_17001_Wrap) Get() interface{} {
	return uint8(d.src)
}
