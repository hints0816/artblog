<template>
  <div>
    <q-input v-if="logged" autogrow v-model="text" label="Label" :dense="dense">
      <template v-slot:before>
        <q-avatar rounded>
          <img :src="meinfo.avatar" />
        </q-avatar>
      </template>
      <template v-slot:after>
        <q-btn push color="primary" icon="mood">
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
        <q-btn push color="primary" @click="addComment()" label="Comment">
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
          <q-avatar rounded>
            <img alt="avatar" :src="comment.Profile.avatar" />
          </q-avatar>
        </q-item-section>
        <q-item-section>
          <q-item-label lines="1">
            <span class="text-weight-bold">{{ comment.Profile.name }}</span>
            <span class="text-gray-light"> {{ comment.CreatedAt }}</span>
          </q-item-label>
          <q-item-label v-html="comment.content" class="q-pt-sm"></q-item-label>
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
            v-if="comment.digg == 0 ? true : false"
            size="10px"
            round
            flat
            color="#ccc"
            icon="favorite"
            @click="favorite(idx)"
          />
          <q-btn
            v-else
            size="10px"
            round
            flat
            color="red"
            icon="favorite"
            @click="favorite(idx)"
          />
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
              v-model="text1"
              :label="'回复：' + comment.Profile.name"
              :dense="dense"
            >
              <template v-slot:after>
                <q-btn push color="primary" icon="mood">
                  <q-popup-edit
                    max-width="158px"
                    style="padding: 4px 8px"
                    self="top start"
                    cover="false"
                  >
                    <div class="q-gutter-sm" style="margin-top: 0px">
                      <a
                        href="javascript:void(0);"
                        @click="getEmo1(index)"
                        style="text-decoration: none"
                        v-for="(item, index) in faceList"
                        :key="index"
                        class="emotionItem"
                        >{{ item }}</a
                      >
                    </div>
                  </q-popup-edit>
                </q-btn>
                <q-btn push color="primary" @click="addCommentChild()" label="Comment">
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
        v-for="(rev, index) in comment.Commentchild"
        class="q-mx-lg"
        :key="index"
      >
        <q-item class="q-pa-md">
          <q-item-section avatar>
            <q-avatar rounded>
              <img alt="avatar" :src="rev.FromProfile.avatar" />
            </q-avatar>
          </q-item-section>
          <q-item-section>
            <q-item-label lines="1">
              <span class="text-weight-bold">{{ rev.FromProfile.name }}</span>
              <span class="text-weight-bold"> 回复：</span>
              <span class="text-gray-light"> {{ rev.ToProfile == null?'':rev.ToProfile.name }}</span>
            </q-item-label>
            <q-item-label v-html="rev.content"></q-item-label>
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
              v-if="rev.digg == 0 ? true : false"
              size="10px"
              round
              flat
              color="#ccc"
              icon="favorite"
              @click="favorite(idx)"
            />
            <q-btn
              v-else
              size="10px"
              round
              flat
              color="red"
              icon="favorite"
              @click="favorite(idx)"
            />
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
                v-model="text1"
                :label="'回复：' + rev.FromProfile.name"
                :dense="dense"
              >
                <template v-slot:after>
                  <q-btn push color="primary" icon="mood">
                    <q-popup-edit
                      max-width="158px"
                      style="padding: 4px 8px"
                      self="top start"
                      cover="false"
                    >
                      <div class="q-gutter-sm" style="margin-top: 0px">
                        <a
                          href="javascript:void(0);"
                          @click="getEmo1(index)"
                          style="text-decoration: none"
                          v-for="(item, index) in faceList"
                          :key="index"
                          class="emotionItem"
                          >{{ item }}</a
                        >
                      </div>
                    </q-popup-edit>
                  </q-btn>
                  <q-btn push color="primary" @click="addCommentChild(rev.ID)"  label="Comment">
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
import { getCurrentInstance, reactive, toRefs, onMounted } from 'vue';
import { LocalStorage } from 'quasar';
import { listComment, addComment, addCommentChild } from '../api/test/index';
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
      text1: '',
      isshow: false,
      isfav: false,
      isshowindex: 0,
      isshowindex1: 0,
      comments: [],
    }) as any;
    const { ctx } = getCurrentInstance() as any;
    const route = useRoute() as any;
    const router = useRouter() as any;
    console.log(ctx);
    const method = {
      commentss(fatherindex, childindex): void {
        console.log(fatherindex);
        console.log(childindex);
        data.text1 = '';
        if (childindex == null) {
          data.comments[data.isshowindex].Commentchild[
            data.isshowindex1
          ].isshow = false;
          data.comments[data.isshowindex].isshow = false;
          data.comments[fatherindex].isshow = true;
          data.isshowindex = fatherindex;
        } else {
          data.comments[data.isshowindex].Commentchild[
            data.isshowindex1
          ].isshow = false;
          data.comments[data.isshowindex].isshow = false;
          data.comments[fatherindex].Commentchild[childindex].isshow = true;
          data.isshowindex = fatherindex;
          data.isshowindex1 = childindex;
        }
      },
      favorite(comment): void {
        data.comments[comment].isfav = true;
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
        data.comments = res.data.map((item) => {
          item.isshow = false;
          const res1 = item.Commentchild;
          res1.forEach((element) => {
            element.isshow = false;
          });
          return item;
        });
        console.log(data.comments);
      },
      async addComment(): Promise<any> {
        const paramss = {
          article_id: (route.params.id * 1000) / 1000,
          content: data.text,
        };
        let res = (await addComment(paramss)) as any;
        console.log(res);
        await method.getTalk();
      },
      async addCommentChild(id: number): Promise<any> {
        const paramss = {
          article_id: (route.params.id * 1000) / 1000,
          content: data.text1,
          parent_id: id,
        };
        let res = (await addCommentChild(paramss)) as any;
        console.log(res);
        await method.getTalk();
      },
      getEmo(index): void {
        const face = data.faceList[index] as string;
        data.text = (data.text as string) + face;
      },
      getEmo1(index): void {
        const face = data.faceList[index] as string;
        data.text1 = (data.text1 as string) + face;
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
