// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.8.4;

import "Errors.sol";
import "@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol";

contract ErrorContract {
    // DoubleEndedQueue
    function ThrowEmpty() public pure {
        revert DoubleEndedQueue.Empty();
    }

    function ThrowOutOfBounds() public pure {
        revert DoubleEndedQueue.OutOfBounds();
    }

    // BALANCE
    function ThrowInsufficientFunds(
        uint256 available,
        uint256 required
    ) public pure {
        revert InsufficientFunds(available, required);
    }

    // JOB
    function ThrowNoJob() public pure {
        revert NoJob();
    }

    function ThrowInvalidJob() public pure {
        revert InvalidJob();
    }

    function ThrowInvalidJobDefinition() public pure {
        revert InvalidJobDefinition();
    }

    function ThrowJobHotStatusOnly(JobStatus current) public pure {
        revert JobHotStatusOnly(current);
    }

    function ThrowJobColdStatusOnly(JobStatus current) public pure {
        revert JobColdStatusOnly(current);
    }

    function ThrowRunningScheduledStatusOnly(JobStatus current) public pure {
        revert RunningScheduledStatusOnly(current);
    }

    function ThrowMetaScheduledScheduledStatusOnly(
        JobStatus current
    ) public pure {
        revert MetaScheduledScheduledStatusOnly(current);
    }

    function ThrowRunningColdStatusOnly(JobStatus current) public pure {
        revert RunningColdStatusOnly(current);
    }

    function ThrowInvalidNNodes(uint256 current) public pure {
        revert InvalidNNodes(current);
    }

    function ThrowInvalidNCpu(uint256 current) public pure {
        revert InvalidNCpu(current);
    }

    function ThrowInvalidNMem(uint256 current) public pure {
        revert InvalidNMem(current);
    }

    function ThrowCustomerOnly(address current, address expected) public pure {
        revert CustomerOnly(current, expected);
    }

    // PERMISSION
    function ThrowJobProviderOnly(
        address current,
        address expected
    ) public pure {
        revert JobProviderOnly(current, expected);
    }

    function ThrowJobProviderThisOnly(
        address current,
        address expected
    ) public pure {
        revert JobProviderThisOnly(current, expected);
    }

    function ThrowOwnerOnly(address current, address expected) public pure {
        revert OwnerOnly(current, expected);
    }

    function ThrowCustomerMetaSchedulerProviderOnly() public pure {
        revert CustomerMetaSchedulerProviderOnly();
    }

    function ThrowMetashedulerProviderOnly() public pure {
        revert MetashedulerProviderOnly();
    }

    // PROVIDER
    function ThrowProviderAddrIsZero() public pure {
        revert ProviderAddrIsZero();
    }

    function ThrowProviderNotJoined() public pure {
        revert ProviderNotJoined();
    }

    function ThrowNoProvider() public pure {
        revert NoProvider();
    }

    function ThrowWaitingApprovalOnly() public pure {
        revert WaitingApprovalOnly();
    }

    function ThrowBanned() public pure {
        revert Banned();
    }

    // TIME
    function ThrowRemainingTimeAboveLimit(
        uint256 remaining,
        uint256 limit
    ) public pure {
        revert RemainingTimeAboveLimit(remaining, limit);
    }

    // OTHER
    function ThrowCreditAddrIsZero() public pure {
        revert CreditAddrIsZero();
    }

    function ThrowNoSpendingAuthority() public pure {
        revert NoSpendingAuthority();
    }

    function ThrowDivisionByZeroError() public pure {
        revert DivisionByZeroError();
    }

    // PROVIDERQUEUE
    function ThrowUninitialized() public pure {
        revert Uninitialized();
    }

    // STATE MACHINE
    function ThrowSameStatusError() public pure {
        revert SameStatusError();
    }

    function ThrowInvalidTransitionFromPending() public pure {
        revert InvalidTransitionFromPending();
    }

    function ThrowInvalidTransitionFromMetascheduled() public pure {
        revert InvalidTransitionFromMetascheduled();
    }

    function ThrowInvalidTransitionFromScheduled() public pure {
        revert InvalidTransitionFromScheduled();
    }

    function ThrowInvalidTransitionFromRunning() public pure {
        revert InvalidTransitionFromRunning();
    }
}
