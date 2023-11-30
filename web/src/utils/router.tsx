import {
    HTMLProps,
    MouseEventHandler,
    createElement,
    useCallback,
    useEffect,
    useMemo,
} from "react";
import * as zustand from "zustand";

export interface Route {
    pathname: string | RegExp;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    element: (...any: any[]) => JSX.Element | null;
}

interface RouterState {
    pathname: string;
    setPathname: (pathname: string) => void;
}

const useRouterStore = zustand.create<RouterState>((set) => ({
    pathname: window.location.pathname,
    setPathname: (pathname: string) => set({ pathname }),
}));

export const usePathname = () => {
    const pathname = useRouterStore((s) => s.pathname);
    return pathname;
};

export const usePush = () => {
    const setPathname = useRouterStore((s) => s.setPathname);

    const f = useCallback((url: string) => {
        setPathname(url);
        history.pushState({}, "", url);
    }, []);

    return f;
};

export const useRedirect = (url: string) => {
    const setPathname = useRouterStore((s) => s.setPathname);

    useEffect(() => {
        setPathname(url);
        history.replaceState({}, "", url);
    }, []);
};

export interface RouterProps {
    routes: Route[];
    NotFoundElement?: () => JSX.Element;
}

export function Router({ routes, NotFoundElement }: RouterProps) {
    const { pathname, setPathname } = useRouterStore();

    useEffect(() => {
        window.addEventListener("popstate", () => {
            setPathname(decodeURI(window.location.pathname));
        });
    }, []);

    const Page = useMemo(() => {
        for (const route of routes) {
            if (route.pathname instanceof RegExp) {
                const match = route.pathname.exec(pathname);
                if (match) {
                    return createElement(route.element, { params: match.groups })
                }
            } else {
                if (route.pathname === pathname) {
                    return createElement(route.element);
                }
            }
        }

        return NotFoundElement ? createElement(NotFoundElement) : null;
    }, [pathname]);

    return Page;
}

export interface LinkProps extends HTMLProps<HTMLAnchorElement> {}

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
