<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-10 col-md-8">
      <q-list padding>
        <Item :postList="postList" />
      </q-list>
    </q-page>
  </div>
</template>
<script lang="ts">
import Item from '../../components/ItemCard.vue';
import { getCurrentInstance, reactive, onBeforeMount, toRefs } from 'vue'
import { listArticle } from '../../api/test/index'
export default {
  name: 'List',
  components: { Item },
  setup () {
    let data = reactive({
      onload: false,
      postList: [
        {
          number: 1,
          title: '1',
          updated_at: '1',
          body_html: '1'
        },
        {
          number: 2,
          title: '2',
          updated_at: '13',
          body_html: '13'
        }
      ],
    })
    const {ctx} = getCurrentInstance() as any
    const method = {
      async commentss():Promise<any> {
        let datas  = await listArticle(1) as any
        data.onload = true
        console.log(datas)
       /*  data.postList.number = datas.data.id
        data.postList.title = datas.data.title
        data.postList.updated_at = datas.data.CreatedAt
        data.postList.body_html = datas.data.content */
        console.log(ctx)
      }
    }
     onBeforeMount(async()=>{
        await method.commentss()
    })
    return {
      ...toRefs(data),
      ...method
    }
  }
};
</script>

