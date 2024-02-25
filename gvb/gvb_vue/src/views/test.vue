<!-- <template>
  <div class="container">
    <a-form
    :model="formdata"
    name="formdata"
    :label-col="{ span: 8 }"
    :wrapper-col="{ span: 16 }"
    :rules="rules"
    autocomplete=align
    class="loginBox"
    style="display: flex; flex-direction: column; justify-content: center; align-items: center;"
  >
    <a-form-item name="username">
      <a-input 
      :model:value="formdata.username" 
      placeholder="Username"
      style="width: 200%;"      
      >
        <template #prefix><UserOutlined style="color: rgba(0, 0, 0, 0.25)" /></template>
      </a-input>
    </a-form-item>

    <a-form-item name="password">
      <a-input 
      :model:value="formdata.password" 
      type="password" 
      placeholder="Password"
      style="width: 200%;"
      >
        <template #prefix><LockOutlined style="color: rgba(0, 0, 0, 0.25)" /></template>
      </a-input>
    </a-form-item>

    <a-form-item>
      <a-button
        type="primary"
        class="loginBtn"
      >
        Log in
      </a-button>
    </a-form-item>
  </a-form>
  </div>
</template>



<script>
export default({
  name: 'LoginView',
  data() {
    return {
      formdata: {
        username:'',
        password:''
      },
      
      rules: {
        username: [
          { required: true, message: "请输入用户名", trigger: 'blur' },
          { min:4,max:12,message:"限制4-12位",trigger:'blur' },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: 'blur' },
          { min:6,max:20,message:"限制6-20位",trigger:'blur' },
        ],
      },
    }
  }
});
</script>
<style scoped>
.container{
  height: 100%;
  background-color: aqua;
}
.loginBox{
  width: 35%;
  height: 45%;
  background-color:antiquewhite;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%,-50%);
  border-radius: 9px;

}
.loginForm{
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0,20px;
  box-sizing: border-box;
} 
.loginBtn{
  display: flex;
  justify-content: flex-end;

}
</style> -->
