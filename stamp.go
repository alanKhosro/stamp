// This package provides functionality to stamp events and calendar periods 
// in a simple, portable, fast, and human-readable way with int type such as 20171201. 
//
// One common way to stamp date-time events is to use int type.
// Examples: 20171201 for Year-Month-Day; 20171 for  year-Quarter; 20171201130100 for YMDHmS.
// It is widely used since it is simple, portable, and human-readable and you can sort them correctly.
//
// The most comprehensive stamp can be YYYYQMMDDHWhhmmssff which 
// ff is the decimal fraction of second, W is the weekday, H is if it is holiday.  
//
// Type "Stamp" is int, stamp.Year to stamp.Second are constants to pass as arguments to functions.
// you can extract any "Part" like Month, or "AllParts"; or "Convert" one Stamp to another Stamp format; 
//ToTime change Stamp to package "time" type, FromTime works the way round.
//
// For common tasks such as retriving a part of date, it is ~50 times faster than time.Time type. 
package stamp
import (
	"time"
)

type Stamp int // format: YYYYQMMDDHHmmSS like 201701201 044500 in which 0 for Q means not recorded quarterly
type Stamps []int
type Periods struct {
	Year, Quarter, Month, Day, Weekday, Holiday, Hour, Minute, Second, Frac int
}

// Constants to pass as argument to functions in this package.
const (Year=1; Quarter = 1e1; Month = 1e3; Day = 1e5; Holiday = 1e6; Weekday= 1e7; Cal = 1e7; Hour = 1e9; Minute=1e11; Second = 1e13; Frac = 1e15)

// To extract any part of the Stamp such as year, month, quarter, ...
func (s Stamp) Part (period, format int) int {
	out := int(s) / (format / period)
	switch period {
	case Year: // whatever remains
		out = out
	case Quarter, Weekday, Holiday: // one-digit placehoder
		out = out % 10 
	case Month, Day, Hour, Minute, Second: // two-digit place holder
		out = out % 100
	case Frac: // whatever remains
		out = out
	}
	return out
}

// Converts s Stamp from one format to another format
func (s Stamp) Convert (from, to Stamp) Stamp {
	return to / from * s
}

// Returns all period parts of receiver s, with given argument format. 
func (s Stamp) AllParts (format int) Periods {
	return Periods{s.Part(Year, format), s.Part(Quarter, format), s.Part(Month, format), s.Part(Day, format), s.Part(Weekday, format), s.Part(Holiday, format), s.Part(Hour, format), s.Part(Minute, format), s.Part(Second, format), s.Part(Frac, format)}
} 

// Converts s, with given format, to time.Time type.
func (s Stamp) ToTime (format int) time.Time {
	parts :=s.AllParts (format)
	return time.Date (parts.Year, time.Month(parts.Month), parts.Day, parts.Hour, parts.Minute, parts.Second, parts.Frac*1e6, time.UTC)
}

// Converts from time.Time to Stamp 
func FromTime (t time.Time) Stamp {
	return Stamp(1e15*(t.Year()/Year + int(t.Month())/Month + t.Day()/Day + t.Hour()/Hour + t.Minute()/Minute + t.Second()/Second))
} 


// Return period part of a given format for the receiver Stamps. 
func (s Stamps) Part (period, format int) []int {
	out := make([]int, len(s))
	part := 1000000
	switch period {
	case Quarter, Weekday, Holiday: // one-digit placehoder
		part = 10 
	case Month, Day, Hour, Minute, Second: // two-digit place holder
		part = 100
	}
	for i,v := range s {
		out[i] = v / (format / period) % part
	}
	return out
}

