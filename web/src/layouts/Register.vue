<template>
  <q-layout>
    <q-page padding>
      <div class="row justify-center">
          <q-card class="col-md-4 col-sm-12" style="padding: 10px">
            <q-form
              @submit="onSubmit"
              @reset="onReset"
              class="q-gutter-md"
            >
              <q-input
                v-model="name"
                label="邮箱 *"
                hint="邮箱将作为你的用户名"
                lazy-rules
                :rules="[ val => val && val.length > 0 || '请输入你的邮箱']"
              >
              <template v-slot:after>
                  <q-btn dense flat label="发送验证码" icon="send" @click="onClick"/>
                </template>
              </q-input>

              <!-- <q-input
                type="number"
                v-model="model.authCode"
                label="邮箱验证码 *"
                lazy-rules
                :rules="[ val => val && val.length > 0 || '请输入验证码']"
              /> -->

              <q-input label="密码 *" v-model="model.password" :type="isPwd ? 'password' : 'text'" 
                lazy-rules
                :rules="[ val => val && val.length >= 6 || '密码长度不能少于6位']">
                <template v-slot:append>
                  <q-icon
                    :name="isPwd ? 'visibility_off' : 'visibility'"
                    class="cursor-pointer"
                    @click="isPwd = !isPwd"
                  />
                </template>
              </q-input>

              <q-input label="再次输入密码 *" v-model="model.repassword" :type="isPwd ? 'password' : 'text'" 
                lazy-rules
                :rules="[ val => val === this.model.password || '两次密码不同']">
                <template v-slot:append>
                  <q-icon
                    :name="isPwd ? 'visibility_off' : 'visibility'"
                    class="cursor-pointer"
                    @click="isPwd = !isPwd"
                  />
                </template>
              </q-input>
              <div>
                <q-btn label="Submit" type="submit" color="primary"/>
                <q-btn label="Reset" type="reset" color="primary" flat class="q-ml-sm" />
              </div>
            </q-form>
          </q-card>
      </div>
    </q-page>
  </q-layout>
</template>

<script>
import { onMounted } from 'vue'
import { getCart } from '../api/test'

export default {
  data() {
    return {
      isPwd: true,
      model: {

      }
    };
  },
  setup() {
    onMounted(() => {
      return init()
    })

   const init = async () => {
      let data  = await getCart()
      console.log(data)
    }
  }
};
</script>

<style lang="sass" scoped>
</style>