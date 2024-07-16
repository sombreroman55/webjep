<script lang="ts">
  import type Player from "$lib/Player";

  export let player: Player;
</script>

<div>
  <div>{player.name}</div>
  <div>{player.score}</div>
</div>
