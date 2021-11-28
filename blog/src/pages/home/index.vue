<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-10 col-md-8">
      <q-infinite-scroll @load="onLoad" :offset="250">
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
  </div>
</template>
<script lang="ts">
import Item from '../../components/ItemCard.vue';
import { getCurrentInstance, reactive, onBeforeMount, toRefs } from 'vue';
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
    const { ctx } = getCurrentInstance() as any;
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
        console.log(datas.data);
        console.log(data.postList);
        console.log(ctx);
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
            console.log(data.postList);
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

