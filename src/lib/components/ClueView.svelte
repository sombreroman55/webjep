<script lang="ts">
import { scale } from 'svelte/transition';
	import type { Clue } from '$lib/Clues';
	import { resetClue } from '$lib/gameStore';
    import "$lib/assets/common.css";

	export let clue: Clue;
	let showAnswer = false;

	function revealAnswer() {
		if (!showAnswer) {
			showAnswer = true;
		} else {
            clue.value = 0;
			resetClue();
		}
	}
</script>

<div transition:scale={{duration: 750}} class="jep-blue" on:click={revealAnswer}>
	<p>{clue.question}</p>
	{#if showAnswer}
		<hr />
		<p>{clue.answer}</p>
	{/if}
</div>

<style>
</style>
