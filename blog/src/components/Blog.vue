<template>
    <div class="container">
         <section v-for="post in posts" v-bind:key="post">
            <h2>{{ post.title }}</h2>
            <hr>
        </section>
    </div>
</template>

<script>
export default {
    name: 'blog',
    data() {
        return{
            posts: [],
        }
    },
    mounted(){
      this.getPosts()
    },
    methods: {
        async getPosts(){
        try {
            const response = await fetch('http://localhost:8888/posts')
            const data = await response.json()
            this.posts = data
        } catch (error){
            console.error(error)
      }
    } 
  },
}
</script>