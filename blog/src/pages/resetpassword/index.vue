<template>
  <q-layout>
    <q-header>
      <q-toolbar>
        <q-toolbar-title shrink class="text-weight-bold">
          <router-link to="/" style="color: #fff">ArtBlog</router-link>
        </q-toolbar-title>
      </q-toolbar>
    </q-header>
    <q-page-container class="bg-primary">
      <div class="row justify-center">
        <q-page padding>
          <q-card flat style="background-color: transparent">
            <q-card-actions class="text-white text-weight-bolder text-h5" align="center">
              Reset your password
            </q-card-actions>
          </q-card>
          <q-card style="width: 380px">
            <q-card-section>
              <q-form @submit="login">
                <q-input
                  outlined
                  dense
                  bg-color="grey-2"
                  v-model="formdata.username"
                  type="text"
                  label="username"
                  lazy-rules
                  :rules="[(val) => !!val || 'username is required']"
                />
                <q-btn
                  class="full-width"
                  @click="login()"
                  style="padding-top: 0"
                  label="sign up"
                  type="submit"
                  color="primary"
                />
              </q-form>
            </q-card-section>
          </q-card>
        </q-page>
      </div>
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { Notify } from 'quasar';
import { getCurrentInstance, reactive, toRefs } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { login } from '../../api/test/index';
import { LocalStorage, Cookies } from 'quasar';
export default {
  name: 'Login',
  components: {},
  setup() {
    const route = useRoute() as any;
    const router = useRouter() as any;
    let data = reactive({
      onload: false,
      isPwd: true,
      formdata: {
        username: '',
        password: '',
      },
    });
    const { ctx } = getCurrentInstance() as any;
    const method = {
      async login(): Promise<any> {
        let res = (await login(data.formdata)) as any;
        if (res.status != 200)
          return Notify.create({
            message: res.message,
            color: 'negative',
            icon: 'report_problem',
            position: 'top',
            timeout: 2000,
          });
        Cookies.set('token', res.token);
        LocalStorage.set('logged_in', 'yes');
        data.onload = true;
        router.push('/');
      },
    };
    return {
      ...toRefs(data),
      ...method,
    };
  },
};
</script>
<style lang="scss" scoped>
a {
  text-decoration: none;
}
a:hover {
  text-decoration: underline;
}
</style>