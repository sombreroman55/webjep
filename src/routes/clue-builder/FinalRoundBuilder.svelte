<script lang="ts">
	import type { FinalRound } from '$lib/Clues';

	export let final: FinalRound;
</script>

<div>
  <h2>Final Round</h2>
  <label>Category: <input name="fr-category" bind:value={final.category}/></label>
  <label>Question: <input name="fr-question" bind:value={final.question}/></label>
  <label>Answer: <input name="fr-answer" bind:value={final.answer}/></label>
</div>
