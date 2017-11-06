/* package stamp provides types and functions to work with date 
and time that is stored as an integer like 201701201 as YYYQMMDD.
It provides a very memory efficient, super fast, and simple way to
stamp calendar event that is sortable, comparable, and human-readable.
Type "Stamp" is uint64, there are contansts from Year to Second to pass as argument.
you can compare stamp1<stamp2 as they are simple integer from Year to Second,
you can "Extract" a period like Month, or all periods, Convert to other formats,
make it ToTime in package time, or FromTime.
*/
package stamp()
import (
	"time"
)

type Stamp uint64 // format: YYYYQMMDDHHmmSS like 201701201 044500 in which 0 for Q means not recorded quarterly
type Stamps []uint64
type Periods struct {
	Year, Quarter, Month, Day, Weekday, Holiday, Hour, Minute, Second, Frac int
}

const (Year=1; Quarter = 1e1; Month = 1e3; Day = 1e5; Holiday = 1e6; Weekday= 1e7; Cal = 1e7; Hour = 1e9; Minute=1e11; Second = 1e13; Frac = 1e15)

/* to retrieve any part of the Stamp such as year, month, quarter, ... */
func (s Stamp) Part (period, format Stamp) int {
	out := s / (format / period)
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
	return int(out)
}

/* Converts s Stamp from one format to another format */
func (s Stamp) Convert (from, to Stamp) Stamp {
	return to / from * s
}

/* Returns all period parts of receiver s, with given argument format. */
func (s Stamp) AllParts (format Stamp) Periods {
	return Periods{s.Part(Year, format), s.Part(Quarter, format), s.Part(Month, format), s.Part(Day, format), s.Part(Weekday, format), s.Part(Holiday, format), s.Part(Hour, format), s.Part(Minute, format), s.Part(Second, format), s.Part(Frac, format)}
} 

/* Converts s, with given format, to time.Time type. */
func (s Stamp) ToTime (format Stamp) time.Time {
	parts :=s.AllParts (format)
	return time.Date (parts.Year, time.Month(parts.Month), parts.Day, parts.Hour, parts.Minute, parts.Second, parts.SS*1e6, time.UTC)
}

/* Converts from time.Time to Stamp */
func FromTime (t time.Time) Stamp {
	return Stamp(1e15*(t.Year()/Year + int(t.Month())/Month + t.Day()/Day + t.Hour()/Hour + t.Minute()/Minute + t.Second()/Second))
} 


/* Return Stamp parts for a given integer. 
func Part (what int, period int) []int {
	//pass
	return []int{0}
}
*/

/* Return period part of a given format for the receiver Stamps. */
func (s Stamps) Part (period, format int) []int {
	part := 1000000
	out := make([]int, len(s))
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

