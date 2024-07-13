<script lang="ts">
	import GameBuilder from './GameBuilder.svelte';
	import type { Snapshot } from './$types';
	import type { Game } from '$lib/Clues';
	import { newGame } from '$lib/Clues';

    let game: Game = newGame();

	export const snapshot: Snapshot<Game> = {
      capture: () => game,
      restore: (g) => (game = g),
    };
</script>

<GameBuilder game={game}/>
