<template>
  <div :class="$style.main">
    <ui-button
      @click="
        $store.dispatch('modal/show', {
          name: 'add/word',
          data: { name: '', category: '', translate: {} },
          func: () => {
            $store.dispatch('word/add');
          },
        })
      "
      text="Add"
      icon="plus"
      style="flex: 0"
    />

    <div :class="$style.list">
      <div :class="$style.word" v-for="x in $store.state.word.list" :key="x.name">
        <div>{{ x.name }} - {{ x.transcription }} - {{ x.translate }} -</div>
        <ui-button :class="$style.button" @click="editWord(x)" icon="pencil" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  components: {},
  async mounted() {
    await this.$store.dispatch('word/getList');
  },
  methods: {
    editWord(x: any) {
      this.$store.dispatch('modal/show', {
        name: 'edit/word',
        data: {
          name: x.name,
          transcription: x.transcription,
          category: x.category.join(', '),
          translate: {
            noun: x.translate.noun.join(', '),
            verb: x.translate.verb.join(', '),
            adjective: x.translate.adjective.join(', '),
            adverb: x.translate.adverb.join(', '),
            particle: x.translate.particle.join(', '),
            preposition: x.translate.preposition.join(', '),
          },
          isExists: false,
        },
        func: () => {
          this.$store.dispatch('word/edit');
        },
      });
    },
  },
  data: () => {
    return {};
  },
});
</script>

<style lang="scss" module>
.main {
  .list {
    .word {
      display: flex;
      align-items: center;
      padding: 5px;

      .button {
        flex: none;
        margin-left: auto;
        padding: 1px;
      }
    }
  }
}
</style>
