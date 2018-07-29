<template>
  <div>
    <section class="hero is-primary">
      <div class="container">
        <div class="hero-body">
          <h2 class="title">sharqen</h2>
          <h3 class="subtitle">クイズ検索ワード作成ツール</h3>
        </div>
      </div>
    </section>
    <section class="section">
      <div class="container">
        <div class="columns">

          <div class="column">
            <!-- クイズ表示 -->
            <article class="message is-primary">
              <div class="message-header">
                <p><span class="icon"><i class="fas fa-question-circle"></i></span>{{ quiz.num }}.{{ quiz.question }} ({{ quiz.channel_name }})</p>
              </div>
              <div class="message-body">
                <div class="content">
                  <ol>
                    <li v-for="choice in quiz.choices">
                      {{ choice }}
                    </li>
                  </ol>
                  答え:（{{ quiz.correct }}）<strong>{{ correctChoice }}</strong>
                </div>
              </div>
            </article>

            <!-- 検索ワードフォーム -->
            <div>
              <div class="field">
                <label class="label">検索ワード</label>
                <div class="control is-expanded">
                  <textarea
                    class="textarea"
                    type="text"
                    placeholder="新規検索ワード"
                    autofocus
                    v-model="query"
                    ></textarea>
                </div>
              </div>
              <div class="field is-grouped is-grouped-right">
                <div class="control">
                  <button
                    class="button is-danger"
                    :disabled="!query"
                    @click="deleteQuestion"
                  ><span class="icon"><i class="fas fa-trash-alt"></i></span>&nbsp;&nbsp;削除</button>
                </div>
                <div class="control">
                  <button
                    class="button is-link"
                    :disabled="!query"
                    @click="search"
                  ><span class="icon"><i class="fas fa-search"></i></span>&nbsp;&nbsp;検索</button>
                </div>
                <div class="control">
                  <button
                    class="button is-success"
                    type="submit"
                    :disabled="!query"
                    @click="saveQuery"
                    ><span class="icon"><i class="fas fa-save"></i></span>&nbsp;&nbsp;登録</button>
                </div>
              </div>
            </div>

          </div>

          <div class="column">
            <div v-if="existAnswer">
              <!-- 検索結果(メトリクス) -->
              <div>
                回答:（{{ answer.index + 1 }}）<strong>{{ answer.choice }}</strong>
                <span v-if="isCorrectAnswer" class="tag is-success">正答</span>
                <span v-else class="tag is-danger">誤答</span>
              </div>
              <div>
                <search-metrics
                  :answer="answer"
                  :choices="choices"
                  :matchCounts="matchCounts"
                ></search-metrics>
              </div>
              <!-- 検索結果(本文) -->
              <div>
                <search-result
                  v-for="(searchResult, index) in searchResults"
                  :key="index"
                  :text="searchResult"
                ></search-result>
              </div>
            </div>
            <div v-else>
              <div class="has-text-centered">
                <h1 class="title has-text-grey">
                  <span class="icon"><i class="fas fa-magic"></i></span>&nbsp;&nbsp;empty!
                </h1>
              </div>
            </div>
          </div>

        </div>
      </div>
    </section>
  </div>
</template>

<script>
import axios from 'axios';
import * as request from './../request';
import searchResult from './searchResult.vue';
import searchMetrics from './searchMetrics.vue';


export default {
  components: {
    searchMetrics,
    searchResult,
  },
  created() {
    this.fetchQuiz();
  },
  data() {
    return {
      answer: {},
      choices: [],
      matchCounts: {},
      query: null,
      quiz: {
        id: null,
        channel_name: null,
        num: null,
        question: null,
        choices: [],
        correct: null,  // index + 1 で持つ
      },
      searchResults: [],
    };
  },
  computed: {
    existAnswer() {
      return Object.keys(this.answer).length !== 0;
    },
    correctChoice() {
      return this.quiz.choices[this.quiz.correct - 1] || '';
    },
    isCorrectAnswer() {
      if (this.answer) {
        return this.quiz.correct -1 === this.answer.index;
      }
      return false;
    },
  },
  methods: {
    fetchQuiz() {
      request.get('/quiz/next').then(res => {
        this.$set(this, 'quiz', res.data.quiz);
        this.$set(this, 'query', res.data.query);
      }).catch(error => {
        console.error(error);
      });
    },
    deleteQuestion() {
      request.del(`/quiz/${this.quiz.id}`).then(res => {
        this.fetchQuiz();
      }).catch(error => {
        console.error(error);
      });
    },
    search() {
      request.get('/search-result', {
        question: this.query,
        choices: this.quiz.choices,
      }).then(res => {
        this.$set(this, 'answer', res.data.answer);
        this.$set(this, 'choices', res.data.choices);
        this.$set(this, 'matchCounts', res.data.match_counts);
        this.$set(this, 'searchResults', res.data.search_results);
      }).catch(error => {
        console.error(error);
      });
    },
    saveQuery() {
      // NOTE: 本当はCSRF対策が必要だがローカル利用を前提としているため入れない
      request.post('/search-query', {
        quiz_id: this.quiz.id,
        search_query: this.query,
      }).then(res => {
        // 次のクイズを取得する
        this.fetchQuiz();
      }).catch(error => {
        console.error(error);
      });
    },
  },
}
</script>

<style lang="scss" scoped>
table {
  font-size: .9em;
}
</style>
