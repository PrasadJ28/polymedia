import { createBrowserRouter, RouteObject } from "react-router-dom";
import { lazyLoad } from "./lazyLoad";
import { dashboardRoutes } from "./routes/dashboardRoutes";
import { loginRoutes } from "./routes/loginRoutes";

// Using the helper
const App = lazyLoad(() => import("../App"));
const ErrorLayout = lazyLoad(() => import("../Layouts/ErrorLayout"));

const routes: RouteObject[] = [
    {
        element: <App />,
        errorElement: <ErrorLayout />,
        children: [
            ...dashboardRoutes,
            ...loginRoutes,
            // You can add more arrays here easily
        ],
    }
];

const router = createBrowserRouter(routes);

export default router;
