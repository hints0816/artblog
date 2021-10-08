<template>
  <q-page class="flex">
    <div class="row">
      <div class="col-xs-12 col-sm-6 col-md-4">
        <q-card class="my-card" flat bordered>
          <q-item>
            <q-item-section avatar>
              <q-avatar>
                <img src="https://cdn.quasar.dev/img/boy-avatar.png">
              </q-avatar>
            </q-item-section>
            <q-item-section>
              <q-item-label>Title</q-item-label>
              <q-item-label caption>
                Subhead
              </q-item-label>
            </q-item-section>
            <div class="col-auto">
                <q-btn color="grey-7" round flat icon="more_vert">
                  <q-menu cover auto-close>
                    <q-list>
                      <q-item clickable>
                        <q-item-section>Remove Card</q-item-section>
                      </q-item>
                      <q-item clickable>
                        <q-item-section>Send Feedback</q-item-section>
                      </q-item>
                      <q-item clickable>
                        <q-item-section>Share</q-item-section>
                      </q-item>
                    </q-list>
                  </q-menu>
                </q-btn>
              </div>
            </q-item>
          <img src="https://cdn.quasar.dev/img/mountains.jpg">
          <q-card-actions align="right">
            <q-btn flat round color="red" icon="favorite" />
            <q-btn flat round color="teal" icon="bookmark" />
            <q-btn flat round color="primary" icon="comment" />
          </q-card-actions>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<style>
</style>

<script>
export default {
  name: 'HelloWorld',
  data() {
    return {
      artList: [],
      queryParam: {
        pagesize: 5,
        pagenum: 1
      },
      total: 0,
      isLoad: false
    }
  },
  watch: {
  },
  mounted() {
     this.getArtList()
  },
  methods: {
    async getArtList() {
      const { data: res } = await this.$http.get(`article/list/${this.cid}`, {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      this.artList = res.data
      this.total = res.total
      this.isLoad = true
    }
  }
}
</script>
