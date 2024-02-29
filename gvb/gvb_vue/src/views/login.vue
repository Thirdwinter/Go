<template>
  <div class="body">
    <div class="login-panel">
        <n-card title="后台登录" class="login-box">
            <n-form :rules="rules" :model="admin">
                <n-form-item path="username" label="用户名">
                    <n-input v-model:value="admin.username" placeholder="请输入用户名" autocomplete="username"></n-input>
                </n-form-item>

                <n-form-item path="password" label="密码">
                    <n-input v-model:value="admin.password" placeholder="请输入密码" autocomplete="current-password"></n-input>
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
  background-color: bisque;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-panel {
  width: 450px;
  margin: 0 auto;
  height: 100%;
  position: absolute;


}

.login-box {
  background-color: aqua;
  top: 60%;
  left: 60%;
  transform: translate(-0%,-60%);
  border-radius: 9px;
}
</style>

<script setup>
import { reactive ,inject} from 'vue'
import {AdminStore} from "../stores/AdminStore";
import {useRouter} from "vue-router";

// const route = useRoute()
const  router= useRouter()
const  adminstore = AdminStore()
const axios = inject("axios")
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
    let result = await axios.post("/login", {
        username: admin.username,
        password: admin.password,
    });
    console.log(result);
    if(admin.username==''||admin.password==''){
      msg.info("请输入登录信息")
    }else {
      if (result.data.status==200){
        adminstore.token = result.data.token
        adminstore.username = admin.username
        localStorage.setItem("token",result.data.token)
        adminstore.remember = admin.remember
        msg.info(result.data.msg)
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
        router.push("/admin")
      }else {
        msg.error(result.data.msg)
      }
    }

}


</script>