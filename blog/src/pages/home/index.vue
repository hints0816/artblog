<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-10 col-md-8">
      <q-infinite-scroll @load="onLoad" :offset="250">
        <q-list padding>
          <Item :postList="postList" />
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
      onload: false,
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
        data.onload = true;
        datas.data.forEach((element) => {
          let timeStamp = new Date(element.UpdatedAt);
          let formattedString = date.formatDate(
            timeStamp,
            'YYYY-MM-DD HH:mm:ss'
          );
          element.UpdatedAt = formattedString;
        });
        data.postList = datas.data;

        console.log(datas.data);
        console.log(data.postList);
        console.log(ctx);
      },
      async onLoad(index, done): Promise<any> {
        data.pagesize += 10;
        const paramss = {
          pagenum: 1,
          pagesize: data.pagesize,
        };
        let datas = (await listArticle(paramss)) as any;
        setTimeout(() => {
          console.log(datas);
          datas.data.forEach((element) => {
            let timeStamp = new Date(element.UpdatedAt);
            let formattedString = date.formatDate(
              timeStamp,
              'YYYY-MM-DD HH:mm:ss'
            );
            element.UpdatedAt = formattedString;
          });
          data.postList = datas.data;
          console.log(data.postList);
          done();
        }, 2000);
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

