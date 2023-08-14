<script setup lang="ts">
import { ref } from 'vue';
import client from '../client';
import { Tag } from '../api/pb/api/v1/api_pb';

const user = new Headers().get("X-Forwarded-User")
const title = ref("")
const text = ref("")
const tagName = ref("")

const createSodan = async () => {
    const tag = new Tag()
    tag.name = tagName.value
    const tags = [tag]

    const res = await client.createSodan({
        title: title.value,
        text: text.value,
        createrId: user ? user : "anonymous",
        tags: tags
    })
    console.log(res)
    title.value = ""
    text.value = ""
    tagName.value = ""
    window.alert("作成しました")
}
</script>

<template>
    <div>
        <label for="title">タイトル</label>
        <br>
        <input type="text" id="title" v-model="title">
        <br>
        <label for="text">本文</label>
        <br>
        <textarea id="text" v-model="text"></textarea>
        <br>
        <label for="tags">タグ</label>
        <br>
        <input type="text" id="tags" v-model="tagName">
        <br>
        <button @click="createSodan">作成する</button>
    </div>
</template>