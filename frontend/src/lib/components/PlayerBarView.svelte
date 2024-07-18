<script lang="ts">
	import type Player from '$lib/Player';
	import Podium from './Podium.svelte';

	export let players: Player[];
</script>

<div class="player-bar">
	{#each players as player}
		<Podium {player} />
	{/each}
</div>

<style>
	.player-bar {
		display: grid;
		grid-auto-columns: minmax(0, 1fr);
		grid-auto-flow: column;
	}
</style>
