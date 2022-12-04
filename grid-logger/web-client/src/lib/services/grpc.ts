import type { ReadResponse } from '@gen/ts/logger/v1alpha1/log';
import type { LoggerAPIClient } from '@gen/ts/logger/v1alpha1/log.client';

export class GRPCService {
	private abortReadAndWatch: AbortController | null = null;

	constructor(private loggerClient: LoggerAPIClient) {}

	readAndWatch(
		address: string,
		logName: string,
		timestamp: bigint,
		signedHash: Uint8Array
	): AsyncIterable<ReadResponse> {
		this.abortReadAndWatch = new AbortController();
		const call = this.loggerClient.read(
			{
				address: address,
				logName: logName,
				timestamp: timestamp,
				signedHash: signedHash
			},
			{
				abort: this.abortReadAndWatch.signal
			}
		);
		return call.responses;
	}

	stopReadAndWatch() {
		if (this.abortReadAndWatch !== null) {
			this.abortReadAndWatch.abort();
			this.abortReadAndWatch = null;
		}
	}
}
