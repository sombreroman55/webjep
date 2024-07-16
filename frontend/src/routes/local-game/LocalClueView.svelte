<script lang="ts">
// TODO: Use svelte store to store status of clues in game
  import type { Clue } from "$lib/Clues";

  export let clue: Clue;
  let showAnswer = false;

  function revealAnswer() {
    showAnswer = true;
  }
</script>

<div on:click={revealAnswer}>
{#if !showAnswer}
  <h3>{clue.question}</h3>
{:else}
  <h3>{clue.question}</h3>
  <h3>{clue.answer}</h3>
{/if}
</div>


<style scoped>
/* TODO: Add styling to this page to handle full screen and not and stuff */
</style>
