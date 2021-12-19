<template>
  <div>
    <q-table
      flat
      :rows="rows"
      :columns="columns"
      :rows-per-page-options="[]"
      row-key="name"
    >
      <template v-slot:header="props">
        <q-tr :props="props">
          <q-th auto-width />
          <q-th v-for="col in props.cols" :key="col.name" :props="props">
            {{ col.label }}
          </q-th>
        </q-tr>
      </template>
      <template v-slot:body="props">
        <q-tr :props="props">
          <q-td auto-width>
            <q-checkbox v-model="props.expand" />
          </q-td>
          <q-td key="id" :props="props">
            {{ props.row.id }}
          </q-td>
          <q-td key="name" :props="props">
            {{ props.row.name }}
          </q-td>
          <q-td key="color" :props="props">
            {{ props.row.color }}
          </q-td>
          <q-td key="text_color" :props="props">
            {{ props.row.text_color }}
          </q-td>
          <q-td key="icon" :props="props">
            {{ props.row.icon }}
          </q-td>
          <q-td key="icon" class="q-gutter-xs" :props="props">
            <q-btn
              dense
              round
              color="primary"
              @click="alertEdit(props.row)"
              icon="edit"
            />
            <q-dialog v-model="alert">
              <q-card style="min-width: 350px">
                <q-card-section>
                  <div class="text-h6">Edit Category</div>
                </q-card-section>
                <q-card-section class="q-pt-none q-gutter-y-xs">
                  <q-input
                    dense
                    outlined
                    disable
                    v-model="category.id"
                    label="id"
                  />
                  <q-input
                    dense
                    outlined
                    v-model="category.name"
                    label="name"
                  />
                  <q-input
                    dense
                    outlined
                    v-model="category.color"
                    label="color"
                  />
                  <q-input
                    dense
                    outlined
                    v-model="category.text_color"
                    label="text_color"
                  />
                  <q-input
                    dense
                    outlined
                    v-model="category.icon"
                    label="icon"
                  />
                </q-card-section>
                <q-card-actions align="right">
                  <q-btn
                    flat
                    label="OK"
                    @click="editCategory"
                    color="primary"
                    v-close-popup
                  />
                </q-card-actions>
              </q-card>
            </q-dialog>
            <q-btn
              dense
              round
              @click="showDelNotify(props.row.id)"
              color="deep-orange"
              icon="delete"
            />
          </q-td>
        </q-tr>
      </template>
    </q-table>
  </div>
</template>
<script lang="ts">
import {
  listCategory,
  updateCategory,
  deleteCategory,
} from '../api/test/index';
import { reactive, toRefs, onBeforeMount } from 'vue';
import { useRoute } from 'vue-router';
import { Notify } from 'quasar';
const columns = [
  { name: 'id', align: 'left', label: 'id', field: 'id' },
  { name: 'name', align: 'center', label: 'name', field: 'name' },
  { name: 'color', align: 'center', label: 'color', field: 'color' },
  {
    name: 'text_color',
    align: 'center',
    label: 'text_color',
    field: 'text_color',
  },
  { name: 'icon', align: 'center', label: 'icon', field: 'icon' },
  { name: 'operate', align: 'center', label: 'operate', field: 'operate' },
];
export default {
  name: 'Category',
  setup() {
    let data = reactive({
      alert: false,
      rows: [],
      category: {
        id: 0,
        name: '',
        color: '',
        text_color: '',
        icon: '',
      },
    });
    const route = useRoute() as any;
    const method = {
      async commentAlert(): Promise<void> {
        const params = {
          pagenum: 1,
          pagesize: 20,
        };
        let res = (await listCategory(params)) as any;
        data.rows = res.data;
      },
      alertEdit(propss: any): void {
        data.category = propss;
        data.alert = true;
      },
      async editCategory(): Promise<void> {
        let res = (await updateCategory(data.category)) as any;
        if (res.status == 200) {
          Notify.create({
            message: 'Edit Successful',
            type: 'positive',
            position: 'top',
            timeout: 2000,
          });
        }
      },
      showDelNotify(id: number) {
        Notify.create({
          message: 'Are you absolutely sure?',
          color: 'negative',
          icon: 'report_problem',
          position: 'top',
          actions: [
            {
              label: 'yes',
              color: 'white',
              handler: async () => {
                let res = (await deleteCategory(id)) as any;
                if (res.status == 200) {
                  return method.commentAlert();
                }
              },
            },
          ],
        });
      },
    };
    onBeforeMount(async () => {
      await method.commentAlert();
    });
    return {
      ...toRefs(data),
      ...method,
      columns,
    };
  },
};
</script>