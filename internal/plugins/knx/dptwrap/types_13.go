package dptwrap

import (
	"github.com/ibuilding-x/driver-box/internal/plugins/knx/knx/dpt"
	"github.com/spf13/cast"
)

// DPT_13001 represents DPT 13.001 / counter value (pulses).
// type DPT_13001 int32
type DPT_13001_Wrap struct {
	src dpt.DPT_13001
}

func (d *DPT_13001_Wrap) Name() string {
	return "counter pulses(int32)"
}

func (d *DPT_13001_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_13001_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_13001_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_13001_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_13001_Wrap) Set(i interface{}) {
	b := cast.ToInt32(i)
	d.src = dpt.DPT_13001(b)
}

func (d *DPT_13001_Wrap) Get() interface{} {
	return int32(d.src)
}

// DPT_13002 represents DPT 13.002 / flow rate (m^3/h).
// type DPT_13002 int32
type DPT_13002_Wrap struct {
	src dpt.DPT_13002
}

func (d *DPT_13002_Wrap) Name() string {
	return "flow rate(m^3/h)"
}

func (d *DPT_13002_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_13002_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_13002_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_13002_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_13002_Wrap) Set(i interface{}) {
	b := cast.ToInt32(i)
	d.src = dpt.DPT_13002(b)
}

func (d *DPT_13002_Wrap) Get() interface{} {
	return int32(d.src)
}

// DPT_13010 represents DPT 13.010 / active energy (Wh).
// type DPT_13010 int32
type DPT_13010_Wrap struct {
	src dpt.DPT_13010
}

func (d *DPT_13010_Wrap) Name() string {
	return "active energy(Wh)"
}

func (d *DPT_13010_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_13010_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_13010_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_13010_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_13010_Wrap) Set(i interface{}) {
	b := cast.ToInt32(i)
	d.src = dpt.DPT_13010(b)
}

func (d *DPT_13010_Wrap) Get() interface{} {
	return int32(d.src)
}

// DPT_13011 represents DPT 13.011 / apparant energy (VAh).
// type DPT_13011 int32
type DPT_13011_Wrap struct {
	src dpt.DPT_13011
}

func (d *DPT_13011_Wrap) Name() string {
	return "apparant energy(VAh)"
}

func (d *DPT_13011_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_13011_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_13011_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_13011_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_13011_Wrap) Set(i interface{}) {
	b := cast.ToInt32(i)
	d.src = dpt.DPT_13011(b)
}

func (d *DPT_13011_Wrap) Get() interface{} {
	return int32(d.src)
}

// DPT_13013 represents DPT 13.013 / active energy (kWh).
// type DPT_13013 int32
type DPT_13013_Wrap struct {
	src dpt.DPT_13013
}

func (d *DPT_13013_Wrap) Name() string {
	return "active energy(kWh)"
}

func (d *DPT_13013_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_13013_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_13013_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_13013_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_13013_Wrap) Set(i interface{}) {
	b := cast.ToInt32(i)
	d.src = dpt.DPT_13013(b)
}

func (d *DPT_13013_Wrap) Get() interface{} {
	return int32(d.src)
}

// DPT_13014 represents DPT 13.014 / apparant energy (kVAh).
// type DPT_13014 int32
type DPT_13014_Wrap struct {
	src dpt.DPT_13014
}

func (d *DPT_13014_Wrap) Name() string {
	return "apparant energy(kVAh)"
}

func (d *DPT_13014_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_13014_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_13014_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_13014_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_13014_Wrap) Set(i interface{}) {
	b := cast.ToInt32(i)
	d.src = dpt.DPT_13014(b)
}

func (d *DPT_13014_Wrap) Get() interface{} {
	return int32(d.src)
}

// DPT_13016 represents DPT 13.016 / active energy (MWh).
// type DPT_13016 int32
type DPT_13016_Wrap struct {
	src dpt.DPT_13016
}

func (d *DPT_13016_Wrap) Name() string {
	return "active energy(MWh)"
}

func (d *DPT_13016_Wrap) Pack() []byte {
	return d.src.Pack()
}

func (d *DPT_13016_Wrap) Unpack(data []byte) error {
	return d.src.Unpack(data)
}

func (d *DPT_13016_Wrap) Unit() string {
	return d.src.Unit()
}

func (d *DPT_13016_Wrap) String() string {
	return d.src.String()
}

func (d *DPT_13016_Wrap) Set(i interface{}) {
	b := cast.ToInt32(i)
	d.src = dpt.DPT_13016(b)
}

func (d *DPT_13016_Wrap) Get() interface{} {
	return int32(d.src)
}
