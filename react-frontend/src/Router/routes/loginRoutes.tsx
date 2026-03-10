import { RouteObject } from "react-router-dom";
import Login from "../../Layouts/Login";
import CreateAccount from "../../Layouts/CreateAccount";

export const loginRoutes: RouteObject[] = [
    {
        path: "login",
        element: <Login />
    },
    {
        path: "create-account",
        element: <CreateAccount/>
    }
];
