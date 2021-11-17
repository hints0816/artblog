<template>
  <q-layout>
    <q-header elevated>
      <q-toolbar>
        <!-- <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          @click="toggleLeftDrawer"
        /> -->

        <q-toolbar-title shrink class="text-weight-bold">
            <router-link to="/" style="color: #fff;">blogName</router-link>
          </q-toolbar-title>
         <q-tabs v-model="tab">
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
        <q-btn dense flat no-wrap>
            <q-avatar color="orange" text-color="white">
              <q-badge color="red" floating>4</q-badge>
              J
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
        </div>
      </q-toolbar>
    </q-header>

    <!-- <q-drawer
      v-model="leftDrawerOpen"
      show-if-above
      bordered
    >
      <q-list>
        <q-item-label
          header
        >
          Essential Links
        </q-item-label>

        <EssentialLink
          v-for="link in essentialLinks"
          :key="link.title"
          v-bind="link"
        />
      </q-list>
    </q-drawer> -->

    <q-page-container> 
      <router-view />
    </q-page-container>
    <q-page-sticky elevated reveal position="bottom-right" :scroll-offset="150" :offset="[18, 18]">
      <q-fab
        icon="unfold_more"
        color="cyan-9"
        direction="up"
      >
        <q-fab-action icon="home" color="cyan-9" @click="$router.push('/12321');" />
        <q-fab-action icon="keyboard_arrow_up" color="cyan-9" @click="backToTop" />
      </q-fab>
    </q-page-sticky>
  </q-layout>
</template>

<script lang="ts">
// import EssentialLink from 'components/EssentialLink.vue'

const linksList = [
  {
    title: 'Docs',
    caption: 'quasar.dev',
    icon: 'school',
    link: 'https://quasar.dev'
  },
  {
    title: 'Github',
    caption: 'github.com/quasarframework',
    icon: 'code',
    link: 'https://github.com/quasarframework'
  },
  {
    title: 'Discord Chat Channel',
    caption: 'chat.quasar.dev',
    icon: 'chat',
    link: 'https://chat.quasar.dev'
  },
  {
    title: 'Forum',
    caption: 'forum.quasar.dev',
    icon: 'record_voice_over',
    link: 'https://forum.quasar.dev'
  },
  {
    title: 'Twitter',
    caption: '@quasarframework',
    icon: 'rss_feed',
    link: 'https://twitter.quasar.dev'
  },
  {
    title: 'Facebook',
    caption: '@QuasarFramework',
    icon: 'public',
    link: 'https://facebook.quasar.dev'
  },
  {
    title: 'Quasar Awesome',
    caption: 'Community Quasar projects',
    icon: 'favorite',
    link: 'https://awesome.quasar.dev'
  }
];

import { defineComponent, ref, getCurrentInstance } from 'vue'

export default defineComponent({
  name: 'MainLayout',

  components: {
    // EssentialLink
  },

  setup () {
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
    }
    return {
      essentialLinks: linksList,
      leftDrawerOpen,
      toggleLeftDrawer () {
        leftDrawerOpen.value = !leftDrawerOpen.value
      },
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
