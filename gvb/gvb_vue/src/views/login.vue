<template>
  <div class="body">
    <div class="login-panel">
      <n-card title="后台登录" class="login-box">
        <n-form :rules="rules" :model="admin">
          <n-form-item path="username" label="用户名">
            <n-input v-model:value="admin.username" placeholder="请输入用户名" autocomplete="username"></n-input>
          </n-form-item>

          <n-form-item path="password" label="密码">
            <n-input v-model:value="admin.password" placeholder="请输入密码" autocomplete="current-password" type="password"></n-input>
          </n-form-item>
        </n-form>
        <template #footer>
          <n-button type="primary" @click="login">登录</n-button>
          <div></div>
          <n-checkbox v-model:checked="admin.remember" label="记住我" style="display: flex;justify-content: flex-end;"/>
        </template>
      </n-card>
    </div>
  </div>
</template>

<style scoped>
html, body, .app {
  height: 100%;
  margin: 0;
  padding: 0;
}

.body {
  background-color: #f0f0f0;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-panel {
  width: 400px;
  margin: 0 auto;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.login-box {
  background-color: #ffffff;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.n-input {
  border: 1px solid #ccc;
  border-radius: 4px;
  padding: 8px;
}

.n-button {
  background-color: #007bff;
  color: #fff;
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.n-checkbox {
  color: #333;
  font-size: 14px;
}

</style>

<script setup>
import { reactive ,inject} from 'vue'
import {AdminStore} from "../stores/AdminStore";
import {useRouter} from "vue-router";
import as from "../plugins/axios"
// const route = useRoute()
const  router= useRouter()
const  adminstore = AdminStore()
//const axios = inject("axios")
const  msg = inject("message")



let rules = {
    username: [
        { required: true, message: "请输入用户名", trigger: "blur" },
        {min:4,max:12,message:"长度限制4-12位",trigger:"blur"},
    ],
    password: [
        { required: true, message: "请输入密码", trigger: "blur" },
        {min:6,max:20,message:"长度限制6-20位",trigger:"blur"},
    ],
}

const admin = reactive({
    username: localStorage.getItem("username")||"",
    password: localStorage.getItem("password")||"",
    remember: localStorage.getItem("remember")==1,
})

const login = async () => {
    let result = await as.post("/login", {
        username: admin.username,
        password: admin.password,
    });
    console.log(result);
    if(admin.username==''||admin.password==''){
      msg.info("请输入登录信息")
    }else {
      if (result.data.code==200){
        // adminstore.atoken = result.data.atoken
        // adminstore.rtoken = result.data.rtoken
        adminstore.username = admin.username
        // localStorage.setItem("atoken",result.data.atoken)
        // localStorage.setItem("rtoken",result.data.rtoken)
        adminstore.remember = admin.remember
        //console.log(adminstore.username,admin.password,adminstore.token,admin.remember)
        if (admin.remember){
          localStorage.setItem("username",adminstore.username)
          localStorage.setItem("password",admin.password)
          localStorage.setItem("remember",admin.remember?1:0)
        }else {
          localStorage.setItem("remember", 0)
          localStorage.setItem("username","")
          localStorage.setItem("password","")
        }
        router.push("admin")
        msg.info(result.data.msg)
      }else {
        msg.error(result.data.msg)
      }
    }

}


</script>