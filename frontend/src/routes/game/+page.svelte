<script lang="ts">
	import type { Game } from '$lib/Clues';
	import GameView from '$lib/components/GameView.svelte';
	// import type Player from '$lib/Player';

	let game: Game | null = null;
	// let payers: Player[] = [{ name: 'test', score: 123 }];

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
		<h1>online game</h1>
		<div class="tutorial">
			online games connect you with friends over the internet to play a round of jep. by creating a
			game, you are the host of that game. uploading clues will create a game on the server and send
			you a link to give to friends. podiums have video chat capability to join and play in a video
			chat setting, all built into webjep. as host, you will be able to adjust player scores, though
			the server will determine the winner of the buzzer. upload your clues below to create a room
			and have fun hosting your friends :)
		</div>
		<label>Load clues for game <input type="file" name="localLoad" id="localLoad" /></label>
		<button id="localLoadButton" on:click={loadClues}>Load clue file</button>
	{:else}
		<GameView {game} />
	{/if}
</div>

<style>
	.tutorial {
		padding: 30px;
	}
</style>
