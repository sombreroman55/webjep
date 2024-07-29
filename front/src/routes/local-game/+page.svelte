<script lang="ts">
	import type { Game } from '$lib/Clues';
	import GameView from '$lib/components/GameView.svelte';
	import type Player from '$lib/Player';

	let game: Game | null = null;
	let ready: boolean = false;
	let players: Player[] = [];

	function addPlayer() {
		let playerInput = <HTMLInputElement>document.getElementById('newPlayerName');
		let playerName = playerInput.value;
		let newPlayer: Player = { name: playerName, score: 0, isHost: false };
		players = [...players, newPlayer];
		playerInput.value = '';
	}

	async function loadClues() {
		const files = (<HTMLInputElement>document.getElementById('localLoad')).files;
		if (!files || files.length == 0) {
			// TODO: Do something
			return;
		}
		const selectedFile = files[0];
		if (!selectedFile) {
			// TODO: Do something
			return;
		}
		const text = await selectedFile.text();
		// TODO: try-catch this
		const gameObj = JSON.parse(text);
		// TODO: Validate clues
		game = gameObj as Game;
	}
</script>

<div>
	{#if !game}
		<nav>
			<a href="/">home</a>
			<a href="/clue-builder">clue builder</a>
			<a href="/local-game">new local game</a>
			<a href="/game">new online game</a>
		</nav>

		<h1>local game</h1>
		<div class="tutorial">
			local games are completely hosted in your browser. no data gets sent to the server, not even
			the clues you load below. there is no need for videos on the podium or for a host podium. this
			is designed for games where you have everyone in the room and just need a good layout to show
			the clues and the scores. the game will randomly assign daily double clues automatically in
			each round. this is the best way to play. load your clues below to get started. have fun :)
		</div>
		<label>Load clues for game <input type="file" name="localLoad" id="localLoad" /></label>
		<button id="localLoadButton" on:click={loadClues}>Load clue file</button>
	{:else if !ready}
		<h1>local game</h1>
		<div class="tutorial">add players to your game</div>
		<div>
			current players: {players.map((p) => p.name).join(',')}
		</div>
		<label>name: <input name="newPlayerName" id="newPlayerName" /></label>
		<button on:click={addPlayer}>add player</button>
		<div>
			<button on:click={() => (ready = true)}>start game</button>
		</div>
	{:else}
		<GameView {game} {players} />
	{/if}
</div>

<style>
	.tutorial {
		padding: 30px;
	}
</style>
