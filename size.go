package size

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// A Size represent ...
type Size int64

const (
	minSize Size = -1 << 63
	maxSize Size = 1<<63 - 1
)

// Common size. ...
const (
	Byte  Size = 1
	KByte      = 1024 * Byte
	MByte      = 1024 * KByte
	GByte      = 1024 * MByte
	TByte      = 1024 * GByte
)

// Used to convert user input to ByteSize
var unitMap = map[string]func(float64) Size{
	"B": func(f float64) Size {
		return Size(f)
	},
	"KB": func(f float64) Size {
		return Size(f * float64(KByte))
	},
	"MB": func(f float64) Size {
		return Size(f * float64(MByte))
	},
	"GB": func(f float64) Size {
		return Size(f * float64(GByte))
	},
	"TB": func(f float64) Size {
		return Size(f * float64(TByte))
	},
}

func (d Size) Bytes() int64 {
	return int64(d)
}

func (d Size) KBytes() float64 {
	hour := d / KByte
	//nsec := d % KByte
	return float64(hour) //+ float64(nsec)/(60*60*1e9)
}

func (d Size) MBytes() float64 {
	mbytes := d / MByte
	//bytes := d % MByte
	return float64(mbytes) //+ float64(bytes)/(60*1e9)
}

func (d Size) GBytes() float64 {
	gbytes := d / GByte
	//nsec := d % Hour
	return float64(gbytes) //+ float64(nsec)/(60*60*1e9)
}

func (d Size) TBytes() float64 {
	tbytes := d / TByte
	//nsec := d % Hour
	return float64(tbytes) //+ float64(nsec)/(60*60*1e9)
}

func (d Size) String() string {
	if t := d.TBytes(); t >= 1 {
		return fmt.Sprintf("%vTB", t)
	} else if t := d.GBytes(); t >= 1 {
		return fmt.Sprintf("%GGB", t)
	} else if t := d.MBytes(); t >= 1 {
		return fmt.Sprintf("%vMB", t)
	} else if t := d.KBytes(); t >= 1 {
		return fmt.Sprintf("%vKB", t)
	}

	return fmt.Sprintf("%vB", d.Bytes())
}

func Parse(s string) (Size, error) {
	// Remove leading and trailing whitespace
	s = strings.TrimSpace(s)

	split := make([]string, 0)
	for i, r := range s {
		if !unicode.IsDigit(r) {
			// Split the string by digit and size designator, remove whitespace
			split = append(split, strings.TrimSpace(s[:i]))
			split = append(split, strings.TrimSpace(s[i:]))
			break
		}
	}

	// Check to see if we split successfully
	if len(split) != 2 {
		return 0, errors.New("Unrecognized size suffix")
	}

	cv, ok := unitMap[strings.ToUpper(split[1])]
	if !ok {
		return 0, errors.New("Unrecognized size suffix " + split[1])

	}

	value, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
		return 0, err
	}

	return cv(value), nil
}
