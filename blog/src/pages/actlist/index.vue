<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12" style="width:750px">
      <div v-if="$q.screen.gt.sm" style="position: fixed;margin-left: -5rem;top: 140px;z-index: 2;">
        <div><q-btn style="margin-bottom: 1.667rem" round color="primary" icon="thumb_up_alt" /></div>
        <div><q-btn style="margin-bottom: 1.667rem" round color="primary" icon="textsms" /></div>
      </div>
      <div v-if="onload" class="q-mb-lg">
        <h4 class="text-cyan-9">{{ post.title }}</h4>
        <q-item>
          <q-item-section avatar>
            <q-btn @click="toRepositories(post.profile.id)" round>
              <q-avatar>
                <img :src="post.profile.avatar">
              </q-avatar>
            </q-btn>
          </q-item-section>
          <q-item-section>
            <q-item-label caption>{{ post.profile.name }}</q-item-label>
            <q-item-label>
              <code class="text-italic">
                Updated by {{ $store.getters.username }} {{ post.CreatedAt }}
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
        <v-md-preview :text="post.content"></v-md-preview>
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
    <q-page v-if="$q.screen.gt.sm" padding style="width:250px;background: antiquewhite">
    </q-page>
  </div>
</template>
<script lang="ts">
import { getCurrentInstance, reactive, onBeforeMount, toRefs } from 'vue'
import Comment from '../../components/CommentList.vue';
import { getArticle } from '../../api/test/index'
import { useRoute, useRouter } from 'vue-router'
import { date } from 'quasar'
export default {
  name: 'Post',
  components: { Comment },
  setup () {
    const route = useRoute() as any
    const router = useRouter() as any;
    let data = reactive({
      onload: false,
      post: {
        id:'',
        title:'',
        CreatedAt:'',
        content:'',
        profile: null
      },
    })
    const {ctx} = getCurrentInstance() as any
    const method = {
      async commentss():Promise<any> {
        let datas  = await getArticle(route.params.id) as any
        data.onload = true
        data.post.id = datas.data.ID
        data.post.title = datas.data.title
        let timeStamp = new Date(datas.data.CreatedAt)
        let formattedString = date.formatDate(timeStamp, 'YYYY-MM-DD HH:mm:ss')
        data.post.CreatedAt = formattedString
        data.post.content = datas.data.content
        data.post.profile = datas.data.Profile
      },
      toRepositories(id: number): void {
        router.push(`/repository/${id}`)
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

<style lang="sass">
blockquote
  margin: 0 10px
.github-markdown-body
  padding: 0 10px
</style>
