// Copyright (c) 2014 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rddnet

import (
	"errors"
	"math/big"

	"github.com/reddcoin-project/rddwire"
)

// These variables are the chain proof-of-work limit parameters for each default
// network.
var (
	// bigOne is 1 represented as a big.Int.  It is defined here to avoid
	// the overhead of creating it multiple times.
	bigOne = big.NewInt(1)

	// mainPowLimit is the highest proof of work value a Bitcoin block can
	// have for the main network.  It is the value 2^236 - 1.
	mainPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 236), bigOne)

	// regressionPowLimit is the highest proof of work value a Bitcoin block
	// can have for the regression test network.  It is the value 2^255 - 1.
	regressionPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 255), bigOne)

	// testNet3PowLimit is the highest proof of work value a Bitcoin block
	// can have for the test network (version 3).  It is the value
	// 2^224 - 1.
	testNet3PowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 224), bigOne)

	// simNetPowLimit is the highest proof of work value a Bitcoin block
	// can have for the simulation test network.  It is the value 2^255 - 1.
	simNetPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 255), bigOne)
)

// Checkpoint identifies a known good point in the block chain.  Using
// checkpoints allows a few optimizations for old blocks during initial download
// and also prevents forks from old blocks.
//
// Each checkpoint is selected based upon several factors.  See the
// documentation for btcchain.IsCheckpointCandidate for details on the selection
// criteria.
type Checkpoint struct {
	Height int64
	Hash   *rddwire.ShaHash
}

// Params defines a Bitcoin network by its parameters.  These parameters may be
// used by Bitcoin applications to differentiate networks as well as addresses
// and keys for one network from those intended for use on another network.
type Params struct {
	Name        string
	Net         rddwire.BitcoinNet
	DefaultPort string

	// Chain parameters
	GenesisBlock           *rddwire.MsgBlock
	GenesisHash            *rddwire.ShaHash
	PowLimit               *big.Int
	PowLimitBits           uint32
	SubsidyHalvingInterval int32
	ResetMinDifficulty     bool

	// Checkpoints ordered from oldest to newest.
	Checkpoints []Checkpoint

	// Reject version 1 blocks once a majority of the network has upgraded.
	// This is part of BIP0034.
	BlockV1RejectNumRequired uint64
	BlockV1RejectNumToCheck  uint64

	// Ensure coinbase starts with serialized block heights for version 2
	// blocks or newer once a majority of the network has upgraded.
	CoinbaseBlockHeightNumRequired uint64
	CoinbaseBlockHeightNumToCheck  uint64

	// Mempool parameters
	RelayNonStdTxs bool

	// Address encoding magics
	PubKeyHashAddrID byte // First byte of a P2PKH address
	ScriptHashAddrID byte // First byte of a P2SH address
	PrivateKeyID     byte // First byte of a WIF private key

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID [4]byte
	HDPublicKeyID  [4]byte

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType uint32
}

// MainNetParams defines the network parameters for the main Bitcoin network.
var MainNetParams = Params{
	Name:        "mainnet",
	Net:         rddwire.MainNet,
	DefaultPort: "45444",

	// Chain parameters
	GenesisBlock:           &genesisBlock,
	GenesisHash:            &genesisHash,
	PowLimit:               mainPowLimit,
	PowLimitBits:           0x1e0fffff,
	ResetMinDifficulty:     false,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []Checkpoint{
		{10,     newShaHashFromStr("a198c38a77555a9fbff0b147bf7ce0660416d6abdaa86adaa3a9be97092592ed")},
		{1000,   newShaHashFromStr("9d849e078deac30d58372db898318186cf5073a7f0b109b4776393b21b7b4e5a")},
		{2000,   newShaHashFromStr("4674c50137c1d9bf47d96dee5e8c38c41895d494a0bb71e243d1c8a6c805e1f7")},
		{3000,   newShaHashFromStr("0deff246b8dfc102ccdbc3a306649be82c441e1da0fba8ca1075cf6b5d7f3c01")},
		{4000,   newShaHashFromStr("ad880a4c23d511f04311e98ee77f5163e54cd92f80433e9f3bd6bc2261d50592")},
		{5000,   newShaHashFromStr("3f673ef045f4a7d71fb841ce96ed190b20569182bd3dfe628527ec3a7934d08f")},
		{6000,   newShaHashFromStr("1222056e58dce70c0a6638e07415bd6190fa5ccdd6d5e7f6af68abb30ebd4eb8")},
		{6150,   newShaHashFromStr("e221b12cf8b0c264697e9bb3c9c9f0f7ada5f2736e054cbd53b94852908cdbd3")},
		{10000,  newShaHashFromStr("35d5f9cbd94c15771d5ebebf55ea4bfc649c473c0a868fe4d1832f5b45bd5d0c")},
		{15000,  newShaHashFromStr("87a8c4289e661720095f2ab6a077151bc84b9615a53c5e7207ba1c20418bd178")},
		{20000,  newShaHashFromStr("6a86a4cbbcea694027591ba416ae3831b4f3f9aa3cc6a0a1b5627f920dd765bb")},
		{44878,  newShaHashFromStr("d81a3724a81b78e762821d16bfe23565576905685b2c072ea9a3fa7d36dbad8b")},
		{45189,  newShaHashFromStr("d10b5da162b922d3cf09c44ea3d533a203c9ab1c442015d7e72f21062d20aeb4")},
		{45239,  newShaHashFromStr("e14dba7c7d1ed1a7566e23b0ca0989e3e26099b7beaa673d324001d1291223f7")},
		{114834, newShaHashFromStr("dc5c776ca006c6d40e48c90aeeb43bf6493de589e28f779b486e64aa3403344a")},
		{184000, newShaHashFromStr("e22e6b027cd49cd9aa2ba6df0e0c264c2eed5656b1fa39052c8235d52f2b04d6")},
		{244999, newShaHashFromStr("0b7bb56edfae2f2f1e71ac39daab16614fccf1a1e02c58d4169521d76d880b42")},		
	},

	// Reject version 1 blocks once a majority of the network has upgraded.
	// 95% (950 / 1000)
	// This is part of BIP0034.
	BlockV1RejectNumRequired: 950,
	BlockV1RejectNumToCheck:  1000,

	// Ensure coinbase starts with serialized block heights for version 2
	// blocks or newer once a majority of the network has upgraded.
	// 75% (750 / 1000)
	// This is part of BIP0034.
	CoinbaseBlockHeightNumRequired: 750,
	CoinbaseBlockHeightNumToCheck:  1000,

	// Mempool parameters
	RelayNonStdTxs: false,

	// Address encoding magics
	PubKeyHashAddrID: 0x3d, // starts with R
	ScriptHashAddrID: 0x05, // starts with 3
	PrivateKeyID:     0xbd, // starts with V (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x88, 0xad, 0xe4}, // starts with xprv
	HDPublicKeyID:  [4]byte{0x04, 0x88, 0xb2, 0x1e}, // starts with xpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 4,
}

// RegressionNetParams defines the network parameters for the regression test
// Bitcoin network.  Not to be confused with the test Bitcoin network (version
// 3), this network is sometimes simply called "testnet".
var RegressionNetParams = Params{
	Name:        "regtest",
	Net:         rddwire.TestNet,
	DefaultPort: "18444",

	// Chain parameters
	GenesisBlock:           &regTestGenesisBlock,
	GenesisHash:            &regTestGenesisHash,
	PowLimit:               regressionPowLimit,
	PowLimitBits:           0x207fffff,
	ResetMinDifficulty:     true,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: nil,

	// Reject version 1 blocks once a majority of the network has upgraded.
	// 75% (75 / 100)
	// This is part of BIP0034.
	BlockV1RejectNumRequired: 75,
	BlockV1RejectNumToCheck:  100,

	// Ensure coinbase starts with serialized block heights for version 2
	// blocks or newer once a majority of the network has upgraded.
	// 51% (51 / 100)
	// This is part of BIP0034.
	CoinbaseBlockHeightNumRequired: 51,
	CoinbaseBlockHeightNumToCheck:  100,

	// Mempool parameters
	RelayNonStdTxs: true,

	// Address encoding magics
	PubKeyHashAddrID: 0x6f, // starts with m or n
	ScriptHashAddrID: 0xc4, // starts with 2
	PrivateKeyID:     0xef, // starts with 9 (uncompressed) or c (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 1,
}

// TestNet3Params defines the network parameters for the test Bitcoin network
// (version 3).  Not to be confused with the regression test network, this
// network is sometimes simply called "testnet".
var TestNet3Params = Params{
	Name:        "testnet3",
	Net:         rddwire.TestNet3,
	DefaultPort: "18333",

	// Chain parameters
	GenesisBlock:           &testNet3GenesisBlock,
	GenesisHash:            &testNet3GenesisHash,
	PowLimit:               testNet3PowLimit,
	PowLimitBits:           0x1d00ffff,
	ResetMinDifficulty:     true,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: nil,

	// Reject version 1 blocks once a majority of the network has upgraded.
	// 75% (75 / 100)
	// This is part of BIP0034.
	BlockV1RejectNumRequired: 75,
	BlockV1RejectNumToCheck:  100,

	// Ensure coinbase starts with serialized block heights for version 2
	// blocks or newer once a majority of the network has upgraded.
	// 51% (51 / 100)
	// This is part of BIP0034.
	CoinbaseBlockHeightNumRequired: 51,
	CoinbaseBlockHeightNumToCheck:  100,

	// Mempool parameters
	RelayNonStdTxs: true,

	// Address encoding magics
	PubKeyHashAddrID: 0x6f, // starts with m or n
	ScriptHashAddrID: 0xc4, // starts with 2
	PrivateKeyID:     0xef, // starts with 9 (uncompressed) or c (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 1,
}

// SimNetParams defines the network parameters for the simulation test Bitcoin
// network.  This network is similar to the normal test network except it is
// intended for private use within a group of individuals doing simulation
// testing.  The functionality is intended to differ in that the only nodes
// which are specifically specified are used to create the network rather than
// following normal discovery rules.  This is important as otherwise it would
// just turn into another public testnet.
var SimNetParams = Params{
	Name:        "simnet",
	Net:         rddwire.SimNet,
	DefaultPort: "18555",

	// Chain parameters
	GenesisBlock:           &simNetGenesisBlock,
	GenesisHash:            &simNetGenesisHash,
	PowLimit:               simNetPowLimit,
	PowLimitBits:           0x207fffff,
	ResetMinDifficulty:     true,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: nil,

	// Reject version 1 blocks once a majority of the network has upgraded.
	// 75% (75 / 100)
	BlockV1RejectNumRequired: 75,
	BlockV1RejectNumToCheck:  100,

	// Ensure coinbase starts with serialized block heights for version 2
	// blocks or newer once a majority of the network has upgraded.
	// 51% (51 / 100)
	CoinbaseBlockHeightNumRequired: 51,
	CoinbaseBlockHeightNumToCheck:  100,

	// Mempool parameters
	RelayNonStdTxs: true,

	// Address encoding magics
	PubKeyHashAddrID: 0x3f, // starts with S
	ScriptHashAddrID: 0x7b, // starts with s
	PrivateKeyID:     0x64, // starts with 4 (uncompressed) or F (compressed)

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x20, 0xb9, 0x00}, // starts with sprv
	HDPublicKeyID:  [4]byte{0x04, 0x20, 0xbd, 0x3a}, // starts with spub

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 115, // ASCII for s
}

var (
	// ErrDuplicateNet describes an error where the parameters for a Bitcoin
	// network could not be set due to the network already being a standard
	// network or previously-registered into this package.
	ErrDuplicateNet = errors.New("duplicate Bitcoin network")

	// ErrUnknownHDKeyID describes an error where the provided id which
	// is intended to identify the network for a hierarchical deterministic
	// private extended key is not registered.
	ErrUnknownHDKeyID = errors.New("unknown hd private extended key bytes")
)

var (
	registeredNets = map[rddwire.BitcoinNet]struct{}{
		MainNetParams.Net:       struct{}{},
		TestNet3Params.Net:      struct{}{},
		RegressionNetParams.Net: struct{}{},
		SimNetParams.Net:        struct{}{},
	}

	pubKeyHashAddrIDs = map[byte]struct{}{
		MainNetParams.PubKeyHashAddrID:  struct{}{},
		TestNet3Params.PubKeyHashAddrID: struct{}{}, // shared with regtest
		SimNetParams.PubKeyHashAddrID:   struct{}{},
	}

	scriptHashAddrIDs = map[byte]struct{}{
		MainNetParams.ScriptHashAddrID:  struct{}{},
		TestNet3Params.ScriptHashAddrID: struct{}{}, // shared with regtest
		SimNetParams.ScriptHashAddrID:   struct{}{},
	}

	// Testnet is shared with regtest.
	hdPrivToPubKeyIDs = map[[4]byte][]byte{
		MainNetParams.HDPrivateKeyID:  MainNetParams.HDPublicKeyID[:],
		TestNet3Params.HDPrivateKeyID: TestNet3Params.HDPublicKeyID[:],
		SimNetParams.HDPrivateKeyID:   SimNetParams.HDPublicKeyID[:],
	}
)

// Register registers the network parameters for a Bitcoin network.  This may
// error with ErrDuplicateNet if the network is already registered (either
// due to a previous Register call, or the network being one of the default
// networks).
//
// Network parameters should be registered into this package by a main package
// as early as possible.  Then, library packages may lookup networks or network
// parameters based on inputs and work regardless of the network being standard
// or not.
func Register(params *Params) error {
	if _, ok := registeredNets[params.Net]; ok {
		return ErrDuplicateNet
	}
	registeredNets[params.Net] = struct{}{}
	pubKeyHashAddrIDs[params.PubKeyHashAddrID] = struct{}{}
	scriptHashAddrIDs[params.ScriptHashAddrID] = struct{}{}
	hdPrivToPubKeyIDs[params.HDPrivateKeyID] = params.HDPublicKeyID[:]
	return nil
}

// IsPubKeyHashAddrID returns whether the id is an identifier known to prefix a
// pay-to-pubkey-hash address on any default or registered network.  This is
// used when decoding an address string into a specific address type.  It is up
// to the caller to check both this and IsScriptHashAddrID and decide whether an
// address is a pubkey hash address, script hash address, neither, or
// undeterminable (if both return true).
func IsPubKeyHashAddrID(id byte) bool {
	_, ok := pubKeyHashAddrIDs[id]
	return ok
}

// IsScriptHashAddrID returns whether the id is an identifier known to prefix a
// pay-to-script-hash address on any default or registered network.  This is
// used when decoding an address string into a specific address type.  It is up
// to the caller to check both this and IsPubKeyHashAddrID and decide whether an
// address is a pubkey hash address, script hash address, neither, or
// undeterminable (if both return true).
func IsScriptHashAddrID(id byte) bool {
	_, ok := scriptHashAddrIDs[id]
	return ok
}

// HDPrivateKeyToPublicKeyID accepts a private hierarchical deterministic
// extended key id and returns the associated public key id.  When the provided
// id is not registered, the ErrUnknownHDKeyID error will be returned.
func HDPrivateKeyToPublicKeyID(id []byte) ([]byte, error) {
	if len(id) != 4 {
		return nil, ErrUnknownHDKeyID
	}

	var key [4]byte
	copy(key[:], id)
	pubBytes, ok := hdPrivToPubKeyIDs[key]
	if !ok {
		return nil, ErrUnknownHDKeyID
	}

	return pubBytes, nil
}

// newShaHashFromStr converts the passed big-endian hex string into a
// rddwire.ShaHash.  It only differs from the one available in rddwire in that
// it panics on an error since it will only (and must only) be called with
// hard-coded, and therefore known good, hashes.
func newShaHashFromStr(hexStr string) *rddwire.ShaHash {
	sha, err := rddwire.NewShaHashFromStr(hexStr)
	if err != nil {
		// Ordinarily I don't like panics in library code since it
		// can take applications down without them having a chance to
		// recover which is extremely annoying, however an exception is
		// being made in this case because the only way this can panic
		// is if there is an error in the hard-coded hashes.  Thus it
		// will only ever potentially panic on init and therefore is
		// 100% predictable.
		panic(err)
	}
	return sha
}
