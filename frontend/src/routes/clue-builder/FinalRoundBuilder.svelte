<script lang="ts">
	import type { FinalRound } from '$lib/Clues';

	export let final: FinalRound;
</script>

<div>
  <h2>Final Round</h2>
  <label >Category: <input bind:value={final.category}/></label>
  <label >Question: <input bind:value={final.question}/></label>
  <label >Answer: <input bind:value={final.answer}/></label>
</div>
