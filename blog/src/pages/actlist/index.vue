<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12" style="width: 750px">
      <div
        v-if="$q.screen.gt.sm"
        style="position: fixed; margin-left: -5rem; top: 140px; z-index: 2"
      >
        <div>
          <q-btn
            style="margin-bottom: 1.667rem"
            round
            color="primary"
            @click="favorite()"
            icon="thumb_up_alt"
          >
            <q-badge
              :color="post.FavourArr.length == 0 ? '#ccc' : 'orange'"
              rounded
              floating
              >{{ post.favour }}</q-badge
            >
          </q-btn>
        </div>
      </div>
      <div v-if="onload" class="q-mb-lg">
        <h4 class="text-cyan-9" style="margin-top: 10px">{{ post.title }}</h4>
        <q-item>
          <q-item-section avatar>
            <q-btn @click="toRepositories(post.profile.id)" round>
              <q-avatar>
                <img :src="post.profile.avatar" />
              </q-avatar>
            </q-btn>
          </q-item-section>
          <q-item-section>
            <q-item-label caption>{{ post.profile.name }}</q-item-label>
            <q-item-label>
              <code class="text-italic">
                Updated by {{ $store.getters.username }} {{ post.CreatedAt }}
              </code>
              <q-icon style="padding-left: 15px" name="visibility"/>
                {{ post.viewCount }}
            </q-item-label>
          </q-item-section>
        </q-item>
      </div>
      <div v-else class="q-mb-lg">
        <q-skeleton width="50%" />
        <q-skeleton type="text" width="30%" />
      </div>
      <q-scroll-observer @scroll="onScroll" />
      <v-md-preview :text="post.content" ref="previewRef"></v-md-preview>

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
      <q-separator color="cyan-9" style="height: 1px" />
      <Comment :idss="post.id" />
    </q-page>
    <q-page v-if="$q.screen.gt.sm" style="width: 250px; padding: 16px">
      <q-card
        style="position: fixed; width: 218px; top: inherit; z-index: 2"
        flat
        bordered
        class="my-card"
      >
        <q-card-section>
          <div class="text-h6">目录</div>
        </q-card-section>
        <q-separator inset />
        <q-card-section ref="titleRef">
          <div
            v-for="(anchor, item) in titles"
            :class="`line${anchor.lineIndex}`"
            :key="item"
            :style="{ padding: `10px 0 10px ${anchor.indent * 20}px` }"
            @click="handleAnchorClick(anchor)"
          >
            <a style="cursor: pointer">{{ anchor.title }}</a>
          </div>
        </q-card-section>
      </q-card>
    </q-page>
  </div>
</template>
<script lang="ts">
import {
  getCurrentInstance,
  reactive,
  ref,
  onBeforeMount,
  toRefs,
  onMounted,
  nextTick,
} from 'vue';
import Comment from '../../components/CommentList.vue';
import imagesLoaded from 'imagesloaded';
import { getArticle, favour, unfavour } from '../../api/test/index';
import { useRoute, useRouter } from 'vue-router';
import { date } from 'quasar';
export default {
  name: 'Post',
  components: { Comment },
  setup() {
    const route = useRoute() as any;
    const router = useRouter() as any;
    let data = reactive({
      onload: false,
      post: {
        id: '',
        title: '',
        CreatedAt: '',
        viewCount: '',
        content: '',
        profile: null,
        favour: 0,
        FavourArr: [],
      },
      titles: [],
      isActive: true,
      closetitle: null,

    });
    const previewRef = ref(null);
    const titleRef = ref(null);
    // const {ctx} = getCurrentInstance() as any
    const method = {
      async commentss(): Promise<any> {
        let datas = (await getArticle(route.params.id)) as any;
        data.onload = true;
        data.post.content = datas.data.content;
        data.post.id = datas.data.ID;
        data.post.title = datas.data.title;
        let timeStamp = new Date(datas.data.CreatedAt);
        let formattedString = date.formatDate(timeStamp, 'YYYY-MM-DD HH:mm:ss');
        data.post.CreatedAt = formattedString;
        data.post.viewCount = datas.data.view_count;
        data.post.profile = datas.data.Profile;
        data.post.favour = datas.data.favour;
        data.post.FavourArr = datas.data.FavourArr;
      },
      toRepositories(id: number): void {
        router.push(`/repository/${id}`);
      },
      handleAnchorClick(anchor) {
        const { lineIndex } = anchor;
        const heading = previewRef.value.$el.querySelector(
          `[data-v-md-line="${String(lineIndex)}"]`
        );

        if (heading) {
          previewRef.value.scrollToTarget({
            target: heading,
            scrollContainer: window,
            top: 60,
          });
        }
      },
      processTitles() {
        const anchors = previewRef.value.$el.querySelectorAll(
          'h1,h2,h3,h4,h5,h6'
        );
        const titles = Array.from(anchors).filter(
          (title) => !!(title as HTMLElement).innerText.trim()
        );
        if (!titles.length) {
          data.titles = [];
          return;
        }
        const hTags = Array.from(
          new Set(titles.map((title) => (title as HTMLElement).tagName))
        ).sort();
        data.titles = titles.map((el) => ({
          title: (el as HTMLElement).innerText,
          lineIndex: (el as HTMLElement).getAttribute('data-v-md-line'),
          indent: hTags.indexOf((el as HTMLElement).tagName),
          top: (el as HTMLElement).offsetTop,
        }));
        console.log(data.titles);
      },
      async favorite(): Promise<any> {
        const params = {
          article_id: parseInt(data.post.id),
        };
        if (data.post.FavourArr.length == 0) {
          let res = (await favour(params)) as any;
          if (res.status == 200) {
            data.post.FavourArr[0] = true;
            data.post.favour++;    
          }
        } else {  
          let res = (await unfavour(params)) as any;
          if (res.status == 200) {
            data.post.FavourArr = [];
            data.post.favour--;
          }
        }
      },
    };
    onMounted(async () => {
      await method.commentss();
      await nextTick(() => {
        let loadedCount = 0;
        const images = previewRef.value.$el.querySelectorAll('img');
        console.log(images);
        images.forEach((img) => {
          if (img.complete) {
            // 图片已经加载
            loadedCount++;
            if (loadedCount === images.length) {
              // 所有图片都已加载
              method.processTitles();
            }
          } else {
            // 图片未加载，添加事件监听器
            img.addEventListener('load', () => {
              loadedCount++;
              if (loadedCount === images.length) {
                // 所有图片都已加载
                method.processTitles();
              }
            });

            img.addEventListener('error', () => {
              console.error('图片加载失败:', img.src);
              // 即使有图片加载失败，也继续计算高度
              loadedCount++;
              if (loadedCount === images.length) {
                method.processTitles();
              }
            });
          }
        });

        if (!images.length) {
          // 如果没有图片，直接计算高度
          method.processTitles();
        }
      });
    });
    return {
      ...toRefs(data),
      ...method,
      previewRef,
      titleRef,
      onScroll(info) {
        var ftop = Number(info.position.top) + 12;
        data.titles.forEach((element) => {
          const title = titleRef.value.$el.querySelector(
            `.line${String(element.lineIndex)}`
          );
          title.style.backgroundColor = '';
          if (ftop >= element.top) {
            data.closetitle = element;
          }
        });
        if (data.closetitle) {
          const activeTitleElement = titleRef.value.$el.querySelector(
            `.line${String(data.closetitle.lineIndex)}`
          );
          activeTitleElement.style.backgroundColor = '#d2eaf9';
        }
      },
    };
  },
};
</script>

<style lang="sass">
blockquote
  margin: 0 10px
.github-markdown-body
  padding: 0 10px
</style>
