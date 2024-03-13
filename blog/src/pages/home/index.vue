<template>
  <div class="row justify-center">
    <q-page padding class="row col-xs-12"  style="width:750px">
      <q-infinite-scroll class="col-12" @load="onLoad" :offset="250">
        <q-list padding>
          <Item @reloadart="reloadArt" :postList="postList" />
        </q-list>
        <template v-slot:loading>
          <div class="row justify-center q-my-md">
            <q-spinner-dots color="primary" size="40px" />
          </div>
        </template>
      </q-infinite-scroll>
    </q-page>
    <q-page v-if="$q.screen.gt.sm" padding style="width:250px;">
      <q-card flat style="border-top: none;border-left: none;border-right: none;" class="my-card">
         <q-item style="padding:8px 0;">
          <q-item-section avatar>
            <q-avatar size="60px">
              <img :src="webMasterProfile.avatar"/>
            </q-avatar>
          </q-item-section>
          <q-item-section>
            <q-item-label>{{ webMasterProfile.name }}</q-item-label>
            <q-item-label caption>{{ webMasterProfile.desc }}</q-item-label>
          </q-item-section>
        </q-item>
      </q-card>
      <q-card flat style="border-top: none;border-left: none;border-right: none;" class="my-card">
        <q-card-section style="padding:8px 0;">
          <div class="text-h6">标签</div>
        </q-card-section>
        <q-separator inset />
        <q-card-section style="padding:8px 0;" class="q-pt-none">
          <q-list dense  v-for="cate in cateList" :key="cate.index">
            <q-item clickable v-ripple @click="reloadArt(cate.id)">
              <q-item-section avatar>
                <q-icon color="primary" :name="cate.icon" />
              </q-item-section>
              <q-item-section>{{ cate.name }}</q-item-section>
            </q-item>
          </q-list>
        </q-card-section>
      </q-card>
      <!-- <q-card flat style="border-top: none;border-left: none;border-right: none;" class="my-card">
        <q-card-section style="padding:8px 0;">
          <div class="text-h6">热门文章</div>
        </q-card-section>
        <q-separator inset />
        <q-card-section style="padding:8px 0;" class="q-pt-none">
          <q-list v-for="cate in cateList" :key="cate.index">
            <q-item clickable v-ripple @click="reloadArt(cate.id)">
              <q-item-section avatar>
                <q-icon color="primary" name="H" />
              </q-item-section>
              <q-item-section>{{ cate.name }}</q-item-section>
            </q-item>
          </q-list>
        </q-card-section>
      </q-card> -->
    </q-page>
  </div>
</template>
<script lang="ts">
import Item from '../../components/ItemCard.vue';
import { reactive, onBeforeMount, toRefs, watch, ref } from 'vue';
import { listArticle, getWebMaster } from '../../api/test/index';
import { useRoute } from 'vue-router';
import { date } from 'quasar';
export default {
  name: 'List',
  components: { Item },
  setup() {
    let data = reactive({
      pagesize: 10,
      total: 10,
      cid: 0,
      postList: [],
      webMasterProfile: {
        id: '',
        name: '',
        desc: '',
        avatar: null,
      },
      cateList: []
    });

    const route = useRoute() as any;
    const queryValue = ref(route.query.q);
    // const { ctx } = getCurrentInstance() as any;
    const method = {
      async commentss(): Promise<any> {
        const paramss = {
          pagenum: 1,
          pagesize: data.pagesize,
          keyword: queryValue.value
        };
        let datas = (await listArticle(paramss)) as any;
        datas.data.forEach((element) => {
          let timeStamp = new Date(element.UpdatedAt);
          let formattedString = date.formatDate(
            timeStamp,
            'YYYY-MM-DD HH:mm:ss'
          );
          element.UpdatedAt = formattedString;
        });
        data.postList = datas.data;
        data.total = datas.total;

        let res = (await getWebMaster()) as any;
        data.webMasterProfile.avatar = res.data.avatar;
        data.webMasterProfile.name = res.data.name;
        data.webMasterProfile.desc = res.data.desc;
        data.cateList = res.cate;
      },
      async onLoad(index, done): Promise<any> {
        if (data.total > data.pagesize) {
          data.pagesize += 10;
          const paramss = {
            pagenum: 1,
            pagesize: data.pagesize,
            id: data.cid,
          };
          let datas = (await listArticle(paramss)) as any;
          setTimeout(() => {
            datas.data.forEach((element) => {
              let timeStamp = new Date(element.UpdatedAt);
              let formattedString = date.formatDate(
                timeStamp,
                'YYYY-MM-DD HH:mm:ss'
              );
              element.UpdatedAt = formattedString;
            });
            data.postList = datas.data;
            data.total = datas.total;
            done();
          }, 2000);
        } else {
          done();
        }
      },
      async reloadArt(id: number): Promise<any> {
        data.cid = id;
        const paramss = {
          pagenum: 1,
          pagesize: data.pagesize,
          id: id,
        };
        let datas = (await listArticle(paramss)) as any;
        datas.data.forEach((element) => {
          let timeStamp = new Date(element.UpdatedAt);
          let formattedString = date.formatDate(
            timeStamp,
            'YYYY-MM-DD HH:mm:ss'
          );
          element.UpdatedAt = formattedString;
        });
        data.postList = datas.data;
        data.total = datas.total;
      }
    };
    onBeforeMount(async () => {
      await method.commentss();
    });
    watch(() => route.query.q, (newQuery, oldQuery) => {
      queryValue.value = newQuery;
      void method.commentss();
    });
    return {
      ...toRefs(data),
      ...method,
    };
  },
};
</script>

