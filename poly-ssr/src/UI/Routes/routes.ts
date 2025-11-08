import { createBrowserRouter } from "react-router-dom";
import { dashboardRoutes } from "./dashboardRoutes.tsx";
const routes = [
    dashboardRoutes
];
const router = createBrowserRouter(routes);

export default router;