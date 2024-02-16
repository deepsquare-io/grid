// Copyright (C) 2024 DeepSquare Association
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package metascheduler

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	errorsabi "github.com/deepsquare-io/grid/cli/types/abi/errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	errorsABI *abi.ABI
)

// PanicError is an native EVM error.
type PanicError byte

const (
	// PanicErrorAssertionError happens when `assert` fails.
	PanicErrorAssertionError PanicError = 0x1
	// PanicErrorArithmeticUnderOrOverflow happens when a number under or overflows.
	PanicErrorArithmeticUnderOrOverflow PanicError = 0x11
	// PanicErrorDivisionByZero happens when a number is divided by a zero denominator.
	PanicErrorDivisionByZero PanicError = 0x12
	// PanicErrorEnumConversionOutOfBounds happens when a number is out of bounds of an enum.
	PanicErrorEnumConversionOutOfBounds PanicError = 0x21
	// PanicErrorIncorrectlyEncodedStorageByteArray happens when byte array is badly encoded.
	PanicErrorIncorrectlyEncodedStorageByteArray PanicError = 0x22
	// PanicErrorPopOnEmptyArray happens when calling pop on an empty array.
	PanicErrorPopOnEmptyArray PanicError = 0x31
	// PanicErrorArrayAccessOutOfBounds happens when calling an element from an index out of bounds of an array.
	PanicErrorArrayAccessOutOfBounds PanicError = 0x32
	// PanicErrorTooMuchMemoryAllocated happens when there is too much memory allocated for the EVM.
	PanicErrorTooMuchMemoryAllocated PanicError = 0x41
	// PanicErrorZeroInitializedVariable happens when a variable is not initialized and it is forbidden to use the zero value of that variable.
	PanicErrorZeroInitializedVariable PanicError = 0x51
)

// IsPanicError checks if the byte of an error data is a panic error code.
func IsPanicError(value byte) bool {
	switch value {
	case byte(PanicErrorAssertionError),
		byte(PanicErrorArithmeticUnderOrOverflow),
		byte(PanicErrorDivisionByZero),
		byte(PanicErrorEnumConversionOutOfBounds),
		byte(PanicErrorIncorrectlyEncodedStorageByteArray),
		byte(PanicErrorPopOnEmptyArray),
		byte(PanicErrorArrayAccessOutOfBounds),
		byte(PanicErrorTooMuchMemoryAllocated),
		byte(PanicErrorZeroInitializedVariable):
		return true
	}
	return false
}

func (e PanicError) Error() string {
	switch e {
	case PanicErrorAssertionError:
		return "Assertion error"
	case PanicErrorArithmeticUnderOrOverflow:
		return "Arithmetic operation underflowed or overflowed outside of an unchecked block"
	case PanicErrorDivisionByZero:
		return "Division or modulo division by zero"
	case PanicErrorEnumConversionOutOfBounds:
		return "Tried to convert a value into an enum, but the value was too big or negative"
	case PanicErrorIncorrectlyEncodedStorageByteArray:
		return "Incorrectly encoded storage byte array"
	case PanicErrorPopOnEmptyArray:
		return ".pop() was called on an empty array"
	case PanicErrorArrayAccessOutOfBounds:
		return "Array accessed at an out-of-bounds or negative index"
	case PanicErrorTooMuchMemoryAllocated:
		return "Too much memory was allocated, or an array was created that is too large"
	case PanicErrorZeroInitializedVariable:
		return "Called a zero-initialized variable of internal function type"
	}
	return "Unknown error"
}

type DoubleEndedQueueEmpty struct{}

func ParseDoubleEndedQueueEmpty(inputs []interface{}) *DoubleEndedQueueEmpty {
	if len(inputs) != 0 {
		return nil
	}
	return &DoubleEndedQueueEmpty{}
}

func (e *DoubleEndedQueueEmpty) Error() string {
	return "QueueEmpty"
}

type DoubleEndedQueueOutOfBounds struct{}

func ParseDoubleEndedQueueOutOfBounds(inputs []interface{}) *DoubleEndedQueueOutOfBounds {
	if len(inputs) != 0 {
		return nil
	}
	return &DoubleEndedQueueOutOfBounds{}
}

func (e *DoubleEndedQueueOutOfBounds) Error() string {
	return "OutOfBounds"
}

type InvalidJob struct{}

func ParseInvalidJob(inputs []interface{}) *InvalidJob {
	if len(inputs) != 0 {
		return nil
	}
	return &InvalidJob{}
}

func (e *InvalidJob) Error() string {
	return "job not found or not valid"
}

type NoJob struct{}

func (e *NoJob) Error() string {
	return "no job"
}

func ParseNoJob(inputs []interface{}) error {
	if len(inputs) != 0 {
		return nil
	}
	return &NoJob{}
}

type InvalidNodesCount struct{}

func (e *InvalidNodesCount) Error() string {
	return "InvalidNodesCount"
}

func ParseInvalidNodesCount(inputs []interface{}) error {
	if len(inputs) != 0 {
		return nil
	}
	return &InvalidNodesCount{}
}

type ArrayLengthMismatch struct{}

func (e *ArrayLengthMismatch) Error() string {
	return "ArrayLengthMismatch"
}

func ParseArrayLengthMismatch(inputs []interface{}) error {
	if len(inputs) != 0 {
		return nil
	}
	return &ArrayLengthMismatch{}
}

type InvalidTotalMem struct{}

func (e *InvalidTotalMem) Error() string {
	return "InvalidTotalMem"
}

func ParseInvalidTotalMem(inputs []interface{}) error {
	if len(inputs) != 0 {
		return nil
	}
	return &InvalidTotalMem{}
}

type InvalidTotalCpus struct{}

func (e *InvalidTotalCpus) Error() string {
	return "InvalidTotalCpus"
}

func ParseInvalidTotalCpus(inputs []interface{}) error {
	if len(inputs) != 0 {
		return nil
	}
	return &InvalidTotalCpus{}
}

type NoProvider struct{}

func ParseNoProvider(_ []interface{}) *NoProvider {
	return &NoProvider{}
}

func (e *NoProvider) Error() string {
	return "NoProvider"
}

type WaitingApprovalOnly struct{}

func ParseWaitingApprovalOnly(_ []interface{}) *WaitingApprovalOnly {
	return &WaitingApprovalOnly{}
}

func (e *WaitingApprovalOnly) Error() string {
	return "WaitingApprovalOnly"
}

type Banned struct{}

func ParseBanned(_ []interface{}) *Banned {
	return &Banned{}
}

func (e *Banned) Error() string {
	return "Banned"
}

type AlreadyDone struct{}

func ParseAlreadyDone(_ []interface{}) *AlreadyDone {
	return &AlreadyDone{}
}

func (e *AlreadyDone) Error() string {
	return "AlreadyDone"
}

type JobHotStatusOnly struct {
	Current JobStatus
}

func (e *JobHotStatusOnly) Error() string {
	return fmt.Sprintf(
		"only applies to pending, meta-scheduled, scheduled and running job (current state: %s)",
		e.Current,
	)
}

func ParseJobHotStatusOnly(inputs []interface{}) *JobHotStatusOnly {
	if len(inputs) != 1 {
		return nil
	}
	return &JobHotStatusOnly{
		Current: JobStatus(inputs[0].(uint8)),
	}
}

type InvalidTransition struct {
	From JobStatus
	To   JobStatus
}

func (e *InvalidTransition) Error() string {
	return fmt.Sprintf(
		"InvalidTransition{From: %s, To: %s}",
		e.From,
		e.To,
	)
}

func ParseInvalidTransition(inputs []interface{}) *InvalidTransition {
	if len(inputs) != 2 {
		return nil
	}
	return &InvalidTransition{
		From: JobStatus(inputs[0].(uint8)),
		To:   JobStatus(inputs[1].(uint8)),
	}
}

type SameStatusError struct{}

func ParseSameStatusError(inputs []interface{}) *SameStatusError {
	if len(inputs) != 0 {
		return nil
	}
	return &SameStatusError{}
}

func (e *SameStatusError) Error() string {
	return "SameStatusError"
}

type InsufficientFunds struct {
	Available *big.Int
	Required  *big.Int
}

func (e *InsufficientFunds) Error() string {
	return fmt.Sprintf(
		"insufficient funds (available: %s, required: %s)",
		e.Available,
		e.Required,
	)
}

func ParseInsufficientFunds(inputs []interface{}) *InsufficientFunds {
	if len(inputs) != 2 {
		return nil
	}
	return &InsufficientFunds{
		Available: inputs[0].(*big.Int),
		Required:  inputs[1].(*big.Int),
	}
}

type InvalidJobDefinition struct{}

func ParseInvalidJobDefinition(
	inputs []interface{},
) *InvalidJobDefinition {
	if len(inputs) != 0 {
		return nil
	}
	return &InvalidJobDefinition{}
}

func (e *InvalidJobDefinition) Error() string {
	return "invalid job definition"
}

type RunningScheduledStatusOnly struct {
	Current JobStatus
}

func (e *RunningScheduledStatusOnly) Error() string {
	return fmt.Sprintf(
		"only applies to running and scheduled job (current state: %s)",
		e.Current,
	)
}

func ParseRunningScheduledStatusOnly(
	inputs []interface{},
) *RunningScheduledStatusOnly {
	if len(inputs) != 1 {
		return nil
	}
	return &RunningScheduledStatusOnly{
		Current: JobStatus(inputs[0].(uint8)),
	}
}

type MetaScheduledScheduledStatusOnly struct {
	Current JobStatus
}

func (e *MetaScheduledScheduledStatusOnly) Error() string {
	return fmt.Sprintf(
		"only applies to meta-scheduled and scheduled job (current state: %s)",
		e.Current,
	)
}

func ParseMetaScheduledScheduledStatusOnly(
	inputs []interface{},
) *MetaScheduledScheduledStatusOnly {
	if len(inputs) != 1 {
		return nil
	}
	return &MetaScheduledScheduledStatusOnly{
		Current: JobStatus(inputs[0].(uint8)),
	}
}

type RunningColdStatusOnly struct {
	Current JobStatus
}

func (e *RunningColdStatusOnly) Error() string {
	return fmt.Sprintf(
		"only applies to running or terminated job (current state: %s)",
		e.Current,
	)
}

func ParseRunningColdStatusOnly(
	inputs []interface{},
) *RunningColdStatusOnly {
	if len(inputs) != 1 {
		return nil
	}
	return &RunningColdStatusOnly{
		Current: JobStatus(inputs[0].(uint8)),
	}
}

type CustomerOnly struct {
	Current  common.Address
	Expected common.Address
}

func (e *CustomerOnly) Error() string {
	return fmt.Sprintf(
		"CustomerOnly{Current: %s, Expected: %s}",
		e.Current.Hex(),
		e.Expected.Hex(),
	)
}

func ParseCustomerOnly(inputs []interface{}) *CustomerOnly {
	if len(inputs) != 2 {
		return nil
	}
	return &CustomerOnly{
		Current:  inputs[0].(common.Address),
		Expected: inputs[1].(common.Address),
	}
}

type JobProviderOnly struct {
	Current  common.Address
	Expected common.Address
}

func (e *JobProviderOnly) Error() string {
	return fmt.Sprintf(
		"JobProviderOnly{Current: %s, Expected: %s}",
		e.Current.Hex(),
		e.Expected.Hex(),
	)
}

func ParseJobProviderOnly(inputs []interface{}) *JobProviderOnly {
	if len(inputs) != 2 {
		return nil
	}
	return &JobProviderOnly{
		Current:  inputs[0].(common.Address),
		Expected: inputs[1].(common.Address),
	}
}

type CustomerMetaSchedulerProviderOnly struct{}

func (e *CustomerMetaSchedulerProviderOnly) Error() string {
	return "CustomerMetaSchedulerProviderOnly"
}

func ParseCustomerMetaSchedulerProviderOnly(
	inputs []interface{},
) *CustomerMetaSchedulerProviderOnly {
	if len(inputs) != 0 {
		return nil
	}
	return &CustomerMetaSchedulerProviderOnly{}
}

type ProviderNotJoined struct{}

func ParseProviderNotJoined(inputs []interface{}) *ProviderNotJoined {
	if len(inputs) != 0 {
		return nil
	}
	return &ProviderNotJoined{}
}

func (e *ProviderNotJoined) Error() string {
	return "ProviderNotJoined"
}

type RemainingTimeAboveLimit struct {
	Remaining *big.Int
	Limit     *big.Int
}

func (e *RemainingTimeAboveLimit) Error() string {
	return fmt.Sprintf(
		"RemainingTimeAboveLimit{Remaining: %s, Limit: %s}",
		e.Remaining,
		e.Limit,
	)
}

func ParseRemainingTimeAboveLimit(
	inputs []interface{},
) *RemainingTimeAboveLimit {
	if len(inputs) != 2 {
		return nil
	}
	return &RemainingTimeAboveLimit{
		Remaining: inputs[0].(*big.Int),
		Limit:     inputs[1].(*big.Int),
	}
}

type NoSpendingAuthority struct{}

func ParseNoSpendingAuthority(inputs []interface{}) *NoSpendingAuthority {
	if len(inputs) != 0 {
		return nil
	}
	return &NoSpendingAuthority{}
}

func (e *NoSpendingAuthority) Error() string {
	return "NoSpendingAuthority"
}

type NewJobRequestDisabled struct{}

func ParseNewJobRequestDisabled(inputs []interface{}) *NewJobRequestDisabled {
	if len(inputs) != 0 {
		return nil
	}
	return &NewJobRequestDisabled{}
}

func (e *NewJobRequestDisabled) Error() string {
	return "NewJobRequestDisabled"
}

func init() {
	var err error
	errorsABI, err = errorsabi.ErrorContractMetaData.GetAbi()
	if err != nil {
		panic(fmt.Errorf("failed to read abi: %w", err))
	}
}

func WrapError(originalErr error) (newErr error) {
	if originalErr == nil {
		return nil
	}

	// Check if it's an RPC error
	var target rpc.DataError
	if ok := errors.As(originalErr, &target); !ok {
		return originalErr
	}

	// Fetch the rpc error hexData
	hexData, ok := target.ErrorData().(string)
	if !ok {
		return originalErr
	}
	data := common.FromHex(hexData)

	// Look for the error type in the ABI for
	if err := parseABIError(errorsABI.Errors, data, ParseDoubleEndedQueueError); err != nil {
		return err
	}

	if err := parseABIError(errorsABI.Errors, data, ParseError); err != nil {
		return err
	}

	// Check for panic error
	if IsPanicError(data[len(data)-1]) {
		return PanicError(data[len(data)-1])
	}

	return fmt.Errorf("%w, data: %s", originalErr, data)
}

func parseABIError(
	abiErrors map[string]abi.Error,
	data []byte,
	parseFunc func(string, []interface{}) error,
) error {
	for key, abiError := range abiErrors {
		parsedAbiError, err := abiError.Unpack(data)
		if err != nil {
			continue
		}

		// Check if error contains inputs
		inputs, ok := parsedAbiError.([]interface{})
		if !ok {
			return fmt.Errorf("%w, data: %+v", err, parsedAbiError)
		}

		// Parse the error
		if parsedErr := parseFunc(key, inputs); parsedErr != nil {
			return parsedErr
		}
	}
	return nil
}

func ParseDoubleEndedQueueError(name string, inputs []interface{}) error {
	switch name {
	case "Empty":
		return ParseDoubleEndedQueueEmpty(inputs)
	case "OutOfBounds":
		return ParseDoubleEndedQueueOutOfBounds(inputs)
	}
	return nil
}

func ParseError(name string, inputs []interface{}) error {
	switch name {
	case "InvalidJob":
		return ParseInvalidJob(inputs)
	case "NoJob":
		return ParseNoJob(inputs)
	case "InvalidNodesCount":
		return ParseInvalidNodesCount(inputs)
	case "ArrayLengthMismatch":
		return ParseArrayLengthMismatch(inputs)
	case "InvalidTotalMem":
		return ParseInvalidTotalMem(inputs)
	case "InvalidTotalCpus":
		return ParseInvalidTotalCpus(inputs)
	case "NoProvider":
		return ParseNoProvider(inputs)
	case "WaitingApprovalOnly":
		return ParseWaitingApprovalOnly(inputs)
	case "Banned":
		return ParseBanned(inputs)
	case "AlreadyDone":
		return ParseAlreadyDone(inputs)
	case "JobHotStatusOnly":
		return ParseJobHotStatusOnly(inputs)
	case "InvalidTransition":
		return ParseInvalidTransition(inputs)
	case "SameStatusError":
		return ParseSameStatusError(inputs)
	case "InsufficientFunds":
		return ParseInsufficientFunds(inputs)
	case "InvalidJobDefinition":
		return ParseInvalidJobDefinition(inputs)
	case "RunningScheduledStatusOnly":
		return ParseRunningScheduledStatusOnly(inputs)
	case "MetaScheduledScheduledStatusOnly":
		return ParseMetaScheduledScheduledStatusOnly(inputs)
	case "RunningColdStatusOnly":
		return ParseRunningColdStatusOnly(inputs)
	case "CustomerOnly":
		return ParseCustomerOnly(inputs)
	case "JobProviderOnly":
		return ParseJobProviderOnly(inputs)
	case "CustomerMetaSchedulerProviderOnly":
		return ParseCustomerMetaSchedulerProviderOnly(inputs)
	case "ProviderNotJoined":
		return ParseProviderNotJoined(inputs)
	case "RemainingTimeAboveLimit":
		return ParseRemainingTimeAboveLimit(inputs)
	case "NoSpendingAuthority":
		return ParseNoSpendingAuthority(inputs)
	case "NewJobRequestDisabled":
		return ParseNewJobRequestDisabled(inputs)
	}
	return nil
}

func CheckReceiptError(
	ctx context.Context,
	client bind.ContractCaller,
	tx *types.Transaction,
	receipt *types.Receipt,
) error {
	if receipt.Status == 1 {
		return nil
	}
	// Try to find reason in the receipt
	// Check gas
	if receipt.GasUsed == tx.Gas() {
		return vm.ErrOutOfGas
	}

	// Replay transaction (without mutation) to find the error
	from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		from, err = types.Sender(types.HomesteadSigner{}, tx)
		if err != nil {
			return fmt.Errorf(
				"tx failed (+unable to find 'from' address), tx: %s, status: %d",
				receipt.TxHash,
				receipt.Status,
			)
		}
	}

	// Replay transaction to find error reason
	_, err = client.CallContract(ctx, ethereum.CallMsg{
		To:         tx.To(),
		From:       from,
		Gas:        tx.Gas(),
		GasPrice:   tx.GasPrice(),
		GasFeeCap:  tx.GasFeeCap(),
		GasTipCap:  tx.GasTipCap(),
		Value:      tx.Value(),
		Data:       tx.Data(),
		AccessList: tx.AccessList(),
	}, receipt.BlockNumber)
	if err != nil {
		return fmt.Errorf(
			"tx failed, tx: %s, status: %d, error: %w",
			receipt.TxHash,
			receipt.Status,
			err,
		)
	}

	return fmt.Errorf("tx failed, tx: %s, status: %d", receipt.TxHash, receipt.Status)
}
