package cobra

// BoolFlag -----------------------------------------------------------------

type BoolFlag struct {
	SerpentFlag
}

func NewBoolFlag(n string, opts ...*FlagOptList) *BoolFlag {
	f := &BoolFlag{}
	var d bool
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *BoolFlag) OptAbbr(a string) *BoolFlag {
	f.Abbr = a
	return f
}

func (f *BoolFlag) OptDesc(a string) *BoolFlag {
	f.Desc = a
	return f
}

func (f *BoolFlag) OptReq(a bool) *BoolFlag {
	f.Req = a
	return f
}

func (f *BoolFlag) OptUbiq(a bool) *BoolFlag {
	f.Ubiq = a
	return f
}

func (f *BoolFlag) OptDefault(a interface{}) *BoolFlag {
	f.Default = a.(bool)
	return f
}

func (f *BoolFlag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().Bool(f.Name, f.Default.(bool), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().BoolP(f.Name, f.Abbr, f.Default.(bool), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().Bool(f.Name, f.Default.(bool), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().BoolP(f.Name, f.Abbr, f.Default.(bool), f.Desc)
	}
	f.postAdd(c)
}

// BytesHexFlag -----------------------------------------------------------------

type BytesHexFlag struct {
	SerpentFlag
}

func NewBytesHexFlag(n string, opts ...*FlagOptList) *BytesHexFlag {
	f := &BytesHexFlag{}
	var d []byte
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *BytesHexFlag) OptAbbr(a string) *BytesHexFlag {
	f.Abbr = a
	return f
}

func (f *BytesHexFlag) OptDesc(a string) *BytesHexFlag {
	f.Desc = a
	return f
}

func (f *BytesHexFlag) OptReq(a bool) *BytesHexFlag {
	f.Req = a
	return f
}

func (f *BytesHexFlag) OptUbiq(a bool) *BytesHexFlag {
	f.Ubiq = a
	return f
}

func (f *BytesHexFlag) OptDefault(a interface{}) *BytesHexFlag {
	f.Default = a.([]byte)
	return f
}

func (f *BytesHexFlag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().BytesHex(f.Name, f.Default.([]byte), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().BytesHexP(f.Name, f.Abbr, f.Default.([]byte), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().BytesHex(f.Name, f.Default.([]byte), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().BytesHexP(f.Name, f.Abbr, f.Default.([]byte), f.Desc)
	}
	f.postAdd(c)
}

// Float32Flag -----------------------------------------------------------------

type Float32Flag struct {
	SerpentFlag
}

func NewFloat32Flag(n string, opts ...*FlagOptList) *Float32Flag {
	f := &Float32Flag{}
	var d float32
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *Float32Flag) OptAbbr(a string) *Float32Flag {
	f.Abbr = a
	return f
}

func (f *Float32Flag) OptDesc(a string) *Float32Flag {
	f.Desc = a
	return f
}

func (f *Float32Flag) OptReq(a bool) *Float32Flag {
	f.Req = a
	return f
}

func (f *Float32Flag) OptUbiq(a bool) *Float32Flag {
	f.Ubiq = a
	return f
}

func (f *Float32Flag) OptDefault(a interface{}) *Float32Flag {
	f.Default = a.(float32)
	return f
}

func (f *Float32Flag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().Float32(f.Name, f.Default.(float32), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().Float32P(f.Name, f.Abbr, f.Default.(float32), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().Float32(f.Name, f.Default.(float32), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().Float32P(f.Name, f.Abbr, f.Default.(float32), f.Desc)
	}
	f.postAdd(c)
}

// Float64Flag -----------------------------------------------------------------

type Float64Flag struct {
	SerpentFlag
}

func NewFloat64Flag(n string, opts ...*FlagOptList) *Float64Flag {
	f := &Float64Flag{}
	var d float64
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *Float64Flag) OptAbbr(a string) *Float64Flag {
	f.Abbr = a
	return f
}

func (f *Float64Flag) OptDesc(a string) *Float64Flag {
	f.Desc = a
	return f
}

func (f *Float64Flag) OptReq(a bool) *Float64Flag {
	f.Req = a
	return f
}

func (f *Float64Flag) OptUbiq(a bool) *Float64Flag {
	f.Ubiq = a
	return f
}

func (f *Float64Flag) OptDefault(a interface{}) *Float64Flag {
	f.Default = a.(float64)
	return f
}

func (f *Float64Flag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().Float64(f.Name, f.Default.(float64), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().Float64P(f.Name, f.Abbr, f.Default.(float64), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().Float64(f.Name, f.Default.(float64), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().Float64P(f.Name, f.Abbr, f.Default.(float64), f.Desc)
	}
	f.postAdd(c)
}

// IntFlag -----------------------------------------------------------------

type IntFlag struct {
	SerpentFlag
}

func NewIntFlag(n string, opts ...*FlagOptList) *IntFlag {
	f := &IntFlag{}
	var d int
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *IntFlag) OptAbbr(a string) *IntFlag {
	f.Abbr = a
	return f
}

func (f *IntFlag) OptDesc(a string) *IntFlag {
	f.Desc = a
	return f
}

func (f *IntFlag) OptReq(a bool) *IntFlag {
	f.Req = a
	return f
}

func (f *IntFlag) OptUbiq(a bool) *IntFlag {
	f.Ubiq = a
	return f
}

func (f *IntFlag) OptDefault(a interface{}) *IntFlag {
	f.Default = a.(int)
	return f
}

func (f *IntFlag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().Int(f.Name, f.Default.(int), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().IntP(f.Name, f.Abbr, f.Default.(int), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().Int(f.Name, f.Default.(int), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().IntP(f.Name, f.Abbr, f.Default.(int), f.Desc)
	}
	f.postAdd(c)
}

// Int16Flag -----------------------------------------------------------------

type Int16Flag struct {
	SerpentFlag
}

func NewInt16Flag(n string, opts ...*FlagOptList) *Int16Flag {
	f := &Int16Flag{}
	var d int16
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *Int16Flag) OptAbbr(a string) *Int16Flag {
	f.Abbr = a
	return f
}

func (f *Int16Flag) OptDesc(a string) *Int16Flag {
	f.Desc = a
	return f
}

func (f *Int16Flag) OptReq(a bool) *Int16Flag {
	f.Req = a
	return f
}

func (f *Int16Flag) OptUbiq(a bool) *Int16Flag {
	f.Ubiq = a
	return f
}

func (f *Int16Flag) OptDefault(a interface{}) *Int16Flag {
	f.Default = a.(int16)
	return f
}

func (f *Int16Flag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().Int16(f.Name, f.Default.(int16), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().Int16P(f.Name, f.Abbr, f.Default.(int16), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().Int16(f.Name, f.Default.(int16), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().Int16P(f.Name, f.Abbr, f.Default.(int16), f.Desc)
	}
	f.postAdd(c)
}

// Int32Flag -----------------------------------------------------------------

type Int32Flag struct {
	SerpentFlag
}

func NewInt32Flag(n string, opts ...*FlagOptList) *Int32Flag {
	f := &Int32Flag{}
	var d int32
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *Int32Flag) OptAbbr(a string) *Int32Flag {
	f.Abbr = a
	return f
}

func (f *Int32Flag) OptDesc(a string) *Int32Flag {
	f.Desc = a
	return f
}

func (f *Int32Flag) OptReq(a bool) *Int32Flag {
	f.Req = a
	return f
}

func (f *Int32Flag) OptUbiq(a bool) *Int32Flag {
	f.Ubiq = a
	return f
}

func (f *Int32Flag) OptDefault(a interface{}) *Int32Flag {
	f.Default = a.(int32)
	return f
}

func (f *Int32Flag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().Int32(f.Name, f.Default.(int32), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().Int32P(f.Name, f.Abbr, f.Default.(int32), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().Int32(f.Name, f.Default.(int32), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().Int32P(f.Name, f.Abbr, f.Default.(int32), f.Desc)
	}
	f.postAdd(c)
}

// Int64Flag -----------------------------------------------------------------

type Int64Flag struct {
	SerpentFlag
}

func NewInt64Flag(n string, opts ...*FlagOptList) *Int64Flag {
	f := &Int64Flag{}
	var d int64
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *Int64Flag) OptAbbr(a string) *Int64Flag {
	f.Abbr = a
	return f
}

func (f *Int64Flag) OptDesc(a string) *Int64Flag {
	f.Desc = a
	return f
}

func (f *Int64Flag) OptReq(a bool) *Int64Flag {
	f.Req = a
	return f
}

func (f *Int64Flag) OptUbiq(a bool) *Int64Flag {
	f.Ubiq = a
	return f
}

func (f *Int64Flag) OptDefault(a interface{}) *Int64Flag {
	f.Default = a.(int64)
	return f
}

func (f *Int64Flag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().Int64(f.Name, f.Default.(int64), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().Int64P(f.Name, f.Abbr, f.Default.(int64), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().Int64(f.Name, f.Default.(int64), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().Int64P(f.Name, f.Abbr, f.Default.(int64), f.Desc)
	}
	f.postAdd(c)
}

// Int8Flag -----------------------------------------------------------------

type Int8Flag struct {
	SerpentFlag
}

func NewInt8Flag(n string, opts ...*FlagOptList) *Int8Flag {
	f := &Int8Flag{}
	var d int8
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *Int8Flag) OptAbbr(a string) *Int8Flag {
	f.Abbr = a
	return f
}

func (f *Int8Flag) OptDesc(a string) *Int8Flag {
	f.Desc = a
	return f
}

func (f *Int8Flag) OptReq(a bool) *Int8Flag {
	f.Req = a
	return f
}

func (f *Int8Flag) OptUbiq(a bool) *Int8Flag {
	f.Ubiq = a
	return f
}

func (f *Int8Flag) OptDefault(a interface{}) *Int8Flag {
	f.Default = a.(int8)
	return f
}

func (f *Int8Flag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().Int8(f.Name, f.Default.(int8), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().Int8P(f.Name, f.Abbr, f.Default.(int8), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().Int8(f.Name, f.Default.(int8), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().Int8P(f.Name, f.Abbr, f.Default.(int8), f.Desc)
	}
	f.postAdd(c)
}

// StringFlag -----------------------------------------------------------------

type StringFlag struct {
	SerpentFlag
}

func NewStringFlag(n string, opts ...*FlagOptList) *StringFlag {
	f := &StringFlag{}
	var d string
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *StringFlag) OptAbbr(a string) *StringFlag {
	f.Abbr = a
	return f
}

func (f *StringFlag) OptDesc(a string) *StringFlag {
	f.Desc = a
	return f
}

func (f *StringFlag) OptReq(a bool) *StringFlag {
	f.Req = a
	return f
}

func (f *StringFlag) OptUbiq(a bool) *StringFlag {
	f.Ubiq = a
	return f
}

func (f *StringFlag) OptDefault(a interface{}) *StringFlag {
	f.Default = a.(string)
	return f
}

func (f *StringFlag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().String(f.Name, f.Default.(string), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().StringP(f.Name, f.Abbr, f.Default.(string), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().String(f.Name, f.Default.(string), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().StringP(f.Name, f.Abbr, f.Default.(string), f.Desc)
	}
	f.postAdd(c)
}

// StringSliceFlag -----------------------------------------------------------------

type StringSliceFlag struct {
	SerpentFlag
}

func NewStringSliceFlag(n string, opts ...*FlagOptList) *StringSliceFlag {
	f := &StringSliceFlag{}
	var d []string
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *StringSliceFlag) OptAbbr(a string) *StringSliceFlag {
	f.Abbr = a
	return f
}

func (f *StringSliceFlag) OptDesc(a string) *StringSliceFlag {
	f.Desc = a
	return f
}

func (f *StringSliceFlag) OptReq(a bool) *StringSliceFlag {
	f.Req = a
	return f
}

func (f *StringSliceFlag) OptUbiq(a bool) *StringSliceFlag {
	f.Ubiq = a
	return f
}

func (f *StringSliceFlag) OptDefault(a interface{}) *StringSliceFlag {
	f.Default = a.([]string)
	return f
}

func (f *StringSliceFlag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().StringSlice(f.Name, f.Default.([]string), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().StringSliceP(f.Name, f.Abbr, f.Default.([]string), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().StringSlice(f.Name, f.Default.([]string), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().StringSliceP(f.Name, f.Abbr, f.Default.([]string), f.Desc)
	}
	f.postAdd(c)
}

// UintFlag -----------------------------------------------------------------

type UintFlag struct {
	SerpentFlag
}

func NewUintFlag(n string, opts ...*FlagOptList) *UintFlag {
	f := &UintFlag{}
	var d uint
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *UintFlag) OptAbbr(a string) *UintFlag {
	f.Abbr = a
	return f
}

func (f *UintFlag) OptDesc(a string) *UintFlag {
	f.Desc = a
	return f
}

func (f *UintFlag) OptReq(a bool) *UintFlag {
	f.Req = a
	return f
}

func (f *UintFlag) OptUbiq(a bool) *UintFlag {
	f.Ubiq = a
	return f
}

func (f *UintFlag) OptDefault(a interface{}) *UintFlag {
	f.Default = a.(uint)
	return f
}

func (f *UintFlag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().Uint(f.Name, f.Default.(uint), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().UintP(f.Name, f.Abbr, f.Default.(uint), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().Uint(f.Name, f.Default.(uint), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().UintP(f.Name, f.Abbr, f.Default.(uint), f.Desc)
	}
	f.postAdd(c)
}

// Uint16Flag -----------------------------------------------------------------

type Uint16Flag struct {
	SerpentFlag
}

func NewUint16Flag(n string, opts ...*FlagOptList) *Uint16Flag {
	f := &Uint16Flag{}
	var d uint16
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *Uint16Flag) OptAbbr(a string) *Uint16Flag {
	f.Abbr = a
	return f
}

func (f *Uint16Flag) OptDesc(a string) *Uint16Flag {
	f.Desc = a
	return f
}

func (f *Uint16Flag) OptReq(a bool) *Uint16Flag {
	f.Req = a
	return f
}

func (f *Uint16Flag) OptUbiq(a bool) *Uint16Flag {
	f.Ubiq = a
	return f
}

func (f *Uint16Flag) OptDefault(a interface{}) *Uint16Flag {
	f.Default = a.(uint16)
	return f
}

func (f *Uint16Flag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().Uint16(f.Name, f.Default.(uint16), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().Uint16P(f.Name, f.Abbr, f.Default.(uint16), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().Uint16(f.Name, f.Default.(uint16), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().Uint16P(f.Name, f.Abbr, f.Default.(uint16), f.Desc)
	}
	f.postAdd(c)
}

// Uint32Flag -----------------------------------------------------------------

type Uint32Flag struct {
	SerpentFlag
}

func NewUint32Flag(n string, opts ...*FlagOptList) *Uint32Flag {
	f := &Uint32Flag{}
	var d uint32
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *Uint32Flag) OptAbbr(a string) *Uint32Flag {
	f.Abbr = a
	return f
}

func (f *Uint32Flag) OptDesc(a string) *Uint32Flag {
	f.Desc = a
	return f
}

func (f *Uint32Flag) OptReq(a bool) *Uint32Flag {
	f.Req = a
	return f
}

func (f *Uint32Flag) OptUbiq(a bool) *Uint32Flag {
	f.Ubiq = a
	return f
}

func (f *Uint32Flag) OptDefault(a interface{}) *Uint32Flag {
	f.Default = a.(uint32)
	return f
}

func (f *Uint32Flag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().Uint32(f.Name, f.Default.(uint32), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().Uint32P(f.Name, f.Abbr, f.Default.(uint32), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().Uint32(f.Name, f.Default.(uint32), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().Uint32P(f.Name, f.Abbr, f.Default.(uint32), f.Desc)
	}
	f.postAdd(c)
}

// Uint64Flag -----------------------------------------------------------------

type Uint64Flag struct {
	SerpentFlag
}

func NewUint64Flag(n string, opts ...*FlagOptList) *Uint64Flag {
	f := &Uint64Flag{}
	var d uint64
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *Uint64Flag) OptAbbr(a string) *Uint64Flag {
	f.Abbr = a
	return f
}

func (f *Uint64Flag) OptDesc(a string) *Uint64Flag {
	f.Desc = a
	return f
}

func (f *Uint64Flag) OptReq(a bool) *Uint64Flag {
	f.Req = a
	return f
}

func (f *Uint64Flag) OptUbiq(a bool) *Uint64Flag {
	f.Ubiq = a
	return f
}

func (f *Uint64Flag) OptDefault(a interface{}) *Uint64Flag {
	f.Default = a.(uint64)
	return f
}

func (f *Uint64Flag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().Uint64(f.Name, f.Default.(uint64), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().Uint64P(f.Name, f.Abbr, f.Default.(uint64), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().Uint64(f.Name, f.Default.(uint64), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().Uint64P(f.Name, f.Abbr, f.Default.(uint64), f.Desc)
	}
	f.postAdd(c)
}

// Uint8Flag -----------------------------------------------------------------

type Uint8Flag struct {
	SerpentFlag
}

func NewUint8Flag(n string, opts ...*FlagOptList) *Uint8Flag {
	f := &Uint8Flag{}
	var d uint8
	f.Default = d
	f.SerpentFlag.populate(n, d, opts...)
	return f
}

func (f *Uint8Flag) OptAbbr(a string) *Uint8Flag {
	f.Abbr = a
	return f
}

func (f *Uint8Flag) OptDesc(a string) *Uint8Flag {
	f.Desc = a
	return f
}

func (f *Uint8Flag) OptReq(a bool) *Uint8Flag {
	f.Req = a
	return f
}

func (f *Uint8Flag) OptUbiq(a bool) *Uint8Flag {
	f.Ubiq = a
	return f
}

func (f *Uint8Flag) OptDefault(a interface{}) *Uint8Flag {
	f.Default = a.(uint8)
	return f
}

func (f *Uint8Flag) AddTo(c *Command) {
	switch {
	case f.Ubiq && len(f.Abbr) == 0:
		c.PersistentFlags().Uint8(f.Name, f.Default.(uint8), f.Desc)
	case f.Ubiq && len(f.Abbr) == 1:
		c.PersistentFlags().Uint8P(f.Name, f.Abbr, f.Default.(uint8), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 0:
		c.Flags().Uint8(f.Name, f.Default.(uint8), f.Desc)
	case !f.Ubiq && len(f.Abbr) == 1:
		c.Flags().Uint8P(f.Name, f.Abbr, f.Default.(uint8), f.Desc)
	}
	f.postAdd(c)
}
