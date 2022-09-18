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

  const url = ref()
  const random = ref()
  const golies = ref<Goly[]>([])
  const checked = ref(false)

  const getGolies = async() => {
    const res = await fetch('http://localhost:3000/goly')

    res.json().then(json => {
      golies.value = json
    })
  }

  async function createGoly(data: any) {
        const json = {
            redirect: url.value,
            goly: random.value,
            random: checked.value,
        }

        await fetch('http://localhost:3000/goly', {
            method: 'POST',
            headers: {"Content-Type": "application/json"},
            body: JSON.stringify(json),   
        }).then(async(res) => {
            await getGolies()
        })
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
        v-model="random">
      <input type="checkbox" v-model="checked">
      <button @click="createGoly">Add</button>
  </div>
  <div v-for="goly in golies" :key="goly.id">
    <Goly :goly="goly"/>
  </div>
</template>

<style scoped>
.modal-wrapper {
  display: flex;
  gap: 1em;
  align-items: center;
}

.modal-input {
  padding: 0.75em 1.2em;
}

button {
  padding: 0.75em 1.2em;
}
</style>
