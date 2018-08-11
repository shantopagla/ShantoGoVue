const API_URL = 'https://hn.algolia.com/api/v1/'
const TOP_LIST_URL = API_URL + 'search?query=...&tags=front_page&hitsPerPage=50'
const SIGNUP_URL = API_URL + 'signup'

import Vue from 'vue'
// import TreeView from '@/components/TreeView'

Vue.directive('tree-node', {
  mounted: function () {
    this.$nextTick(function () {
      // Code that will run only after the
      // entire view has been rendered
      window.console.warn('what loads...')
    })
  },
  beforeRouteUpdate (to, from, next) {
    // called when the route that renders this component has changed,
    // but this component is reused in the new route.
    // For example, for a route with dynamic params `/foo/:id`, when we
    // navigate between `/foo/1` and `/foo/2`, the same `Foo` component instance
    // will be reused, and this hook will be called when that happens.
    // has access to `this` component instance.
    next()
  },
  bind: function (el, binding, vnode) {
    // var s = JSON.stringify
    const data = vnode.data.model.value
    window.console.warn(data)
    el.innerHTML = '<ul>'

    data.forEach(function (comment) {
      el.innerHTML += '<li class="row tree-list">' +
        (comment.children.length ? '[+ ' + comment.children.length + ']' : '[-]') +
      comment.text +
      ' <span class="badge pull-right alert-default"> ' +
      comment.author + ' </span> ' +
      ' </li> '
    })
    el.innerHTML += '</ul>'

    /*
                <i class="glyphicon glyphicon-plus" v-if="!isExpandedPost(post.objectID)" ></i>
                <i class="glyphicon glyphicon-minus" v-if="isExpandedPost(post.objectID)" ></i>
      window.console.warn()
      'value: ' + s(binding.value) + '<br>' +
      'expression: ' + s(binding.expression) + '<br>' +
      'argument: ' + s(binding.arg) + '<br>' +
      'modifiers: ' + s(binding.modifiers) + '<br>' +
      'vnode keys: ' + Object.keys(vnode).join(', ')
    */
  }
})

/*
*/

export default {

  /*
  getNewTreeNode () {
    return treeNode
  },
  */

  getTopResultsUrl () {
    return TOP_LIST_URL
    /*
    context.$http.get(TOP_LIST_URL, params).then(response => {
      localStorage.setItem('id_token', response.body.id_token)
      if (redirect) {
        context.$router.replace(redirect)
      }
    }, response => {
      context.error = response.statusText
    })
    */
  },

  signup (context, creds, redirect) {
    context.$http.post(SIGNUP_URL, creds).then(response => {
      localStorage.setItem('id_token', response.body.id_token)
      if (redirect) {
        context.$router.replace(redirect)
      }
    }, response => {
      context.error = response.statusText
    })
  },

  logout (context) {
    localStorage.removeItem('id_token')
    context.$router.replace('/home')
  },

  isAuthenticated () {
    var jwt = localStorage.getItem('id_token')
    if (jwt) {
      return true
    }
    return false
  },

  getAuthor () {
    return {
      'Authorization': 'Bearer ' + localStorage.getItem('id_token')
    }
  }
}
