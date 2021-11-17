<template>
  <div class="row justify-center">
    <q-page v-show="post.id" padding class="col-xs-12 col-sm-10 col-md-8">
      <div class="q-mb-lg">
        <h4 class="text-cyan-9">{{ post.title }}</h4>
        <code class="text-italic">
          Updated by {{ $store.getters.username }} {{ post.updated_at }}
        </code>
      </div>
      <div class="q-mt-lg">
        <v-md-preview :text="post.body_html"></v-md-preview>
      </div>
      <div>
        <q-chip
          v-for="label in post.labels"
          outline
          square
          clickable
          class="label"
          :key="label.index"
          :style="`border-color: #${label.color}; color: #${label.color};`"
          @click="chipClickHandler(label.name)"
        >
          {{ label.name }}
        </q-chip>
      </div>
      <q-separator color="cyan-9" style="height: 1px;" />
      <Comment />
    </q-page>
  </div>
</template>
<script lang="ts">
import { getCurrentInstance, reactive, onBeforeMount, toRefs } from 'vue'
import Comment from '../../components/CommentList.vue';
import { getCart1 } from '../../api/test/index'
export default {
  name: 'Post',
  components: { Comment },
  setup () {
    let data = reactive({
      post: {
        'id':'123',
        'title':'为什么你应该停止阅读新闻？',
        'updated_at':'123',
        'body_html':''
      },
    })
    const {ctx} = getCurrentInstance() as any
    const method = {
      async commentss():Promise<any> {
        let aaa  = await getCart1() as any
        data.post.body_html = aaa.msg
        console.log(ctx)
      }
    }
     onBeforeMount(async()=>{
        await method.commentss()
    })
    return {
      ...toRefs(data),
      ...method
    }
  }
};
</script>

