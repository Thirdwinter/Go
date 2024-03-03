import {defineStore} from  'pinia'
export const AdminStore = defineStore("admin",{
    state:()=>{
        return {
            username:"",
            //password:"",
            atoken:"",
            rtoken:"",
            //remember:"",
        }
    },
    actions:{},
    getters:{},
})
