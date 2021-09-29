<template style="background-color:#1b1b1b">
  <v-col>
    <div class="text-center">
        {{ msg }}1
    </div>
  </v-col>
</template>
<script>
export default {
  data() {
    return {
      artList: [],
      queryParam: {
        pagesize: 5,
        pagenum: 1
      },

      count: 0,
      total: 0
    }
  },
  props: {
    msg: String
  },
  created() {},
  mounted() {
    this.getArtList()
  },
  filters: {},
  methods: {
    async getArtList() {
      console.log(1)
      const { data: res } = await this.$http.get('user/1')
      this.artList = res.data
      this.total = res.total
    }
  }
}
</script>
<style lang=""></style>
