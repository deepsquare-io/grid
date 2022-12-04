<script lang="ts">
	import { GRPCService } from '$lib/services/grpc';
	import { token } from '$lib/stores/jwt';
	import { AuthAPIClient } from '@gen/ts/auth/v1alpha1/auth.client';
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
	let authClient: AuthAPIClient;
	let loggerClient: LoggerAPIClient;
	let grpcService: GRPCService;

	// Fields
	let logName: string;

	// Refresh login button
	$: if ($token) {
		onboarding.stopOnboarding();
	}

	// Refresh terminal
	$: if ($token) {
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

			// Do gRPC
			try {
				token.set(await grpcService.registerOrSignIn(address, signer));
				console.log($token);
			} catch (e) {
				// TODO: handle error
				console.log(e);
			}
		}
	}

	// Read and watch button action
	async function doReadAndWatch() {
		isTermHidden = false;
		terminal.clear();
		grpcService.stopReadAndWatch();

		try {
			const responses = grpcService.readAndWatch(logName, $token);
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
		authClient = new AuthAPIClient(transport);
		loggerClient = new LoggerAPIClient(transport);
		grpcService = new GRPCService(authClient, loggerClient);

		// Configure metamask
		onboarding = new MetaMaskOnboarding();
		provider = new ethers.providers.Web3Provider(window.ethereum);
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
		{#if !$token}
			<dialog transition:fade open>
				<article>
					<button on:click={doLogin}>Login with MetaMask</button>
				</article>
			</dialog>
		{/if}

		{#if $token}
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

		<div transition:fade bind:this={terminalElement} hidden={isTermHidden} />
	</section>
</main>
