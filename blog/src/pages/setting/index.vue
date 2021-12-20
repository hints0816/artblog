<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-12" style="max-width: 1200px">
      <div class="row q-col-gutter-x-lg" style="margin-left: 0px;">
        <div class="col-xs-12 col-md-3" style="height: 100%;padding-left: 0;">
          <q-card style="background-color:transparent;" flat class="my-card">
            <q-card-actions align="center">
              <div class="useravatar full-width">
                 <q-list bordered separator>
                  <q-item clickable v-ripple :active="true" :active-class="$q.dark.isActive?'bg-grey-9':'bg-grey-3'">
                    <q-item-section>Article</q-item-section>
                  </q-item>

                  <q-item clickable v-ripple to="/setting/category">
                    <q-item-section>
                      <q-item-label>Category</q-item-label>
                    </q-item-section>
                  </q-item>

                  <q-item clickable v-ripple to="/setting/comment">
                    <q-item-section>
                      <q-item-label>Comment</q-item-label>
                    </q-item-section>
                  </q-item>
                </q-list>
              </div>
            </q-card-actions>
          </q-card>
        </div>
        <div class="col-xs-12 col-md-9 q-col-gutter-x-lg" style="height: 100%;margin-left:0;padding-left:0;">
          <q-card style="padding-right: 24px;" flat>
                <q-card-actions class="row q-col-gutter-x-xs justify-between">
                  <q-input
                    class="col-grow col-xs-12 col-sm-8"
                    dense
                    outlined
                    v-model="text"
                    label="Outlined"
                  />
                  <q-select
                    dense
                    outlined
                    v-model="model"
                    class="col-grow"
                    :options="['Google', 'Facebook']"
                    style="width: 130px"
                    label="Outlined"
                  />
                  <q-btn outline to="/edit" class="text-purple q-mx-xs">
                    <div class="row items-center no-wrap">
                      <q-icon left name="book" />
                      <div class="text-center">New</div>
                    </div>
                  </q-btn>
                </q-card-actions>
                <q-separator />
                <q-card-section>
                  <Edit @reloadart="listart" :postList="postList" :edit="edit" />
                </q-card-section>
          </q-card>
        </div>
      </div>
    </q-page>
  </div>
</template>
<script lang="ts">
import Edit from '../../components/EditCard.vue';
import {
  listArticle,
} from '../../api/test/index';
import {
  reactive,
  toRefs,
  onBeforeMount,
} from 'vue';
import { date } from 'quasar';
import { useRoute } from 'vue-router';

export default {
  name: 'Post',
  components: { Edit },
  setup() {
    let data = reactive({
      alert: false,
      name: '',
      model: null,
      pagesize: 10,
      total: 10,
      edit: false,
      postList: [],
    });
    const route = useRoute() as any;
    const method = {
      async listart(): Promise<any> {
        const paramss = {
          pagenum: 1,
          pagesize: data.pagesize,
          user_id: route.params.id,
        };
        let datas = (await listArticle(paramss)) as any;
        datas.data.forEach((element) => {
          let timeStamp = new Date(element.UpdatedAt);
          let formattedString = date.formatDate(
            timeStamp,
            'YYYY-MM-DD HH:mm:ss'
          );
          element.UpdatedAt = formattedString;
        });
        data.postList = datas.data;
        data.total = datas.total;
        data.edit = datas.edit;
      }
    };
    onBeforeMount(async () => {
      await method.listart();
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