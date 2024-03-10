<template>
  <div>
    <n-table :bordered="false" :single-line="false">
      <thead>
      <tr>
        <th>编号</th>
        <th>分类</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(category,index) in categorylist" :key="index">
        <td>{{category.id}}</td>
        <td>{{category.name}}</td>
      </tr>
      </tbody>
    </n-table>
  </div>
</template>
<script setup>
import {ref,onMounted} from "vue";
import as from "@/plugins/axios";
const categorylist = ref([])

onMounted(()=>{
  loadDatas()
})
const loadDatas = async ()=>{
  let res = await as.get("/category")
  categorylist.value = res.data.data
  console.log(res)
}

</script>