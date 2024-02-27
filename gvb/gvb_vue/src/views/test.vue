<template>
  <div class="menus">
    <n-anchor
        affix
        listen-to=".document-scroll-container"
        :trigger-top="24"
        :top="88"
        style="z-index: 1"
        :bound="24"
    >
      <div style="height: 100px">
        <n-anchor-link :title="'你好!  '+myname" style="margin-bottom: 55px" />
      </div>
      <n-anchor-link title="用户管理" style="margin-bottom: 25px" href="admin/test"/>
      <n-anchor-link title="分类管理" style="margin-bottom: 25px" />
      <n-anchor-link title="文章管理" style="margin-bottom: 25px" />
      <n-anchor-link title="退出" style="margin-bottom: 25px" @click="liuleliule"/>
    </n-anchor>
  </div>
  <div style="padding: 10px;width: 100%;background-color: black">
    <router-view>

    </router-view>>
  </div>
</template>
<style>
.menus{
  height: 100vh;
  width: 150px;
  background-color: bisque;
  div{
    cursor:pointer;
    &:hover{
      color: #42b983;
    }
  }
}
</style>


<script setup>
import {useRouter} from "vue-router";
const router = useRouter()
const myname = localStorage.getItem("username");
const liuleliule=()=>{
  localStorage.removeItem('token');
  router.push("/login")

}
</script>