package errors

import (
	"errors"
	"fmt"

	pkgerrors "github.com/pkg/errors"
)

var New = errors.New
var Wrap = pkgerrors.Wrap
var Wrapf = pkgerrors.Wrapf

func Newf(err *Error, format string, args ...interface{}) error {
	return NewError(err.Code, fmt.Sprintf(format, args...))
}

var (
	BlockAlreadyExists                        = NewError(100, "already exists in block")
	HashDoesNotMatch                          = NewError(101, "`Hash` does not match")
	SignatureVerificationFailed               = NewError(102, "signature verification failed")
	BadPublicAddress                          = NewError(103, "failed to parse public address")
	InvalidFee                                = NewError(104, "invalid fee")
	InvalidOperation                          = NewError(105, "invalid operation")
	NewButKnownMessage                        = NewError(106, "received new, but known message")
	InvalidState                              = NewError(107, "found invalid state")
	InvalidVotingThresholdPolicy              = NewError(108, "invalid `VotingThresholdPolicy`")
	BallotEmptyMessage                        = NewError(109, "init state ballot does not have `Message`")
	InvalidHash                               = NewError(110, "invalid `Hash`")
	InvalidMessage                            = NewError(111, "invalid `Message`")
	BallotHasMessage                          = NewError(112, "none-init state ballot must not have `Message`")
	VotingResultAlreadyExists                 = NewError(113, "`VotingResult` already exists")
	VotingResultNotFound                      = NewError(114, "`VotingResult` not found")
	VotingResultFailedToSetState              = NewError(115, "failed to set the new state to `VotingResult`")
	VotingResultNotInBox                      = NewError(116, "ballot is not in here")
	BallotNoVoting                            = NewError(118, "ballot has no `Voting`")
	BallotNoNodeKey                           = NewError(119, "ballot has no `NodeKey`")
	VotingThresholdInvalidValidators          = NewError(120, "invalid validators")
	BallotHasInvalidState                     = NewError(121, "ballot has invalid state")
	VotingResultFailedToClose                 = NewError(122, "failed to close `VotingResult`")
	TransactionEmptyOperations                = NewError(123, "operations needs in transaction")
	AlreadySaved                              = NewError(124, "already saved")
	DuplicatedOperation                       = NewError(125, "duplicated operations in transaction")
	UnknownOperationType                      = NewError(126, "unknown operation type")
	TypeOperationBodyNotMatched               = NewError(127, "operation type and it's type does not match")
	BlockAccountDoesNotExists                 = NewError(128, "account does not exists in block")
	BlockAccountAlreadyExists                 = NewError(129, "account already exists in block")
	AccountBalanceUnderZero                   = NewError(130, "account balance will be under zero")
	MaximumBalanceReached                     = NewError(131, "monetary amount would be greater than the total supply of coins")
	StorageRecordDoesNotExist                 = NewError(132, "record does not exist in storage")
	TransactionInvalidSequenceID              = NewError(133, "invalid sequenceID found")
	BlockTransactionDoesNotExists             = NewError(134, "transaction does not exists in block")
	BlockOperationDoesNotExists               = NewError(135, "operation does not exists in block")
	RoundVoteNotFound                         = NewError(136, "`RoundVote` not found")
	BlockNotFound                             = NewError(137, "Block not found")
	TransactionExcessAbilityToPay             = NewError(138, "Transaction requests over ability to pay")
	TransactionSameSourceInBallot             = NewError(139, "Same transaction source found in ballot")
	TransactionNotFound                       = NewError(140, "Transaction not found")
	BallotFromUnknownValidator                = NewError(141, "ballot from unknown validator")
	InvalidVotingBasis                        = NewError(142, "invalid voting basis")
	BallotAlreadyVoted                        = NewError(143, "ballot already voted")
	BallotHasOverMaxTransactionsInBallot      = NewError(144, "too many transactions in ballot")
	MessageHasIncorrectTime                   = NewError(145, "time in message is not correct")
	InvalidQueryString                        = NewError(146, "found invalid query string")
	InvalidContentType                        = NewError(147, "found invalid 'Content-Type'")
	StorageRecordAlreadyExists                = NewError(148, "record already exists in storage")
	StorageCoreError                          = NewError(149, "storage error")
	ContentTypeNotJSON                        = NewError(150, "`Content-Type` must be 'application/json'")
	TransactionHasOverMaxOperations           = NewError(151, "too many operations in transaction")
	OperationAmountUnderflow                  = NewError(152, "invalid `Amount`: lower than 1")
	FrozenAccountNoDeposit                    = NewError(153, "frozen account can not receive payment")
	FrozenAccountCreationWholeUnit            = NewError(154, "frozen account balance must be a whole number of units (10k)")
	FrozenAccountMustWithdrawEverything       = NewError(155, "frozen account can only withdraw the full amount (minus tx fee)")
	InsufficientAmountNewAccount              = NewError(156, "insufficient amount for new account")
	OperationBodyInsufficient                 = NewError(157, "operation body insufficient")
	OperationAmountOverflow                   = NewError(158, "invalid `Amount`: over than expected")
	WrongBlockFound                           = NewError(159, "wrong Block found")
	InvalidProposerTransaction                = NewError(160, "invalid proposer transaction found")
	InvalidInflationRatio                     = NewError(161, "invalid inflation ratio found")
	NotImplemented                            = NewError(162, "not implemented")
	HTTPProblem                               = NewError(163, "http failed to get response")
	InvalidTransaction                        = NewError(164, "invalid transaction")
	NotMatcHTTPRouter                         = NewError(165, "doesn't match http router")
	UnfreezingFromInvalidAccount              = NewError(166, "unfreezing should be done from a frozen account")
	UnfreezingToInvalidLinkedAccount          = NewError(167, "unfreezing should be done to a valid linked account")
	UnfreezingNotReachedExpiration            = NewError(168, "unfreezing should pass 241920 blockheight from unfreezing request")
	FrozenAccountMustCreatedFromLinkedAccount = NewError(169, "frozen account create-transaction must be generated from the linked account")
	UnfreezingRequestNotRequested             = NewError(170, "unfreezing must be generated after the unfreezing request")
	UnfreezingRequestAlreadyReceived          = NewError(171, "unfreezing request already received from client")
	TooManyRequests                           = NewError(172, "too many requests; reached limit")
	HTTPServerError                           = NewError(173, "Internal Server Error")
	BlockTransactionHistoryDoesNotExists      = NewError(174, "transaction history does not exists in block")
	NotCommittable                            = NewError(175, "not Committable")
	TransactionSameSourceInPool               = NewError(176, "Same transaction source found in pool")
	AlreadyCommittable                        = NewError(177, "already Committable")
	FailedToSaveBlockOperaton                 = NewError(178, "failed to save BlockOperation")
	NodeNotFound                              = NewError(179, "Node not found")
	InvalidMessageVersion                     = NewError(180, "message version is invalid")
	InvalidGenesisOption                      = NewError(181, "--genesis expects '<genesis address>,<common account>[,balance]")
	NotPublicKey                              = NewError(182, "not public key")
	InflationPFResultMissed                   = NewError(183, "voting result cannot be retrieved")
	InflationPFAmountMissMatched              = NewError(184, "inflation pf amount missmatched")
	InflationPFFundingAddressMissMatched      = NewError(185, "inflation pf funding address missmatched")
	CongressAddressMisMatched                 = NewError(186, "congress address mismatched")
	PageQueryLimitMaxExceed                   = NewError(187, "maximum value of limit query parameter exceed")
	TransactionPoolFull                       = NewError(188, "transaction pool is full")
	TransactionAlreadyExistsInPool            = NewError(189, "transaction already exists in pool")
)
