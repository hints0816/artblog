<template>
  <q-layout>
    <q-header>
      <q-toolbar :style="{ 'background-color': hex }">
        <q-btn
          v-if="!$q.screen.gt.sm"
          breakpoint="500"
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        />

        <q-toolbar-title shrink class="text-weight-bold">
          <router-link to="/" style="color: #fff">ArtBlog</router-link>
        </q-toolbar-title>
        <q-tabs v-if="$q.screen.gt.sm" v-model="tab">
          <div style="height: 100%" v-for="link in links1" :key="link.text">
            <q-route-tab
              v-if="link.child == undefined"
              :name="link.text"
              :label="link.text"
              :to="link.to"
              exact
            />
            <q-btn-dropdown
              style="height: 100%"
              v-if="link.child != undefined"
              auto-close
              stretch
              flat
              :label="link.text"
            >
              <q-list v-if="link.child != undefined">
                <q-item
                  v-for="childmenu in link.child"
                  :key="childmenu.text"
                  clickable
                  @click="tab = 'movies'"
                >
                  <q-item-section>{{ childmenu.text }}</q-item-section>
                </q-item>
              </q-list>
            </q-btn-dropdown>
          </div>
        </q-tabs>
        <q-space></q-space>
        <div class="q-gutter-sm row items-center no-wrap">
          <q-btn @click="openGithub" dense round flat icon="ion-logo-github"/>
          <q-btn @click="alertAdd = true" dense round flat icon="add_circle_outline"/>
          <q-dialog class="maindia" style="background-color: rgba(0,0,0,0.6);" persistent v-model="alertAdd">
            <q-btn @click="closeUploader" class="text-white absolute-top-right q-mr-lg q-mt-lg" icon="close" v-close-popup flat round dense/>
            <q-card class="container full-width" :style="$q.screen.gt.xs?bigStyle:smallStyle
                          ">
            <q-resize-observer @resize="onResize" />
            <q-card-section style="padding: 0px;" class="full-width full-height">
                  <q-item class="relative-position">
                    <q-item-section align="center">
                      <q-item-label>
                        <q-toolbar-title><span class="text-weight-bold">Create Your New Picture</span></q-toolbar-title>
                      </q-item-label>
                    </q-item-section>
                    <q-btn class="absolute-right" v-if="imageFiles.length!=0" @click="next1" style="top: initial !important;bottom: initial !important;" icon="navigate_next" round dense flat/>
                  </q-item>
                  <q-separator />
                  <q-item style="padding: 0px;height: calc(100% - 49px);">
                    <q-item-section>
                      <q-card-section horizontal style="padding: 0px;" class="toolbar-upload full-height">
                        <q-card-section style="padding: 0px;" :class="alertNext?'col-6':'col-12'">
                          <q-uploader
                            @added="addImageS"
                            @removed="removeImage"
                            class="full-height"
                            style="width:initial!important;max-height:initial!important;"
                            hide-upload-btn=false
                            flat
                            multiple
                            bordered
                          >
                            <template v-slot:header="scope">
                              <div class="no-wrap" align="center">
                                <q-btn class="absolute-center vertical-middle	" v-if="scope.canAddFiles" type="a" icon="add_box" round dense flat>
                                  <q-uploader-add-trigger />
                                  <q-tooltip>Pick Files</q-tooltip>
                                </q-btn>
                              </div>
                            </template>
                          </q-uploader>
                        </q-card-section>
                        <q-card-section v-if="alertNext" style="padding: 0px;" class="col-6">
                          <q-item>
                            <q-item-section avatar>
                              <q-avatar>
                                <img src="https://cdn.quasar.dev/img/boy-avatar.png">
                              </q-avatar>
                            </q-item-section>
                            <q-item-section>
                              <q-item-label>Title</q-item-label>
                              <q-item-label caption>
                                Subhead
                              </q-item-label>
                            </q-item-section>
                            <q-item-section side top>
                            </q-item-section>
                          </q-item>
                          <q-item>
                            <q-item-section style="max-height: 50vh" class="scroll hide-scrollbar">
                            <q-input
                              class="full-width"
                              v-model="text"
                              filled
                              autogrow
                            />
                            </q-item-section>
                          </q-item>
                        </q-card-section>
                      </q-card-section>
                    </q-item-section>
                  </q-item>
                </q-card-section>
            </q-card>
          </q-dialog>
          <q-toggle
            v-model="darkValue"
            color="orange"
          />
          <q-btn-dropdown v-if="$q.screen.gt.sm" stretch flat icon="color_lens">
            <q-list>
              <q-color v-model="hex" no-header class="my-picker" />
            </q-list>
          </q-btn-dropdown>
          <div v-if="$q.screen.gt.sm">
            <q-btn v-if="profile.name" dense flat no-wrap>
              <q-avatar color="white" text-color="white">
                <q-badge color="red" floating>4</q-badge>
                <img :src="profile.avatar == ''?'https://cdn.quasar.dev/logo-v2/svg/logo.svg':profile.avatar" />
              </q-avatar>
              <q-icon name="arrow_drop_down" size="16px" />
              <q-menu auto-close>
                <q-list style="min-width: 100px">
                  <q-item clickable @click="toRepositories(profile.id)" class="GL__menu-link-signed-in">
                    <q-item-section>
                      <div>
                        Sign in as <strong> {{ profile.name }}</strong>
                      </div>
                    </q-item-section>
                  </q-item>
                  <q-separator />
                  <q-item class="GL__menu-link">
                    <q-item-section>
                      <q-btn
                        outline
                        @click="alert = true"
                        color="primary"
                        :label="profile.emoji == ''?'Set status':profile.emoji+profile.emoji_text"
                      >
                        <q-icon v-if="profile.emoji == ''" name="tag_faces" color="blue-9" size="18px" />
                      </q-btn>
                    </q-item-section>
                  </q-item>

                  <q-separator />
                  <!-- <q-item clickable to="/profile" class="GL__menu-link">
                    <q-item-section>Your profile</q-item-section>
                  </q-item> -->
                  <!-- <q-item clickable @click="toRepositories(profile.id)" class="GL__menu-link">
                    <q-item-section>Your repositories</q-item-section>
                  </q-item>
                  <q-separator /> -->
                  <!-- <q-item clickable class="GL__menu-link">
                    <q-item-section>Settings</q-item-section>
                  </q-item> -->
                  <q-item clickable @click="loginOut" class="GL__menu-link">
                    <q-item-section>Sign out</q-item-section>
                  </q-item>
                </q-list>
              </q-menu>
            </q-btn>
            <div v-else>
              <router-link
                v-if="$q.screen.gt.sm"
                class="text-weight-bold q-mr-lg"
                to="/login"
                style="color: #fff"
                >SIGN IN</router-link
              >
              <q-btn
                color="white"
                to="/signup"
                text-color="black"
                push
                label="sign up"
              />
            </div>
          </div>
        </div>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-if="!$q.screen.gt.sm"
      v-model="leftDrawerOpen"
      :breakpoint="500"
      bordered
    >
      <q-scroll-area style="height: calc(100% - 150px); margin-top: 150px">
        <q-list padding>
          <div v-for="link in links1" :key="link.text">
            <q-item v-if="link.child == undefined" :to="link.to" v-ripple clickable>
              <q-item-section avatar>
                <q-icon color="grey" :name="link.icon" />
              </q-item-section>
              <q-item-section>
                <q-item-label>{{ link.text }}</q-item-label>
              </q-item-section>
            </q-item>
            <q-expansion-item
              v-if="link.child != undefined"
              expand-separator
              :icon="link.icon"
              :label="link.text"
            >
             <q-item
                :header-inset-level="1"
                icon="receipt"
                label="Receipts"
              >
              </q-item>
            </q-expansion-item>
          </div>
        </q-list>
      </q-scroll-area>
      <q-img
        class="absolute-top"
        contain
        :src="profile.img == ''?'https://cdn.quasar.dev/img/material.png':profile.img"
        style="height: 150px"
      >
        <div v-if="profile.name" class="absolute-bottom bg-transparent">
          <q-btn round @click="toRepositories(profile.id)">
            <q-avatar color="white">
              <img :src="profile.avatar == ''?'https://cdn.quasar.dev/logo-v2/svg/logo.svg':profile.avatar" />
            </q-avatar>
          </q-btn>
          <div class="text-weight-bold">{{ profile.name }}</div>
          <div>{{ profile.email }}</div>
        </div>
        <div v-else class="absolute-bottom text-subtitle1 text-center">
          <router-link
            class="text-weight-bold q-mr-lg"
            to="/login"
            style="color: #fff"
            >SIGN IN</router-link
          >
          <q-btn
            cover
            color="white"
            to="/signup"
            text-color="black"
            push
            label="sign up"
          />
        </div>
      </q-img>
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
    <q-page-sticky
      elevated
      reveal
      position="bottom-right"
      :scroll-offset="150"
      :offset="[18, 18]"
    >
      <q-fab icon="unfold_more" color="cyan-9" direction="up">
        <q-fab-action icon="comment" color="cyan-9" @click="open('right')" />
        <q-fab-action
          icon="home"
          color="cyan-9"
          @click="$router.push('/12321')"
        />
        <q-fab-action
          icon="keyboard_arrow_up"
          color="cyan-9"
          @click="backToTop"
        />
      </q-fab>
      <q-dialog v-model="dialog" wi persistent :position="position">
        <q-card>
          <q-bar>
            <q-space />

            <q-btn dense flat icon="close" v-close-popup>
              <q-tooltip>Close</q-tooltip>
            </q-btn>
          </q-bar>

          <q-card-section>
            <q-scroll-area style="height: 400px; width: 400px">
              <q-chat-message
                v-for="label in talk"
                :key="label.index"
                :name="label.name"
                :avatar="label.avatar"
                :text="label.text"
                :stamp="label.stamp"
                :sent="label.send"
                :bg-color="label.send != '' ? 'primary' : 'amber-7'"
                :text-color="label.send != '' ? 'white' : ''"
              />
            </q-scroll-area>
          </q-card-section>
          <q-card-section>
            <q-input
              filled
              bottom-slots
              v-model="text"
              label="Label"
              :dense="dense"
            >
              <template v-slot:append>
                <q-btn round dense flat icon="mood">
                  <q-popup-edit
                    max-width="158px"
                    style="padding: 4px 8px"
                    self="top start"
                    cover="false"
                  >
                    <div class="q-gutter-sm" style="margin-top: 0px">
                      <a
                        href="javascript:void(0);"
                        @click="getEmo(index)"
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
              <template v-slot:after>
                <q-btn @click="send()" round dense flat icon="send" />
              </template>
            </q-input>
          </q-card-section>
        </q-card>
      </q-dialog>
      <q-dialog v-model="alert">
        <q-card style="min-width: 350px">
          <q-card-section>
            <div class="text-h6">Edit status</div>
          </q-card-section>
          <q-card-section class="q-pt-none">
             <q-input outlined v-model="profile.emoji_text" label="What's happening?" dense>
              <template v-slot:before>
                <q-btn style="color:rgba(0, 0, 0, 1);"  round dense flat :label="profile.emoji">
                  <q-icon v-if="profile.emoji == ''" name="mood"/>
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
            <q-btn flat label="OK" @click="editStatus()" color="primary" v-close-popup />
          </q-card-actions>
        </q-card>
      </q-dialog>
    </q-page-sticky>
  </q-layout>
</template>

<script lang="ts">
// import EssentialLink from 'components/EssentialLink.vue'
import emoji from '../css/emoji.json';
import {
  onBeforeMount,
  defineComponent,
  ref,
  reactive,
  toRefs,
  onMounted,
  watch
} from 'vue';
import { getProfileMe, editEmoji, logout } from '../api/test/index';
import { LocalStorage, Notify, Dark, Cookies } from 'quasar';
import { useRouter } from 'vue-router';
import { Screen } from 'quasar';

export default defineComponent({
  name: 'MainLayout',

  components: {},

  setup() {
    let darkdata = reactive({
      darkValue:false,
    })
    let data = reactive({
      imageFiles: [],
      smallStyle: 'min-width: 90%;min-height: 450px;max-width: 850px;max-height: 850px;height: 55vw !important;width: 55vw !important;',
      bigStyle: 'min-width: 450px;min-height: 450px;max-width: 850px;max-height: 80%;height: 45vw !important;width: 30vw !important;',
      alertNext: false,
      alertAdd: false,
      alert: false,
      hex: 'rgb(99, 24, 105)',
      tab: 'Home',
      profile: {
        name: '',
        avatar: '',
        email: '',
        img: '',
        id: 0,
        emoji: '',
        emoji_text: ''
      },
      links1: [
        { icon: 'home', text: 'Home', to: '/' },
        { icon: 'whatshot', text: 'ArtList', to: '/imglist' },
        { icon: 'whatshot', text: 'Setting', to: '/setting/article' },
        // {
        //   icon: 'subscriptions',
        //   text: 'Subscriptions',
        //   child: [
        //     { icon: 'subscriptions', text: 'Subscriptions', to: '/' },
        //     { icon: 'whatshot', text: 'Trending', to: '/' },
        //   ],
        // },
      ],
      talk: [
        {
          name: 'me',
          avatar: 'https://cdn.quasar.dev/img/avatar1.jpg',
          text: ['hi'],
          stamp: '7 minutes ago',
          send: '',
        },
        {
          name: 'Jane',
          avatar: 'https://cdn.quasar.dev/img/avatar5.jpg',
          text: ['doing fine, how r you?'],
          stamp: '4 minutes ago',
          send: 'send',
        },
      ],
      faceList: [],
      text: '',
      icon: false,
      bar: false,
      bar2: false,
      toolbar: false,
      onload: false,
      dialog: false,
      persistent: false,
      position: 'top',
    });
    const router = useRouter() as any;
    let leftDrawerOpen = ref(false);
    const method = {
      backToTop() {
        let timer: any = '';
        const gotoTop = () => {
          let currentPosition =
            document.documentElement.scrollTop || document.body.scrollTop;

          currentPosition -= 100;
          if (currentPosition > 0) {
            window.scrollTo(0, currentPosition);
          } else {
            window.scrollTo(0, 0);
            clearInterval(timer);
            timer = null;
          }
        };
        timer = setInterval(gotoTop, 15);
      },
      open(position): void {
        data.position = position;
        data.dialog = true;
      },
      addImageS(files: Array<any>): void {
        data.imageFiles = data.imageFiles.concat(files)
        console.log(data.imageFiles)
      },
      removeImage(files: Array<any>): void {
        data.imageFiles.splice(data.imageFiles.indexOf(files[0]),1)
        console.log(data.imageFiles)
        data.alertNext = false
        data.bigStyle ='min-width: 450px;min-height: 450px;max-width: 850px;max-height: 80%;height: 45vw !important;width: 30vw !important;'
      },
      closeUploader(): void {
        data.imageFiles = []
        data.alertNext = false
        data.bigStyle ='min-width: 450px;min-height: 450px;max-width: 850px;max-height: 80%;height: 45vw !important;width: 30vw !important;'
      },
      toRepositories(id: number): void {
        router.push(`/repository/${id}`)
      },
      openGithub(): void {
        window.open('https://github.com/hints0816/artblog','_blank')
      },
      getEmo(index: number): void {
        const face = data.faceList[index] as string;
        data.text = data.text + face;
      },
      next1(): void {
        data.alertNext = true
        data.bigStyle = 'min-width: 450px;min-height: 450px;max-width: 950px;max-height: 80%;height: 45vw !important;width: 70vw !important;'
      },
      setEmojiStatus(index: number): void {
        const face = data.faceList[index] as string;
        data.profile.emoji = face;
      },
      async editStatus(): Promise<any> {
        const params = {
          emoji: data.profile.emoji,
          emoji_text: data.profile.emoji_text
        };
        let res = await editEmoji(params) as any;
        if (res.status == 200) {
          router.go(0)
          return Notify.create({
            message: 'Edit Successful',
            type: 'positive',
            position: 'top',
            timeout: 2000
          })
        }
      },
      async getProfileMe(): Promise<any> {
        let res = (await getProfileMe()) as any;
        if (res.status == 200) {
          data.profile.avatar = res.data.avatar;
          data.profile.emoji = res.data.emoji;
          data.profile.emoji_text = res.data.emoji_text;
          data.profile.img = res.data.img;
          data.profile.id = res.data.id;
          data.profile.name = res.data.name;
          data.profile.email = res.data.email;
          LocalStorage.set('avatar', res.data.avatar);
        } else {
          data.links1.splice(2,1)
        }
      },
      async loginOut(): Promise<any>  {
        LocalStorage.remove('token')
        LocalStorage.set('logged_in', 'no');
        let res = await logout() as any;
        if(res.status == 302) {
          window.location.href = '/';
        }
      },
      send(): void {
        const params = {
          name: 'me',
          avatar: 'https://cdn.quasar.dev/img/avatar1.jpg',
          text: [data.text],
          stamp: '4 minutes ago',
          send: '',
        };
        data.talk.push(params);
        data.text = '';
      },
    };
    onBeforeMount(async () => {
      await method.getProfileMe();
    });
    onMounted(() => {
      emoji.forEach((element) => {
        data.faceList.push(element.char);
      });
    });
    watch(darkdata,(newVal,oldVal)=>{
      Dark.toggle()
    })
     const style = ref({ width: '200px', height: '200px' })
    const report = ref(null)

    return {
      style,
      report,

      onResize (size) {
        report.value = size
        // {
        //   width: 20 // width of container (in px)
        //   height: 50 // height of container (in px)
        // }
      },

      setRandomSize () {
        style.value = {
          width: String(Math.floor(100 + Math.random() * 200)) + 'px',
          height: String(Math.floor(100 + Math.random() * 200)) + 'px'
        }
      },
      leftDrawerOpen,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value;
      },
      ...toRefs(data),
      ...toRefs(darkdata),
      ...method,
    };
  },
});
</script>
<style lang="sass">
.container
  transition: width .3s, height .3s
a 
  text-decoration: none
.toolbar-upload 
  .q-uploader__header
    background-color: initial !important 
    height: 100%
.toolbar-upload 
  .q-uploader__header
    .q-btn-item
      z-index: 1000
      color: #1976D2
      font-size: 30px
.toolbar-upload 
  .q-uploader__file-header
    background: initial !important
.toolbar-upload 
  .q-uploader__file-header-content
    .q-uploader__title
      display: none
    .q-uploader__subtitle
      display: none
.toolbar-upload      
  .q-uploader__list
    padding: 0px
    position: absolute
    width: 100%
    height: 100%
.toolbar-upload          
  .q-uploader__subtitle
    display: none
.toolbar-upload        
  .q-uploader__file
    height: 100%
.maindia .no-pointer-events
  pointer-events: initial !important
</style>