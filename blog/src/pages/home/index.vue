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
import { date } from 'quasar'
export default {
  name: 'List',
  components: { Item },
  setup () {
    let data = reactive({
      onload: false,
      postList: [
      ],
    })
    const {ctx} = getCurrentInstance() as any
    const method = {
      async commentss():Promise<any> {
        let datas  = await listArticle(1) as any
        data.onload = true
        datas.data.forEach(element => {
          let timeStamp = new Date(element.UpdatedAt)
          let formattedString = date.formatDate(timeStamp, 'YYYY-MM-DD HH:mm:ss')
          element.UpdatedAt = formattedString
        });  
        data.postList = datas.data
         
        console.log(datas.data)
        console.log(data.postList)
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

