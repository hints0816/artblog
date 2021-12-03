<template>
  <q-layout>
    <q-header>
      <q-toolbar>
        <q-toolbar-title shrink class="text-weight-bold">
          <router-link to="/" style="color: #fff;">blogName</router-link>
        </q-toolbar-title>
      </q-toolbar>
    </q-header>
    <q-page-container class="bg-primary"> 
      <div class="row justify-center">
        <q-page padding>
           <q-card style="width:380px">
              <q-card-section>
                <q-form
                  @submit="onSubmit"
                >
                  <q-input outlined dense bg-color="grey-2" 
                    v-model="formdata.password" 
                    :type="isPwd ? 'password' : 'text'" label="password"
                    lazy-rules
                    :rules="[val => !!val || 'password is required']">
                    <template v-slot:append>
                      <q-icon
                        :name="isPwd ? 'visibility_off' : 'visibility'"
                        class="cursor-pointer"
                        @click="isPwd = !isPwd"
                      />
                    </template>
                  </q-input>
                  <q-input outlined dense bg-color="grey-2" 
                    v-model="formdata.password" 
                    :type="isPwd ? 'password' : 'text'" label="comfirm password"
                    lazy-rules
                    :rules="[val => !!val || 'comfirm password is required']">
                    <template v-slot:append>
                      <q-icon
                        :name="isPwd ? 'visibility_off' : 'visibility'"
                        class="cursor-pointer"
                        @click="isPwd = !isPwd"
                      />
                    </template>
                  </q-input>
                  <q-btn class="full-width" @click="login" style="padding-top: 0" label="sign up" type="submit" color="primary"/>
                </q-form>
              </q-card-section>
            </q-card>
        </q-page>
      </div>
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { getCurrentInstance, reactive, toRefs } from 'vue'
import { useRoute, useRouter } from 'vue-router'
export default {
  name: 'Login',
  components: {
  },
  setup () {
    const route = useRoute() as any
    const router = useRouter() as any
    console.log(route)
    let data = reactive({
      isPwd: true,
      onload: false,
      formdata: {
        'username':'',
        'password':''
      },
    })
    const {ctx} = getCurrentInstance() as any
    console.log(ctx)
    const method = {
      
    }
    return {
      ...toRefs(data),
      ...method
    }
  }
}
</script>
<style lang="scss" scoped>
a{
  text-decoration: none
}
a:hover{
  text-decoration: underline
}
</style>