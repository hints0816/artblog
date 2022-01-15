<template>
  <div>
    <q-input v-if="logged" autogrow v-model="text" label="Comment" :dense="dense">
      <template v-slot:before>
        <q-avatar rounded>
          <img :src="meinfo.avatar" />
        </q-avatar>
      </template>
      <template v-slot:after>
        <q-btn flat round color="primary" icon="mood">
          <q-popup-edit
            max-width="250px"
            style="padding: 4px 8px"
            self="top start"
            cover="false"
          >
            <div class="q-gutter-sm" style="margin-top: 0px">
              <a
                href="javascript:void(0);"
                @click="getEmo(index)"
                style="font-size: 22px;texttext-decoration: none"
                v-for="(item, index) in faceList"
                :key="index"
                class="emotionItem"
                >{{ item }}</a
              >
            </div>
          </q-popup-edit>
        </q-btn>
        <q-btn flat color="primary" class="text-weight-bolder" @click="addComment(0)" label="comment">
          <!-- <q-popup-proxy :breakpoint="1024">
                <a href="javascript:void(0);" v-for="(item,index) in faceList" :key="index" class="emotionItem">{{item}}</a>
            </q-popup-proxy> -->
        </q-btn>
      </template>
    </q-input>
    <q-card v-else class="bg-primary text-white">
      <q-card-section>
        <div class="text-bold">
          <q-btn
            push
            color="white"
            to="/signup"
            text-color="black"
            label="Sign up for free"
          />
          to join this conversation on GitHub. Already have an account?
          <router-link
            class="text-yellow-8 text-weight-bold"
            to="/login"
            style="color: #fff"
            >Sign in to comment</router-link
          >
        </div>
      </q-card-section>
    </q-card>
    <q-list
      v-for="(comment, idx) in comments"
      bordered
      class="rounded-borders q-mt-lg"
      :key="idx"
    >
      <q-item class="q-pa-md">
        <q-item-section avatar>
          <q-avatar class="cursor-pointer" @click="toRepositories(comment.FromProfile.id)" rounded>
            <img alt="avatar" :src="comment.FromProfile.avatar" />
          </q-avatar>
        </q-item-section>
        <q-item-section>
          <q-item-label lines="1">
            <span class="text-weight-bold">{{ comment.FromProfile.name }}&#12288;</span>
            <span class="text-gray-light">{{ comment.CreatedAt }}</span>
          </q-item-label>
          <q-item-label style="word-break:break-all;" v-html="comment.content" class="q-pt-sm"></q-item-label>
        </q-item-section>
        <div class="q-gutter-sm">
          <q-btn
            v-if="logged"
            size="10px"
            round
            flat
            color="primary"
            icon="comment"
            @click="commentss(idx, null)"
          />
          <q-btn
            size="10px"
            round
            flat
            :color="comment.DiggArr.length == 0 ? '#ccc' : 'red'"
            icon="favorite"
            @click="favorite(idx, null, comment.ID)"
          >
            <q-badge
              :color="comment.DiggArr.length == 0 ? '#ccc' : 'orange'"
              rounded
              floating
              >{{ comment.digg }}</q-badge
            >
          </q-btn>
        </div>
      </q-item>
      <div v-if="comment.isshow">
        <q-item>
          <q-item-section avatar>
            <q-avatar rounded>
              <img :src="meinfo.avatar" />
            </q-avatar>
          </q-item-section>
          <q-item-section>
            <q-input
              autogrow
              v-model="childtext"
              :label="'回复：' + comment.FromProfile.name"
              :dense="dense"
            >
              <template v-slot:after>
                <q-btn flat round color="primary" icon="mood">
                  <q-popup-edit
                    max-width="250px"
                    style="padding: 4px 8px"
                    self="top start"
                    cover="false"
                  >
                    <div class="q-gutter-sm" style="margin-top: 0px">
                      <a
                        href="javascript:void(0);"
                        @click="getEmo1(index)"
                        style="font-size: 22px;text-decoration: none"
                        v-for="(item, index) in faceList"
                        :key="index"
                        class="emotionItem"
                        >{{ item }}</a
                      >
                    </div>
                  </q-popup-edit>
                </q-btn>
                <q-btn
                  flat
                  class="text-weight-bolder"
                  color="primary"
                  @click="addComment(comment.ID)"
                  label="Comment"
                >
                  <!-- <q-popup-proxy :breakpoint="1024">
                        <a href="javascript:void(0);" v-for="(item,index) in faceList" :key="index" class="emotionItem">{{item}}</a>
                    </q-popup-proxy> -->
                </q-btn>
              </template>
            </q-input>
          </q-item-section>
        </q-item>
      </div>
      <q-list
        v-for="(rev, index) in comment.CommentChild"
        class="q-mx-lg"
        :key="index"
      >
        <q-item class="q-pa-md">
          <q-item-section avatar>
            <q-avatar class="cursor-pointer" @click="toRepositories(rev.FromProfile.id)" rounded>
              <img alt="avatar" :src="rev.FromProfile.avatar" />
            </q-avatar>
          </q-item-section>
          <q-item-section>
            <q-item-label lines="1">
              <span class="text-weight-bold">{{ rev.FromProfile.name }}</span>
              <span class="text-weight-bold"> 回复：</span>
              <span class="text-gray-light"> {{ rev.ToProfile.name }}</span>
            </q-item-label>
            <q-item-label style="word-break:break-all;" v-html="rev.content"></q-item-label>
          </q-item-section>
          <div class="q-gutter-sm">
            <q-btn
              v-if="logged"
              size="10px"
              round
              flat
              color="primary"
              icon="comment"
              @click="commentss(idx, index)"
            />
            <q-btn
              size="10px"
              round
              flat
              :color="rev.DiggArr.length == 0 ? '#ccc' : 'red'"
              icon="favorite"
              @click="favorite(idx, index, rev.ID)"
            >
              <q-badge
                :color="rev.DiggArr.length == 0 ? '#ccc' : 'orange'"
                rounded
                floating
                >{{ rev.digg }}</q-badge
              >
            </q-btn>
          </div>
        </q-item>
        <div v-if="rev.isshow">
          <q-item>
            <q-item-section avatar>
              <q-avatar rounded>
                <img :src="meinfo.avatar" />
              </q-avatar>
            </q-item-section>
            <q-item-section>
              <q-input
                autogrow
                v-model="childtext"
                :label="'回复：' + rev.ToProfile.name"
                :dense="dense"
              >
                <template v-slot:after>
                  <q-btn flat round color="primary" icon="mood">
                    <q-popup-edit
                      max-width="250px"
                      style="padding: 4px 8px"
                      self="top start"
                      cover="false"
                    >
                      <div class="q-gutter-sm" style="margin-top: 0px">
                        <a
                          href="javascript:void(0);"
                          @click="getEmo1(index)"
                          style="font-size: 22px;text-decoration: none"
                          v-for="(item, index) in faceList"
                          :key="index"
                          class="emotionItem"
                          >{{ item }}</a
                        >
                      </div>
                    </q-popup-edit>
                  </q-btn>
                  <q-btn
                    flat
                    color="primary"
                    class="text-weight-bolder"
                    @click="addComment(rev.ID)"
                    label="Comment"
                  >
                    <!-- <q-popup-proxy :breakpoint="1024">
                        <a href="javascript:void(0);" v-for="(item,index) in faceList" :key="index" class="emotionItem">{{item}}</a>
                    </q-popup-proxy> -->
                  </q-btn>
                </template>
              </q-input>
            </q-item-section>
          </q-item>
        </div>
      </q-list>
    </q-list>
    <div class="q-pa-md q-gutter-sm">
      <!-- <q-btn outline no-caps color="cyan-9" @click="goAddComment">
        <q-icon left size="2rem" name="add_comment" />
        <div>Add Comment</div>
      </q-btn> -->
    </div>

    <!-- <v-md-preview-html :html="marktext" preview-class="vuepress-markdown-body"></v-md-preview-html> -->
  </div>
</template>

<script lang="ts">
import emoji from '../css/emoji.json';
import { reactive, toRefs, onMounted } from 'vue';
import { LocalStorage, Notify, date } from 'quasar';
import { listComment, addComment, digg, undigg } from '../api/test/index';
import { useRoute, useRouter } from 'vue-router';
export default {
  name: 'Comment',
  components: {},
  props: {
    idss: {
      type: String,
      required: true,
    },
  },
  setup() {
    let data = reactive({
      meinfo: {},
      logged: false,
      marktext: '',
      faceList: [],
      text: '',
      childtext: '',
      isshow: false,
      isfav: false,
      isshowindex: 0,
      isshowindex1: 0,
      comments: [],
    }) as any;
    const route = useRoute() as any;
    const router = useRouter() as any;
    const method = {
      commentss(fatherindex, childindex): void {
        data.childtext = '';
        if (childindex == null) {
          if (data.comments[data.isshowindex].CommentChild.length != 0) {
            data.comments[data.isshowindex].CommentChild[
              data.isshowindex1
            ].isshow = false;
          }
          data.comments[data.isshowindex].isshow = false;
          data.comments[fatherindex].isshow = true;
          data.isshowindex = fatherindex;
        } else {
          data.comments[data.isshowindex].CommentChild[
            data.isshowindex1
          ].isshow = false;
          data.comments[data.isshowindex].isshow = false;
          data.comments[fatherindex].CommentChild[childindex].isshow = true;
          data.isshowindex = fatherindex;
          data.isshowindex1 = childindex;
        }
      },
      toRepositories(id: number): void {
        router.push(`/repository/${id}`)
      },
      async favorite(
        fatherindex: number,
        childindex: number,
        comment_id: number
      ): Promise<any> {
        const params = {
          comment_id: comment_id,
        };

        if (childindex == null) {
          if (data.comments[fatherindex].DiggArr.length == 0) {
            let res = (await digg(params)) as any;
            if (res.status == 200) {
              data.comments[fatherindex].DiggArr[0] = true;
              data.comments[fatherindex].digg++;
            }
          } else {
            let res = (await undigg(params)) as any;
            if (res.status == 200) {
              data.comments[fatherindex].DiggArr = [];
              data.comments[fatherindex].digg--;
            }
          }
        } else {
          if (
            data.comments[fatherindex].CommentChild[childindex].DiggArr
              .length == 0
          ) {
            let res = (await digg(params)) as any;
            if (res.status == 200) {
              data.comments[fatherindex].CommentChild[
                childindex
              ].DiggArr[0] = true;
              data.comments[fatherindex].CommentChild[childindex].digg++;
            }
          } else {
            let res = (await undigg(params)) as any;
            if (res.status == 200) {
              data.comments[fatherindex].CommentChild[childindex].DiggArr = [];
              data.comments[fatherindex].CommentChild[childindex].digg--;
            }
          }
        }
      },
      async getTalk(): Promise<any> {
        if (LocalStorage.getItem('logged_in') === 'yes') {
          data.logged = true;
        }
        const paramss = {
          pagenum: 1,
          pagesize: 10,
        };
        let res = (await listComment(route.params.id, paramss)) as any;
        if(res.data != null)
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
      async addComment(id: number): Promise<any> {
        let paramss = {
          article_id: (route.params.id * 1000) / 1000,
          content: data.text,
          comment_id: 0,
        };
        if (id != 0) {
          paramss.content = data.childtext;
          paramss.comment_id = id;
        } else {
          delete paramss.comment_id;
        }
        if (paramss.content == '') {
          Notify.create({
            message: 'comment can\'t be empty',
            color: 'negative',
            icon: 'report_problem',
            position: 'top',
            timeout: 2000,
          });
          return;
        }
        let res = (await addComment(paramss)) as any;
        if (res.status == 200 ) {
          data.text = ''
          data.childtext = ''
        }
        await method.getTalk();
      },
      getEmo(index): void {
        const face = data.faceList[index] as string;
        data.text = (data.text as string) + face;
      },
      getEmo1(index): void {
        const face = data.faceList[index] as string;
        data.childtext = (data.childtext as string) + face;
      },

      async getEmo2(): Promise<void> {
        /* const params = {
          addressId: data.marktext
        } */
      },
    };
    onMounted(async () => {
      data.meinfo.avatar = LocalStorage.getItem('avatar');
      await method.getTalk();
      emoji.forEach((element) => {
        data.faceList.push(element.char);
      });
    });
    return {
      ...toRefs(data),
      ...method,
    };
  },
};
</script>
