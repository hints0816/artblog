<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-12" style="max-width: 1200px">
      <div class="row q-col-gutter-x-lg" style="margin-left: 0px;">
        <div class="col-xs-12 col-md-3" style="height: 100%;padding-left: 0;">
          <q-card style="background-color:transparent;" flat class="my-card">
            <q-card-actions align="center">
              <div class="useravatar" style="width: 200px">
                <q-file
                  v-if="edit"
                  display-value=""
                  @input="pick"
                  v-model="filesPng"
                  class="cursor-pointer"
                  outlined
                >
                  <q-tooltip> Some text as content of Tooltip </q-tooltip>
                  <q-avatar round size="200px">
                    <img :src="profile.avatar" />
                  </q-avatar>
                  <q-badge
                    style="
                      background-color: initial !important;
                      color: initial !important;
                      top: 144px;
                    "
                    floating
                  >
                    <q-btn @click="alert = true, emoji = profile.emoji" class="bg-white" round>
                      {{ profile.emoji }}
                    </q-btn>
                  </q-badge>
                </q-file>
                <q-avatar v-else round size="200px">
                  <img :src="profile.avatar" />
                  <q-badge
                    @mouseenter="mouseover()"
                    style="
                      background-color: initial !important;
                      color: initial !important;
                      top: 144px;
                    "
                    floating
                  >
                    <q-btn class="bg-white" round>
                      {{ profile.emoji }}
                    </q-btn>
                  </q-badge>
                </q-avatar>
              </div>
            </q-card-actions>
            <q-card-actions class="text-weight-bold text-h6" align="left">
              {{ profile.name }}
            </q-card-actions>
            <q-card-actions v-if="!openEditProfile">
              <q-btn
                outline
                v-if="edit"
                @click="openEditProfile = true"
                class="text-cyan-9 full-width"
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
                <q-btn outline dense class="text-cyan-9" @click="updateProfile" label="Save" />
                <q-btn
                  outline
                  dense
                  @click="openEditProfile = false"
                  label="Cancel"
                />
              </div>
            </q-card-actions>
          </q-card>
          <q-dialog v-model="alert">
            <q-card style="min-width: 350px">
              <q-card-section>
                <div class="text-h6">Edit status</div>
              </q-card-section>
              <q-card-section class="q-pt-none">
                <q-input
                  outlined
                  v-model="profile.emoji_text"
                  label="What's happening?"
                  dense
                >
                  <template v-slot:before>
                    <q-btn
                      style="color: rgba(0, 0, 0, 1)"
                      round
                      dense
                      flat
                      :label="emoji"
                    >
                      <q-icon v-if="emoji == ''" name="mood" />
                      <q-popup-edit
                        max-width="320px"
                        style="padding: 4px 8px"
                        self="top start"
                        cover="false"
                      >
                        <div class="q-gutter-sm" style="margin-top: 0px">
                          <a
                            href="javascript:void(0);"
                            @click="setEmojiStatus(index)"
                            style="text-decoration: none"
                            v-for="(item, index) in faceList"
                            :key="index"
                            class="emotionItem"
                            >{{ item }}</a
                          >
                        </div>
                      </q-popup-edit>
                    </q-btn>
                  </template>
                </q-input>
              </q-card-section>

              <q-card-actions align="right">
                <q-btn
                  flat
                  label="OK"
                  @click="editStatus()"
                  color="primary"
                  v-close-popup
                />
              </q-card-actions>
            </q-card>
          </q-dialog>
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
        <div class="col-xs-12 col-md-9 q-col-gutter-x-lg" style="height: 100%;margin-left:0;padding-left:0;">
          <q-card style="padding-right: 24px;" flat>
            <q-tabs v-model="tab" inline-label class="q-my-md" align="left">
              <q-tab name="article" icon="feed" label="ArtList" />
              <q-tab name="comment" v-if="edit" icon="chat" label="Comment" />
              <q-tab name="artImage" icon="collections" label="ArtImage" />
            </q-tabs>
            <q-tab-panels v-model="tab" animated>
              <q-tab-panel name="article">
                <q-card-actions class="row q-col-gutter-x-xs justify-between">
                  <q-input style="width:80%" dense outlined v-model="searchText" input-class="text-left" placeholder="Search">
                    <template v-slot:append>
                      <q-icon name="search" class="cursor-pointer" @click="listart();" />
                    </template>
                  </q-input>
                  <q-btn v-if="edit" flat to="/edit" icon="post_add" class="text-cyan-9 q-mx-xs">
                  </q-btn>
                </q-card-actions>
                <q-card-section>
                  <Edit @reloadart="listart" :postList="postList" :edit="edit" />
                </q-card-section>
              </q-tab-panel>
              <q-tab-panel name="comment">
                <CommentCard/>
              </q-tab-panel>
              <q-tab-panel name="artImage">
                <ArtImgCard/>
              </q-tab-panel>
            </q-tab-panels>
          </q-card>
        </div>
      </div>
    </q-page>
  </div>
</template>
<script lang="ts">
import emoji from '../../css/emoji.json';
import { updateProfile, editEmoji } from '../../api/test/index';
import Edit from '../../components/EditCard.vue';
import CommentCard from '../../components/CommentCard.vue';
import ArtImgCard from '../../components/ArtImgCard.vue';
import {
  listArticle,
  getProfile,
  uploadImage,
  uploadAvatarImage,
} from '../../api/test/index';
import {
  ref,
  reactive,
  toRefs,
  onBeforeMount,
} from 'vue';
import { date } from 'quasar';
import { useRoute, useRouter } from 'vue-router';
import { Notify } from 'quasar';

export default {
  name: 'Post',
  components: { Edit, CommentCard, ArtImgCard },
  setup() {
    let data = reactive({
      alert: false,
      openEditProfile: false,
      name: '',
      faceList: [],
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
        emoji_text: ''
      },
      emoji: '',
      tab: 'article',
      filesPng: null,
      cropperAvatarDialog: false,
      model: null,
      option: {
        img: 'https://8.134.9.37:9999/blog/473562851655553024.png', // 原图文件
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
      searchText: ''
    });
    const route = useRoute() as any;
    const router = useRouter() as any;
    const cropper = ref(null);
    const method = {
      async listart(): Promise<any> {
        const paramss = {
          pagenum: 1,
          pagesize: data.pagesize,
          user_id: route.params.id,
          keyword: data.searchText
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
      },
      mouseover():void {
      },
      async pick(value: any): Promise<void> {
          let param = new FormData();
          param.append('file', value.target.files[0]);
          let res = (await uploadImage(param)) as any;
          if (res.status == 200) {
            data.option.img = res.url;
            data.cropperAvatarDialog = true;
          }
      },
      async updateProfile(value: any): Promise<void> {
        const params = {
          name: data.profile.name,
          wechat: data.profile.wechat,
          qq_chat: data.profile.qq_chat,
          weibo: data.profile.weibo,
          bili: data.profile.bili,
        };
        let res = (await updateProfile(params)) as any;
        if (res.status == 200) {
          Notify.create({
            message: 'Edit Successful',
            color: 'positive',
            position: 'top',
            timeout: 2000
          })
        }
      },
      upload(): void {
        let avatarfile = null;

        cropper.value.getCropData(async (data: any) => {
          // do something
          var arr = data.split(','),
            mime = arr[0].match(/:(.*?);/)[1],
            bstr = atob(arr[1]),
            n = bstr.length,
            u8arr = new Uint8Array(n);
          while (n--) {
            u8arr[n] = bstr.charCodeAt(n);
          }
          avatarfile = new File([u8arr], 'filename.jpeg', { type: mime });
          let param = new FormData();
          param.append('file', avatarfile);
          let res = (await uploadAvatarImage(param)) as any;
          if (res.status == 200) {
            router.go(0);
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
      setEmojiStatus(index: number): void {
        const face = data.faceList[index] as string;
        data.emoji = face;
      },
      async editStatus(): Promise<any> {
        const params = {
          emoji: data.emoji,
          emoji_text: data.profile.emoji_text
        };
        let res = await editEmoji(params) as any;
        if (res.status == 200) {
          router.go(0)
          return Notify.create({
            message: 'Edit Successful',
            color: 'positive',
            position: 'top',
            timeout: 2000
          })
        }
      },
    };
    onBeforeMount(async () => {
      await method.getProfile(route.params.id);
      await method.listart();
      emoji.forEach((element) => {
        data.faceList.push(element.char);
      });
    });
    return {
      cropper,
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