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

  const socket = ref(new WebSocket('ws://localhost:8080/echo'))

  socket.value.onopen = () => {
    console.log('WebSocket Client Connected')
  }

  socket.value.onmessage = (message) => {
    console.log(message)
  }

  function send(message: string) {
    socket.value.send(message)
  }
</script>

<template>
  <h1>Lobby</h1>
  <h2>Host: {{ host }}</h2>
  <button @click="send('Hello')">Send</button>
</template>
