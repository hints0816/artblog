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
                  v-model="filesPng"
                  outlined
                >
                  <q-tooltip> Some text as content of Tooltip </q-tooltip>
                  <q-avatar round size="200px">
                    <img :src="profile.avatar" />
                  </q-avatar>
                   <q-badge 
                    style="background-color:initial!important;color:initial!important;top:144px;" 
                    floating>
                     <q-btn class="bg-white" round>
                        {{ profile.emoji }}
                     </q-btn>
                   </q-badge>
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
              <Edit @reloadart="listart" :postList="postList" :edit="edit" />
            </q-card-section>
          </q-card>
        </div>
      </div>
    </q-page>
  </div>
</template>
<script lang="ts">
import Edit from '../../components/EditCard.vue';
import { listArticle, getProfile, uploadImage,uploadAvatarImage } from '../../api/test/index';
import { getCurrentInstance, reactive, toRefs, onBeforeMount } from 'vue';
import { date } from 'quasar';
import { useRoute, useRouter } from 'vue-router';

export default {
  name: 'Post',
  components: { Edit },
  setup() {
    let data = reactive({
      openEditProfile: false,
      name:'',
      profile: {
        avatar: '',
        emoji: '😚',
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
        img: 'http://47.119.167.128:9999/blog/473562851655553024.png', // 原图文件
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
      edit: false,
      postList: [],
    });
    const route = useRoute() as any;
    const router = useRouter() as any;
    const { ctx } = getCurrentInstance() as any;
    console.log(ctx);
    const method = {
      async listart(): Promise<any> {
        const paramss = {
          pagenum: 1,
          pagesize: data.pagesize,
          user_id: route.params.id
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
        data.edit = datas.edit;
        console.log(ctx);
      },
      async pick(value: any): Promise<void> {
        let param = new FormData()
        param.append('file', value.target.files[0])
        let res = await uploadImage(param) as any
        if(res.status == 200) {
          data.option.img = res.url
          data.cropperAvatarDialog = true;
        } else {

        }
      },
      upload(): void {
        let avatarfile = null
        this.$refs.cropper.getCropData(async (data:any) => {
          console.log(data)
          // do something
          var arr = data.split(','),
          mime = arr[0].match(/:(.*?);/)[1],
          bstr = atob(arr[1]),
          n = bstr.length,
          u8arr = new Uint8Array(n)
          while (n--) {
            u8arr[n] = bstr.charCodeAt(n)
          }
          avatarfile = new File([u8arr], 'filename.jpeg', { type: mime })
          let param = new FormData()
          param.append('file', avatarfile)
          let res = await uploadAvatarImage(param) as any
          if(res.status == 200) {
            router.go(0)
          } else {

          }
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