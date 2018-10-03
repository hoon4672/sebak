package operation

import (
	"encoding/json"

	"github.com/btcsuite/btcutil/base58"

	"boscoin.io/sebak/lib/common"
	"boscoin.io/sebak/lib/error"
)

type OperationType string

const (
	TypeCreateAccount        OperationType = "create-account"
	TypePayment              OperationType = "payment"
	TypeCongressVoting       OperationType = "congress-voting"
	TypeCongressVotingResult OperationType = "congress-voting-result"
	TypeCollectTxFee         OperationType = "collect-tx-fee"
	TypeInflation            OperationType = "inflation"
)

var KindsNormalTransaction map[OperationType]struct{} = map[OperationType]struct{}{
	TypeCreateAccount:        struct{}{},
	TypePayment:              struct{}{},
	TypeCongressVoting:       struct{}{},
	TypeCongressVotingResult: struct{}{},
}

// Limit is the number of operations to be included in a transaction
var Limit = 1000

type Operation struct {
	H Header
	B Body
}

func NewOperation(opb Body) (op Operation, err error) {
	var t OperationType
	switch opb.(type) {
	case CreateAccount:
		t = TypeCreateAccount
	case Payment:
		t = TypePayment
	case CollectTxFee:
		t = TypeCollectTxFee
	case Inflation:
		t = TypeInflation
	default:
		err = errors.ErrorUnknownOperationType
		return
	}

	op = Operation{
		H: Header{Type: t},
		B: opb,
	}

	return
}

type Header struct {
	Type OperationType `json:"type"`
}

type Body interface {
	//
	// Check that this transaction is self consistent
	//
	// This routine is used by the transaction checker and thus is part of consensus
	//
	// Params:
	//   networkid = Network id this operation was emitted on
	//
	// Returns:
	//   An `error` if that transaction is invalid, `nil` otherwise
	//
	IsWellFormed([]byte) error
	Serialize() ([]byte, error)
}

type Payable interface {
	Body
	TargetAddress() string
	GetAmount() common.Amount
}

func (o Operation) MakeHash() []byte {
	return common.MustMakeObjectHash(o)
}

func (o Operation) MakeHashString() string {
	return base58.Encode(o.MakeHash())
}

func (o Operation) IsWellFormed(networkID []byte) (err error) {
	return o.B.IsWellFormed(networkID)
}

func (o Operation) Serialize() (encoded []byte, err error) {
	return json.Marshal(o)
}

func (o Operation) String() string {
	encoded, _ := json.MarshalIndent(o, "", "  ")

	return string(encoded)
}

type envelop struct {
	H Header
	B interface{}
}

func (o *Operation) UnmarshalJSON(b []byte) (err error) {
	var raw json.RawMessage
	oj := envelop{
		B: &raw,
	}
	if err = json.Unmarshal(b, &oj); err != nil {
		return
	}

	o.H = oj.H

	var body Body
	if body, err = UnmarshalBodyJSON(oj.H.Type, raw); err != nil {
		return
	}
	o.B = body

	return
}

func UnmarshalBodyJSON(t OperationType, b []byte) (body Body, err error) {
	switch t {
	case TypeCreateAccount:
		var ob CreateAccount
		if err = json.Unmarshal(b, &ob); err != nil {
			return
		}
		body = ob
	case TypePayment:
		var ob Payment
		if err = json.Unmarshal(b, &ob); err != nil {
			return
		}
		body = ob
	case TypeCongressVoting:
		var ob CongressVoting
		if err = json.Unmarshal(b, &ob); err != nil {
			return
		}
		body = ob
	case TypeCongressVotingResult:
		var ob CongressVotingResult
		if err = json.Unmarshal(b, &ob); err != nil {
			return
		}
		body = ob
	case TypeCollectTxFee:
		var ob CollectTxFee
		if err = json.Unmarshal(b, &ob); err != nil {
			return
		}
		body = ob
	case TypeInflation:
		var ob Inflation
		if err = json.Unmarshal(b, &ob); err != nil {
			return
		}
		body = ob
	default:
		err = errors.ErrorInvalidOperation
		return
	}

	return
}