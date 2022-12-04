<script lang="ts">
	import { GRPCService } from '$lib/services/grpc';
	import { LoggerAPIClient } from '@gen/ts/logger/v1alpha1/log.client';
	import MetaMaskOnboarding from '@metamask/onboarding';
	import { GrpcStatusCode, GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
	import { RpcError } from '@protobuf-ts/runtime-rpc';
	import { ethers } from 'ethers';
	import { onMount } from 'svelte';
	import { fade } from 'svelte/transition';
	import type { Terminal } from 'xterm';
	import type { FitAddon } from 'xterm-addon-fit';

	// Web3
	let onboarding: MetaMaskOnboarding;
	let address: string = '';
	let signer: ethers.providers.JsonRpcSigner;
	let provider: ethers.providers.Web3Provider;

	// Terminal
	let terminalElement: HTMLElement;
	let terminal: Terminal;
	let fitAddon: FitAddon;
	let isTermHidden = true;

	// gRPC
	const url = 'http://localhost:9000';
	let loggerClient: LoggerAPIClient;
	let grpcService: GRPCService;

	// Fields
	let logName: string;

	// Refresh login button
	$: if (address) {
		onboarding.stopOnboarding();
	}

	// Refresh terminal
	$: if (address) {
		terminal.open(terminalElement);
		fitAddon.fit();
	}

	// Login button action
	async function doLogin() {
		if (!MetaMaskOnboarding.isMetaMaskInstalled()) {
			onboarding.startOnboarding();
		} else {
			await provider.send('eth_requestAccounts', []);
			signer = provider.getSigner();
			address = await signer.getAddress();
		}
	}

	// Read and watch button action
	async function doReadAndWatch() {
		terminal.clear();
		grpcService.stopReadAndWatch();

		try {
			const timestamp = Date.now();
			const msg = `${address.toLowerCase()}/${logName}/${timestamp}`;
			const sig = await signer.signMessage(msg);
			const responses = grpcService.readAndWatch(
				address.toLowerCase(),
				logName,
				BigInt(timestamp),
				ethers.utils.arrayify(sig)
			);
			isTermHidden = false;
			for await (let message of responses) {
				terminal.writeln(message.data);
			}
		} catch (e) {
			if (e instanceof RpcError) {
				if (e.code == GrpcStatusCode[GrpcStatusCode.ABORTED]) {
					console.log(e);
				} else {
					// TODO: handle error
					console.error(e);
				}
			}
			// TODO: handle error
			console.error(e);
		}
	}

	onMount(async () => {
		const { Terminal } = await import('xterm');
		const { FitAddon } = await import('xterm-addon-fit');

		// Configure terminal
		terminal = new Terminal({
			rows: 50
		});
		fitAddon = new FitAddon();
		terminal.loadAddon(fitAddon);

		// Configure gRPC
		const transport = new GrpcWebFetchTransport({ baseUrl: url });
		loggerClient = new LoggerAPIClient(transport);
		grpcService = new GRPCService(loggerClient);

		// Configure metamask
		onboarding = new MetaMaskOnboarding();
		provider = new ethers.providers.Web3Provider(window.ethereum);

		if (MetaMaskOnboarding.isMetaMaskInstalled()) {
			await provider.send('eth_requestAccounts', []);
			signer = provider.getSigner();
			address = await signer.getAddress();
		}
	});
</script>

<header class="container">
	<hgroup>
		<h1>DeepSquare Grid Logs</h1>
		<h2>Live display of running and finished jobs!</h2>
	</hgroup>
</header>

<main class="container">
	<section>
		{#if !address}
			<dialog transition:fade open>
				<article>
					<button on:click={doLogin}>Login with MetaMask</button>
				</article>
			</dialog>
		{/if}

		{#if address}
			<article>
				<form transition:fade on:submit|preventDefault={doReadAndWatch}>
					<label for="logname">
						Log Name
						<input name="logname" bind:value={logName} />
					</label>
					<footer>
						<button>Fetch</button>
					</footer>
				</form>
			</article>
		{/if}

		{#key !address}
			<div in:fade bind:this={terminalElement} hidden={isTermHidden} />
		{/key}
	</section>
</main>
