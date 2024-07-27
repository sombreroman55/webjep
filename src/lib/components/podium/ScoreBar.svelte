<script lang="ts">
	import type Player from '$lib/Player';

	export let player: Player;
</script>

<div>
	<button>-</button>
	{player.score}
	<button>+</button>
</div>
