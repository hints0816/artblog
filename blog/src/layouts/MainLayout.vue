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
          Quasar App
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
        <q-btn round flat>
          <q-avatar color="orange" text-color="white">
            <q-badge color="red" floating>4</q-badge>
            J
          </q-avatar>
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
