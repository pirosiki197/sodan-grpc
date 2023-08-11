<script setup lang="ts">
import { createPromiseClient } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";
import { APIService } from "./api/pb/api/v1/api_connect";
import { Sodan,Reply } from "./api/pb/api/v1/api_pb";
import { onMounted, ref } from "vue";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8080",
})

const client = createPromiseClient(APIService, transport)

const header = new Headers()

const sodan = ref<Sodan>()
const replies = ref<Reply[]>([])
const userID = ref<string | null>()
const text = ref("")

const SendReply = async () => {
  console.log(text.value)
  await client.createReply({
    sodanId: BigInt(1),
    text: text.value,
    createrId: userID.value ? userID.value : "pirosiki"
  })
  text.value = ""
}

onMounted(async () => {
  userID.value = header.get("X-Forwarded-User")

  var sodanRes = await client.getSodan({id: BigInt(1)})
  sodan.value = sodanRes.sodan
  var repliesRes = await client.getReplies({sodanId: BigInt(1)})
  replies.value = repliesRes.replies

  for await (const res of client.subscribeSodan({id: BigInt(1)})) {
    console.log(res)
    replies.value.push(res.reply ? res.reply : new Reply())
  }
})

</script>

<template>
  <p>こんにちは</p>
  <h1>{{ sodan?.title }}</h1>
  <p>作成者: {{ sodan?.createrId }}</p>
  <p>{{ sodan?.text }}</p>
  <div>
    <h2>返信</h2>
    <ul>
      <li v-for="reply in replies" :key="reply.id.toString">
        <p>{{ reply.createrId }}</p>
        <p>{{ reply.text }}</p>
      </li>
    </ul>
  </div>
  <div>
    <input type="text" v-model="text">
    <button @click="SendReply">返信</button>
  </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
