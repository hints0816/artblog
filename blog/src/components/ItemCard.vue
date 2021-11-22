<template>
  <div>
    <div v-for="post in postList" :key="post.index">
      <q-item clickable v-ripple class="q-mt-lg">
        <q-item-section @click="toPostDetail(post.ID)">
          <q-item-label>
            <div class="text-h6 text-cyan-9">{{ post.title }}</div>
          </q-item-label>
          <q-item-label lines="4" caption class="text-body1 text-gray-light text-justify">
            {{ post.desc }}
          </q-item-label>
        </q-item-section>
        <q-item-section side top>
            <q-item-label caption>
                {{ post.UpdatedAt }}
            </q-item-label>
          <q-chip
            v-for="label in post.labels"
            clickable
            class="label"
            :color="label.color" 
            :text-color="label.textcolor" 
            :icon="label.icon"
            :key="label.index"
            @click="chipClickHandler(label.name)"
          >
            {{ label.name }}
          </q-chip>
        </q-item-section>
      </q-item>
        <!-- <q-card-actions align="left">
            <q-btn flat color="red" icon="favorite" />
            <q-btn flat color="teal" icon="comment" />
        </q-card-actions> -->
      <q-separator spaced inset />
    </div>
  </div>
</template>

<script lang="ts">
import { getCurrentInstance } from 'vue'
import { useRouter } from 'vue-router'

export default {
  name: 'Item',
  props: {
    postList: {
      type: Array,
      required: true,
    },
  },
  setup () {
    const router = useRouter() as any
    const {ctx} = getCurrentInstance() as any
    const method = {
      toPostDetail(id: number):void {
        router.push(`/posts/${id}`);
        console.log(ctx)
      }
    }
    return {
      ...method
    }
  },
  methods: {
  
  },
};
</script>

<style lang="scss" scoped>
.markdown-body hr{
    height: unset
}
</style>