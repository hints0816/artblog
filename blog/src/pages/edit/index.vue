<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-10">
      <q-input autogrow v-model="text" label="title" :dense="dense">
        <q-btn push color="primary" @click="push()" label="Comment"/>
      </q-input>
      <v-md-editor v-model="marktext" height="90%"></v-md-editor>
      <v-md-preview :text="marktext"></v-md-preview>
    </q-page>
  </div>
</template>
<script lang="ts">
import { getCurrentInstance, reactive, toRefs } from 'vue'
import { addArt } from '../../api/test/index'
export default {
  name: 'Post',
  setup () {
    let data = reactive({
      marktext: '',
      text: '',
    })
    const {ctx} = getCurrentInstance() as any
    console.log(ctx)
    const method = {
      async push(): Promise<void> {
        const params = {
          title: data.text,
          content: data.marktext
        }
        let data2  = await addArt(params)
        console.log(data2)
      }
    }
    return {
      ...toRefs(data),
      ...method
    }
  }
};
</script>

