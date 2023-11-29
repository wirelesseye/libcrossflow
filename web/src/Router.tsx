import { HTMLProps, MouseEventHandler, useCallback, useEffect, useMemo } from "react";
import FilesPage from "./pages/Files";
import HomePage from "./pages/Home";
import NotFoundPage from "./pages/NotFound";
import * as zustand from "zustand";

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
        } else if (pathname.startsWith("/files")) {
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
    const setPathname = useRouterStore((s) => s.setPathname);

    const handleClick = useCallback<MouseEventHandler<HTMLAnchorElement>>(
        (e) => {
            e.preventDefault();
            if (href) {
                setPathname(href);
                history.pushState({}, "", href);
            }
            if (onClick) onClick(e);
        },
        [href],
    );

    return <a onClick={handleClick} href={href} {...other} />;
}
