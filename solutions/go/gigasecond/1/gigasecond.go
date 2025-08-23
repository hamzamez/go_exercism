// This adds gigasecond to a date.
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	// we will just use Time.Add function
    
	return t.Add(time.Second * 1000000000)
}
