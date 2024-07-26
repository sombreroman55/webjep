<script lang="ts">
	import GameBuilder from './GameBuilder.svelte';
	import type { Snapshot } from './$types';
	import type { Game } from '$lib/Clues';
	import { newGame } from '$lib/Clues';

	let game: Game = newGame();

	export const snapshot: Snapshot<Game> = {
		capture: () => game,
		restore: (g) => (game = g)
	};

	let showSplash = true;
</script>

{#if showSplash}
	<h1>clue builder</h1>
	<div class="tutorial">
		the clue builder is a simple tool to help you make a set of clues for your own game of jep.
		simply fill out the form then click "Download" to download your clues in JSON format which you
		can then use to start either a <a href="/local-game">local game</a> or an
		<a href="/game">online game</a>. click the button below to get started. alternatively,
		you can download a template file here and fill it out in your favorite text editor.
	</div>
	<button on:click={() => (showSplash = false)}>Build clues</button>
{:else}
	<GameBuilder {game} />
{/if}

<style>
	.tutorial {
		padding: 30px;
	}
</style>
