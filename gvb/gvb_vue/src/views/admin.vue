<template>
  <div class="main-container">
    <div class="main-panel">
      <div class="menus">
        <div v-for="(menu, index) in menus" :key="index" @click="topage(menu)">
          {{ menu.name }}
        </div>
      </div>
    </div>
    <div class="content">
      <router-view></router-view>
    </div>
    <div class="title">后台管理</div>
  </div>
</template>

<style>
.main-container {
  display: flex;
  position: relative;
}

.main-panel {
  padding: 20px 0;
  box-sizing: border-box;
  line-height: 55px;
  text-align: center;
  width: 180px;
  height: 95vh;
  border-right: 1px solid #7842b9;
}

.menus div {
  cursor: pointer;
}

.menus div:hover {
  color: chartreuse;
}

.content {
  flex: 1;
  padding: 20px;
}

.title {
  font-size: 65px;
  font-weight: bold;
  text-align: right;
  position: fixed;
  color: rgba(0, 0, 0, 0.47);
  bottom: 20px;
  right: calc((100vw - 1000px) / 2);
}
</style>

<script setup>
import {useRouter} from "vue-router";
const router = useRouter()

let menus = [
  {name:"用户管理",href:"/admin/user"},
  {name:"分类管理",href: "/admin/category"},
  {name:"文章管理",href: "/admin/article"},
  {name:"登出",href: "logout"}
];
const topage=(menu)=>{
  if(menu.href=="logout"){
    localStorage.removeItem("token")
    router.push("/login")
  }else {
    router.push(menu.href)
  }
}

</script>