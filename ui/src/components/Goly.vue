<script setup lang="ts">
import { ref } from 'vue'

interface Goly {
    id: number
    redirect: string
    goly: string
    clicked: number
    random: string
}

interface Props {
    goly: Goly
}

const props = defineProps<Props>()

const show = ref(true)

const deleteGoly = async() => {
    await fetch("http://localhost:3000/goly/" + props.goly.id, {
        method: "DELETE"
    }).then(() => {
        show.value = false
    })
}
</script>

<template>
    <div v-if="show">
        id: {{ props.goly.id }}
        url: <a :href="`http://localhost:3000/r/${props.goly.goly}`">http://localhost:3000/r/{{ props.goly.goly }}</a>
        redirect: {{ props.goly.redirect }}
        clicked: {{ props.goly.clicked }}
    <button @click="deleteGoly">Delete</button>
    </div>
</template>

<style scoped>

</style>