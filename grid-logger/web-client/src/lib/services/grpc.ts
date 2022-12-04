import type { ReadResponse, WatchListResponse } from '@gen/ts/logger/v1alpha1/log';
import type { LoggerAPIClient } from '@gen/ts/logger/v1alpha1/log.client';
import { ethers } from 'ethers';

export class GRPCService {
	private abortReadAndWatch: AbortController | null = null;

	constructor(private loggerClient: LoggerAPIClient, private signer: ethers.Signer) {}

	async readAndWatch(address: string, logName: string): Promise<AsyncIterable<ReadResponse>> {
		this.abortReadAndWatch = new AbortController();
		const timestamp = Date.now();
		const msg = `read:${address.toLowerCase()}/${logName}/${timestamp}`;
		const signedHash = await this.signer.signMessage(msg);

		const call = this.loggerClient.read(
			{
				address: address,
				logName: logName,
				timestamp: BigInt(timestamp),
				signedHash: ethers.utils.arrayify(signedHash)
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

	async listAndWatch(address: string): Promise<AsyncIterable<WatchListResponse>> {
		const timestamp = Date.now();
		const msg = `watchList:${address.toLowerCase()}/${timestamp}`;
		const signedHash = await this.signer.signMessage(msg);

		const { responses } = this.loggerClient.watchList({
			address: address,
			timestamp: BigInt(timestamp),
			signedHash: ethers.utils.arrayify(signedHash)
		});
		return responses;
	}
}
