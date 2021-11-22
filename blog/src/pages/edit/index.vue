<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-10">
      <q-input autogrow v-model="title_text" label="title" :dense="dense">
        <template v-slot:after>
            <q-btn push color="primary" @click="save()" label="Comment"/>
        </template>
      </q-input>
      <v-md-editor v-model="content_text" height="90%"></v-md-editor>
      <v-md-preview :text="marktext"></v-md-preview>
    </q-page>
  </div>
</template>
<script lang="ts">
import { Notify } from 'quasar'
import { ArticleInfo } from '../../api/test/article.model';
import { addArticle } from '../../api/test/index'
import { getCurrentInstance, reactive, toRefs } from 'vue'
export default {
  name: 'Post',
  setup () {
    let data = reactive({
      title_text: '',
      content_text: '',
    })
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
        let params: ArticleInfo = {
          title: data.title_text,
          content: data.content_text, 
        }
        let datas  = await addArticle(params) as any
        console.log(datas)
      }
    }
    return {
      ...toRefs(data),
      ...method
    }
  }
};
</script>

