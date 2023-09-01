// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

import "interfaces/IJobRepository.sol";
import "interfaces/IProviderJobQueues.sol";
import "interfaces/IProviderManager.sol";
import "Tools.sol";
import "Metascheduler.sol";
import "@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol";

contract ErrorContract {
    // DoubleEndedQueue
    function ThrowEmpty() public pure {
        revert DoubleEndedQueue.Empty();
    }

    function ThrowOutOfBounds() public pure {
        revert DoubleEndedQueue.OutOfBounds();
    }

    // IJobRepository
    function ThrowInvalidJob() public pure {
        revert InvalidJob();
    }

    // IProviderJobQueues
    function ThrowNoJob() public pure {
        revert NoJob();
    }

    // IProviderManager
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

    function ThrowNoProvider() public pure {
        revert NoProvider();
    }

    function ThrowWaitingApprovalOnly() public pure {
        revert WaitingApprovalOnly();
    }

    function ThrowBanned() public pure {
        revert Banned();
    }

    // Tools
    function ThrowJobHotStatusOnly(JobStatus current) public pure {
        revert JobHotStatusOnly(current);
    }

    function ThrowInvalidTransition(JobStatus from, JobStatus to) public pure {
        revert InvalidTransition(from, to);
    }

    function ThrowSameStatusError() public pure {
        revert SameStatusError();
    }

    // Metascheduler
    function ThrowInsufficientFunds(
        uint256 available,
        uint256 required
    ) public pure {
        revert InsufficientFunds(available, required);
    }

    function ThrowInvalidJobDefinition() public pure {
        revert InvalidJobDefinition();
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

    function ThrowCustomerOnly(address current, address expected) public pure {
        revert CustomerOnly(current, expected);
    }

    function ThrowJobProviderOnly(
        address current,
        address expected
    ) public pure {
        revert JobProviderOnly(current, expected);
    }

    function ThrowCustomerMetaSchedulerProviderOnly() public pure {
        revert CustomerMetaSchedulerProviderOnly();
    }

    function ThrowProviderNotJoined() public pure {
        revert ProviderNotJoined();
    }

    function ThrowRemainingTimeAboveLimit(
        uint256 remaining,
        uint256 limit
    ) public pure {
        revert RemainingTimeAboveLimit(remaining, limit);
    }

    function ThrowNoSpendingAuthority() public pure {
        revert NoSpendingAuthority();
    }

    function ThrowNewJobRequestDisabled() public pure {
        revert NewJobRequestDisabled();
    }
}
