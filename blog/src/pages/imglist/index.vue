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
            <q-img v-if="!checkClass(index)" :src="src" />
            <div v-else :class="inClass + ' column col-4'">
            <div class="col-6">
              <q-img
                src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
              />
            </div>
            <div class="col-6">
              <q-img
                src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
              />
            </div>
          </div>
          </div>
          <div class="col-8">
            <q-img
              src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
            />
          </div>
          <div :class="inClass + ' column col-4'">
            <div class="col-6">
              <q-img
                src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
              />
            </div>
            <div class="col-6">
              <q-img
                src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
              />
            </div>
          </div>
          <div class="col-4">
            <q-img
              src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
            />
          </div>
          <div class="col-4">
            <q-img
              src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
            />
          </div>
          <div class="col-4">
            <q-img
              src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
            />
          </div>
          <div class="col-4">
            <q-img
              src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
            />
          </div>
          <div class="col-4">
            <q-img
              src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
            />
          </div>
          <div class="col-4">
            <q-img
              src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
            />
          </div>
          <div :class="inClass + ' column col-4'">
            <div class="col-6">
              <q-img
                src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
              />
            </div>
            <div class="col-6">
              <q-img
                src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
              />
            </div>
          </div>
          <div class="col-8">
            <q-img
              src="http://47.119.167.128:9999/blog/479772348744994816.jpg"
            />
          </div>
        </div>
      </q-infinite-scroll>
    </q-page>
  </div>
</template>
<script lang="ts">
import { ref, reactive, toRefs, onMounted, getCurrentInstance } from 'vue';
import { Screen } from 'quasar';

export default {
  setup() {
    const { ctx } = getCurrentInstance() as any;
    let data = reactive({
      outClass: Screen.lt.lg
        ? 'q-col-gutter-' + Screen.name
        : 'q-col-gutter-lg',
      inClass: Screen.lt.lg
        ? 'q-col-gutter-y-' + Screen.name
        : 'q-col-gutter-y-lg',
    });
    const images = ref(
      Array(24)
        .fill(null)
        .map((_, i) => 'http://47.119.167.128:9999/blog/479772348744994816.jpg')
    );
    const method = {
      checkNum(nums: number): boolean {
        const r = /^\+?[0-9][0-9]*$/;
        if (r.test(String(nums))) {
          return true;
        } else {
          return false;
        }
      },
      loadClass(nums: number): string {
        if (method.checkNum(nums / 9)) {
          return 'col-8';
        } else if (
          method.checkNum((nums - 1) / 9) ||
          method.checkNum((nums - 2) / 9)
        ) {
          return data.inClass + ' column col-4';
        } else if (method.checkNum((nums - 2) / 9)) {
          return 'col-4';
        }
      },
      checkClass(nums: number): boolean {
        if (
          method.checkNum((nums - 1) / 9) ||
          method.checkNum((nums - 2) / 9)
        ) {
          return true;
        } else {
          return false;
        }
      },
    };
    onMounted(() => {
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
      images,
      ...method,
    };
  },
};
</script>
<style lang="sass">
</style>