import React from "react";
import {Box, Button, Card, CardBody, CardHeader, Form, FormField, Heading, TextInput} from "grommet";
import {useNavigate} from "react-router-dom";
import {signUp} from "../apis/accountApis";

export default function SignUpPage(props){

    const [userInfo,setUserInfo]=React.useState({})

    const sendSignUpInfo=()=>{
        signUp(userInfo)
    }

    const navigate=useNavigate()

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
                            <FormField name="userName" htmlFor="text-input-userName" label="昵称">
                                <TextInput id="text-input-name" name="userName" />
                            </FormField>
                            <FormField name="password" htmlFor="text-input-password" label="密码">
                                <TextInput id="text-input-password" name="password" />
                            </FormField>
                            <Box direction="row" gap="medium" margin={{top:'2em'}}>
                                <Button onClick={sendSignUpInfo} color={"#0069aa"} primary label="注册" />
                                <Button onClick={()=>navigate("/")} color={"#0069aa"} label="返回" />
                            </Box>
                        </Form>

                    </CardBody>
                </Card>
            </Box>



        </div>
    )
}