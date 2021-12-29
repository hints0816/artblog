<template>
  <div>
    <q-tabs v-model="tab" inline-label class="q-my-md" align="left">
      <q-tab name="article" icon="mail" label="ArtList" />
      <q-tab name="setting" icon="movie" label="Setting" />
    </q-tabs>
    <q-separator />
    <q-tab-panels v-model="tab" animated>
      <q-tab-panel name="article">
        <q-card-section>
          <Edit @reloadart="listart" :postList="postList" :edit="edit" />
        </q-card-section>
      </q-tab-panel>
      <q-tab-panel name="setting">
        <Category />
      </q-tab-panel>
    </q-tab-panels>
  </div>
</template>
<script lang="ts">
import Edit from '../components/EditCard.vue';
import Category from '../components/CategoryCard.vue';
import {
  listArticle,
} from '../api/test/index';
import {
  ref,
  reactive,
  toRefs,
  onBeforeMount,
} from 'vue';
import { date } from 'quasar';
import { useRoute, useRouter } from 'vue-router';
import { Notify } from 'quasar';

export default {
  name: 'Post',
  components: { Edit, Category },
  setup() {
    let data = reactive({
      alert: false,
      name: '',
      tab: 'article',
      model: null,
      pagesize: 10,
      total: 10,
      edit: false,
      postList: [],
    });
    const route = useRoute() as any;
    const router = useRouter() as any;
    const method = {
      async listart(): Promise<any> {
        const paramss = {
          pagenum: 1,
          pagesize: data.pagesize,
          user_id: route.params.id,
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
        data.edit = datas.edit;
      }
    };
    onBeforeMount(async () => {
      await method.listart();
    });
    return {
      ...toRefs(data),
      ...method,
    };
  },
};
</script>