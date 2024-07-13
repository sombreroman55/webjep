<script lang="ts">
	import type { Clue } from '$lib/Clues';

	export let clue: Clue;
</script>

<div>
  ${clue.value} Clue
  <label >Question: <input bind:value={clue.question}/></label>
  <label >Answer: <input bind:value={clue.answer}/></label>
</div>
