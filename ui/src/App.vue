<script setup lang="ts">
  import { onMounted, ref } from 'vue'
  import Goly from './components/Goly.vue'

  interface Goly {
    id: number
    redirect: string
    goly: string
    clicked: number
    random: string
}

  const url = ref<string>()
  const random = ref<string>()
  const golies = ref<Goly[]>([])
  const checked = ref(true)

  const getGolies = async() => {
    const res = await fetch('http://localhost:3000/goly')

    res.json().then(json => {
      golies.value = json
    })
  }

  async function createGoly() {
        const json = {
            redirect: url.value,
            goly: random.value,
            random: checked.value,
        }

        if (json.redirect) {
          await fetch('http://localhost:3000/goly', {
              method: 'POST',
              headers: {"Content-Type": "application/json"},
              body: JSON.stringify(json),   
          }).then(async() => {
              await getGolies()
          })
      }
    }

  onMounted( () => {
    getGolies()
  })

</script>

<template>
  <h1>URL Shortener</h1>

  <div class="modal-wrapper">
      <input 
        class="modal-input"
        type="text" id="url"
        placeholder="url" 
        v-model="url">
      <input 
        class="modal-input"
        type="text"
        id="random" 
        placeholder="random"
        :disabled="checked"
        v-model="random">
      <input type="checkbox" v-model="checked">
      <button @click="createGoly">Add</button>
  </div>
  <div class="golies-wrapper">
    <div v-for="goly in golies" :key="goly.id">
      <Goly :goly="goly"/>
    </div>
  </div>
</template>

<style scoped>
.modal-wrapper {
  display: flex;
  gap: 1em;
  align-items: center;
}

.golies-wrapper {
  max-height: 500px;
  overflow: auto;
  margin-top: 10px;
}

.modal-input {
  padding: 0.75em 1.2em;
}

button {
  padding: 0.75em 1.2em;
}
</style>
