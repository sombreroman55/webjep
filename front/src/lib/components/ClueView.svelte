<script lang="ts">
	// import { scale } from 'svelte/transition';
	import type { Clue } from '$lib/types';
	import { resetClue } from '$lib/gameStore';
	import '$lib/assets/common.css';

	export let clue: Clue;
	let showAnswer = false;

	// TODO: Make custom clue reveal animation

	function revealAnswer() {
		if (!showAnswer) {
			showAnswer = true;
		} else {
			clue.answered = true;
			resetClue();
		}
	}
</script>

<div class="jep-blue" on:click={revealAnswer}>
	<p>{clue.question}</p>
	{#if showAnswer}
		<hr />
		<p>{clue.answer}</p>
	{/if}
</div>

<style>
</style>
