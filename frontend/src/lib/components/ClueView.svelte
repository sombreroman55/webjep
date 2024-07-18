<script lang="ts">
  import type { Clue } from "$lib/Clues";

  export let clue: Clue;
  let showAnswer = false;

  function revealAnswer() {
    showAnswer = true;
    // TODO: Mark clue as spent
  }
</script>

<div on:click={revealAnswer}>
  <p>{clue.question}</p>
  {#if showAnswer}
    <hr />
    <p>{clue.answer}</p>
  {/if}
</div>
