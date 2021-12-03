<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-12" style="max-width: 1200px">
      <div class="row q-col-gutter-x-lg">
        <div class="col-xs-12 col-sm-3" style="height: 100%">
          <q-card flat class="my-card">
            <q-card-actions align="center">
              <div class="useravatar" style="width: 200px">
                <q-file
                  display-value=""
                  @input="pick"
                  @rejected="upload"
                  v-model="filesPng"
                  outlined
                >
                  <q-tooltip> Some text as content of Tooltip </q-tooltip>
                  <q-avatar round size="200px">
                    <img :src="profile.avatar" />
                  </q-avatar>
                </q-file>
              </div>
            </q-card-actions>
            <q-card-actions class="text-weight-bold text-h6" align="left">
              {{ profile.name }}
            </q-card-actions>
            <q-card-actions v-if="!openEditProfile">
              <q-btn
                @click="openEditProfile = true"
                color="primary"
                class="full-width"
                label="Edit profile"
              />
            </q-card-actions>
            <q-card-actions v-else class="row q-col-gutter-y-xs">
              <q-input
                outlined
                class="full-width"
                v-model="profile.name"
                dense
                label="name"
              />
              <q-input
                outlined
                class="full-width"
                v-model="profile.email"
                dense
                label="email"
                readonly 
              />
              <q-input
                outlined
                class="full-width"
                v-model="profile.wechat"
                dense
                label="weixin"
              >
                <template v-slot:before>
                  <q-avatar>
                    <q-icon size="20px" name="fab fa-weixin" />
                  </q-avatar>
                </template>
              </q-input>
              <q-input
                outlined
                class="full-width"
                v-model="profile.qq_chat"
                dense
                label="qq"
              >
                <template v-slot:before>
                  <q-avatar>
                    <q-icon size="20px" name="fab fa-qq" />
                  </q-avatar>
                </template>
              </q-input>
              <q-input
                outlined
                class="full-width"
                v-model="profile.weibo"
                dense
                label="weibo"
              >
                <template v-slot:before>
                  <q-avatar>
                    <q-icon size="20px" name="fab fa-weibo" />
                  </q-avatar>
                </template>
              </q-input>
              <q-input
                outlined
                class="full-width"
                v-model="profile.bili"
                dense
                label="bilibili"
              >
                <template v-slot:before>
                  <q-avatar>
                    <q-icon
                      size="25px"
                      name="img:icons/icon_bilibili-circle.svg"
                    />
                  </q-avatar>
                </template>
              </q-input>
              <div class="q-gutter-sm">
                <q-btn dense color="primary" label="Save" />
                <q-btn
                  dense
                  @click="openEditProfile = false"
                  color="white"
                  text-color="black"
                  label="Cancel"
                />
              </div>
            </q-card-actions>
          </q-card>
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
              <q-btn
                class="full-width"
                color="primary"
                @click="upload"
                label="Set new profile picture"
              />
            </q-card-section>
          </q-card>
        </q-dialog>
        <div class="col-xs-12 col-sm-9 q-col-gutter-x-lg" style="height: 100%">
          <q-card flat>
            <q-tabs v-model="tab" inline-label class="q-my-md" align="left">
              <q-tab name="Article" icon="mail" label="Mails" />
              <q-tab name="Movie" icon="alarm" label="Alarms" />
              <q-tab name="Music" icon="movie" label="Movies" />
            </q-tabs>
            <q-separator />
            <q-card-actions class="row q-col-gutter-x-xs justify-between">
              <q-input class="col-grow col-xs-12 col-sm-8" dense outlined v-model="text" label="Outlined" />
              <q-select
                dense
                outlined
                v-model="model"
                class="col-grow"
                :options="['Google', 'Facebook']"
                style="width: 130px"
                label="Outlined"
              />
              <q-btn to="/edit" color="deep-orange q-mx-xs" push>
                <div class="row items-center no-wrap">
                  <q-icon left name="book" />
                  <div class="text-center">New</div>
                </div>
              </q-btn>
            </q-card-actions>
            <q-separator />
            <q-card-section>
              <Edit @reloadart="listart" :postList="postList" />
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
import Edit from '../../components/EditCard.vue';
import { ArticleInfo } from '../../api/test/article.model';
import { addArticle, listArticle, getProfile } from '../../api/test/index';
import { getCurrentInstance, reactive, toRefs, onBeforeMount } from 'vue';
import { date, Dark } from 'quasar';
import { useRoute } from 'vue-router';

export default {
  name: 'Post',
  components: { Edit },
  setup() {
    let data = reactive({
      openEditProfile: false,
      name:'',
      profile: {
        avatar: '',
        bili: '',
        desc: '',
        email: '',
        id: 0,
        img: '',
        name: '',
        qq_chat: '',
        wechat: '',
        weibo: '',
      },
      tab: 'mails',
      filesPng: null,
      cropperAvatarDialog: false,
      model: null,
      option: {
        img: 'https://img-blog.csdnimg.cn/00f2b25ffece44768f17f9bb892f42d1.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBA56qd6JuL5Lq6,size_20,color_FFFFFF,t_70,g_se,x_16#pic_center', // 原图文件
        size: 1,
        outputType: 'jpeg',
        info: true,
        canScale: true,
        autoCrop: true,
        fixedBox: true,
      },
      title_text: '',
      content_text: '',
      pagesize: 10,
      total: 10,
      postList: [],
    });
    const route = useRoute() as any;
    const { ctx } = getCurrentInstance() as any;
    console.log(ctx);
    const method = {
      async listart(): Promise<any> {
        const paramss = {
          pagenum: 1,
          pagesize: data.pagesize,
        };
        let datas = (await listArticle(paramss)) as any;
        datas.data.forEach((element) => {
          let timeStamp = new Date(element.UpdatedAt);
          let formattedString = date.formatDate(
            timeStamp,
            'YYYY-MM-DD HH:mm:ss'
          );
          element.UpdatedAt = formattedString;
        });
        data.postList = datas.data;
        data.total = datas.total;
        console.log(datas.data);
        console.log(data.postList);
        console.log(ctx);
      },
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
      async getProfile(id: number): Promise<any> {
        let res = (await getProfile(id)) as any;
        if (res.status == 200) {
          data.profile = res.data;
        } else {
        }
      },
    };
    onBeforeMount(async () => {
      await method.getProfile(route.params.id);
      await method.listart();
    });
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