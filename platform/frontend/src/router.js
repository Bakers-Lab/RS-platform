import {createBrowserRouter} from "react-router-dom";
import LoginPage from "./accountPages/LoginPage";
import SignUpPage from "./accountPages/SignUpPage";
import HomePage from "./accountPages/HomePage";
import DatasetListPage from "./datasetPages/DatasetListPage";
import DatasetDetailPage from "./datasetPages/DatasetDetailPage";
import EvaluationListPage from "./evaluationPages/EvaluationListPage";
import EvaluationDetailPage from "./evaluationPages/EvaluationDetailPage";
export const router = createBrowserRouter([
    {
        path: "/",
        element: <LoginPage/>,
    },
    {
        path: "/signup",
        element: <SignUpPage/>,
    },
    {
        path: "/home",
        element: <HomePage/>,
        children:[
            {
                path:"/home/datasetlist",
                element:<DatasetListPage/>
            },
            {
                path: "/home/datasetdetail",
                element: <DatasetDetailPage/>
            },
            {
                path: "/home/evaluationlist",
                element: <EvaluationListPage/>
            },
            {
                path: "/home/evaluationDetail",
                element: <EvaluationDetailPage/>
            }
        ]
    }
]);