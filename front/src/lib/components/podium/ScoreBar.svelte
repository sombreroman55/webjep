<script lang="ts">
	import { currentClue } from '$lib/gameStore';
	import type { Player } from '$lib/types';
	import '$lib/assets/common.css';

	export let player: Player;

	$: scorePrefix = player.score < 0 ? '-$' : '$';
	$: scoreStyle = player.score < 0 ? 'color: red' : 'color: white';

	function updateScore(correct: boolean) {
		player.score += $currentClue.value! * (correct ? 1 : -1);
	}
</script>

<div class="jep-blue">
	<button on:click={() => updateScore(false)}>-</button>
	<span style={scoreStyle} class="score">
		{scorePrefix}{Math.abs(player.score).toLocaleString()}
	</span>
	<button on:click={() => updateScore(true)}>+</button>
</div>

<style>
	.score {
		font-family: sans-serif;
	}
</style>
