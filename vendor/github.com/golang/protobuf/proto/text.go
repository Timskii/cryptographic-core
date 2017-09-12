// Go support for Protocol Buffers - Google's data interchange format
//
// Copyright 2010 The Go Authors.  All rights reserved.
// https://github.com/golang/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package proto

// Functions for writing the text protocol buffer format.

import (

	"bytes"
	"io"
	"log"
	"strings"
)

var (
	newline         = []byte("\n")
	spaces          = []byte("                                        ")
)

type writer interface {
	io.Writer
	WriteByte(byte) error
}

// textWriter is an io.Writer that tracks its indentation level.
type textWriter struct {
	ind      int
	complete bool // if the current position is a complete line
	compact  bool // whether to write out as a one-liner
	w        writer
}

func (w *textWriter) WriteString(s string) (n int, err error) {
	if !strings.Contains(s, "\n") {
		if !w.compact && w.complete {
			w.writeIndent()
		}
		w.complete = false
		return io.WriteString(w.w, s)
	}
	// WriteString is typically called without newlines, so this
	// codepath and its copy are rare.  We copy to avoid
	// duplicating all of Write's logic here.
	return w.Write([]byte(s))
}

func (w *textWriter) Write(p []byte) (n int, err error) {
	newlines := bytes.Count(p, newline)
	if newlines == 0 {
		if !w.compact && w.complete {
			w.writeIndent()
		}
		n, err = w.w.Write(p)
		w.complete = false
		return n, err
	}

	frags := bytes.SplitN(p, newline, newlines+1)
	if w.compact {
		for i, frag := range frags {
			if i > 0 {
				if err := w.w.WriteByte(' '); err != nil {
					return n, err
				}
				n++
			}
			nn, err := w.w.Write(frag)
			n += nn
			if err != nil {
				return n, err
			}
		}
		return n, nil
	}

	for i, frag := range frags {
		if w.complete {
			w.writeIndent()
		}
		nn, err := w.w.Write(frag)
		n += nn
		if err != nil {
			return n, err
		}
		if i+1 < len(frags) {
			if err := w.w.WriteByte('\n'); err != nil {
				return n, err
			}
			n++
		}
	}
	w.complete = len(frags[len(frags)-1]) == 0
	return n, nil
}

func (w *textWriter) WriteByte(c byte) error {
	if w.compact && c == '\n' {
		c = ' '
	}
	if !w.compact && w.complete {
		w.writeIndent()
	}
	err := w.w.WriteByte(c)
	w.complete = c == '\n'
	return err
}

func (w *textWriter) indent() { w.ind++ }

func (w *textWriter) unindent() {
	if w.ind == 0 {
		log.Print("proto: textWriter unindented too far")
		return
	}
	w.ind--
}

type int32Slice []int32

func (s int32Slice) Len() int           { return len(s) }
func (s int32Slice) Less(i, j int) bool { return s[i] < s[j] }
func (s int32Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }



func (w *textWriter) writeIndent() {
	if !w.complete {
		return
	}
	remain := w.ind * 2
	for remain > 0 {
		n := remain
		if n > len(spaces) {
			n = len(spaces)
		}
		w.w.Write(spaces[:n])
		remain -= n
	}
	w.complete = false
}

// TextMarshaler is a configurable text format marshaler.
type TextMarshaler struct {
	Compact   bool // use compact text format (one line).
	ExpandAny bool // expand google.protobuf.Any messages of known types
}

// Marshal writes a given protocol buffer in text format.
// The only errors returned are from w.


// Text is the same as Marshal, but returns the string directly.
func (tm *TextMarshaler) Text(pb Message) string {
	var buf bytes.Buffer
	//tm.Marshal(&buf, pb)
	return buf.String()
}

var (
	defaultTextMarshaler = TextMarshaler{}
	compactTextMarshaler = TextMarshaler{Compact: true}
)

// TODO: consider removing some of the Marshal functions below.



// CompactTextString is the same as CompactText, but returns the string directly.
func CompactTextString(pb Message) string { return compactTextMarshaler.Text(pb) }
