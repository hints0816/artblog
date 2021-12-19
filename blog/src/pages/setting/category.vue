<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-12" style="max-width: 1200px">
      <div class="row q-col-gutter-x-lg" style="margin-left: 0px;">
        <div class="col-xs-12 col-md-3" style="height: 100%;padding-left: 0;">
          <q-card style="background-color:transparent;" flat class="my-card">
            <q-card-actions align="center">
              <div class="useravatar full-width">
                 <q-list bordered separator>
                  <q-item clickable v-ripple to="/setting/article">
                    <q-item-section>Single line item</q-item-section>
                  </q-item>

                  <q-item clickable v-ripple :active="true" active-class="bg-teal-1 text-grey-8">
                    <q-item-section>
                      <q-item-label>Item with caption</q-item-label>
                      <q-item-label caption>Caption</q-item-label>
                    </q-item-section>
                  </q-item>

                  <q-item clickable v-ripple>
                    <q-item-section>
                      <q-item-label overline>OVERLINE</q-item-label>
                      <q-item-label>Item with caption</q-item-label>
                    </q-item-section>
                  </q-item>
                </q-list>
              </div>
            </q-card-actions>
          </q-card>
        </div>
        <div class="col-xs-12 col-md-9 q-col-gutter-x-lg" style="height: 100%;margin-left:0;padding-left:0;">
          <q-card style="padding-right: 24px;" flat>
              <q-tab-panel name="setting">
                <Category/>
              </q-tab-panel>
          </q-card>
        </div>
      </div>
    </q-page>
  </div>
</template>
<script lang="ts">
import emoji from '../../css/emoji.json';
import { updateProfile, editEmoji } from '../../api/test/index';
import Category from '../../components/CategoryCard.vue';
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
  components: { Category },
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
    const cropper = ref(null);
    const method = {
      async listart(): Promise<any> {
        const paramss = {
          pagenum: 1,
          pagesize: data.pagesize,
          user_id: route.params.id,
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
        console.log(1)
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