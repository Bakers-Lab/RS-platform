import React from "react";
import {Box, Button, Card, CardBody, CardHeader, Heading, PageHeader, TextInput} from "grommet";
import {Form} from "grommet";
import {FormField} from "grommet";
import {login} from "../apis/accountApis";
import { useNavigate } from 'react-router-dom'
export default function LoginPage(props){
    const [userInfo,setUserInfo]=React.useState({})
    const navigate=useNavigate()

    const sendLoginInfo=()=>{
        login(userInfo).then(res=>{
            console.log(res.data)
            let info=res.data
            localStorage.setItem("id",JSON.stringify(info.id))
            localStorage.setItem("userName",JSON.stringify(info.userName))
            localStorage.setItem("email",JSON.stringify(info.email))
            navigate('/home/datasetlist')
        })
    }



    return(
        <div>
            <Box
                align={"center"}
                pad={{top:"15em"}}
                animation={{type:"slideDown",duration:1000,size:"medium"}}
            >
                <Card height="auto" width="medium">
                    <CardHeader pad={"medium"} background={"#0069aa"} height={"4em"}>
                        <Heading level={3}>遥感平台</Heading>
                    </CardHeader>
                    <CardBody pad={"medium"} background={{
                        image:"url(https://img2.baidu.com/it/u=3224995901,99255141&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=507)",
                        opacity:"weak"
                    }}>
                        <Form
                            value={userInfo}
                            onChange={newInfo=>setUserInfo(newInfo)}>
                            <FormField name="email" htmlFor="text-input-email" label="邮箱">
                                <TextInput id="text-input-email" name="email" />
                            </FormField>
                            <FormField name="password" htmlFor="text-input-password" label="密码">
                                <TextInput id="text-input-password" name="password" />
                            </FormField>
                            <Box direction="row" gap="medium" margin={{top:'2em'}}>
                                <Button onClick={sendLoginInfo} color={"#0069aa"} primary label="登录" />
                                <Button onClick={()=>navigate("/signup")} color={"#0069aa"} label="注册" />
                            </Box>
                        </Form>

                    </CardBody>
                </Card>
            </Box>



        </div>
    )
}
