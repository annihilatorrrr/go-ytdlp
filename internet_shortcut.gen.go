// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Internet Shortcut Option Group

package ytdlp

type InternetShortcutBuilder struct {
	parent *Command
}

// Then jumps back to the base command builder, if you want to add additional flags
// from another flag builder.
func (ff *InternetShortcutBuilder) Then() *Command {
	return ff.parent
}

// Write an internet shortcut file, depending on the current platform (.url,
// .webloc or .desktop). The URL may be cached by the OS
//
// WriteLink maps to cli flags: --write-link.
func (ff *InternetShortcutBuilder) WriteLink() *InternetShortcutBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "writelink",
		Flag: "--write-link",
		Args: nil,
	})
	return ff
}

// Write a .url Windows internet shortcut. The OS caches the URL based on the file
// path
//
// WriteUrlLink maps to cli flags: --write-url-link.
func (ff *InternetShortcutBuilder) WriteUrlLink() *InternetShortcutBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "writeurllink",
		Flag: "--write-url-link",
		Args: nil,
	})
	return ff
}

// Write a .webloc macOS internet shortcut
//
// WriteWeblocLink maps to cli flags: --write-webloc-link.
func (ff *InternetShortcutBuilder) WriteWeblocLink() *InternetShortcutBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "writewebloclink",
		Flag: "--write-webloc-link",
		Args: nil,
	})
	return ff
}

// Write a .desktop Linux internet shortcut
//
// WriteDesktopLink maps to cli flags: --write-desktop-link.
func (ff *InternetShortcutBuilder) WriteDesktopLink() *InternetShortcutBuilder {
	ff.parent.addFlag(&Flag{
		ID:   "writedesktoplink",
		Flag: "--write-desktop-link",
		Args: nil,
	})
	return ff
}
