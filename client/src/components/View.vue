<template>
  <div class="col-sm-10 col-sm-offset-1">
    <h1>Dynamic Tree View!</h1>
    <button class="btn btn-warning" v-on:click="getList()">Get a List</button>
    <button class="btn btn-warning" v-on:click="getDetails(1)">Get a Detail</button>
    <ul class="view-area col-xs-12" v-if="hacks.hits">
      <li class=" tree-list" v-for="hack in hacks">
        <button v-on:click="getDetails(hack.id)">
          <i class="glyphicon glyphicon-cog" v-if="hack.children.length===0"></i>
          <i class="glyphicon glyphicon-plus" v-if="hack.children.length"></i>
          <i class="glyphicon glyphicon-minus" v-if="!hack.children.length"></i>
          <span class="badge badge-right" v-if="hack.children"></span>
          <span class="tree-node" >{{ hack.author + ' ' +hack.title}}</span>
        </button>
      </li>
    </ul>
  </div>
</template>

<script>
import auth from '../auth'

export default {
  data () {
    return {
      view: '',
      hacks: {},
      model: {
        'author': 'pg',
        'children': [],
        'created_at': '2006-10-09T18:21:51.000Z',
        'created_at_i': 1160418111,
        'id': 1,
        'options': [],
        'parent_id': null,
        'points': 61,
        'story_id': null,
        text: null,
        'title': 'Y Combinator',
        'type': 'story',
        'url': 'http://ycombinator.com'
      }
    }
  },
  methods: {
    getView () {
      this.$http.get('/api/view/protected/random', {headers: auth.getAuthHeader()})
      .then(response => {
        this.view = response.body
      }, response => {
        if (response.status === 401) {
          auth.logout(this)
        }
        console.log(response)
      })
    },
    getList () {
      // this.$http.get('https://hn.algolia.com/api/v1/search?query=micro&tags=story', null)
      this.$http.get('https://hn.algolia.com/api/v1/items/' + 1, null)
      .then(response => {
        this.hacks = response.body
      }, response => {
        if (response.status === 401) {
          auth.logout(this)
        }
        console.log(response)
      })
    },
    isArray (h) {
      return (h && h.length && Array.isArray(h))
    },
    getDetails (id) {
      this.$http.get('https://hn.algolia.com/api/v1/items/' + id, null)
      .then(response => {
        this.hacks = response.body
      }, response => {
        if (response.status === 401) {
          auth.logout(this)
        }
        console.log(response)
      })
    }
  }
}
</script>

<style>
  .tree-list {
    list-style-type: none;
  }
  .tree-node {
    margin-left: 1em;
  }
</style>
