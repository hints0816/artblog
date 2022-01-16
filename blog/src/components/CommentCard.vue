<template>
  <q-infinite-scroll @load="onLoad" :offset="200">
    <div v-for="data in commentList" :key="data.index">
      <q-item clickable v-ripple class="q-mt-lg">
        <q-item-section @click="toPostDetail(data.article_id)">
          <q-item-label
            caption
            class="text-justify"
          >
          {{ data.FromProfile.name }} 回复：
            {{ data.Title }}
          </q-item-label>
          <q-item-label>
            <div class="text-h6 text-cyan-9">{{ data.content }}</div>
          </q-item-label>
        </q-item-section>
        <q-item-section @click="toPostDetail(data.ID)" class="q-mt-sm" side top>
          <q-item-label caption>
            <q-chip v-if="data.status == 0" dense square icon="bookmark">draft</q-chip>
            {{ data.UpdatedAt }}
          </q-item-label>
        </q-item-section>
        <q-item-section top side>
          <div class="text-grey-8">
            <q-btn class="gt-xs" @click="showDelNotify(data.ID)"  size="12px" flat dense round icon="delete">
            </q-btn>
          </div>
        </q-item-section>
      </q-item>
      <q-separator spaced inset />
    </div>
  </q-infinite-scroll>
</template>
<script lang="ts">
import {
  listSettingComment,
  delComment
} from '../api/test/index';
import {
  ref,
  reactive,
  toRefs,
  onBeforeMount,
} from 'vue';
import { date, Notify } from 'quasar';
import { useRoute, useRouter } from 'vue-router';

export default {
  name: 'Post',
  setup() {
    let data = reactive({
      alert: false,
      name: '',
      tab: 'article',
      model: null,
      pagesize: 10,
      total: 10,
      commentList: [],
    });
    const route = useRoute() as any;
    const router = useRouter() as any;
    const method = {
      toPostDetail(id: number): void {
        router.push(`/posts/${id}`);
      },
      showDelNotify(id: number): void {
        Notify.create({
          message: 'Are you absolutely sure?',
          color: 'negative',
          icon: 'report_problem',
          position: 'top',
          actions: [
            { label: 'yes', color: 'white', handler: async () => { 
              let res = (await delComment(id)) as any;
              if(res.status == 200) {
                return method.listComment()
              }
            }}
          ]
        })
      },
      async listComment(): Promise<any> {
        const paramss = {
          pagenum: 1,
          pagesize: data.pagesize,
          user_id: route.params.id,
        };
        let datas = (await listSettingComment(paramss)) as any;
        datas.data.forEach((element) => {
          let timeStamp = new Date(element.UpdatedAt);
          let formattedString = date.formatDate(
            timeStamp,
            'YYYY-MM-DD HH:mm:ss'
          );
          element.UpdatedAt = formattedString;
        });
        data.commentList = datas.data;
        data.total = datas.total;
      }
    };
    onBeforeMount(async () => {
      await method.listComment();
    });
    return {
      ...toRefs(data),
      ...method,
    };
  },
};
</script>