<script setup lang="ts">
  function postClues() {
    const url = "http://localhost:8080/api/clues"
    const form = document.querySelector('form');
    const formData = new FormData(form);
    const request = new Request(url, {
      method: 'POST',
      body: formData,
    });
    fetch(request)
      .then((response) => {
        if (response.status === 200) {
          alert('Clues uploaded successfully!');
        } else {
          alert('Something went wrong. Please try again.');
        }
      })
      .catch((error) => {
        console.error(error);
      });
  }
</script>

<template>
  <div>
    <form @submit.prevent="postClues">
      <label for="clues">Upload Clues</label>
      <input type="file" name="clues" accept=".json" />
      <button type="submit">Submit</button>
    </form>
  </div>
</template>
