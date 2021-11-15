<template>
  <div>
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
        <div class="q-pa-md q-gutter-sm">
            <q-btn size="10px" round flat color="primary" icon="comment"  @click="commentss(idx)" />
            <q-btn v-if="comment.isshow" size="10px" round flat color="#ccc" icon="favorite"  @click="favorite(idx)" />
            <q-btn v-else size="10px" round flat color="red" icon="favorite"  @click="favorite(idx)" />
        </div>
      </q-item>
      <div v-if="comment.isshow">
        <q-item class="q-pa-md">
            <q-item-section avatar>
                <q-avatar rounded>
                    <img alt="avatar" :src="comment.user.avatar_url" />
                </q-avatar>
            </q-item-section>
            <q-item-section>
                <q-editor v-model="editor" min-height="5rem" />
            </q-item-section>
        </q-item>
      </div>
    </q-list>
    <div class="q-pa-md q-gutter-sm">
      <!-- <q-btn outline no-caps color="cyan-9" @click="goAddComment">
        <q-icon left size="2rem" name="add_comment" />
        <div>Add Comment</div>
      </q-btn> -->
    </div>
  </div>
</template>

<script lang="ts">
import { getCurrentInstance, reactive, toRefs, onMounted } from 'vue'
export default {
  name: 'Comment',
  data() {
    return {
      comments: [
          {
            'index':'123',
            'user':{
                'avatar_url':'avatar_url',
                'login':'login',
            },
            'updated_at':'123',
            'body_html':'为了开阔视野我们花很多时间去看新闻。但是这些时间花的值么？根据新闻的定义，新闻一般不会持续太久。而且由于现在新闻变得很容易发布和生产成本更低，其质量也下降了很多。我们很少停下来问自己关于我们所花费时间的事物：这个东西真的重要么？是否经得起时间的考验 —— 一周或者一年之内？写这个东西的人是否对此问题了如指掌？'
          },
          {
            'index':'12',
            'user':{
                'avatar_url':'avatar_url',
                'login':'login',
            },
            'updated_at':'123',
            'body_html':'为了开阔视野我们花很多时间去看新闻。但是这些时间花的值么？根据新闻的定义，新闻一般不会持续太久。而且由于现在新闻变得很容易发布和生产成本更低，其质量也下降了很多。我们很少停下来问自己关于我们所花费时间的事物：这个东西真的重要么？是否经得起时间的考验 —— 一周或者一年之内？写这个东西的人是否对此问题了如指掌？'
          }
      ]
    };
  },
  setup () {
      
    const data = reactive({
        isshow: false,
        isfav: false,
        isshow1: 0,
        comments:[]
    }) as any
    const {ctx} = getCurrentInstance() as any
    const method = {
      commentss(comment):void {
data.comments[data.isshow1].isshow = false
          data.comments[comment].isshow = true
          data.isshow1=comment
      },
      favorite(comment):void {
          data.comments[comment].isfav = true

      },
      getTalk(): void {
          const res = ctx.comments
          data.comments = res.map(item=>{
              item.isshow = false
              return item
          })
      }
    }
    onMounted(()=>{
          method.getTalk()
      })
    return {
        ...toRefs(data),
      ...method
    }
  },
};
</script>
