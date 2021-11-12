<template>
  <div class="row justify-center">
    <q-page padding class="col-xs-12 col-sm-10 col-md-8">
      <q-list padding>
        <q-item clickable v-ripple class="q-mt-lg">
          <q-item-section>
            <q-item-label>Single line item</q-item-label>
            <q-item-label caption lines="2">Secondary line text. Lorem ipsum dolor sit amet, consectetur adipiscit elit.
              Secondary line text. Lorem ipsum dolor sit amet, consectetur adipiscit elit.
              Secondary line text. Lorem ipsum dolor sit amet, consectetur adipiscit elit.
              Secondary line text. Lorem ipsum dolor sit amet, consectetur adipiscit elit.
              Secondary line text. Lorem ipsum dolor sit amet, consectetur adipiscit elit.
            </q-item-label>
          </q-item-section>

          <q-item-section side top>
            <q-item-label caption>5 min ago</q-item-label>
            <q-icon name="star" color="yellow" />
            <q-chip clickable @click="onClick" color="primary" text-color="white" icon="event">
              Add to calendar
            </q-chip>
            <q-chip clickable @click="onClick" icon="bookmark">
              Bookmark
            </q-chip>
            <q-chip clickable @click="onClick" color="teal" text-color="white" icon="bookmark">
              Bookmark
            </q-chip>
            <q-chip clickable @click="onClick" color="red" text-color="white" icon="alarm" label="Set alarm" />
            <q-chip clickable @click="onClick" color="orange" text-color="white" icon="directions">
              Get directions
            </q-chip>
          </q-item-section>
        </q-item>

        <q-separator spaced inset />

        <q-item>
          <q-item-section>
            <q-item-label>Single line item</q-item-label>
            <q-item-label caption>Secondary line text. Lorem ipsum dolor sit amet, consectetur adipiscit elit.</q-item-label>
          </q-item-section>

          <q-item-section side top>
            <q-item-label caption>Voted!</q-item-label>
          </q-item-section>
        </q-item>

        <q-separator spaced inset />

      </q-list>
    </q-page>
  </div>
</template>

<style lang="scss" scoped>
.active-class:hover{
    box-shadow: 4px 4px 2px #888
}
</style>
