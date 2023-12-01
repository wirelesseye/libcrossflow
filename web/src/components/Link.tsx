import { HTMLProps, MouseEventHandler, useCallback } from "react";
import { usePush } from "../utils/router";

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