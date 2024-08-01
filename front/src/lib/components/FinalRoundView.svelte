<script lang="ts">
	// import { scale } from 'svelte/transition';
	import type { FinalRound } from '$lib/types';
	import { resetClue } from '$lib/gameStore';
	import '$lib/assets/common.css';

	export let clue: FinalRound;
	let showCategory = true;
	let showAnswer = false;

	function revealAnswer() {
		if (showCategory) {
			showCategory = false;
		} else {
			if (!showAnswer) {
				showAnswer = true;
			} else {
				resetClue();
			}
		}
	}
</script>

<div class="jep-blue" on:click={revealAnswer}>
	{#if showCategory}
		<p>{clue.category}</p>
	{:else}
		<p>{clue.question}</p>
		{#if showAnswer}
			<hr />
			<p>{clue.answer}</p>
		{/if}
	{/if}
</div>

<style>
</style>
