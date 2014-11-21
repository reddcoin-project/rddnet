// Copyright (c) 2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rddnet_test

import (
	"bytes"
	"testing"

	"github.com/reddcoin-project/rddnet"
	"github.com/davecgh/go-spew/spew"
)

// TestGenesisBlock tests the genesis block of the main network for validity by
// checking the encoded bytes and hashes.
func TestGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := rddnet.MainNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), genesisBlockBytes) {
		t.Fatalf("TestGenesisBlock: Genesis block does not appear valid - "+
			"got %v, want %v", spew.Sdump(buf.Bytes()),
			spew.Sdump(genesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash, err := rddnet.MainNetParams.GenesisBlock.BlockSha()
	if err != nil {
		t.Fatalf("BlockSha: %v", err)
	}
	if !rddnet.MainNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestGenesisBlock: Genesis block hash does not "+
			"appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(rddnet.MainNetParams.GenesisHash))
	}
}

// TestRegTestGenesisBlock tests the genesis block of the regression test
// network for validity by checking the encoded bytes and hashes.
func TestRegTestGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := rddnet.RegressionNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestRegTestGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), regTestGenesisBlockBytes) {
		t.Fatalf("TestRegTestGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(regTestGenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash, err := rddnet.RegressionNetParams.GenesisBlock.BlockSha()
	if err != nil {
		t.Errorf("BlockSha: %v", err)
	}
	if !rddnet.RegressionNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestRegTestGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(rddnet.RegressionNetParams.GenesisHash))
	}
}

// TestTestNet3GenesisBlock tests the genesis block of the test network (version
// 3) for validity by checking the encoded bytes and hashes.
func TestTestNet3GenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := rddnet.TestNet3Params.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestTestNet3GenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), testNet3GenesisBlockBytes) {
		t.Fatalf("TestTestNet3GenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(testNet3GenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash, err := rddnet.TestNet3Params.GenesisBlock.BlockSha()
	if err != nil {
		t.Fatalf("BlockSha: %v", err)
	}
	if !rddnet.TestNet3Params.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestTestNet3GenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(rddnet.TestNet3Params.GenesisHash))
	}
}

// TestSimNetGenesisBlock tests the genesis block of the simulation test network
// for validity by checking the encoded bytes and hashes.
func TestSimNetGenesisBlock(t *testing.T) {
	// Encode the genesis block to raw bytes.
	var buf bytes.Buffer
	err := rddnet.SimNetParams.GenesisBlock.Serialize(&buf)
	if err != nil {
		t.Fatalf("TestSimNetGenesisBlock: %v", err)
	}

	// Ensure the encoded block matches the expected bytes.
	if !bytes.Equal(buf.Bytes(), simNetGenesisBlockBytes) {
		t.Fatalf("TestSimNetGenesisBlock: Genesis block does not "+
			"appear valid - got %v, want %v",
			spew.Sdump(buf.Bytes()),
			spew.Sdump(simNetGenesisBlockBytes))
	}

	// Check hash of the block against expected hash.
	hash, err := rddnet.SimNetParams.GenesisBlock.BlockSha()
	if err != nil {
		t.Fatalf("BlockSha: %v", err)
	}
	if !rddnet.SimNetParams.GenesisHash.IsEqual(&hash) {
		t.Fatalf("TestSimNetGenesisBlock: Genesis block hash does "+
			"not appear valid - got %v, want %v", spew.Sdump(hash),
			spew.Sdump(rddnet.SimNetParams.GenesisHash))
	}
}

// genesisBlockBytes are the wire encoded bytes for the genesis block of the
// main network as of protocol version 70003.
var genesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0xff, 0x79, 0xaf, 0x16,
	0xa9, 0xff, 0xeb, 0x1b, 0x82, 0x6d, 0xe1, 0xea,
	0x7f, 0x24, 0x53, 0x9a, 0x2f, 0xe3, 0x70, 0x2f,
	0xe9, 0x87, 0x91, 0x2b, 0x09, 0x07, 0x2b, 0xc4,
	0x1d, 0xbc, 0x02, 0xb5, 0xd0, 0xfe, 0xdd, 0x52,
	0xf0, 0xff, 0x0f, 0x1e, 0xb3, 0x5a, 0x44, 0x0d,
	0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xff, 0xff, 0x30, 0x04, 0xff, 0xff, 0x00, 0x1d,
	0x01, 0x04, 0x28, 0x4a, 0x61, 0x6e, 0x75, 0x61,
	0x72, 0x79, 0x20, 0x32, 0x31, 0x73, 0x74, 0x20,
	0x32, 0x30, 0x31, 0x34, 0x20, 0x77, 0x61, 0x73,
	0x20, 0x73, 0x75, 0x63, 0x68, 0x20, 0x61, 0x20,
	0x6e, 0x69, 0x63, 0x65, 0x20, 0x64, 0x61, 0x79,
	0x2e, 0x2e, 0x2e, 0xff, 0xff, 0xff, 0xff, 0x01,
	0x00, 0x10, 0xa5, 0xd4, 0xe8, 0x00, 0x00, 0x00,
	0x43, 0x41, 0x04, 0x01, 0x84, 0x71, 0x0f, 0xa6,
	0x89, 0xad, 0x50, 0x23, 0x69, 0x0c, 0x80, 0xf3,
	0xa4, 0x9c, 0x8f, 0x13, 0xf8, 0xd4, 0x5b, 0x8c,
	0x85, 0x7f, 0xbc, 0xbc, 0x8b, 0xc4, 0xa8, 0xe4,
	0xd3, 0xeb, 0x4b, 0x10, 0xf4, 0xd4, 0x60, 0x4f,
	0xa0, 0x8d, 0xce, 0x60, 0x1a, 0xaf, 0x0f, 0x47,
	0x02, 0x16, 0xfe, 0x1b, 0x51, 0x85, 0x0b, 0x4a,
	0xcf, 0x21, 0xb1, 0x79, 0xc4, 0x50, 0x70, 0xac,
	0x7b, 0x03, 0xa9, 0xac, 0x00, 0x00, 0x00, 0x00,
}

// regTestGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the regression test network as of protocol version 60002.
var regTestGenesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x3b, 0xa3, 0xed, 0xfd, /* |....;...| */
	0x7a, 0x7b, 0x12, 0xb2, 0x7a, 0xc7, 0x2c, 0x3e, /* |z{..z.,>| */
	0x67, 0x76, 0x8f, 0x61, 0x7f, 0xc8, 0x1b, 0xc3, /* |gv.a....| */
	0x88, 0x8a, 0x51, 0x32, 0x3a, 0x9f, 0xb8, 0xaa, /* |..Q2:...| */
	0x4b, 0x1e, 0x5e, 0x4a, 0xda, 0xe5, 0x49, 0x4d, /* |K.^J)._I| */
	0xff, 0xff, 0x7f, 0x20, 0x02, 0x00, 0x00, 0x00, /* |......+|| */
	0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, /* |........| */
	0xff, 0xff, 0x4d, 0x04, 0xff, 0xff, 0x00, 0x1d, /* |..M.....| */
	0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, 0x54, /* |..EThe T| */
	0x69, 0x6d, 0x65, 0x73, 0x20, 0x30, 0x33, 0x2f, /* |imes 03/| */
	0x4a, 0x61, 0x6e, 0x2f, 0x32, 0x30, 0x30, 0x39, /* |Jan/2009| */
	0x20, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x6c, /* | Chancel| */
	0x6c, 0x6f, 0x72, 0x20, 0x6f, 0x6e, 0x20, 0x62, /* |lor on b| */
	0x72, 0x69, 0x6e, 0x6b, 0x20, 0x6f, 0x66, 0x20, /* |rink of | */
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x20, 0x62, /* |second b| */
	0x61, 0x69, 0x6c, 0x6f, 0x75, 0x74, 0x20, 0x66, /* |ailout f| */
	0x6f, 0x72, 0x20, 0x62, 0x61, 0x6e, 0x6b, 0x73, /* |or banks| */
	0xff, 0xff, 0xff, 0xff, 0x01, 0x00, 0xf2, 0x05, /* |........| */
	0x2a, 0x01, 0x00, 0x00, 0x00, 0x43, 0x41, 0x04, /* |*....CA.| */
	0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, 0x48, 0x27, /* |g....UH'| */
	0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, 0xb7, 0x10, /* |.g..q0..| */
	0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, 0x09, 0xa6, /* |\..(.9..| */
	0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, 0xde, 0xb6, /* |yb...a..| */
	0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, 0x38, 0xc4, /* |I..?L.8.| */
	0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, 0x12, 0xde, /* |.U......| */
	0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, 0x8d, 0x57, /* |\8M....W| */
	0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, 0x1d, 0x5f, /* |.Lp+k.._|*/
	0xac, 0x00, 0x00, 0x00, 0x00, /* |.....|    */
}

// testNet3GenesisBlockBytes are the wire encoded bytes for the genesis block of
// the test network (version 3) as of protocol version 60002.
var testNet3GenesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x3b, 0xa3, 0xed, 0xfd, /* |....;...| */
	0x7a, 0x7b, 0x12, 0xb2, 0x7a, 0xc7, 0x2c, 0x3e, /* |z{..z.,>| */
	0x67, 0x76, 0x8f, 0x61, 0x7f, 0xc8, 0x1b, 0xc3, /* |gv.a....| */
	0x88, 0x8a, 0x51, 0x32, 0x3a, 0x9f, 0xb8, 0xaa, /* |..Q2:...| */
	0x4b, 0x1e, 0x5e, 0x4a, 0xda, 0xe5, 0x49, 0x4d, /* |K.^J)._I| */
	0xff, 0xff, 0x00, 0x1d, 0x1a, 0xa4, 0xae, 0x18, /* |......+|| */
	0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, /* |........| */
	0xff, 0xff, 0x4d, 0x04, 0xff, 0xff, 0x00, 0x1d, /* |..M.....| */
	0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, 0x54, /* |..EThe T| */
	0x69, 0x6d, 0x65, 0x73, 0x20, 0x30, 0x33, 0x2f, /* |imes 03/| */
	0x4a, 0x61, 0x6e, 0x2f, 0x32, 0x30, 0x30, 0x39, /* |Jan/2009| */
	0x20, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x6c, /* | Chancel| */
	0x6c, 0x6f, 0x72, 0x20, 0x6f, 0x6e, 0x20, 0x62, /* |lor on b| */
	0x72, 0x69, 0x6e, 0x6b, 0x20, 0x6f, 0x66, 0x20, /* |rink of | */
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x20, 0x62, /* |second b| */
	0x61, 0x69, 0x6c, 0x6f, 0x75, 0x74, 0x20, 0x66, /* |ailout f| */
	0x6f, 0x72, 0x20, 0x62, 0x61, 0x6e, 0x6b, 0x73, /* |or banks| */
	0xff, 0xff, 0xff, 0xff, 0x01, 0x00, 0xf2, 0x05, /* |........| */
	0x2a, 0x01, 0x00, 0x00, 0x00, 0x43, 0x41, 0x04, /* |*....CA.| */
	0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, 0x48, 0x27, /* |g....UH'| */
	0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, 0xb7, 0x10, /* |.g..q0..| */
	0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, 0x09, 0xa6, /* |\..(.9..| */
	0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, 0xde, 0xb6, /* |yb...a..| */
	0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, 0x38, 0xc4, /* |I..?L.8.| */
	0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, 0x12, 0xde, /* |.U......| */
	0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, 0x8d, 0x57, /* |\8M....W| */
	0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, 0x1d, 0x5f, /* |.Lp+k.._|*/
	0xac, 0x00, 0x00, 0x00, 0x00, /* |.....|    */
}

// simNetGenesisBlockBytes are the wire encoded bytes for the genesis block of
// the simulation test network as of protocol version 70002.
var simNetGenesisBlockBytes = []byte{
	0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x3b, 0xa3, 0xed, 0xfd, /* |....;...| */
	0x7a, 0x7b, 0x12, 0xb2, 0x7a, 0xc7, 0x2c, 0x3e, /* |z{..z.,>| */
	0x67, 0x76, 0x8f, 0x61, 0x7f, 0xc8, 0x1b, 0xc3, /* |gv.a....| */
	0x88, 0x8a, 0x51, 0x32, 0x3a, 0x9f, 0xb8, 0xaa, /* |..Q2:...| */
	0x4b, 0x1e, 0x5e, 0x4a, 0x45, 0x06, 0x86, 0x53, /* |K.^J)._I| */
	0xff, 0xff, 0x7f, 0x20, 0x02, 0x00, 0x00, 0x00, /* |......+|| */
	0x01, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* |........| */
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, /* |........| */
	0xff, 0xff, 0x4d, 0x04, 0xff, 0xff, 0x00, 0x1d, /* |..M.....| */
	0x01, 0x04, 0x45, 0x54, 0x68, 0x65, 0x20, 0x54, /* |..EThe T| */
	0x69, 0x6d, 0x65, 0x73, 0x20, 0x30, 0x33, 0x2f, /* |imes 03/| */
	0x4a, 0x61, 0x6e, 0x2f, 0x32, 0x30, 0x30, 0x39, /* |Jan/2009| */
	0x20, 0x43, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x6c, /* | Chancel| */
	0x6c, 0x6f, 0x72, 0x20, 0x6f, 0x6e, 0x20, 0x62, /* |lor on b| */
	0x72, 0x69, 0x6e, 0x6b, 0x20, 0x6f, 0x66, 0x20, /* |rink of | */
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x20, 0x62, /* |second b| */
	0x61, 0x69, 0x6c, 0x6f, 0x75, 0x74, 0x20, 0x66, /* |ailout f| */
	0x6f, 0x72, 0x20, 0x62, 0x61, 0x6e, 0x6b, 0x73, /* |or banks| */
	0xff, 0xff, 0xff, 0xff, 0x01, 0x00, 0xf2, 0x05, /* |........| */
	0x2a, 0x01, 0x00, 0x00, 0x00, 0x43, 0x41, 0x04, /* |*....CA.| */
	0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, 0x48, 0x27, /* |g....UH'| */
	0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, 0xb7, 0x10, /* |.g..q0..| */
	0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, 0x09, 0xa6, /* |\..(.9..| */
	0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, 0xde, 0xb6, /* |yb...a..| */
	0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, 0x38, 0xc4, /* |I..?L.8.| */
	0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, 0x12, 0xde, /* |.U......| */
	0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, 0x8d, 0x57, /* |\8M....W| */
	0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, 0x1d, 0x5f, /* |.Lp+k.._|*/
	0xac, 0x00, 0x00, 0x00, 0x00, /* |.....|    */
}
