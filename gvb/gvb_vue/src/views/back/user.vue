<template>
  <div style="padding: 5px">
    <n-input placeholder="搜索" v-model:value="search_name">
      <template #prefix>
        <n-icon :component="FlashOutline" />
      </template>
    </n-input>
    <div style="display: flex;">
      <n-button type="info" @click="search" style="margin-right: 10px;">
        让我康康!
      </n-button>
      <n-button type="info" @click="letnil">
        清空
      </n-button>
    </div>

  </div>
  <div>
    <n-table :bordered="false" :single-line="false">
      <thead>
      <tr>
        <th>用户名</th>
        <th>角色码</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(user, index) in userlist" :key="index">
        <td>{{ user.username }}</td>
        <td>{{ user.role }}</td>
      </tr>
      </tbody>
    </n-table>
    <n-pagination
        v-model:page="page"
        v-model:page-size="pageSize"
        :page-count="pagecount"
        show-size-picker
        :page-sizes="pagesize"
        @update:page="handlePageUpdate"
        @update:page-size="handlePageSizeUpdate"
    />
  </div>
</template>

<script setup>
import { ref, onMounted, inject} from "vue";
import { FlashOutline } from "@vicons/ionicons5";
import {useRouter} from "vue-router";
//import as from "D:/Download/GOCODE/src/Go/gvb/gvb_vue/src/plugins/axios.js"
import as from "../../plugins/axios"

const router = useRouter()
//const axios = inject("axios");
let userlist = ref([]);
let pagecount = ref(0);
let page = ref(1);
let pageSize = ref(1);
const msg = inject("message")

const pagesize = [
  { label: "1 每页", value: 1 },
  { label: "2 每页", value: 2 },
  { label: "3 每页", value: 3 },
  { label: "4 每页", value: 4 }
];

onMounted(() => {
  loadDatas(pageSize.value, page.value);
});

const loadDatas = async (pagesize, pagenum) => {
  try {
    //console.log("Sending request with parameters: pagesize =", pagesize, "pagenum =", pagenum);
    const res = await as.get("/users", {
      params: {
        pagesize: pagesize,
        pagenum: pagenum,
      }
    });
    //console.log("Response data:", res.data);
    userlist.value = res.data.data;
    pagecount.value = Math.ceil(res.data.total / pagesize);
  } catch (error) {
    console.error("Error loading data:", error);
  }
};


const handlePageUpdate = (newPage) => {
  page.value = newPage;
  loadDatas(pageSize.value, page.value);
};

const handlePageSizeUpdate = (newPageSize) => {
  pageSize.value = newPageSize;
  // page.value=1
  loadDatas(newPageSize, 1);
};

const search_name = ref('');

const searchname = async (username) => {
  let uinfo = await as.get("/searchuser", {
    params: {
      username: username,
    },

  });
  console.log(uinfo)
  return uinfo.data;
};



const search = () => {
  if (search_name.value.trim() !== '') {
    console.log("search", search_name.value);
    let uinfoPromise = searchname(search_name.value); // 返回一个Promise对象
    uinfoPromise.then(uinfo => {

        if (uinfo.data.ID === 0) {
          msg.error("未查找到用户");
          letnil();
          return;
        }
        userlist.value = Array.isArray(uinfo.data) ? uinfo.data : [uinfo.data];
        const resultIndex = userlist.value.findIndex(user => user.username === search_name.value);
        if (resultIndex !== -1) {
          page = ref(Math.ceil((resultIndex + 1) / pageSize.value));
          console.log("定位");
        }
        return;
    }).catch(error => {
      console.error("Error occurred:", error);
    });
  } else {
    console.log("搜索内容为空，不执行查询操作");
  }
};


const letnil=()=>{
  search_name.value=''
}
</script>
<style>
</style>