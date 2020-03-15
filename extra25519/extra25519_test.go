// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package extra25519

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/teserakt-io/golang-ed25519"
	"golang.org/x/crypto/curve25519"
)

func TestCurve25519Conversion(t *testing.T) {
	public_orig, private_orig, _ := ed25519.GenerateKey(rand.Reader)

	var private [64]byte
	var public [32]byte
	copy(private[:], private_orig)
	copy(public[:], public_orig)

	var curve25519Public, curve25519Public2, curve25519Private [32]byte
	PrivateKeyToCurve25519(&curve25519Private, &private)
	curve25519.ScalarBaseMult(&curve25519Public, &curve25519Private)

	if !PublicKeyToCurve25519(&curve25519Public2, &public) {
		t.Fatalf("PublicKeyToCurve25519 failed")
	}

	if !bytes.Equal(curve25519Public[:], curve25519Public2[:]) {
		t.Errorf("Values didn't match: curve25519 produced %x, conversion produced %x", curve25519Public[:], curve25519Public2[:])
	}
}
