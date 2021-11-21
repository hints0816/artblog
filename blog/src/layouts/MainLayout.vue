<template>
  <q-layout>
    <q-header elevated>
      <q-toolbar>
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
            <router-link to="/" style="color: #fff;">blogName</router-link>
          </q-toolbar-title>
         <q-tabs v-if="$q.screen.gt.sm" v-model="tab">
          <q-tab name="videos" label="Videos" />
          <q-tab name="articles" label="Articles" />
          <q-btn-dropdown auto-close stretch flat label="More...">
          <q-list>
            <q-item clickable @click="tab = 'movies'">
              <q-item-section>Movies</q-item-section>
            </q-item>

            <q-item clickable @click="tab = 'photos'">
              <q-item-section>Photos</q-item-section>
            </q-item>
          </q-list>
        </q-btn-dropdown>
        </q-tabs>
        <q-space></q-space>
        <div class="q-gutter-sm row items-center no-wrap">
        <iframe src="https://ghbtns.com/github-btn.html?user=hints0816&repo=gofile3&type=star&count=true&size=large" frameborder="0" scrolling="0" width="170" height="30" title="GitHub"></iframe>  
        <q-btn v-if="avatar" dense flat no-wrap>
            <q-avatar color="orange" text-color="white">
              <q-badge color="red" floating>4</q-badge>
              <img :src="avatar">
            </q-avatar>
            <q-icon name="arrow_drop_down" size="16px" />

            <q-menu auto-close>
              <q-list dense=false>
                <q-item class="GL__menu-link-signed-in">
                  <q-item-section>
                    <div>Signed in as <strong>Mary</strong></div>
                  </q-item-section>
                </q-item>
                <q-separator />
                <q-item clickable class="GL__menu-link-status">
                  <q-item-section>
                    <div>
                      <q-icon name="tag_faces" color="blue-9" size="18px" />
                      Set your status
                    </div>
                  </q-item-section>
                </q-item>
                <q-separator />
                <q-item clickable class="GL__menu-link">
                  <q-item-section>Your profile</q-item-section>
                </q-item>
                <q-item clickable to="/edit" class="GL__menu-link">
                  <q-item-section>Your repositories</q-item-section>
                </q-item>
                <q-separator />
                <q-item clickable class="GL__menu-link">
                  <q-item-section>Help</q-item-section>
                </q-item>
                <q-item clickable class="GL__menu-link">
                  <q-item-section>Settings</q-item-section>
                </q-item>
                <q-item clickable class="GL__menu-link">
                  <q-item-section>Sign out</q-item-section>
                </q-item>
              </q-list>
            </q-menu>
        </q-btn>
        <div v-else>
          <router-link class="text-weight-bold" to="/" style="color: #fff;">blogName</router-link>
          <q-btn outline label="sign up" />
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
      <q-scroll-area class="fit">
        <q-list padding>
          <q-item v-for="link in links1" :key="link.text" v-ripple clickable>
            <q-item-section avatar>
              <q-icon color="grey" :name="link.icon" />
            </q-item-section>
            <q-item-section>
              <q-item-label>{{ link.text }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-separator class="q-my-md" />

          <q-item v-for="link in links2" :key="link.text" v-ripple clickable>
            <q-item-section avatar>
              <q-icon color="grey" :name="link.icon" />
            </q-item-section>
            <q-item-section>
              <q-item-label>{{ link.text }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-separator class="q-mt-md q-mb-xs" />

          <q-item-label header class="text-weight-bold text-uppercase">
            More from Youtube
          </q-item-label>

          <q-item v-for="link in links3" :key="link.text" v-ripple clickable>
            <q-item-section avatar>
              <q-icon color="grey" :name="link.icon" />
            </q-item-section>
            <q-item-section>
              <q-item-label>{{ link.text }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-separator class="q-my-md" />

          <q-item v-for="link in links4" :key="link.text" v-ripple clickable>
            <q-item-section avatar>
              <q-icon color="grey" :name="link.icon" />
            </q-item-section>
            <q-item-section>
              <q-item-label>{{ link.text }}</q-item-label>
            </q-item-section>
          </q-item>

          <q-separator class="q-mt-md q-mb-lg" />

          <div class="q-px-md text-grey-9">
            <div class="row items-center q-gutter-x-sm q-gutter-y-xs">
              <a
                v-for="button in buttons1"
                :key="button.text"
                class="YL__drawer-footer-link"
                href="javascript:void(0)"
              >
                {{ button.text }}
              </a>
            </div>
          </div>
          <div class="q-py-md q-px-md text-grey-9">
            <div class="row items-center q-gutter-x-sm q-gutter-y-xs">
              <a
                v-for="button in buttons2"
                :key="button.text"
                class="YL__drawer-footer-link"
                href="javascript:void(0)"
              >
                {{ button.text }}
              </a>
            </div>
          </div>
        </q-list>
      </q-scroll-area>
    </q-drawer>

    <q-page-container> 
      <router-view />
    </q-page-container>
    <q-page-sticky elevated reveal position="bottom-right" :scroll-offset="150" :offset="[18, 18]">
      <q-fab
        icon="unfold_more"
        color="cyan-9"
        direction="up"
      >
        <q-fab-action icon="comment" color="cyan-9" @click="open('right')" />
        <q-fab-action icon="home" color="cyan-9" @click="$router.push('/12321');" />
        <q-fab-action icon="keyboard_arrow_up" color="cyan-9" @click="backToTop" />
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
            <q-scroll-area
              style="height: 400px; width: 400px;">
              <q-chat-message
                v-for="label in talk"
                :key="label.index"
                :name="label.name"
                :avatar="label.avatar"
                :text="label.text"
                :stamp="label.stamp"
                :sent="label.send"
                :bg-color="label.send!=''?'primary':'amber-7'"
                :text-color="label.send!=''?'white':''"
              />
            </q-scroll-area>
        </q-card-section>
        <q-card-section>
          <q-input filled bottom-slots v-model="text" label="Label" :dense="dense">
            <template v-slot:append>
              <q-btn round dense flat icon="mood">
                <q-popup-edit max-width="158px" style="padding: 4px 8px;" self="top start" cover="false">
                  <div class="q-gutter-sm" style="margin-top:0px">
                    <a href="javascript:void(0);" @click="getEmo(index)" style="text-decoration:none" v-for="(item,index) in faceList" :key="index" class="emotionItem">{{item}}</a>
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
    </q-page-sticky>
  </q-layout>
</template>

<script lang="ts">
// import EssentialLink from 'components/EssentialLink.vue'
import emoji from '../css/emoji.json'
import { onBeforeMount, defineComponent, ref, getCurrentInstance, reactive, toRefs, onMounted } from 'vue'
import { getProfileMe } from '../api/test/index'
export default defineComponent({
  name: 'MainLayout',

  components: {
  },

  setup () {
    let data = reactive({
      links1: [
        { icon: 'home', text: 'Home' },
        { icon: 'whatshot', text: 'Trending' },
        { icon: 'subscriptions', text: 'Subscriptions' }
      ],
      links2: [
        { icon: 'folder', text: 'Library' },
        { icon: 'restore', text: 'History' },
        { icon: 'watch_later', text: 'Watch later' },
        { icon: 'thumb_up_alt', text: 'Liked videos' }
      ],
      links3: [
        { icon: 'local_movies', text: 'YouTube Premium' },
        { icon: 'local_movies', text: 'Movies & Shows' },
        { icon: 'videogame_asset', text: 'Gaming' },
        { icon: 'live_tv', text: 'Live' }
      ],
      links4: [
        { icon: 'settings', text: 'Settings' },
        { icon: 'flag', text: 'Report history' },
        { icon: 'help', text: 'Help' },
        { icon: 'feedback', text: 'Send feedback' }
      ],
      avatar: '',
      talk:[
        {
          name: 'me',
          avatar: 'https://cdn.quasar.dev/img/avatar1.jpg',
          text: [
            'hi'
          ],
          stamp: '7 minutes ago',
          send: ''
        },
         {
          name: 'Jane',
          avatar: 'https://cdn.quasar.dev/img/avatar5.jpg',
          text: [
            'doing fine, how r you?'
          ],
          stamp: '4 minutes ago',
          send: 'send'
        }
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
      position: 'top'
    })
    const {ctx} = getCurrentInstance() as any
    let leftDrawerOpen = ref(false)
    const method = {
      backToTop() {
        console.log(ctx)
        let timer: any = '';
        const gotoTop = () => {
            let currentPosition = document.documentElement.scrollTop || document.body.scrollTop;
            
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
        data.position = position
        data.dialog = true
      },
      getEmo(index): void {
        const face = data.faceList[index] as string
        data.text = data.text + face
      },
      async getProfileMe():Promise<any> {
        let res  = await getProfileMe() as any
        if(res.status == 200){
          data.avatar = res.data.avatar
        }
      },
      send(): void {
        const params = {
          name: 'me',
          avatar: 'https://cdn.quasar.dev/img/avatar1.jpg',
          text: [
             data.text
          ],
          stamp: '4 minutes ago',
          send: ''
        }
        data.talk.push(params)
        data.text = ''
      },
    }
    onBeforeMount(async()=>{
      await method.getProfileMe()
    })
    onMounted(()=>{
           emoji.forEach(element => {
             data.faceList.push(element.char);
         });                   
      })
    return {
      leftDrawerOpen,
      toggleLeftDrawer () {
        leftDrawerOpen.value = !leftDrawerOpen.value
      },
      ...toRefs(data),
      ...method
    }
  },
  methods: {
    getUserInfo():void {
      // axiosInstance.get(`https://api.github.com/users/${this.$store.getters.username}`)
      //   .then((res) => {
      //     this.$set(this, 'user', res.data);
      //   });
    },
    backToTop():void {
      // let timer;
      // const gotoTop = () => {
      //   let currentPosition = document.documentElement.scrollTop || document.body.scrollTop;
      //   currentPosition -= 100;
      //   if (currentPosition > 0) {
      //     window.scrollTo(0, currentPosition);
      //   } else {
      //     window.scrollTo(0, 0);
      //     clearInterval(timer);
      //     timer = null;
      //   }
      // };
      // timer = setInterval(gotoTop, 1);
    }
  },
})
</script>

<style lang="scss" scoped>
a{
    text-decoration:none
}
</style>