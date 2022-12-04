import type { AuthAPIClient } from '@gen/ts/auth/v1alpha1/auth.client';
import type { LoggerAPIClient } from '@gen/ts/logger/v1alpha1/log.client';
import { RpcError } from '@protobuf-ts/runtime-rpc';
import type { ReadResponse } from '@gen/ts/logger/v1alpha1/log';
import { ethers } from 'ethers';
import { GrpcStatusCode } from '@protobuf-ts/grpcweb-transport';

export class GRPCService {
	private abortReadAndWatch: AbortController | null = null;

	constructor(private authClient: AuthAPIClient, private loggerClient: LoggerAPIClient) {}

	async getNonce(address: string): Promise<Uint8Array> {
		const { response } = await this.authClient.nonce({ address: address });
		return response.nonce;
	}

	async signIn(address: string, nonce: Uint8Array, signature: Uint8Array): Promise<string> {
		const { response } = await this.authClient.signIn({
			address: address,
			nonce: nonce,
			sig: signature
		});
		return response.accessToken;
	}

	async register(address: string) {
		await this.authClient.register({ address: address });
	}

	async registerOrSignIn(address: string, signer: ethers.Signer): Promise<string> {
		let nonce: Uint8Array;
		try {
			nonce = await this.getNonce(address);
		} catch (e) {
			if (e instanceof RpcError) {
				if (e.code === GrpcStatusCode[GrpcStatusCode.NOT_FOUND]) {
					console.debug('user not found, trying to register');
					await this.register(address);
					nonce = await this.getNonce(address);
				} else {
					throw e;
				}
			} else {
				throw e;
			}
		}
		const signature = await signer.signMessage(nonce);
		return this.signIn(address, nonce, ethers.utils.arrayify(signature));
	}

	readAndWatch(logName: string, token: string): AsyncIterable<ReadResponse> {
		this.abortReadAndWatch = new AbortController();
		const call = this.loggerClient.read(
			{ logName: logName },
			{
				meta: {
					authorization: `Bearer ${token}`
				},
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
