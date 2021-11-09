<template>
  <q-page padding>
    <div class="row justify-center">
          <q-card class="col-md-4 col-sm-12">
            <q-form
              @submit="onSubmit"
              @reset="onReset"
              class="q-gutter-md"
            >
              <q-input
                filled
                v-model="name"
                label="Your name *"
                hint="Name and surname"
                lazy-rules
                :rules="[ val => val && val.length > 0 || 'Please type something']"
              />

              <q-input
                filled
                type="number"
                v-model="age"
                label="Your age *"
                lazy-rules
                :rules="[
                  val => val !== null && val !== '' || 'Please type your age',
                  val => val > 0 && val < 100 || 'Please type a real age'
                ]"
              />

              <q-toggle v-model="accept" label="I accept the license and terms" />

              <div>
                <q-btn label="Submit" type="submit" color="primary"/>
                <q-btn label="Reset" type="reset" color="primary" flat class="q-ml-sm" />
              </div>
            </q-form>
          </q-card>
      </div>
  </q-page>
</template>

<script>
export default {
  // name: 'PageName',
}
</script>
