<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-10">
      <q-input autogrow v-model="title_text" label="title" :dense="dense">
        <template v-slot:after>
          <q-btn push color="primary" @click="save()" label="Save"/>
          <q-btn push color="primary" @click="alert = true" label="Comment"/>
        </template>
      </q-input>
      <v-md-editor v-model="content_text" height="90%"></v-md-editor>
      <v-md-preview :text="marktext"></v-md-preview>
      <q-dialog v-model="alert">
        <q-card>
          <q-toolbar>
            <q-avatar>
              <img src="https://cdn.quasar.dev/logo-v2/svg/logo.svg">
            </q-avatar>

            <q-toolbar-title><span class="text-weight-bold">Quasar</span> Framework</q-toolbar-title>

            <q-btn flat round dense icon="close" v-close-popup />
          </q-toolbar>
          <q-card-section class="q-pt-none">
            <q-uploader
              url="http://localhost:4444/upload"
              style="max-width: 300px"
              hide-upload-btn=false
              color="teal"
              flat
              bordered
            />
            <div class="q-gutter-xs row" style="max-width: 300px" :class="{ 'truncate-chip-labels': truncate }">
                <q-chip
                  removable
                  v-model="vanilla"
                  color="primary"
                  text-color="white"
                  icon="cake"
                  :label="vanillaLabel"
                  :title="vanillaLabel"
                />
                <q-chip
                  removable
                  v-model="chocolate"
                  color="teal"
                  text-color="white"
                  icon="cake"
                  :label="chocolateLabel"
                >
                  <q-tooltip>{{ chocolateLabel }}</q-tooltip>
                </q-chip>
                <q-chip
                  removable
                  v-model="strawberry"
                  color="orange"
                  text-color="white"
                  icon="cake"
                >
                  <div class="ellipsis">
                    {{ strawberryLabel }}
                    <q-tooltip>{{ strawberryLabel }}</q-tooltip>
                  </div>
                </q-chip>
                <q-chip
                  removable
                  v-model="cookies"
                  color="red"
                  text-color="white"
                >
                  <q-avatar>
                    <img src="https://cdn.quasar.dev/img/boy-avatar.png">
                  </q-avatar>
                  <div class="ellipsis">
                    {{ cookiesLabel }}
                    <q-tooltip>{{ cookiesLabel }}</q-tooltip>
                  </div>
                </q-chip>
              </div>
          </q-card-section>

          <q-card-actions align="right">
            <q-btn flat label="OK" color="primary" v-close-popup />
          </q-card-actions>
        </q-card>
      </q-dialog>
    </q-page>
  </div>
</template>
<script lang="ts">
import { Notify } from 'quasar'
import { ArticleInfo } from '../../api/test/article.model';
import { addArticle, getArticle } from '../../api/test/index'
import { getCurrentInstance, reactive, toRefs, onBeforeMount } from 'vue'
import 'vue-cropper/dist/index.css'
import { VueCropper }  from 'vue-cropper'
import { useRoute } from 'vue-router';

export default {
  name: 'Post',
  components: [
    VueCropper
  ],
  setup () {
    let data = reactive({
       options: {
        img: '', // 原图文件
        autoCrop: true, // 默认生成截图框
        fixedBox: false, // 固定截图框大小
        canMoveBox: true, // 截图框可以拖动
        autoCropWidth: 200, // 截图框宽度
        autoCropHeight: 200, // 截图框高度
        fixed: true, // 截图框宽高固定比例
        fixedNumber: [1, 1], // 截图框的宽高比例
        centerBox: true, // 截图框被限制在图片里面
        canMove: false, // 上传图片不允许拖动
        canScale: false // 上传图片不允许滚轮缩放
      },

      vanilla: true,
      chocolate: true,
      strawberry: true,
      cookies: true,
      alert: false,
      title_text: '',
      content_text: '',
    })
    const route = useRoute() as any;
    const {ctx} = getCurrentInstance() as any
    console.log(ctx)
    const method = {
      async save(): Promise<void> {
        if(data.title_text === '') {
          Notify.create({
            message: '请填写title',
            color: 'negative',
            icon: 'report_problem',
            position: 'top',
            timeout: 2000
          })
          return 
        }
        if(data.content_text === '') {
          Notify.create({
            message: '请填写content',
            color: 'negative',
            icon: 'report_problem',
            position: 'top',
            timeout: 2000
          })
          return 
        }
        let params: ArticleInfo = {
          title: data.title_text,
          content: data.content_text, 
        }
        let datas  = await addArticle(params) as any
        console.log(datas)
      },
      async commentss(id: number):Promise<any> {
        let datas  = await getArticle(id) as any
        data.title_text = datas.data.title
        data.content_text = datas.data.content
      }
    }
    onBeforeMount(async()=>{
        await method.commentss(route.params.id)
    })
    return {
      ...toRefs(data),
      ...method
    }
  }
};
</script>

