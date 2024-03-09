<script setup>
import {useRouter}  from "vue-router";

const router = useRouter()

const  toLogin=()=>{

  router.push("/admin")
}
</script>


<template>
<h1>制作中...</h1>
  <n-button type="primary" @click="toLogin">前往后台</n-button>
</template>

<style>

</style>