<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12" style="width:750px">
      <div v-if="$q.screen.gt.sm" style="position: fixed;margin-left: -5rem;top: 140px;z-index: 2;">
        <div><q-btn style="margin-bottom: 1.667rem" round color="primary" icon="thumb_up_alt" /></div>
        <div><q-btn style="margin-bottom: 1.667rem" round color="primary" icon="textsms" /></div>
      </div>
      <div v-if="onload" class="q-mb-lg">
        <h4 class="text-cyan-9" style="margin-top: 10px;">{{ post.title }}</h4>
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
      <div v-if="onload" class="q-mt-lg container">
        <q-scroll-observer @scroll="onScroll" />
        <v-md-preview :text="post.content" ref="preview"></v-md-preview>
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
    <q-page v-if="$q.screen.gt.sm"  
      style="width:250px;padding:16px;">
      <q-card style="position: fixed; width:218px; top: inherit;z-index: 2;" flat bordered class="my-card">
        <q-card-section>
          <div class="text-h6">目录</div>
        </q-card-section>
        <q-separator inset />
        <q-card-section ref="title">
          <div v-for="(anchor, item) in titles" 
                  :class="`line${anchor.lineIndex}`"
                  :key="item" :style="{ padding: `10px 0 10px ${anchor.indent * 20}px` }"
                  @click="handleAnchorClick(anchor)">
                  <a style="cursor: pointer">{{ anchor.title }}</a>
          </div>
        </q-card-section>
      </q-card>
    </q-page>
  </div>
</template>
<script lang="ts">
import { getCurrentInstance, reactive, onBeforeMount, toRefs, ref ,onMounted } from 'vue'
import Comment from '../../components/CommentList.vue';
import imagesLoaded from 'imagesloaded';
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
      titles: [],
      isActive: true,
      closetitle: null
    })
    const {ctx} = getCurrentInstance() as any
    const method = {
      async commentss():Promise<any> {
        let datas  = await getArticle(route.params.id) as any
        data.onload = true
        data.post.content = datas.data.content
        data.post.id = datas.data.ID
        data.post.title = datas.data.title
        let timeStamp = new Date(datas.data.CreatedAt)
        let formattedString = date.formatDate(timeStamp, 'YYYY-MM-DD HH:mm:ss')
        data.post.CreatedAt = formattedString
        data.post.profile = datas.data.Profile
      },
      toRepositories(id: number): void {
        router.push(`/repository/${id}`)
      },
      handleAnchorClick(anchor) {
        const { preview } = ctx.$refs;
        const { lineIndex } = anchor;
        const heading = preview.$el.querySelector(`[data-v-md-line="${String(lineIndex)}"]`);

        if (heading) {
          preview.scrollToTarget({
            target: heading,
            scrollContainer: window,
            top: 60,
          });
        }
      },
    };
    onMounted(async()=>{
       await method.commentss();
       ctx.$nextTick(() => {
        const anchors =ctx.$refs.preview.$el.querySelectorAll('h1,h2,h3,h4,h5,h6');
        const titles = Array.from(anchors).filter((title) => !!((title as HTMLElement).innerText.trim()));
        if (!titles.length) {
          data.titles = [];
          return;
        }
        const hTags = Array.from(new Set(titles.map((title) => (title as HTMLElement).tagName))).sort();
        imagesLoaded('.container', function() {
          // 所有图片加载完成后的回调函数
          data.titles = titles.map((el) => ({
            title: (el as HTMLElement).innerText,
            lineIndex: (el as HTMLElement).getAttribute('data-v-md-line'),
            indent: hTags.indexOf((el as HTMLElement).tagName),
            top: (el as HTMLElement).offsetTop,
          }));
        });
      });
    });
    return {
      ...toRefs(data),
      ...method,
      
      onScroll (info) {
        var ftop = Number(info.position.top)+12;
        data.titles.forEach((element) => {
          const title = ctx.$refs.title.$el.querySelector(`.line${String(element.lineIndex)}`);
          title.style.backgroundColor = '';
          if(ftop>= element.top){
            data.closetitle = element;
          }
        });
        if (data.closetitle) {
            const activeTitleElement = ctx.$refs.title.$el.querySelector(`.line${String(data.closetitle.lineIndex)}`);
            activeTitleElement.style.backgroundColor = '#d2eaf9';
        }
      }
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
