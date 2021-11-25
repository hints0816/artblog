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
          <q-card flat class="bg-primary">
            <q-card-section style="padding:5px" align="right">
              <span class="text-blue-1">Already have an account? </span>
              <router-link class="text-blue-1 text-weight-bold"
                 to="/login">Sign in →</router-link>
            </q-card-section>
          </q-card>
           <q-card style="width:380px">
              <q-card-section>
                <q-form
                  @submit="onSubmit"
                >
                <q-input outlined dense bg-color="grey-2" 
                  v-model="formdata.username" type="text" label="username"
                  lazy-rules
                  :rules="[val => !!val || 'username is required']"/>
                <q-input outlined dense bg-color="grey-2" 
                  v-model="formdata.email" type="email" label="email"
                  lazy-rules
                  :rules="[val => !!val || 'email is required']">
                  <template v-slot:prepend>
                    <q-icon name="mail" />
                  </template>
                </q-input>
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
import { Notify } from 'quasar'
import { getCurrentInstance, reactive, toRefs } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { adduser } from '../../api/test/index'
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
      isPwd: true,
      onload: false,
      formdata: {
        'username':'',
        'password':''
      },
    })
    const {ctx} = getCurrentInstance() as any
    const method = {
      async login():Promise<any> {
        let res  = await adduser(data.formdata) as any
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