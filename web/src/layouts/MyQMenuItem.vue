<template>
  <div>
  
    <div v-for="(item,index) in myRouter" :key="index">

	  <!-- 当菜单item标识为 ItemLabel 时，显示Label_text，不渲染为菜单 -->
      <q-item-label v-if="item.isItemLabel" header class="text-weight-bold text-uppercase">
        {{item.Label_text}}
      </q-item-label>

      <!-- 先判断当前菜单是否需要显示 -->
      <div v-if="!item.hidden && !item.isItemLabel">
      
        <!-- 当没有子菜单时，结束递归 -->
        <div v-if="!item.children">
          <q-expansion-item
            :header-inset-level="initLevel"
            expand-icon="1"
            :to="item.path"
            :icon="item.icon"
            :label="item.label"
          />
        </div>
        
        <!-- 子菜单依然存在时，继续递归 -->
        <div v-else>
          <q-expansion-item
            expand-separator
            :header-inset-level="initLevel"
            :to="item.path"
            :icon="item.icon"
            :label="item.label"
          >
            <!-- 进入下一次递归 -->
            <my-q-menu-item :my-router="item.children" :init-level="initLevel+0.3"/>
          </q-expansion-item>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
export default {
  name: 'my-q-menu-item',
  data () {
    return {}
  },
  props: ['myRouter', 'initLevel']
}
</script>