// Copyright (c) 2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rddnet

import (
	"time"

	"github.com/reddcoin-project/rddwire"
)

// genesisCoinbaseTxBtc is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTxBtc = rddwire.MsgTx{
	Version: 1,
	TxIn: []*rddwire.TxIn{
		{
			PreviousOutPoint: rddwire.OutPoint{
				Hash:  rddwire.ShaHash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x45, /* |.......E| */
				0x54, 0x68, 0x65, 0x20, 0x54, 0x69, 0x6d, 0x65, /* |The Time| */
				0x73, 0x20, 0x30, 0x33, 0x2f, 0x4a, 0x61, 0x6e, /* |s 03/Jan| */
				0x2f, 0x32, 0x30, 0x30, 0x39, 0x20, 0x43, 0x68, /* |/2009 Ch| */
				0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x6f, 0x72, /* |ancellor| */
				0x20, 0x6f, 0x6e, 0x20, 0x62, 0x72, 0x69, 0x6e, /* | on brin| */
				0x6b, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x65, 0x63, /* |k of sec|*/
				0x6f, 0x6e, 0x64, 0x20, 0x62, 0x61, 0x69, 0x6c, /* |ond bail| */
				0x6f, 0x75, 0x74, 0x20, 0x66, 0x6f, 0x72, 0x20, /* |out for |*/
				0x62, 0x61, 0x6e, 0x6b, 0x73, /* |banks| */
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*rddwire.TxOut{
		{
			Value: 0x12a05f200,
			PkScript: []byte{
				0x41, 0x04, 0x67, 0x8a, 0xfd, 0xb0, 0xfe, 0x55, /* |A.g....U| */
				0x48, 0x27, 0x19, 0x67, 0xf1, 0xa6, 0x71, 0x30, /* |H'.g..q0| */
				0xb7, 0x10, 0x5c, 0xd6, 0xa8, 0x28, 0xe0, 0x39, /* |..\..(.9| */
				0x09, 0xa6, 0x79, 0x62, 0xe0, 0xea, 0x1f, 0x61, /* |..yb...a| */
				0xde, 0xb6, 0x49, 0xf6, 0xbc, 0x3f, 0x4c, 0xef, /* |..I..?L.| */
				0x38, 0xc4, 0xf3, 0x55, 0x04, 0xe5, 0x1e, 0xc1, /* |8..U....| */
				0x12, 0xde, 0x5c, 0x38, 0x4d, 0xf7, 0xba, 0x0b, /* |..\8M...| */
				0x8d, 0x57, 0x8a, 0x4c, 0x70, 0x2b, 0x6b, 0xf1, /* |.W.Lp+k.| */
				0x1d, 0x5f, 0xac, /* |._.| */
			},
		},
	},
	LockTime: 0,
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = rddwire.MsgTx{
	Version: 1,
	TxIn: []*rddwire.TxIn{
		{
			PreviousOutPoint: rddwire.OutPoint{
				Hash:  rddwire.ShaHash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x28, /* |.......(| */
				0x4a, 0x61, 0x6e, 0x75, 0x61, 0x72, 0x79, 0x20, /* |January | */
				0x32, 0x31, 0x73, 0x74, 0x20, 0x32, 0x30, 0x31, /* |21st 201| */
				0x34, 0x20, 0x77, 0x61, 0x73, 0x20, 0x73, 0x75, /* |4 was su| */
				0x63, 0x68, 0x20, 0x61, 0x20, 0x6e, 0x69, 0x63, /* |ch a nic| */
				0x65, 0x20, 0x64, 0x61, 0x79, 0x2e, 0x2e, 0x2e, /* |e day...| */
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*rddwire.TxOut{
		{
			Value: 0xe8d4a51000, // 1,000,000,000,000
			PkScript: []byte{
				0x41, 0x04, 0x01, 0x84, 0x71, 0x0f, 0xa6, 0x89,
				0xad, 0x50, 0x23, 0x69, 0x0c, 0x80, 0xf3, 0xa4,
				0x9c, 0x8f, 0x13, 0xf8, 0xd4, 0x5b, 0x8c, 0x85,
				0x7f, 0xbc, 0xbc, 0x8b, 0xc4, 0xa8, 0xe4, 0xd3,
				0xeb, 0x4b, 0x10, 0xf4, 0xd4, 0x60, 0x4f, 0xa0,
				0x8d, 0xce, 0x60, 0x1a, 0xaf, 0x0f, 0x47, 0x02,
				0x16, 0xfe, 0x1b, 0x51, 0x85, 0x0b, 0x4a, 0xcf,
				0x21, 0xb1, 0x79, 0xc4, 0x50, 0x70, 0xac, 0x7b,
				0x03, 0xa9, 0xac,
			},
		},
	},
	LockTime: 0,
	Timestamp: time.Unix(0, 0),
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
var genesisHash = rddwire.ShaHash([rddwire.HashSize]byte{ // Make go vet happy.
	0xcc, 0xde, 0xc1, 0x74, 0xeb, 0xd4, 0xfa, 0x10,
	0x31, 0x4b, 0x3b, 0x9e, 0xf9, 0xcb, 0x8a, 0xdc,
	0xf9, 0xaa, 0x87, 0xe5, 0x7e, 0xc6, 0xad, 0x0d,
	0x0e, 0x3c, 0x3c, 0x5a, 0xd9, 0xe0, 0x68, 0xb8,
})

// genesisMerkleRootBtc is the hash of the first transaction in the genesis block
// for the bitcoin main network.
var genesisMerkleRootBtc = rddwire.ShaHash([rddwire.HashSize]byte{ // Make go vet happy.
	0x3b, 0xa3, 0xed, 0xfd, 0x7a, 0x7b, 0x12, 0xb2,
	0x7a, 0xc7, 0x2c, 0x3e, 0x67, 0x76, 0x8f, 0x61,
	0x7f, 0xc8, 0x1b, 0xc3, 0x88, 0x8a, 0x51, 0x32,
	0x3a, 0x9f, 0xb8, 0xaa, 0x4b, 0x1e, 0x5e, 0x4a,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
var genesisMerkleRoot = rddwire.ShaHash([rddwire.HashSize]byte{ // Make go vet happy.
	0xff, 0x79, 0xaf, 0x16, 0xa9, 0xff, 0xeb, 0x1b,
	0x82, 0x6d, 0xe1, 0xea, 0x7f, 0x24, 0x53, 0x9a,
	0x2f, 0xe3, 0x70, 0x2f, 0xe9, 0x87, 0x91, 0x2b,
	0x09, 0x07, 0x2b, 0xc4, 0x1d, 0xbc, 0x02, 0xb5,
})

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = rddwire.MsgBlock{
	Header: rddwire.BlockHeader{
		Version:    1,
		PrevBlock:  rddwire.ShaHash{},        // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot,        // b502bc1dc42b07092b9187e92f70e32f9a53247feae16d821bebffa916af79ff
		Timestamp:  time.Unix(1390280400, 0), // 2014-01-21 05:00:00 +0000 UTC
		Bits:       0x1e0ffff0,               // 504365040 [00000ffff0000000000000000000000000000000000000000000000000000000]
		Nonce:      0x0d445ab3,               // 222583475
	},
	Transactions: []*rddwire.MsgTx{&genesisCoinbaseTx},
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
var regTestGenesisHash = rddwire.ShaHash([rddwire.HashSize]byte{ // Make go vet happy.
	0x06, 0x22, 0x6e, 0x46, 0x11, 0x1a, 0x0b, 0x59,
	0xca, 0xaf, 0x12, 0x60, 0x43, 0xeb, 0x5b, 0xbf,
	0x28, 0xc3, 0x4f, 0x3a, 0x5e, 0x33, 0x2a, 0x1f,
	0xc7, 0xb2, 0xb7, 0x3c, 0xf1, 0x88, 0x91, 0x0f,
})

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = genesisMerkleRootBtc

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = rddwire.MsgBlock{
	Header: rddwire.BlockHeader{
		Version:    1,
		PrevBlock:  rddwire.ShaHash{},        // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: regTestGenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1296688602, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*rddwire.MsgTx{&genesisCoinbaseTxBtc},
}

// testNet3GenesisHash is the hash of the first block in the block chain for the
// test network (version 3).
var testNet3GenesisHash = rddwire.ShaHash([rddwire.HashSize]byte{ // Make go vet happy.
	0x43, 0x49, 0x7f, 0xd7, 0xf8, 0x26, 0x95, 0x71,
	0x08, 0xf4, 0xa3, 0x0f, 0xd9, 0xce, 0xc3, 0xae,
	0xba, 0x79, 0x97, 0x20, 0x84, 0xe9, 0x0e, 0xad,
	0x01, 0xea, 0x33, 0x09, 0x00, 0x00, 0x00, 0x00,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
var testNet3GenesisMerkleRoot = genesisMerkleRootBtc

// testNet3GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 3).
var testNet3GenesisBlock = rddwire.MsgBlock{
	Header: rddwire.BlockHeader{
		Version:    1,
		PrevBlock:  rddwire.ShaHash{},         // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNet3GenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1296688602, 0),  // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x1d00ffff,                // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0x18aea41a,                // 414098458
	},
	Transactions: []*rddwire.MsgTx{&genesisCoinbaseTxBtc},
}

// simNetGenesisHash is the hash of the first block in the block chain for the
// simulation test network.
var simNetGenesisHash = rddwire.ShaHash([rddwire.HashSize]byte{ // Make go vet happy.
	0xf6, 0x7a, 0xd7, 0x69, 0x5d, 0x9b, 0x66, 0x2a,
	0x72, 0xff, 0x3d, 0x8e, 0xdb, 0xbb, 0x2d, 0xe0,
	0xbf, 0xa6, 0x7b, 0x13, 0x97, 0x4b, 0xb9, 0x91,
	0x0d, 0x11, 0x6d, 0x5c, 0xbd, 0x86, 0x3e, 0x68,
})

// simNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network.  It is the same as the merkle root for
// the main network.
var simNetGenesisMerkleRoot = genesisMerkleRootBtc

// simNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var simNetGenesisBlock = rddwire.MsgBlock{
	Header: rddwire.BlockHeader{
		Version:    1,
		PrevBlock:  rddwire.ShaHash{},        // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: simNetGenesisMerkleRoot,  // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1401292357, 0), // 2014-05-28 15:52:37 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*rddwire.MsgTx{&genesisCoinbaseTxBtc},
}
