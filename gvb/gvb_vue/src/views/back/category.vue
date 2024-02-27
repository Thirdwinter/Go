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
import {ref,onMounted,inject} from "vue";

const axios = inject("axios")
const categorylist = ref([])

onMounted(()=>{
  loadDatas()
})
const loadDatas = async ()=>{
  let res = await axios.get("/category")
  categorylist.value = res.data.data
  console.log(res)
}

</script>