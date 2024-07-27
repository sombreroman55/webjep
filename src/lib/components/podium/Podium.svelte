<script lang="ts">
	import type Player from '$lib/Player';
	import ScoreBar from '$lib/components/podium/ScoreBar.svelte';
    import '$lib/assets/common.css'
	export let player: Player;

    let selected = false;

    // TODO: Write custom CSS transition for when player buzzes in
</script>

<div class="podium">
	<ScoreBar {player} />
    <div class="light-box" class:active={selected} on:click={() => (selected = !selected)}>
	{#if player.videoEnabled}
      <video></video>
	{:else}
      <canvas>{player.name}</canvas>
	{/if}
    </div>
</div>

<style>
.podium {
  border: 1px solid black;
}

.light-box {
  background: #1122DD;
}

.active {
  background: #777777;
}
</style>
