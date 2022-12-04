<script lang="ts">
	import { onMount } from 'svelte';
	import { ethers } from 'ethers';
	import MetaMaskOnboarding from '@metamask/onboarding';
	import type { Terminal } from 'xterm';
	import { GRPCService } from '$lib/services/grpc';
	import { GrpcStatusCode, GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
	import { AuthAPIClient } from '@gen/ts/auth/v1alpha1/auth.client';
	import { LoggerAPIClient } from '@gen/ts/logger/v1alpha1/log.client';
	import type { FitAddon } from 'xterm-addon-fit';
	import { token } from '$lib/stores/jwt';
	import { RpcError } from '@protobuf-ts/runtime-rpc';

	// Web3
	let onboarding: MetaMaskOnboarding;
	let accounts: string[] = [];
	let address: string = '';
	let signer: ethers.providers.JsonRpcSigner;
	let provider: ethers.providers.Web3Provider;
	let loginButtonText: string;

	// Terminal
	let terminalElement: HTMLElement;
	let terminal: Terminal;
	let fitAddon: FitAddon;

	// gRPC
	const url = 'http://localhost:9000';
	let authClient: AuthAPIClient;
	let loggerClient: LoggerAPIClient;
	let grpcService: GRPCService;

	// Fields
	let logName: string;

	// Refresh login button
	$: if (accounts && accounts.length > 0) {
		onboarding.stopOnboarding();
		loginButtonText = `Connected with ${address}`;
	} else {
		loginButtonText = 'Login with MetaMask';
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
			accounts = await provider.send('eth_requestAccounts', []);
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

		// Check if already logged in
		signer = provider.getSigner();
		if (signer) {
			accounts = await provider.send('eth_requestAccounts', []);
			address = await signer.getAddress();

			// Do gRPC
			if (!$token) {
				try {
					token.set(await grpcService.registerOrSignIn(address, signer));
					console.log($token);
				} catch (e) {
					// TODO: handle error
					console.log(e);
				}
			} else {
				console.log('already logged in');
			}
		}
	});
</script>

<main class="container">
	<section>
		<h1>Terminal</h1>
		<button on:click={doLogin}>{loginButtonText}</button>

		{#if $token}
			<form on:submit|preventDefault={doReadAndWatch}>
				<label for="logname">
					Log Name
					<input name="logname" bind:value={logName} />
				</label>
				<button>Fetch</button>
			</form>
		{/if}

		<div bind:this={terminalElement} />
	</section>
</main>
