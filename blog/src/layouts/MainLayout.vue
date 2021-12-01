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
          <router-link to="/" style="color: #fff">blogName</router-link>
        </q-toolbar-title>
        <q-tabs v-if="$q.screen.gt.sm" v-model="tab">
          <div style="height: 100%" v-for="link in links1" :key="link.text">
            <q-tab
              v-if="link.child == undefined"
              :name="link.text"
              :label="link.text"
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
          <iframe
            src="https://ghbtns.com/github-btn.html?user=hints0816&repo=gofile3&type=star&count=true&size=large"
            frameborder="0"
            scrolling="0"
            width="170"
            height="30"
            title="GitHub"
          ></iframe>
          <q-btn-dropdown v-if="$q.screen.gt.sm" stretch flat icon="color_lens">
            <q-list>
              <q-color v-model="hex" no-header class="my-picker" />
            </q-list>
          </q-btn-dropdown>
          <div v-if="$q.screen.gt.sm">
            <q-btn v-if="profile.name" dense flat no-wrap>
              <q-avatar color="orange" text-color="white">
                <q-badge color="red" floating>4</q-badge>
                <img :src="profile.avatar" />
              </q-avatar>
              <q-icon name="arrow_drop_down" size="16px" />
              <q-menu auto-close>
                <q-list style="min-width: 100px">
                  <q-item class="GL__menu-link-signed-in">
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
                        label="Set status"
                      >
                        <q-icon name="tag_faces" color="blue-9" size="18px" />
                      </q-btn>
                    </q-item-section>
                  </q-item>

                  <q-separator />
                  <q-item clickable to="/profile" class="GL__menu-link">
                    <q-item-section>Your profile</q-item-section>
                  </q-item>
                  <q-item clickable @click="toRepositories(profile.id)" class="GL__menu-link">
                    <q-item-section>Your repositories</q-item-section>
                  </q-item>
                  <q-separator />
                  <q-item clickable class="GL__menu-link">
                    <q-item-section>Settings</q-item-section>
                  </q-item>
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
          <q-item v-for="link in links1" :key="link.text" v-ripple clickable>
            <q-item-section avatar>
              <q-icon color="grey" :name="link.icon" />
            </q-item-section>
            <q-item-section>
              <q-item-label>{{ link.text }}</q-item-label>
            </q-item-section>
          </q-item>
        </q-list>
      </q-scroll-area>
      <q-img
        class="absolute-top"
        contain
        src="https://cdn.quasar.dev/img/material.png"
        style="height: 150px"
      >
        <div v-if="profile.name" class="absolute-bottom bg-transparent">
          <q-btn round to="/profile">
            <q-avatar>
              <img :src="profile.avatar" />
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
        <q-card>
          <q-card-section>
            <div class="text-h6">Alert</div>
          </q-card-section>

          <q-card-section class="q-pt-none">
             <q-input bottom-slots v-model="text" label="Label" counter maxlength="12" :dense="dense">
              <template v-slot:before>
                <q-btn round dense flat icon="send" />
              </template>
              <template v-slot:append>
                <q-icon v-if="text !== ''" name="close" @click="text = ''" class="cursor-pointer" />
                <q-icon name="schedule" />
              </template>
            </q-input>
          </q-card-section>

          <q-card-actions align="right">
            <q-btn flat label="OK" color="primary" v-close-popup />
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
  getCurrentInstance,
  reactive,
  toRefs,
  onMounted,
} from 'vue';
import { getProfileMe } from '../api/test/index';
import { LocalStorage } from 'quasar';
import { useRouter } from 'vue-router';

export default defineComponent({
  name: 'MainLayout',

  components: {},

  setup() {
    let data = reactive({
      alert: false,
      hex: '',
      tab: 'Home',
      profile: {
        name: '',
        avatar: '',
        email: '',
        id: 0
      },
      links1: [
        { icon: 'home', text: 'Home' },
        { icon: 'whatshot', text: 'Trending' },
        {
          icon: 'subscriptions',
          text: 'Subscriptions',
          child: [
            { icon: 'subscriptions', text: 'Subscriptions' },
            { icon: 'whatshot', text: 'Trending' },
          ],
        },
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
    const { ctx } = getCurrentInstance() as any;
    let leftDrawerOpen = ref(false);
    const method = {
      backToTop() {
        console.log(ctx);
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
      toRepositories(id: number): void {
        router.push(`/repository/${id}`)
      },
      getEmo(index): void {
        const face = data.faceList[index] as string;
        data.text = data.text + face;
      },
      async getProfileMe(): Promise<any> {
        let res = (await getProfileMe()) as any;
        if (res.status == 200) {
          data.profile.avatar = res.data.avatar;
          data.profile.id = res.data.id;
          data.profile.name = res.data.name;
          data.profile.email = res.data.email;
          LocalStorage.set('avatar', res.data.avatar);
        } else {
        }
      },
      loginOut(): void {
        LocalStorage.remove('token');
        LocalStorage.set('logged_in', 'no');
        window.location.href = '/';
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
      console.log(data.links1[1].child);
      emoji.forEach((element) => {
        data.faceList.push(element.char);
      });
    });
    return {
      leftDrawerOpen,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value;
      },
      ...toRefs(data),
      ...method,
    };
  },
});
</script>

<style scoped>
/* /deep/ .q-btn__content {
  display: flex;
  flex-direction: row;
  width: 100px;
} */
</style>
<style lang="scss" scoped>
a {
  text-decoration: none;
}
</style>