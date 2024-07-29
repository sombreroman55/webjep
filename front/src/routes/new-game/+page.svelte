<script lang="ts">
	import { redirect } from '@sveltejs/kit';

	async function sendClues() {
		const files = (<HTMLInputElement>document.getElementById('onlineLoad')).files;
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
      var formData = new FormData();
      formData.append("clues", text);
		const response = await fetch("/new-game", { method: 'post', body: formData });
		console.log(response);
	}
</script>

<div>
	<nav>
		<a href="/">home</a>
		<a href="/clue-builder">clue builder</a>
		<a href="/local-game">new local game</a>
		<a href="/game">new online game</a>
	</nav>

	<h1>online game</h1>
	<div class="tutorial">
		online games connect you with friends over the internet to play a round of jep. by creating a
		game, you are the host of that game. uploading clues will create a game on the server and send
		you a link to give to friends. podiums have video chat capability to join and play in a video
		chat setting, all built into webjep. as host, you will be able to adjust player scores, though
		the server will determine the winner of the buzzer. upload your clues below to create a room and
		have fun hosting your friends :)
	</div>
	<form method="POST" action="/new-game" on:submit|preventDefault={sendClues}>
		<label>Load clues for game <input type="file" name="onlineLoad" id="onlineLoad" /></label>
		<input type="submit" value="Load clue file" />
	</form>
</div>

<style>
	.tutorial {
		padding: 30px;
	}
</style>
