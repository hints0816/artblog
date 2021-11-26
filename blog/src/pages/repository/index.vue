<template>
  <div class="row justify-center">
    <q-page
      padding
      class="q-mt-lg col-xs-12 col-sm-12"
      style="max-width: 1200px"
    >
      <q-item>
        <q-item-section avatar>
          <q-btn round>
            <q-avatar>
              <img src="https://cdn.quasar.dev/img/avatar2.jpg" />
            </q-avatar>
          </q-btn>
        </q-item-section>
        <q-item-section>
          <q-item-label caption>Caption</q-item-label>
          <q-item-label>
            <code class="text-italic"> </code>
          </q-item-label>
        </q-item-section>
      </q-item>
      <div class="row q-col-gutter-x-lg">
        <div class="useravatar col-xs-12 col-sm-3" style="height: 100%">
          <q-file
                    display-value=""
                    @input="pick"
                    @rejected="upload"
                    v-model="filesPng"
                    outlined
                  >
                  <q-tooltip>
                    Some text as content of Tooltip
                  </q-tooltip>
                    <q-avatar round size="200px">
                      <img src="https://cdn.quasar.dev/img/avatar5.jpg" />
                      <q-badge color="red" floating>4</q-badge>
                    </q-avatar>
                  </q-file>
        </div>
         <q-dialog v-model="cropperAvatarDialog">
            <q-card>
              <q-toolbar>
                <q-toolbar-title>
                  <span class="text-weight-bold">
                    Crop your new profile picture
                  </span>
                </q-toolbar-title>
                <q-btn flat round dense icon="close" v-close-popup />
              </q-toolbar>
              <q-card-section style="height: 300px">
                <vue-cropper
                  ref="cropper"
                  :img="option.img"
                  :outputSize="option.size"
                  :outputType="option.outputType"
                  :info="option.info"
                  :canScale="option.canScale"
                  :autoCrop="option.autoCrop"
                  :fixedBox="option.fixedBox"
                  autoCropWidth="200"
                  autoCropHeight="200"
                ></vue-cropper>
              </q-card-section>
              <q-card-section>
                <q-btn class="full-width" color="primary" @click="upload" label="Set new profile picture" />
              </q-card-section>
            </q-card>
          </q-dialog>
        <div class="col-sm-9 q-col-gutter-x-lg" style="height: 100%">
          <q-card flat>
            <q-tabs
              v-model="tab"
              inline-label
              class="q-my-md bg-purple text-white shadow-2"
              align="left"
            >
              <q-tab name="mails" icon="mail" label="Mails" />
              <q-tab name="alarms" icon="alarm" label="Alarms" />
              <q-tab name="movies" icon="movie" label="Movies" />
            </q-tabs>
            <q-separator inset />
             <q-card-actions align="right">
              <q-input dense outlined v-model="text" label="Outlined" />
                <q-select
                    dense
                    outlined
                    v-model="model"
                    :options="['Google','Facebook'
                    ]"
                    style="min-width: 130px;"
                    label="Outlined"
                />
                <q-select
                    dense
                    outlined
                    v-model="model"
                    :options="['Google','Facebook'
                    ]"
                    style="min-width: 130px;"
                    label="Outlined"
                />
                <q-select
                    dense
                    outlined
                    v-model="model"
                    :options="['Google','Facebook'
                    ]"
                    style="min-width: 130px;"
                    label="Outlined"
                />
                 <q-btn color="deep-orange" push>
                    <div class="row items-center no-wrap">
                        <q-icon left name="book" />
                        <div class="text-center">
                        New
                        </div>
                    </div>
                    </q-btn>
             </q-card-actions>
            <q-separator inset />
            <q-card-section class="row q-col-gutter-x-lg">
              <div>
                
              </div>
            </q-card-section>
          </q-card>
          <q-dialog v-model="cropperAvatarDialog">
            <q-card>
              <q-toolbar>
                <q-toolbar-title>
                  <span class="text-weight-bold">
                    Crop your new profile picture
                  </span>
                </q-toolbar-title>
                <q-btn flat round dense icon="close" v-close-popup />
              </q-toolbar>
              <q-card-section style="height: 300px">
                <vue-cropper
                  ref="cropper"
                  :img="option.img"
                  :outputSize="option.size"
                  :outputType="option.outputType"
                  :info="option.info"
                  :canScale="option.canScale"
                  :autoCrop="option.autoCrop"
                  :fixedBox="option.fixedBox"
                  autoCropWidth="200"
                  autoCropHeight="200"
                ></vue-cropper>
              </q-card-section>
              <q-card-section>
                <q-btn
                  class="full-width"
                  color="primary"
                  @click="upload"
                  label="Set new profile picture"
                />
              </q-card-section>
            </q-card>
          </q-dialog>
        </div>
      </div>
    </q-page>
  </div>
</template>
<script lang="ts">
import { ArticleInfo } from '../../api/test/article.model';
import { addArticle } from '../../api/test/index';
import { getCurrentInstance, reactive, toRefs } from 'vue';
export default {
  name: 'Post',
  setup() {
    let data = reactive({
      filesPng: null,
      cropperAvatarDialog: false,
      model: null,
      option: {
        img:
          'https://img-blog.csdnimg.cn/00f2b25ffece44768f17f9bb892f42d1.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA56qd6JuL5Lq6,size_20,color_FFFFFF,t_70,g_se,x_16#pic_center', // 原图文件
        size: 1,
        outputType: 'jpeg',
        info: true,
        canScale: true,
        autoCrop: true,
        fixedBox: true,
      },
      title_text: '',
      content_text: '',
    });
    const { ctx } = getCurrentInstance() as any;
    console.log(ctx);
    const method = {
      async save(): Promise<void> {
        let params: ArticleInfo = {
          title: data.title_text,
          content: data.content_text,
        };
        let datas = (await addArticle(params)) as any;
        console.log(datas);
      },
      pick(value: any): void {
        console.log(value);
        data.cropperAvatarDialog = true;
      },
      upload(): void {
        ctx.$refs.cropper.getCropData((data) => {
          // do something
          console.log(data);
        });
      },
    };
    return {
      ...toRefs(data),
      ...method,
    };
  },
};
</script>
<style lang="sass">
.my-menu-link
    color: white
    background: #F2C037
.useravatar .q-field--outlined .q-field__control
    border-radius: 100px
    padding: 0px
</style>