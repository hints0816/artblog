<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-12" style="max-width: 1200px">
      <div class="row q-col-gutter-x-lg" style="margin-left: 0px;">
        <div class="col-xs-12 col-md-3" style="height: 100%;padding-left: 0;">
          <q-card style="background-color:transparent;" flat class="my-card">
            <q-card-actions align="center">
              <div class="useravatar full-width">
                 <q-list>
                  <q-item clickable v-ripple to="/setting/article">
                    <q-item-section avatar>
                      <q-icon name="feed" />
                    </q-item-section>
                    <q-item-section>
                      <q-item-label class="text-weight-bold">Article</q-item-label>
                    </q-item-section>
                  </q-item>

                  <q-item clickable v-ripple :active="true"  :active-class="$q.dark.isActive?'bg-grey-9':'bg-grey-3'">
                    <q-item-section avatar>
                      <q-icon name="category" />
                    </q-item-section>
                    <q-item-section>
                      <q-item-label class="text-weight-bold">Category</q-item-label>
                    </q-item-section>
                  </q-item>

                  <q-item clickable v-ripple to="/setting/comment">
                    <q-item-section avatar>
                      <q-icon name="chat" />
                    </q-item-section>
                    <q-item-section>
                      <q-item-label class="text-weight-bold">Comment</q-item-label>
                    </q-item-section>
                  </q-item>

                  <q-item clickable v-ripple to="/setting/artimg">
                    <q-item-section avatar>
                      <q-icon name="collections" />
                    </q-item-section>
                    <q-item-section>
                      <q-item-label class="text-weight-bold">ArtImage</q-item-label>
                    </q-item-section>
                  </q-item>
                </q-list>
              </div>
            </q-card-actions>
          </q-card>
        </div>
        <div class="col-xs-12 col-md-9 q-col-gutter-x-lg" style="height: 100%;margin-left:0;padding-left:0;">
          <q-card style="padding-right: 24px;" flat>
              <q-tab-panel name="setting">
                <Category/>
              </q-tab-panel>
          </q-card>
        </div>
      </div>
    </q-page>
  </div>
</template>
<script lang="ts">
import Category from '../../components/CategoryCard.vue';
import {
  ref,
  reactive,
  toRefs,
  onBeforeMount,
} from 'vue';

export default {
  name: 'Post',
  components: { Category },
  setup() {
    let data = reactive({
      alert: false,
      name: '',
      model: null,
      pagesize: 10,
      total: 10,
      edit: false,
    });
    const method = {
    };
    onBeforeMount(async () => {
    });
    return {
      ...toRefs(data),
      ...method,
    };
  },
};
</script>
<style lang="sass">
.my-menu-link
  color: white
  background: #F2C037
.useravatar .q-field--outlined .q-field__control
  border-radius: 100px
  padding: 0px
</style>