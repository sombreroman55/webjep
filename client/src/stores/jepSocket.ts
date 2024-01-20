import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useJepSocketStore = defineStore('jepSocket', () => {
  const socket = ref<WebSocket | null>(null);
  function connect() {
    socket.value = new WebSocket('ws://localhost:8080/jep');

    socket.value.onopen = function () {
      console.log('Jep WebSocket is connected!');
    };

    socket.value.onmessage = function (event) {
      console.log('Jep Message: ' + event.data);
    };

    socket.value.onclose = function () {
      console.log('Jep WebSocket is closed!');
    };
  }

  function send(message: string) {
    socket.value?.send(message);
  }

  function disconnect() {
    socket.value?.close();
    socket.value = null;
  }

  return { socket, connect, send, disconnect };
})
