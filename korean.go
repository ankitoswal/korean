// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package korean provides Korean encodings such as EUC-KR and CP949.
// It is a wrapper of code.google.com/p/go.text/encoding/korean package
// for easy to use.

package korean

import (
	"reflect"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

const (
	defaultBufSize = 4096
)

// trans writes to dst the transformed bytes read from src, and
// returns dst bytes slice written and src bytes read.
func trans(t transform.Transformer, src []byte) (dst []byte, err error) {
	switch reflect.TypeOf(t).Name() {
	case "eucKRDecoder":
		dst = make([]byte, len(src)+len(src)/2)
	case "eucKREncoder":
		dst = make([]byte, len(src))
	}

	for {
		nDst, _, err := t.Transform(dst, src, true)
		if err != nil {
			// Destination buffer was too short to receive
			// all of the transformed bytes.
			if err == transform.ErrShortDst {
				dst = make([]byte, len(dst)+defaultBufSize)
				continue
			} else {
				return nil, err
			}
		}
		return dst[:nDst], nil
	}
}

// UTF8 converts from EUC-KR bytes to UTF-8 bytes.
func UTF8(src []byte) (dst []byte, err error) {
	return trans(korean.EUCKR.NewDecoder(), src)
}

// EUCKR converts from UTF-8 bytes to EUC-KR bytes.
func EUCKR(src []byte) (dst []byte, err error) {
	return trans(korean.EUCKR.NewEncoder(), src)
}
