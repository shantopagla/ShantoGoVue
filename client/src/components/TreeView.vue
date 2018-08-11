<template>
  <div class=" row-fluid">
    <div class=" col-md-4" >
      <div class="list-group sidebar" >
        <span v-for="post in posts.hits" class="list-group-item" v-bind:id="post.objectID" >
          <router-link :to="{ name: 'view.detail', params: { id: post.objectID }}" >
            <h6 class="list-group-item-heading alert-default" v-bind:title="post.title" v-on:click="getDetails(post.objectID)">{{post.title}}
              <p v-if="post.text" v-bind:html="post.text" class="list-group-item-text"></p>
              [{{post.points}}]
            </h6>
            
              <span class="label alert-info"> 
                <span class="" >{{post.author}} </span>
              </span>
              [
                {{post.num_comments + ' comments'}}
              ] <a v-bind:href="post.url">
                  <i class="glyphicon glyphicon-share" v-if="post.url" ></i>
                </a>
          </router-link>
        </span>
      </div>
    </div>
    <section class="panel panel-info col-md-8" v-if="currentStory">
        <span class="pull-right badge alert-default">{{Date(currentStory.created_at)}}</span>
        <h4 class=" panel-heading">{{currentStory.title}} (<em> {{currentStory.children.length}} comments </em>)
        </h4>
        <div class="panel-body" >
          <div class="container" v-tree-node="currentStory.id" v-model="currentStory.children">
          </div>
        </div>
    </section>
  </div>
</template>

<script>
import blue from '../tree'
import auth from '../auth'

export default {
  computed: {
    isActive () {
      return this.href === this.$root.currentRoute
    }
  },
  data () {
    return {
      posts: {},
      expanded: {},
      store: {},
      currentStory: null
    }
  },
  created () {
    // fetch the data when the view is created and the data is
    // already being observed
    this.getView()
  },
  watch: {
    // call again the method if the route changes
    // '$route': 'getDetail'
  },
  mounted () {
    return this.getview
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
    isExpandedPost (id) {
      (!this.store[id] ? this.store[id] = {isExpanded: false} : id)
      return this.store[id].isExpanded === true
    },
    getView () {
      this.posts = {}
      const url = blue.getTopResultsUrl()
      var s = JSON.stringify
      this.$http.get(url, {hitsPerPage: 50})
      .then(response => {
        this.posts = response.body
        localStorage.setItem('topstories', s(response.body.hits))
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
      // this.posts = {}
      // this.comments = {}
    },
    getDetails (id) {
      if (this.expanded[id] === true) {
        this.expanded[id] = false
        return
      }
      this.currentStory = null
      this.$http.get('https://hn.algolia.com/api/v1/items/' + id, {hitsPerPage: 100})
      .then(response => {
        this.resetModels()
        this.store[id] = response.body
        this.expanded[id] = true
        this.$root.currentState = '/view/:id'
        this.currentStory = response.body
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
  .sidebar {
    height: 88vh;
    margin: 0em 1em 0em -1em;
    overflow: none;
    overflow-y: scroll; 
  }
  .tree-list {
    list-style-type: none;
    display: inline-block;
  }
  a:hover {
    cursor: pointer;
    text-decoration: none;
  }
  .tree-node {
    margin-left: 1em;
  }
</style>
