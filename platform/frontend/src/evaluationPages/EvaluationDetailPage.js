import React from "react";
import {useLocation} from "react-router-dom";
import {useNavigate} from "react-router-dom";
import {getEvaluationDetail} from "../apis/evaluationApis";
import {Anchor, Card, CardBody, Heading, NameValueList, NameValuePair, Text} from "grommet";
export default function EvaluationDetailPage(props){
    const location=useLocation()
    const navigate=useNavigate()
    const [id,setId]=React.useState(location.state.id)
    const [evaluationData,setEvaluationData]=React.useState({})

    const init=()=>{
        getEvaluationDetail(id).then(res=>{
            setEvaluationData(res.data)
        })
    }

    const jumpToInferJob=()=>{
        //todo:写完预测页面后添加
    }

    const jumpToDataset=()=>{
        navigate("/home/datasetdetail",{state:{id:evaluationData.datasetId}})
    }

    React.useEffect(init,[])
    return(
        <div>
            <Card height="auto" width="large" background="light-1" animation={{type:"slideDown",delay:0,duration:1000,size:"medium"}}>
                <CardBody pad={"medium"}>
                    <Heading level={2}>{evaluationData.name}</Heading>
                    <NameValueList>
                        <NameValuePair name={"评估序号"}>
                            <Text>{evaluationData.id}</Text>
                        </NameValuePair>
                        <NameValuePair name={"预测任务序号"}>
                            <Anchor label={evaluationData.inferJobId} color={"#0069aa"}/>
                        </NameValuePair>
                        <NameValuePair name={"数据集序号"}>
                            <Anchor label={evaluationData.datasetId} color={"#0069aa"} onClick={()=>jumpToDataset()}/>
                        </NameValuePair>
                        <NameValuePair name={"评估路径"}>
                            <Text>{evaluationData.path}</Text>
                        </NameValuePair>
                        <NameValuePair name={"评估文件大小"}>
                            <Text>{evaluationData.fileSize}</Text>
                        </NameValuePair>
                        <NameValuePair name={"评估状态"}>
                            <Text>{evaluationData.state}</Text>
                        </NameValuePair>
                        <NameValuePair name={"创建时间"}>
                            <Text>{evaluationData.createdAt}</Text>
                        </NameValuePair>
                        <NameValuePair name={"更新时间"}>
                            <Text>{evaluationData.updatedAt}</Text>
                        </NameValuePair>
                        <NameValuePair name={"删除时间"}>
                            <Text>{evaluationData.deletedAt}</Text>
                        </NameValuePair>
                    </NameValueList>
                </CardBody>
            </Card>
        </div>
    )
}