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
    <div class="goly-wrapper" v-if="show">
        <p>url: <a :href="`http://localhost:3000/r/${props.goly.goly}`">http://localhost:3000/r/{{ props.goly.goly }}</a></p>
        <p>redirect: {{ props.goly.redirect }}</p>
        <p>clicked: {{ props.goly.clicked }}</p>
    <button @click="deleteGoly">Delete</button>
    </div>
</template>

<style scoped>
.goly-wrapper {
    background-color: rgb(46, 45, 45);
    display: flex;
    padding: 1em;
    border-radius: 4px;
    flex-direction: column;
    gap: 5px;
    margin-top: 10px;
    max-width: 580px;
}

p {
    margin: unset;
}
</style>