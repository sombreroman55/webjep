<script lang="ts">
	import type { Player } from '$lib/types';

	export let winner: Player;
</script>

<div>
	<b>{winner.name} wins with {winner.score}</b>
</div>

<style>
</style>
