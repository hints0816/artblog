<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-10">
      <q-input autogrow v-model="title_text" label="title" :dense="dense">
        <template v-slot:after>
          <q-btn v-if="content_status == 0" push color="primary" @click="save(0)" label="Save"/>
          <q-btn push color="primary" @click="commentAlert()" label="Comment"/>
        </template>
      </q-input>
      <v-md-editor 
        v-model="content_text" 
        height="90%"
            :disabled-menus="[]"
        @upload-image="handleUploadImage"
      ></v-md-editor>
      <q-dialog v-model="alert">
        <q-card style="width: 300px">
          <q-toolbar>
            <q-avatar>
              <img src="https://cdn.quasar.dev/logo-v2/svg/logo.svg">
            </q-avatar>

            <q-toolbar-title><span class="text-weight-bold">Quasar</span> Framework</q-toolbar-title>

            <q-btn flat round dense icon="close" v-close-popup />
          </q-toolbar>
          <q-card-section class="q-pt-none">
            <q-uploader
              @added="addImage"
              style="width:initial!important;"
              hide-upload-btn=false
              flat
              bordered
            />
            <div style="max-width: 300px" :class="{ 'truncate-chip-labels': truncate }">
                 <q-chip
                  v-for="label,index in select_tag"
                  removable
                  class="label"
                  :color="label.color"
                  :text-color="label.text_color"
                  :icon="label.icon"
                  :key="index"
                  @remove="removeChipClickHandler(index)"
                >
                  {{ label.name }}
                </q-chip>
              </div>
          </q-card-section>
          <q-card-section>
            <q-form
              @submit="onSubmit"
              class="q-gutter-md"
            >
              <q-input outlined dense v-model="category_text" label="Label" />
            </q-form>
            <q-chip
            v-for="label,index in unselect_tag"
            clickable
            class="label"
            :color="label.color"
            :text-color="label.text_color"
            :icon="label.icon"
            :key="index"
            @click="addChipClickHandler(index)"
          >
            {{ label.name }}
          </q-chip>
          </q-card-section>
          
          <q-card-actions align="right">
            <q-btn flat label="OK" @click="save(1)" color="primary" v-close-popup />
          </q-card-actions>
        </q-card>
      </q-dialog>
    </q-page>
  </div>
</template>
<script lang="ts">
import { Notify } from 'quasar'
import { ArticleInfo } from '../../api/test/article.model';
import { addArticle, getArticle, uploadImage, listCategory } from '../../api/test/index'
import { getCurrentInstance, reactive, toRefs, onBeforeMount } from 'vue'
import { useRoute, useRouter } from 'vue-router';
import VueMarkdownEditor, { xss } from '@kangc/v-md-editor';

export default {
  name: 'Post',
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
      content_img: '',
      content_id: 0,
      content_status: 0,
      select_tag: [],
      unselect_tag: [],
      category_text:''
    })
    const route = useRoute() as any;
    const router = useRouter() as any;
    const {ctx} = getCurrentInstance() as any
    const method = {
      onSubmit(): void {
        data.select_tag.push({
          name: data.category_text
        })
      },
      addChipClickHandler(index: number): void{
        data.select_tag.push(data.unselect_tag[index])
        data.unselect_tag.splice(index,1)
      },
      removeChipClickHandler(index: number): void{
        if(data.select_tag[index].id != undefined) {
          data.unselect_tag.push(data.select_tag[index])
        }
        data.select_tag.splice(index,1)  
      },
      async commentAlert(): Promise<void> {
        data.alert = true
        const paramss = {
          pagenum: 1,
          pagesize: 20,
        };
        let res  = await listCategory(paramss) as any
        data.unselect_tag = res.data
        data.select_tag = []
      },
      async save(status: number): Promise<void> {
        if(data.title_text === '') {
          Notify.create({
            message: 'title can\'t be empty',
            color: 'negative',
            icon: 'report_problem',
            position: 'top',
            timeout: 2000
          })
          return 
        }
        if(data.content_text === '') {
          Notify.create({
            message: 'content can\'t be empty',
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
          status: status,
          id: data.content_id,
          img: data.content_img,
          desc: xss.process(VueMarkdownEditor.vMdParser.themeConfig.markdownParser.render(data.content_text))
        }
        let res  = await addArticle(params) as any
        if (res.status == 200) {
          if (res.data.status == 0) {
            if (data.content_id == 0) {
              // let contentId = res.data.ID as string
              // let url = route.path as string
              // history.replaceState('','','#'+url+'/'+contentId)
              data.content_id = res.data.ID
            }
            Notify.create({
              message: 'Save Successful',
              color: 'positive',
              icon: 'report_problem',
              position: 'top',
              timeout: 3000
            })
          } else {
            Notify.create({
              message: 'Publish Successful',
              type: 'positive',
              position: 'top',
              timeout: 5000,
              actions: [
                { label: 'Dismiss', color: 'white', handler: () => { /* ... */ } }
              ]
            })
            router.go(-1)
          }
        } 
      },
      async commentss(id: number):Promise<any> {
        let datas  = await getArticle(id) as any
        if (datas != undefined) {
          data.content_status = datas.data.status
          data.content_id = datas.data.ID
          data.title_text = datas.data.title
          data.content_text = datas.data.content
        }
      },
      async handleUploadImage(event, insertImage, files):Promise<any> {
        let param = new FormData()
        param.append('file', files[0])
        let res = await uploadImage(param) as any
        if(res.status == 200) {
          insertImage({
            url: res.url,
            desc: '七龙珠',
            // width: 'auto',
            // height: 'auto',
          });
        } else {

        }
      },
      async addImage(files: Array<File>): Promise<any> {
        let param = new FormData()
        param.append('file', files[0])
        let res = await uploadImage(param) as any
        if(res.status == 200) {
          data.content_img = res.url
        }
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

