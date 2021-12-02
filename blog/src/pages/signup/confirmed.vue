<template>
  <q-page padding>
    <q-card v-if="email == ''" class="q-pa-lg" style="width: 380px">
        <q-item align="center">
        <q-item-section>
          <q-item-label class="text-weight-bold text-red-6">
            OOPSIES!
          </q-item-label>
          <q-item-label class="text-h6">Something Went Wrong</q-item-label>
        </q-item-section>
      </q-item>
      <q-card-actions align="center">
        <div align="center" class="text-subtitle2">Looks like we couldn't complete the next steps of sending you an invite to claim a Font Awesome account and kit.</div>
      </q-card-actions>
    </q-card>
    <q-card v-else class="q-pa-lg" style="width: 380px">
        <q-item align="center">
        <q-item-section>
          <q-item-label class="text-weight-bold text-blue-10">
            <q-icon size="18px" name="mail"/>You've got mail
          </q-item-label>
          <q-item-label class="text-h6">Check Your Email</q-item-label>
        </q-item-section>
      </q-item>
      <q-card-actions align="center">
        <div align="center" class="text-subtitle2">We just sent a confirmation link to {{ email }}. Verify your address and we'll get you all set up!</div>
      </q-card-actions>
      <q-card-section>
        <q-btn
          class="full-width"
          @click="login"
          style="padding-top: 0"
          label="Resend Confirmation Email"
          type="submit"
          icon-right="send"
          color="primary"
        />
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script lang="ts">
import { Notify } from 'quasar';
import { getCurrentInstance, reactive, toRefs } from 'vue';
import { useRoute } from 'vue-router';
import { sendvalidator } from '../../api/test/index';
export default {
  name: 'Login',
  components: {},
  props: {
    email: {
      type: String,
      required: true,
    },
  },
  setup() {
    const route = useRoute() as any;
    console.log(route);
    let data = reactive({
    });
    const { ctx } = getCurrentInstance() as any;
    const method = {
      async login():Promise<any> {
        let res = (await sendvalidator({email:ctx.email})) as any;
        if (res.status != 200)
        return Notify.create({
          message: res.message,
          color: 'negative',
          icon: 'report_problem',
          position: 'top',
          timeout: 2000,
        });
        console.log(ctx);
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