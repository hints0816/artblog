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
          <q-btn @click="openGithub" dense round flat icon="ion-logo-github" />
          <q-btn
            @click="alertAdd = true"
            dense
            round
            flat
            icon="add_circle_outline"
          />
          <q-dialog
            class="maindia"
            style="background-color: rgba(0, 0, 0, 0.6)"
            persistent
            v-model="alertAdd"
          >
            <q-btn
              @click="closeUploader"
              class="text-white absolute-top-right q-mr-lg q-mt-lg"
              icon="close"
              v-close-popup
              flat
              round
              dense
            />
            <q-card
              class="container full-width"
              :style="$q.screen.gt.xs ? bigStyle : smallStyle"
            >
              <q-resize-observer @resize="onResize" />
              <q-card-section
                style="padding: 0px"
                class="full-width full-height"
              >
                <q-item class="relative-position">
                  <q-item-section align="center">
                    <q-item-label>
                      <q-toolbar-title
                        ><span class="text-weight-bold"
                          >Create Your New Picture</span
                        ></q-toolbar-title
                      >
                    </q-item-label>
                  </q-item-section>
                  <q-btn
                    v-if="alertNext"
                    class="absolute-left"
                    @click="back"
                    style="top: initial !important; bottom: initial !important"
                    icon="arrow_back"
                    round
                    dense
                    flat
                  />
                  <q-btn
                    class="absolute-right"
                    v-if="imageFiles.length != 0 && !alertNext"
                    @click="next1"
                    style="top: initial !important; bottom: initial !important"
                    icon="navigate_next"
                    round
                    dense
                    flat
                  />
                </q-item>
                <q-separator />
                <q-item style="padding: 0px; height: calc(100% - 49px)">
                  <q-item-section>
                    <q-card-section
                      horizontal
                      style="padding: 0px"
                      class="toolbar-upload full-height"
                    >
                      <q-card-section
                        id="imgCard"
                        @click="addTag($event)"
                        :style="{
                          position: 'relative',
                          padding: '0px',
                          cursor: alertNext ? 'crosshair' : 'default',
                        }"
                        :class="alertNext ? 'col-6' : 'col-12'"
                      >
                        <q-chip
                          removable
                          v-for="(item, index) in chipPositions"
                          :key="index"
                          :id="'tagChip' + index"
                          clickable
                          @remove="removeTag(index)"
                          :style="{
                            'z-index': 99,
                            position: 'absolute',
                            top: item.chiptop + 'px',
                            left: item.chipleft + 'px',
                          }"
                          dense
                          @mousedown="moveFabIndex(index)"
                          v-touch-pan.prevent.mouse="moveFab"
                          icon="event"
                          >{{item.chipText}}</q-chip
                        >
                        <q-uploader
                          @added="addImageS"
                          @removed="removeImage"
                          class="full-height"
                          :style="{
                            'width': 'initial !important',
                            'max-height': 'initial !important',
                            'display': option.img == ''?'':'none'
                            }
                          "
                          hide-upload-btn="false"
                          flat
                          multiple
                          bordered
                        >
                          <template v-slot:header="scope">
                            <div class="no-wrap" align="center">
                              <q-btn
                                :style="
                                  imageFiles.length != 0
                                    ? 'top:-26px;left:22px'
                                    : ''
                                "
                                class="absolute-center vertical-middle"
                                v-if="scope.canAddFiles"
                                type="a"
                                icon="add_box"
                                round
                                dense
                                flat
                              >
                                <q-uploader-add-trigger />
                                <q-tooltip>Pick Files</q-tooltip>
                              </q-btn>
                            </div>
                          </template>
                        </q-uploader>
                        <vue-cropper
                          ref="cropper"
                          v-if="option.img!='' && !alertNext" 
                          :key="cropperkey" 
                          :img="option.img"
                          :outputSize="option.size"
                          :outputType="option.outputType"
                          :full="option.full"
                          :info="option.info"
                          :canScale="option.canScale"
                          :autoCrop="option.autoCrop"
                          :fixedBox="option.fixedBox"
                          autoCropWidth="300"
                          autoCropHeight="300"
                        ></vue-cropper>
                        <q-img
                          v-if="alertNext"
                          :src="imgdisurl"
                          :ratio="1"
                        />
                      </q-card-section>
                      <q-dialog v-model="alertTag">
                        <q-card style="min-width: 350px">
                          <q-card-section>
                            <div class="text-h6">Edit Tag</div>
                          </q-card-section>
                          <q-card-section class="q-pt-none">
                            <q-input
                              outlined
                              v-model="tagText"
                              label="Set Something at this Tag"
                              dense
                            >
                              <template v-slot:before>
                                <q-btn
                                  style="color: rgba(0, 0, 0, 1)"
                                  round
                                  dense
                                  flat
                                  :label="profile.emoji"
                                >
                                  <q-icon
                                    v-if="profile.emoji == ''"
                                    name="mood"
                                  />
                                  <q-popup-edit
                                    max-width="320px"
                                    style="padding: 4px 8px"
                                    self="top start"
                                    cover="false"
                                  >
                                    <div
                                      class="q-gutter-sm"
                                      style="margin-top: 0px"
                                    >
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
                              @click="editTag()"
                              color="primary"
                              v-close-popup
                            />
                          </q-card-actions>
                        </q-card>
                      </q-dialog>
                      <q-card-section
                        v-if="alertNext"
                        style="padding: 0px"
                        class="col-6"
                      >
                        <q-item>
                          <q-item-section avatar>
                            <q-avatar>
                              <img
                                :src="profile.avatar"
                              />
                            </q-avatar>
                          </q-item-section>
                          <q-item-section>
                            <q-item-label>{{ profile.name }}</q-item-label>
                            <q-item-label caption> Subhead </q-item-label>
                          </q-item-section>
                          <q-item-section side top> </q-item-section>
                        </q-item>
                        <q-item>
                          <q-item-section
                            style="max-height: 45vh"
                            class="scroll hide-scrollbar"
                          >
                            <q-input
                              style="height: 100%"
                              class="full-width"
                              v-model="arttext"
                              borderless
                              autogrow
                            />
                          </q-item-section>
                        </q-item>
                        <q-item class="full-width" style="position:absolute;bottom:0px;">
                          <q-btn round flat icon="mood">
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
                          <q-btn style="position: absolute; right:5px;" push color="primary" @click="postArt()" label="POST">
                          </q-btn>
                        </q-item>
                        <q-item style="position: absolute; bottom: 0px">
                          <q-item-section> </q-item-section>
                        </q-item>
                      </q-card-section>
                    </q-card-section>
                  </q-item-section>
                </q-item>
              </q-card-section>
            </q-card>
          </q-dialog>
          <q-toggle v-model="darkValue" color="orange" />
          <q-btn-dropdown v-if="$q.screen.gt.sm" stretch flat icon="color_lens">
            <q-list>
              <q-color v-model="hex" no-header class="my-picker" />
            </q-list>
          </q-btn-dropdown>
          <div v-if="$q.screen.gt.sm">
            <q-btn v-if="profile.name" dense flat no-wrap>
              <q-avatar color="white" text-color="white">
                <q-badge color="red" floating>4</q-badge>
                <img
                  :src="
                    profile.avatar == ''
                      ? 'https://cdn.quasar.dev/logo-v2/svg/logo.svg'
                      : profile.avatar
                  "
                />
              </q-avatar>
              <q-icon name="arrow_drop_down" size="16px" />
              <q-menu auto-close>
                <q-list style="min-width: 100px">
                  <q-item
                    clickable
                    @click="toRepositories(profile.id)"
                    class="GL__menu-link-signed-in"
                  >
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
                        :label="
                          profile.emoji == ''
                            ? 'Set status'
                            : profile.emoji + profile.emoji_text
                        "
                      >
                        <q-icon
                          v-if="profile.emoji == ''"
                          name="tag_faces"
                          color="blue-9"
                          size="18px"
                        />
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
            <q-item
              v-if="link.child == undefined"
              :to="link.to"
              v-ripple
              clickable
            >
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
              <q-item :header-inset-level="1" icon="receipt" label="Receipts">
              </q-item>
            </q-expansion-item>
          </div>
        </q-list>
      </q-scroll-area>
      <q-img
        class="absolute-top"
        contain
        :src="
          profile.img == ''
            ? 'https://cdn.quasar.dev/img/material.png'
            : profile.img
        "
        style="height: 150px"
      >
        <div v-if="profile.name" class="absolute-bottom bg-transparent">
          <q-btn round @click="toRepositories(profile.id)">
            <q-avatar color="white">
              <img
                :src="
                  profile.avatar == ''
                    ? 'https://cdn.quasar.dev/logo-v2/svg/logo.svg'
                    : profile.avatar
                "
              />
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
                  :label="profile.emoji"
                >
                  <q-icon v-if="profile.emoji == ''" name="mood" />
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
  watch,
} from 'vue';
import { getProfileMe, editEmoji, logout, postArt, uploadImage } from '../api/test/index';
import { LocalStorage, Notify, Dark, Cookies } from 'quasar';
import { useRouter } from 'vue-router';
import { dom } from 'quasar';

export default defineComponent({
  name: 'MainLayout',

  components: {},

  setup() {
    let darkdata = reactive({
      darkValue: false,
    });
    let data = reactive({
      option: {
        img: '', // 原图文件
        size: 1,
        outputType: 'jpeg',
        info: true,
        full: true,  
        canScale: true,
        autoCrop: true,
        fixedBox: true,
      },
      tagText: '',
      tagEvent: null,
      moveIndex: 0,
      alertTagEdit: false,
      alertTag: false,
      chipPositions: [],
      chiptop: 0,
      chipleft: 0,
      imageFiles: [],
      smallStyle:
        'min-width: 90%;min-height: 450px;max-width: 850px;max-height: 850px;height: 55vw !important;width: 55vw !important;',
      bigStyle:
        'min-width: 450px;min-height: 450px;max-width: 850px;max-height: 80%;height: 45vw !important;width: 30vw !important;',
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
        emoji_text: '',
      },
      arttext: '',
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
      cropperkey: 1,
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
      imgdisurl: ''
    });
    const router = useRouter() as any;
    let leftDrawerOpen = ref(false);
    const cropper = ref(null);
    const method = {
      moveFabIndex(index: number) {
        data.moveIndex = index;
      },
      moveFab(ev) {
        console.log(data.moveIndex);
        const h1 = dom.height(
          document.getElementById('tagChip' + data.moveIndex.toString())
        );
        const w1 = dom.width(
          document.getElementById('tagChip' + data.moveIndex.toString())
        );
        const h2 = dom.height(document.getElementById('imgCard'));
        const w2 = dom.width(document.getElementById('imgCard'));
        // console.log(data.chipPositions[0].chiptop)
        if (data.chipPositions[data.moveIndex].chiptop - -ev.delta.y > 0) {
          if (
            data.chipPositions[data.moveIndex].chiptop - -ev.delta.y <
            h2 - h1 - 7
          ) {
            data.chipPositions[data.moveIndex].chiptop =
              data.chipPositions[data.moveIndex].chiptop - -ev.delta.y;
          } else {
            data.chipPositions[data.moveIndex].chiptop = h2 - h1 - 7;
          }
        } else {
          data.chipPositions[data.moveIndex].chiptop = 0;
        }
        if (data.chipPositions[data.moveIndex].chipleft - -ev.delta.x > 0) {
          if (
            data.chipPositions[data.moveIndex].chipleft - -ev.delta.x <
            w2 - w1 - 7
          ) {
            data.chipPositions[data.moveIndex].chipleft =
              data.chipPositions[data.moveIndex].chipleft - -ev.delta.x;
          } else {
            data.chipPositions[data.moveIndex].chipleft = w2 - w1 - 7;
          }
        } else {
          data.chipPositions[data.moveIndex].chipleft = 0;
        }
      },
      removeTag(index: number) {
        data.chipPositions.splice(index, 1);
      },
      addTag(e) {
        if (data.alertNext) {
          data.alertTag = true;
          data.chipleft = e.offsetX;
          data.chiptop = e.offsetY;
        }
      },
      editTag() {
        if (data.alertNext) {
          data.chipPositions.push({
            chiptop: data.chiptop,
            chipleft: data.chipleft,
            chipText: data.tagText
          });
          data.tagText = ''
        }
      },
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
        data.imageFiles = data.imageFiles.concat(files);
        data.option.img = data.imageFiles[0].__img.src;
        // let param = new FormData();
        // param.append('file', files[0]);
        // let res = (await uploadImage(param)) as any;
        // if (res.status == 200) {
        // }
      },
      removeImage(files: Array<any>): void {
        data.imageFiles.splice(data.imageFiles.indexOf(files[0]), 1);
        console.log(data.imageFiles);
        data.chipPositions = [];
        data.alertNext = false;
        data.bigStyle =
          'min-width: 450px;min-height: 450px;max-width: 850px;max-height: 80%;height: 45vw !important;width: 30vw !important;';
      },
      closeUploader(): void {
        data.imageFiles = [];
        data.alertNext = false;
        data.option.img = '';
        data.bigStyle =
          'min-width: 450px;min-height: 450px;max-width: 850px;max-height: 80%;height: 45vw !important;width: 30vw !important;';
      },
      toRepositories(id: number): void {
        router.push(`/repository/${id}`);
      },
      openGithub(): void {
        window.open('https://github.com/hints0816/artblog', '_blank');
      },
      getEmo(index: number): void {
        const face = data.faceList[index] as string;
        data.text = data.text + face;
      },
      back(): void {
        data.bigStyle = 'min-width: 450px;min-height: 450px;max-width: 850px;max-height: 80%;height: 45vw !important;width: 30vw !important;';
        data.cropperkey = 2;
        data.alertNext = false;
      },
      next1(): void {
        let avatarfile = null;

        cropper.value.getCropData(async (data1: any) => {
          var arr = data1.split(','),
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
          let res = (await uploadImage(param)) as any;
          console.log(res.url)
          if (res.status == 200) {
            data.imgdisurl = res.url
          } else {
          }
        });
        data.alertNext = true;
        data.bigStyle =
          'min-width: 450px;min-height: 450px;max-width: 950px;max-height: 80%;height: 45vw !important;width: 70vw !important;';
      },
      setEmojiStatus(index: number): void {
        const face = data.faceList[index] as string;
        data.profile.emoji = face;
      },
      async editStatus(): Promise<any> {
        const params = {
          emoji: data.profile.emoji,
          emoji_text: data.profile.emoji_text,
        };
        let res = (await editEmoji(params)) as any;
        if (res.status == 200) {
          router.go(0);
          return Notify.create({
            message: 'Edit Successful',
            type: 'positive',
            position: 'top',
            timeout: 2000,
          });
        }
      },
      async postArt(): Promise<any> {
        // let param = new FormData()
        // param.append('file', data.imageFiles[0])
        // param.append('content', data.arttext)
        const param = {
          imgurl: data.imgdisurl,
          content: data.arttext,
          Imglist: [
            {
              img_url: data.imgdisurl
            }
          ]
        };
        let res = await postArt(param) as any
        if(res.status == 200) {
          router.go(0);
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
          data.links1.splice(2, 1);
        }
      },
      async loginOut(): Promise<any> {
        LocalStorage.remove('token');
        LocalStorage.set('logged_in', 'no');
        let res = (await logout()) as any;
        if (res.status == 302) {
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
    watch(darkdata, (newVal, oldVal) => {
      Dark.toggle();
    });
    const report = ref(null);

    return {
      cropper,
      report,
      onResize(size) {
        report.value = size;
        // {
        //   width: 20 // width of container (in px)
        //   height: 50 // height of container (in px)
        // }
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
.q-dialog__inner > div
  border-radius: 14px
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
    display: none
.toolbar-upload200
  .q-uploader__subtitle
    display: none
.toolbar-upload
  .q-uploader__file
    height: 100%
.maindia .no-pointer-events
  pointer-events: initial !important
</style>