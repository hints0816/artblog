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
import { ActicleInfo } from '../../api/test/acticle.model';
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
        let params: ActicleInfo = {
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

