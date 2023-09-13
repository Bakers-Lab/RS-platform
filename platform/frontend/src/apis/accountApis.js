import {request} from "../request";
// 示例:
// export const changeNickname=(content)=>{
//     return request({url:"/account/updateClientInfo",method:"post",params:{userId:JSON.parse(localStorage.getItem("client_Id")),type:0,content:content}})
// }
export const signUp=(signUpInfo)=>{
    // return request({url:"",method:"post",data:signUpInfo})
    console.log(signUpInfo)
}
export const login=(loginInfo)=>{
    // return request({url:"",method:"post",data:loginInfo})
    console.log(loginInfo)
    return new Promise((resolve, reject)=>{
        resolve({data:{
                id:1,
                userName:"zjp",
                email:"914856774@qq.com"
            }})
    })
}