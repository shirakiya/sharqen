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

            <article class="error-message" v-if="existError">
              <div class="notification is-danger">
                <button class="delete" @click="closeError"></button>
                <p v-if="error.statusText && error.status">
                  {{ this.error.statusText }} [{{ this.error.status }}]
                </p>
                <p>
                  {{ this.error.message }}
                </p>
              </div>
            </article>

          </div>

          <div class="column">
            <div v-if="existAnswer">
              <!-- 検索結果(メトリクス) -->
              <div>
                回答:（{{ answerIndex + 1 }}）<strong>{{ answer }}</strong>
                <span v-if="isCorrectAnswer" class="tag is-success">正答</span>
                <span v-else class="tag is-danger">誤答</span>
              </div>
              <div>
                <search-metrics
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
      answer: null,
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
      error: {},
    };
  },
  computed: {
    existError() {
      return Object.keys(this.error).length !== 0;
    },
    existAnswer() {
      return Boolean(this.answer);
    },
    correctChoice() {
      return this.quiz.choices[this.quiz.correct - 1] || '';
    },
    answerIndex() {
      return (this.answer) ? this.quiz.choices.indexOf(this.answer) : null;
    },
    isCorrectAnswer() {
      if (this.existAnswer) {
        return this.quiz.correct === this.answerIndex + 1;
      }
      return false;
    },
  },
  methods: {
    showError(error) {
      if (error.response) {
        this.$set(this, 'error', {
          message: error.message,
          status: error.response.status,
          statusText: error.response.statusText,
        });
        console.log(error.response);
      } else {
        this.$set(this, 'error', {
          message: error.message,
        });
      }
      console.log(error.request);
    },
    closeError() {
      if (this.error) {
        this.error = {};
      }
    },
    clearSearchResult() {
      this.answer = null;
      this.choices = [];
      this.matchCounts = {};
      this.searchResults = [];
    },
    fetchQuiz() {
      request.get('/quiz/next').then(res => {
        this.$set(this, 'quiz', res.data.quiz);
        this.$set(this, 'query', res.data.query);
      }).catch(error => {
        this.showError(error);
      });
    },
    deleteQuestion() {
      request.del(`/quiz/${this.quiz.id}`).then(res => {
        this.fetchQuiz();
      }).catch(error => {
        this.showError(error);
      });
    },
    search() {
      request.get('/search-result', {
        question: this.query,
        choices: this.quiz.choices,
      }).then(res => {
        console.log('search: res.data', res.data);
        this.$set(this, 'answer', res.data.answer);
        this.$set(this, 'choices', res.data.choices);
        this.$set(this, 'matchCounts', res.data.match_counts);
        this.$set(this, 'searchResults', res.data.search_results);
      }).catch(error => {
        this.showError(error);
      });
    },
    saveQuery() {
      // NOTE: 本当はCSRF対策が必要だがローカル利用を前提としているため入れない
      request.post('/search-query', {
        quiz_id: this.quiz.id,
        search_query: this.query,
      }).then(res => {
        // 次のクイズを取得する
        this.clearSearchResult();
        this.fetchQuiz();
      }).catch(error => {
        this.showError(error);
      });
    },
  },
}
</script>

<style lang="scss" scoped>
table {
  font-size: .9em;
}

.error-message {
  margin-top: 1.5rem;
}
</style>
