
/*
 This file was autogenerated via
 --------------------------------------------------------
 ldetool generate --little-endian --package main rule.lde
 --------------------------------------------------------
 do not touch it with bare hands!
*/

package main

import ()

// Line ...
type Line struct {
	rest  []byte
	Name  []byte
	Count []byte
}

// Extract ...
func (p *Line) Extract(line []byte) (bool, error) {
	p.rest = line
	var pos int

	// Take until '|' as Name(string)
	pos = -1
	for i, char := range p.rest {
		if char == '|' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.Name = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Looking for '|' and then pass it
	pos = -1
	for i, char := range p.rest {
		if char == '|' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Looking for '|' and then pass it
	pos = -1
	for i, char := range p.rest {
		if char == '|' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Looking for '|' and then pass it
	pos = -1
	for i, char := range p.rest {
		if char == '|' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Take until '|' as Count(string)
	pos = -1
	for i, char := range p.rest {
		if char == '|' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.Count = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	return true, nil
}
