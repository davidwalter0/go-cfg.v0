package main

import (
	"time"
)

// Specification example test struct
type Specification struct {
	Debug        bool           `envconfig:"Debug" usage:"enable debug mode"`
	Port         int            `short:"p" default:"8080" usage:"primary ip port"`
	User         string         `usage:"user for ..."`
	UserName     string         `envconfig:"USER_NAME"`
	Users        []string       ``
	UserArray    []string       `default:"x,y,z,0,1"`
	IntArray     []int64        `default:"0,1,2,3,4"`
	Rate         float64        ``
	RateOfTravel float32        `short:"rt"  default:"3.14"`
	Timeout      time.Duration  `short:"t1"  default:"720h1m3s"`
	Timeout2     time.Duration  `short:"t2"  default:"720h1m3s" usage:"Timeout2 something"`
	Int8         int8           `short:"i8"  default:"127"      usage:"int8   test"`
	Nint8        int8           `short:"n8"  default:"-128"     usage:"nint8  test"`
	Uint8        uint8          `short:"u8"  default:"255"      usage:"uint8  test"`
	Int16        int16          `short:"i16" default:"32767"    usage:"int16  test"`
	Nint16       int16          `short:"n16" default:"-32768"   usage:"nint16 test"`
	Uint16       uint16         `short:"u16" default:"65535"    usage:"uint16 test"`
	Int32        int32          `short:"i32" default:"1048576"  usage:"int32  test"`
	Nint32       int32          `short:"n32" default:"-1232"    usage:"nint32 test"`
	Uint32       uint32         `short:"u32" default:"255"      usage:"uint32 test" required:""`
	ColorCodes   map[string]int `short:"cc"  default:"white:0xfff,black:000,red:f00,green:0f0,blue:00f"`
	// ColorCodes map[string]int     `short:"cc"  default:"white:0,black:1,red:,green:2,blue:3"`
	Map map[string]float64 `short:"m"   default:"π:3.14159,ξ:1,ρ:.01,φ:1.2,β:3,α:.01,δ:3,ε:.001,φ:.1,ψ:.9,ω:2.1"`
}
