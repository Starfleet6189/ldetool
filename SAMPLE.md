```go
/* This file was autogenerated via
 ----------------------------------------
 ldetool generate --package main line.lde
 ----------------------------------------
do not touch it with bare hands!
*/

package main

import (
	"bytes"
	"fmt"
	"strconv"
	"unsafe"
)

var countryLsbrck = []byte("country[")
var firstLsbrck = []byte("first[")
var spaceFETCHSpace = []byte(" FETCH ")
var spaceFormatLsbrck = []byte(" format[")
var spaceHiddenLsbrck = []byte(" hidden[")
var spaceUserUscoreAgentLsbrck = []byte(" user_agent[")

// Line autogenerated parser
type Line struct {
	rest   []byte
	Time   []byte
	First  uint8
	Format []byte
	Hidden struct {
		Valid bool
		Value uint8
	}
	UserAgent []byte
	Country   []byte
}

// Extract autogenerated method of Line
func (p *Line) Extract(line []byte) (bool, error) {
	var err error
	var hiddenRest []byte
	var pos int
	var tmp []byte
	var tmpUint uint64
	p.rest = line

	// Checks if the rest starts with '[' symbol and pass it
	if len(p.rest) > 0 && p.rest[0] == '[' {
		p.rest = p.rest[1:]
	} else {
		return false, nil
	}

	// Put data before ']' into Time
	if pos = bytes.IndexByte(p.rest, ']'); pos >= 0 {
		p.Time = p.rest[:pos]

		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Checks if the rest starts with " FETCH " and pass it
	if bytes.HasPrefix(p.rest, spaceFETCHSpace) {
		p.rest = p.rest[len(spaceFETCHSpace):]
	} else {
		return false, nil
	}

	// Checks if the rest starts with "first[" and pass it
	if bytes.HasPrefix(p.rest, firstLsbrck) {
		p.rest = p.rest[len(firstLsbrck):]
	} else {
		return false, nil
	}

	// Put data before ']' into First
	if pos = bytes.IndexByte(p.rest, ']'); pos >= 0 {
		tmp = p.rest[:pos]
		if tmpUint, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&tmp)), 10, 8); err != nil {
			return false, fmt.Errorf("Error parsing \033[1m%s\033[0m value as uint8 for field \033[1mFirst\033[0m: %s", string(tmp), err)
		}
		p.First = uint8(tmpUint)

		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Checks if the rest starts with " format[" and pass it
	if bytes.HasPrefix(p.rest, spaceFormatLsbrck) {
		p.rest = p.rest[len(spaceFormatLsbrck):]
	} else {
		return false, nil
	}

	// Put data before ']' into Format
	if pos = bytes.IndexByte(p.rest, ']'); pos >= 0 {
		p.Format = p.rest[:pos]

		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}
	hiddenRest = p.rest

	// Checks if the rest starts with " hidden[" and pass it
	if bytes.HasPrefix(hiddenRest, spaceHiddenLsbrck) {
		hiddenRest = hiddenRest[len(spaceHiddenLsbrck):]
	} else {
		p.Hidden.Valid = false
		goto hiddenLabel
	}

	// Put data before ']' into Hidden.Value
	if pos = bytes.IndexByte(hiddenRest, ']'); pos >= 0 {
		tmp = hiddenRest[:pos]
		if tmpUint, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&tmp)), 10, 8); err != nil {
			return false, fmt.Errorf("Error parsing \033[1m%s\033[0m value as uint8 for field \033[1mHidden.Value\033[0m: %s", string(tmp), err)
		}
		p.Hidden.Value = uint8(tmpUint)

		hiddenRest = hiddenRest[pos+1:]
	} else {
		p.Hidden.Valid = false
		goto hiddenLabel
	}
	p.Hidden.Valid = true
	p.rest = hiddenRest
hiddenLabel:

	// Checks if the rest starts with " user_agent[" and pass it
	if bytes.HasPrefix(p.rest, spaceUserUscoreAgentLsbrck) {
		p.rest = p.rest[len(spaceUserUscoreAgentLsbrck):]
	} else {
		return false, nil
	}

	// Put data before ']' into UserAgent
	if pos = bytes.IndexByte(p.rest, ']'); pos >= 0 {
		p.UserAgent = p.rest[:pos]

		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Looking for "country[" and then pass it
	pos = bytes.Index(p.rest, countryLsbrck)
	if pos >= 0 {
		p.rest = p.rest[pos+len(countryLsbrck):]
	} else {
		return false, nil
	}

	// Put data before ']' into Country
	if pos = bytes.IndexByte(p.rest, ']'); pos >= 0 {
		p.Country = p.rest[:pos]

		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	return true, nil
}

// GetHiddenValue retrieves optional value for HiddenValue.Name
func (p *Line) GetHiddenValue() (res uint8) {
	if !p.Hidden.Valid {
		return
	}
	return p.Hidden.Value
}
```
