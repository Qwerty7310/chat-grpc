import {createBrowserRouter} from "react-router-dom";
import {LoginPage} from "./pages/LoginPage.tsx";
import {ChatPage} from "./pages/ChatPage.tsx"

export const router = createBrowserRouter([
    {
        path: "/",
        element: <LoginPage/>,
    },
    {
        path: "/chat",
        element: <ChatPage/>,
    }
]);