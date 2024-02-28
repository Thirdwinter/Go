<template>
  <div style="padding: 5px">
    <n-input placeholder="搜索" v-model:value="search_name">
      <template #prefix>
        <n-icon :component="FlashOutline" />
      </template>
    </n-input>
    <n-button type="info" @click="search">
      让我康康!
    </n-button>
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

const axios = inject("axios");
const userlist = ref([]);
const pagecount = ref(0);
const page = ref(1);
const pageSize = ref(1);

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
    console.log("Sending request with parameters: pagesize =", pagesize, "pagenum =", pagenum);
    const res = await axios.get("/users", {
      params: {
        pagesize: pagesize,
        pagenum: pagenum,
      }
    });
    console.log("Response data:", res.data);
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
  page.value=1
  loadDatas(newPageSize, 1);
};

const search_name = ref('');

const searchname = async (username) => {
  // 从本地存储获取JWT
  const jwt = localStorage.getItem('token');

  const uinfo = await axios.get("/searchuser", {
    params: {
      username: username,
    },
    headers: {
      Authorization: `Bearer ${jwt}` // 设置Authorization头部为Bearer加上JWT
    }
  });

  return uinfo.data; // 返回获取的数据
}


  const search = () => {
    console.log("search", search_name.value);
    const uinfoPromise = searchname(search_name.value); // 返回一个Promise对象

    uinfoPromise.then(uinfo => {
      console.log(uinfo);
      console.log("ok");
    }).catch(error => {
      console.error("Error:", error);
    });
  };
</script>
