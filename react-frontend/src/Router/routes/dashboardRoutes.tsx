import { RouteObject } from "react-router-dom";
import Dashboard from "../../Layouts/Dashboard";

export const dashboardRoutes: RouteObject[] = [
    {
        index: true,
        element: <Dashboard />
    },
    {
        path: "settings",
        element: <div>Settings</div>
    }
];
