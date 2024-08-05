<script lang="ts">
	import type { Player } from '$lib/types';

	// TODO: Figure this out at some point
	export let player: Player;

	async function enableMediaForPlayer() {
		player.mediaEnabled = true;
		try {
			const constraints = { video: true, audio: true };
			const stream = await navigator.mediaDevices.getUserMedia(constraints);
			const videoElement = document.querySelector(`video#video-stream-${player.name}`);
			videoElement.srcObject = stream;
		} catch (error) {
			console.error('Error opening video camera.', error);
		}
	}
</script>

<div>
	{#if player.mediaEnabled}
		<video class="video-stream" id="video-stream-{player.name}" autoplay playsinline></video>
	{:else}
		<button on:click={enableMediaForPlayer}>Enable media</button>
	{/if}
</div>

<style>
	.video-stream {
	}
</style>
