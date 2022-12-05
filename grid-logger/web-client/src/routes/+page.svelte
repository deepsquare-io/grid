<script lang="ts">
	import { GRPCService } from '$lib/services/grpc';
	import type { ReadResponse } from '@gen/ts/logger/v1alpha1/log';
	import { LoggerAPIClient } from '@gen/ts/logger/v1alpha1/log.client';
	import MetaMaskOnboarding from '@metamask/onboarding';
	import { GrpcStatusCode, GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
	import { RpcError } from '@protobuf-ts/runtime-rpc';
	import { ethers } from 'ethers';
	import { onDestroy, onMount } from 'svelte';
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

	onDestroy(() => {
		terminal?.dispose();
		grpcService?.stopReadAndWatch();
	});

	// List logs
	let logs: string[] = [];

	let errorMessage = '';

	// gRPC
	const url = 'http://localhost:9000';
	let loggerClient: LoggerAPIClient;
	let grpcService: GRPCService;
	let logStream: AsyncIterable<ReadResponse> | null = null;

	// onAddressDefined
	$: if (address) {
		onboarding.stopOnboarding();
		listAndWatch();
	}

	$: if (terminalElement && logStream !== null) {
		terminal.open(terminalElement);
		fitAddon.fit();
	}

	// Login button action
	async function doLogin() {
		if (!MetaMaskOnboarding.isMetaMaskInstalled()) {
			onboarding.startOnboarding();
		} else {
			await login();
		}
	}

	async function login() {
		await provider.send('eth_requestAccounts', []);
		signer = provider.getSigner();
		address = await signer.getAddress();
		grpcService = new GRPCService(loggerClient, signer);
	}

	// Selected table data action
	async function onLogNameSelected(log: string) {
		terminal.clear();
		isTermHidden = true;
		grpcService.stopReadAndWatch();

		try {
			logStream = await grpcService.readAndWatch(address.toLowerCase(), log);
			isTermHidden = false;

			for await (let resp of logStream) {
				terminal.writeln(resp.data);
			}
		} catch (e) {
			if (e instanceof RpcError) {
				if (e.code == GrpcStatusCode[GrpcStatusCode.ABORTED]) {
					console.log(e);
				} else {
					if (e instanceof Error) {
						errorMessage = e.message;
					}
					console.error(e);
				}
			}
			if (e instanceof Error) {
				errorMessage = e.message;
			}
			console.error(e);
		}
	}

	async function listAndWatch() {
		try {
			const responses = await grpcService.listAndWatch(address.toLowerCase());
			for await (let resp of responses) {
				logs = resp.logNames;
			}
		} catch (e) {
			if (e instanceof Error) {
				errorMessage = e.message;
			}

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

		// Configure metamask
		onboarding = new MetaMaskOnboarding();
		provider = new ethers.providers.Web3Provider(window.ethereum);

		if (MetaMaskOnboarding.isMetaMaskInstalled()) {
			await login();
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
				<table>
					<thead>
						<tr>
							<th>Available logs</th>
						</tr>
					</thead>
					<tbody>
						{#each logs as log}
							<tr>
								<!-- svelte-ignore a11y-click-events-have-key-events -->
								<td class="clickable" on:click={() => onLogNameSelected(log)}>{log}</td>
							</tr>
						{/each}
					</tbody>
				</table>

				{#if errorMessage || !isTermHidden}
					<footer transition:fade>
						{#if errorMessage}
							<span transition:fade class="error">{errorMessage}</span>
						{/if}

						<div transition:fade bind:this={terminalElement} />
					</footer>
				{/if}
			</article>
		{/if}
	</section>
</main>
