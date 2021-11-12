<template>
  <q-page v-show="post.id" padding>
    <div class="q-mb-lg">
      <h1 class="text-cyan-9">{{ post.title }}</h1>
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
</template>
<script>
export default {
  name: 'Post',
  data() {
    return {
      post: {
        'id':'123',
        'title':'123',
        'updated_at':'123',
        'body_html':'123'
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

