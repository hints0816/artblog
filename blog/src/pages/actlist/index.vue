<template>
  <div class="row justify-center">
    <q-page v-show="post.id" padding class="col-xs-12 col-sm-10 col-md-8">
      <div class="q-mb-lg">
        <h4 class="text-cyan-9">{{ post.title }}</h4>
        <code class="text-italic">
          Updated by {{ $store.getters.username }} {{ post.updated_at }}
        </code>
      </div>
      <div v-html="post.body_html" class="q-mt-lg" />
      <div>
        <q-chip
          v-for="label in post.labels"
          outline
          square
          clickable
          class="label"
          :key="label.index"
          :style="`border-color: #${label.color}; color: #${label.color};`"
          @click="chipClickHandler(label.name)"
        >
          {{ label.name }}
        </q-chip>
      </div>
      <q-separator color="cyan-9" style="height: 1px;" />
      <Comment />
    </q-page>
  </div>
</template>
<script>
import Comment from '../../components/CommentList.vue';
export default {
  name: 'Post',
  components: { Comment },
  data() {
    return {
      post: {
        'id':'123',
        'title':'为什么你应该停止阅读新闻？',
        'updated_at':'123',
        'body_html':'<h3>为什么你应该停止阅读新闻？</h3>'
      },
    };
  },
  watch: {
  },
  methods: {
  },
  created() {
  },
};
</script>

