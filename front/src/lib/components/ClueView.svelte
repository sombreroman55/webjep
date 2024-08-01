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

<div class="clue box-center-v shadow jep-blue" on:click={revealAnswer}>
	<p>{clue.question.toUpperCase()}</p>
	{#if showAnswer}
		<hr />
		<p>{clue.answer.toUpperCase()}</p>
	{/if}
</div>

<style>
  @font-face {
    font-family: 'itc';
    font-style: normal;
    font-weight: 700;
    src: url('/fonts/korinna/itc.woff') format('woff');
  }

  @font-face {
    font-family: 'itc-std';
    font-style: normal;
    font-weight: 700;
    src: url('/fonts/korinna/itc-std.woff') format('woff');
  }

  @font-face {
    font-family: 'itc-lt';
    font-style: normal;
    font-weight: 700;
    src: url('/fonts/korinna/itc-lt.woff') format('woff');
  }

  hr {
    color: white;
  }

  .clue {
    font-family: 'itc';
    font-size: 72px;
    text-align: center;
    color: white;
    width: 100%;
    height: 100%;
  }
</style>
