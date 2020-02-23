package uint8_tree

import "errors"

// code common to the IPv4/IPv6 trees

// MatchesFunc is called to check if tag data matches the input value
type MatchesFunc func(payload uint8, val uint8) bool

// FilterFunc is called on each result to see if it belongs in the resulting set
type FilterFunc func(payload uint8) bool
var errStop = errors.New("stop")
