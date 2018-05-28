package sebak

import (
	"encoding/json"
	"fmt"

	"github.com/spikeekips/sebak/lib/common"
	"github.com/spikeekips/sebak/lib/error"
	"github.com/spikeekips/sebak/lib/storage"
)

/*
BlockTransaction is `Transaction` data for block. the storage should support,
 * find by `Hash`

 * //get list by `Checkpoint` and created order
 * get list by `Source` and created order
 * get list by `Confirmed` order
*/

const (
	BlockTransactionPrefixHash string = "bt-hash-" // bt-hash-<BlockTransaction.Hash>
	//BlockTransactionPrefixCheckpoint string = "bt-checkpoint-" // bt-hash-<BlockTransaction.Checkpoint>
	BlockTransactionPrefixSource    string = "bt-source-"    // bt-hash-<BlockTransaction.Source>
	BlockTransactionPrefixConfirmed string = "bt-confirmed-" // bt-hash-<BlockTransaction.Confirmed>
)

// TODO(BlockTransaction): support counting

type BlockTransaction struct {
	Hash string

	Checkpoint string
	Signature  string
	Source     string
	Fee        Amount
	Operations []string
	Amount     Amount

	Confirmed string
	Created   string
	Message   []byte

	transaction Transaction
}

func NewBlockTransactionFromTransaction(tx Transaction, message []byte) BlockTransaction {
	var opHashes []string
	for _, op := range tx.B.Operations {
		opHashes = append(opHashes, op.MakeHashString())
	}

	return BlockTransaction{
		Hash:       tx.H.Hash,
		Checkpoint: tx.B.Checkpoint,
		Signature:  tx.H.Signature,
		Source:     tx.B.Source,
		Fee:        tx.B.Fee,
		Operations: opHashes,
		Amount:     tx.TotalAmount(true),

		Created: tx.H.Created,
		Message: message,

		transaction: tx,
	}
}

/*
func (bt BlockTransaction) NewBlockTransactionKeyCheckpoint() string {
	return fmt.Sprintf(
		"%s%s",
		GetBlockTransactionKeyPrefixCheckpoint(bt.Checkpoint),
		sebakcommon.GetUniqueIDFromUUID(),
	)
}
*/

func (bt BlockTransaction) NewBlockTransactionKeySource() string {
	return fmt.Sprintf(
		"%s%s",
		GetBlockTransactionKeyPrefixSource(bt.Source),
		sebakcommon.GetUniqueIDFromUUID(),
	)
}

func (bt BlockTransaction) NewBlockTransactionKeyConfirmed() string {
	return fmt.Sprintf(
		"%s%s",
		GetBlockTransactionKeyPrefixConfirmed(bt.Confirmed),
		sebakcommon.GetUniqueIDFromUUID(),
	)
}

func (bt BlockTransaction) Save(st *sebakstorage.LevelDBBackend) (err error) {
	key := GetBlockTransactionKey(bt.Hash)

	var exists bool
	exists, err = st.Has(key)
	if err != nil {
		return
	} else if exists {
		return sebakerror.ErrorBlockAlreadyExists
	}

	bt.Confirmed = sebakcommon.NowISO8601()
	if err = st.New(GetBlockTransactionKey(bt.Hash), bt); err != nil {
		return
	}
	/*
		if err = st.New(bt.NewBlockTransactionKeyCheckpoint(), bt.Hash); err != nil {
			return
		}
	*/
	if err = st.New(bt.NewBlockTransactionKeySource(), bt.Hash); err != nil {
		return
	}
	if err = st.New(bt.NewBlockTransactionKeyConfirmed(), bt.Hash); err != nil {
		return
	}

	for _, op := range bt.transaction.B.Operations {
		bo := NewBlockOperationFromOperation(op, bt.transaction)
		if err = bo.Save(st); err != nil {
			return
		}
	}

	return nil
}

func (bt BlockTransaction) Serialize() (encoded []byte, err error) {
	encoded, err = sebakcommon.EncodeJSONValue(bt)
	return
}

func (bt BlockTransaction) Transaction() Transaction {
	return bt.transaction
}

/*
func GetBlockTransactionKeyPrefixCheckpoint(checkpoint string) string {
	return fmt.Sprintf("%s%s-", BlockTransactionPrefixCheckpoint, checkpoint)
}
*/

func GetBlockTransactionKeyPrefixSource(source string) string {
	return fmt.Sprintf("%s%s-", BlockTransactionPrefixSource, source)
}

func GetBlockTransactionKeyPrefixConfirmed(confirmed string) string {
	return fmt.Sprintf("%s%s-", BlockTransactionPrefixConfirmed, confirmed)
}

func GetBlockTransactionKey(hash string) string {
	return fmt.Sprintf("%s%s", BlockTransactionPrefixHash, hash)
}

func GetBlockTransaction(st *sebakstorage.LevelDBBackend, hash string) (bt BlockTransaction, err error) {
	if err = st.Get(GetBlockTransactionKey(hash), &bt); err != nil {
		return
	}

	return
}

func ExistBlockTransaction(st *sebakstorage.LevelDBBackend, hash string) (bool, error) {
	return st.Has(GetBlockTransactionKey(hash))
}

func LoadBlockTransactionsInsideIterator(
	st *sebakstorage.LevelDBBackend,
	iterFunc func() (sebakstorage.IterItem, bool),
	closeFunc func(),
) (
	func() (BlockTransaction, bool),
	func(),
) {

	return (func() (BlockTransaction, bool) {
			item, hasNext := iterFunc()
			if !hasNext {
				return BlockTransaction{}, false
			}

			var hash string
			json.Unmarshal(item.Value, &hash)

			bt, err := GetBlockTransaction(st, hash)
			if err != nil {
				return BlockTransaction{}, false
			}

			return bt, hasNext
		}), (func() {
			closeFunc()
		})
}

/*
// GetBlockTransactionsByCheckpoint is maybe useless, because the only one
// transaction could be possible from a transaction.
func GetBlockTransactionsByCheckpoint(st *sebakstorage.LevelDBBackend, checkpoint string, reverse bool) (
	func() (BlockTransaction, bool),
	func(),
) {
	iterFunc, closeFunc := st.GetIterator(GetBlockTransactionKeyPrefixCheckpoint(checkpoint), reverse)

	return LoadBlockTransactionsInsideIterator(st, iterFunc, closeFunc)
}
*/

func GetBlockTransactionsBySource(st *sebakstorage.LevelDBBackend, source string, reverse bool) (
	func() (BlockTransaction, bool),
	func(),
) {
	iterFunc, closeFunc := st.GetIterator(GetBlockTransactionKeyPrefixSource(source), reverse)

	return LoadBlockTransactionsInsideIterator(st, iterFunc, closeFunc)
}

func GetBlockTransactionsByConfirmed(st *sebakstorage.LevelDBBackend, reverse bool) (
	func() (BlockTransaction, bool),
	func(),
) {
	iterFunc, closeFunc := st.GetIterator(BlockTransactionPrefixConfirmed, reverse)

	return LoadBlockTransactionsInsideIterator(st, iterFunc, closeFunc)
}

var GetBlockTransactions = GetBlockTransactionsByConfirmed
