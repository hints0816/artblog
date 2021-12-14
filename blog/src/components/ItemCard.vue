<template>
  <div>
    <div v-for="post in postList" :key="post.index">
      <q-item clickable v-ripple class="q-mt-md">
        <q-item-section @click="toPostDetail(post.ID)">
          <q-item-label>
            <div class="text-h6 text-cyan-9">{{ post.title }}</div>
          </q-item-label>
          <q-item-label
            lines="4"
            caption
            class="text-body1 text-gray-light text-justify"
          >
            <q-item>
              <q-item-section v-if="post.img" top thumbnail class="q-ml-none">
                <img style="width:160px;height:90px;" :src="post.img">
              </q-item-section>
              <q-item-section>
                <q-item-label style="height: 70%;" lines="3">{{ post.desc }}</q-item-label>
                <q-item-label style="height: 25%;" lines="1">
                  <router-link class="text-bold" to="/" :style="{ 'color': $q.dark.isActive?'#fff':'#3b6c9f' }">{{ post.Profile.name }}</router-link>
                </q-item-label>
              </q-item-section>
            </q-item>
          </q-item-label>
        </q-item-section>
        <q-item-section style="position: relative;" side top>
          <q-item-label @click="toPostDetail(post.ID)" caption>
            {{ post.UpdatedAt }}
          </q-item-label>
          <q-chip
            dense
            v-for="label in post.Cateart"
            clickable
            class="label"
            :color="label.Category.color"
            :text-color="label.Category.text_color"
            :icon="label.Category.icon"
            :key="label.index"
            @click="chipClickHandler(label.Category.id)"
          >
            {{ label.Category.name }}
          </q-chip>
          <q-btn style="position: absolute;bottom: 0px" class="gt-xs" size="12px" flat dense round icon="fas fa-ellipsis-h" />
        </q-item-section>
      </q-item>
      <!-- <q-card-actions align="left">
              <q-btn flat color="red" icon="favorite" />
              <q-btn flat color="teal" icon="comment" />
          </q-card-actions> -->
      <q-separator spaced inset />
    </div>
  </div>
</template>

<script lang="ts">
import { getCurrentInstance } from 'vue';
import { useRouter } from 'vue-router';

export default {
  name: 'Item',
  props: {
    postList: {
      type: Array,
      required: true,
    },
  },
  setup() {
    const router = useRouter() as any;
    const { ctx } = getCurrentInstance() as any;
    const method = {
      toPostDetail(id: number): void {
        router.push(`/posts/${id}`);
      },
      chipClickHandler(name: string): void {
        ctx.$emit('reloadart', name)
      }
    };
    return {
      ...method,
    };
  },
  methods: {},
};
</script>

<style lang="scss" scoped>
.markdown-body hr {
  height: unset;
}
</style>