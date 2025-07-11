// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package ytdlp

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func wrapError(r *Result, err error) (*Result, error) {
	if err == nil {
		return r, nil
	}

	if r == nil {
		return nil, &ErrUnknown{wrapped: err}
	}

	err = r.decorateError(err)

	if errors.Is(err, exec.ErrDot) || errors.Is(err, exec.ErrNotFound) {
		return r, &ErrMisconfig{wrapped: err, result: r}
	}

	if r.ExitCode != 0 {
		return r, &ErrExitCode{wrapped: err, result: r}
	}

	if strings.Contains(r.Stderr, "error: no such option") {
		return r, &ErrParsing{wrapped: err, result: r}
	}

	return r, &ErrUnknown{wrapped: err}
}

// ErrExitCode is returned when the exit code of the yt-dlp process is non-zero.
type ErrExitCode struct {
	wrapped error
	result  *Result
}

func (e *ErrExitCode) Unwrap() error {
	return e.wrapped
}

func (e *ErrExitCode) Error() string {
	return fmt.Sprintf("exit code %d: %s", e.result.ExitCode, e.wrapped)
}

// IsExitCodeError returns true when the exit code of the yt-dlp process is non-zero.
func IsExitCodeError(err error) (*ErrExitCode, bool) {
	var e *ErrExitCode
	return e, errors.As(err, &e)
}

// ErrMisconfig is returned when the yt-dlp executable is not found, or is not
// configured properly.
type ErrMisconfig struct {
	wrapped error
	result  *Result
}

func (e *ErrMisconfig) Unwrap() error {
	return e.wrapped
}

func (e *ErrMisconfig) Error() string {
	return fmt.Sprintf("misconfiguration error (executable: %q): %s", e.result.Executable, e.wrapped)
}

// IsMisconfigError returns true when the yt-dlp executable is not found, or is not
// configured properly.
func IsMisconfigError(err error) (*ErrMisconfig, bool) {
	var e *ErrMisconfig
	return e, errors.As(err, &e)
}

// ErrParsing is returned when the yt-dlp process fails due to an invalid flag or
// argument, possibly due to a version mismatch or go-ytdlp bug.
type ErrParsing struct {
	wrapped error
	result  *Result
}

func (e *ErrParsing) Unwrap() error {
	return e.wrapped
}

func (e *ErrParsing) Error() string {
	return fmt.Sprintf(
		"parsing error (yt-dlp version might be too different, go-ytdlp version built with yt-dlp %s/%s): %s",
		Channel,
		Version,
		e.wrapped,
	)
}

// IsParsingError returns true when the yt-dlp process fails due to an invalid flag or
// argument, possibly due to a version mismatch or go-ytdlp bug.
func IsParsingError(err error) (*ErrParsing, bool) {
	var e *ErrParsing
	return e, errors.As(err, &e)
}

// ErrUnknown is returned when the error is unknown according to go-ytdlp.
type ErrUnknown struct {
	wrapped error
}

func (e *ErrUnknown) Unwrap() error {
	return e.wrapped
}

func (e *ErrUnknown) Error() string {
	return e.wrapped.Error()
}

// IsUnknownError returns true when the error is unknown according to go-ytdlp.
func IsUnknownError(err error) (*ErrUnknown, bool) {
	var e *ErrUnknown
	return e, errors.As(err, &e)
}

type ErrJSONParsingFlag struct {
	ID       string `json:"id,omitempty"`
	JSONPath string `json:"json_path,omitempty"`
	Flag     string `json:"flag,omitempty"`
	Err      error  `json:"error,omitempty"`
}

func (e *ErrJSONParsingFlag) Unwrap() error {
	return e.Err
}

func (e *ErrJSONParsingFlag) Error() string {
	return fmt.Sprintf(
		"error while parsing json at path %q (flag: %q, id: %q): %s",
		e.JSONPath,
		e.Flag,
		e.ID,
		e.Err,
	)
}

// IsJSONParsingFlagError returns true when the error is a JSON parsing error.
func IsJSONParsingFlagError(err error) (*ErrJSONParsingFlag, bool) {
	var e *ErrJSONParsingFlag
	return e, errors.As(err, &e)
}

type ErrMultipleJSONParsingFlags struct {
	Errors []*ErrJSONParsingFlag `json:"errors,omitempty"`
}

func (e *ErrMultipleJSONParsingFlags) Error() string {
	return fmt.Sprintf("multiple errors while parsing json: %s", e.Errors)
}

func IsMultipleJSONParsingFlagsError(err error) (*ErrMultipleJSONParsingFlags, bool) {
	var e *ErrMultipleJSONParsingFlags
	return e, errors.As(err, &e)
}
