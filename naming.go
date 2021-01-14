// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package embedx

var (
	namer Namer = NamerFunc(func(name string) string {
		return name
	})
)

// Namer make a new name from an old name.
type Namer interface {
	Naming(string) string
}

// NamerFunc wrap a func as Namer
type NamerFunc func(string) string

// Naming rename give name to a new name.
func (f NamerFunc) Naming(name string) string {
	return f(name)
}

// RegisterNamer register a new namer replace default.
// Note: this function is not concurrent safe.
func RegisterNamer(n Namer) {
	namer = n
}

// Naming rename give name to a new name use default namer
func Naming(name string) string {
	return namer.Naming(name)
}
