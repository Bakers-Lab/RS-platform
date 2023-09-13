import React from "react";
import {Outlet, useLocation, useNavigate} from "react-router-dom"
import {Box, Button, Grid, Header, Heading, Nav, Sidebar} from "grommet";
export default function HomePage(props){

    const [id,setId]=React.useState(0)
    const [userName,setUserName]=React.useState("")
    const [email,setEmail]=React.useState("")
    const location=useLocation()
    const navigate=useNavigate()

    const init=()=>{
        setId(JSON.parse(localStorage.getItem("id")))
        setUserName(JSON.parse(localStorage.getItem("userName")))
        setEmail(JSON.parse(localStorage.getItem("email")))
        console.log(location)
    }
    React.useEffect(init,[])

    const active=(button)=>{
        if(button===1&&(location.pathname==="/home/datasetlist"||location.pathname==="/home/datasetdetail")){
            return true
        }
        //todo:确定路由名称后完善这一函数,用于切换页面时改变按钮状态
        else if(button===3&&(location.pathname==="/home/evaluationlist"||location.pathname==="/home/evaluationdetail")){
            return true
        }
    }

    const jump=(button)=>{
        if (button===1){
            navigate("/home/datasetlist")
        }
        //todo:确定路由名称后完善这一函数，用于确定按钮跳转的路由
        if (button===3){
            navigate("/home/evaluationlist")
        }
    }

    return(
        <div>
            <Box>
                <Header background={"#0069aa"} height={"xsmall"} round={"small"}>
                    <Box margin={{left:"5%"}}>
                        <Heading level={1}>遥感平台</Heading>
                    </Box>
                    <Box margin={{right:"5%"}}>
                        <Heading level={5} margin={"xsmall"}>{userName}</Heading>
                        <Heading level={5} margin={"xsmall"}>{email}</Heading>
                    </Box>
                </Header>
            </Box>
            <Box margin={{top:"1em"}} direction={"row"} >
                    <Sidebar
                        animation={{type:"fadeIn",duration:1000,size:"medium"}}
                             background={{
                        image:"url(https://img2.baidu.com/it/u=3224995901,99255141&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=507)",
                        opacity:"weak"}}
                             round={"small"}
                             height={"large"}
                             width={"10%"}
                        border={{color:"#0069aa"}}>
                        <Nav direction={"column"} margin={{top:"30%"}}>
                            <Button style={{borderRadius:"5px"}} color={"#0069aa"} primary label={"数据集"} hoverIndicator={active(1)} active={active(1)} onClick={()=>jump(1)}/>
                            <Button style={{borderRadius:"5px"}} color={"#0069aa"} primary label={"预测"} hoverIndicator active={false}/>
                            <Button style={{borderRadius:"5px"}} color={"#0069aa"} primary label={"评估"} hoverIndicator={active(3)} active={active(3)} onClick={()=>jump(3)}/>
                        </Nav>
                        {/*todo:仿照上面添加路由跳转*/}
                    </Sidebar>
                    <Box margin={{left:"large"}} width={"90%"}>
                        <Outlet/>
                    </Box>
            </Box>

        </div>
    )
}