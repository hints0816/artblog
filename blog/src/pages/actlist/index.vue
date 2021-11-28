<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-10 col-md-8">
      <div v-if="onload" class="q-mb-lg">
        <h4 class="text-cyan-9">{{ post.title }}</h4>
        <q-item>
          <q-item-section avatar>
            <q-btn round>
              <q-avatar>
                <img src="https://cdn.quasar.dev/img/avatar2.jpg">
              </q-avatar>
            </q-btn>
          </q-item-section>
          <q-item-section>
            <q-item-label caption>Caption</q-item-label>
            <q-item-label>
              <code class="text-italic">
                Updated by {{ $store.getters.username }} {{ post.updated_at }}
              </code>
            </q-item-label>
          </q-item-section>
        </q-item>
      </div>
      <div v-else class="q-mb-lg">
        <q-skeleton width="50%" />
        <q-skeleton type="text" width="30%" />
      </div>
      <div v-if="onload" class="q-mt-lg">
        <v-md-preview :text="post.body_html"></v-md-preview>
      </div>
      <div v-else class="q-mt-lg">
        <q-skeleton type="text" />
        <q-skeleton type="text" />
        <q-skeleton type="text" />
        <q-skeleton type="text" />
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
      <Comment :idss="post.id"/>
    </q-page>
  </div>
</template>
<script lang="ts">
import { getCurrentInstance, reactive, onBeforeMount, toRefs } from 'vue'
import Comment from '../../components/CommentList.vue';
import { getArticle } from '../../api/test/index'
import { useRoute } from 'vue-router'
import { date } from 'quasar'
export default {
  name: 'Post',
  components: { Comment },
  setup () {
    const route = useRoute() as any
    let data = reactive({
      onload: false,
      post: {
        'id':'',
        'title':'',
        'updated_at':'',
        'body_html':''
      },
    })
    const {ctx} = getCurrentInstance() as any
    const method = {
      async commentss():Promise<any> {
        let datas  = await getArticle(route.params.id) as any
        data.onload = true
        console.log(datas)
        data.post.id = datas.data.ID
        data.post.title = datas.data.title
        let timeStamp = new Date(datas.data.CreatedAt)
        let formattedString = date.formatDate(timeStamp, 'YYYY-MM-DD HH:mm:ss')
        data.post.updated_at = formattedString
        data.post.body_html = datas.data.content
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

