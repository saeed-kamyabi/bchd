// Copyright (c) 2017 The bchsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package blockchain

import (
        "math"
	"github.com/bchsuite/bchd/wire"
	"github.com/bchsuite/bchutil"
)


const (
        // DefaultMaxBlockSize is the current concensus for large blocks
        // Can be overridden by ExcessiveBlockSize config parameter 
        DefaultMaxBlockSize = 8000000

        // LegacyMaxBlockSize can be used for any pre-fork checks
        LegacyMaxBlockSize = 1000000

        // MaxBlockSigOps is used to derrive the maximum number of sigops allowed
        // for a block.  
        MaxBlockSigOps = 20000
)


func getMinTransactionSize() uint32 {
      	var msgTx wire.MsgTx
        tx := bchutil.NewTx(&msgTx)
        serializedTxSize := tx.MsgTx().SerializeSize()
        return uint32(serializedTxSize)
}

func GetMaxSigOpsPerBlock(MaxBlockSize uint32) uint32 {
        // BUIP040 MaxSigOps calculation
        return MaxBlockSigOps * uint32(math.Ceil((math.Max(float64(MaxBlockSize), float64(LegacyMaxBlockSize)) / float64(LegacyMaxBlockSize))))
}
