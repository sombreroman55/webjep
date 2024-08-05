<script lang="ts">
	import type { Category, Clue, Round, Session, Player } from '$lib/types';
	import GameView from '$lib/components/GameView.svelte';
	import { currentGame } from '$lib/gameStore';

	let session: Session | null = null;
	let ready: boolean = false;
	let players: Player[] = [];

	function addPlayer() {
		let playerInput = <HTMLInputElement>document.getElementById('newPlayerName');
		let playerName = playerInput.value;
		let newPlayer: Player = { name: playerName, score: 0, isHost: false, mediaEnabled: false };
		players = [...players, newPlayer];
		playerInput.value = '';
	}

	async function loadClues() {
		const files = (<HTMLInputElement>document.getElementById('localLoad')).files;
		if (!files || files.length == 0) {
			return false;
		}
		const selectedFile = files[0];
		if (!selectedFile) {
			return false;
		}
		const text = await selectedFile.text();
		try {
			const sessionObj = JSON.parse(text) as Session;
			const gameObj = sessionObj.game;

			for (var i = 0; i < gameObj.rounds.length; i++) {
				var r: Round = gameObj.rounds[i];
				for (var j = 0; j < r.categories.length; j++) {
					var c: Category = r.categories[j];
					for (var k = 0; k < c.clues.length; k++) {
						var q: Clue = c.clues[k];
						var bcv = sessionObj.settings.baseClueValue;
						var irm = sessionObj.settings.interRoundMultiplier;
						var ici = sessionObj.settings.interClueInterval;
						var crbv = Math.floor(bcv + i * bcv * irm);
						var ccv = Math.floor(crbv + k * crbv * ici);
						q.value = ccv;
						q.answered = false;
					}
				}
			}
			session = sessionObj;
			currentGame.set(session.game);
		} catch {
			console.error('nope');
		}
	}
</script>

<div>
	{#if !session}
		<nav>
			<a href="/">home</a>
			<a href="/game-builder">game builder</a>
			<a href="/local-game">new local game</a>
			<a href="/new-game">new online game</a>
		</nav>

		<h1>local game</h1>
		<div class="tutorial">
			local games are completely hosted in your browser. no data gets sent to the server, not even
			the clues you load below. there is no need for videos on the podium or for a host podium. this
			is designed for games where you have everyone in the room and just need a good layout to show
			the clues and the scores. the game will randomly assign daily double clues automatically in
			each round. this is the best way to play. load your clues below to get started. have fun :)
		</div>
		<label>Load clues for game <input type="file" name="localLoad" id="localLoad" required /></label
		>
		<button id="localLoadButton" on:click={loadClues}>Load clue file</button>
	{:else if !ready}
		<h1>local game</h1>
		<div class="tutorial">add players to your game</div>
		<div>
			current players: {players.map((p) => p.name).join(',')}
		</div>
		<form on:submit|preventDefault={addPlayer}>
			<label>name: <input name="newPlayerName" id="newPlayerName" /></label>
			<input type="submit" value="add player" />
		</form>
		<div>
			<button on:click={() => (ready = true)}>start game</button>
		</div>
	{:else}
		<GameView game={session.game} {players} />
	{/if}
</div>

<style>
	.tutorial {
		padding: 30px;
	}
</style>
