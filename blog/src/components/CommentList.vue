<template>
  <div>
    <!-- <q-input autogrow v-model="text" label="Label" :dense="dense" />
    <q-btn push color="primary" label="Push" /> -->
    <q-input autogrow v-model="text" label="Label" :dense="dense">
        <template v-slot:before>
          <q-avatar rounded>
            <img src="https://cdn.quasar.dev/img/avatar5.jpg">
          </q-avatar>
        </template>
        <template v-slot:after>
          <q-btn push color="primary" icon="mood">
              <q-popup-edit max-width="158px" style="padding: 4px 8px;" self="top start" cover="false">
                <div class="q-gutter-sm" style="margin-top:0px">
                  <a href="javascript:void(0);" @click="getEmo(index)" style="text-decoration:none" v-for="(item,index) in faceList" :key="index" class="emotionItem">{{item}}</a>
                </div>
              </q-popup-edit>
          </q-btn>
          <q-btn push color="primary" label="Comment">
            <!-- <q-popup-proxy :breakpoint="1024">
                <a href="javascript:void(0);" v-for="(item,index) in faceList" :key="index" class="emotionItem">{{item}}</a>
            </q-popup-proxy> -->
          </q-btn>
        </template>
      </q-input>
    <q-list
      v-for="(comment,idx) in comments"
      bordered
      class="rounded-borders q-mt-lg"
      :key="idx"
    >
      <q-item class="q-pa-md">
        <q-item-section avatar>
          <q-avatar rounded>
            <img alt="avatar" :src="comment.user.avatar_url" />
          </q-avatar>
        </q-item-section>
        <q-item-section>
          <q-item-label lines="1">
            <span class="text-weight-bold">{{ comment.user.login }}</span>
            <span class="text-gray-light"> {{ comment.updated_at }}</span>
          </q-item-label>
          <q-item-label v-html="comment.body_html" class="q-pt-sm"></q-item-label>
        </q-item-section>
        <div class="q-gutter-sm">
            <q-btn size="10px" round flat color="primary" icon="comment"  @click="commentss(idx,null)" />
            <q-btn v-if="comment.digg==0?true:false" size="10px" round flat color="#ccc" icon="favorite"  @click="favorite(idx)" />
            <q-btn v-else size="10px" round flat color="red" icon="favorite"  @click="favorite(idx)" />
        </div>
      </q-item>
      <div v-if="comment.isshow">
        <q-item>
            <q-item-section avatar>
                <q-avatar rounded>
                    <img src="https://cdn.quasar.dev/img/avatar5.jpg">
                  </q-avatar>
            </q-item-section>
            <q-item-section>
                <q-input autogrow v-model="text1" :label="'回复：'+comment.user.login" :dense="dense">
                <template v-slot:after>
                  <q-btn push color="primary" icon="mood">
                    <q-popup-edit max-width="158px" style="padding: 4px 8px;" self="top start" cover="false">
                      <div class="q-gutter-sm" style="margin-top:0px">
                        <a href="javascript:void(0);" @click="getEmo1(index)" style="text-decoration:none" v-for="(item,index) in faceList" :key="index" class="emotionItem">{{item}}</a>
                      </div>
                    </q-popup-edit>
                  </q-btn>
                  <q-btn push color="primary" label="Comment">
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
        v-for="(rev,index) in comment.rev"
        class="q-mx-lg"
        :key="index"
      >
      <q-item class="q-pa-md">
        <q-item-section avatar>
          <q-avatar rounded>
            <img alt="avatar" :src="rev.from_avatar" />
          </q-avatar>
        </q-item-section>
        <q-item-section>
          <q-item-label lines="1">
            <span class="text-weight-bold">{{ rev.from }}</span>
            <span class="text-weight-bold"> 回复：</span>
            <span class="text-gray-light"> {{ rev.to }}</span>
          </q-item-label>
          <q-item-label v-html="rev.texthtml"></q-item-label>
        </q-item-section>
        <div class="q-gutter-sm">
            <q-btn size="10px" round flat color="primary" icon="comment"  @click="commentss(idx,index)" />
            <q-btn v-if="rev.digg==0?true:false" size="10px" round flat color="#ccc" icon="favorite"  @click="favorite(idx)" />
            <q-btn v-else size="10px" round flat color="red" icon="favorite"  @click="favorite(idx)" />
        </div>
      </q-item>
      <div v-if="rev.isshow">
        <q-item>
            <q-item-section avatar>
                <q-avatar rounded>
                    <img src="https://cdn.quasar.dev/img/avatar5.jpg">
                  </q-avatar>
            </q-item-section>
            <q-item-section>
                <q-input autogrow v-model="text1" :label="'回复：'+rev.from" :dense="dense">
                <template v-slot:after>
                  <q-btn push color="primary" icon="mood">
                    <q-popup-edit max-width="158px" style="padding: 4px 8px;" self="top start" cover="false">
                      <div class="q-gutter-sm" style="margin-top:0px">
                        <a href="javascript:void(0);" @click="getEmo1(index)" style="text-decoration:none" v-for="(item,index) in faceList" :key="index" class="emotionItem">{{item}}</a>
                      </div>
                    </q-popup-edit>
                  </q-btn>
                  <q-btn push color="primary" label="Comment">
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
import emoji from '../css/emoji.json'
import { getCurrentInstance, reactive, toRefs, onMounted } from 'vue'
export default {
  name: 'Comment',
  components: {
  },
  setup () {
    let data = reactive({
        marktext: '',
        faceList: [],
        text: '',
        text1: '',
        isshow: false,
        isfav: false,
        isshowindex: 0,
        isshowindex1: 0,
        comments: [
          {
            'index':'123',
            'user':{
                'avatar_url':'https://cdn.quasar.dev/img/avatar2.jpg',
                'login':'login',
            },
            'updated_at':'123',
            'body_html':'为了开阔视野我们花很多时间去看新闻。但是这些时间花的值么？根据新闻的定义，新闻一般不会持续太久。而且由于现在新闻变得很容易发布和生产成本更低，其质量也下降了很多。我们很少停下来问自己关于我们所花费时间的事物：这个东西真的重要么？是否经得起时间的考验 —— 一周或者一年之内？写这个东西的人是否对此问题了如指掌？',
            'digg':0,
            'rev':[
              {
                'from':'fromuser',
                'from_avatar':'https://cdn.quasar.dev/img/avatar3.jpg',
                'to':'touser',
                'texthtml':'texthtml',
                'time':'2019-09-09',
                'digg':1
              },
              {
                'from':'touser',
                'from_avatar':'https://cdn.quasar.dev/img/avatar4.jpg',
                'to':'fromuser',
                'texthtml':'texthtml',
                'time':'2019-09-09',
                'digg':0
              }
            ],
          },
          {
            'index':'12',
            'user':{
                'avatar_url':'https://cdn.quasar.dev/img/avatar6.jpg',
                'login':'login',
            },
            'updated_at':'123',
            'body_html':'为了开阔视野我们花很多时间去看新闻。但是这些时间花的值么？根据新闻的定义，新闻一般不会持续太久。而且由于现在新闻变得很容易发布和生产成本更低，其质量也下降了很多。我们很少停下来问自己关于我们所花费时间的事物：这个东西真的重要么？是否经得起时间的考验 —— 一周或者一年之内？写这个东西的人是否对此问题了如指掌？',
            'digg':0,
            'rev':[]
          }
      ]
    }) as any
    const {ctx} = getCurrentInstance() as any
    console.log(ctx)
    const method = {
      commentss(fatherindex, childindex):void {
          console.log(fatherindex)
          console.log(childindex)
          data.text1 = ''
          if(childindex == null){
             data.comments[data.isshowindex].rev[data.isshowindex1].isshow = false
             data.comments[data.isshowindex].isshow = false
             data.comments[fatherindex].isshow = true
             data.isshowindex=fatherindex
          }else{
            data.comments[data.isshowindex].rev[data.isshowindex1].isshow = false
            data.comments[data.isshowindex].isshow = false
            data.comments[fatherindex].rev[childindex].isshow = true
            data.isshowindex=fatherindex
            data.isshowindex1=childindex
          }
      },
      favorite(comment):void {
          data.comments[comment].isfav = true

      },
      getTalk(): void {
          const res = data.comments
          data.comments = res.map(item=>{
              item.isshow = false
              const res1 = item.rev
              res1.forEach(element => {
                element.isshow = false
              });   
              return item
          })
      },
      getEmo(index): void {
        const face = data.faceList[index] as string
        data.text = data.text as string + face
      },
      getEmo1(index): void {
        const face = data.faceList[index] as string
        data.text1 = data.text1 as string + face
      },
      async getEmo2(): Promise<void> {
        /* const params = {
          addressId: data.marktext
        } */
      },
    }
    onMounted(()=>{
         method.getTalk()
         emoji.forEach(element => {
             data.faceList.push(element.char);
         });                   
      })
    return {
        ...toRefs(data),
      ...method
    }
  },
};
</script>
