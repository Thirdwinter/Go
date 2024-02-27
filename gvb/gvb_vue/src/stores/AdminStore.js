import {defineStore} from  'pinia'
export const AdminStore = defineStore("admin",{
    state:()=>{
        return {
            username:"",
            //password:"",
            token:"",
            //remember:"",
        }
    },
    actions:{},
    getters:{},
})
