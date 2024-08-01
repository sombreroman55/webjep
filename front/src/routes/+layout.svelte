<script lang="ts">
export const prerender = true;
export const ssr = false;
</script>

<slot></slot>

<style>
:global(body) {
  margin: 0px;
  padding: 0px;
}
</style>
