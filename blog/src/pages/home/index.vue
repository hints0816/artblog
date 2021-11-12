<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-10 col-md-8">
      <q-list padding>
        <Item :postList="postList" />
      </q-list>
    </q-page>
  </div>
</template>
<script>
import Item from '../../components/ItemCard.vue';
export default {
  name: 'List',
  components: { Item },
  data() {
    return {
      postList: [
        {
          'number':1,
          'title':'tilte',
          'created_at':'2021-09-02',
          'body_html':'Secondary line text. Lorem ipsum dolor sit amet, consectetur adipiscit elit.Secondary line text. Lorem ipsum dolor sit amet, consectet',
          'labels':[
            {
              'icon':'bookmark',
              'name':'bookmark'
            },
            {
              'icon':'event',
              'textcolor':'white',
              'color':'primary',
              'name':'Add to calendar'
            }
          ]
        }
      ],
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

