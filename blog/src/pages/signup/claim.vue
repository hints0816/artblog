<template>
  <q-layout>
    <q-header>
      <q-toolbar :class="$q.dark.isActive ? 'bg-dark' : 'bg-primary'">
        <q-toolbar-title shrink class="text-weight-bold">
          <router-link to="/" style="color: #fff">ArtBlog</router-link>
        </q-toolbar-title>
      </q-toolbar>
    </q-header>
    <q-page-container :class="$q.dark.isActive ? 'bg-dark' : 'bg-primary'">
      <div class="row justify-center">
        <q-page padding>
          <q-card flat v-if="valid" style="width: 380px">
            <q-card-section>
              <q-form @submit="onSubmit">
                <q-input
                  outlined
                  dense
                  :bg-color="$q.dark.isActive ? 'dark' : 'grey-2'"
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
                <q-input
                  outlined
                  dense
                  :bg-color="$q.dark.isActive ? 'dark' : 'grey-2'"
                  v-model="formdata.comfirm_password"
                  :type="isPwd ? 'password' : 'text'"
                  label="comfirm password"
                  lazy-rules
                  :rules="[
                    (val) => !!val || 'comfirm password is required',
                    (val) =>
                      val == formdata.password ||
                      'Make sure the password twice',
                  ]"
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
                  style="padding-top: 0"
                  label="sign up"
                  type="submit"
                  color="primary"
                />
              </q-form>
            </q-card-section>
          </q-card>
          <q-card v-else class="q-pa-lg" style="width: 380px">
            <q-item align="center">
              <q-item-section>
                <q-item-label class="text-weight-bold text-red-6">
                  OOPSIES! OUR BAD
                </q-item-label>
                <q-item-label class="text-h6">
                  This Set-up Link is Bad
                </q-item-label>
              </q-item-section>
            </q-item>
            <q-card-actions align="center">
              <div align="center" class="text-subtitle2">
                Hmm, this account set up link isn't valid.
              </div>
            </q-card-actions>
          </q-card>
        </q-page>
      </div>
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { getCurrentInstance, reactive, toRefs, onBeforeMount } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { signup, checkToken } from '../../api/test/index';
import { Notify } from 'quasar';
export default {
  name: 'Login',
  components: {},
  setup() {
    const router = useRouter() as any;
    const route = useRoute() as any;
    let data = reactive({
      isPwd: true,
      valid: true,
      formdata: {
        password: '',
      },
    });
    const { ctx } = getCurrentInstance() as any;
    const method = {
      async onSubmit(): Promise<any> {
        const params = {
          token: route.query.token as string,
          password: data.formdata.password,
        };
        let res = (await signup(params)) as any;
        if (res.status == 200) {
          Notify.create({
            message: 'Sign Up Successful',
            color: 'positive',
            icon: 'report_problem',
            position: 'top',
            timeout: 3000,
          });
          router.push({ path: '/login', query: { login_redirect: 'claim' } });
        }
      },
      async checkToken(): Promise<any> {
        const params = {
          token: route.query.token as string,
        };
        let res = (await checkToken(params)) as any;
        if (res.status == 200) {
          data.valid = true;
        }
      },
    };
    onBeforeMount(async () => {
      await method.checkToken();
    });
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