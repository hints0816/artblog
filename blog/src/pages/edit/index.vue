<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-10">
      <q-input autogrow v-model="title_text" label="title" :dense="dense">
        <template v-slot:after>
          <q-btn push color="primary" @click="save()" label="Save" />
          <q-btn push color="primary" @click="alert = true" label="Comment" />
        </template>
      </q-input>
      <v-md-editor v-model="content_text" height="90%"></v-md-editor>
      <v-md-preview :text="marktext"></v-md-preview>
      <q-dialog persistent v-model="alert">
        <q-card>
          <q-toolbar>
            <q-avatar>
              <img src="https://cdn.quasar.dev/logo-v2/svg/logo.svg" />
            </q-avatar>
            <q-toolbar-title
              ><span class="text-weight-bold">Quasar</span>
              Framework</q-toolbar-title
            >
            <q-btn flat round dense icon="close" v-close-popup />
          </q-toolbar>

          <q-card-section class="q-pt-none">
            <div class="q-gutter-sm">
              <q-radio dense v-model="shape" val="picture" label="picture" />
              <q-radio
                dense
                v-model="shape"
                val="nopicture"
                label="nopicture"
              />
            </div>
            <div class="q-py-md">
              <q-btn style="padding: 0px">
                <img
                  style="max-width: 400px; height: 200px"
                  src="https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fwww.pianshen.com%2Fimages%2F886%2F002965543cfb901f26f46cd7865740d6.png&refer=http%3A%2F%2Fwww.pianshen.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1640349629&t=3fed3b7c5857985581dee4f416a909b3"
                />
              </q-btn>
            </div>
            <div
              class="q-gutter-xs row"
              style="max-width: 300px"
              :class="{ 'truncate-chip-labels': truncate }"
            >
              <q-chip
                v-for="(item, index) in chip"
                :key="index"
                removable
                :v-model="item.status"
                color="primary"
                text-color="white"
                icon="cake"
                :label="item.name"
              />
              <q-chip
                clickable
                @click="cropperAvatarDialog = true"
                color="teal"
                text-color="white"
                icon="bookmark"
              >
                Add Label
              </q-chip>
            </div>
          </q-card-section>

          <q-card-actions align="right">
            <q-btn flat label="OK" color="primary" v-close-popup />
          </q-card-actions>
        </q-card>
      </q-dialog>
      <q-dialog v-model="cropperAvatarDialog" persistent>
        <q-card style="width: 300px">
          <q-toolbar>
            <q-space />
            <q-btn flat round dense icon="close" v-close-popup />
          </q-toolbar>
          <q-card-section>
            <q-input v-model="text" label="Standard" />
          </q-card-section>
          <q-separator inset />
          <q-card-section>
            
          </q-card-section>
        </q-card>
      </q-dialog>
    </q-page>
  </div>
</template>

<script lang="ts">
import { Notify } from 'quasar';
import { ArticleInfo } from '../../api/test/article.model';
import { addArticle } from '../../api/test/index';
import { getCurrentInstance, reactive, toRefs } from 'vue';
export default {
  name: 'Post',
  setup() {
    let data = reactive({
      text: '',
      shape: 'nopicture',
      chip: [
        {
          name: 'vanilla',
          status: true,
        },
        {
          name: 'chocolate',
          status: true,
        },
        {
          name: 'strawberry',
          status: true,
        },
      ],
      alert: false,
      cropperAvatarDialog: false,
      title_text: '',
      content_text: '',
    });
    const { ctx } = getCurrentInstance() as any;
    console.log(ctx);
    const method = {
      async save(): Promise<void> {
        if (data.title_text === '') {
          Notify.create({
            message: '请填写title',
            color: 'negative',
            icon: 'report_problem',
            position: 'top',
            timeout: 2000,
          });
          return;
        }
        if (data.content_text === '') {
          Notify.create({
            message: '请填写content',
            color: 'negative',
            icon: 'report_problem',
            position: 'top',
            timeout: 2000,
          });
          return;
        }
        let params: ArticleInfo = {
          title: data.title_text,
          content: data.content_text,
        };
        let datas = (await addArticle(params)) as any;
        console.log(datas);
      },
    };
    return {
      ...toRefs(data),
      ...method,
    };
  },
};
</script>

