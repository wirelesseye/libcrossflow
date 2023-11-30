import Layout from "./Layout";
import { Route, Router } from "./utils/router";

import FilesPage from "./pages/Files";
import HomePage from "./pages/Home";
import NotFoundPage from "./pages/NotFound";
import SharespacesPage from "./pages/Sharespaces";

import "./styles.css";

const routes: Route[] = [
    {
        pathname: "/",
        element: HomePage,
    },
    {
        pathname: /^\/files\/?$/,
        element: SharespacesPage,
    },
    {
        pathname: /^\/files\/(?<filepath>.+)$/,
        element: FilesPage
    }
];

export default function App() {
    return (
        <Layout>
            <Router routes={routes} NotFoundElement={NotFoundPage} />
        </Layout>
    );
}
