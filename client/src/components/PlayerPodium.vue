<script setup lang="ts">
import { ref } from "vue";
const timeToAnswer = ref(5);
let timerId = null;

function startCountdown() {
  timerId = setInterval(() => {
    if (timeToAnswer.value > 0) {
      timeToAnswer.value--;
    } else {
      clearInterval(timerId);
    }
  }, 1000);
}

function buzzedIn() {
  startCountdown();
}

function correctAnswer() {
  // TODO: Update score
}

function incorrectAnswer() {
  // TODO: Update score
}
</script>

<template>
  <div class="player-podium" @keydown.space="buzzedIn">
  <div class="timer-bar">
    <div>
    </div>
  </div>
  </div>
</template>

<style></style>
