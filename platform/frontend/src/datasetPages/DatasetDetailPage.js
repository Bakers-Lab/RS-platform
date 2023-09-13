import React from "react";
import {useLocation} from "react-router-dom";
import {
    Box, Button,
    Card,
    CardBody,
    CardHeader,
    DataTable,
    Heading, Layer,
    NameValueList,
    NameValuePair,
    Paragraph,
    Text, TextInput,
    Tip,
    Form,
    FormField, FileInput, Anchor, Grommet
} from "grommet";
import {getDatasetDetail} from "../apis/datasetApis";
import {Add} from "grommet-icons";
export default function DatasetDetailPage(props){
    const myTheme={
        anchor:{
            color:"#0069aa"
        }
    }
    const location=useLocation()
    const [id,setId]=React.useState(location.state.id)
    const [datasetInfo,setDatasetInfo]=React.useState({})
    const [batches,setBatches]=React.useState([])
    const [show,setShow]=React.useState(false)
    const [showContent,setShowContent]=React.useState(false)


    const [batchName,setBatchName]=React.useState("")
    const [batchFile,setBatchFile]=React.useState(null)
    const init=()=>{
        getDatasetDetail(id).then(res=>{
            setDatasetInfo(res.data)
            setBatches(res.data.batches)
        })
    }
    React.useEffect(init,[])
    return(
        <Grommet theme={myTheme}>
            <Card height="medium" width="large" background="light-1" animation={{type:"slideDown",delay:0,duration:1000,size:"medium"}}>
                <CardBody pad={"medium"}>
                    <Heading level={2}>{datasetInfo.name}</Heading>
                    <Box direction={"row"} gap={"xlarge"}>
                        <NameValueList>
                            <NameValuePair name={"数据集序号"}>
                                <Text>{datasetInfo.id}</Text>
                            </NameValuePair>
                            <NameValuePair name={"数据集路径"}>
                                <Text>"{datasetInfo.path}"</Text>
                            </NameValuePair>
                            <NameValuePair name={"数据集存储格式"}>
                                <Text>{datasetInfo.storeFormat}</Text>
                            </NameValuePair>
                            <NameValuePair name={"数据集状态"}>
                                <Text >{datasetInfo.state}</Text>
                            </NameValuePair>
                            <NameValuePair name={"数据集描述"}>
                                <Paragraph margin={"none"} maxLines={4}>{datasetInfo.comment}</Paragraph>
                                <Anchor label={"点击查看全部"} onClick={()=>setShowContent(true)} color={"#0069aa"}/>
                                <Box>
                                    {showContent && (
                                        <Layer
                                            onClickOutside={() => {
                                                setShowContent(false)
                                            }}>
                                            <Box width={"medium"} margin={"medium"}>
                                                {datasetInfo.comment}
                                            </Box>

                                        </Layer>)}
                                </Box>
                            </NameValuePair>
                        </NameValueList>
                    </Box>
                </CardBody>
            </Card>
            <Box animation={{type:"fadeIn",delay:200,duration:1000,size:"medium"}} width={"small"}>
                <Button onClick={()=>setShow(true)} hoverIndicator primary color={"#0069aa"} size={"small"} label={"添加batch"} margin={{vertical:"medium"}}></Button>
            </Box>
            <Box animation={{type:"fadeIn",delay:400,duration:1000,size:"medium"}}>
                <DataTable
                    data={batches}
                    pad={{left:"xlarge",vertical:"xsmall"}}
                    size={"small"}
                    margin={{right:"large"}}
                    columns={[
                        {
                            property:'id',
                            header:"序号",
                            render:batch=>(<Text weight={"lighter"}>{batch.id}</Text>),
                            size:"xsmall"
                        },
                        {
                            property:'name',
                            header:"名称",
                            render:batch=>(<Text weight={"lighter"}>{batch.name}</Text>),
                            size:"xsmall"
                        },
                        {
                            property:'state',
                            header:"状态",
                            render:batch=>(<Text weight={"lighter"}>{batch.state}</Text>),
                            size:"xsmall"
                        },
                        {
                            property:'samplesNum',
                            header:"样本量",
                            render:batch=>(<Text weight={"lighter"}>{batch.samplesNum}</Text>),
                            size:"xsmall"
                        },
                        {
                            property:'createdAt',
                            header:"创建时间",
                            render:batch=>(<Text weight={"lighter"}>{batch.createdAt}</Text>),
                            size:"xsmall"
                        },
                        {
                            property:'updatedAt',
                            header:"更新时间",
                            render:batch=>(<Text weight={"lighter"}>{batch.updatedAt}</Text>),
                            size:"xsmall"
                        },
                        {
                            property:'deletedAt',
                            header:"删除时间",
                            render:batch=>(<Text weight={"lighter"}>{batch.deletedAt}</Text>),
                            size:"xsmall"
                        },
                        {
                            property:'fileSize',
                            header:"文件大小",
                            render:batch=>(<Text weight={"lighter"}>{batch.fileSize}</Text>),
                            size:"xsmall"
                        },
                    ]}
                >

                </DataTable>
            </Box>
            <Box>
                {show && (
                    <Layer
                        onClickOutside={() => {
                            setShow(false)
                            setBatchName("")
                            setBatchFile(null)
                        }}
                    >
                        <Box pad={"medium"}>
                            <FormField label={"名称"}>
                                <TextInput
                                    value={batchName}
                                    onChange={event => {setBatchName(event.target.value)}}/>
                            </FormField>
                            <FormField label={"文件"}>
                                <FileInput
                                    name={"batchFile"}
                                    onChange={event => {setBatchFile(event.target.files[0])}}/>
                            </FormField>
                            <Box direction={"row"} gap={"small"} margin={{top:"medium"}}>
                                <Button primary label={"提交"} color={"#0069aa"}/>
                                {/*todo:完成上传batch的api*/}
                                <Button label={"取消"} color={"#0069aa"} onClick={()=>{
                                    setShow(false)
                                    setBatchName("")
                                    setBatchFile(null)
                                }}/>
                            </Box>
                        </Box>
                    </Layer>
                )}
            </Box>

        </Grommet>
    )
}