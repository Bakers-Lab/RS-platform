import React from "react";
import {getAllEvaluation} from "../apis/evaluationApis";
import {useEffect} from "react";
import {
    Box,
    Button,
    DataTable,
    FileInput,
    Grommet,
    Layer,
    RadioButtonGroup,
    Select,
    Text,
    TextInput,
    WorldMap
} from "grommet";
import {Add} from "grommet-icons";
import { useNavigate} from "react-router-dom";
import {Form} from "grommet";
import {FormField} from "grommet";
import {getAllDataset} from "../apis/datasetApis";
export default function EvaluationListPage(props){
    const myTheme={
        select:{
            icons:{
                color:"#0069aa"
            }
        },
        anchor:{
            color:"#0069aa"
        }
    }
    const [evaluationList,setEvaluationList]=React.useState([])
    const [show,setShow]=React.useState(false)
    const inidata={
        name:"",
        inferJobId:-1,
        datasetId:-1,
        createrUserId:JSON.parse(localStorage.getItem("id")),//todo:提醒一下加一下createrUserId
    }
    const [file,setFile]=React.useState(null)
    const [newData,setNewData]=React.useState(inidata)
    const [myInferJob,setMyInferJob]=React.useState([])
    const [myDataset,setMyDataset]=React.useState([])
    const init=()=>{
        getAllEvaluation(JSON.parse(localStorage.getItem("id"))).then(res=>{
            setEvaluationList(res.data)
        })
        getAllDataset(JSON.parse(localStorage.getItem("id"))).then(res=>{
            let datasetIds=[]
            for (let i = 0; i < res.data.length; i++) {
                datasetIds.push(res.data[i].id)
            }
            setMyDataset(datasetIds)
        })
        setMyInferJob([2,3])//todo:预测写完后补全
    }

    useEffect(init,[])

    const navigate=useNavigate()

    const jump=(data)=>{
        navigate("/home/evaluationdetail",{state:{id:data.id}})
    }

    return(
        <Grommet theme={myTheme}>
            <Box  animation={{type:"fadeIn",delay:200,duration:1000,size:"medium"}} width={"xxsmall"}>
                <Button hoverIndicator primary color={"#0069aa"} size={"small"} icon={<Add size={"medium"}/>} onClick={()=>setShow(true)}/>
            </Box>
            <Box animation={{type:"fadeIn",delay:200,duration:1000,size:"medium"}}>
                <DataTable
                    pad={{left:"medium",vertical:"xsmall"}}
                    onClickRow={(e)=>jump(e.datum)}
                    margin={{right:"large"}}
                    columns={[
                        {
                            property:'id',
                            header:"序号",
                            render:evaluation=>(<Text weight={"lighter"}>{evaluation.id}</Text>),
                            size:"xsmall"
                        },
                        {
                            property:"name",
                            header:"名称",
                            render:evaluation=><Text truncate weight={"lighter"}>{evaluation.name}</Text>,
                            size:"xsmall",
                        },
                        {
                            property:"state",
                            header:"状态",
                            render:evaluation=><Text weight={"lighter"}>{evaluation.state}</Text>,
                            size:"xsmall"
                        },
                    ]}
                    data={evaluationList}
                    size={"small"}/>
            </Box>
            <Box animation={{type:"fadeIn",delay:400,duration:1000,size:"medium"}}>
                <WorldMap
                    margin={{top:"3%"}}
                    color="#a4b8c9"
                    continents={[
                        {
                            name: 'Asia',
                            color: "#0069aa",
                            onClick: (name) => {},
                        },
                    ]}
                    onSelectPlace={(lat, lon) => {}}
                    selectColor="brand"
                />
            </Box>
            <Box>
                {show &&(
                    <Layer
                        onClickOutside={() => {
                            setShow(false)
                            setNewData(inidata)
                            setFile(null)
                        }}>
                        <Box pad={"medium"}>
                            <Form
                                value={newData}
                                onChange={nextValue=>setNewData(nextValue)}>
                                <FormField name={"name"} htmlFor={"newtext-input-name"} label={"名称"}>
                                    <TextInput id={"newtext-input-name"} name={"name"}/>
                                </FormField>
                                <FormField name={"inferJobId"} htmlFor={"newtext-input-inferJobId"} label={"预测任务"}>
                                    <Select options={myInferJob} name={"inferJobId"}/>
                                </FormField>
                                <FormField name={"datasetId"} htmlFor={"newtext-input-datasetId"} label={"数据集"}>
                                    <Select options={myDataset} name={"datasetId"}/>
                                </FormField>
                                <FormField label={"评估相关文件"}>
                                    <FileInput onChange={event => setFile(event.target.files[0])}/>
                                </FormField>
                            </Form>
                            <Box direction={"row"} gap={"small"} margin={{top:"medium"}}>
                                <Button primary label={"提交"} color={"#0069aa"}/>
                                {/*todo:完成创建评估的api*/}
                                <Button label={"取消"} color={"#0069aa"} onClick={()=>{
                                    setShow(false)
                                    setNewData(inidata)
                                    setFile(null)
                                }}/>
                            </Box>
                        </Box>
                    </Layer>
                )}
            </Box>
        </Grommet>

    )
}