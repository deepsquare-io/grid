package metascheduler

import (
	"errors"
	"fmt"
	"math/big"

	errorsabi "github.com/deepsquare-io/the-grid/cli/internal/abi/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	errorsABI *abi.ABI
)

type PanicError byte

const (
	PanicErrorAssertionError                     PanicError = 0x1
	PanicErrorArithmeticUnderOrOverflow          PanicError = 0x11
	PanicErrorDivisionByZero                     PanicError = 0x12
	PanicErrorEnumConversionOutOfBounds          PanicError = 0x21
	PanicErrorIncorrectlyEncodedStorageByteArray PanicError = 0x22
	PanicErrorPopOnEmptyArray                    PanicError = 0x31
	PanicErrorArrayAccessOutOfBounds             PanicError = 0x32
	PanicErrorTooMuchMemoryAllocated             PanicError = 0x41
	PanicErrorZeroInitializedVariable            PanicError = 0x51
)

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

type InsufficientFunds struct {
	Available *big.Int
	Required  *big.Int
}

func (e *InsufficientFunds) Error() string {
	return fmt.Sprintf(
		"InsufficientFunds{Available: %s, Required: %s}",
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

type NoJob struct{}

func (e *NoJob) Error() string {
	return "NoJob"
}

func ParseNoJob(inputs []interface{}) error {
	if len(inputs) != 0 {
		return nil
	}
	return &NoJob{}
}

type InvalidJob struct{}

func ParseInvalidJob(inputs []interface{}) *InvalidJob {
	if len(inputs) != 0 {
		return nil
	}
	return &InvalidJob{}
}

func (e *InvalidJob) Error() string {
	return "InvalidJob"
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
	return "InvalidJobDefinition"
}

type JobHotStatusOnly struct {
	Current JobStatus
}

func (e *JobHotStatusOnly) Error() string {
	return fmt.Sprintf(
		"JobHotStatusOnly{Current: %s}",
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

type RunningScheduledStatusOnly struct {
	Current JobStatus
}

func (e *RunningScheduledStatusOnly) Error() string {
	return fmt.Sprintf(
		"RunningScheduledStatusOnly{Current: %s}",
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
		"MetaScheduledScheduledStatusOnly{Current: %s}",
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
		"RunningColdStatusOnly{Current: %s}",
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

type InvalidNNodes struct {
	Current *big.Int
}

func (e *InvalidNNodes) Error() string {
	return fmt.Sprintf(
		"InvalidNNodes{Current: %s}",
		e.Current,
	)
}

func ParseInvalidNNodes(
	inputs []interface{},
) *InvalidNNodes {
	if len(inputs) != 1 {
		return nil
	}
	return &InvalidNNodes{
		Current: inputs[0].(*big.Int),
	}
}

type InvalidNCpu struct {
	Current *big.Int
}

func (e *InvalidNCpu) Error() string {
	return fmt.Sprintf(
		"InvalidNCpu{Current: %s}",
		e.Current,
	)
}

func ParseInvalidNCpu(
	inputs []interface{},
) *InvalidNCpu {
	if len(inputs) != 1 {
		return nil
	}
	return &InvalidNCpu{
		Current: inputs[0].(*big.Int),
	}
}

type InvalidNMem struct {
	Current *big.Int
}

func (e *InvalidNMem) Error() string {
	return fmt.Sprintf(
		"InvalidNMem{Current: %s}",
		e.Current,
	)
}

func ParseInvalidNMem(
	inputs []interface{},
) *InvalidNMem {
	if len(inputs) != 1 {
		return nil
	}
	return &InvalidNMem{
		Current: inputs[0].(*big.Int),
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

type JobProviderThisOnly struct {
	Current  common.Address
	Expected common.Address
}

func (e *JobProviderThisOnly) Error() string {
	return fmt.Sprintf(
		"JobProviderThisOnly{Current: %s, Expected: %s}",
		e.Current.Hex(),
		e.Expected.Hex(),
	)
}

func ParseJobProviderThisOnly(inputs []interface{}) *JobProviderThisOnly {
	if len(inputs) != 2 {
		return nil
	}
	return &JobProviderThisOnly{
		Current:  inputs[0].(common.Address),
		Expected: inputs[1].(common.Address),
	}
}

type OwnerOnly struct {
	Current  common.Address
	Expected common.Address
}

func (e *OwnerOnly) Error() string {
	return fmt.Sprintf(
		"OwnerOnly{Current: %s, Expected: %s}",
		e.Current.Hex(),
		e.Expected.Hex(),
	)
}

func ParseOwnerOnly(inputs []interface{}) *OwnerOnly {
	if len(inputs) != 2 {
		return nil
	}
	return &OwnerOnly{
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
	return &CustomerMetaSchedulerProviderOnly{}
}

type MetashedulerProviderOnly struct{}

func (e *MetashedulerProviderOnly) Error() string {
	return "MetashedulerProviderOnly"
}

func ParseMetashedulerProviderOnly(
	inputs []interface{},
) *MetashedulerProviderOnly {
	return &MetashedulerProviderOnly{}
}

type ProviderAddrIsZero struct{}

func ParseProviderAddrIsZero(inputs []interface{}) *ProviderAddrIsZero {
	return &ProviderAddrIsZero{}
}

func (e *ProviderAddrIsZero) Error() string {
	return "ProviderAddrIsZero"
}

type ProviderNotJoined struct{}

func ParseProviderNotJoined(inputs []interface{}) *ProviderNotJoined {
	return &ProviderNotJoined{}
}

func (e *ProviderNotJoined) Error() string {
	return "ProviderNotJoined"
}

type NoProvider struct{}

func ParseNoProvider(inputs []interface{}) *NoProvider {
	return &NoProvider{}
}

func (e *NoProvider) Error() string {
	return "NoProvider"
}

type WaitingApprovalOnly struct{}

func ParseWaitingApprovalOnly(inputs []interface{}) *WaitingApprovalOnly {
	return &WaitingApprovalOnly{}
}

func (e *WaitingApprovalOnly) Error() string {
	return "WaitingApprovalOnly"
}

type Banned struct{}

func ParseBanned(inputs []interface{}) *Banned {
	return &Banned{}
}

func (e *Banned) Error() string {
	return "Banned"
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

type CreditAddrIsZero struct{}

func ParseCreditAddrIsZero(inputs []interface{}) *CreditAddrIsZero {
	if len(inputs) != 0 {
		return nil
	}
	return &CreditAddrIsZero{}
}

func (e *CreditAddrIsZero) Error() string {
	return "CreditAddrIsZero"
}

type NoSpendingAuthority struct{}

func ParseNoSpendingAuthority(inputs []interface{}) *NoSpendingAuthority {
	return &NoSpendingAuthority{}
}

func (e *NoSpendingAuthority) Error() string {
	return "NoSpendingAuthority"
}

type DivisionByZeroError struct{}

func ParseDivisionByZeroError(inputs []interface{}) *DivisionByZeroError {
	if len(inputs) != 0 {
		return nil
	}
	return &DivisionByZeroError{}
}

func (e *DivisionByZeroError) Error() string {
	return "DivisionByZeroError"
}

type Uninitialized struct{}

func (e *Uninitialized) Error() string {
	return "Uninitialized"
}

func ParseUninitialized(inputs []interface{}) error {
	if len(inputs) != 0 {
		return nil
	}
	return &Uninitialized{}
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

type InvalidTransitionFromPending struct{}

func ParseInvalidTransitionFromPending(
	inputs []interface{},
) *InvalidTransitionFromPending {
	if len(inputs) != 0 {
		return nil
	}
	return &InvalidTransitionFromPending{}
}

func (e *InvalidTransitionFromPending) Error() string {
	return "InvalidTransitionFromPending"
}

type InvalidTransitionFromMetascheduled struct{}

func ParseInvalidTransitionFromMetascheduled(
	inputs []interface{},
) *InvalidTransitionFromMetascheduled {
	if len(inputs) != 0 {
		return nil
	}
	return &InvalidTransitionFromMetascheduled{}
}

func (e *InvalidTransitionFromMetascheduled) Error() string {
	return "InvalidTransitionFromMetascheduled"
}

type InvalidTransitionFromScheduled struct{}

func ParseInvalidTransitionFromScheduled(
	inputs []interface{},
) *InvalidTransitionFromScheduled {
	if len(inputs) != 0 {
		return nil
	}
	return &InvalidTransitionFromScheduled{}
}

func (e *InvalidTransitionFromScheduled) Error() string {
	return "InvalidTransitionFromScheduled"
}

type InvalidTransitionFromRunning struct{}

func ParseInvalidTransitionFromRunning(
	inputs []interface{},
) *InvalidTransitionFromRunning {
	if len(inputs) != 0 {
		return nil
	}
	return &InvalidTransitionFromRunning{}
}

func (e *InvalidTransitionFromRunning) Error() string {
	return "InvalidTransitionFromRunning"
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
	case "InsufficientFunds":
		return ParseInsufficientFunds(inputs)
	case "InvalidJob":
		return ParseInvalidJob(inputs)
	case "InvalidJobDefinition":
		return ParseInvalidJobDefinition(inputs)
	case "JobHotStatusOnly":
		return ParseJobHotStatusOnly(inputs)
	case "RunningScheduledStatusOnly":
		return ParseRunningScheduledStatusOnly(inputs)
	case "MetaScheduledScheduledStatusOnly":
		return ParseMetaScheduledScheduledStatusOnly(inputs)
	case "RunningColdStatusOnly":
		return ParseRunningColdStatusOnly(inputs)
	case "InvalidNNodes":
		return ParseInvalidNNodes(inputs)
	case "InvalidNCpu":
		return ParseInvalidNCpu(inputs)
	case "InvalidNMem":
		return ParseInvalidNMem(inputs)
	case "CustomerOnly":
		return ParseCustomerOnly(inputs)
	case "JobProviderOnly":
		return ParseJobProviderOnly(inputs)
	case "JobProviderThisOnly":
		return ParseJobProviderThisOnly(inputs)
	case "OwnerOnly":
		return ParseOwnerOnly(inputs)
	case "CustomerMetaSchedulerProviderOnly":
		return ParseCustomerMetaSchedulerProviderOnly(inputs)
	case "MetashedulerProviderOnly":
		return ParseMetashedulerProviderOnly(inputs)
	case "ProviderAddrIsZero":
		return ParseProviderAddrIsZero(inputs)
	case "ProviderNotJoined":
		return ParseProviderNotJoined(inputs)
	case "NoProvider":
		return ParseNoProvider(inputs)
	case "WaitingApprovalOnly":
		return ParseWaitingApprovalOnly(inputs)
	case "Banned":
		return ParseBanned(inputs)
	case "RemainingTimeAboveLimit":
		return ParseRemainingTimeAboveLimit(inputs)
	case "CreditAddrIsZero":
		return ParseCreditAddrIsZero(inputs)
	case "NoSpendingAuthority":
		return ParseNoSpendingAuthority(inputs)
	case "DivisionByZeroError":
		return ParseDivisionByZeroError(inputs)
	case "SameStatusError":
		return ParseSameStatusError(inputs)
	case "NoJob":
		return ParseNoJob(inputs)
	case "Uninitialized":
		return ParseUninitialized(inputs)
	case "InvalidTransitionFromPending":
		return ParseInvalidTransitionFromPending(inputs)
	case "InvalidTransitionFromMetascheduled":
		return ParseInvalidTransitionFromMetascheduled(inputs)
	case "InvalidTransitionFromScheduled":
		return ParseInvalidTransitionFromScheduled(inputs)
	case "InvalidTransitionFromRunning":
		return ParseInvalidTransitionFromRunning(inputs)
	}
	return nil
}
