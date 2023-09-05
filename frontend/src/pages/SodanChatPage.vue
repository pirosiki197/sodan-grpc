<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useClient } from '../client';
import { Reply, Sodan } from '../api/pb/api/v1/api_pb';
import { APIService } from '../api/pb/api/v1/api_connect';

const header = new Headers()

const props = defineProps<{ id: string }>()
const sodan = ref<Sodan>()
const replies = ref<Reply[]>([])
const replyText = ref<string>("")
const client = useClient(APIService)

const sendReply = async () => {
    const user = header.get("X-Showcase-User")
    await client.createReply({
        sodanId: BigInt(props.id),
        text: replyText.value,
        createrId: user ? user : "guest",
    })
    replyText.value = ""
}

onMounted(async () => {
    const sodanRes = await client.getSodan({ id: BigInt(props.id) })
    sodan.value = sodanRes.sodan

    const repliesRes = await client.getReplies({ sodanId: BigInt(props.id) })
    replies.value = repliesRes.replies

    for await (const res of client.subscribeSodan({ id: BigInt(props.id) })){
        console.log(res.reply)
        if (res.reply) {
            replies.value.push(res.reply)
            console.log(res.reply)
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
                <h3 v-if="sodan?.tags.length">タグ</h3>
                <li v-for="tag in sodan?.tags" :key="tag.name">
                    {{ tag.name }}
                </li>
            </ul>
        </div>
        <div>
            <ul>
                <h2>返信</h2>
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