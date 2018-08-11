<template>
  <div class=" col-sm-10 col-sm-offset-1">
    <h1>Dynamic Tree View!</h1>
    <button class="btn btn-warning" v-on:click="getList()">Get a Tree List</button>
    <button class="btn btn-warning" v-on:click="getDetails(1)">Get a Detail</button>
    <ul class="view-area col-xs-12" v-if="posts && posts.hits">
      <li class=" tree-list" v-for="post in posts.hits">
        <a v-on:click="getDetails(post.objectID)" v-bind:title="post.story_title" >
          [{{post.points}}]
          <span v-if="post.num_comments">
            <i class="glyphicon glyphicon-plus" ></i>
            {{post.num_comments}}
          </span>
          <span>{{post.story_title}}</span>
          <span class="badge badge-default">
            <span class="label" >{{post.author}}</span>
          </span>
        </a>
      </li>
    </ul>

    <ul class="view-area col-xs-12" v-if="comments">
      <li class=" comment-list" v-for="comment in comments.children">
        <button v-on:click="toggleExpandedState(comment)" v-bind:title="comment.comment_text" >
          <i class="glyphicon glyphicon-cog" v-if="comment.children.length===0"></i>
          <i class="glyphicon glyphicon-plus" v-if="comment.children.length"></i>
          [{{comment.points}}]
          <span v-if="comment.num_comments">
            <i class="glyphicon glyphicon-plus" ></i>
            {{comment.num_comments}}
          </span>
          <span>{{comment.text}}</span>
          <span class="badge badge-success">
            <span class="label" >{{comment.author}} - {{comment.children.length}}</span>
          </span>
        </button>
      </li>
    </ul>
    <pre v-bind:html="comments.children">{{comments}}</pre>
  </div>
</template>

<script>
import auth from '../auth'

export default {
  computed: {
    isActive () {
      return this.href === this.$root.currentRoute
    }
  },
  data () {
    return {
      view: '',
      posts: {},
      comments: {}
    }
  },
  methods: {
    go (event) {
      event.preventDefault()
      this.$root.currentRoute = this.href
      window.history.pushState(
        null,
        window.location.pathname,
        this.href
      )
    },
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
    toggleExpandedState (comment) {
      // toggleExpandedState
      this.is
    },
    getList () {
      this.resetModels()
      this.$http.get('https://hn.algolia.com/api/v1/search?query=...&tags=front_page', null)
      .then(response => {
        this.$root.currentState = this.href + '/story'
        this.posts = response.body
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
    resetModels () {
      this.posts = {}
      this.comments = {}
    },
    getDetails (id) {
      this.$http.get('https://hn.algolia.com/api/v1/items/' + id, {hitsPerPage: 100})
      .then(response => {
        this.resetModels()
        this.comments = response.body
        this.$root.currentState = '/view/:id'
        window.console.warn(this.posts)
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
    display: inline-block;
  }
  .tree-list:hover {
    cursor: pointer;
    text-decoration: none;
  }
  .tree-node {
    margin-left: 1em;
  }
</style>
