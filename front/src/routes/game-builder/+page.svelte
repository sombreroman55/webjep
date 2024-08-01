<script lang="ts">
	import { fade } from 'svelte/transition';
	import GameBuilder from './GameBuilder.svelte';
	import type { Game, GameSettings, Session } from '$lib/types';
	import { newGame } from '$lib/utils';

	let numRoundsPerGame = 2;
	let numCategoriesPerRound = 6;
	let numCluesPerCategory = 5;
	let baseClueValue = 200;
	let interClueInterval = 1;
	let interRoundMultiplier = 1;
	let playFinalJep = true;

	let session: Session;

	let showSplash = true;
	let showAdvancedOptions = false;

	function createForm() {
		let settings: GameSettings = {
			playFinalJep: playFinalJep,
			baseClueValue: baseClueValue,
			interClueInterval: interClueInterval,
			interRoundMultiplier: interRoundMultiplier,
			numberOfRounds: numRoundsPerGame,
			categoriesPerRound: numCategoriesPerRound,
			cluesPerCategory: numCluesPerCategory
		};
		let innerGame: Game = newGame(settings);
		session = {
			game: innerGame,
			settings: settings
		};
		showSplash = false;
	}
</script>

{#if showSplash}
	<h1>clue builder</h1>
	<div class="tutorial">
		the clue builder is a simple tool to help you make a set of clues for your own game of jep.
		simply fill out the form then click "Download" to download your clues in JSON format which you
		can then use to start either a <a href="/local-game">local game</a> or an
		<a href="/game">online game</a>. click the button below to get started. alternatively, you can
		download a template file here and fill it out in your favorite text editor.
	</div>
	<div>
		<label>
			<input bind:checked={showAdvancedOptions} type="checkbox" /> Show Advanced Options
		</label>
	</div>
	{#if showAdvancedOptions}
		<div transition:fade>
			<div class="tutorial">
				while this strays from the purity of the traditional jep experience, you can configure the
				game to be the way you want to play it. here you can configure nearly every part of the
				game, from the number of rounds, to the number of daily doubles on the board. you can also
				select from a list of pre-made configurations that change the normal jep ruleset in
				different ways.
			</div>
			<div id="advancedOptions">
				<label>
					Play final round <input bind:checked={playFinalJep} type="checkbox" />
				</label>
				<label>Base Clue Value: <input type="number" bind:value={baseClueValue} /></label>
				<label>Number of rounds: <input type="number" bind:value={numRoundsPerGame} /></label>
				<label
					>Number of categories per round: <input
						type="number"
						bind:value={numCategoriesPerRound}
					/></label
				>
				<label
					>Number of clues per category: <input
						type="number"
						bind:value={numCluesPerCategory}
					/></label
				>
				<label>Inter-clue interval: <input type="number" bind:value={interClueInterval} /></label>
				<label
					>Inter-round interval: <input type="number" bind:value={interRoundMultiplier} /></label
				>
			</div>
		</div>
	{/if}
	<button on:click={createForm}>Create Clue Template</button>
{:else}
	<GameBuilder {session} />
{/if}

<style>
	.tutorial {
		padding: 30px;
	}

	#advancedOptions {
		display: flex;
		flex-direction: column;
	}
</style>
