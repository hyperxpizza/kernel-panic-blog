<template>
  <div>
    <Navigation />
    <Blog />
  </div>
</template>

<script>
import Navigation from './components/Navigation.vue';
import Blog from './components/Blog.vue'

export default {
  name: 'app',
  components: {
    Navigation,
    Blog,
  },
}
</script>