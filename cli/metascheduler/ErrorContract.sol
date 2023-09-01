// SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.8.4;

import 'Errors.sol';
import '@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol';

contract ErrorContract {
  // DoubleEndedQueue
  function ThrowEmpty() public pure {
    revert DoubleEndedQueue.Empty();
  }

  function ThrowOutOfBounds() public pure {
    revert DoubleEndedQueue.OutOfBounds();
  }

  // BALANCE
  function ThrowInsufficientFunds(uint256 available, uint256 required) public pure {
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

  function ThrowMetaScheduledScheduledStatusOnly(JobStatus current) public pure {
    revert MetaScheduledScheduledStatusOnly(current);
  }

  function ThrowRunningColdStatusOnly(JobStatus current) public pure {
    revert RunningColdStatusOnly(current);
  }

  function ThrowInvalidNodesCount() public pure {
    revert InvalidNodesCount();
  }

  function ThrowArrayLengthMismatch() public pure {
    revert ArrayLengthMismatch();
  }

  function ThrowInvalidTotalMem() public pure {
    revert InvalidTotalMem();
  }

  function ThrowInvalidTotalCpus() public pure {
    revert InvalidTotalCpus();
  }

  function ThrowCustomerOnly(address current, address expected) public pure {
    revert CustomerOnly(current, expected);
  }

  // PERMISSION
  function ThrowJobProviderOnly(address current, address expected) public pure {
    revert JobProviderOnly(current, expected);
  }

  function ThrowOwnerOnly(address current, address expected) public pure {
    revert OwnerOnly(current, expected);
  }

  function ThrowCustomerMetaSchedulerProviderOnly() public pure {
    revert CustomerMetaSchedulerProviderOnly();
  }

  function ThrowMetaschedulerProviderOnly() public pure {
    revert MetaschedulerProviderOnly();
  }

  // PROVIDER
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
  function ThrowRemainingTimeAboveLimit(uint256 remaining, uint256 limit) public pure {
    revert RemainingTimeAboveLimit(remaining, limit);
  }

  // OTHER
  function ThrowNoSpendingAuthority() public pure {
    revert NoSpendingAuthority();
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

  function ThrowNewJobRequestDisabled() public pure {
    revert NewJobRequestDisabled();
  }
}
