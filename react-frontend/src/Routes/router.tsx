import React from "react";
import { createBrowserRouter } from "react-router-dom";
import Dashboard from "../Layouts/Dashboard";
const App = React.lazy(() => import("../App"))
const ErrorLayout = React.lazy(() => import("../Layouts/ErrorLayout"))
const router = createBrowserRouter([
    {
        element: (
            <React.Suspense fallback={<div>Loading app…</div>}>
                <App />
            </React.Suspense>
            ),
            errorElement: (
            <React.Suspense fallback={<div>Loading error…</div>}>
                <ErrorLayout />
            </React.Suspense>
            ),
        children: [
            { index: true, element: <Dashboard /> },
        ],
    }
]);


export default router;