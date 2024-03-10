<template>
    <div>
      <h1>User Authentication Test</h1>
  
      <div id="loginForm">
        <h2>Login</h2>
        <label for="username">Username:</label>
        <input type="text" id="username" name="username" v-model="username" required>
        <label for="password">Password:</label>
        <input type="password" id="password" name="password" v-model="password" required>
        <button @click="login">Login</button>
      </div>
  
      <div id="searchForm">
        <h2>Search User</h2>
        <label for="searchUsername">Search Username:</label>
        <input type="text" id="searchUsername" name="searchUsername" v-model="searchUsername" required>
        <button @click="searchUser">Search</button>
      </div>
  
      <div id="result">{{ searchResult }}</div>
    </div>
  </template>
  
  <script>
  //import as from '../plugins/axios';

  import axios from 'axios';
  //创建axios实例
  const service = axios.create({
    baseURL: '/Api',
    timeout: 10000
  });
  
  export default {
    data() {
      return {
        username: '',
        password: '',
        searchUsername: '',
        searchResult: ''
      };
    },
    methods: {
      login() {
        service.post('/login', { username: this.username, password: this.password })
          .then(response => {
            console.log(response);
            // Store the generated token in a cookie
          })
          .catch(error => {
            console.error(error);
          });
      },
      searchUser() {
        service.get('/searchuser', {
          withCredentials: true,
          params: {
            username: this.searchUsername
          }
        })
          .then(response => {
            console.log(response.data);
            this.searchResult = "Search result: " + response.data;
          })
          .catch(error => {
            console.error(error);
          });
      }
    }
  };
  </script>
  
  <style>
  /* Add your styles here */
  </style>
  