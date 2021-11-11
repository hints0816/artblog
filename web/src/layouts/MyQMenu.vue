<template>
  <div>
    <!-- 动态菜单 -->
    <q-list>
      <my-q-menu-item :my-router="menuList" :init-level="0"/>
    </q-list>
  </div>
</template>

<script>
import MyQMenuItem from './myQMenuItem'

export default {
  name: 'my-q-menu',
  components: { MyQMenuItem },
  data () {
    return {
      menuList: [
        {
          icon: 'inbox',
          label: '菜单1',
          children: [
            {
              icon: 'perm_identity',
              label: '菜单1-1',
              path: '/recommend'
            },
            {
              icon: 'delete',
              label: '菜单1-2',
              children: [
                {
                  icon: 'settings',
                  label: '菜单2-1',
                  path: '/song'
                }
              ]
            }
          ]
        },
        {
          isItemLabel: true,
          Label_text: 'item-label '
        },
        {
          icon: 'feedback',
          label: '菜单数据结构',
          path: '/test',
          hidden: false
        }
      ]
    }
  }
}
</script>
