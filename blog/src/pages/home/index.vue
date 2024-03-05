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
              <img src="https://cdn.quasar.dev/img/avatar2.jpg">
            </q-avatar>
          </q-item-section>
          <q-item-section>
            <q-item-label>Title</q-item-label>
            <q-item-label caption>Subhead</q-item-label>
          </q-item-section>
        </q-item>
      </q-card>
      <q-card flat style="border-top: none;border-left: none;border-right: none;" class="my-card">
        <q-card-section style="padding:8px 0;">
          <div class="text-h6">Our Changing Planet</div>
        </q-card-section>

        <q-card-section style="padding:8px 0;" class="q-pt-none">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
          tempor incididunt ut labore et dolore magna aliqua.
        </q-card-section>

        <q-separator inset />

        <q-card-section style="padding:8px 0;">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
          tempor incididunt ut labore et dolore magna aliqua.
        </q-card-section>
      </q-card>
    </q-page>
  </div>
</template>
<script lang="ts">
import Item from '../../components/ItemCard.vue';
import { reactive, onBeforeMount, toRefs } from 'vue';
import { listArticle } from '../../api/test/index';
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
    });
    // const { ctx } = getCurrentInstance() as any;
    const method = {
      async commentss(): Promise<any> {
        const paramss = {
          pagenum: 1,
          pagesize: data.pagesize,
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
      },
    };
    onBeforeMount(async () => {
      await method.commentss();
    });
    return {
      ...toRefs(data),
      ...method,
    };
  },
};
</script>

