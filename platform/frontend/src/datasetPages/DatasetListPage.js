import React from "react";
import {Box, Button, DataTable, Layer, Select, Text, TextArea, TextInput, WorldMap} from "grommet";
import {getAllDataset} from "../apis/datasetApis";
import { useNavigate} from "react-router-dom";
import {Form} from "grommet";
import {FormField} from "grommet";
import {Add} from "grommet-icons";
export default function DatasetListPage(props){
    const [datasets,setDatasets]=React.useState([])
    const [show,setShow]=React.useState(false)
    const inidata={
        name:"",
        comment:"",
        id:JSON.parse(localStorage.getItem("id")),//todo:改一下
        storeFormat:""
    }
    const [newData,setNewData]=React.useState(inidata)
    const init=()=>{
        getAllDataset(JSON.parse(localStorage.getItem("id"))).then(res=>{
            setDatasets(res.data)
        })
    }
    React.useEffect(init,[])
    const navigate=useNavigate()
    const jump=(data)=>{
        navigate("/home/datasetdetail",{state:{id:data.id}})
    }
    return(
        <div>
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
                            render:datasets=>(<Text weight={"lighter"}>{datasets.id}</Text>),
                            size:"xsmall"
                        },
                        {
                            property:"name",
                            header:"名称",
                            render:datasets=><Text truncate weight={"lighter"}>{datasets.name}</Text>,
                            size:"xsmall",
                        },
                        {
                            property:"state",
                            header:"状态",
                            render:datasets=><Text weight={"lighter"}>{datasets.state}</Text>,
                            size:"xsmall"
                        },
                        {
                            property:"comment",
                            header:"描述",
                            render:datasets=><Text weight={"lighter"} truncate>{datasets.comment}</Text>,
                            size:"medium"
                        }
                    ]}
                    data={datasets}
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
                {
                    show && (
                        <Layer
                            onClickOutside={() => {
                                setShow(false)
                                setNewData(inidata)
                            }}
                        >
                            <Box pad={"medium"}>
                                <Form
                                value={newData}
                                onChange={nextValue=>setNewData(nextValue)}
                                >
                                    <FormField name={"name"} htmlFor={"newtext-input-name"} label={"名称"}>
                                        <TextInput id={"newtext-input-name"} name={"name"}/>
                                    </FormField>
                                    <FormField name={"storeFormat"} htmlFor={"newtext-input-storeFormat"} label={"文件类型"}>
                                        <Select options={['.npy','.png','.tif']} id={"newtext-input-storeFormat"} name={"storeFormat"}/>
                                    </FormField>
                                    <FormField name={"comment"} htmlFor={"newtext-input-comment"} label={"描述"}>
                                        <TextArea id={"newtext-input-comment"} name={"comment"} resize={"vertical"}/>
                                    </FormField>
                                </Form>
                                <Box direction={"row"} gap={"small"} margin={{top:"medium"}}>
                                    <Button primary label={"提交"} color={"#0069aa"} />
                                    {/*todo:完成创建数据集的api*/}
                                    <Button label={"取消"} color={"#0069aa"} onClick={()=>{
                                        setShow(false)
                                        setNewData(inidata)
                                    }}/>
                                </Box>
                            </Box>
                        </Layer>
                    )
                }
            </Box>
        </div>
    )
}