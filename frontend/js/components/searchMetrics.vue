<template>
  <table class="table is-fullwidth">
    <thead>
      <tr>
        <th></th>
        <th v-for="choice in unNormalizedChoices">
          {{ choice }}
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(counts, key) in matchCounts">
        <th>マッチ数({{ key }})</th>
        <td v-for="count in counts">{{ count }}</td>
      </tr>
      <tr>
        <th>選択肢(非正規化)</th>
        <td v-for="unNormalizedChoice in unNormalizedChoices">{{ unNormalizedChoice }}</td>
      </tr>
      <tr>
        <th>選択肢(正規化後)</th>
        <td v-for="normalizedChoice in normalizedChoices">{{ normalizedChoice }}</td>
      </tr>
      <tr>
        <th>選択肢(かな)</th>
        <td v-for="choiceSurface in choiceSurfaces">{{ choiceSurface }}</td>
      </tr>
      <tr>
        <th>選択肢(候補)</th>
        <td v-for="choiceCandidates in allChoiceCandidates">
          <div class="content">
            <ul>
              <li v-for="candidate in choiceCandidates">{{candidate}}</li>
            </ul>
          </div>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  props: {
    answer: {
      type: Object,
      required: true,
    },
    choices: {
      type: Array,
      required: true,
    },
    matchCounts: {
      type: Object,
      required: true,
    },
  },
  computed: {
    unNormalizedChoices() {
      return this.choices.map(choice => choice.unnormalized);
    },
    normalizedChoices() {
      return this.choices.map(choice => choice.normalized);
    },
    choiceSurfaces() {
      return this.choices.map(choice => choice.surface);
    },
    allChoiceCandidates() {
      return this.choices.map(choice => choice.candidates);
    },
  },
}
</script>

<style lang="scss" scoped>
</style>
