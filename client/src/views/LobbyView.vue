<script setup lang="ts">
  import { ref, onMounted } from 'vue'
  const host = ref('')

  onMounted(async () => {
    await getHost()
  })

  async function getHost() {
    const response = await fetch('http://localhost:8080/api/host')
    const data = await response.json()
    host.value = data.host
  }

  const socket = ref(new WebSocket('ws://localhost:8080'))
</script>

<template>
  <h1>Lobby</h1>
  <h2>Host: {{ host }}</h2>
</template>
