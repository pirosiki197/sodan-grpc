<script setup lang="ts">
import { onMounted, ref } from 'vue';
import client from '../client';
import { Reply, Sodan } from '../api/pb/api/v1/api_pb';

const header = new Headers()

const props = defineProps<{ id: bigint }>()
const sodan = ref<Sodan>()
const replies = ref<Reply[]>([])
const replyText = ref<string>("")

const sendReply = async () => {
    const user = header.get("X-Forwarded-User")
    await client.createReply({
        sodanId: props.id,
        text: replyText.value,
        createrId: user ? user : "guest",
    })
    replyText.value = ""
}

onMounted(async () => {
    const sodanRes = await client.getSodan({ id: props.id })
    sodan.value = sodanRes.sodan

    const repliesRes = await client.getReplies({ sodanId: props.id })
    replies.value = repliesRes.replies

    for await (const res of client.subscribeSodan({ id: props.id })){
        if (res.reply) {
            replies.value.push(res.reply)
        }
    }
})
</script>

<template>
    <div>
        <div>
            <h1>{{ sodan?.title }}</h1>
            <p>作成者: {{ sodan?.createrId }}</p>
            <p>{{ sodan?.text }}</p>
            <ul>
                <li v-for="tag in sodan?.tags" :key="tag.name">
                    {{ tag.name }}
                </li>
            </ul>
        </div>
        <div>
            <ul>
                <li v-for="reply in replies" :key="reply.id.toString">
                    <p>{{ reply.createrId }}</p>
                    <p>{{ reply.text }}</p>
                </li>
            </ul>
        </div>

        <div>
            <input type="text" v-model="replyText">
            <button @click="sendReply">送信</button>
        </div>
    </div>
</template>