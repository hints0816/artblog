<template>
  <div>
    <div v-for="post in postList" :key="post.index">
      <q-item clickable v-ripple class="q-mt-lg">
        <q-item-section @click="toPostDetail(post.ID)">
          <q-item-label>
            <div class="text-h6 text-cyan-9">{{ post.title }}</div>
          </q-item-label>
          <q-item-label
            lines="4"
            caption
            class="text-body1 text-gray-light text-justify"
          >
            {{ post.desc }}
          </q-item-label>
        </q-item-section>
        <q-item-section @click="toPostDetail(post.ID)" class="q-mt-sm" side top>
          <q-item-label caption>
            <q-chip v-if="post.status == 0" dense square icon="bookmark">draft</q-chip>
            {{ post.UpdatedAt }}
          </q-item-label>
        </q-item-section>
        <q-item-section v-if="edit" top side>
          <div class="text-grey-8">
            <q-btn class="gt-xs" @click="toEdit(post.ID)" size="12px" flat dense round icon="edit" />
            <q-btn class="gt-xs" @click="showDelNotify(post.ID)"  size="12px" flat dense round icon="delete">
              <!-- <q-popup-proxy>
                <q-banner>
                  Are you absolutely sure?
                  <q-btn color="negative" @click="delArt(post.ID)" label="DELETE" />
                </q-banner>
              </q-popup-proxy> -->
            </q-btn>
          </div>
        </q-item-section>
      </q-item>
      <q-separator spaced inset />
    </div>
  </div>
</template>

<script lang="ts">
import { getCurrentInstance } from 'vue';
import { Notify } from 'quasar'
import { useRouter } from 'vue-router';
import { delArticle } from '../api/test/index';
export default {
  name: 'Item',
  props: {
    postList: {
      type: Array,
      required: true,
    },
    edit: {
      type: Boolean,
      required: true,
    },
  },
  setup() {
    const router = useRouter() as any;
    const { ctx } = getCurrentInstance() as any;
    const method = {
      toPostDetail(id: number): void {
        router.push(`/posts/${id}`);
      },
      toEdit(id: number): void {
        router.push(`/edit/${id}`)
      },
      showDelNotify(id: number) {
        Notify.create({
          message: 'Are you absolutely sure?',
          color: 'negative',
          icon: 'report_problem',
          position: 'top',
          actions: [
            { label: 'yes', color: 'white', handler: async () => { 
              let res = (await delArticle(id)) as any;
              if(res.status == 200) {
                ctx.$emit('reloadart')
              }
            }}
          ]
        })
      }
    };
    return {
      ...method,
    };
  },
  methods: {},
};
</script>

<style lang="scss" scoped>
.markdown-body hr {
  height: unset;
}
</style>