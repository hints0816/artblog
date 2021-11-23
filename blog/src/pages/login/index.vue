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
                <q-input dense v-model="formdata.username" filled type="email" hint="Email" />
                  <q-input dense v-model="formdata.password" filled :type="isPwd ? 'password' : 'text'" hint="Password with toggle">
                    <template v-slot:append>
                      <q-icon
                        :name="isPwd ? 'visibility_off' : 'visibility'"
                        class="cursor-pointer"
                        @click="isPwd = !isPwd"
                      />
                    </template>
                  </q-input>
                  <q-btn class="q-mt-md full-width" @click="login()" label="sign in" type="submit" color="primary"/>
              </q-card-section>
            </q-card>
            <q-card class="q-mt-md" style="width:380px" align="center">
              <q-card-section>
                <span class="text-blue-grey-10">New to GitHub? </span>
                <router-link class="text-blue-10 text-weight-bold"
                 to="/signup" style="color: #fff;">Create an account.</router-link>
              </q-card-section>
            </q-card>
        </q-page>
      </div>
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { Notify } from 'quasar'
import { getCurrentInstance, reactive, toRefs } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { login } from '../../api/test/index'
import { LocalStorage } from 'quasar'
export default {
  name: 'Login',
  components: {
  },
  setup () {
    const route = useRoute() as any
    const router = useRouter() as any
    console.log(route)
    let data = reactive({
      onload: false,
      formdata: {
        'username':'',
        'password':''
      },
    })
    const {ctx} = getCurrentInstance() as any
    const method = {
      async login():Promise<any> {
        let res  = await login(data.formdata) as any
        if (res.status != 200) 
        return Notify.create({
          message: res.message,
          color: 'negative',
          icon: 'report_problem',
          position: 'top',
          timeout: 2000
        })
        console.log(res.token)
        LocalStorage.set('token', res.token)
        console.log(LocalStorage.getItem('token'))
        data.onload = true
        router.push('/');
        console.log(ctx)
      }
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