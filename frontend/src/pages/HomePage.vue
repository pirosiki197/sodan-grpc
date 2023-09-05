<script setup lang="ts">
import { APIService } from '../api/pb/api/v1/api_connect';
import { Sodan } from '../api/pb/api/v1/api_pb';
import { useClient } from '../client'
import { onMounted, ref } from 'vue';

const sodans = ref<Sodan[]>([])
const client = useClient(APIService)

onMounted(async () => {
    const res = await client.getSodanList({});
    sodans.value = res.sodans;
    console.log(sodans.value)
})

const findByTag = async (tag: string) => {
    const res = await client.getSodansByTag({ tagName: tag})
    sodans.value = res.sodans
}
</script>

<template>
    <h1>最新の質問</h1>
    <div>
        <ul>
            <li v-for="sodan in sodans" :key="sodan.id.toString">
                <router-link :to="`/sodan/${sodan.id.toString}`">{{ sodan.title }}</router-link>
                <p v-if="sodan.tags.length">タグ</p>
                <ul>
                    <li v-for="tag in sodan.tags" :key="tag.name">
                        <button @click="findByTag(tag.name)">{{ tag.name }}</button>
                    </li>
                </ul>
            </li>
        </ul>
    </div>
</template>