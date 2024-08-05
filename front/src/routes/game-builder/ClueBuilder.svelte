<script lang="ts">
	import type { Clue } from '$lib/types';

	export let roundNumber: number;
	export let categoryNumber: number;
	export let clueNumber: number;
	export let clue: Clue;
</script>

<div>
	Clue {clueNumber + 1}
	<label>
		Question: <input
			name="r{roundNumber}-c{categoryNumber}-l{clueNumber}-question"
			bind:value={clue.question}
		/></label
	>
	<label>
		Answer: <input
			name="r{roundNumber}-c{categoryNumber}-l{clueNumber}-answer"
			bind:value={clue.answer}
		/></label
	>
</div>
