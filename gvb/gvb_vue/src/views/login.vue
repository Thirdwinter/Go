<template>
    <div class="login-panel">
        <n-card title="后台登录">
            <n-form :rules="rules" :model="admin">
                <n-form-item path="account" label="用户名">
                    <n-input v-model:value="admin.account" placeholder="请输入用户名"></n-input>
                </n-form-item>

                <n-form-item path="password" label="密码">
                    <n-input v-model:value="admin.password" placeholder="请输入密码"></n-input>
                </n-form-item>
            </n-form>
            <template #footer>
                <n-button @click="login">登录</n-button>
                <n-checkbox v-model:checked="admin.rember" label="记住我"/>
            </template>
        </n-card>
    </div>
</template>

<style scoped>
.login-panel{
    width: 500px;
    margin: 0 auto;
    margin-top: 130px;
}
</style>

<script setup>

import { reactive ,inject} from 'vue'

const axios = inject("axios")

let rules = {
    account: [
        { required: true, message: "请输入用户名", trigger: "blur" },
        {min:4,max:12,message:"长度限制4-12位",trigger:"blur"},
    ],
    password: [
        { required: true, message: "请输入密码", trigger: "blur" },
        {min:6,max:20,message:"长度限制6-20位",trigger:"blur"},
    ],
}

const admin = reactive({
    account: '',
    password: '',
    rember:false,
})

const login = async () => {
    let result = await axios.post("/login", {
        username: admin.account,
        password: admin.password,
    });
    console.log(result)
}
</script>