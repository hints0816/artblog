<template>
  <q-layout>
    <q-header>
      <q-toolbar :style="{ 'background-color': $q.dark.isActive?'#1D1D1D':'primary' }">
        <q-toolbar-title shrink class="text-weight-bold">
          <router-link to="/" style="color: #fff">blogName</router-link>
        </q-toolbar-title>
      </q-toolbar>
    </q-header>
    <q-page-container :class="$q.dark.isActive?'bg-dark':'bg-primary'">
      <div class="row justify-center">
        <q-page padding>
          <q-card flat style="width: 380px">
            <q-card-section>
              <q-form @submit="login">
                <q-input
                  outlined
                  dense
                  :bg-color="$q.dark.isActive?'dark':'grey-2'"
                  v-model="formdata.username"
                  type="text"
                  label="username"
                  lazy-rules
                  :rules="[(val) => !!val || 'username is required']"
                />
                <div class="column items-end">
                  <div class="col">
                    <router-link
                      class="text-blue-10 text-weight-bold"
                      to="/password_reset"
                      style="color: #fff"
                      >Forgot password?
                    </router-link>
                  </div>
                </div>
                <q-input
                  outlined
                  dense
                  :bg-color="$q.dark.isActive?'dark':'grey-2'"
                  v-model="formdata.password"
                  :type="isPwd ? 'password' : 'text'"
                  label="password"
                  lazy-rules
                  :rules="[(val) => !!val || 'password is required']"
                >
                  <template v-slot:append>
                    <q-icon
                      :name="isPwd ? 'visibility_off' : 'visibility'"
                      class="cursor-pointer"
                      @click="isPwd = !isPwd"
                    />
                  </template>
                </q-input>
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
          <q-card flat class="q-mt-md" style="width: 380px;" align="center">
            <q-card-section style="padding: 5px;">
              <q-btn @click="toGithub" dense round flat icon="ion-logo-github"/>
              <span :class="$q.dark.isActive?'text-grey-4':'text-blue-grey-10'">New to GitHub? </span>
              <router-link
                :class="$q.dark.isActive?'text-grey-4 text-weight-bold':'text-blue-10 text-weight-bold'"
                to="/signup"
                style="color: #fff"
                >Create an account.</router-link
              >
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
    const { ctx } = getCurrentInstance() as any;

    let data = reactive({
      onload: false,
      isPwd: true,
      formdata: {
        username: '',
        password: '',
      },
    });
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
        if(route.query.login_redirect == 'claim') {
          router.push('/')
        } else {
          router.go(-1)
        }
      },
      toGithub():void {
        window.location.href='https://github.com/login/oauth/authorize?client_id=6c85823c1a121313be52&state=s76de&redirect_uri=http://www.arthins.com:3000/api/blog/github'
      }
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