/* eslint-disable @typescript-eslint/consistent-type-imports */

// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
// and what to do when importing types
declare namespace App {
	// interface Error {}
	// interface Locals {}
	// interface PageData {}
	// interface Platform {}
}

interface Window {
	ethereum: import('ethers').providers.ExternalProvider;
}
