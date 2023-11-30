import { HTMLProps, MouseEventHandler, useCallback, useEffect, useMemo } from "react";
import * as zustand from "zustand";
import FilesPage from "./pages/Files";
import HomePage from "./pages/Home";
import NotFoundPage from "./pages/NotFound";
import SharespacesPage from "./pages/Sharespaces";

interface RouterState {
    pathname: string;
    setPathname: (pathname: string) => void;
}

const useRouterStore = zustand.create<RouterState>((set) => ({
    pathname: window.location.pathname,
    setPathname: (pathname: string) => set({ pathname }),
}));

export const usePathname = () => {
    const pathname = useRouterStore(s => s.pathname);
    return pathname;
}

export const usePush = () => {
    const setPathname = useRouterStore(s => s.setPathname);
    
    const f = useCallback((url: string) => {
        setPathname(url);
        history.pushState({}, "", url);
    }, []);

    return f;
}

export const useRedirect = (url: string) => {
    const setPathname = useRouterStore(s => s.setPathname);

    useEffect(() => {
        setPathname(url);
        history.replaceState({}, "", url);
    }, []);
}

export function Router() {
    const {pathname, setPathname} = useRouterStore();

    useEffect(() => {
        window.addEventListener('popstate', () => {
            setPathname(window.location.pathname);
        });
    }, []);

    const Page = useMemo(() => {
        if (pathname === "/") {
            return HomePage;
        } else if (pathname === "/files" || pathname === "/files/") {
            return SharespacesPage;
        } else if (pathname.startsWith("/files/")) {
            return FilesPage;
        } else {
            return NotFoundPage;
        }
    }, [pathname]);

    return <Page />;
}

export interface LinkProps extends HTMLProps<HTMLAnchorElement> {
}

export function Link(props: LinkProps) {
    const { onClick, href, ...other } = props;
    const push = usePush();

    const handleClick = useCallback<MouseEventHandler<HTMLAnchorElement>>(
        (e) => {
            e.preventDefault();
            if (href) {
                push(href);
            }
            if (onClick) onClick(e);
        },
        [href],
    );

    return <a onClick={handleClick} href={href} {...other} />;
}
