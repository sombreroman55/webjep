<script setup lang="ts">
  import { ref, onMounted, onUnmounted } from 'vue'
  import { useJepSocketStore } from '@/stores/jepSocket';
  const host = ref('')
  const { connect, send, disconnect } = useJepSocketStore()

  onMounted(async () => {
    await getHost()
    connect();
  });

  onUnmounted(() => {
    disconnect();
  });

  async function getHost() {
    const response = await fetch('http://localhost:8080/api/host')
    const data = await response.json()
    host.value = data.host
  }
</script>

<template>
  <h1>Lobby</h1>
  <h2>Host: {{ host }}</h2>
  <button @click="send('Hello')">Send</button>
</template>
