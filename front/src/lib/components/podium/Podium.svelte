<script lang="ts">
	import type { Player } from '$lib/types';
	import ScoreBar from '$lib/components/podium/ScoreBar.svelte';
	import '$lib/assets/common.css';
	import VideoCanvas from './VideoCanvas.svelte';
	export let player: Player;

	let selected = false;

	// TODO: Write custom CSS transition for when player buzzes in
</script>

<div class="podium">
	<ScoreBar {player} />
	<div class="light-box" class:active={selected} on:click={() => (selected = !selected)}>
		<div>
			<VideoCanvas {player} />
		</div>
	</div>
</div>

<style>
	.podium {
		border: 1px solid black;
	}

	.light-box {
		background: #1122dd;
	}

	.active {
		background: #777777;
	}
</style>
