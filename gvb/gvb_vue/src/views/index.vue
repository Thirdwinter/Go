<script setup>
import {useRouter}  from "vue-router";

const router = useRouter()

const  toLogin=()=>{
  router.push("/login")
}
</script>


<template>
<h1>监修中...</h1>
  <n-button type="primary" @click="toLogin">前往登录</n-button>
</template>

<style>

</style>