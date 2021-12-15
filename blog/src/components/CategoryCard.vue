<template>
  <div>
   <q-table
      flat 
      :rows="rows"
      :columns="columns"
      title="CategoryEdit"
      :rows-per-page-options="[]"
      row-key="name"
    >
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td key="id" :props="props">
            {{ props.row.id }}
          </q-td>
          <q-td key="name" :props="props">
            {{ props.row.name }}
            <q-popup-edit v-model="props.row.name" buttons v-slot="scope" @save="updateCategory">
              <q-input v-model="scope.value" dense autofocus/>
            </q-popup-edit>
          </q-td>
          <q-td key="color" :props="props">
            {{ props.row.color }}
            <q-popup-edit v-model="props.row.color" buttons v-slot="scope">
              <q-input v-model="scope.value" dense autofocus @keyup.enter="scope.set" />
            </q-popup-edit>
          </q-td>
          <q-td key="text_color" :props="props">
            {{ props.row.text_color }}
            <q-popup-edit v-model="props.row.text_color" buttons v-slot="scope">
              <q-input v-model="scope.value" dense autofocus @keyup.enter="scope.set" />
            </q-popup-edit>
          </q-td>
          <q-td key="icon" :props="props">
            {{ props.row.icon }}
            <q-popup-edit v-model="props.row.icon" buttons v-slot="scope">
              <q-input v-model="scope.value" dense autofocus @keyup.enter="scope.set" />
            </q-popup-edit>
          </q-td>
        </q-tr>
      </template>
    </q-table>
  </div>
</template>
<script lang="ts">
import { listCategory, updateCategory } from '../api/test/index'
import { reactive, toRefs, onBeforeMount } from 'vue'
import { useRoute } from 'vue-router';
const columns = [
  { name: 'id', align: 'left', label: 'id', field: 'id' },
  { name: 'name', align: 'center', label: 'name', field: 'name' },
  { name: 'color', align: 'center', label: 'color', field: 'color' },
  { name: 'text_color', align: 'center', label: 'text_color', field: 'text_color' },
  { name: 'icon', align: 'center', label: 'icon', field: 'icon' }
]
export default {
  name: 'Category',
setup () {
    let data = reactive({
      rows: []
    })
    const route = useRoute() as any;
    const method = {
      async commentAlert(): Promise<void> {
        const params = {
          pagenum: 1,
          pagesize: 20,
        };
        let res  = await listCategory(params) as any
        data.rows = res.data
      },
      updateCategory(value: any, initialValue: any): void {
        console.log(value)
                console.log(initialValue)
      }
    }
    onBeforeMount(async()=>{
        await method.commentAlert()
    })
    return {
      ...toRefs(data),
      ...method,
      columns
    }
  }
}
</script>