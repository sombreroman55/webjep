<script lang="ts">
	import type { Game } from '$lib/Clues';
	import GameView from '$lib/components/GameView.svelte';
	// import type Player from '$lib/Player';

	let game: Game | null = null;
	// let players: Player[] = [{ name: 'test', score: 123 }];

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
		console.log(game);
	}
</script>

<div>
	{#if !game}
		<label>Load clues for game <input type="file" name="localLoad" id="localLoad" /></label>
		<button id="localLoadButton" on:click={loadClues}>Load clue file</button>
	{:else}
		<GameView {game} />
	{/if}
</div>
