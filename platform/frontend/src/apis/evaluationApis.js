import {request} from "../request";
// 示例:
// export const changeNickname=(content)=>{
//     return request({url:"/account/updateClientInfo",method:"post",params:{userId:JSON.parse(localStorage.getItem("client_Id")),type:0,content:content}})
// }
export const getAllEvaluation=(id)=>{
    return new Promise((resolve, reject) => {
        resolve({
            data:[
                {
                    id:1,
                    name:"评估1",
                    state:"导入",
                },
                {
                    id:2,
                    name:"评估2",
                    state:"导入",
                },
                {
                    id:3,
                    name:"评估3",
                    state:"导入",
                },
                {
                    id:4,
                    name:"评估4",
                    state:"导入",
                },
            ]
        })
    })
}
export const getEvaluationDetail=(id)=>{
    return new Promise((resolve, reject) => {
        resolve({
            data:{
                id:id,
                name:"评估"+id,
                inferJobId:1,
                datasetId:1,
                path:"c:/data",
                fileSize:60,
                state:"导入",
                createdAt:100,
                updatedAt:100,
                deletedAt:100
            }
        })
    })
}