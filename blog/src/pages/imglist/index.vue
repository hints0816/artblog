<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-12 col-md-8 col-lg-6">
      <q-infinite-scroll @load="onLoad" :offset="250">
        <div :class="outClass + ' row'">
          <div
            v-for="(src, index) in images"
            :key="index"
            :class="loadClass(index)"
          >
            <template v-if="checkClass(index) == 0 || checkClass(index) == -1">
              <q-img @click="getDetail(src.id)" @mouseenter="alertNum = index" @mouseleave="alertNum = -1" :src="src.url">
                <div align="center" v-if="alertNum == index" class="relative-position cursor-pointer indexz">
                  <div class="absolute-center q-col-gutter-x-md" style="font-size: 2em">
                  </div>
                </div>
              </q-img>
            </template>
            <template v-if="checkClass(index) == 1">
              <div class="col-6">
                <q-img @click="getDetail(images[index].id)" @mouseenter="alertNum = index" @mouseleave="alertNum = -1" :src="images[index].url">
                  <div align="center" v-if="alertNum == index" class="relative-position cursor-pointer indexz">
                    <div class="absolute-center q-col-gutter-x-md" style="font-size: 2em">
                    </div>
                  </div>
                </q-img>
              </div>
              <div class="col-6">
                <q-img @click="getDetail(images[index+1].id)" @mouseenter="alertNum = index+1" @mouseleave="alertNum = -1" :src="images[index+1].url">
                  <div align="center" v-if="alertNum == index+1" class="relative-position cursor-pointer indexz">
                    <div class="absolute-center q-col-gutter-x-md" style="font-size: 2em">
                    </div>
                  </div>
                </q-img>
              </div>
            </template>
          </div>
        </div>
      </q-infinite-scroll>
      <q-dialog class="maindia" style="background-color: rgba(0,0,0,0.6);" persistent v-model="alert">
      <q-btn class="text-white absolute-top-right q-mr-lg q-mt-lg" icon="close" v-close-popup flat round dense/>
      <q-btn class="text-white absolute-right" style="top: initial !important;bottom: initial !important;" icon="navigate_next" flat round dense/>
      <q-btn class="text-white absolute-left" style="top: initial !important;bottom: initial !important;" icon="chevron_left" flat round dense/>
      <q-card class="full-width" style="width: 700px; max-width: 80vw;height: 80% !important">
        <q-card-section v-if="$q.screen.gt.sm" horizontal class="full-height">
          <q-carousel
            animated
            v-model="imgDetail.slide"
            transition-prev="slide-right"
            transition-next="slide-left"
            :arrows="imgDetail.Imglist.length>1"
            :navigation="imgDetail.Imglist.length>1"
            infinite
            class="col-5 full-height"
          >
            <q-carousel-slide  v-for="(data, index) in imgDetail.Imglist"
            :key="index" :name="index+1" :img-src="data.img_url" />
          </q-carousel>
          <q-card-section style="padding: 0px;" class="full-width">
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
                <q-btn @click="morealert = true" class="gt-xs" size="12px" flat dense round icon="fas fa-ellipsis-h"/>
              </q-item-section>
            </q-item>
             <q-separator />
            <q-item>
              <q-item-section>
                <q-card-section style="max-height: 55vh" class="scroll hide-scrollbar">
                  <p>{{imgDetail.content}}</p>
                </q-card-section>
              </q-item-section>
            </q-item>
             <q-separator />
            <q-card-actions class="full-width" style="position:absolute;bottom:0px;padding: 0px;">
              <q-input class="full-width q-px-sm" autogrow v-model="text" label="Comment" :dense="dense">
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
                    <q-btn push color="primary" @click="addComment(0)" label="Comment">
                    </q-btn>
                  </template>
                </q-input>
            </q-card-actions>
          </q-card-section>
        </q-card-section>
        <q-card-section style="padding: 0px;" v-else class="full-width">
          <q-card-section style="padding: 0px;" class="full-width">
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
                <q-btn @click="morealert = true" class="gt-xs" size="12px" flat dense round icon="fas fa-ellipsis-h"/>
              </q-item-section>
            </q-item>
             <q-separator />
             <q-img
            class="col-12 full-height"
            src="https://cdn.quasar.dev/img/parallax1.jpg"
          />
            <q-item>
              <q-item-section>
                <q-card-section style="max-height: 50vh" class="scroll hide-scrollbar">
                  <p v-for="n in 15" :key="n">Lorem ipsum dolor sit amet consectetur adipisicing elit. Rerum repellendus sit voluptate voluptas eveniet porro. Rerum blanditiis perferendis totam, ea at omnis vel numquam exercitationem aut, natus minima, porro labore.</p>
                </q-card-section>
              </q-item-section>
            </q-item>
             <q-separator />
            <q-card-actions style="padding: 0px;">
              <q-input class="full-width q-px-sm" autogrow v-model="text" label="Comment" :dense="dense">
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
                    <q-btn push color="primary" @click="addComment(0)" label="Comment">
                    </q-btn>
                  </template>
                </q-input>
            </q-card-actions>
          </q-card-section>
        </q-card-section>
      </q-card>
    </q-dialog>
    <q-dialog v-model="morealert" class="moredia" transition-show="scale" transition-hide="scale">
      <q-card style="width: 300px;">
        <q-card-section style="padding: 0px;">
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
    </q-page>
  </div>
</template>
<script lang="ts">
import { ref, reactive, toRefs, onMounted, getCurrentInstance } from 'vue';
import { Screen } from 'quasar';
import { listArtList, getImgContent } from '../../api/test/index';
import emoji from '../../css/emoji.json';

export default {
  setup() {
    const { ctx } = getCurrentInstance() as any;
    let data = reactive({
      slide: 1,
      alert: false,
      morealert: false,
      alertNum: -1,
      text: '',
      faceList: [],
      outClass: Screen.lt.lg
        ? 'q-col-gutter-' + Screen.name
        : 'q-col-gutter-lg',
      inClass: Screen.lt.lg
        ? 'q-col-gutter-y-' + Screen.name
        : 'q-col-gutter-y-lg',
      images: [],
      imgDetail: null
    });
    const method = {
      checkNum(nums: number): boolean {
        const r = /^\+?[0-9][0-9]*$/;
        if (r.test(String(nums))) {
          return true;
        } else {
          return false;
        }
      },
      getEmo(index: number): void {
        const face = data.faceList[index] as string;
        data.text = data.text + face;
      },
      loadClass(nums: number): string {
        if (method.checkNum(nums / 9)) {
          if (method.checkNum(nums / 9 / 2)) {
            return 'col-8';
          } else {
            return data.inClass + ' column col-4';
          }
        } else if (method.checkNum((nums - 1) / 9)) {
          if (method.checkNum((nums - 1) / 9 / 2)) {
            return data.inClass + ' column col-4';
          } else {
            return 'nopadding';
          }
        } else if (method.checkNum((nums - 2) / 9)) {
          if (method.checkNum((nums - 2) / 9 / 2)) {
            return 'nopadding';
          } else {
            return 'col-8';
          }
        } else {
          return 'col-4';
        }
      },
      // alertClass(nums: number): boolean {
      //   return 0;
      // },
      checkClass(nums: number): number {
        if (method.checkNum(nums / 9)) {
          if (method.checkNum(nums / 9 / 2)) {
            return 0;
          } else {
            return 1;
          }
        } else if (method.checkNum((nums - 1) / 9)) {
          if (method.checkNum((nums - 1) / 9 / 2)) {
            return 1;
          } else {
            return 2;
          }
        } else if (method.checkNum((nums - 2) / 9)) {
          if (method.checkNum((nums - 2) / 9 / 2)) {
            return 2;
          } else {
            return 0;
          }
        } else {
          return -1;
        }
      },
      async listArt():Promise<any> {
        const paramss = {
          pagenum: 1,
          pagesize: 20,
        };
        let datas = (await listArtList(paramss)) as any;
        datas.data.forEach((element) => {
          data.images.push({url:element.imgurl,id:element.ID})
        });
      },
      async getDetail(id: number):Promise<any> {
        let res  = await getImgContent(id) as any
        data.imgDetail = res.data
        data.imgDetail.slide = 1
        data.alert = true
      }
    };
    onMounted(async () => {
      await method.listArt()
      emoji.forEach((element) => {
        data.faceList.push(element.char);
      });
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
.maindia .no-pointer-events
  pointer-events: initial !important
</style>