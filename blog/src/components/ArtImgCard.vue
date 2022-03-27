<template>
  <q-infinite-scroll @load="onLoad" :offset="200">
        <div :class="outClass + ' row'">
          <div
            v-for="(src, index) in images"
            :key="index"
            :class="loadClass(index)"
          >
            <template v-if="checkClass(index) == -1">
              <q-img
                @click="getDetail(src.id)"
                @mouseenter="alertNum = index"
                @mouseleave="alertNum = -1"
                :src="src.url"
              >
                <div
                  align="center"
                  v-if="alertNum == index"
                  class="relative-position cursor-pointer indexz"
                >
                  <div
                    class="absolute-center q-col-gutter-x-md"
                    style="font-size: 2em"
                  ></div>
                </div>
              </q-img>
            </template>
          </div>
        </div>
        <template v-slot:loading>
          <div class="row justify-center q-my-md">
            <q-spinner-dots color="primary" size="40px" />
          </div>
        </template>
      <q-dialog
        class="maindia"
        style="background-color: rgba(0, 0, 0, 0.6)"
        persistent
        v-model="alert"
      >
        <q-btn
          class="text-white absolute-top-right q-mr-lg q-mt-lg"
          icon="close"
          v-close-popup
          flat
          round
          dense
        />
        <q-btn
          class="text-white absolute-right"
          @click="nextimg"
          style="top: initial !important; bottom: initial !important"
          icon="navigate_next"
          flat
          round
          dense
        />
        <q-btn
          class="text-white absolute-left"
          @click="backimg"
          style="top: initial !important; bottom: initial !important"
          icon="chevron_left"
          flat
          round
          dense
        />
        <q-card
          class="full-width"
          style="width: 700px; max-width: 70vw; height: 90% !important"
        >
          <q-card-section v-if="$q.screen.gt.sm" horizontal class="full-height">
            <q-carousel
              animated
              v-model="imgDetail.slide"
              transition-prev="slide-right"
              transition-next="slide-left"
              :arrows="imgDetail.Imglist.length > 1"
              :navigation="imgDetail.Imglist.length > 1"
              infinite
              class="col-7 full-height"
            >
              <q-carousel-slide
                v-for="(data, index) in imgDetail.Imglist"
                :key="index"
                :name="index + 1"
                class="column justify-center bg-black cursor-pointer	"
                style="padding: 0px;"
                @click="showTag()"
              >
              <q-img
                :src="data.img_url"
                :ratio="1"
                spinner-color="white"
              >
              <q-chip
                  v-for="(item, index2) in data.Imgtag"
                  :key="index2"
                  class="chiptag bg-dark text-white"
                  :style="{
                    'z-index': 99,
                    position: 'absolute',
                    top: item.top*100 + '%',
                    left: item.left*100 + '%',
                  }"
                  dense
                  >{{item.tag}}</q-chip
                >
              </q-img> 
              </q-carousel-slide>
            </q-carousel>
            <q-card-section style="padding: 0px" class="full-width">
              <q-item>
                <q-item-section avatar>
                  <q-btn
                    size="0px"
                    @click="toRepositories(imgDetail.Profile.id)"
                    round
                  >
                    <q-avatar>
                      <img :src="imgDetail.Profile.avatar" />
                    </q-avatar>
                  </q-btn>
                </q-item-section>
                <q-item-section>
                  <q-item-label>{{ imgDetail.Profile.name }}</q-item-label>
                </q-item-section>
                <q-item-section side>
                  <q-btn @click="deleteImg(imgDetail.ID)" class="text-red" flat dense round icon="delete_outline"/>
                </q-item-section>
              </q-item>
              <q-separator />
              <q-item
                style="position: absolute; top: 56px; width: 100%; bottom: 56px"
              >
                <q-item-section
                  style="max-height: 55vh; position: relative"
                  class="scroll hide-scrollbar"
                >
                  <q-card-section
                    class="full-width"
                    style="position: absolute; top: 0"
                  >
                    <p>{{ imgDetail.content }}</p>
                    <q-item
                      v-for="(data, index) in comments"
                      :key="index"
                      style="padding: 0px"
                    >
                      <q-item-section avatar>
                        <q-btn
                          size="0px"
                          @click="toRepositories(data.FromProfile.id)"
                          round
                        >
                          <q-avatar size="30px">
                            <img :src="data.FromProfile.avatar" />
                          </q-avatar>
                        </q-btn>
                      </q-item-section>
                      <q-item-section class="full-width">
                        <q-item-label class="text-weight-bold">{{
                          data.FromProfile.name
                        }}</q-item-label>
                        <q-item-label style="word-break: break-all">{{
                          data.content
                        }}</q-item-label>
                      </q-item-section>
                      <q-item-section side top>
                        <q-item-label class="text-caption">{{
                          data.CreatedAt
                        }}</q-item-label>
                        <!-- <q-btn @click="morealert = true" class="gt-xs" size="12px" flat dense round icon="fas fa-ellipsis-h"/> -->
                      </q-item-section>
                    </q-item>
                  </q-card-section>
                </q-item-section>
              </q-item>
              <q-card-actions
                class="full-width"
                style="position: absolute; bottom: 0px; padding: 0px"
              >
                <q-separator />
                <q-input
                  style="height: 56px"
                  class="scroll hide-scrollbar full-width q-px-sm"
                  autogrow
                  v-model="text"
                  label="Comment"
                  :dense="dense"
                >
                  <template v-slot:before>
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
                  </template>
                  <template v-slot:after>
                    <q-btn
                      push
                      color="primary"
                      @click="addImgComment(0)"
                      label="Comment"
                    >
                    </q-btn>
                  </template>
                </q-input>
              </q-card-actions>
            </q-card-section>
          </q-card-section>
          <q-card-section style="padding: 0px" v-else class="full-width">
            <q-card-section style="padding: 0px" class="full-width">
              <q-item>
                <q-item-section avatar>
                  <q-avatar>
                    <img :src="imgDetail.Profile.avatar" />
                  </q-avatar>
                </q-item-section>
                <q-item-section>
                  <q-item-label>{{ imgDetail.Profile.name }}</q-item-label>
                </q-item-section>
                <q-item-section side top>
                  <q-btn
                    @click="morealert = true"
                    class="gt-xs"
                    size="12px"
                    flat
                    dense
                    round
                    icon="fas fa-ellipsis-h"
                  />
                </q-item-section>
              </q-item>
              <q-separator />
              <q-carousel
                animated
                v-model="imgDetail.slide"
                transition-prev="slide-right"
                transition-next="slide-left"
                :arrows="imgDetail.Imglist.length > 1"
                :navigation="imgDetail.Imglist.length > 1"
                infinite
                class="col-5 full-height"
              >
                <q-carousel-slide
                  v-for="(data, index) in imgDetail.Imglist"
                  :key="index"
                  :name="index + 1"
                  :img-src="data.img_url"
                >
                </q-carousel-slide>
              </q-carousel>
              <q-item>
                <q-item-section>
                  <q-card-section
                    style="max-height: 50vh"
                    class="scroll hide-scrollbar"
                  >
                    <p v-for="n in 15" :key="n">
                      Lorem ipsum dolor sit amet consectetur adipisicing elit.
                      Rerum repellendus sit voluptate voluptas eveniet porro.
                      Rerum blanditiis perferendis totam, ea at omnis vel
                      numquam exercitationem aut, natus minima, porro labore.
                    </p>
                  </q-card-section>
                </q-item-section>
              </q-item>
              <q-separator />
              <q-card-actions style="padding: 0px">
                <q-input
                  class="full-width q-px-sm"
                  autogrow
                  v-model="text"
                  label="Comment"
                  :dense="dense"
                >
                  <template v-slot:before>
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
                  </template>
                  <template v-slot:after>
                    <q-btn
                      push
                      color="primary"
                      @click="addImgComment(0)"
                      label="Comment"
                    >
                    </q-btn>
                  </template>
                </q-input>
              </q-card-actions>
            </q-card-section>
          </q-card-section>
        </q-card>
      </q-dialog>
      <q-dialog
        v-model="morealert"
        class="moredia"
        transition-show="scale"
        transition-hide="scale"
      >
        <q-card style="width: 300px">
          <q-card-section style="padding: 0px">
            <q-list bordered separator>
              <q-item clickable v-ripple>
                <q-item-section>Single line item</q-item-section>
              </q-item>
              <q-item clickable v-ripple>
                <q-item-section>
                  <q-item-label>Item with caption</q-item-label>
                </q-item-section>
              </q-item>
              <q-item clickable v-ripple>
                <q-item-section>
                  <q-item-label overline>OVERLINE</q-item-label>
                </q-item-section>
              </q-item>
            </q-list>
          </q-card-section>
        </q-card>
      </q-dialog>
      </q-infinite-scroll>
</template>
<script lang="ts">
import {
  listArtList,
  getImgContent,
  listImgComment,
  delImg
} from '../api/test/index';
import {
  ref,
  reactive,
  toRefs,
  onBeforeMount,
} from 'vue';
import { Screen, date, Notify } from 'quasar';
import { getCurrentInstance } from 'vue';
import { useRoute, useRouter } from 'vue-router';

export default {
  name: 'Post',
  setup() {
    let data = reactive({
      alertNum: -1,
      images: [],
      alert: false,
      imgSort: [],
      comments: [],
      imgDetail: null,
      pagesize: 18,
      pagenum: 1,
      total: 18,
      outClass: Screen.lt.lg
        ? 'q-col-gutter-' + Screen.name
        : 'q-col-gutter-lg',
      inClass: Screen.lt.lg
        ? 'q-col-gutter-y-' + Screen.name
        : 'q-col-gutter-y-lg',
    });
    const route = useRoute() as any;
    const router = useRouter() as any;
    const { ctx } = getCurrentInstance() as any;
    const method = {
      async getDetail(id: number): Promise<any> {
        let res = (await getImgContent(id)) as any;
        data.imgDetail = res.data;
        data.imgDetail.slide = 1;
        await method.getImgComment();
        data.alert = true;
      },
      deleteImg(id: number): void {
        Notify.create({
          message: 'Are you absolutely sure?',
          color: 'negative',
          icon: 'report_problem',
          position: 'top',
          actions: [
            { label: 'yes', color: 'white', handler: async () => { 
              let res = (await delImg(id)) as any;
              if(res.status == 200) {
                return router.go(0)
              }
            }}
          ]
        })
      },
      async getImgComment(): Promise<any> {
        const params = {
          pagenum: 1,
          pagesize: 10,
        };
        let res = (await listImgComment(data.imgDetail.ID, params)) as any;
        if (res.data != null)
          data.comments = res.data.map((item) => {
            item.isshow = false;
            let timeStamp1 = new Date(item.CreatedAt);
            let formattedString1 = date.formatDate(
              timeStamp1,
              'YYYY-MM-DD HH:mm:ss'
            );
            item.CreatedAt = formattedString1;
            const res1 = item.CommentChild;
            if (res1 != null) {
              if (res1.length != 0) {
                res1.forEach((element) => {
                  element.isshow = false;
                  let timeStamp2 = new Date(element.CreatedAt);
                  let formattedString2 = date.formatDate(
                    timeStamp2,
                    'YYYY-MM-DD HH:mm:ss'
                  );
                  element.CreatedAt = formattedString2;
                });
              }
            }
            return item;
          });
      },
      async onLoad(index, done): Promise<any> {
        if (data.total > data.pagesize*data.pagenum) {
          // data.pagesize += 10;
          data.pagenum += 1;
          const params = {
            pagenum: data.pagenum,
            pagesize: data.pagesize,
          };
          let datas = (await listArtList(params)) as any;
          setTimeout(() => {
            datas.data.forEach((element: any) => {
              data.images.push({ url: element.imgurl, id: element.ID });
              data.imgSort.push(element.ID);
            });
            done();
          }, 2000);
        } else {
          done();
        }
      },
      checkClass(nums: number): number {
        return -1;
      },
      checkNum(nums: number): boolean {
        const r = /^\+?[0-9][0-9]*$/;
        if (r.test(String(nums))) {
          return true;
        } else {
          return false;
        }
      },
      loadClass(nums: number): string {
        return 'col-4';
      },
      async listArt(): Promise<any> {
        const paramss = {
          pagenum: 1,
          pagesize: data.pagesize,
        };
        let res = (await listArtList(paramss)) as any;
        res.data.forEach((element: any) => {
          data.images.push({ url: element.imgurl, id: element.ID });
          data.imgSort.push(element.ID);
        });
        data.total = res.total
      },
    };
    onBeforeMount(async () => {
      await method.listArt();
      ctx.$nextTick(() => {
        window.addEventListener('resize', () => {
          if (Screen.lt.lg) {
            (data.outClass = 'q-col-gutter-' + Screen.name),
              (data.inClass = 'q-col-gutter-y-' + Screen.name);
          }
        });
      });
    });
    return {
      ...toRefs(data),
      ...method,
    };
  },
};
</script>
<style lang="sass">
.nopadding
  padding: 0px
.indexz
  position: fixed
  top: 0
  left: 0
  bottom: 0
  right: 0
  width: 100%
  height: 100%
  background: rgba(0,0,0,0.8)
</style>